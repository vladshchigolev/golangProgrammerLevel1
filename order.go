package main

import "fmt"

func main() {
	var drink, meal, subMeal, time string
	fmt.Scan(&drink, &meal, &subMeal, &time)
	fmt.Print("I wanna some '", drink, "', my favorite meal - '", meal, "'. Give me also '", subMeal, "'. I will come soon... '", time, "'")
}