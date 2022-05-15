package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search a word from WordsAPI",
	Long: `Search and display a word from WordsAPI (not Notion database)`,
	Args: cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return search(cmd, args)
	},
	SilenceErrors: true,
	SilenceUsage: true,
}

func init() {
	rootCmd.AddCommand(searchCmd)
}

func search(cmd *cobra.Command, args []string) error {
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

	wdef, err := getWordDefinition(word)
	if err != nil {
		return err
	}

	index, err := showSelectPrompt(wdef)
	if err != nil {
		return err
	}

	fmt.Printf("- Definition: %v\n", wdef.Results[index].Definition)
	fmt.Printf("- Examples: %v\n", wdef.Results[index].Examples)
	fmt.Printf("- Frequency: %v\n", wdef.Frequency)
	fmt.Printf("- Pronounciation: %v\n", getPronunciation(wdef, index))
	fmt.Printf("- Synonyms: %v\n", wdef.Results[index].Synonyms)

	return nil
}
