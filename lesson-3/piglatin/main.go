package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

const consonants = "bcdfghjklmnpqrstvwxz"
const suffix = "ay"
const validationRegexp = "[^a-zA-Z .,!?-()]" // it's better to prepare a regexp object right here
const punctuationMarks = " .,!?-()"

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Printf("Enter string: ")

	for scanner.Scan() {
		input := strings.Trim(scanner.Text(), " ")

		if !validInput(input) {
			fmt.Println("Invalid input. Please, use latin letters and punctuation marks only.")
			break
		}

		fmt.Println("Your result: ", translate(input))
		break // if you make only one scan, just don't use for-loop
	}

	if scanner.Err() != nil {
		fmt.Println("Oops... Some error occurred: ", scanner.Text())
	}
}

func validInput(s string) bool {
	match, _ := regexp.MatchString(validationRegexp, s) // it's better to have a regexp-object prepared previously, to make this heavy operation only once

	return !match && len(s) > 0
}

func translate(s string) string {
	var translated string
	var lexeme string

	for _, v := range strings.Split(s, "") {
		if strings.Contains(punctuationMarks, v) {
			translated += translateToPigLatin(lexeme)
			translated += v
			lexeme = ""
		} else {
			lexeme += v
		}
	}

	translated += translateToPigLatin(lexeme)

	return translated
}

func translateToPigLatin(s string) string {
	if len(s) == 0 {
		return s
	}
	stringSlice := strings.Split(s, "")


	for i, v := range stringSlice {
		if strings.Contains(consonants, v) || strings.Contains(strings.ToUpper(consonants), v) { // you can just convert v to lower case and compare it to consonants
			continue
		} else { // you don't need else-clause here.
			stringSlice = append(stringSlice, stringSlice[:i]...) // it's a bit strange to change iterated collection inside the loop. it works, but should be refactored
			stringSlice = stringSlice[i:]

			break
		}
	}

	return strings.Join(stringSlice, "") + suffix
}
