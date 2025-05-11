/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"strconv" 

	"github.com/spf13/cobra"
	"github.com/raff/kv4p-go/kv4pht"
)

// debugCmd represents the debug command
var debugCmd = &cobra.Command{
	Use:   "debug",
	Short: "Enable debug mode",
	Aliases: []string{"dbg", "d"},
	Long: `Enable debug mode for the application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			fmt.Println("Invalid number of arguments. Usage: kv4p debug <true|false>")
			return
		}
		debugMode, err := strconv.ParseBool(args[0])
		if err != nil {
			fmt.Println("Invalid argument. Usage: kv4p debug <true|false>")
			return
		}
		fmt.Println("debug called")
		kv4pht.Debug = debugMode
		fmt.Println("Debug mode set to:", debugMode)
	},
}

func init() {
	RootCmd.AddCommand(debugCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// debugCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// debugCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
