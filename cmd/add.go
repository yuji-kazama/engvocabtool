package cmd

import (
	"engvocabtool/notion"
	"engvocabtool/words"
	"fmt"
	"os"
	"strconv"
	"time"

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

	nc := notion.NewClient()
	if nc.Exist(word) {
		return fmt.Errorf("input word alredy exists")
	}

	wc := words.NewClient()
	res, err := wc.GetEverything(word)
	if err != nil {
		return err
	}

	index, err := showSelectPrompt(res)
	if err != nil {
		return err
	}

	json := createJsonForAdd(index, res)
	pr, err := nc.CreatePage(json)
	if err != nil {
		return err
	}
	fmt.Println(pr.URL)
	return nil
}



func getToday() string {
	return time.Now().String()[0:10]
}

func createJsonForAdd(index int, res *words.Response) string {
	name := res.Word
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
		"parent": {
			"database_id": "` + os.Getenv("NOTION_DATABASE_ID") + `"
		},
		"properties": {
			"Name": {
				"title": [
					{
						"text": {
							"content": "` + name + `" 
						}
					}
				]
			},
			"Status": {
				"select": {
					"name": "1: New"
				}
			},
			"Check Num": {
				"select": {
					"name": "1"
				}
			},
			"Study Date": {
				"date": {
					"start": "` + getToday() + `",
					"end": null,
					"time_zone": null
				}
			},
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
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
