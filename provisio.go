package main

import (
	"fmt"
	"math"

	"./model"
)

func main() {

	//var testPhase = 1
	var proj = MockProject()
	ProcessProject(proj, false)
	proj.Print(nil)
}

func ProcessProject(prj model.Project, debug bool) {
	for phaseIndex := 0; phaseIndex < len(prj.Phases); phaseIndex++ {
		switch phaseIndex {
		case len(prj.Phases) - 1:
			FillPhaseWithLastPhase(prj.Phases[phaseIndex-1], &prj.Phases[phaseIndex], debug)
			CompleteLastPhase(&prj.Phases[phaseIndex], debug)
			break
		case 0:
			ProcessPhase(&prj.Phases[phaseIndex], debug)
			break
		default:
			FillPhaseWithLastPhase(prj.Phases[phaseIndex-1], &prj.Phases[phaseIndex], debug)
			ProcessPhase(&prj.Phases[phaseIndex], debug)
		}
	}
}

func FillPhaseWithLastPhase(lastPhase model.Phase, currentPhase *model.Phase, debug bool) {
	for taskIndex := 0; taskIndex < len(lastPhase.Tasks); taskIndex++ {
		currentPhase.Tasks[taskIndex].ValueOfRealAdvance = lastPhase.Tasks[taskIndex].ValueOfRealAdvance
	}
}

func CompleteLastPhase(phase *model.Phase, debug bool) {
	for taskIndex := 0; taskIndex < len(phase.Tasks); taskIndex++ {
		phase.Tasks[taskIndex].ValueOfRealAdvance = phase.Tasks[taskIndex].ValueOfTheWholeTask()
	}
}

func ProcessPhase(phase *model.Phase, debug bool) {
	phase.Tasks.OrderByMinorValueDesc()
	addValuesBasedOnWeights(phase, debug)
	fmt.Printf("Phase %02d/%d => Best result with weights: %d\n", phase.Month, phase.Year, phase.Tasks.SumValuOfAdvance())
	fmt.Println("Phase => Adjust in precision...")
	addValuesBasedOnMinorValues(phase, debug)
	fmt.Printf("Phase %02d/%d => Best result with minor value: %d\n", phase.Month, phase.Year, phase.Tasks.SumValuOfAdvance())
	fmt.Println("Phase => Adjust subtract and in precision...")
	tryToFit(phase, debug)
	fmt.Printf("Phase %02d/%d => Best result with fit: %d\n\n", phase.Month, phase.Year, phase.Tasks.SumValuOfAdvance())
}

func addValuesBasedOnWeights(phase *model.Phase, debug bool) {
	var control = 0
	for phase.Tasks.SumValuOfAdvance() < phase.AccGoal && control < (len(phase.Tasks)-1) {
		for index := 0; index < len(phase.Tasks); index++ {
			var e = &phase.Tasks[index]
			var add = e.TaskWeight() * e.MinorValue()
			var newPerc = e.IntDivisionInPerc((e.ValueOfRealAdvance + add), e.ValueOfTheWholeTask())
			if phase.Tasks.SumValuOfAdvance()+add < phase.AccGoal && e.TaskWeight() > 0 && newPerc < e.TopOfRealAdvanceInt() {
				e.ValueOfRealAdvance += add
				if debug {
					fmt.Printf("Task: %s Total added: %12d Total: %12d NewPerc: %12d Top: %12d\n", e.Name, add, phase.Tasks.SumValuOfAdvance(), newPerc, e.TopOfRealAdvanceInt())
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
		fmt.Printf("Total obtained in addValuesBasedOnWeights: %12.2f\n", phase.IntDivisionByScale(phase.Tasks.SumValuOfAdvance()))
	}
}

func addValuesBasedOnMinorValues(phase *model.Phase, debug bool) {
	var control = 0
	for phase.Tasks.SumValuOfAdvance() < phase.AccGoal && control < (len(phase.Tasks)-1) {
		for index := 0; index < len(phase.Tasks); index++ {
			var e = &phase.Tasks[index]
			if phase.Tasks.SumValuOfAdvance()+e.MinorValue() < phase.AccGoal {
				for weightIndex := 0; weightIndex < e.TaskWeight(); weightIndex++ {
					var newPerc = e.IntDivisionInPerc((e.ValueOfRealAdvance + e.MinorValue()), e.ValueOfTheWholeTask())
					if phase.Tasks.SumValuOfAdvance()+e.MinorValue() < phase.AccGoal && newPerc < e.TopOfRealAdvanceInt() {
						e.ValueOfRealAdvance += e.MinorValue()
						if debug {
							fmt.Printf("Task: %s Total added: %12d Total: %12d NewPerc: %12d Top: %12d\n", e.Name, e.MinorValue(), phase.Tasks.SumValuOfAdvance(), newPerc, e.TopOfRealAdvanceInt())
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
		fmt.Printf("Total obtained in addValuesBasedOnMinorValues: %.4f\n", phase.Tasks.SumValuOfAdvance())
	}
}

func tryToFit(phase *model.Phase, debug bool) {
	var newGoal = phase.AccGoal - phase.Tasks.SumValuOfAdvance()
	for index := 0; index < len(phase.Tasks); index++ {
		var e = &phase.Tasks[index]
		for index2 := index + 1; index2 < len(phase.Tasks); index2++ {
			var e2 = &phase.Tasks[index2]

			n, m := e.GetFirstSwapCombination(*e2, newGoal)
			if debug {
				fmt.Printf("%2d * %2d - %2d * %2d = %2d\n", e.MinorValue(), n, e2.MinorValue(), m, e.MinorValue()*n-e2.MinorValue()*m)
			}

			if int(math.Abs(float64(e.MinorValue()*n-e2.MinorValue()*m))) == newGoal {
				e.ValueOfRealAdvance -= n * e.MinorValue()
				e2.ValueOfRealAdvance += m * e2.MinorValue()
				return
			}
		}
	}
}
