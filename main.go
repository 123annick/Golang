package main

import (
	"fmt"

	"go-learning-task/models"
)

func main() {
	student := models.Student{
		ID:     1,
		Name:   "Annick",
		Email:  "annick@example.com",
		Course: "Software Engineering",
		Score:  90,
	}

	fmt.Println(student)
}
