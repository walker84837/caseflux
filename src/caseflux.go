package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
	"unicode"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// Case conversion functions
func toSentenceCase(text string) string {
	if len(text) == 0 {
		return text
	}
	runes := []rune(text)
	runes[0] = unicode.ToUpper(runes[0])
	for i := 1; i < len(runes); i++ {
		runes[i] = unicode.ToLower(runes[i])
	}
	return string(runes)
}

func toLowerCase(text string) string {
	return strings.ToLower(text)
}

func toUpperCase(text string) string {
	return strings.ToUpper(text)
}

func toCapitalizedCase(text string) string {
	words := strings.Fields(text)
	for i, word := range words {
		if len(word) > 0 {
			runes := []rune(word)
			runes[0] = unicode.ToUpper(runes[0])
			for j := 1; j < len(runes); j++ {
				runes[j] = unicode.ToLower(runes[j])
			}
			words[i] = string(runes)
		}
	}
	return strings.Join(words, " ")
}

func toAlternatingCase(text string) string {
	runes := []rune(text)
	for i := 0; i < len(runes); i++ {
		if i%2 == 0 {
			runes[i] = unicode.ToUpper(runes[i])
		} else {
			runes[i] = unicode.ToLower(runes[i])
		}
	}
	return string(runes)
}

func toTitleCase(text string) string {
	caser := cases.Title(language.Und)
	return caser.String(strings.ToLower(text))
}

func toInverseCase(text string) string {
	runes := []rune(text)
	for i, ch := range runes {
		if unicode.IsUpper(ch) {
			runes[i] = unicode.ToLower(ch)
		} else if unicode.IsLower(ch) {
			runes[i] = unicode.ToUpper(ch)
		}
	}
	return string(runes)
}

// Caesar cipher function
func cipher(text string, shift int) string {
	shift = shift % 26
	runes := []rune(text)
	for i, ch := range runes {
		if ch >= 'A' && ch <= 'Z' {
			runes[i] = 'A' + ((ch - 'A' + int32(shift) + 26) % 26)
		} else if ch >= 'a' && ch <= 'z' {
			runes[i] = 'a' + ((ch - 'a' + int32(shift) + 26) % 26)
		}
	}
	return string(runes)
}

func main() {
	var input, transform string
	var stdin bool
	var shift int

	flag.StringVar(&input, "input", "", "Input file")
	flag.BoolVar(&stdin, "stdin", false, "Read from stdin")
	flag.StringVar(&transform, "transform", "", "Transformation type (sentence, lower, upper, capitalized, alternating, title, inverse, cipher)")
	flag.IntVar(&shift, "shift", 0, "Shift value for cipher (if transform is cipher)")
	flag.Parse()

	var content string
	if stdin {
		bytes, err := io.ReadAll(os.Stdin)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to read from stdin: %v\n", err)
			os.Exit(1)
		}
		content = string(bytes)
	} else {
		fileContent, err := os.ReadFile(input)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to read input file: %v\n", err)
			os.Exit(1)
		}
		content = string(fileContent)
	}

	var transformedContent string
	switch transform {
	case "sentence":
		transformedContent = toSentenceCase(content)
	case "lower":
		transformedContent = toLowerCase(content)
	case "upper":
		transformedContent = toUpperCase(content)
	case "capitalized":
		transformedContent = toCapitalizedCase(content)
	case "alternating":
		transformedContent = toAlternatingCase(content)
	case "title":
		transformedContent = toTitleCase(content)
	case "inverse":
		transformedContent = toInverseCase(content)
	case "cipher":
		transformedContent = cipher(content, shift)
	default:
		fmt.Fprintln(os.Stderr, "Invalid transformation type!")
		os.Exit(1)
	}

	fmt.Println(transformedContent)
}
