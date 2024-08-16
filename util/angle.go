package util

import "math"

func DegreesToRadians(angle float64) float64 {
	return angle * (math.Pi / 180)
}

func RadiansToDegress(angle float64) float64 {
	return angle * (180 / math.Pi)
}
