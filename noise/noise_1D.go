package noise

import (
	"encoding/binary"
	"fmt"
	"os"
)

func Noise1(x float32) (result float32) {
	x00 := (int)(x / 1)
	Xdot := (x - (float32)(x00))

	d0 := x00
	d1 := x00 + 1

	g0 := GetGradient1(d0)
	g1 := GetGradient1(d1)

	p0 := Xdot
	p1 := Xdot - 1

	s0 := g0 * p0
	s1 := g1 * p1
	result = Interpolation(s0, s1, Xdot)
	return result + 0.5
}

func Noise1D(len float32, step float32) {
	Arr(array1)
	Shuffle(array1)
	name := "data/1D/" + fmt.Sprint(int(len/step)) + "_" + fmt.Sprint(len) + "_" + fmt.Sprint(step) + ".bin"
	f, err := os.Create(name)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer f.Close()
	for j := float32(0.0); j < len; j += step {
		byteVal := make([]byte, 8)
		binary.BigEndian.PutUint64(byteVal, uint64(100*Noise1((float32)(j))))
		_, err = f.Write(byteVal)
		if err != nil {
			fmt.Println("Ошибка при записи в файл:", err)
			return
		}
	}
	fmt.Println("good job")
}
