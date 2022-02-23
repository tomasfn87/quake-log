package quakeLog

import (
	"encoding/json"
	"fmt"
)

func (ql QuakeLogFile) GetJsonQuakeLogs() {
	quakeGameLogFile := NewQuakeLogFile()
	quakeGameLogFile.Games = ql.OpenQuakeLog()

	fmt.Println("\nIndented JSON Data:")
	json, err := json.Marshal(quakeGameLogFile.Games)
	PanicIf(err)

	fmt.Printf("%s\n", json)
}

func (ql QuakeLogFile) GetIndentedJsonQuakeLogs(indentation string) {
	quakeGameLogFile := NewQuakeLogFile()
	quakeGameLogFile.Games = ql.OpenQuakeLog()

	indentedJson, err := json.MarshalIndent(quakeGameLogFile.Games, "", indentation)
	PanicIf(err)

	fmt.Printf("%s\n", indentedJson)
}
