package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type celsius float64
type fahrenheit float64

type pounds float64
type kilos float64

func cToF(c celsius) fahrenheit { return fahrenheit(c*9/5 + 32) }
func fToC(f fahrenheit) celsius { return celsius((f - 32) * 5 / 9) }

func pToK(p pounds) kilos { return kilos(0.454592 * p) }
func kToP(k kilos) pounds { return pounds(2.20462 * k) }

func conversions(valueStr string) [4]string {
	value, err := strconv.ParseFloat(valueStr, 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "\"%v\" is not a valid float64. Error: %s\n", valueStr, err)
		os.Exit(1)
	}

	baseString := fmt.Sprint(value, " %v is %v %v")

	cfString := fmt.Sprintf(baseString, "°C", cToF(celsius(value)), "°F")
	fcString := fmt.Sprintf(baseString, "°F", fToC(fahrenheit(value)), "°C")

	pkString := fmt.Sprintf(baseString, "lbs", pToK(pounds(value)), "kg")
	kpString := fmt.Sprintf(baseString, "kg", kToP(kilos(value)), "lbs")

	return [4]string{cfString, fcString, pkString, kpString}
}

func main() {
	var scanner *bufio.Scanner
	if len(os.Args) == 1 {
		scanner = bufio.NewScanner(os.Stdin)
	} else {
		joinedString := strings.Join(os.Args, " ")
		scanner = bufio.NewScanner(strings.NewReader(joinedString))
		scanner.Split(bufio.ScanWords)
		scanner.Scan() // Get rid of program arg.
	}

	for scanner.Scan() {
		for _, outstring := range conversions(scanner.Text()) {
			fmt.Println(outstring)
		}
		fmt.Println()
	}
}
