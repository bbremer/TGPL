package main

import (
	"fmt"
	"os"
	"strconv"
)

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func PopCount(x uint64) int {
	count := byte(0)
	for i := 0; i < 8; i++ {
		count += pc[byte(x>>(i*8))]
	}
	return int(count)
}

func main() {
	for _, xStr := range os.Args[1:] {
		x, err := strconv.ParseUint(xStr, 10, 64)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Println(x, PopCount(x))
	}
}
