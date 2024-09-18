package noise

import (
	"math/rand"
	"os"
	"os/exec"
)

func Arr(array []int) {
	for i := 0; i < len(array); i++ {
		array[i] = i
	}
}

func Shuffle(array []int) {
	rand.Seed(17)
	//rand.Seed(time.Now().UnixNano())
	for i := len(array) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		array[i], array[j] = array[j], array[i]
	}
}

func Interpolation(a, b, t float32) float32 {
	res := a + (b-a)*Curve(t)
	return res
}

func Curve(t float32) float32 {
	return t * t * t * (t*(t*6-15) + 10)
}
func GetGradient(x int, y int) (v [2]float32) {

	//var buf int
	iX := x % 16
	iY := y % 16

	switch (Array[(iX*16+iY)%256]) % 4 {
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

func Noise(x, y float32) (result float32) {
	//целое значение точки
	x00 := (int)(x / 1)
	y00 := (int)(y / 1)

	//дробное значение точки
	Xdot := (x - (float32)(x00))
	Ydot := (y - (float32)(y00))

	//координаты вершин квадрата
	d00 := [2]int{x00, y00}
	d01 := [2]int{x00, y00 + 1}
	d10 := [2]int{x00 + 1, y00}
	d11 := [2]int{x00 + 1, y00 + 1}

	//задаем градиенты вершинам квадрата
	g00 := GetGradient(d00[0], d00[1])
	g01 := GetGradient(d01[0], d01[1])
	g10 := GetGradient(d10[0], d10[1])
	g11 := GetGradient(d11[0], d11[1])
	/*
		g00 := GetGradientHash(d00[0], d00[1])
		g01 := GetGradientHash(d01[0], d01[1])
		g10 := GetGradientHash(d10[0], d10[1])
		g11 := GetGradientHash(d11[0], d11[1])
	*/
	//векторы до точки в квадрате
	p00 := [2]float32{Xdot, Ydot}
	p01 := [2]float32{Xdot, Ydot - 1}
	p10 := [2]float32{Xdot - 1, Ydot}
	p11 := [2]float32{Xdot - 1, Ydot - 1}

	//скалярное произведение двух векторов
	s00 := g00[0]*p00[0] + g00[1]*p00[1]
	s01 := g01[0]*p01[0] + g01[1]*p01[1]
	s10 := g10[0]*p10[0] + g10[1]*p10[1]
	s11 := g11[0]*p11[0] + g11[1]*p11[1]

	//интерполяция
	I0 := Interpolation(s00, s10, Xdot)
	I1 := Interpolation(s01, s11, Xdot)
	result = Interpolation(I0, I1, Ydot)

	result = (result + 0.54) / 1.08
	return result
}

func Maxmin(r float32) int {
	if r > max {
		max = r
		return 1
	}
	if r < min {
		min = r
		return 2
	}
	return 3
}

func VisualizeData(filename string) error {
	cmd := exec.Command("python", "E:/учеба/Noise/py/visualaze.py", filename)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
