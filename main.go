package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"go-learning-task/models"
	"go-learning-task/service"
)

func main() {
	var students []models.Student
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\n--- Student Management CLI ---")
		fmt.Println("1. Add student")
		fmt.Println("2. View all students")
		fmt.Println("3. Find student by ID")
		fmt.Println("4. Update student")
		fmt.Println("5. Delete student")
		fmt.Println("6. Calculate average score")
		fmt.Println("7. Show passed students")
		fmt.Println("8. Exit")
		fmt.Print("Choose an option: ")

		choice := readLine(reader)

		switch choice {
		case "1":
			students = handleAdd(reader, students)
		case "2":
			handleViewAll(students)
		case "3":
			handleFind(reader, students)
		case "4":
			handleUpdate(reader, students)
		case "5":
			students = handleDelete(reader, students)
		case "6":
			fmt.Printf("Average score: %.2f\n", service.CalculateAverage(students))
		case "7":
			handlePassed(students)
		case "8":
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid option, try again.")
		}
	}
}

// readLine reads one line of input and trims the trailing newline/whitespace.
func readLine(reader *bufio.Reader) string {
	text, _ := reader.ReadString('\n')
	return strings.TrimSpace(text)
}

// readInt reads a line and parses it as an int, looping until valid.
func readInt(reader *bufio.Reader, prompt string) int {
	for {
		fmt.Print(prompt)
		value, err := strconv.Atoi(readLine(reader))
		if err != nil {
			fmt.Println("Please enter a valid whole number.")
			continue
		}
		return value
	}
}

// readFloat reads a line and parses it as a float64, looping until valid.
func readFloat(reader *bufio.Reader, prompt string) float64 {
	for {
		fmt.Print(prompt)
		value, err := strconv.ParseFloat(readLine(reader), 64)
		if err != nil {
			fmt.Println("Please enter a valid number.")
			continue
		}
		return value
	}
}

func handleAdd(reader *bufio.Reader, students []models.Student) []models.Student {
	id := readInt(reader, "ID: ")
	fmt.Print("Name: ")
	name := readLine(reader)
	fmt.Print("Email: ")
	email := readLine(reader)
	fmt.Print("Course: ")
	course := readLine(reader)
	score := readFloat(reader, "Score: ")

	updated, err := service.AddStudent(students, models.Student{
		ID:     id,
		Name:   name,
		Email:  email,
		Course: course,
		Score:  score,
	})
	if err != nil {
		fmt.Println("Error:", err)
		return students
	}
	fmt.Println("Student added.")
	return updated
}

func handleViewAll(students []models.Student) {
	if len(students) == 0 {
		fmt.Println("No students yet.")
		return
	}
	for _, s := range students {
		fmt.Printf("%+v\n", s)
	}
}

func handleFind(reader *bufio.Reader, students []models.Student) {
	id := readInt(reader, "ID to find: ")
	student, err := service.GetStudentByID(students, id)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("%+v\n", *student)
}

func handleUpdate(reader *bufio.Reader, students []models.Student) {
	id := readInt(reader, "ID to update: ")
	fmt.Print("New name: ")
	name := readLine(reader)
	fmt.Print("New email: ")
	email := readLine(reader)
	fmt.Print("New course: ")
	course := readLine(reader)
	score := readFloat(reader, "New score: ")

	err := service.UpdateStudent(students, id, models.Student{
		ID:     id,
		Name:   name,
		Email:  email,
		Course: course,
		Score:  score,
	})
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Student updated.")
}

func handleDelete(reader *bufio.Reader, students []models.Student) []models.Student {
	id := readInt(reader, "ID to delete: ")
	updated, err := service.DeleteStudent(students, id)
	if err != nil {
		fmt.Println("Error:", err)
		return students
	}
	fmt.Println("Student deleted.")
	return updated
}

func handlePassed(students []models.Student) {
	passed := service.PassedStudents(students)
	if len(passed) == 0 {
		fmt.Println("No students have passed yet.")
		return
	}
	for _, s := range passed {
		fmt.Printf("%+v\n", s)
	}
}
