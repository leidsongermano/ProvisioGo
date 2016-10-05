package model

import (
	"fmt"
	"os"

	"../sdk"
)

type PhaseTask struct {
	*Task
	TopOfRealAdvance          float64
	MinimunValueOfRealAdvance int
	ValueOfRealAdvance        int
}

func (a PhaseTask) TaskWeight() int {
	return a.TopOfRealAdvanceInt() / a.Weight()
}

func (a PhaseTask) PercOfRealAdvance() int {
	return a.IntDivisionInPerc(a.ValueOfRealAdvance, a.ValueOfTheWholeTask())
}

func (a PhaseTask) TopOfRealAdvanceInt() int {
	return sdk.FloatTimesInt(a.TopOfRealAdvance, a.Scale())
}

func (a PhaseTask) TopValueOfRealAdvance() int {
	return a.IntDivisionInPerc(a.ValueOfRealAdvance, a.ValueOfTheWholeTask())
}

func (e PhaseTask) ValueOfTheWholeTaskScaled() float64 {
	return e.IntDivisionByScale(e.ValueOfTheWholeTask())
}

func (e PhaseTask) PercOfTheTaskInProjectIntScaled() float64 {
	return e.IntDivisionByScale(e.PercOfTheTaskInProjectInt())
}

func (e PhaseTask) PercOfRealAdvanceScaled() float64 {
	return e.IntDivisionByScale(e.PercOfRealAdvance())
}

func (e PhaseTask) ValueOfRealAdvanceScaled() float64 {
	return e.IntDivisionByScale(e.ValueOfRealAdvance)
}

func (e PhaseTask) MinorValueScaled() float64 {
	return e.IntDivisionByScale(e.MinorValue())
}

func (a PhaseTask) GetFirstSwapCombination(b PhaseTask, goal int) (n, m int) {
	if a.MinorValue() > b.MinorValue() {
		return b.GetFirstSwapCombination(a, goal)
	}

	res2 := goal % b.MinorValue()
	for i := 1; i < b.MinorValue(); i++ {
		var newM = ((i*a.MinorValue()-res2)/b.MinorValue() - goal/b.MinorValue())
		if i*a.MinorValue()%b.MinorValue() == res2 {
			if (a.ValueOfRealAdvance-i*a.MinorValue()) < a.MinimunValueOfRealAdvance || (newM*b.MinorValue()+b.ValueOfRealAdvance) > b.TopValueOfRealAdvance() {
				continue
			}
			n = i
			m = newM
		}
	}
	return
}

func (e PhaseTask) Print(file *os.File) {
	fmt.Fprintf(file, "Name: %s - TotalPrjt: %8.2f - TopPerc: %8.2f - PercInPrj: %8.2f - PercAdv: %8.2f - ValAdv: %8.2f - MinorStep: %8.2f - Weight: %d\n",
		e.Name,
		e.ValueOfTheWholeTaskScaled(),
		e.TopOfRealAdvance,
		e.PercOfTheTaskInProjectIntScaled(),
		e.PercOfRealAdvanceScaled(),
		e.ValueOfRealAdvanceScaled(),
		e.MinorValueScaled(),
		e.TaskWeight())
}

type PhaseTasks []PhaseTask

func (a PhaseTasks) Len() int           { return len(a) }
func (a PhaseTasks) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a PhaseTasks) Less(i, j int) bool { return a[i].MinorValue() < a[j].MinorValue() }
func (a PhaseTasks) SortByMinorValue() {
	for index := 0; index < len(a); index++ {
		for index2 := 0; index2 < len(a); index2++ {
			if a[index].MinorValue() < a[index2].MinorValue() {
				a[index], a[index2] = a[index2], a[index]
			}
		}
	}
}

func (a PhaseTasks) SortByName() {
	for index := 0; index < len(a); index++ {
		for index2 := 0; index2 < len(a); index2++ {
			if a[index].Name < a[index2].Name {
				a[index], a[index2] = a[index2], a[index]
			}
		}
	}
}

func (a PhaseTasks) Reverse() {
	var size = len(a) / 2
	for index := 0; index < size; index++ {
		a[index], a[len(a)-index-1] = a[len(a)-index-1], a[index]
	}
}

func (a PhaseTasks) OrderByMinorValueDesc() {
	a.SortByMinorValue()
	a.Reverse()
}

func (a PhaseTasks) SumValueOfAdvance() int {
	var total = 0
	for index := 0; index < len(a); index++ {
		total += a[index].ValueOfRealAdvance
	}
	return total
}

func (a PhaseTasks) Print(file *os.File) {
	for index := 0; index < len(a); index++ {
		a[index].Print(file)
	}
}
