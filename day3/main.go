package main

import (
	"bufio"
	"errors"
	"io"
	"log"
	"os"
)

func main() {
	file, err := os.DirFS(".").Open("input.txt")
	if err != nil {
		log.Fatalln("problem reading file", err)
	}

	reader := bufio.NewReader(file)

	total := 0
	for {
		end := false
		data, err := reader.ReadString('\n')

		if err != nil {
			if errors.Is(err, io.EOF) {
				end = true
			} else {
				log.Fatalln("problem reading file", err)
			}
			if len(data) == 0 {
				break
			}
		}
		data = data[:len(data)-1]
		log.Println("line: :", data)

		maxVal := data[0] - '0'
		pos := 0

		for i := 1; i < (len(data) - 1); i++ {
			val := data[i] - '0'
			if val > maxVal {
				pos = i
				maxVal = val
			}
		}

		maxVal2 := data[pos+1] - '0'
		pos2 := pos + 1
		for i := pos2; i < len(data); i++ {
			val := data[i] - '0'
			if val > maxVal2 {
				pos2 = i
				maxVal2 = val
			}
		}

		res := (maxVal * 10) + (maxVal2)
		log.Println("max:", maxVal, "pos: ", pos, "max2", maxVal2, "pos2: ", pos2, "res: ", res)

		total += int(res)

		if end {
			break
		}
	}
	log.Println("total:", total)
}
