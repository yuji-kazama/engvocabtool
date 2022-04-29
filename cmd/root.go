/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)



// rootCmd represents the base command when called without any subcommands
// var rootCmd = &cobra.Command{
// 	Use:   "notion-wordsapi-test",
// 	Short: "A brief description",
// 	Long: `A longer description`,
// 	// Uncomment the following line if your bare application
// 	// has an action associated with it:
// 	// Run: func(cmd *cobra.Command, args []string) { },
// }

func NewRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "notion-wordsapi-test",
		Short: "A brief description",
		Long: `A longer description`,
	}
	cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	cmd.AddCommand(NewAddCmd())
	return cmd
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	// err := rootCmd.Execute()
	// if err != nil {
	// 	os.Exit(1)
	// }
	cmd := NewRootCmd()
	cmd.SetOutput(os.Stdout)
	if err := cmd.Execute(); err != nil {
		cmd.SetOutput(os.Stderr)
		cmd.Println(err)
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.notion-wordsapi-test.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}


