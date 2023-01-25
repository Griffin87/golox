package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func main() {
	if len(os.Args) > 2 {
		fmt.Println("Usage: lox [script]")
		os.Exit(64)
	} else if len(os.Args) == 1 {
		runFile(os.Args[0])
	} else {
		runPrompt()
	}
}

func runFile(path string) int {
	data, err := io.ReadAll(path)

	if err != nil {
		panic(error)
	}

	run(string(data))
	if hadError {
		os.Exit(65)
	}
	}

}

// interactive prompt
func runPrompt() int {
	for {
		fmt.Println("> ")
		data, err := reader
	}
}

func run(source string) {
	scanner := Scanner.new(source)
	tokens := scanner.scanTokens()

	for _, token range := tokens {
		os.Println(token)
	}
}

func error(line int, err_msg string) {
	report(line, "", err_msg)
}

func report(line int, where string, msg string) {
	fmt.Println("[line " + line + "] Error" + where ": " + message)
	hadError = true
}