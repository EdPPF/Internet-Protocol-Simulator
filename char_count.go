package main

import (
	"fmt"
    "errors"
)

func char_count_client(data []int) []int {
    // Rewrite the data array with its size at the first position
    var x int
    x = len(data)
    data = append(data, x)
    for i:= len(data)-1; i>= 1; i--{
        data[i] = data[i-1]
    }
    data[0] = x
	return data
}

func char_count_server(data []int) ([]int, error) {
    // Rewrite the data array removing its size from the first position
    if (data[0] == len(data)-1){
        for i:= 0; i < len(data)-1; i++{
            data[i] = data[i+1]
        }
    } else {
        return data, errors.New("Transmission failed")

    }
    data = data[:len(data)-1]
	return data, nil
}

func main() {

	// Example input data
	data := []int{1, 0, 1, 1, 0}

	signal_client := char_count_client(data)
	fmt.Println("Client signal:", signal_client)

	signal_server, err := char_count_server(signal_client)
    if err != nil {
	    fmt.Println(err)
    }
    fmt.Println("Server signal:", signal_server)
}
