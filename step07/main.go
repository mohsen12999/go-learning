package main

import "fmt"

type Developer struct {
	Name       string
	Role       string
	Experience int
	Language   string
	Salary     float64
	Active     bool
}

func promoteDeveloper(developer *Developer) {
	developer.Role = "Senior Developer"
	developer.Experience++
}

func printDeveloper(developer Developer) {
	fmt.Println("Name:", developer.Name)
	fmt.Println("Role", developer.Role)
	fmt.Println("Experience:", developer.Experience)
}

func main() {

	mohsen := Developer{
		Name:       "Mohsen",
		Role:       "Developer",
		Experience: 10,
		Language:   "GO",
		Salary:     5000.5,
		Active:     true,
	}

	promoteDeveloper(&mohsen)

	printDeveloper(mohsen)
}
