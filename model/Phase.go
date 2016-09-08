package model

import (
	"fmt"

	"../sdk"
)

type Phase struct {
	*BaseModel
	Month   int
	Year    int
	AccGoal int
	Goal    float64
	Tasks   PhaseTasks
}

func (a Phase) GoalInt() int {
	return sdk.FloatTimesInt(a.Goal, a.Scale())
}

func (a Phase) Print() {
	fmt.Printf("Phase => Month: %d - Year: %d - Goal: %8.2f\n", a.Month, a.Year, a.Goal)
	a.Tasks.Print()
}

type Phases []Phase

func (a Phases) SumFirstNGoals(n int) float64 {
	var total = 0.0
	for index := 0; index < n; index++ {
		total += a[index].Goal
	}
	return total
}

func (a Phases) Print(qtyPhases *int) {
	var size = len(a)
	if qtyPhases != nil {
		size = *qtyPhases
	}
	for index := 0; index < size; index++ {
		fmt.Println()
		a[index].Print()
	}
}
