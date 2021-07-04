package main

import "fmt"

func main() {
	var firstNum, secondNum float32
	fmt.Scan(&firstNum, &secondNum)
	sum := int(firstNum + secondNum)
	if (sum % 2) == 0 {
		fmt.Print("ЧЁТНОЕ")
	} else {
		fmt.Print("НЕЧЁТНОЕ")
	}
}
