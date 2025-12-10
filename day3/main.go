package main

import (
	"bufio"
	"errors"
	"io"
	"log"
	"math"
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

		var numbers [12]uint8
		// 171741365185404 is too low

		result := 0
		lastPos := 0
		for offset := 11; offset >= 0; offset-- {
			maxVal := data[lastPos] - '0'

			for i := lastPos + 1; i < (len(data) - offset); i++ {
				val := data[i] - '0'
				if val > maxVal {
					lastPos = i
					maxVal = val
				}
			}
			lastPos++
			result += int(maxVal) * int(math.Pow10(offset))
			numbers[11-offset] = maxVal
		}

		//maxVal2 := data[pos+1] - '0'
		//pos2 := pos + 1
		//for i := pos2; i < len(data); i++ {
		//	val := data[i] - '0'
		//	if val > maxVal2 {
		//		pos2 = i
		//		maxVal2 = val
		//	}
		//}

		//res := (maxVal * 10) + (maxVal2)
		total += result
		log.Println("numbers: ", numbers, "result: ", result, "total: ", total)

		//total += int(res)

		if end {
			break
		}
	}
	log.Println("total:", total)
}
