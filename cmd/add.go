/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"notion-wordsapi-test/words"

	"github.com/spf13/cobra"
)

func NewAddCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add",
		Short: "Add a word to Notion database",
		Long: `Add a word to Notion database.`,
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			word := args[0]
			cw := words.NewClient()

			res, err := cw.GetEverything(word)
			if err != nil {
				return err
			}
		cmd.Printf("SUCCESS: %v is added", res.Word)
			return nil
		},
		SilenceErrors: true,
		SilenceUsage: true,
	}
	return cmd
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
