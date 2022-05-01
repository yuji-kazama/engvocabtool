package cmd

import (
	"engvocabtool/notion"
	"engvocabtool/words"
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

func NewUpdateCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update",
		Short: "Update a word of Notion database",
		Long: `Update a word of Notion database`,
		Args: cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return update(args)
		},
		SilenceErrors: true,
		SilenceUsage: true,
	}
	return cmd
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

	nc := notion.NewClient()
	if !nc.Exist(word) {
		return fmt.Errorf("input word does not exist")
	}

	sr, err := notion.NewClient().GetPageByName(word)
	if err != nil {
		return err
	}

	wc := words.NewClient()
	ar, err := wc.GetEverything(word)
	if err != nil {
		return err
	}

	json := createJson(ar, sr)
	up, err := nc.UpdatePage(sr.Results[0].ID, json)
	if err != nil {
		return err
	}
	fmt.Println(up.URL)
	return nil
}

func createJson(wr *words.AllResults, nr *notion.SearchResult) string {
	frequency := strconv.FormatFloat(wr.Frequency, 'f', -1, 64)
	json := `{
		"parent": {
			"database_id": "` + os.Getenv("NOTION_DATABASE_ID") + `"
		},
		"properties": {
			"Frequency": {
				"number": ` + frequency + `
			}
		} 
	}`
	return json
}

func init() {
}
