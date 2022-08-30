package main

import (
	"projectone/assignment1/helper"
	"fmt"
	"os"
	"strconv"
)

func main() {

	if toGetAllArgs := os.Args[1:]; len(toGetAllArgs) > 0 {
		
		getIndex, errIndex := strconv.Atoi(toGetAllArgs[0])
		if errIndex != nil {
			fmt.Println(errIndex.Error())
			return
		}

		getArray := helper.GetTeman(getIndex)
		fmt.Println(getArray)
	}
	

}