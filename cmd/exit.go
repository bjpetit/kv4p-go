/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/raff/kv4p-go/common"
)

// exitCmd represents the exit command
var exitCmd = &cobra.Command{
	Use:   "exit",
	Short: "Exit the program",
	Long: `Shutdown the radio and exit the program.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("exit called")
		common.ShutdownRadio()
		os.Exit(0)
	},
}

func init() {
	RootCmd.AddCommand(exitCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// exitCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// exitCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
