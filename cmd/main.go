package main

import (
	"fmt"

	quakeLog "code.repo/quakeLog/src"
)

func main() {
	quakeLogFilePath := "./Quake.txt"
	file := quakeLog.QuakeLogFile{Path: quakeLogFilePath}

	quakeGamesLog := file.OpenQuakeLog()
	for _, v := range quakeGamesLog {
		fmt.Println(v)
	}
}
