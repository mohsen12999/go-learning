package main

import (
	"fmt"
	"myapp/developer"
)

func main() {
	mohsen := developer.NewDeveloper(
		"Mohsen",
		"Backend Developer",
		10,
	)

	fmt.Println(mohsen.Name)
	fmt.Println(mohsen.IsSenior())
}
