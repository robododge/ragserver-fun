package main

import (
	"fmt"

	"github.com/robododge/ragserver-fun/leet"
)

func main() {
	teamHeapTrial()
	teamHeapTrialByWins()
}

func teamHeapTrial() {
	teams := generateTeams()
	teamsHeap := leet.NewTeamsHeap(teams)

	topTeam := teamsHeap.Peek()
	fmt.Println("Found team from the heap: ", topTeam)

	i := 0
	for len(*teamsHeap) > 0 {
		teamPopped := teamsHeap.HelperPop()
		fmt.Printf(" Popped team[%d]: %v\n", i, teamPopped)
		i++
	}
}

func teamHeapTrialByWins() {
	teams := generateTeams()
	teamsHeap := leet.NewTeamsHeapByWin(teams)

	topWinTeam := teamsHeap.Peek()
	fmt.Println("Found top winning team: ", topWinTeam)

	i := 0
	for len(teamsHeap.Teams) > 0 {
		teamPopped := teamsHeap.HelperPopMaxWin()
		fmt.Printf(" Popped team[%d]: %v\n", i, teamPopped)
		i++
	}

}

func generateTeams() []leet.Team {
	return leet.MakeTeamData()
}
