/*
Copyright © 2025 Alireza Arzehgar <alirezaarzehgar82@gmail.com>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/alirezaarzehgar/git-llm/internal/commit"
	"github.com/alirezaarzehgar/git-llm/internal/llm"
	"github.com/spf13/cobra"
)

var message string

// commitfixCmd represents the commitfix command
var commitfixCmd = &cobra.Command{
	Use:   "commitfix",
	Short: "fix grammar and structure of commit",
	Run: func(cmd *cobra.Command, args []string) {
		err := commit.FixCommitMessage(llm.Groq{}, message)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(commitfixCmd)
	commitfixCmd.PersistentFlags().StringVarP(&message, "message", "m", "", "Use the given message as the commit message")
	commitfixCmd.MarkFlagRequired("message")
}
