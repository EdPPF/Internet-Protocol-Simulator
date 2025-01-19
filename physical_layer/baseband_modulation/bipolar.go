package baseband_modulation

import (
	"fmt"
)

func bipolarModulation(data []int, v float64) []float64 {
	// Placeholder for modulation logic
	var modulatedSignal []float64
	var lastBit float64 = -v
	for _, bit := range data {
		if bit == 1 {
			if lastBit > 0 {
				modulatedSignal = append(modulatedSignal, -v)
			} else {
				modulatedSignal = append(modulatedSignal, v)
			}
			lastBit = -1 * lastBit
		} else {
			modulatedSignal = append(modulatedSignal, 0)
		}
	}
	return modulatedSignal
}

func BipolarModulationWrapper(input interface{}) (interface{}, error) {
	params, ok := input.(struct {
		Data []int
		V    float64
	})
	if !ok {
		return nil, fmt.Errorf("invalid input type for bipolarModulation")
	}

	result := bipolarModulation(params.Data, params.V)
	return result, nil
}

func bipolarDemodulation(signal []float64) []int {
	// Placeholder for demodulation logic
	var demodulatedSignal []int
	for _, value := range signal {
		if value != 0 {
			demodulatedSignal = append(demodulatedSignal, 1)
		} else {
			demodulatedSignal = append(demodulatedSignal, 0)
		}
	}
	return demodulatedSignal
}

func BipolarDemodulationWrapper(input interface{}) (interface{}, error) {
	params, ok := input.(struct {
		Data []float64
	})
	if !ok {
		return nil, fmt.Errorf("invalid input type for bipolarDemodulation")
	}

	result := bipolarDemodulation(params.Data)
	return result, nil
}

/*
func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run bipolar.go <modulation_value>")
		return
	}

	modulationValue, err := strconv.ParseFloat(os.Args[1], 64)
	if err != nil {
		fmt.Println("Invalid modulation value. Please provide a valid number.")
		return
	}

	// Example input data
	data := []int{1, 0, 1, 1, 0}
	modulatedSignal := bipolarModulation(data, modulationValue)

	fmt.Println("Modulated Signal:", modulatedSignal)

	demodulatedSignal := bipolarDemodulation(modulatedSignal)
	fmt.Println("Demodulated Signal:", demodulatedSignal)
}
*/
