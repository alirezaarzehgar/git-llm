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
		if err := commit.Generate(llm.Qrok{}); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	},
}

func init() {
	rootCmd.AddCommand(commitGenCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// commitCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// commitCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
