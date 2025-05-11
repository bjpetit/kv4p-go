/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strconv"

	"github.com/raff/kv4p-go/common"
	"github.com/spf13/cobra"
)

// volumeCmd represents the volume command
var volumeCmd = &cobra.Command{
	Use:     "volume",
	Short:   "Set the volume",
	Aliases: []string{"vol", "v"},
	Long: `Set the volume of the radio.
	Argument is the volume level (0-100).
	Example: kv4p volume 50`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("volume called")
		if len(args) != 1 {
			fmt.Println("Invalid number of arguments")
			return
		}
		volume, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Invalid volume format")
			return
		}
		if volume < 0 || volume > 100 {
			fmt.Println("Invalid volume level")
			return
		}
		// Set the volume
		common.CurrentRadioSettings.Volume = volume
		if common.Kv4phtInstance != nil {
			common.Kv4phtInstance.SetVolume(float64(common.CurrentRadioSettings.Volume) / 100)
		} else {
			fmt.Println("Radio not initialized")
			return
		}
		fmt.Println("Volume set to", volume)
	},
}

func init() {
	RootCmd.AddCommand(volumeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// volumeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// volumeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
