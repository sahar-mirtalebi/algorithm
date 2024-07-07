package main

import "fmt"

func main() {
	var numA int
	var numB int

	fmt.Scan(&numA)
	fmt.Scan(&numB)

	var poweredNum = numA
	var addedNum = numA

	for n := 1; n < numB; n++ {
		for m := 1; m < numA; m++ {
			poweredNum = poweredNum + addedNum
		}
		addedNum = poweredNum
	}
	fmt.Println(poweredNum)
}
