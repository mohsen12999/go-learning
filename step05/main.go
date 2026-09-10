package main

import (
	"fmt"
	"strconv"
)

func main() {

	name := "Mohsen"
	exp := 5

	languages := []string{
		"Go",
		"C#",
		"TypeScript",
		"PHP",
		"Python",
	}

	for _, language := range languages {
		fmt.Println(language)
	}

	developer := map[string]string{
		"name":       name,
		"role":       "Backend Developer",
		"language":   languages[0],
		"experience": strconv.Itoa(exp),
	}

	fmt.Println(developer["language"])

	languages = append(languages, "Java")

	for _, language := range languages {
		fmt.Println(language)
	}

}
