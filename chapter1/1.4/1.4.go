package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)
	duplines_to_files := make(map[string][]string)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines("stdin", counts, duplines_to_files)
	} else {
		for _, arg := range files {
			countLines(arg, counts, duplines_to_files)
		}
	}
	for line, n := range counts {
		if n > 1 {
            fns := strings.Join(duplines_to_files[line], " ")
			fmt.Printf("%d\t%s\t%s\n", n, line, fns)
		}
	}
}

func countLines(fn string, counts map[string]int, duplines_to_files map[string][]string) {
    var f *os.File
    if fn == "stdin" {
        f = os.Stdin
    } else {
        var err error 
        f, err = os.Open(fn)
        if err != nil{
            fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
        }
        return
    }
	input := bufio.NewScanner(f)
	for input.Scan() {
        line := input.Text()
		counts[line]++
		duplines_to_files[line] = append(duplines_to_files[line], fn)
	}
	if fn != "stdin" {
	    f.Close()
	}
}
