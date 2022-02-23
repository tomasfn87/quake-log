package quakeLog

import (
	"encoding/json"
	"fmt"
)

func (ql QuakeLogFile) GetJson() {
	quakeGameLog := ql.OpenQuakeLog()
	/*for _, v := range quakeGameLog {
		fmt.Println(v)
	}
	fmt.Printf("type of var quakeGames: %T\n", quakeGameLog)
	*/
	jsonArray := make([]string, 0, len(quakeGameLog))
	for i := 0; i < len(quakeGameLog); i++ {
		json, err := json.Marshal(quakeGameLog[i])
		PanicIf(err)
		if i == len(quakeGameLog)-1 {
			jsonArray = append(jsonArray, string(json))
		} else {
			jsonArray = append(jsonArray, fmt.Sprintf("%s,", json))
		}
	}
	fmt.Println(jsonArray)
}
