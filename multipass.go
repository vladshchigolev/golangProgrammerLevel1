package main

import "fmt"

func main() {
	var firstName, lastName string
	var age int
	fmt.Scan(&firstName, &lastName, &age)
	fmt.Println("Имя:", firstName, ", Фамилия:", lastName, ", Возраст:", age)
}
