/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"notion-wordsapi-test/notion"
	"notion-wordsapi-test/words"
	"os"
	"strconv"
	"time"

	"github.com/spf13/cobra"
)

var partOfSpeechToClass = map[string]string {
	"noun": "Noun",
	"adjective": "Adjective",
	"adverb": "Adverb",
	"verb": "Verb",
	"conjunction": "Conjunction",
}

func NewAddCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add",
		Short: "Add a word to Notion database",
		Long: `Add a word to Notion database.`,
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			word := args[0]
			wc := words.NewClient()
			res, err := wc.GetEverything(word)
			if err != nil {
				return err
			}

			nc := notion.NewClient()
			json := createPostJson(res)

			fmt.Println(json)

			if err := nc.PostPage(json); err != nil {
				return err
			}

			cmd.Println("Success: word has been added")
			return nil
		},
	}
	return cmd
}

func getToday() string {
	return time.Now().String()[0:10] // e.g. 2022-04-09
}

func createPostJson(res *words.AllResults) (string) {
	name := res.Word
	frequency := strconv.FormatFloat(res.Frequency, 'f', -1, 64)
	examples := res.Results[0].Examples
	synonyms := res.Results[0].Synonyms
	meaning := res.Results[0].Definition
	class := partOfSpeechToClass[res.Results[0].PartOfSpeech]

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
			if i == 1 {
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
	// rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
