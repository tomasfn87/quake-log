package quakeLog

import (
	"encoding/json"
	"fmt"
)

func (ql QuakeLogFile) GetJsonQuakeLogs() {
	quakeGameLog := ql.OpenQuakeLog()
	/*for _, v := range quakeGameLog {
		fmt.Println(v)
	}
	fmt.Printf("type of var quakeGames: %T\n", quakeGameLog)
	*/
	jsonString := "["
	for i := 0; i < len(quakeGameLog); i++ {
		json, err := json.Marshal(quakeGameLog[i])
		PanicIf(err)
		if i == len(quakeGameLog) {
			jsonString += string(json)
		} else {
			jsonString += fmt.Sprintf("%s,", json)
		}
	}
	jsonString += "]"
	fmt.Println(jsonString)
}
