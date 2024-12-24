package main

import (
	"fmt"
	"os"
	"strconv"
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
			lastBit = -1*lastBit
		} else {
			modulatedSignal = append(modulatedSignal, 0)
		}
	}
	return modulatedSignal
}

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
}