package main

import "fmt"

func main() {

	var numA = 0
	var numB = 1
	var numList = []int{}

	for num := 0; num < 18; num++ {
		var nuwNum = numA + numB
		numList = append(numList, nuwNum)
		numA = numB
		numB = nuwNum
	}
	fmt.Println(numList)
}
