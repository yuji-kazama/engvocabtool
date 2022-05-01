package cmd

import (
	"engvocabtool/words"
	"fmt"
	"os"
	"strconv"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
		Use:   "engvocabtool",
		Short: "A brief description",
		Long:  `A longer description`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}


func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.engvocabtool.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func showInputPrompt() (string, error) {
	prompt := promptui.Prompt{
		Label: "Input Word",
	}
	return prompt.Run()
}

func showSelectPrompt(res *words.AllResults) (int, error) {
	var items []string
	for i, s := range res.Results {
		items = append(items, strconv.Itoa(i+1)+". "+"["+s.PartOfSpeech+"] "+s.Definition)
	}
	prompt := promptui.Select{
		Label: "Select Definition",
		Items: items,
	}
	index, _, err := prompt.Run()
	return index, err
}
