package leet

import "fmt"

type Team struct {
	name           string
	wins           int
	draws          int
	pointsGained   int
	pointsConceded int
}

func (t Team) String() string {
	return fmt.Sprintf("Name: %s, Wins: %d, Draws: %d, Points Gained: %d, Points Conceded: %d, Point Factor: %d, Diff Gain-Conceed: %d",
		t.name, t.wins, t.draws, t.pointsGained, t.pointsConceded, t.wins*3+t.draws*1, t.pointsGained-t.pointsConceded)
}

type ByWins []Team

func (a ByWins) Len() int {
	return len(a)
}
func (a ByWins) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}
func (a ByWins) Less(i, j int) bool {
	fmt.Println("Running by wins less")
	return a[i].wins < a[j].wins
}

// type ByAllFactors struct {
// ByWins
// }
type ByAllFactors ByWins

func (a ByAllFactors) Less(i, j int) bool {
	factor1i := a[i].wins*3 + a[i].draws*1
	factor1j := a[j].wins*3 + a[j].draws*1
	factor2i := a[i].pointsGained - a[i].pointsConceded
	factor2j := a[j].pointsGained - a[j].pointsConceded

	fmt.Println("print in Less!!")
	switch {
	case factor1i < factor1j:
		return true
	case factor1i == factor1j && factor2i < factor2j:
		return true
	}
	return false
}

func MakeTeamData() []Team {
	teams := make([]Team, 5)
	names := []string{"Team A", "Team B", "Team C", "Team D", "Team E"}
	wins := []int{10, 5, 8, 10, 9}
	draws := []int{2, 3, 1, 2, 5}
	pointsGained := []int{32, 18, 25, 30, 22}
	pointsConceded := []int{20, 25, 18, 22, 2}

	for i := 0; i < len(names); i++ {
		teams[i] = Team{
			name:           names[i],
			wins:           wins[i],
			draws:          draws[i],
			pointsGained:   pointsGained[i],
			pointsConceded: pointsConceded[i],
		}
	}
	return teams

}
