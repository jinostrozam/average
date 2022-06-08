// average calculates the average of several numbers.

package main

import (
	"fmt"
	"log"

	"github.com/jinostrozam/datafile"
)

func main() {

	// numbers := [3]float64{71.8, 56.2, 89.5}
	numbers, err := datafile.GetFloats("data.txt")
	if err != nil {
		log.Fatal(err)
	}

	var sum float64 = 0

	for _, number := range numbers {
		sum += number
	}

	sampleCount := float64(len(numbers))
	fmt.Printf("Average: %0.2f\n", sum/sampleCount)

}
