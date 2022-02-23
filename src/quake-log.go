package quakeLog

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
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

	quakeGames := []QuakeGameLog{}

	var ClientUserinfoChangedRE *regexp.Regexp
	ClientUserinfoChangedRE, err = regexp.Compile(`(.*)( n\\)(.+)(\\t\\)(.*)`)

	if err != nil {
		panic(err)
	}

	gameCount := 0
	totalKillCount := 0
	for _, v := range logLines {
		values := strings.Split(v, " ")
		for _, y := range values {
			switch y {
			case "InitGame:":
				{
					quakeGames = append(quakeGames, QuakeGameLog{Game: gameCount + 1})
				}
			case "ShutdownGame:":
				{
					gameCount++
				}
			case "ClientUserinfoChanged:":
				{
					fmt.Println(ClientUserinfoChangedRE.ReplaceAllString(v, `$3`))
				}
			case "Kill:":
				{
					quakeGames[gameCount].Status.TotalKills++
					totalKillCount++
				}
			}
		}
	}

	/*
		for _, v := range quakeGames {
			fmt.Println(v)
		}
		fmt.Printf("totalKillCount: %d\n", totalKillCount)
	*/
}

func getFileValues(file *os.File) []string {
	scanner := bufio.NewScanner(file)
	logLines := []string{}
	for scanner.Scan() {
		logLines = append(logLines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return logLines
}
