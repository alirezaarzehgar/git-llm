/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/alirezaarzehgar/git-llm/internal/commit"
	"github.com/spf13/cobra"
)

// commitCmd represents the commit command
var commitCmd = &cobra.Command{
	Use:   "commit",
	Short: "generate commit message and open editor",
	Long: `generate commit message using configured LLM and
open $EDITOR to change it. Then commit that message.
For example:

EDITOR=vim git llm commit`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := commit.Generate(); err != nil {
			fmt.Fprintln(os.Stderr, "failed to generate commit message:", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(commitCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// commitCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// commitCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
