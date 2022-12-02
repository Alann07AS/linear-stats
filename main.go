package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	dataByte, err := os.ReadFile("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	data := strings.Fields(string(dataByte))
	someX := 0.0
	someXQqrt := 0.0
	someY := 0.0
	someYQqrt := 0.0
	someXY := 0.0
	for i, strNb := range data {
		nb, _ := strconv.Atoi(strNb)
		fNb := float64(nb)
		fI := float64(i)
		someY += fNb
		someYQqrt += fNb * fNb
		someX += fI
		someXQqrt += fI * fI
		someXY += fNb * fI
	}
	n := float64(len(data))
	b := (n*someXY - someX*someY) / (n*someXQqrt - someX*someX)
	a := (someY - b*(someX)) / (n)
	c := (n*someXY - someX*someY) / math.Sqrt((n*(someXQqrt)-someX*someX)*(n*(someYQqrt)-someY*someY))
	fmt.Printf("Linear Regression Line: y = %fx + %f \nPearson Correlation Coefficient: "+fmt.Sprint(round10(c))+"\n", round6(b), round6(a))
}

func round10(nb float64) float64 {
	return math.Round(nb*math.Pow10(10)) / math.Pow10(10)
}

func round6(nb float64) float64 {
	return math.Round(nb*math.Pow10(6)) / math.Pow10(6)
}
