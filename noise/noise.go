package noise

import (
	"fmt"
	"os"
)

var (
	Array  []int = make([]int, 256)
	max          = float32(0.5)
	min          = float32(-0.5)
	array1 []int = make([]int, 16)
)

func MatrixNoiseWithOctaves(lenght float32, step float32, octaves int, persistence float32, frequency float32) {
	var i, j float32
	pix := 0
	name := "data/" + fmt.Sprint(lenght) + "_" + fmt.Sprint(step) + "_" + fmt.Sprint(octaves) + "_" + fmt.Sprint(persistence) + "_" + fmt.Sprint(frequency) + ".txt"
	f, err := os.Create(name)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer f.Close()

	for i = 0.0; i < lenght; i += step {
		for j = 0.0; j < lenght; j += step {
			pix = (int)(255 * NoiseWithOctaves((float32)(i), (float32)(j), octaves, persistence, frequency))
			fmt.Fprintf(f, "%d ", pix)
		}
		fmt.Fprintln(f)
	}

	err = VisualizeData(name)
	if err != nil {
		fmt.Println("Failed to execute Python script")
	}

	fmt.Println("Data generated and visualized successfully")
}
func NoiseWithOctaves(fx float32, fy float32, octaves int, persistence float32, frequency float32) float32 {
	var (
		amplitude float32
		max       float32
		result    float32
	)
	amplitude = 1
	max = 0
	result = 0
	for octaves > 0 {
		max += amplitude
		result += Noise(fx, fy) * amplitude
		amplitude *= persistence
		fx *= frequency
		fy *= frequency
		octaves--
	}
	return result / max

}
