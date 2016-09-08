package model

import (
	"../sdk"
)

type Task struct {
	*BaseModel
	Name                   string
	Value                  float64
	PercOfTheTaskInProject float64
}

func (a Task) ValueOfTheWholeTask() int {
	return sdk.FloatTimesInt(a.Value, a.Scale())
}

func (a Task) MinorValue() int {
	return a.IntDivisionByScaleInt(a.ValueOfTheWholeTask())
}

func (a Task) PercOfTheTaskInProjectInt() int {
	return sdk.FloatTimesInt(a.PercOfTheTaskInProject, a.Scale())
}

type Tasks []Task

func (a Tasks) Len() int           { return len(a) }
func (a Tasks) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a Tasks) Less(i, j int) bool { return a[i].MinorValue() < a[j].MinorValue() }
func (a Tasks) SortByMinorValue() {
	for index := 0; index < len(a); index++ {
		for index2 := 0; index2 < len(a); index2++ {
			if a[index].MinorValue() < a[index2].MinorValue() {
				a[index], a[index2] = a[index2], a[index]
			}
		}
	}
}

func (a Tasks) SortByName() {
	for index := 0; index < len(a); index++ {
		for index2 := 0; index2 < len(a); index2++ {
			if a[index].Name < a[index2].Name {
				a[index], a[index2] = a[index2], a[index]
			}
		}
	}
}

func (a Tasks) Reverse() {
	var size = len(a) / 2
	for index := 0; index < size; index++ {
		a[index], a[len(a)-index-1] = a[len(a)-index-1], a[index]
	}
}

func (a Tasks) OrderByMinorValueDesc() {
	a.SortByMinorValue()
	a.Reverse()
}
