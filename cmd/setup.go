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

const (
	PROMPT_FORMAT_GENERATE_COMMIT = `Analyze the git diff --cached output in the # DIFF section and generate a standard Git commit.

	The format is: <type>[(scope)]: <description>
	
	The output must follow this exact JSON structure. Keep the commit clean, short, and informative.
	
	- type: One of fix, feature, doc, refactor, etc.
	- scope: The basename of the folder, or general context like all, maint, etc.
	- description: A concise summary of the change (max 50 characters).
	- body: Start with a single-line explanation of what the commit does, then itemize the specific changes using - at the beginning of each line.
	
	- Each bullet point should not exceed 100 characters.
	- Total body lines (including the explanation) must be 3 to 10 lines.
	
	# Example:
	
	feature(parser): add support for YAML config
	
	Add YAML parsing support to config loader
	- Introduced yaml.Unmarshal in parser.go
	- Updated config_test.go for YAML cases
	- Adjusted README to document YAML support
	
	Output JSON format:
	{
	  "type": "what the commit does or enhances: fix, feature, doc, refactor",
	  "scope": "basename of folder, all, maint, context of change",
	  "description": "short description of what this diff does (max 50 characters)",
	  "body": "one-line explanation followed by bullet-pointed changes with dashes"
	}
	
	# DIFF
	%s
	`
	PROMPT_FORMAT_FIX_COMMIT = `
	Please correct the grammar of the content following the "# TEXT" marker without altering the style,
	structure, or formatting. It is crucial to maintain the original line breaks and not to change the
	size of the text. Only correct grammar mistakes. Don't include "# TEXT" to the message.
	
	Ensure that the output adheres to the following JSON structure:
	{
		"text": "grammar corrected text"
	}
	
	# TEXT
	%s
	`
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
	viper.Set("PROMPT_FORMAT_GENERATE_COMMIT", PROMPT_FORMAT_GENERATE_COMMIT)
	viper.Set("PROMPT_FORMAT_FIX_COMMIT", PROMPT_FORMAT_FIX_COMMIT)

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
