package main

import (
	"fmt"

	quakeLog "code.repo/quakeLog/src"
)

func fillWithTimes(char string, size int) string {
	fill := ""
	for i := 0; i < size; i++ {
		fill += fmt.Sprintf("%c", char[0])
	}
	return fill
}

func main() {
	quakeLogFilePath := "./Quake.txt"
	file := quakeLog.QuakeLogFile{Path: quakeLogFilePath}

	file.OpenQuakeLog()
}
