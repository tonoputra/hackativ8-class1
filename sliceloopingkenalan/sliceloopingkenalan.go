package sliceloopingkenalan

import "fmt"

func GetSliceKenalan() {
	var temenList = []string{
		"Mocha Fiqri",
		"Rijal Khawarizmi",
		"Aditya Kristianto",
		"Clara Velita Pranolo",
		"Satrio Utomo",
		"Eka Widiantara",
		"Aulia Nurhady",
		"Brahmantyo Adi",
		"Ivan Muhanov",
		"Fatur Ewing Fadh",
	}

	for _, v := range temenList {
		fmt.Println(v)
	}
}