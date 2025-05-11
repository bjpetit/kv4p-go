/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:           "kv4p",
	SilenceUsage:  true, // Only print usage when defined in command.
	SilenceErrors: true,
}
