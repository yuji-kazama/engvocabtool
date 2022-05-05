package cmd

import (
	"context"
	"engvocabtool/words"
	"fmt"
	"os"

	"github.com/jomei/notionapi"
	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
		Use:   "update",
		Short: "Update a word of Notion database",
		Long: `Update a word of Notion database`,
		Args: cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return update(args)
		},
		SilenceErrors: true,
		SilenceUsage:  true,
}

func update(args []string) error {
	var word string
	var err error
	if len(args) == 0 {
		word, err = showInputPrompt()
		if err != nil {
			return err
		}
	} else {
		word = args[0]
	}

	nc := notionapi.NewClient(notionapi.Token(os.Getenv("NOTION_INTEGRATION_TOKEN")))
	page, err := getPage(word, nc)
	if err != nil {
		return err
	}
	if len(page.Results) == 0 {
		return fmt.Errorf("input word does not exist")
	}

	wdef, err := getWordDefinition(word)
	if err != nil {
		return err
	}

	index, err := showSelectPrompt(wdef)
	if err != nil {
		return err
	}

	pur := getPageUpdateRequest(wdef, index)
	up, err := nc.Page.Update(context.Background(), notionapi.PageID(page.Results[0].ID), pur)
	if err != nil {
		return err
	}

	fmt.Println(up.URL)
	return nil
}

func getPage(word string, nc *notionapi.Client) (*notionapi.DatabaseQueryResponse, error) {
	sr := &notionapi.DatabaseQueryRequest{
		PropertyFilter: &notionapi.PropertyFilter{
			Property: "Name",
			RichText: &notionapi.TextFilterCondition{
				Equals: word,
			},
		},
	}
	return nc.Database.Query(
		context.Background(), notionapi.DatabaseID(os.Getenv("NOTION_DATABASE_ID")), sr)
}

func getPageUpdateRequest(wres *words.Response, index int) *notionapi.PageUpdateRequest {
	example := getExample(wres.Results[index].Examples)
	synonym := getSynonym(wres.Results[index].Synonyms)
	pronunciation := getPronunciation(wres, index)

	pur := &notionapi.PageUpdateRequest{
		Properties: notionapi.Properties{
			"Class": notionapi.SelectProperty{
				Select: notionapi.Option{
					Name: wres.Results[index].PartOfSpeech,
				},
			},
			"Frequency": notionapi.NumberProperty{
				Number: wres.Frequency,
			},
			"Meaning": notionapi.RichTextProperty{
				RichText: []notionapi.RichText{
					{Text: notionapi.Text{Content: wres.Results[index].Definition}},
				},
			},
			"Sentence": notionapi.RichTextProperty{
				RichText: []notionapi.RichText{
					{Text: notionapi.Text{Content: example}},
				},
			},
			"Synonyms": notionapi.RichTextProperty{
				RichText: []notionapi.RichText{
					{Text: notionapi.Text{Content: synonym}},
				},
			},
			"Pronunciation": notionapi.RichTextProperty{
				RichText: []notionapi.RichText{
					{Text: notionapi.Text{Content: pronunciation}},
				},
			},
		},
	}
	return pur
}

func init() {
	rootCmd.AddCommand(updateCmd)
}
