package main

import (
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

	output, err := os.Create(fileDir + fileName + ".py")
	if err != nil {
		log.Fatal(err)
	}

	_, err = output.WriteString(code)
	if err != nil {
		log.Fatal(err)
	}

	cmd := exec.Command("python", fileDir+fileName+".py")

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}
