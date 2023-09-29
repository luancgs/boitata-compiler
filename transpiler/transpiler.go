package transpiler

import (
	"encoding/json"
	"io"
	"log"
	"os"
	"regexp"
	"strings"
)

type Token struct {
	Btt string `json:"btt"`
	Py  string `json:"py"`
}

const placeholder = "${LITERAL_PLACEHOLDER}$"

func Transpile(code string) string {
	file, err := os.Open("tokens.json")
	if err != nil {
		log.Fatal("Error opening file:", err)
	}
	defer file.Close()

	byteValue, err := io.ReadAll(file)
	if err != nil {
		log.Fatal("Error reading file:", err)
	}

	var tokens []Token

	err = json.Unmarshal(byteValue, &tokens)
	if err != nil {
		log.Fatal("Error unmarshaling JSON:", err)
	}

	// Replace and remove the double-quoted strings from the original text
	code, literals := RemoveStringLiterals(code)

	for _, token := range tokens {
		code = strings.ReplaceAll(code, token.Btt, token.Py)
	}

	// Return the double-quoted strings to the original text
	code = ReturnStringLiterals(code, literals)

	return code
}

func RemoveStringLiterals(code string) (string, []string) {
	regex := regexp.MustCompile(`"[^"]*"`)
	var literals []string

	code = regex.ReplaceAllStringFunc(code, func(matched string) string {
		literals = append(literals, matched)
		return placeholder
	})

	return code, literals
}

func ReturnStringLiterals(code string, literals []string) string {
	for _, str := range literals {
		code = strings.Replace(code, placeholder, str, 1)
	}

	return code
}
