/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"strconv"

	"github.com/c-bata/go-prompt"
	"github.com/spf13/cobra"
	cobraprompt "github.com/stromland/cobra-prompt"

	"github.com/raff/kv4p-go/common"
)

var getFreqDynamicAnnotationValue = "Frequency"

var GetFreqDynamic = func(annotationValue string) []prompt.Suggest {
	if annotationValue != getFreqDynamicAnnotationValue {
		return nil
	}

	return []prompt.Suggest{
		{Text: "146.520", Description: "Calling frequency"},
		{Text: "146.580", Description: "Adventure frequency"},
	}
}

// frequencyCmd represents the frequency command
var frequencyCmd = &cobra.Command{
	Use:     "frequency [tx frequency] [rx frequency]",
	Short:   "Get/Set the frequency",
	Aliases: []string{"freq", "f"},
	Long: `Get/Set the recieve frequency of the radio.:
 	  If no arguments are given, the current tx and rx frequencies are printed.
 	  If one argument is given, it is set as the tx frequency and the rx frequency is set to the same value.
 	  If two arguments are given, the first is set as the tx frequency and the second as the rx frequency.
	  Argument is the frequency in MHz.
 	    Example: kv4p frequency 162.4`,
	Annotations: map[string]string{
		cobraprompt.DynamicSuggestionsAnnotation: getFreqDynamicAnnotationValue,
	},

	Run: func(cmd *cobra.Command, args []string) {
		// fmt.Println("frequency called", args)
		if len(args) == 0 {
			fmt.Println("Current tx frequency:", common.CurrentRadioSettings.Tx_frequency)
			fmt.Println("Current rx frequency:", common.CurrentRadioSettings.Rx_frequency)
			return
		}
		if len(args) > 2 {
			fmt.Println("Invalid number of arguments")
			return
		}
		if len(args) >= 1 {
			txFreq, err := strconv.ParseFloat(args[0], 64)
			if err != nil {
				fmt.Println("Invalid tx frequency format")
				return
			}
			band := common.CheckFrequency(txFreq)
			if band < 0 {
				fmt.Println("Invalid tx frequency")
				return
			}
			common.CurrentRadioSettings.Tx_frequency = txFreq
			common.CurrentRadioSettings.Band = band
		}
		if len(args) == 2 {
			rxFreq, err := strconv.ParseFloat(args[1], 64)
			if err != nil {
				fmt.Println("Invalid rx frequency format")
				return
			}
			band := common.CheckFrequency(rxFreq)
			if band < 0 {
				fmt.Println("Invalid tx frequency")
				return
			}
			common.CurrentRadioSettings.Rx_frequency = rxFreq
		} else {
			common.CurrentRadioSettings.Rx_frequency = common.CurrentRadioSettings.Tx_frequency
		}

		if err := common.Kv4phtInstance.SendConfig(common.CurrentRadioSettings.Band); err != nil {
			log.Fatalf("Send CONFIG: %v", err)
			return
		}

		if err := common.Kv4phtInstance.SendGroup(
			common.CurrentRadioSettings.Bandwidth,
			common.CurrentRadioSettings.Tx_frequency,
			common.CurrentRadioSettings.Rx_frequency,
			common.CurrentRadioSettings.Squelch); err != nil {
			log.Fatalf("Send GROUP: %v", err)
			return
		}
	},
}

func init() {
	RootCmd.AddCommand(frequencyCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// frequencyCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// frequencyCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
