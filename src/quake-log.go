package quakeLog

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type QuakeLogFile struct {
	Path  string
	Games []QuakeGameLog `json:"quake_game_logs"`
}

type QuakeGameLog struct {
	Game   int    `json:"game"`
	Status Status `json:"status"`
}

type Status struct {
	TotalKills int      `json:"total_kills"`
	Players    []Player `json:"players"`
}

type Player struct {
	Id       int      `json:"id"`
	Nome     string   `json:"nome"`
	Kills    int      `json:"kills"`
	OldNames []string `json:"old_names"`
}

func NewQuakeGameLog() *QuakeGameLog {
	return &QuakeGameLog{}
}

func NewPlayer(p Player) *Player {
	return &Player{Nome: p.Nome, Id: p.Id}
}

func NewQuakeLogFile() *QuakeLogFile {
	return &QuakeLogFile{}
}

func (ql QuakeLogFile) OpenQuakeLog() []QuakeGameLog {
	file, err := os.Open(ql.Path)
	PanicIf(err)
	defer file.Close()

	quakeLogFileLines := getDataFromFileLines(file)
	quakeGames := []QuakeGameLog{}
	quakeGames = parseDataFromFileLines(quakeLogFileLines, quakeGames)

	return quakeGames
}

func PanicIf(err error) {
	if err != nil {
		panic(err)
	}
}

func parseDataFromFileLines(logLines []string, qgl []QuakeGameLog) []QuakeGameLog {
	var gameCount int
	for _, v := range logLines {
		lineContent := strings.Split(v, " ")
		for _, y := range lineContent {
			if y == "InitGame:" {
				qgl = append(qgl, QuakeGameLog{Game: gameCount + 1})
			} else if y == "ShutdownGame:" {
				gameCount++
			} else if y == "ClientUserinfoChanged:" {
				var clientUserinfoChangedRE *regexp.Regexp
				clientUserinfoChangedRE, err := regexp.Compile(`^(.*) (\d+)( n\\)(.+)(\\t\\)(.*)$`)
				PanicIf(err)

				clientUserinfoChanged := strings.Split(clientUserinfoChangedRE.ReplaceAllString(v, `$2   :   $4`), "   :   ")
				id, err := strconv.Atoi(clientUserinfoChanged[0])
				PanicIf(err)
				id--
				nome := clientUserinfoChanged[1]

				if playerListContainsId(qgl[gameCount].Status.Players, id) {
					pl := getPlayerIndexByIdFromPlayerList(qgl[gameCount].Status.Players, id)
					qgl[gameCount].Status.Players[pl].OldNames = []string{}

					if !playerListContainsNome(qgl[gameCount].Status.Players, nome) {
						if !oldNamesContainsNome(qgl[gameCount].Status.Players[pl].OldNames, qgl[gameCount].Status.Players[pl].Nome) {
							qgl[gameCount].Status.Players[pl].OldNames =
								append(qgl[gameCount].Status.Players[pl].OldNames, qgl[gameCount].Status.Players[pl].Nome)
						}
						qgl[gameCount].Status.Players[pl].Nome = nome
					}

					if !playerListContainsNome(qgl[gameCount].Status.Players, nome) {
						qgl[gameCount].Status.Players[pl].Nome = nome
					}

					qgl[gameCount].Status.Players[pl].OldNames =
						cleanRepeatedOldNames(qgl[gameCount].Status.Players[pl].OldNames, qgl[gameCount].Status.Players[pl].Nome)
				} else {
					newPlayer := NewPlayer(Player{Nome: nome, Id: id})
					newPlayer.OldNames = []string{}
					qgl[gameCount].Status.Players = append(qgl[gameCount].Status.Players, *newPlayer)
				}
			} else if y == "Kill:" {
				qgl[gameCount].Status.TotalKills++

				var killRE *regexp.Regexp
				killRE, err := regexp.Compile(`^(.*)( Kill: )(\d+) (\d+) (\d+:) (.+) killed (.+) by (.*)$`)
				PanicIf(err)

				kill := strings.Split(killRE.ReplaceAllString(v, `$6   :   $7   :   $8`), "   :   ")
				killer, victim := kill[0], kill[1]

				// "cause" stores the source weapon / "<world>" inflinged cause of death
				// cause := kill[2]
				// fmt.Printf("v{%s} k{%s} c{%s}\n", victim, killer, cause)
				for k, p := range qgl[gameCount].Status.Players {
					if killer != "<world>" && killer == p.Nome {
						qgl[gameCount].Status.Players[k].Kills++
					}
				}
				if killer == "<world>" {
					for k, p := range qgl[gameCount].Status.Players {
						if p.Nome == victim {
							qgl[gameCount].Status.Players[k].Kills--
						}
					}
				}
			}
		}
	}
	return qgl
}

func getDataFromFileLines(file *os.File) []string {
	scanner := bufio.NewScanner(file)
	fileLines := []string{}
	for scanner.Scan() {
		fileLines = append(fileLines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		PanicIf(err)
	}
	return fileLines
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
	var playerIndex int
	for k, v := range arr {
		if v.Id == id {
			playerIndex = k
		}
	}
	return playerIndex
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
