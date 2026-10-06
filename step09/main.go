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

type Worker interface {
	Describe() string
}

func (developer Developer) Describe() string {
	return developer.Name + " is a " + developer.Role
}

func printWorker(worker Worker) {
	fmt.Println(worker.Describe())
}

type Tester struct {
	Name string
}

func (tester Tester) Describe() string {
	return tester.Name + " is a Tester"
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

	ali := Tester{ Name: "Ali"}

	printWorker(mohsen)
	printWorker(ali)
}
