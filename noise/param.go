package noise

import (
	"fmt"
	"os"
)

type Param struct {
	Lenght      float32
	Step        float32
	Octaves     int
	Persistence float32
	Frequency   float32
}

func (p Param) NewParam(lenght float32, step float32, octaves int, persistence float32, frequency float32) *Param {
	return &Param{
		Lenght:      lenght,
		Step:        step,
		Octaves:     octaves,
		Persistence: persistence,
		Frequency:   frequency,
	}
}
func (p Param) Print() {
	fmt.Println("\nТекущее значения параметров:")
	fmt.Printf("Size : %.0f\n", p.Lenght)
	fmt.Printf("Step : %.2f\n", p.Step)
	fmt.Printf("Octaves : %d\n", p.Octaves)
	fmt.Printf("Persistence : %.2f\n", (float64)(p.Persistence))
	fmt.Printf("Frequency : %.1f\n\n", (float64)(p.Frequency))
}

func (p *Param) Change(choise int) {

	fmt.Printf("Введите новое значение")
	switch choise {
	case 1:
		{
			fmt.Println(" size")
			fmt.Fscan(os.Stdin, &p.Lenght)
		}
	case 2:
		{
			fmt.Println(" step")
			fmt.Fscan(os.Stdin, &p.Step)
		}
	case 3:
		{
			fmt.Println(" octaves")
			fmt.Fscan(os.Stdin, &p.Octaves)
		}
	case 4:
		{
			fmt.Println(" persistance")
			fmt.Fscan(os.Stdin, &p.Persistence)
		}
	case 5:
		{
			fmt.Println(" frequency")
			fmt.Fscan(os.Stdin, &p.Frequency)
		}
	}
	p.Print()
}
