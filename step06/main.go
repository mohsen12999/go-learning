package main

import (
	"fmt"
)

type Developer struct {
	Name       string
	Role       string
	Experience int
	Language   string
	Salary     float64
	Active     bool
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
	ali := Developer{
		Name:       "Ali",
		Role:       "Tester",
		Experience: 5,
		Language:   "C#",
		Salary:     25000.5,
		Active:     false,
	}

	developers := []Developer{mohsen, ali}

	for _, developer := range developers {
		fmt.Println(developer.Name)
		fmt.Println(developer.Language)
	}
}
