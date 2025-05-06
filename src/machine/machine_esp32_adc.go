//go:build esp32

package machine

var (
	ADC1_CHANNEL_0 = ADC{Pin: GPIO36}
	ADC1_CHANNEL_1 = ADC{Pin: GPIO37}
	ADC1_CHANNEL_2 = ADC{Pin: GPIO38}
	ADC1_CHANNEL_3 = ADC{Pin: GPIO39}
	ADC1_CHANNEL_4 = ADC{Pin: GPIO32}
	ADC1_CHANNEL_6 = ADC{Pin: GPIO34}
	ADC1_CHANNEL_7 = ADC{Pin: GPIO35}

	ADC2_CHANNEL_0 = ADC{Pin: GPIO4}
	ADC2_CHANNEL_1 = ADC{Pin: GPIO0}
	ADC2_CHANNEL_2 = ADC{Pin: GPIO2}
	ADC2_CHANNEL_3 = ADC{Pin: GPIO15}
	ADC2_CHANNEL_4 = ADC{Pin: GPIO13}
	ADC2_CHANNEL_5 = ADC{Pin: GPIO12}
	ADC2_CHANNEL_6 = ADC{Pin: GPIO14}
	ADC2_CHANNEL_7 = ADC{Pin: GPIO27}
	ADC2_CHANNEL_8 = ADC{Pin: GPIO25}
	ADC2_CHANNEL_9 = ADC{Pin: GPIO26}
)

const (
	ADC_ATTEN_DB_0   = iota // 100 mV ~ 950 mV
	ADC_ATTEN_DB_2_5        // 100 mV ~ 1250 mV
	ADC_ATTEN_DB_6          // 150 mV ~ 1750 mV
	ADC_ATTEN_DB_11         // 150 mV ~ 2450 mV
)

// InitADC initializes the ADC.
func InitADC() {

}

// Configure configures a ADC pin to be able to be used to read data.
func (a ADC) Configure(config ADCConfig) {

}

// Get returns the current value of a ADC pin, in the range 0..0xffff.
func (a ADC) Get() uint16 {
	return 0
}

func (a ADC) getADCChannel() uint8 {
	return 0
}

func waitADCSync() {

}
