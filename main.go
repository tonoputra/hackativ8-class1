package main

import (
	"projectone/closurepointerstruct"
	"projectone/ganjilgenap"
	"projectone/sliceloopingkenalan"
)

func main() {
	// print ganjil genap
	ganjilgenap.GetGanjilGenap(10)

	// slice + loop + kenalan
	sliceloopingkenalan.GetSliceKenalan()

	// closure + pointer + struct
	closurepointerstruct.GetTemanCall()
}