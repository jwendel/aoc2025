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

	file, err := os.DirFS(".").Open("sample_input.txt")
	if err != nil {
		log.Fatalln("problem reading file", err)
	}

	reader := bufio.NewReader(file)

	end := false
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
			log.Fatalln("problem parsing string: ", data)
		}

		lower, err := strconv.Atoi(before)
		if err != nil {
			log.Fatalln("problem parsing string: ", before)
		}

		after = strings.Trim(after, ",")
		upper, err := strconv.Atoi(after)
		if err != nil {
			log.Fatalln("problem parsing string: ", after)
		}

		log.Println(lower, upper)

		if end {
			break
		}
	}
}
