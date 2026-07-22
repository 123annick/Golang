package main

import (
	"fmt"

	"go-learning-task/models"
	"go-learning-task/service"
)

func main() {
	var students []models.Student

	students, err := service.AddStudent(students, models.Student{
		ID:     1,
		Name:   "Annick",
		Email:  "annick@example.com",
		Course: "Software Engineering",
		Score:  90,
	})
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Students:", students)
	fmt.Println("Average score:", service.CalculateAverage(students))
}
