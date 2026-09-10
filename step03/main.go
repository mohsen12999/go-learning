package main

import "fmt"

func main()  {

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
fmt.Println("Years until 40:" , 40-age)

if age < 18 {
	fmt.Println("Status: Child")
} else {
	fmt.Println("Status: Adult")
}

if developer {
	fmt.Println("Profession: Developer")
}

switch exp {
	case 0:
		fmt.Println("Level: Beginner")
	case 1, 2:
		fmt.Println("Level: Junior")
	case 3, 4, 5, 6:
		fmt.Println("Level: Mid-level")
	default:
		fmt.Println("Level: Senior")
}

fmt.Println("\nExperience timeline:")

for i := 1; i <= 5; i++ {
	fmt.Println("Year", i)
}

fmt.Println("")

fmt.Print("Experience category: ")


switch exp {
case 0:
    fmt.Print("Beginner")
case 1, 2:
    fmt.Print("Junior")
case 3, 4, 5, 6:
    fmt.Print("Mid-level")
default:
    fmt.Print("Senior")
}

fmt.Println("\n========================")

}