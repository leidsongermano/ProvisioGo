package main

import (
	"fmt"
	"log"
	"math"
	"os"

	"./model"
)

func main() {
	f, err := os.Create("test.log")
	if err != nil {
		log.Fatal("Cannot create file", err)
	}
	defer f.Close()
	//var testPhase = 3
	var proj = model.MockProject()
	ProcessProject(proj, f, false)
	proj.Print(nil, f)
}

func ProcessProject(prj model.Project, file *os.File, debug bool) {
	for phaseIndex := 0; phaseIndex < len(prj.Phases); phaseIndex++ {
		switch phaseIndex {
		case len(prj.Phases) - 1:
			FillPhaseWithLastPhase(prj.Phases[phaseIndex-1], &prj.Phases[phaseIndex], file)
			CompleteLastPhase(&prj.Phases[phaseIndex], file)
			break
		case 0:
			ProcessPhase(&prj.Phases[phaseIndex], file, debug)
			break
		default:
			FillPhaseWithLastPhase(prj.Phases[phaseIndex-1], &prj.Phases[phaseIndex], file)
			ProcessPhase(&prj.Phases[phaseIndex], file, debug)
		}
	}
}

func FillPhaseWithLastPhase(lastPhase model.Phase, currentPhase *model.Phase, file *os.File) {
	for taskIndex := 0; taskIndex < len(lastPhase.Tasks); taskIndex++ {
		currentPhase.Tasks[taskIndex].ValueOfRealAdvance = lastPhase.Tasks[taskIndex].ValueOfRealAdvance
		currentPhase.Tasks[taskIndex].MinimunValueOfRealAdvance = lastPhase.Tasks[taskIndex].ValueOfRealAdvance
	}
}

func CompleteLastPhase(phase *model.Phase, file *os.File) {
	for taskIndex := 0; taskIndex < len(phase.Tasks); taskIndex++ {
		phase.Tasks[taskIndex].ValueOfRealAdvance = phase.Tasks[taskIndex].ValueOfTheWholeTask()
	}
}

func ProcessPhase(phase *model.Phase, file *os.File, debug bool) {
	fmt.Fprintln(file, "Start process...")
	phase.Tasks.OrderByMinorValueDesc()
	addValuesBasedOnWeights(phase, file, debug)
	fmt.Fprintf(file, "Phase %02d/%d => Best result with weights: %d\n", phase.Month, phase.Year, phase.Tasks.SumValueOfAdvance())
	fmt.Fprintln(file, "Phase => Adjust in precision...")
	addValuesBasedOnMinorValues(phase, file, debug)
	fmt.Fprintf(file, "Phase %02d/%d => Best result with minor value: %d\n", phase.Month, phase.Year, phase.Tasks.SumValueOfAdvance())
	fmt.Fprintln(file, "Phase => Adjust subtract and in precision...")
	tryToFit(phase, file, true)
	fmt.Fprintf(file, "Phase %02d/%d => Best result with fit: %d\n\n", phase.Month, phase.Year, phase.Tasks.SumValueOfAdvance())
	phase.Tasks.SortByName()
}

func addValuesBasedOnWeights(phase *model.Phase, file *os.File, debug bool) {
	var control = 0
	for phase.Tasks.SumValueOfAdvance() < phase.AccGoal && control < (len(phase.Tasks)-1) {
		for index := 0; index < len(phase.Tasks); index++ {
			var e = &phase.Tasks[index]
			var add = e.TaskWeight() * e.MinorValue()
			var newPerc = e.IntDivisionInPerc((e.ValueOfRealAdvance + add), e.ValueOfTheWholeTask())
			if phase.Tasks.SumValueOfAdvance()+add <= phase.AccGoal && e.TaskWeight() > 0 && newPerc <= e.TopOfRealAdvanceInt() {
				e.ValueOfRealAdvance += add
				if debug {
					fmt.Fprintf(file, "Task: %s Total added: %12d Total: %12d NewPerc: %12d Top: %12d\n", e.Name, add, phase.Tasks.SumValueOfAdvance(), newPerc, e.TopOfRealAdvanceInt())
				}
				control = 0
			} else {
				control++
				if debug {
					fmt.Fprintf(file, "Control: %d\n", control)
				}
			}
		}
	}
	if debug {
		fmt.Fprintf(file, "Total obtained in addValuesBasedOnWeights: %12.2f\n", phase.IntDivisionByScale(phase.Tasks.SumValueOfAdvance()))
	}
}

func addValuesBasedOnMinorValues(phase *model.Phase, file *os.File, debug bool) {
	var control = 0
	for phase.Tasks.SumValueOfAdvance() < phase.AccGoal && control < (len(phase.Tasks)-1) {
		for index := 0; index < len(phase.Tasks); index++ {
			var e = &phase.Tasks[index]
			var newPerc = e.IntDivisionInPerc((e.ValueOfRealAdvance + e.MinorValue()), e.ValueOfTheWholeTask())
			if phase.Tasks.SumValueOfAdvance()+e.MinorValue() <= phase.AccGoal && newPerc <= e.TopOfRealAdvanceInt() {
				for weightIndex := 0; weightIndex < e.TaskWeight(); weightIndex++ {
					newPerc = e.IntDivisionInPerc((e.ValueOfRealAdvance + e.MinorValue()), e.ValueOfTheWholeTask())
					if phase.Tasks.SumValueOfAdvance()+e.MinorValue() <= phase.AccGoal && newPerc <= e.TopOfRealAdvanceInt() {
						e.ValueOfRealAdvance += e.MinorValue()
						if debug {
							fmt.Fprintf(file, "Task: %s Total added: %12d Total: %12d NewPerc: %12d Top: %12d\n", e.Name, e.MinorValue(), phase.Tasks.SumValueOfAdvance(), newPerc, e.TopOfRealAdvanceInt())
						}
						control = 0
					}
				}
			} else {
				control++
				if debug {
					fmt.Fprintf(file, "Control: %d\n", control)
				}
			}
		}
	}
	if debug {
		fmt.Fprintf(file, "Total obtained in addValuesBasedOnMinorValues: %.4f\n", phase.Tasks.SumValueOfAdvance())
	}
}

func tryToFit(phase *model.Phase, file *os.File, debug bool) {
	var newGoal = phase.AccGoal - phase.Tasks.SumValueOfAdvance()
	for index := 0; index < len(phase.Tasks); index++ {
		var e = &phase.Tasks[index]
		for index2 := index + 1; index2 < len(phase.Tasks); index2++ {
			var e2 = &phase.Tasks[index2]

			n, m := e.GetFirstSwapCombination(*e2, newGoal)
			if debug {
				fmt.Fprintf(file, "%2d * %2d - %2d * %2d = %2d\n", e.MinorValue(), n, e2.MinorValue(), m, e.MinorValue()*n-e2.MinorValue()*m)
			}

			if int(math.Abs(float64(e.MinorValue()*n-e2.MinorValue()*m))) == newGoal {
				e.ValueOfRealAdvance -= n * e.MinorValue()
				e2.ValueOfRealAdvance += m * e2.MinorValue()
				return
			}
		}
	}
}
