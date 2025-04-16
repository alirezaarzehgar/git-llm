/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "manage config",
	Run: func(cmd *cobra.Command, args []string) {
		// fmt.Println("config called")
	},
}

var configSetCmd = &cobra.Command{
	Use:   "set",
	Short: "set config",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("set config called")
	},
}

var configGetCmd = &cobra.Command{
	Use:   "get",
	Short: "get config",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("get config called")
	},
}

func init() {
	rootCmd.AddCommand(configCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// configCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// configCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	configCmd.AddCommand(configSetCmd)
	configCmd.AddCommand(configGetCmd)
}
