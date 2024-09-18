package noise

import (
	f "fmt"
	"os"
)

func Interface() {
	Arr(Array)
	Shuffle(Array)
	var choise = -1
	var subchoise = -1
	var p Param
	p = *p.NewParam(4.0, 0.01, 4, 0.5, 2)
	p.Print()

	for choise != 0 {
		p.Print()
		f.Printf("1. Изменить параметры.\n2. Генерировать\n0.  Выход\n\n")
		f.Fscan(os.Stdin, &choise)
		f.Println()

		switch choise {
		case 1:
			{
				p.Print()
				f.Println("1. size")
				f.Println("2. step")
				f.Println("3. octaves")
				f.Println("4. persistance")
				f.Println("5. frequency")
				f.Println("0. Отмена")
				f.Fscan(os.Stdin, &subchoise)
				if subchoise != 0 {
					p.Change(subchoise)
				} else {
					subchoise = -1
				}
			}

		case 2:
			{
				f.Printf("2. Генерация значений шума:")
				p.Print()
				MatrixNoiseWithOctaves(p.Lenght, p.Step, p.Octaves, p.Persistence, p.Frequency)
			}
		}
	}
}
