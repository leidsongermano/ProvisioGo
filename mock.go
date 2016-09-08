package main

import "./model"

func MockProject() (prj model.Project) {

	var taskValueList = []float64{15000, 5000, 6000, 6500, 10500, 1500, 3000, 3000, 1500, 1500, 1500, 3500, 4000, 1500, 1500, 1500, 1706}
	var goalPhases = []float64{38000, 10000, 12000}
	var topLists = [][]float64{
		[]float64{100, 100, 100, 90, 40, 40, 0, 0, 40, 40, 40, 0, 0, 40, 40, 40, 0},
		[]float64{100, 100, 100, 100, 85, 80, 0, 0, 80, 80, 80, 0, 0, 80, 80, 80, 0},
		[]float64{100, 100, 100, 100, 95, 95, 80, 80, 90, 90, 90, 0, 0, 90, 90, 90, 0}}

	prj = model.Project{
		Name:        "CSA",
		TotalAmount: 68706,
		BaseModel: &model.BaseModel{
			Perc:      100,
			Precision: 4,
			Base:      10}}

	prj.FillTasks(taskValueList)
	prj.FillPhases(goalPhases, topLists)
	return
}
