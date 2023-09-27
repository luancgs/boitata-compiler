package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func main() {

	if len(os.Args) < 2 {
		log.Fatal("O caminho do arquivo .btt não foi informado")
	}

	if _, err := os.Stat(os.Args[1]); os.IsNotExist(err) {
		log.Fatal("O arquivo informado não existe")
	}

	if !strings.HasSuffix(os.Args[1], ".btt") {
		log.Fatal("O arquivo informado não é um arquivo .btt")
	}

	filePath := os.Args[1]
	fileDir, fileName := filepath.Split(filePath)
	fileName = fileName[:len(fileName)-4]

	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	code := Transpile(string(data))

	tmpFile, err := os.CreateTemp(fileDir, fmt.Sprintf("boitata-%s-*.py", fileName))
	if err != nil {
		log.Fatal("Error creating temporary Python file:", err)
	}
	defer os.Remove(tmpFile.Name())

	_, err = tmpFile.WriteString(code)
	if err != nil {
		log.Fatal("Error writing to temporary Python file:", err)
	}

	tmpFile.Close()

	cmd := exec.Command("python", tmpFile.Name())

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}
