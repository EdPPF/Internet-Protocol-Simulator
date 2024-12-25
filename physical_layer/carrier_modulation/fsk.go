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

func fskModulation(A float64, f1 float64, f2 float64, bitStream []int) []float64 {
	var sigSize = len(bitStream)
	var modulatedSignal = make([]float64, sigSize*100)
	for i := 0; i < sigSize; i++ {
		if bitStream[i] == 1 {
			for j := 0; j < 100; j++ {
				modulatedSignal[i*100+j] = A * math.Sin(2 * math.Pi * f1 * float64(j) / 100)
			}
		} else {
			for j := 0; j < 100; j++ {
				modulatedSignal[i*100+j] = A * math.Sin(2 * math.Pi * f2 * float64(j) / 100)
			}
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

	p.Title.Text = "FSK Modulation"
	p.X.Min = 0
	p.X.Max = float64(len(signal))
	p.Y.Min = -A
	p.Y.Max = A

	line, err := plotter.NewLine(points)
	if err != nil {
		log.Fatal(err)
	}
	p.Add(line)
	if err := p.Save(4*vg.Inch, 4*vg.Inch, "fsk.png"); err != nil {
		log.Fatal(err)
	}
}

func main() {
	fmt.Println(os.Args)
	if len(os.Args) < 4 {
		fmt.Println("Usage: go run fsk.go <amplitude> <frequency1> <frequency2>")
		return
	}

	amplitude, err := strconv.ParseFloat(os.Args[1], 64)
	if err != nil {
		fmt.Println("Invalid amplitude. Please provide a valid number.")
		return
	}

	frequency1, err := strconv.ParseFloat(os.Args[2], 64)
	if err != nil {
		fmt.Println("Invalid frequency 1. Please provide a valid number.")
		return
	}

	frequency2, err := strconv.ParseFloat(os.Args[3], 64)
	if err != nil {
		fmt.Println("Invalid frequency 2. Please provide a valid number.")
		return
	}

	// Example input data
	data := []int{1, 0, 1, 1, 0}
	modulatedSignal := fskModulation(amplitude, frequency1, frequency2, data)

	fmt.Println("Modulated Signal:", modulatedSignal)
}