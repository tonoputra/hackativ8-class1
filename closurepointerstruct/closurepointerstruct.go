package closurepointerstruct

import (
	"fmt"
)

func GetTemanCall() {

	temanList := []string{
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

	var tmpTemanList []*Teman
	for _, v := range temanList {
		tmpTemanList = append(tmpTemanList, &Teman{
			NamaPanjang: v,
		})
	}

	// create closure func
	GetTeman := func(teman []*Teman) {
		for _, v := range teman {
			fullName := v.NamaPanjang
			fmt.Println("hasil pointer dan closure: ", fullName)
		}
	}

	GetTeman(tmpTemanList)
}
