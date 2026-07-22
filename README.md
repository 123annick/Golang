# Student Management CLI Application

## Description

This project is a command-line Student Management Application developed using the Go programming language.

The application allows users to manage student records by adding, viewing, searching, updating, and deleting students. It also calculates the average student score and displays students who have passed.

A student is considered to have passed when their score is **50 or above**.

---

## Features

The application supports the following operations:

- Add a new student
- View all students
- Find a student by ID
- Update student information
- Delete a student
- Calculate the average student score
- Display students who passed

---

## Technologies Used

- Go Programming Language
- Git
- Visual Studio Code

---

## Project Structure

```
go-learning-task
│
├── main.go
├── go.mod
│
├── models
│   └── student.go
│
└── service
    └── student-service.go
```

### File Description

- **main.go**  
  Contains the main program, user interface, menu, and handles user input.

- **models/student.go**  
  Contains the Student structure that defines student information.

- **service/student-service.go**  
  Contains functions responsible for student operations such as adding, updating, deleting, searching, and calculating scores.

---

# How to Run the Application

## 1. Clone the repository

```bash
git clone https://github.com/123annick/Golang.git
```

## 2. Navigate to the project directory

```bash
cd go-learning-task
```

## 3. Run the application

```bash
go run .
```

The application will start and display a menu where users can select different student management operations.

---

# Go Concepts Explanation

## 1. Slices

A slice is a dynamic collection of elements in Go. Unlike arrays, slices can grow or shrink during program execution.

In this project, a slice is used to store multiple student records:

```go
var students []models.Student
```

The `append()` function is used to add new students to the slice.

Example:

```go
students = append(students, student)
```

This adds a new student to the existing list.

---

## 2. Maps

A map is a data structure in Go that stores information in key-value pairs.

Example:

```go
studentScores := map[string]float64{
    "Annick": 85,
}
```

In this example:

- `"Annick"` is the key.
- `85` is the value.

Maps are useful when data needs to be accessed quickly using a unique key.

Although maps are part of Go fundamentals, this application uses slices because students are managed as a list of records.

---

## 3. Structs

A struct is a custom data type in Go that groups related information together.

In this project, the `Student` struct represents a student:

```go
type Student struct {
    ID     int
    Name   string
    Email  string
    Course string
    Score  float64
}
```

The struct allows all student details to be stored together as one object.

---

## 4. Pointers

A pointer is a variable that stores the memory address of another variable.

In this project, the `GetStudentByID()` function returns a pointer:

```go
func GetStudentByID(students []Student, id int) (*Student, error)
```

The pointer allows the function to return the original student data instead of creating a copy.

Pointers are useful when working with data that needs to be accessed or modified directly.

---

## 5. Error Handling

Go handles errors by returning them from functions instead of using exceptions.

For example:

```go
student, err := service.GetStudentByID(students, id)
```

If the student exists, the function returns the student information.

If the student does not exist, it returns an error message:

```
student with ID not found
```

This allows the application to handle problems properly without crashing.

---

## Author

Developed as part of a Go programming fundamentals learning task.
