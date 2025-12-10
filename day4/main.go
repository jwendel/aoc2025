package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatalln("problem reading file", err)
	}
	sdata := string(data)
	fmt.Println(sdata)
	lineLength := strings.Index(sdata, "\n")
	if lineLength <= 0 {
		log.Fatalln("didn't find newline")
	}
	rows := (len(sdata)) / lineLength
	fmt.Println(lineLength, len(sdata), rows)

	array := make([][]bool, rows)
	for i := 0; i < rows; i++ {
		array[i] = make([]bool, lineLength)
	}

	pos := 0
	for i := 0; i < rows; i++ {
		for j := 0; j < lineLength; j++ {
			if sdata[pos] == '@' {
				array[i][j] = true
			}
			pos++
		}
		pos++
	}

	for i := 0; i < rows; i++ {
		fmt.Println(array[i])
	}

	total := 0
	for i := 0; i < rows; i++ {
		for j := 0; j < lineLength; j++ {
			if array[i][j] && getCount(array, i, j, rows, lineLength) < 4 {
				total++
				fmt.Print("x")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
	fmt.Println("total: ", total)
}

func getCount(array [][]bool, row int, col int, rows int, cols int) int {
	count := 0
	if row > 0 && col > 0 && array[row-1][col-1] {
		count++
	}
	if row > 0 && array[row-1][col] {
		count++
	}
	if row < rows-1 && col > 0 && array[row+1][col-1] {
		count++
	}
	if row < rows-1 && array[row+1][col] {
		count++
	}
	if row < rows-1 && col < cols-1 && array[row+1][col+1] {
		count++
	}
	if col < cols-1 && array[row][col+1] {
		count++
	}
	if row > 0 && col < cols-1 && array[row-1][col+1] {
		count++
	}
	if row > 0 && array[row-1][col] {
		count++
	}

	return count
}
