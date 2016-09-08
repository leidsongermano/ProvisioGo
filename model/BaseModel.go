package model

import (
	"math"

	"../sdk"
)

//BaseModel class to store and organize the configurations
type BaseModel struct {
	Perc      int
	Precision int
	Base      int
}

func (b BaseModel) Scale() int {
	return int(math.Pow(float64(b.Base), float64(b.Precision)))
}

func (b BaseModel) Weight() int {
	return b.Scale() * b.Base
}

func (b BaseModel) IntDivisionByScale(fl1 int) float64 {
	return sdk.IntToFloatDivision(fl1, b.Scale())
}

func (b BaseModel) IntDivisionByScaleInt(fl1 int) int {

	return int(b.IntDivisionByScale(fl1))
}

func (b BaseModel) IntDivisionInScale(fl1, fl2 int) int {
	return sdk.FloatTimesInt(sdk.IntToFloatDivision(fl1, fl2), b.Scale())
}

func (b BaseModel) IntDivisionInPerc(fl1, fl2 int) int {
	return sdk.FloatTimesInt(sdk.IntToFloatDivision(fl1, fl2), b.Scale()*b.Perc)
}

func (b BaseModel) FloatDivisionInPerc(fl1, fl2 float64) float64 {
	return b.IntDivisionByScale(sdk.FloatTimesInt(fl1/fl2, b.Scale()*b.Perc))
}
