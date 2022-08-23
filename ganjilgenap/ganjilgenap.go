package ganjilgenap

import "fmt"

func GetGanjilGenap(n int) {
	for i := 0; i < n; i++ {
		if i%2 == 0 {
			fmt.Println("genap")
		} else {
			fmt.Println("ganjil")
		}
	}
}