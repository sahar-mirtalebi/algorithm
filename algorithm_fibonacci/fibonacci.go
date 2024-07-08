package main

import "fmt"

func main() {

	var numA = 0
	var numB = 1
	var numList = []int{numA, numB}
	var inputNum int
	fmt.Scan(&inputNum)

	for num := 0; num < 18; num++ {
		var nuwNum = numA + numB
		if nuwNum <= inputNum {
			numList = append(numList, nuwNum)
			numA = numB
			numB = nuwNum
		} else {
			break
		}
	}
	fmt.Printf("the resault of %v:\n", inputNum)
	fmt.Printf("series %v\n", numList)
}
