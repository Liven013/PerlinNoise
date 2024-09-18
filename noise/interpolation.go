package noise

import "math"

func Lerp(t float32) float32 {
	return t
}
func CosInter(t float32) float32 {
	return float32((1 - math.Cos(float64(t*math.Pi))) / 2)
}
func CubeInter(t float32) float32 {
	return -2*t*t*t + 3*t*t
}
