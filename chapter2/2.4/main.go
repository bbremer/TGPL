package main

import (
	"fmt"
	"os"
	"strconv"
)

func PopCount(x uint64) int {
	var count uint64 = 0
	for i := 0; i < 64; i++ {
		if x&(1<<i) != 0 {
			count++
		}
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
