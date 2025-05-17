package main

import (
	"fmt"
	"sort"

	"github.com/robododge/ragserver-fun/leet"
)

func main() {

	proveSwitchCaseWithFuncsCase()
	fmt.Println("===============SORTING FIRST TRY!========================")
	sortTeamsWithFirstTryAbstraction()
	fmt.Println("================SORTING PROPER!=======================")
	sortTeamsWithProperAbstraction()

}

// Closed to the sort wrapper approach from pkg.go.dev https://pkg.go.dev/sort@go1.24.2#example-package-SortWrapper
func sortTeamsWithProperAbstraction() {
	teams := leet.Teams(leet.MakeTeamData())

	// Sort the teams by wins
	sort.Sort(leet.ByWinsProper{teams})

	// Print the sorted teams
	fmt.Println("Teams sorted by wins:")
	for i, team := range teams {
		fmt.Printf("%d: %s\n", i+1, team)
	}

	// Sort the teams by all factors in reverse order
	sort.Sort(sort.Reverse(leet.ByAllFactorsProper{teams}))
	// Print the sorted by all factors
	fmt.Println("Teams sorted by all factors:")
	for i, team := range teams {
		fmt.Printf("%d: %s\n", i+1, team)
	}
}

func sortTeamsWithFirstTryAbstraction() {
	teams := leet.MakeTeamData()

	// Sort the teams by wins
	sort.Sort(leet.ByWins(teams))

	// Print the sorted teams
	fmt.Println("Teams sorted by wins:")
	for i, team := range teams {
		fmt.Printf("%d: %s\n", i+1, team)
	}

	// Sort the teams by all factors in reverse order
	//MESSY there is a BUG!  we are sorting using the ByWins interface not the ByAllFactors interface
	sort.Sort(sort.Reverse(leet.ByWins(leet.ByAllFactors(teams))))

	fmt.Println("BUG! here! we are sorting using the ByWins interface not the ByAllFactors interface")
	fmt.Println("BUG! here! we are sorting using the ByWins interface not the ByAllFactors interface")
	// Print the sorted by all factors
	fmt.Println("Teams sorted by all factors:")
	for i, team := range teams {
		fmt.Printf("%d: %s\n", i+1, team)
	}
}

func proveSwitchCaseWithFuncsCase() {
	numsA := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	numsB := []int{1, 3, 5, 7, 9, 11, 13, 15, 17}

	for i := 0; i < len(numsA); i++ {
		switch {
		case iieven(numsA[i], 'A'):
			fmt.Printf("Casey A: %d is even\n", numsA[i])
		case iieven(numsB[i], 'B'):
			fmt.Printf("Casey B: %d is even\n", numsB[i])
		}
	}
}

func iieven(i int, name rune) bool {
	fmt.Printf("Calling casey %v with %d\n", string(name), i)
	return i&1 == 0
}
