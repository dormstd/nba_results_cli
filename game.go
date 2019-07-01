package main

import (
	"fmt"

	"github.com/fatih/color"
)

type scoreBoardResult struct {
	Resource   string `json:"resource"`
	Parameters struct {
		GameDate  string `json:"GameDate"`
		LeagueID  string `json:"LeagueID"`
		DayOffset string `json:"DayOffset"`
	} `json:"parameters"`
	ResultSets []struct {
		Name    string          `json:"name"`
		Headers []string        `json:"headers"`
		RowSet  [][]interface{} `json:"rowSet"`
	} `json:"resultSets"`
}
type gameStruct struct {
	GameID      string
	StatusText  string
	teamLocal   teamScore
	teamVisitor teamScore
}

type teamScore struct {
	teamName string
	score    int
}

func newGames(input scoreBoardResult) ([]gameStruct, error) {
	teams, err := teamScores(input)
	if err != nil {
		return nil, fmt.Errorf("error getting the teams from the json")
	}

	var games []gameStruct
	if !(len(input.ResultSets) > 0) {
		return nil, fmt.Errorf("webservice Result does not contain results")
	}
	gamesJSON := input.ResultSets[0].RowSet
	//For each of the games stored in the slice
	for i := range gamesJSON {
		var game gameStruct

		//ToDo: This is not clean and open to errors, look for a way to fix it
		g := gamesJSON[i]
		game.GameID = g[2].(string)
		game.StatusText = g[4].(string)
		game.teamLocal = teams[int(g[6].(float64))]
		game.teamVisitor = teams[int(g[7].(float64))]

		games = append(games, game)
	}
	return games, nil
}

func teamScores(input scoreBoardResult) (map[int]teamScore, error) {

	m := make(map[int]teamScore)
	if !(len(input.ResultSets) > 0) {
		return nil, fmt.Errorf("webservice Result does not contain results")
	}
	teamsJSON := input.ResultSets[1].RowSet
	//For each of the teams score stored in the slice
	for i := range teamsJSON {
		//ToDo: This is not clean and open to errors, look for a way to fix it
		t := teamsJSON[i]

		ID := int(t[3].(float64))
		name := t[5].(string)
		score := int(t[21].(float64))

		m[ID] = teamScore{name, score}
	}
	return m, nil
}

func printResults(games []gameStruct) {
	// Miami 119 - 109 Philadelphia

	green := color.New(color.FgGreen).SprintFunc()
	red := color.New(color.FgRed).SprintFunc()
	white := color.New(color.FgWhite).SprintFunc()

	for _, game := range games {
		localScoreNum := game.teamLocal.score
		visitorScoreNum := game.teamVisitor.score
		localScore := white(game.teamLocal.score)
		visitorScore := white(game.teamVisitor.score)
		if localScoreNum > visitorScoreNum {
			localScore = green(localScoreNum)
			visitorScore = red(visitorScoreNum)
		}
		if localScoreNum < visitorScoreNum {
			localScore = red(localScoreNum)

			visitorScore = green(visitorScoreNum)
		}
		fmt.Printf("%-14s %4s : %4s %14s \n", game.teamLocal.teamName, localScore, visitorScore, game.teamVisitor.teamName)
	}
}
