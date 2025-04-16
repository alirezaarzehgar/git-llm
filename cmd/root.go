/*
Copyright © 2025 Alireza Arzehgar

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/
package cmd

import (
	"bufio"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "git-llm",
	Short: "utilizing LLM for managing git projects",
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

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.git-llm.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	viper.SetConfigName("git-llm")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("/etc/git-llm/")
	viper.AddConfigPath("/etc/")
	viper.AddConfigPath("$HOME/.git-llm")
	viper.AddConfigPath("$HOME/")
	viper.AddConfigPath("$HOME/.config")
	viper.AddConfigPath("$HOME/.config/git-llm")

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("Config file notfound. Generate it for the first time")
		scnr := bufio.NewScanner(os.Stdin)

		fmt.Print("GROK_API_KEY: ")
		scnr.Scan()
		apiKey := scnr.Text()

		fmt.Print("LLM_MODEL: ")
		scnr.Scan()
		llmModel := scnr.Text()

		fmt.Print("EDITOR: ")
		scnr.Scan()
		editor := scnr.Text()

		setupConfig(apiKey, llmModel, editor)
	}
}
