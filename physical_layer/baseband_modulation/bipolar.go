package baseband_modulation

// import (
// 	"fmt"
// 	"os"
// 	"strconv"
// )

func BipolarModulation(data []int, v float64) []float64 {
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

func BipolarDemodulation(signal []float64) []int {
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

// func main() {
// 	if len(os.Args) < 2 {
// 		fmt.Println("Usage: go run bipolar.go <modulation_value>")
// 		return
// 	}

// 	modulationValue, err := strconv.ParseFloat(os.Args[1], 64)
// 	if err != nil {
// 		fmt.Println("Invalid modulation value. Please provide a valid number.")
// 		return
// 	}

// 	// Example input data
// 	data := []int{1, 0, 1, 1, 0}
// 	modulatedSignal := BipolarModulation(data, modulationValue)

// 	fmt.Println("Modulated Signal:", modulatedSignal)

// 	demodulatedSignal := BipolarDemodulation(modulatedSignal)
// 	fmt.Println("Demodulated Signal:", demodulatedSignal)
// }
