package leet

import "fmt"

// implement sort.Interface on Teams
type Teams []Team

func (a Teams) Len() int {
	return len(a)
}
func (a Teams) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}
func (a Teams) Less(i, j int) bool {
	// Compare the teams by team names
	fmt.Println("Running by name less")
	return a[i].name < a[j].name
}

type ByWinsProper struct {
	Teams
}

func (a ByWinsProper) Less(i, j int) bool {
	// Compare the teams by wins
	fmt.Println("Running byWinsProper by wins less")
	return a.Teams[i].wins < a.Teams[j].wins
}

type ByAllFactorsProper struct {
	Teams
}

func (a ByAllFactorsProper) Less(i, j int) bool {
	fmt.Println("Running byAllFactorsProper by all factors less")
	// Compare the teams by all factors
	factor1i := a.Teams[i].wins*3 + a.Teams[i].draws*1
	factor1j := a.Teams[j].wins*3 + a.Teams[j].draws*1
	factor2i := a.Teams[i].pointsGained - a.Teams[i].pointsConceded
	factor2j := a.Teams[j].pointsGained - a.Teams[j].pointsConceded

	fmt.Printf(" - TRY factor1i == factor1j ? %d == %d\n", factor1i, factor1j)
	switch {
	case factor1i < factor1j:
		return true
	case factor1i == factor1j && factor2i < factor2j:
		fmt.Printf(" -- Equals Match Team %s and Team %s\n", a.Teams[i].name, a.Teams[j].name)
		fmt.Printf(" -- factor1i == factor1j && factor2i < factor2j: %d < %d\n", factor2i, factor2j)
		return true
	}
	if factor1i == factor1j && factor2i > factor2j {
		fmt.Printf(" -- Equals match Team %s and Team %s\n", a.Teams[i].name, a.Teams[j].name)
		fmt.Printf(" -- factor1i == factor1j && factor2i > factor2j: %d > %d\n", factor2i, factor2j)
	}
	return false
}
