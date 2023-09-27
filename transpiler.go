package main

import (
	"encoding/json"
	"io"
	"log"
	"os"
	"strings"
)

type Token struct {
	Btt string `json:"btt"`
	Py  string `json:"py"`
}

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

	for _, token := range tokens {
		code = strings.ReplaceAll(code, token.Btt, token.Py)
	}

	return code
}
