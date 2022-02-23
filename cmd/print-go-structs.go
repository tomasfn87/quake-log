package main

import (
	"fmt"

	quakeLog "code.repo/quakeLog/src"
)

func main() {
	file := quakeLog.NewQuakeLogFile()
	file.Path = "./Quake.txt"

	fmt.Println("GO Structs:")
	quakeGamesLog := file.OpenQuakeLog()
	for _, v := range quakeGamesLog {
		fmt.Println(v)
	}
}
