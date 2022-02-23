package quakeLog

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// ClientUserinfoChanged // this func must return the player name

type QuakeLogFile struct {
	Path string
}

type QuakeGameLog struct {
	Game   int
	Status Status
}

type Status struct {
	TotalKills int
	Players    Players
}

type Players struct {
	List []Player
}

type Player struct {
	Id       int
	Nome     string
	Kills    int
	OldNames []string
}

func NewQuakeLog() *QuakeGameLog {
	return &QuakeGameLog{}
}

func (qlf QuakeLogFile) OpenQuakeLog() {
	file, err := os.Open(qlf.Path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	logLines := getFileValues(file)

	quakeGames := []*QuakeGameLog{}

	var clientUserinfoChangedRE *regexp.Regexp
	clientUserinfoChangedRE, err = regexp.Compile(`^(.*) (\d+)( n\\)(.+)(\\t\\)(.*)$`)
	panicIf(err)

	var killRE *regexp.Regexp
	killRE, err = regexp.Compile(`^(.*)( Kill: )(\d+) (\d+) (\d+:) (.+) killed (.+) by (.*)$`)
	panicIf(err)

	gameCount := 0
	for _, v := range logLines {
		lineContent := strings.Split(v, " ")
		for _, y := range lineContent {
			switch y {
			case "InitGame:":
				{
					quakeGames = append(quakeGames, &QuakeGameLog{Game: gameCount + 1})
				}
			case "ShutdownGame:":
				{
					gameCount++
				}
			case "ClientUserinfoChanged:":
				{
					clientUserinfoChanged := strings.Split(clientUserinfoChangedRE.ReplaceAllString(v, `$2   :   $4`), "   :   ")
					id, err := strconv.Atoi(clientUserinfoChanged[0])
					panicIf(err)
					id--
					nome := clientUserinfoChanged[1]

					if playerListContainsId(quakeGames[gameCount].Status.Players.List, id) {
						pl := getPlayerIndexByIdFromPlayerList(quakeGames[gameCount].Status.Players.List, id)

						if !playerListContainsNome(quakeGames[gameCount].Status.Players.List, nome) {
							if !oldNamesContainsNome(quakeGames[gameCount].Status.Players.List[pl].OldNames, quakeGames[gameCount].Status.Players.List[pl].Nome) {
								quakeGames[gameCount].Status.Players.List[pl].OldNames =
									append(quakeGames[gameCount].Status.Players.List[pl].OldNames, quakeGames[gameCount].Status.Players.List[pl].Nome)
							}
							quakeGames[gameCount].Status.Players.List[pl].Nome = nome
						}
						if !playerListContainsNome(quakeGames[gameCount].Status.Players.List, nome) {
							quakeGames[gameCount].Status.Players.List[pl].Nome = nome
						}
						quakeGames[gameCount].Status.Players.List[pl].OldNames =
							cleanRepeatedOldNames(quakeGames[gameCount].Status.Players.List[pl].OldNames, quakeGames[gameCount].Status.Players.List[pl].Nome)
					} else {
						quakeGames[gameCount].Status.Players.List = append(quakeGames[gameCount].Status.Players.List, Player{Nome: nome, Id: id})
					}

					// fmt.Printf("id{%d} type:%T, nome{%s} type:%T\n", id, id, nome, nome)

				}
			case "Kill:":
				{
					quakeGames[gameCount].Status.TotalKills++
					kill := strings.Split(killRE.ReplaceAllString(v, `$6   :   $7   :   $8`), "   :   ")
					killer, victim := kill[0], kill[1]
					//cause := kill[2]
					// fmt.Printf("v{%s} k{%s} c{%s}\n", victim, killer, cause)
					for k, p := range quakeGames[gameCount].Status.Players.List {
						if killer != "<world>" && killer == p.Nome {
							quakeGames[gameCount].Status.Players.List[k].Kills++
						}
					}
					if killer == "<world>" {
						for k, p := range quakeGames[gameCount].Status.Players.List {
							if p.Nome == victim {
								quakeGames[gameCount].Status.Players.List[k].Kills--
							}

						}
					}
				}
			}
		}
	}

	for _, v := range quakeGames {
		fmt.Println(v)
	}

}

func getFileValues(file *os.File) []string {
	scanner := bufio.NewScanner(file)
	logLines := []string{}
	for scanner.Scan() {
		logLines = append(logLines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		panicIf(err)
	}
	return logLines
}

func panicIf(err error) {
	if err != nil {
		panic(err)
	}
}

func playerListContainsId(arr []Player, item int) bool {
	if len(arr) == 0 {
		return false
	}
	for _, v := range arr {
		if v.Id == item {
			return true
		}
	}
	return false
}

func playerListContainsNome(arr []Player, item string) bool {
	if len(arr) == 0 {
		return false
	}
	for _, v := range arr {
		if v.Nome == item {
			return true
		}
	}
	return false
}

func getPlayerIndexByIdFromPlayerList(arr []Player, id int) int {
	var index int
	for k, v := range arr {
		if v.Id == id {
			index = k
		}
	}
	return index
}

func oldNamesContainsNome(arr []string, nome string) bool {
	if len(arr) == 0 {
		return false
	}
	for _, v := range arr {
		if v == nome {
			return true
		}
	}
	return false
}

func cleanRepeatedOldNames(arr []string, currentName string) []string {
	cleanedOldNames := []string{}
	for _, v := range arr {
		if v != currentName {
			cleanedOldNames = append(cleanedOldNames, v)
		}
	}
	return cleanedOldNames
}
