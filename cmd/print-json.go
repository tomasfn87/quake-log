package main

import (
	"fmt"

	quakeLog "code.repo/quakeLog/src"
)

func main() {
	file := quakeLog.NewQuakeLogFile()
	file.Path = "./Quake.txt"

	fmt.Println("\nJSON String:")
	file.GetJsonQuakeLogs()
}
