package noise

func GetGradient1(x int) (v float32) {
	iX := x % 16

	switch array1[iX] % 2 {
	case 0:
		return float32(-1.0)
	case 1:
		return float32(1.0)
	}
	return
}
func GetGradientHash(x int, y int) (v [2]float32) {
	value := mixedHash(x, y) % 4
	switch value {
	case 0:
		v[0] = 0.0
		v[1] = 1.0
	case 1:
		v[0] = 0.0
		v[1] = -1.0
	case 2:
		v[0] = 1.0
		v[1] = 0.0
	case 3:
		v[0] = -1.0
		v[1] = 0.0
	}
	return v
}
func mixedHash(x, y int) int {
	var h = x*374761393 + y*668265263
	h = (h ^ (h >> 13)) * 1274126177
	return h ^ (h >> 16)
}
