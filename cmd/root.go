/*
Copyright © 2023 Anindya Chowdhury <anindya.chowdhury@protonmail.com>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"oni-commit/utils/console"
	"oni-commit/utils/editor"
	"oni-commit/utils/gitcommand"
	"oni-commit/utils/promptwrapper"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "oni-commit",
	Short: "Commit Carefully",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		commitMessage := ""

		var prompt = promptwrapper.SelectPrompt()

		index, _, err := prompt.Run()

		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}

		selected := promptwrapper.Scopes[index]
		console.PrintSelection("selected", selected.Name)

		var typePrompt = promptwrapper.TypePrompt(selected)

		message, err := typePrompt.Run()

		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}

		console.PrintSelection("message", message)

		commitMessage = selected.Type + ": " + message

		var typeBodyPrompt = promptwrapper.TypeBodyPrompt()

		result, err := typeBodyPrompt.Run()

		addBody := true

		if err != nil || result != "y" {
			addBody = false
		}

		bodyMessage := ""

		if addBody {

			inputData, err := editor.TakeBodyAsInput()

			if err != nil {
				return
			}
			bodyMessage = (string(inputData))
		}

		gitcommand.CommitIt(commitMessage, bodyMessage)

	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.oni-commit.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
