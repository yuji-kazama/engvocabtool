package cmd

import (
	"engvocabtool/notion"
	"engvocabtool/words"
	"fmt"
	"strconv"

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
		SilenceUsage: true,
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

	index, err := showSelectPrompt(ar)
	if err != nil {
		return err
	}

	json := createJsonForUpdate(index, ar)
	up, err := nc.UpdatePage(sr.Results[0].ID, json)
	if err != nil {
		return err
	}
	fmt.Println(up.URL)
	return nil
}

func createJsonForUpdate(index int, res *words.AllResults) string {
	frequency := strconv.FormatFloat(res.Frequency, 'f', -1, 64)
	examples := res.Results[index].Examples
	synonyms := res.Results[index].Synonyms
	meaning := res.Results[index].Definition
	class := partOfSpeechToClass[res.Results[index].PartOfSpeech]

	if class == "" {
		class = "Unknown"
	}

	var example string
	if len(examples) < 1 {
		example = ""
	} else {
		example = examples[0]
	}

	var synonym string
	if len(synonyms) < 1 {
		synonym = ""
	} else {
		for i, s := range synonyms {
			if i == 0 {
				synonym = s
			} else {
				synonym = synonym + ", " + s
			}
		}
	}

	json := `{
		"properties": {
			"Class": {
				"select": {
					"name": "` + class + `"
				}
			},
			"Frequency": {
				"number": ` + frequency + `
			},
			"Meaning": {
				"rich_text": [
					{
						"text": {
							"content": "` + meaning + `"
						}
					}
				]
			},
			"Sentence": {
				"rich_text": [
					{
						"text": {
							"content": "` + example + `" 
						}
					}
				]
			},
			"Synonyms": {
				"rich_text": [
					{
						"text": {
							"content": "` + synonym + `" 
						}
					}
				]
			}
		} 
	}`
	return json
}


func init() {
	rootCmd.AddCommand(updateCmd)
}
