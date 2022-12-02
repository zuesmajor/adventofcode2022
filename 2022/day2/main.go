package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	player_score := 0
	win_point := 6
	rock_point := 1
	paper_point := 2
	scissor_point := 3

	file, err := os.Open("/Users/patrick/go/adventofcode/day2/input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		elf_hand := string(scanner.Text()[0])
		player_hand := string(scanner.Text()[2])

		if elf_hand == "A" && player_hand == "X" {
			player_score += (1 + 3)
		} else if elf_hand == "B" && player_hand == "Y" {
			player_score += (2 + 3)
		} else if elf_hand == "C" && player_hand == "Z" {
			player_score += (3 + 3)
			// end of draw block
		} else if player_hand == "X" {
			if elf_hand == "C" {
				player_score += (win_point + rock_point)
			} else if elf_hand == "B" {
				player_score += rock_point
			}
		} else if player_hand == "Y" {
			if elf_hand == "A" {
				player_score += (win_point + paper_point)
			} else if elf_hand == "C" {
				player_score += paper_point
			}
		} else if player_hand == "Z" {
			if elf_hand == "B" {
				player_score += (win_point + scissor_point)
			} else if elf_hand == "A" {
				player_score += scissor_point
			}
		}
	}
	fmt.Println(player_score)
}
