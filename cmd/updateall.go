/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"engvocabtool/notion"
	"engvocabtool/words"
	"fmt"

	"github.com/spf13/cobra"
)

// updateallCmd represents the updateall command
var updateallCmd = &cobra.Command{
	Use:   "updateall",
	Short: "Update all words",
	Long: `Update all words`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return updateall(args)
	},
}

func updateall(args []string) error {

	wc := words.NewClient()
	nc := notion.NewClient()

	hasmore := true
	cursor := ""
	count := 1
	loop:
	if hasmore {
		sr, err := nc.GetAllPages(cursor)
		if err != nil {
			return err
		}
		for i, s := range sr.Results {
			word := s.Properties.Name.Title[0].Text.Content
			ar, err := wc.GetEverything(word)
			if ar != nil && err == nil {
				json := createJson(ar)
				nc.UpdatePage(sr.Results[i].ID, json)
			}
			fmt.Printf("%v : %v \n", count, word)
			count++
		}
		hasmore = sr.HasMore
		cursor = sr.NextCursor
		goto loop
	}

	return nil
}

func init() {
	rootCmd.AddCommand(updateallCmd)
}
