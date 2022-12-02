package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	max := 0
	elf_total := 0

	file, err := os.Open("/Users/patrick/go/adventofcode/day1/input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if scanner.Text() != "" {
			// fmt.Println(scanner.Text())
			intVar, err := strconv.Atoi(scanner.Text())
			elf_total += intVar
			if err != nil {
				log.Fatal(err)
			}
		} else {
			if elf_total > max {
				max = elf_total
			}
			elf_total = 0
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(max)
}
