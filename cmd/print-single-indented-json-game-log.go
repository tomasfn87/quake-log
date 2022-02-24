package main

import (
	"fmt"
	"os"
	"strconv"

	quakeLog "code.repo/quakeLog/src"
)

func main() {
	file := quakeLog.NewQuakeLogFile(quakeLog.QuakeLogFile{Path: "./Quake.txt"})
	input := os.Args

	var gameInput string

	if len(input) >= 2 {
		gameInput = input[1]
	} else {
		gameInput = "0"
	}

	game, err := strconv.Atoi(gameInput)
	quakeLog.PanicIf(err)
	game--

	fmt.Println("Single Indented JSON Game Log:")
	file.PrintSingleQuakeGameLog(game)
}
