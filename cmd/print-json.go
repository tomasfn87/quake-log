package main

import (
	"fmt"

	quakeLog "code.repo/quakeLog/src"
)

func main() {
	file := quakeLog.NewQuakeLogFile(quakeLog.QuakeLogFile{Path: "./Quake.txt"})

	fmt.Println("JSON String:")
	file.PrintJsonQuakeLogs()
}
