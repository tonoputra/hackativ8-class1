package main

import (
	"fmt"
	"math/rand"
	"time"
)

var (
	breakAt int = 32
	chance  int = 1000
)

func main() {
	ch := make(chan string)
	defer close(ch)

	rand.Seed(time.Now().UnixNano())

	memberList := []string{"tono", "tina", "tini", "tunu"}

	for i := 0; i < chance-1; i++ {
		for _, v := range memberList {
			go func(name string) {
				ch <- name
			}(v)
		}

		setRand := rand.Int()/int(time.Now().UnixNano())
		mssg := ""
		if setRand%breakAt == 0 {
			mssg = fmt.Sprintf("chance ke-%d; member %s kalah", i, <-ch)
			fmt.Println(mssg)
			return
		}
		mssg = fmt.Sprintf("chance ke-%d dan nilai %d; korek di %s", i, setRand, <-ch)
		fmt.Println(mssg)
	}
}
