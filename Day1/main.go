package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

func CSVToArray ( ) []int{
	array := make([]int, 0)

	file, err := os.Open("input1.csv")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	csvReader := csv.NewReader(file)

	for {
		measure, err := csvReader.Read()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal(err)
		}

		number, err2 := strconv.Atoi(measure[0])

		if err2 != nil {
			log.Fatal(err2)
		}
		array = append(array, number)
		
	}
	return array
}

func firstPart (array []int) {
	cont := 0
	prev := array[0]
	for i, measure := range array {
		if i != 0 && measure > prev { // skip the first measurement
			cont ++
		}
		prev = measure
	}

	fmt.Printf("There are %d measurments that are larger than the previous\n", cont)
}

func secondPart (data []int) {
	cont := 0
	prevSum := data[0] + data [1] + data [2]

	for i := 1; i <= len(data) - 3; i++  {
		sum := data[i] + data[i+1] + data[i+2]
		if sum > prevSum{
			cont++
		}
		prevSum = sum
	}
	fmt.Printf("There are %d sums that are larger than the previous\n", cont)
}

func main () {
	measurements := CSVToArray()

	firstPart(measurements)
	secondPart(measurements)
}