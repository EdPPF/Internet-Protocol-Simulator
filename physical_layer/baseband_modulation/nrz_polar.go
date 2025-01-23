package baseband_modulation

// import (
// 	"fmt"
// 	"os"
// 	"strconv"
// )

// PolarNRZModulation applies polar NRZ modulation to the input data.
func PolarNRZModulation(data []int, v float64) []float64 {
	// Placeholder for modulation logic
	var modulatedSignal []float64
	for _, bit := range data {
		if bit == 1 {
			modulatedSignal = append(modulatedSignal, v)
		} else {
			modulatedSignal = append(modulatedSignal, -v)
		}
	}
	return modulatedSignal
}

func PolarNRZDemodulation(signal []float64) []int {
	// Placeholder for demodulation logic
	var demodulatedSignal []int
	for _, value := range signal {
		if value > 0 {
			demodulatedSignal = append(demodulatedSignal, 1)
		} else {
			demodulatedSignal = append(demodulatedSignal, 0)
		}
	}
	return demodulatedSignal
}

// func main() {
// 	if len(os.Args) < 2 {
// 		fmt.Println("Usage: go run nrz_polar.go <modulation_value>")
// 		return
// 	}

// 	modulationValue, err := strconv.ParseFloat(os.Args[1], 64)
// 	if err != nil {
// 		fmt.Println("Invalid modulation value. Please provide a valid number.")
// 		return
// 	}

// 	// Example input data
// 	data := []int{1, 0, 1, 1, 0}
// 	modulatedSignal := PolarNRZModulation(data, modulationValue)

// 	fmt.Println("Modulated Signal:", modulatedSignal)

// 	demodulatedSignal := PolarNRZDemodulation(modulatedSignal)
// 	fmt.Println("Demodulated Signal:", demodulatedSignal)
// }
