package quakeLog

import (
	"encoding/json"
	"fmt"
)

func (ql QuakeLogFile) PrintJsonQuakeLogs() {
	quakeGameLogFile := NewQuakeLogFile()
	quakeGameLogFile.Games = ql.OpenQuakeLog()

	json, err := json.Marshal(quakeGameLogFile.Games)
	PanicIf(err)
	fmt.Printf("%s\n", json)
}

func (ql QuakeLogFile) PrintIndentedJsonQuakeLogs(indentation string) {
	quakeGameLogFile := NewQuakeLogFile()
	quakeGameLogFile.Games = ql.OpenQuakeLog()

	indentedJson, err := json.MarshalIndent(quakeGameLogFile.Games, "", indentation)
	PanicIf(err)
	fmt.Printf("%s\n", indentedJson)
}
