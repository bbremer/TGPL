package main

import (
	"fmt"
	"temperature/tempconv"
)

func main() {
	fmt.Printf("Brrr! %v\n", tempconv.AbsoluteZeroC)
	fmt.Println(tempconv.CToF(tempconv.BoilingC))
	fmt.Printf("0°C is %v.\n", tempconv.CToK(0))
	fmt.Printf("300°K is %v\n", tempconv.KToF(300))
	tempconv.TestF()
}
