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

func (developer Developer) IsSenior() bool {
	return developer.Experience >= 7
}

func (developer *Developer) Promote() {
	developer.Role = "Senior Developer"
	developer.Experience++
}

func (developer Developer) Print() {
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

	mohsen.Promote()

	mohsen.Print()

	fmt.Println("Senior:", mohsen.IsSenior())
}
