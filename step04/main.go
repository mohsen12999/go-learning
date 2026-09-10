package main

import "fmt"

func main() {

	name := "Mohsen"
	age := 30
	exp := 5
	developer := true
	salary := 5000.5

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
	fmt.Println("Years until 40:", yearsUntil40(age))

	if age < 18 {
		fmt.Println("Status: Child")
	} else {
		fmt.Println("Status: Adult")
	}

	if developer {
		fmt.Println("Profession: Developer")
	}

	fmt.Println("Level:", getExperienceLevel(exp))

	fmt.Println("\nExperience timeline:")

	for i := 1; i <= 5; i++ {
		fmt.Println("Year", i)
	}

	fmt.Println("")

	fmt.Print("Experience category:", getExperienceLevel(exp))

	fmt.Println("\n========================")

}

func getExperienceLevel(exp int) string {
	switch exp {
	case 0:
		return "Beginner"
	case 1, 2:
		return "Junior"
	case 3, 4, 5, 6:
		return "Mid-level"
	default:
		return "Senior"
	}
}

func yearsUntil40(age int) int {
	return 40 - age
}
