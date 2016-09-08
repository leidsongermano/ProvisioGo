package sdk

import "math"

//Round function is for round a float64 to next int
func Round(num float64) int {
	return int(math.Floor(num + 0.5))
}

//RoundPlus function is for truncate a float64 to a speficic precision scale
func RoundPlus(num float64, scale int) (newVal float64) {
	var round float64
	pow := math.Pow(10, float64(scale))
	digit := pow * num
	_, div := math.Modf(digit)
	if div >= 0.5 {
		round = math.Ceil(digit)
	} else {
		round = math.Floor(digit)
	}
	newVal = round / pow
	return
}

//FloatTimesInt fuction to multiply a float64 by a int
func FloatTimesInt(fl float64, i int) int {
	return int(fl * float64(i))
}

//IntToFloatDivision divide two ints in float64 format
func IntToFloatDivision(fl1, fl2 int) float64 {
	return float64(fl1) / float64(fl2)
}

//DivisorWithLimit function to found the first combinantion of n,m times a,b to result in -(res)
// a*n - b*m = -(res)
func DivisorWithLimit(a, b, res int, alimit, blimit *int) (n int, m int) {
	if a > b {
		return swap(DivisorWithLimit(b, a, res, blimit, alimit))
	}
	res2 := res % b
	for i := 1; i < b; i++ {
		var newM = ((i*a-res2)/b - res/b)
		if i*a%b == res2 {
			if alimit != nil && blimit != nil && ((i*a > *alimit) || (newM*b > *blimit)) {
				continue
			}
			n = i
			m = newM
			return
		}
	}
	return
}
