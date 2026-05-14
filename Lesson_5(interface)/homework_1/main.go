package main

import "fmt"

type Displayer interface {
	Display()
}
type Student struct {
	FirstName    string
	LastName     string
	Age          int
	Course       int
	AverageGrade float64
}

func (s Student) Display() {
	fmt.Printf("%s %s, курс: %d, возраст: %d, средний бал: %.1f\n",
		s.FirstName, s.LastName, s.Course, s.Age, s.AverageGrade)
}

func addStudent(students []Student, student Student) []Student {
	return append(students, student)
}

func filterByMinGrade(students []Student, minGrade float64) []Student {
	var filtered []Student
	for _, student := range students {
		if student.AverageGrade >= minGrade {
			filtered = append(filtered, student)
		}
	}
	return filtered
}

func printStudents(ds []Displayer) {
	for _, d := range ds {
		d.Display()
	}
}

func main() {
	students := []Student{
		{
			LastName:     "Иванов",
			FirstName:    "Иван",
			Age:          19,
			Course:       2,
			AverageGrade: 4.5,
		},
		{
			LastName:     "Петров",
			FirstName:    "Пётр",
			Age:          20,
			Course:       3,
			AverageGrade: 3.8,
		},
		{
			LastName:     "Сидорова",
			FirstName:    "Анна",
			Age:          18,
			Course:       1,
			AverageGrade: 4.9,
		},
	}

	newStudent := Student{
		LastName:     "Васильев",
		FirstName:    "Петя",
		Age:          19,
		Course:       3,
		AverageGrade: 4.6,
	}

	students = addStudent(students, newStudent)

	fmt.Println("Все студенты:")
	for _, student := range students {
		student.Display()
	}

	minGrade := 4.0
	filterStudents := filterByMinGrade(students, minGrade)

	var displayList []Displayer
	for _, student := range filterStudents {
		displayList = append(displayList, student)
	}
	fmt.Printf("\nСтуденты со средним баллом не ниже: %.1f\n", minGrade)
	printStudents(displayList)
}
