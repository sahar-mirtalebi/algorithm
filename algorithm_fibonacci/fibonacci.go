package main

import "fmt"

var numA = 0
var numB = 1

func main() {
	var inputNum int
	fibonacci(inputNum)
}

func fibonacci(inputNum int) {
	fmt.Scan(&inputNum)
	var numList = []int{numA, numB}
	for num := 0; num < inputNum; num++ {
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
