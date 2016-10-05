package main

import "./model"

// func MockProject() (prj model.Project) {

// 	var taskValueList = []float64{8000, 2500, 5000, 2300, 3000, 6000, 5000, 6500, 5500, 2500, 2000, 9700, 8700, 3000, 1500, 2300, 4000, 2800, 4000, 2300, 2700, 4000, 2300, 2500, 3500, 4500, 2300, 2300, 2300, 2000}
// 	var goalPhases = []float64{23000, 23000, 23000, 23000}
// 	var topLists = [][]float64{
// 		[]float64{90, 90, 87.60, 90, 90, 40, 40, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
// 		[]float64{95, 95, 90, 95, 95, 60, 60, 50, 49.99, 69.99, 69.99, 55.01, 55.01, 0, 0, 26.92, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
// 		[]float64{100, 100, 100, 100, 100, 100, 100, 95, 95, 95, 95, 75, 75, 50, 10, 75, 0, 40, 0, 45, 45, 0, 0, 0, 0, 0, 30, 30, 30, 0},
// 		[]float64{100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 95, 95, 95, 90, 95, 60, 95, 0, 95, 95, 0, 95, 95, 0, 0, 95, 95, 95, 0}}

// 	prj = model.Project{
// 		Name:        "CSA",
// 		TotalAmount: 115000,
// 		BaseModel: &model.BaseModel{
// 			Perc:      100,
// 			Precision: 4,
// 			Base:      10}}

// 	prj.FillTasks(taskValueList)
// 	prj.FillPhases(goalPhases, topLists)
// 	return
// }

func MockProject() (prj model.Project) {

	var taskValueList = []float64{5650, 750, 2300, 2250, 3778, 1800.86, 1450, 2400, 2150}
	var goalPhases = []float64{12000}
	var topLists = [][]float64{
		[]float64{100, 95, 95, 95, 60, 0, 0, 0, 0}}

	prj = model.Project{
		Name:        "CSA",
		TotalAmount: 22528.86,
		BaseModel: &model.BaseModel{
			Perc:      100,
			Precision: 4,
			Base:      10}}

	prj.FillTasks(taskValueList)
	prj.FillPhases(goalPhases, topLists)
	return
}
