package model

import (
	"fmt"
	"strconv"

	"../sdk"
)

type Project struct {
	*BaseModel
	Name        string
	TotalAmount float64
	Tasks       Tasks
	Phases      Phases
}

func (p Project) TotalAmountInt() int {
	return sdk.FloatTimesInt(p.TotalAmount, p.Scale())
}

func (prj *Project) FillTasks(taskValueList []float64) {
	prj.Tasks = make(Tasks, len(taskValueList))
	for taskIndex := 0; taskIndex < len(taskValueList); taskIndex++ {
		prj.Tasks[taskIndex] = Task{
			Name:      "B" + strconv.Itoa(taskIndex+13),
			BaseModel: prj.BaseModel,
			Value:     taskValueList[taskIndex],
			PercOfTheTaskInProject: prj.FloatDivisionInPerc(taskValueList[taskIndex], prj.TotalAmount)}
	}
}

func (prj *Project) FillPhases(goalPhases []float64, topLists [][]float64) {
	prj.Phases = make(Phases, len(goalPhases)+1)
	var accGoal int = 0
	for phaseIndex := 0; phaseIndex < len(goalPhases); phaseIndex++ {
		prj.Phases[phaseIndex] = Phase{
			BaseModel: prj.BaseModel,
			Goal:      goalPhases[phaseIndex],
			Month:     phaseIndex + 1,
			Year:      2016}

		accGoal += prj.Phases[phaseIndex].GoalInt()
		prj.Phases[phaseIndex].AccGoal = accGoal
		prj.FillPhaseTasks(topLists[phaseIndex], phaseIndex)
	}

	prj.Phases[len(goalPhases)] = Phase{
		BaseModel: prj.BaseModel,
		Goal:      prj.TotalAmount - prj.Phases.SumFirstNGoals(len(goalPhases)),
		AccGoal:   prj.TotalAmountInt(),
		Month:     prj.Phases[len(goalPhases)-1].Month + 1,
		Year:      2016}

	prj.FillFinalPhaseTasks()
}

func (prj *Project) FillPhaseTasks(topList []float64, phaseIndex int) {
	prj.Phases[phaseIndex].Tasks = make(PhaseTasks, len(prj.Tasks))
	for index := 0; index < len(prj.Tasks); index++ {
		prj.Phases[phaseIndex].Tasks[index].TopOfRealAdvance = topList[index]
		prj.Phases[phaseIndex].Tasks[index].Task = &prj.Tasks[index]
	}
}

func (prj *Project) FillFinalPhaseTasks() {
	prj.Phases[len(prj.Phases)-1].Tasks = make(PhaseTasks, len(prj.Tasks))
	for index := 0; index < len(prj.Tasks); index++ {
		prj.Phases[len(prj.Phases)-1].Tasks[index].TopOfRealAdvance = 100
		prj.Phases[len(prj.Phases)-1].Tasks[index].Task = &prj.Tasks[index]
	}
}

func (prj Project) Print(qtyPhases *int) {
	fmt.Println()
	fmt.Printf("Project.BaseModel => Perc: %d - Precision: %d - Base: %d - Scale: %d\n", prj.Perc, prj.Precision, prj.Base, prj.Scale())
	fmt.Printf("Project => Name: %s - TotalAmount: %8.2f\n", prj.Name, prj.TotalAmount)
	prj.Phases.Print(qtyPhases)
}
