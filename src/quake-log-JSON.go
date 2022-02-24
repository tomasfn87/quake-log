package quakeLog

import (
	"encoding/json"
	"fmt"
)

func (qlf QuakeLogFile) PrintJsonQuakeLogs() {
	quakeGameLogFile := NewQuakeLogFile(QuakeLogFile{Games: qlf.OpenQuakeLog()})
	json, err := json.Marshal(quakeGameLogFile.Games)
	PanicIf(err)
	fmt.Printf("%s\n", json)
}

func (qlf QuakeLogFile) PrintIndentedJsonQuakeLogs(indentation string) {
	quakeGameLogFile := NewQuakeLogFile(QuakeLogFile{Games: qlf.OpenQuakeLog()})
	indentedJson, err := json.MarshalIndent(quakeGameLogFile.Games, "", indentation)
	PanicIf(err)
	fmt.Printf("%s\n", indentedJson)
}

func (qlf QuakeLogFile) PrintSingleQuakeGameLog(game int) {
	quakeGameLogFile := NewQuakeLogFile(QuakeLogFile{Games: qlf.OpenQuakeLog()})
	if game < 0 {
		game = 0
	} else if game >= len(quakeGameLogFile.Games) {
		game = len(quakeGameLogFile.Games) - 1
	}
	indentedJson, err := json.MarshalIndent(quakeGameLogFile.Games[game:game+1], "", "    ")
	PanicIf(err)
	fmt.Printf("%s\n", indentedJson)

}
