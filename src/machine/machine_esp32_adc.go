//go:build esp32

package machine

const (
	ADC1_0 = ADC{Pin: GPIO36}
	ADC1_3 = ADC{Pin: GPIO39}
	ADC1_4 = ADC{Pin: GPIO32}
	ADC1_6 = ADC{Pin: GPIO34}
	ADC1_7 = ADC{Pin: GPIO35}

	ADC2_8 = ADC{Pin: GPIO25}
	ADC2_9 = ADC{Pin: GPIO26}
	ADC2_7 = ADC{Pin: GPIO14}
	ADC2_6 = ADC{Pin: GPIO12}
	ADC2_5 = ADC{Pin: GPIO13}
	ADC2_4 = ADC{Pin: GPIO9}
	ADC2_0 = ADC{Pin: GPIO4}
	ADC2_1 = ADC{Pin: GPIO0}
	ADC2_2 = ADC{Pin: GPIO2}
	ADC2_3 = ADC{Pin: GPIO15}
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
