package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func CSVToArray ( ) [][]string{
	array := make([][]string, 0)

	file, err := os.Open("input.csv") // Open csv file

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close() // Close csv file after main

	csvReader := csv.NewReader(file)

	for {
		move, err := csvReader.Read() // Read csv file line by line

		if err == io.EOF { // Stop reading at the end of file 
			break
		}

		if err != nil {
			log.Fatal(err)
		}

		course := strings.Split(move[0], " ") // Split strings into slices
		array = append(array, course) // Save slices as a multidimensional array
		
	}
	return array
}

func firstPart (course [][]string) {
	horizontalPosition := 0
	depth := 0

	for i := 0 ; i < len(course) ; i++ {

		number, err := strconv.Atoi(course[i][1]) // Converting string to number
		if err != nil {
			log.Fatal(err)
		}

		// Instructions for navigation of the submarine
		if course[i][0] == "forward" { 
			horizontalPosition += number
		} else if course[i][0] == "down" {
			depth += number
		} else if course[i][0] == "up" {
			depth -= number
		}
	}
	fmt.Printf("Horizontal Position: %d\nDepth: %d\nProduct: %d", horizontalPosition, depth, horizontalPosition*depth)
}

func secondPart (course [][]string) {
	horizontalPosition := 0
	depth := 0
	aim := 0

	for i := 0 ; i < len(course) ; i++ {

		number, err := strconv.Atoi(course[i][1]) // Converting string to number
		if err != nil {
			log.Fatal(err)
		}

		// Instructions for navigation of the submarine
		if course[i][0] == "forward" {
			horizontalPosition += number
			depth += aim * number
		} else if course[i][0] == "down" {
			aim += number
		} else if course[i][0] == "up" {
			aim -= number
		}
	}
	fmt.Printf("Horizontal Position: %d\nDepth: %d\nProduct: %d", horizontalPosition, depth, horizontalPosition*depth)
}

func main () {
	course := CSVToArray()
	firstPart(course)
	secondPart(course)
}