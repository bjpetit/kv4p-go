/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"strconv"

	"github.com/raff/kv4p-go/common"
	"github.com/spf13/cobra"
)

// squelchCmd represents the squelch command
var squelchCmd = &cobra.Command{
	Use:   "squelch",
	Short: "Get or set the squelch level (0-8)",
	Long: `Get or set the squelch level (0-8).
If no argument is given, the current squelch level is displayed.
If an argument is given, it sets the squelch level to that value.`,
	Args: cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// fmt.Println("squelch called")
		if len(args) == 0 {
			fmt.Println("Squelch ", common.CurrentRadioSettings.Squelch)
			return
		}
		if len(args) != 1 {
			fmt.Println("Invalid number of arguments")
			return
		}
		squelch, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Invalid squelch format")
			return
		}
		if squelch < 0 || squelch > 8 {
			fmt.Println("Invalid squelch level")
			return
		}
		// Set the volume
		common.CurrentRadioSettings.Squelch = squelch
		if common.Kv4phtInstance != nil {
			if err := common.Kv4phtInstance.SendGroup(
				common.CurrentRadioSettings.Bandwidth,
				common.CurrentRadioSettings.Tx_frequency,
				common.CurrentRadioSettings.Rx_frequency,
				common.CurrentRadioSettings.Squelch); err != nil {
				log.Fatalf("Send GROUP: %v", err)
				return
			}
		} else {
			fmt.Println("Radio not initialized")
			return
		}
		fmt.Println("Squelch set to", squelch)
	},
}

func init() {
	RootCmd.AddCommand(squelchCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// squelchCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// squelchCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
