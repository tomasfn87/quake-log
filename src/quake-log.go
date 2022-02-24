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

func NewQuakeLogFile(qlf QuakeLogFile) *QuakeLogFile {
	return &QuakeLogFile{Path: qlf.Path, Games: qlf.Games}
}

func NewQuakeGameLog(qgl QuakeGameLog) *QuakeGameLog {
	return &QuakeGameLog{Game: qgl.Game, Status: qgl.Status}
}

func NewPlayer(p Player) *Player {
	return &Player{Nome: p.Nome, Id: p.Id, Kills: p.Kills, OldNames: p.OldNames}
}

func (qlf QuakeLogFile) OpenQuakeLog() []QuakeGameLog {
	file, err := os.Open(qlf.Path)
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

func parseDataFromFileLines(quakeLogFileLines []string, qgl []QuakeGameLog) []QuakeGameLog {
	var gameCount int
	for _, v := range quakeLogFileLines {
		lineContent := strings.Split(v, " ")
		for _, y := range lineContent {
			var clientUserinfoChangedRE *regexp.Regexp
			clientUserinfoChangedRE, err := regexp.Compile(`^(.*) (\d+)( n\\)(.+)(\\t\\)(.*)$`)
			PanicIf(err)

			var killRE *regexp.Regexp
			killRE, err = regexp.Compile(`^(.*)( Kill: )(\d+) (\d+) (\d+:) (.+) killed (.+) by (.*)$`)
			PanicIf(err)

			if y == "InitGame:" {
				// New Game has started
				newGame := NewQuakeGameLog(QuakeGameLog{Game: gameCount + 1})
				qgl = append(qgl, *newGame)
			} else if y == "ShutdownGame:" {
				// Game has ended
				gameCount++
			} else if y == "ClientUserinfoChanged:" {
				// Client may have changed name; store old names avoiding repetition (remove current name if previously stored)
				clientUserinfoChanged := strings.Split(clientUserinfoChangedRE.ReplaceAllString(v, `$2  :  $4`), "  :  ")
				id, err := strconv.Atoi(clientUserinfoChanged[0])
				PanicIf(err)
				id-- // Player IDs start at 2
				nome := clientUserinfoChanged[1]

				players := qgl[gameCount].Status.Players

				if playerListContainsId(players, id) {
					// Current ID already  exists; player might have changed names
					pl := getPlayerIndexByIdFromPlayerList(id, players)
					currentNome := &qgl[gameCount].Status.Players[pl].Nome
					oldNames := &qgl[gameCount].Status.Players[pl].OldNames

					if !playerListContainsNome(players, nome) {
						if !oldNamesContainsNome(*oldNames, *currentNome) {
							*oldNames = append(*oldNames, *currentNome)
						}
						*currentNome = nome
					}

					*oldNames = cleanRepeatedOldNames(*oldNames, *currentNome)
				} else {
					// New ID = New Player Entry
					newPlayer := NewPlayer(Player{Nome: nome, Id: id, OldNames: []string{}})
					qgl[gameCount].Status.Players = append(qgl[gameCount].Status.Players, *newPlayer)
				}
			} else if y == "Kill:" {
				// Add kills for each player kill, remove player kills if killer is "<world>"
				qgl[gameCount].Status.TotalKills++
				players := qgl[gameCount].Status.Players
				kill := strings.Split(killRE.ReplaceAllString(v, `$3  :  $4  :  $6`), "  :  ")

				killerId, err := strconv.Atoi(kill[0])
				PanicIf(err)
				killerIndex := getPlayerIndexByIdFromPlayerList(killerId, players)
				killerName := kill[2]

				victimId, err := strconv.Atoi(kill[1])
				PanicIf(err)
				victimIndex := getPlayerIndexByIdFromPlayerList(victimId, players)

				if killerId != 1022 {
					players[killerIndex].Kills++
				}

				if killerName == "<world>" {
					players[victimIndex].Kills--
				}
			}
		}
	}
	return qgl
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

func getPlayerIndexByIdFromPlayerList(id int, playerList []Player) int {
	var playerIndex int
	for k, v := range playerList {
		if v.Id == id {
			playerIndex = k
		}
	}
	return playerIndex
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
