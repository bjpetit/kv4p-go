package main

import (
	"flag"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/c-bata/go-prompt"
	cobraprompt "github.com/stromland/cobra-prompt"

	cmd "github.com/raff/kv4p-go/cmd"
	"github.com/raff/kv4p-go/common"
	"github.com/raff/kv4p-go/kv4pht"
)

var advancedPrompt = &cobraprompt.CobraPrompt{
	RootCmd:                  cmd.RootCmd,
	PersistFlagValues:        true,
	ShowHelpCommandAndFlags:  true,
	DisableCompletionCommand: true,
	AddDefaultExitCommand:    false,
	GoPromptOptions: []prompt.Option{
		prompt.OptionTitle("kv4p"),
		prompt.OptionPrefix(">ht> "),
		prompt.OptionMaxSuggestion(10),
	},
	DynamicSuggestionsFunc: func(annotationValue string, document *prompt.Document) []prompt.Suggest {
		if suggestions := cmd.GetFreqDynamic(annotationValue); suggestions != nil {
			return suggestions
		}

		return []prompt.Suggest{}
	},
	OnErrorFunc: func(err error) {
		if strings.Contains(err.Error(), "unknown command") {
			cmd.RootCmd.PrintErrln(err)
			return
		}

		cmd.RootCmd.PrintErr(err)
		os.Exit(1)
	},
}

//var simplePrompt = &cobraprompt.CobraPrompt{
//	RootCmd:                  cmd.RootCmd,
//	AddDefaultExitCommand:    true,
//	DisableCompletionCommand: true,
//}

func init_radio(dev *string) int {
	// Initialize the radio
	if *dev == "" {
		log.Println("Starting radio (autodiscover)")
	} else {
		log.Println("Starting radio on", *dev)
	}
	p, err := kv4pht.Start(*dev)
	if err != nil {
		log.Fatalf("Start: %v", err)
	}
	common.Kv4phtInstance = p

	// Wait for HELLO message
	for i := 0; i < 2 && !common.Kv4phtInstance.Hello(); i++ {

		if i == 1 {
			log.Println("Reset board")
			common.Kv4phtInstance.Reset()
		}

		log.Println("Waiting for HELLO message...")
		for j := 0; j < 10 && !common.Kv4phtInstance.Hello(); j++ {
			if common.Kv4phtInstance.Hello() {
				log.Println("HELLO message received")
				break
			}
			time.Sleep(1 * time.Second)
		}
	}
	if !common.Kv4phtInstance.Hello() {
		log.Println("No HELLO message received")
		return -1
	}

	//if err := common.Kv4phtInstance.SendStop(); err != nil {
	//	log.Fatalf("Send STOP: %v", err)
	//	return -1
	//}

	band := common.CheckFrequency(common.CurrentRadioSettings.Tx_frequency)
	if band < 0 {
		log.Println("Invalid tx frequency")
		return -1
	}
	common.CurrentRadioSettings.Band = band
	if err := common.Kv4phtInstance.SendConfig(common.CurrentRadioSettings.Band); err != nil {
		log.Fatalf("Send CONFIG: %v", err)
		return -1
	}

	// Wait for VERSION message
	for i := 0; i < 10; i++ {
		v, _, _ := common.Kv4phtInstance.Version()
		if v != 0 {
			break
		}

		log.Println("Waiting for VERSION message...")
		time.Sleep(1 * time.Second)
	}

	if v, _, _ := common.Kv4phtInstance.Version(); v == 0 {
		log.Println("No VERSION message received")
		return -1
	}

	if err := common.Kv4phtInstance.SendFilters(
		common.CurrentRadioSettings.PreFilter,
		common.CurrentRadioSettings.HighPass,
		common.CurrentRadioSettings.LowPass); err != nil {
		log.Fatalf("Send FILTERS: %v", err)
		return -1
	}

	if common.CurrentRadioSettings.Volume < 0 {
		common.CurrentRadioSettings.Volume = 0
	} else if common.CurrentRadioSettings.Volume > 100 {
		common.CurrentRadioSettings.Volume = 100
	}
	common.Kv4phtInstance.SetVolume(float64(common.CurrentRadioSettings.Volume) / 100)

	if common.CurrentRadioSettings.Bandwidth != kv4pht.DRA818_12K5 {
		common.CurrentRadioSettings.Bandwidth = kv4pht.DRA818_25K
	}

	if common.CurrentRadioSettings.Squelch < 0 {
		common.CurrentRadioSettings.Squelch = 0
	} else if common.CurrentRadioSettings.Squelch > 8 {
		common.CurrentRadioSettings.Squelch = 8
	}

	log.Printf("FREQ: %3.3f", common.CurrentRadioSettings.Tx_frequency)
	if err := common.Kv4phtInstance.SendGroup(
		common.CurrentRadioSettings.Bandwidth,
		common.CurrentRadioSettings.Tx_frequency,
		common.CurrentRadioSettings.Rx_frequency,
		common.CurrentRadioSettings.Squelch); err != nil {
		log.Fatalf("Send GROUP: %v", err)
		return -1
	}

	return 0
}

func set_defaults() {

}

func main() {

	dev := flag.String("dev", "", "Serial device to use (e.g. /dev/ttyUSB0)")
	debug := flag.Bool("debug", false, "Enable debug mode")
	flag.Parse()

	kv4pht.Debug = *debug

	shutdown := func() {
		common.ShutdownRadio()
		os.Exit(0)
	}

	defer shutdown()

	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)
	go func() {
		s := <-sigc
		log.Println(s)
		shutdown()
	}()

	set_defaults()

	// Initialize the radio
	if init_radio(dev) != 0 {
		log.Println("Failed to initialize radio")
		return
	}
	advancedPrompt.Run()
}
