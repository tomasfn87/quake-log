package main

import (
	"fmt"
	"os"

	quakeLog "code.repo/quakeLog/src"
)

func main() {
	file := quakeLog.NewQuakeLogFile(quakeLog.QuakeLogFile{Path: "./Quake.txt"})

	indentation := ",.,."
	char := " "
	input := os.Args

	if len(input) >= 2 {
		indentation = input[1]
	}

	if len(input) >= 3 {
		char = input[2]
	}

	fmt.Println("Indented JSON Data:")
	file.PrintIndentedJsonQuakeLogs(fillWithTimes(char, len(indentation)))
}

func fillWithTimes(char string, size int) string {
	fill := ""
	for i := 0; i < size; i++ {
		fill += fmt.Sprintf("%c", char[0])
	}
	return fill
}
