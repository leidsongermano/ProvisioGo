package main

import (
	"fmt"
	"math"
	"strconv"
)

/************DEFINE****************/
//var scale int = 4
//var precision float64 = 0.01
var minorStep int = 1
var perc int = 100
var scale int = 4
var scaleWeight int = int(math.Pow(10, float64(scale)))
var weight int = 10 * scaleWeight

/*********************************/

var totalOfprojectOriginal float64 = 68706
var goalOriginal float64 = 38000
var taskValueListOriginal = []float64{15000, 5000, 6000, 6500, 10500, 1500, 3000, 3000, 1500, 1500, 1500, 3500, 4000, 1500, 1500, 1500, 1706}
var topListOriginal = []float64{100, 100, 100, 90, 40, 40, 0, 0, 40, 40, 40, 0, 0, 40, 40, 40, 0}

var totalOfproject int = Round(totalOfprojectOriginal * float64(scaleWeight))
var goal int = Round(goalOriginal * float64(scaleWeight))
var taskValueList = make([]int, len(taskValueListOriginal))
var topList = make([]int, len(taskValueListOriginal))
var taskList myTasks = make(myTasks, len(taskValueListOriginal))

func main() {
	normalizeLists()
	fillTaskList(false)
	calculateTheValues(false)
	taskList.SortByName()
	formatTaskList()
	printTaskList()
	fmt.Printf("Total: %12.2f\n", IntToFloatDivisionByScale(taskList.SumValuOfAdvance()))
}

func normalizeLists() {
	for index := 0; index < len(taskValueListOriginal); index++ {
		taskValueList[index] = FloatTimesInt(taskValueListOriginal[index], scaleWeight)
		topList[index] = FloatTimesInt(topListOriginal[index], scaleWeight)
	}
}

func formatTaskList() {
	for index := 0; index < len(taskValueListOriginal); index++ {
		taskValueList[index] = FloatTimesInt(taskValueListOriginal[index], scaleWeight)
		topList[index] = FloatTimesInt(topListOriginal[index], scaleWeight)
	}
}

func printTaskList() {
	for index := 0; index < len(taskValueList); index++ {
		var e = taskList[index]
		fmt.Printf("name: %s - totalPrjt: %8.2f - topPerc: %8.2f - percInPrj: %8.2f - percAdv: %8.2f - valAdv: %8.2f - minorStep: %8.2f - weight: %d\n",
			e.name, IntToFloatDivisionByScale(e.valueOfTheWholeTask), IntToFloatDivisionByScale(e.topOfRealAdvance), IntToFloatDivisionByScale(e.percOfTheTaskInProject),
			IntToFloatDivisionByScale(e.percOfRealAdvance()), IntToFloatDivisionByScale(e.valueOfRealAdvance), IntToFloatDivisionByScale(e.minorValue), e.weight)

	}
}

func fillTaskList(debug bool) {
	var total int
	for index := 0; index < len(taskValueList); index++ {
		var e myTask = myTask{
			name:                   "B" + strconv.Itoa(index+13),
			valueOfTheWholeTask:    taskValueList[index],
			percOfTheTaskInProject: IntToFloatDivisionInPerc(taskValueList[index], totalOfproject),
			minorValue:             IntToFloatDivisionByScaleInt(taskValueList[index]),
			weight:                 topList[index] / weight,
			topOfRealAdvance:       topList[index],
			topValueOfRealAdvance:  IntToFloatDivisionByScaleInt(taskValueList[index]) * topList[index]}
		total += e.percOfTheTaskInProject
		taskList[index] = e
	}
	//sort.Sort(sort.Reverse(taskList))
	taskList.SortByMinorValue()
	taskList.Reverse()
	if debug {
		printTaskList()
		fmt.Printf("Total of perc: %.2f\n", float64(total)/float64(scaleWeight))
	}
}

func addValuesBasedOnWeights(debug bool) {
	var control int = 0
	for taskList.SumValuOfAdvance() < goal && control < (len(taskList)-1) {
		for index := 0; index < len(taskList); index++ {
			var e myTask = taskList[index]
			var add int = e.weight * e.minorValue
			var newPerc int = IntToFloatDivisionInPerc((e.valueOfRealAdvance + add), e.valueOfTheWholeTask)
			if taskList.SumValuOfAdvance() > goal {
				fmt.Printf("Total added: %12d Total: %12d NewPerc: %12d Top: %12d\n", add, taskList.SumValuOfAdvance(), newPerc, e.topOfRealAdvance)
			}
			if taskList.SumValuOfAdvance()+add < goal && e.weight > 0 && newPerc < e.topOfRealAdvance {
				e.valueOfRealAdvance += add
				taskList[index] = e
				if debug {
					fmt.Printf("Total added: %d Total now: %d\n", add, taskList.SumValuOfAdvance())
				}
				control = 0
			} else {
				control++
				if debug {
					fmt.Printf("Control: %d\n", control)
				}
			}
		}
	}
	if debug {
		fmt.Printf("Total obtained in addValuesBasedOnWeights: %12.2f\n", IntToFloatDivisionByScale(taskList.SumValuOfAdvance()))
	}
}

func addValuesBasedOnMinorValues(debug bool) {
	var control int = 0
	for taskList.SumValuOfAdvance() < goal && control < (len(taskList)-1) {
		for index := 0; index < len(taskList); index++ {
			var e myTask = taskList[index]
			if taskList.SumValuOfAdvance()+e.minorValue < goal {
				for weightIndex := 0; weightIndex < e.weight; weightIndex++ {
					var newPerc int = IntToFloatDivisionInPerc((e.valueOfRealAdvance + e.minorValue), e.valueOfTheWholeTask)
					if taskList.SumValuOfAdvance()+e.minorValue < goal && newPerc < e.topOfRealAdvance {
						e.valueOfRealAdvance += e.minorValue
						taskList[index] = e
						if debug {
							fmt.Printf("Total added: %f Total now: %f\n", e.minorValue, taskList.SumValuOfAdvance())
						}
						control = 0
					}
				}
			} else {
				control++
				if debug {
					fmt.Printf("Control: %d\n", control)
				}
			}
		}
	}
	if debug {
		fmt.Printf("Total obtained in addValuesBasedOnMinorValues: %.4f\n", taskList.SumValuOfAdvance())
	}
}

func tryToFit(debug bool) {
	var newGoal = goal - taskList.SumValuOfAdvance()
	for index := 0; index < len(taskList); index++ {
		var e myTask = taskList[index]
		for index2 := index + 1; index2 < len(taskList); index2++ {
			var e2 myTask = taskList[index2]

			n, m := divisor(e, e2, newGoal)
			if debug {
				fmt.Printf("%2d * %2d - %2d * %2d = %2d\n", e.minorValue, n, e2.minorValue, m, e.minorValue*n-e2.minorValue*m)
			}

			if int(math.Abs(float64(e.minorValue*n-e2.minorValue*m))) == newGoal {
				e.valueOfRealAdvance -= n * e.minorValue
				e2.valueOfRealAdvance += m * e2.minorValue
				taskList[index] = e
				taskList[index2] = e2
				return
			}
		}
	}
}

func calculateTheValues(debug bool) {
	//var loop int = 0
	addValuesBasedOnWeights(debug)
	fmt.Printf("Best result with weights: %d\n", taskList.SumValuOfAdvance())
	fmt.Println("Adjust in precision...")
	addValuesBasedOnMinorValues(debug)
	fmt.Printf("Best result with minor value: %d\n", taskList.SumValuOfAdvance())
	fmt.Println("Adjust subtract and in precision...")
	tryToFit(debug)
}
