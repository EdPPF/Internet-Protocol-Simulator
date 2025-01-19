package carrier_modulation

import (
	"fmt"
	"log"
	"math"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func qam8Modulation(A float64, f float64, bitStream []int) []float64 {
	var sigSize = len(bitStream)

	var mappingBinary = map[int]struct {
		amplitude float64
		phase     float64
	}{
		0b000: {amplitude: 1, phase: 0},
		0b001: {amplitude: 1, phase: math.Pi / 2},
		0b010: {amplitude: 1, phase: math.Pi},
		0b011: {amplitude: 1, phase: 3 * math.Pi / 2},
		0b100: {amplitude: 2, phase: math.Pi / 4},
		0b101: {amplitude: 2, phase: 3 * math.Pi / 4},
		0b110: {amplitude: 2, phase: 5 * math.Pi / 4},
		0b111: {amplitude: 2, phase: 7 * math.Pi / 4},
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
				group[j] = 0 // Padding with 0
			}
		}

		coordinate := mappingBinary[int(group[0])*4+int(group[1])*2+int(group[2])]
		for k := 0; k < 100; k++ {
			modulatedSignal[i*100+k] = A * coordinate.amplitude * math.Sin((2*math.Pi*f*float64(k)/100)+coordinate.phase)
		}
	}
	qam_plotSignal(modulatedSignal, A)
	return modulatedSignal
}

func Qam8ModulationWrapper(input interface{}) (interface{}, error) {
	params, ok := input.(struct {
		A         float64
		F         float64
		BitStream []int
	})
	if !ok {
		return nil, fmt.Errorf("invalid input type for qam8Modulation")
	}

	result := qam8Modulation(params.A, params.F, params.BitStream)
	return result, nil
}

func qam_plotSignal(signal []float64, A float64) {
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

func qam8Demodulation(modulatedSignal []float64, A float64, f float64) []int {
	numSymbols := len(modulatedSignal) / 100
	demodulatedBits := make([]int, numSymbols*3)

	for i := 0; i < numSymbols; i++ {
		firstCoordinate := modulatedSignal[i*100]
		secondCoordinate := modulatedSignal[i*100+1]

		fmt.Println("First Coordinate:", firstCoordinate)
		fmt.Println("Second Coordinate:", secondCoordinate)

		// 0:
		if firstCoordinate > -A/2 && firstCoordinate < A/2 {
			if secondCoordinate > firstCoordinate {
				// upwards
				demodulatedBits[i*3] = 0
				demodulatedBits[i*3+1] = 0
				demodulatedBits[i*3+2] = 0
			} else {
				// downwards
				demodulatedBits[i*3] = 0
				demodulatedBits[i*3+1] = 1
				demodulatedBits[i*3+2] = 0
			}
		}

		// A:
		if firstCoordinate <= A && firstCoordinate >= A/2 {
			demodulatedBits[i*3] = 0
			demodulatedBits[i*3+1] = 0
			demodulatedBits[i*3+2] = 1
		}

		// -A:
		if firstCoordinate >= -A && firstCoordinate <= -A/2 {
			demodulatedBits[i*3] = 0
			demodulatedBits[i*3+1] = 1
			demodulatedBits[i*3+2] = 1
		}

		// ~2A:
		if firstCoordinate > A {
			if secondCoordinate > firstCoordinate {
				// upwards
				demodulatedBits[i*3] = 1
				demodulatedBits[i*3+1] = 0
				demodulatedBits[i*3+2] = 0
			} else {
				// downwards
				demodulatedBits[i*3] = 1
				demodulatedBits[i*3+1] = 0
				demodulatedBits[i*3+2] = 1
			}
		}

		// ~-2A:
		if firstCoordinate < -A {
			if secondCoordinate > firstCoordinate {
				// upwards
				demodulatedBits[i*3] = 1
				demodulatedBits[i*3+1] = 1
				demodulatedBits[i*3+2] = 1
			} else {
				// downwards
				demodulatedBits[i*3] = 1
				demodulatedBits[i*3+1] = 1
				demodulatedBits[i*3+2] = 0
			}
		}
	}

	return demodulatedBits
}

func Qam8DemodulationWrapper(input interface{}) (interface{}, error) {
	params, ok := input.(struct {
		ModulatedSignal []float64
		A               float64
		F               float64
	})
	if !ok {
		return nil, fmt.Errorf("invalid input type for qam8Demodulation")
	}

	result := qam8Demodulation(params.ModulatedSignal, params.A, params.F)
	return result, nil
}

// func main() {
// 	if len(os.Args) < 3 {
// 		fmt.Println("Usage: go run qam_8.go <amplitude> <frequency>")
// 		return
// 	}

// 	amplitude, err := strconv.ParseFloat(os.Args[1], 64)
// 	if err != nil {
// 		fmt.Println("Invalid amplitude. Please provide a valid number.")
// 		return
// 	}

// 	frequency, err := strconv.ParseFloat(os.Args[2], 64)
// 	if err != nil {
// 		fmt.Println("Invalid frequency. Please provide a valid number.")
// 		return
// 	}

// 	// Example input data
// 	data := []int{1, 1, 1, 1, 0}
// 	modulatedSignal := qam8Modulation(amplitude, frequency, data)

// 	fmt.Println("Modulated Signal:", modulatedSignal)

// 	demodulatedSignal := qam8Demodulation(modulatedSignal, amplitude, frequency)
// 	fmt.Println("Demodulated Signal:", demodulatedSignal)
// }
