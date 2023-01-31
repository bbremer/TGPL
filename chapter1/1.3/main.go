package main

import (
    "fmt"
    "os"
    "strings"
    "time"
)

func echo1() {
    var s, sep string
    for i := 1; i < len(os.Args); i++ {
        s += sep + os.Args[i]
        sep = " "
    }
    fmt.Println(s)
}

func echo2() {
    s, sep := "", ""
    for _, arg := range os.Args[1:] {
        s += sep + arg
        sep = " "
    }
    fmt.Println(s)
}

func echo3() {
    fmt.Println(strings.Join(os.Args[1:], " "))
}

func main() {
    REPEATS := 10000  // Larger runs into deadlock with w.

    rescueStdout := os.Stdout
    _, w, _ := os.Pipe()
    os.Stdout = w

    s := time.Now()
    for i := 0; i < REPEATS; i++ {
        echo1()
    }
    echo1_time := time.Since(s)
    fmt.Fprintln(rescueStdout, "echo1 done")

    s = time.Now()
    for i := 0; i < REPEATS; i++ {
        echo2()
    }
    echo2_time := time.Since(s)
    fmt.Fprintln(rescueStdout, "echo2 done")

    s = time.Now()
    for i := 0; i < REPEATS; i++ {
        echo3()
    }
    echo3_time := time.Since(s)
    fmt.Fprintln(rescueStdout, "echo3 done")

    os.Stdout = rescueStdout
    fmt.Println("Echo 1:", echo1_time)
    fmt.Println("Echo 2:", echo2_time)
    fmt.Println("Echo 3:", echo3_time)
}
