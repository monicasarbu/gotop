package common

import "math"

func Round(f float64) float64 {
	return math.Floor(f + .5)
}

func Percentage(t1 float64, t2 float64) float64 {
	if t2 > 0 {
		perc := (t1 / t2) * 100
		return Round(perc)
	}
	return 0.0
}
