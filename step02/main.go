package main

import "fmt"

func main()  {
	var name string = "Mohsen"
	var age int = 30
	var exp int = 5
	var developer bool = true
	var salary float64 = 5000.50

	const language = "GO"

	fmt.Println("========================")
	fmt.Println("     GO DEVELOPER")
	fmt.Println("========================")
	fmt.Println("Name:", name)
	fmt.Println("Age:", age)
	fmt.Println("Experience:", exp)
	fmt.Println("Backend Developer:", developer)
	fmt.Println("Salary:", salary)
	fmt.Println("Language:", language)
	fmt.Println("Years until 40:", (40-age))
	fmt.Println("========================")
}