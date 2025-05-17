package leet

import (
	"container/heap"
	"fmt"
)

func (t *Teams) Push(x any) {
	*t = append(*t, x.(Team))
}

func (t *Teams) Pop() any {
	old := *t
	n := len(old)
	x := old[n-1]
	*t = old[0 : n-1]
	return x
}

func (t *Teams) Peek() *Team {
	old := *t
	n := len(old)
	if n == 0 {
		return nil
	}
	return &old[0]
}

func (t *Teams) HelperPush(team Team) {
	heap.Push(t, team)
}
func (t *Teams) HelperPop() *Team {
	if len(*t) == 0 {
		return nil
	}
	x := heap.Pop(t)
	out := x.(Team)
	return &out
}

func NewTeamsHeap(teamsIn []Team) *Teams {
	teams := Teams(teamsIn)
	x := &teams
	heap.Init(x)
	return x
}

func NewTeamsHeapByWin(teamsIn []Team) *MaxWinTeamHeap {
	maxWinTH := MaxWinTeamHeap{teamsIn}
	x := &maxWinTH
	heap.Init(x)
	return x
}

// Priority version of this heap, max heap
type MaxWinTeamHeap struct {
	Teams
}

func (mth MaxWinTeamHeap) Less(i, j int) bool {
	fmt.Println("Max win HEAP sorting!")
	return mth.Teams[j].wins < mth.Teams[i].wins
}

func (mth *MaxWinTeamHeap) HelperPopMaxWin() *Team {
	if len(mth.Teams) == 0 {
		return nil
	}
	x := heap.Pop(mth)
	teamPopped := x.(Team)
	return &teamPopped
}
