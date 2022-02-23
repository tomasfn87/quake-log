package main

import (
	"fmt"
	"os"

	quakeLog "code.repo/quakeLog/src"
)

func main() {
	file := quakeLog.NewQuakeLogFile()
	file.Path = "./Quake.txt"

	indentation := ",.,."
	indentationInput := os.Args
	if len(indentationInput) > 1 {
		indentation = indentationInput[1]
	}
	fmt.Println("\nIndented JSON Data:")
	file.GetIndentedJsonQuakeLogs(fillWithTimes(" ", len(indentation)))
}

func fillWithTimes(char string, size int) string {
	fill := ""
	for i := 0; i < size; i++ {
		fill += fmt.Sprintf("%c", char[0])
	}
	return fill
}
