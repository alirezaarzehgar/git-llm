/*
Copyright © 2025 Alireza Arzehgar <alirezaarzehgar82@gmail.com>
*/
package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	apiKey   string
	editor   string
	llmModel string
)

func setupConfig(apiKey, llmModel, editor string) {
	viper.Set("GROK_API_KEY", apiKey)
	viper.Set("LLM_MODEL", llmModel)
	viper.Set("EDITOR", editor)

	home, _ := os.UserHomeDir()
	path := filepath.Join(home, ".config", "git-llm.yaml")
	err := viper.WriteConfigAs(path)
	if err != nil {
		fmt.Println(err)
	}
}

// setupCmd represents the setup command
var setupCmd = &cobra.Command{
	Use:   "setup",
	Short: "Set configuration",
	Run: func(cmd *cobra.Command, args []string) {
		setupConfig(apiKey, llmModel, editor)
	},
}

func init() {
	rootCmd.AddCommand(setupCmd)

	setupCmd.PersistentFlags().StringVarP(&apiKey, "key", "a", "", "your LLM api key")
	setupCmd.PersistentFlags().StringVarP(&editor, "editor", "e", "nano", "editor for commitgen")
	setupCmd.PersistentFlags().StringVarP(&llmModel, "model", "m", "", "LLM model")

	setupCmd.MarkPersistentFlagRequired("key")
	setupCmd.MarkPersistentFlagRequired("model")
}
