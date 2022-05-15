package cmd

import (
	"context"
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

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}


func init() {
}

func showInputPrompt() (string, error) {
	prompt := promptui.Prompt{
		Label: "Input Word",
	}
	return prompt.Run()
}

func showSelectPrompt(res *words.Response) (int, error) {
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

func getWordDefinition(word string) (*words.Response, error) {
	wc := words.NewClient()
	wres, err := wc.GetEverything(context.Background(), word)
	if err != nil {
		return nil, err
	}
	return wres, nil
}

func getSynonym(synonyms []string) string {
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

func getPronunciation(wres *words.Response, index int) string {
	var pronunciation string
	if wres.Pronunciation.All != "" {
		pronunciation = wres.Pronunciation.All
	} else {
		switch wres.Results[index].PartOfSpeech {
		case "noun":
			pronunciation = wres.Pronunciation.Noun
		case "verb":
			pronunciation = wres.Pronunciation.Verb
		case "adjective":
			pronunciation = wres.Pronunciation.Adjective
		case "adverb":
			pronunciation = wres.Pronunciation.Adverb
		case "conjunction":
			pronunciation = wres.Pronunciation.Conjunction
		default:
			pronunciation = wres.Pronunciation.All
		}
	}
	return pronunciation
}

func getExample(examples []string) string {
	if len(examples) < 1 {
		return ""
	}
	return examples[0]
}
