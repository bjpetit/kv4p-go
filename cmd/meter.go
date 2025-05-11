/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/raff/kv4p-go/common"
	"github.com/spf13/cobra"
)

// meterCmd represents the meter command
var meterCmd = &cobra.Command{
	Use:   "meter",
	Short: "Show S-Meter Level",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("meter called")
		smeter, scount := common.Kv4phtInstance.SMeter()
		fmt.Printf("S-Meter: %d, count %d\n", smeter, scount)
	},
}

func init() {
	RootCmd.AddCommand(meterCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// meterCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// meterCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
