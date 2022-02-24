package main

import (
	"fmt"

	quakeLog "code.repo/quakeLog/src"
)

func main() {
	file := quakeLog.NewQuakeLogFile(quakeLog.QuakeLogFile{Path: "./Quake.txt"})

	fmt.Println("GO Structs:")
	quakeGamesLog := file.OpenQuakeLog()
	for _, v := range quakeGamesLog {
		fmt.Println(v)
	}
}
