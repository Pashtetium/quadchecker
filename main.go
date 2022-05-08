package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//read console
	var input []string
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		line := s.Text()

		input = append(input, line)
	}

	fmt.Println(input)
	if CheckRowsLength(input) && CheckForWhitespace(input) && CheckIfLettersValid(input) {
		fmt.Println("zaebis")

		//checking quadA case
		if input[0][0] == 'o' {
			fmt.Printf("[quadA] [%v] [%v]", len(input[0]), len(input))
		}

		//checking quadB case
		if input[0][0] == '/' {
			fmt.Printf("[quadB] [%v] [%v]", len(input[0]), len(input))
		}

		//checking quad C D E cases
		if input[0][0] == 'A' {
			if len(input) == 1 && len(input[0]) == 1 {
				fmt.Printf("[quadC] [1] [1] || [quadD] [1] [1] || [quadE] [%v] [%v]", len(input[0]), len(input))
				return
			}

			if len(input) == 2 && input[0][0] == 'A' && input[len(input)-1][0] == 'C' {
				fmt.Print("[quadC] [1] [2] || [quadE] [1] [2]")
				return
			}

			if input[0][len(input[0])-1] == 'A' {
				fmt.Printf("[quadC] [%v] [%v]", len(input[0]), len(input))
			}

			if input[0][len(input[0])-1] == 'C' {

				if input[len(input)-1][0] == 'A' {
					fmt.Printf("[quadD] [%v] [%v]", len(input[0]), len(input))
				} else {
					if input[len(input)-1][0] == 'C' {
						fmt.Printf("[quadE] [%v] [%v]", len(input[0]), len(input))
					}
				}
			}
		}
	} else {
		fmt.Println("Not a Raid function")
	}
}

func CheckIfLettersValid(input []string) bool {
	validLetters := []rune{'o', '-', '|', '/', '*', '\\', 'B', 'A', 'C'}
	var nonValidFound bool

	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[i]); j++ {

			for z := 0; z < len(validLetters); z++ {
				if input[i][j] != byte(validLetters[z]) {
					nonValidFound = false
				}
			}
		}
	}

	if nonValidFound {
		return false
	}

	return true
}

func CheckForWhitespace(input []string) bool {
	//checking first row for whitespace
	//valid quad should be without spaces on first row
	for i := 0; i < len(input[0]); i++ {
		if input[0][i] == ' ' {
			return false
		}
	}

	return true
}

func CheckRowsLength(input []string) bool {
	//checking if rows length is same
	for i := 0; i < len(input); i++ {
		if len(input[0]) != len(input[i]) {
			return false
		}
	}

	return true
}
