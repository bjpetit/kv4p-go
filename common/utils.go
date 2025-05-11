package common

import (
	"log"

	"github.com/raff/kv4p-go/kv4pht"
)

func CheckFrequency(freq float64) int {
	if freq >= kv4pht.VHF_MIN_FREQ && freq <= kv4pht.VHF_MAX_FREQ {
		return kv4pht.MODE_VHF
	} else if freq >= kv4pht.UHF_MIN_FREQ && freq <= kv4pht.UHF_MAX_FREQ {
		return kv4pht.MODE_UHF
	} else {
		return -1
	}
}

func ShutdownRadio() {
	if Kv4phtInstance != nil {
		log.Println("Stopping radio")
		Kv4phtInstance.Stop()
	}
}
