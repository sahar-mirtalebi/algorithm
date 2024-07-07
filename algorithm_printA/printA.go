package main

import "fmt"

func main() {
	var num int
	fmt.Scan(&num)
	var output = "A"

	for n := 0; n < num; n++ {
		fmt.Println(output)
		output += "A"
	}
}
