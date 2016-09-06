package main

import "math"

type myTask struct {
	name                   string
	valueOfTheWholeTask    int
	percOfTheTaskInProject int
	topOfRealAdvance       int
	topValueOfRealAdvance  int
	valueOfRealAdvance     int
	minorValue             int
	weight                 int
}

func (a myTask) percOfRealAdvance() int {
	return IntToFloatDivisionInPerc(a.valueOfRealAdvance, a.valueOfTheWholeTask)
}

type myTasks []myTask

func (a myTasks) Len() int           { return len(a) }
func (a myTasks) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a myTasks) Less(i, j int) bool { return a[i].minorValue < a[j].minorValue }
func (a myTasks) SortByMinorValue() {
	for index := 0; index < len(a); index++ {
		for index2 := 0; index2 < len(a); index2++ {
			if a[index].minorValue < a[index2].minorValue {
				a[index], a[index2] = a[index2], a[index]
			}
		}
	}
}

func (a myTasks) SortByName() {
	for index := 0; index < len(a); index++ {
		for index2 := 0; index2 < len(a); index2++ {
			if a[index].name < a[index2].name {
				a[index], a[index2] = a[index2], a[index]
			}
		}
	}
}

func (a myTasks) Reverse() {
	var size int = len(a) / 2
	for index := 0; index < size; index++ {
		a[index], a[len(a)-index-1] = a[len(a)-index-1], a[index]
	}
}

func (a myTasks) SumValuOfAdvance() int {
	var total int = 0
	for index := 0; index < len(a); index++ {
		total += a[index].valueOfRealAdvance
	}
	return total
}

func Round(num float64) int {
	return int(math.Floor(num + 0.5))
}

func RoundPlus(val float64, places int) (newVal float64) {
	var round float64
	pow := math.Pow(10, float64(places))
	digit := pow * val
	_, div := math.Modf(digit)
	if div >= 0.5 {
		round = math.Ceil(digit)
	} else {
		round = math.Floor(digit)
	}
	newVal = round / pow
	return
}

func FloatTimesInt(fl float64, i int) int {
	return int(fl * float64(i))
}

func IntToFloatDivisionByScale(fl1 int) float64 {
	return float64(fl1) / float64(scaleWeight)
}

func IntToFloatDivisionByScaleInt(fl1 int) int {
	return int(float64(fl1) / float64(scaleWeight))
}

func IntToFloatDivisionInScale(fl1, fl2 int) int {
	return FloatTimesInt((float64(fl1) / float64(fl2)), scaleWeight)
}

func IntToFloatDivisionInPerc(fl1, fl2 int) int {
	return FloatTimesInt((float64(fl1) / float64(fl2)), scaleWeight*perc)
}

func divisor(a, b myTask, res int) (n int, m int) {
	if a.minorValue > b.minorValue {
		return swap(divisor(b, a, res))
	}
	res2 := res % b.minorValue
	for i := 1; i < b.minorValue; i++ {
		var newM = ((i*a.minorValue-res2)/b.minorValue - res/b.minorValue)
		if i*a.minorValue%b.minorValue == res2 && (i*a.minorValue <= a.topValueOfRealAdvance) && (newM*b.minorValue <= b.topValueOfRealAdvance) {
			n = i
			m = newM
			return
		}
	}
	return
}
func swap(a int, b int) (c int, d int) { return b, a }
