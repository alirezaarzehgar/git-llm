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

// commitGenCmd represents the commit command
var commitGenCmd = &cobra.Command{
	Use:   "commitgen",
	Short: "generate commit message and open editor",
	Long: `generate commit message using configured LLM and
open configured EDITOR to change it. Then commit that message.
For example:

git llm commitgen`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := commit.Generate(llm.Groq{}); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(commitGenCmd)
}
