package common

import (
	"github.com/raff/kv4p-go/kv4pht"
)

var Kv4phtInstance *kv4pht.CommandProcessor = nil
var CurrentRadioSettings = &RadioSettings{
	Volume:       10,
	Bandwidth:    kv4pht.DRA818_25K, //wideband
	Tx_frequency: 146.520,
	Rx_frequency: 146.520,
	Band:         kv4pht.MODE_VHF,
	Squelch:      0,
	Tx_ctcss:     0,
	Rx_ctcss:     0,
	Tx_dcs:       0,
	Rx_dcs:       0,
	PreFilter:    false,
	HighPass:     true,
	LowPass:      true,
}

type RadioSettings struct {
	Volume       int
	Bandwidth    int
	Tx_frequency float64
	Rx_frequency float64
	Band         int
	Squelch      int
	Tx_ctcss     int
	Rx_ctcss     int
	Tx_dcs       int
	Rx_dcs       int
	PreFilter    bool
	HighPass     bool
	LowPass      bool
}
