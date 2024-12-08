package main

import "fmt"

func add(n1 int8, n2 int8) int8 {
	return n1 + n2
}

//Multi return function
func mathCalc(n1, n2 int8) (int8, int8) {
	sum := n1 + n2
	sub := n1 - n2
	return sum, sub
}

func main() {
	sum := add(10, 20)
	fmt.Println(sum)

	var f = func(txt string) string{
		fmt.Println(txt)
		return txt
	}

	result := f("Function one text")
	fmt.Println(result)

	sumResult, _ := mathCalc(10, 15)
	fmt.Println(sumResult)

	_, subResult := mathCalc(10, 15)
	fmt.Println(subResult)


}
