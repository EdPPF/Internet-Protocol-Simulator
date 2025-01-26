package baseband_modulation

import (
	"fmt"
)

func manchesterModulation(data []int, v float64) []float64 {
	// Placeholder for modulation logic
	var modulatedSignal []float64
	for _, bit := range data {
		if bit^0 == 0 {
			modulatedSignal = append(modulatedSignal, -v)
		} else {
			modulatedSignal = append(modulatedSignal, v)
		}

		if bit^1 == 0 {
			modulatedSignal = append(modulatedSignal, -v)
		} else {
			modulatedSignal = append(modulatedSignal, v)
		}
	}
	return modulatedSignal
}

func ManchesterModulationWrapper(input interface{}) (interface{}, error) {
	params, ok := input.(struct {
		Data []int
		V    float64
	})
	if !ok {
		return nil, fmt.Errorf("invalid input type for manchesterModulation")
	}

	result := manchesterModulation(params.Data, params.V)
	return result, nil
}

func manchesterDemodulation(signal []float64) []int {
	// Placeholder for demodulation logic
	var signalLen = len(signal)
	var demodulatedSignal []int
	for i := 0; i < signalLen; i += 2 {
		if signal[i] > 0 && signal[i+1] < 0 {
			demodulatedSignal = append(demodulatedSignal, 1)
		} else if signal[i] < 0 && signal[i+1] > 0 {
			demodulatedSignal = append(demodulatedSignal, 0)
		}
	}
	return demodulatedSignal
}

func ManchesterDemodulationWrapper(input interface{}) (interface{}, error) {
	params, ok := input.(struct {
		Data []float64
	})
	if !ok {
		return nil, fmt.Errorf("invalid input type for manchesterDemodulation")
	}

	result := manchesterDemodulation(params.Data)
	return result, nil
}

// func main() {
// 	if len(os.Args) < 2 {
// 		fmt.Println("Usage: go run manchester.go <modulation_value>")
// 		return
// 	}

// 	modulationValue, err := strconv.ParseFloat(os.Args[1], 64)
// 	if err != nil {
// 		fmt.Println("Invalid modulation value. Please provide a valid number.")
// 		return
// 	}

// 	// Example input data
// 	data := []int{1, 0, 1, 1, 0}
// 	modulatedSignal := manchesterModulation(data, modulationValue)

// 	fmt.Println("Modulated Signal:", modulatedSignal)

// 	demodulatedSignal := manchesterDemodulation(modulatedSignal)
// 	fmt.Println("Demodulated Signal:", demodulatedSignal)
// }
