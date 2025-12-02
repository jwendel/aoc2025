package main

import (
	"bufio"
	"errors"
	"io"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.DirFS(".").Open("input.txt")
	if err != nil {
		log.Fatalln("problem reading file", err)
	}

	reader := bufio.NewReader(file)

	p := 50
	count := 0
	for {
		data, isPrefix, err := reader.ReadLine()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			log.Fatalln(isPrefix, err)
		}
		d := string(data)
		num, err := strconv.Atoi(d[1:])
		if err != nil {
			log.Fatalln("atoi fail: ", err)
		}
		switch d[0] {
		case 'L':
			p = p - num
		case 'R':
			p = p + num
		default:
			log.Fatalln("bad data? ", d)
		}
		for {
			if p < 0 {
				p += 100
			}
			if p > 99 {
				p -= 100
			}
			if p >= 0 && p <= 99 {
				break
			}
		}
		log.Println(d, "  ", p)
		if p == 0 {
			count++
		}
	}
	log.Println(count)

}
