package main

import (
	"bufio"
	"fmt"
	_ "go/scanner"
	"os"
	"regexp"
	"strings"
)

const consonants = "bcdfghjklmnpqrstvwxz"
const suffix = "ay"
const punctuationMarks = " .,!?-()"

var validationRegexp, _ = regexp.Compile(`[^a-zA-Z .,!?\-()]`)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Printf("Enter string: ")

	scanner.Scan()
	input := strings.Trim(scanner.Text(), " ")

	if !validInput(input) {
		fmt.Println("Invalid input. Please, use latin letters and punctuation marks only.")
		return
	}

	fmt.Println("Your result: ", translate(input))
	if scanner.Err() != nil {
		fmt.Println("Oops... Some error occurred: ", scanner.Text())
	}
}

func validInput(s string) bool {
	return !validationRegexp.MatchString(s) && len(s) > 0
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
		if strings.Contains(consonants, strings.ToLower(v)) {
			continue
		}

		tmpSlice := append(stringSlice, stringSlice[:i]...)
		stringSlice = tmpSlice[i:]

		break
	}

	return strings.Join(stringSlice, "") + suffix
}
