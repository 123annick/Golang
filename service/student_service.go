package service

import (
	"errors"
	"fmt"

	"go-learning-task/models"
)

// AddStudent appends a new student after basic validation.
func AddStudent(students []models.Student, student models.Student) ([]models.Student, error) {
	if student.Name == "" || student.Email == "" {
		return students, errors.New("name and email are required")
	}
	return append(students, student), nil
}

// GetStudentByID returns a pointer to the matching student, or an error.
func GetStudentByID(students []models.Student, id int) (*models.Student, error) {
	for i := range students {
		if students[i].ID == id {
			return &students[i], nil
		}
	}
	return nil, fmt.Errorf("student with ID %d not found", id)
}

// UpdateStudent finds a student by ID and overwrites its fields.
func UpdateStudent(students []models.Student, id int, updated models.Student) error {
	for i := range students {
		if students[i].ID == id {
			students[i] = updated
			return nil
		}
	}
	return fmt.Errorf("student with ID %d not found", id)
}

// DeleteStudent removes a student by ID.
func DeleteStudent(students []models.Student, id int) ([]models.Student, error) {
	for i, s := range students {
		if s.ID == id {
			return append(students[:i], students[i+1:]...), nil
		}
	}
	return students, fmt.Errorf("student with ID %d not found", id)
}

// CalculateAverage returns the mean score, or 0 if there are no students.
func CalculateAverage(students []models.Student) float64 {
	if len(students) == 0 {
		return 0
	}
	var total float64
	for _, s := range students {
		total += s.Score
	}
	return total / float64(len(students))
}

// PassedStudents returns students who scored at least 50.
func PassedStudents(students []models.Student) []models.Student {
	var passed []models.Student
	for _, s := range students {
		if s.Score >= 50 {
			passed = append(passed, s)
		}
	}
	return passed
}
