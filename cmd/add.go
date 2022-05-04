package cmd

import (
	"context"
	"engvocabtool/words"
	"fmt"
	"os"
	"time"

	"github.com/jomei/notionapi"
	"github.com/spf13/cobra"
)

var partOfSpeechToClass = map[string]string{
	"noun":        "Noun",
	"adjective":   "Adjective",
	"adverb":      "Adverb",
	"verb":        "Verb",
	"conjunction": "Conjunction",
}

var addCmd = &cobra.Command{
		Use:   "add",
		Short: "Add a word to Notion database",
		Long:  `Add a word to Notion database. The information of the word such as meanings, synonyms, examples and etc is gotten from WordsAPI.`,
		Args:  cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return add(args)
		},
		SilenceErrors: true,
		SilenceUsage:  true,
}

func add(args []string) error {
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
	if len(page.Results) > 0 {
		return fmt.Errorf("input word alredy exists")
	}

	wc := words.NewClient()
	wres, err := wc.GetEverything(word)
	if err != nil {
		return err
	}

	index, err := showSelectPrompt(wres)
	if err != nil {
		return err
	}

	pcr := createPageCreateRequest(wres, index)
	np, err := nc.Page.Create(context.Background(), pcr)
	if err != nil {
		return err
	}
	fmt.Println(np.URL)
	return nil
}

func createPageCreateRequest(wres *words.Response, index int) *notionapi.PageCreateRequest {
	example := getExample(wres.Results[index].Examples)
	synonym := getSynonum(wres.Results[index].Synonyms)

	dateObj := notionapi.Date(time.Now())
	pcr := &notionapi.PageCreateRequest{
		Parent: notionapi.Parent{
			Type:       notionapi.ParentTypeDatabaseID,
			DatabaseID: notionapi.DatabaseID(os.Getenv("NOTION_DATABASE_ID")),
		},
		Properties: notionapi.Properties{
			"Name": notionapi.TitleProperty{
				Title: []notionapi.RichText{
					{Text: notionapi.Text{Content: wres.Word}},
				},
			},
			"Status": notionapi.SelectProperty{
				Select: notionapi.Option{
					Name: "1: New",
				},
			},
			"Check Num": notionapi.SelectProperty{
				Select: notionapi.Option{
					Name: "1",
				},
			},
			"Study Date": notionapi.DateProperty{
				Date: notionapi.DateObject{
					Start: &dateObj,
				},
			},
			"Class": notionapi.SelectProperty{
				Select: notionapi.Option{
					Name: partOfSpeechToClass[wres.Results[index].PartOfSpeech],
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
		},
	}
	return pcr
}

func getSynonum(synonyms []string) string {
	if len(synonyms) < 1 {
		return ""
	}
	var synonym string
	for i, s := range synonyms {
		if i == 0 {
			synonym = s
		} else {
			synonym = synonym + ", " + s
		}
	}
	return synonym
}

func getExample(examples []string) string {
	if len(examples) < 1 {
		return ""
	}
	return examples[0]
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


func init() {
	rootCmd.AddCommand(addCmd)
}
