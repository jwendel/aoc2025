package main

import (
	"bufio"
	"errors"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

	// 50857215695 is too high

	file, err := os.DirFS(".").Open("sample_input.txt")
	if err != nil {
		log.Fatalln("problem reading file", err)
	}

	reader := bufio.NewReader(file)

	end := false
	total := 0
	for {

		data, err := reader.ReadString(',')
		if err != nil {
			if errors.Is(err, io.EOF) {
				end = true
			} else {
				log.Fatalln("problem reading file", err)
			}
		}

		before, after, found := strings.Cut(data, "-")
		if !found {
			log.Fatalln("1problem parsing string: ", data)
		}

		lower, err := strconv.Atoi(before)
		if err != nil {
			log.Fatalln("2problem parsing string: ", before)
		}

		after = strings.Trim(after, ",")
		upper, err := strconv.Atoi(after)
		if err != nil {
			log.Fatalln("3problem parsing string: ", after)
		}

		log.Println("upper/lower: ", lower, upper)

		for i := lower; i <= upper; i++ {
			found := checkForMatch(i)
			if found {
				total += i
				log.Printf("found match in range %d-%d. Value: %d. Current total: %d", lower, upper, i, total)
			}
		}

		if end {
			break
		}
	}
	log.Printf("total: %v", total)
}

func checkForMatch(value int) (found bool) {
	strValue := strconv.Itoa(value)
	beforeLen := len(strValue)
	maxVal := (beforeLen + 1) / 2

outer:
	for j := 1; j <= maxVal; j++ {
		if beforeLen%j != 0 {
			continue
		}
		pattern := strValue[:j]
		for k := j; k < beforeLen; k += j {
			toMatch := strValue[k : k+j]
			if pattern != toMatch {
				continue outer
			}
		}
		return true
	}
	return false
}
