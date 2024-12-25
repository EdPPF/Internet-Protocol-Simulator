package main

import (
	"os"
	"fmt"
	"strconv"
	"math"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/vg"
	"gonum.org/v1/plot/plotter"
	"log"
)

func qam8Modulation(A float64, f float64, bitStream []int) []float64 {
	var sigSize = len(bitStream)

	var mappingBinary = map[int]struct {
		amplitude float64
		phase     float64
	}{
		0b000: {amplitude: 1, phase: 0},
		0b001: {amplitude: 1, phase: math.Pi/2},
		0b010: {amplitude: 1, phase: math.Pi},
		0b011: {amplitude: 1, phase: 3*math.Pi/2},
		0b100: {amplitude: 2, phase: math.Pi/4},
		0b101: {amplitude: 2, phase: 3*math.Pi/4},
		0b110: {amplitude: 2, phase: 5*math.Pi/4},
		0b111: {amplitude: 2, phase: 7*math.Pi/4},
	}

	numGroups := (sigSize + 2) / 3
	
	var modulatedSignal = make([]float64, numGroups*100)
	for i := 0; i < numGroups; i++ {
		start := i * 3
		group := make([]int, 3)
		for j := 0; j < 3; j++ {
            if start+j < len(bitStream) {
                group[j] = bitStream[start+j]
            } else {
                group[j] = 0  // Padding with 0
            }
        }
		
		coordinate := mappingBinary[int(group[0])*4 + int(group[1])*2 + int(group[2])]
		for k := 0; k < 100; k++ {
			modulatedSignal[i*100+k] = A * coordinate.amplitude * math.Sin(2 * math.Pi * (f + coordinate.phase) * float64(k) / 100)
		}
	}
	plotSignal(modulatedSignal, A)
	return modulatedSignal
}

func plotSignal(signal []float64, A float64) {
	points := make(plotter.XYs, len(signal))
	for i, v := range signal {
		points[i].X = float64(i)
		points[i].Y = v
	}
	p := plot.New()

	p.Title.Text = "8QAM Modulation"
	p.X.Min = 0
	p.X.Max = float64(len(signal))
	p.Y.Min = -A
	p.Y.Max = A

	line, err := plotter.NewLine(points)
	if err != nil {
		log.Fatal(err)
	}
	p.Add(line)
	if err := p.Save(4*vg.Inch, 4*vg.Inch, "qam8.png"); err != nil {
		log.Fatal(err)
	}
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run qam_8.go <amplitude> <frequency>")
		return
	}

	amplitude, err := strconv.ParseFloat(os.Args[1], 64)
	if err != nil {
		fmt.Println("Invalid amplitude. Please provide a valid number.")
		return
	}

	frequency, err := strconv.ParseFloat(os.Args[2], 64)
	if err != nil {
		fmt.Println("Invalid frequency. Please provide a valid number.")
		return
	}

	// Example input data
	data := []int{1, 0, 1, 1, 0}
	modulatedSignal := qam8Modulation(amplitude, frequency, data)

	fmt.Println("Modulated Signal:", modulatedSignal)
}