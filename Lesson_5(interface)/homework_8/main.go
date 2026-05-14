package main

import "fmt"

type Displayable interface {
	Display()
}

type Course struct {
	Name          string
	Instructor    string
	DurationHours int
	Price         float64
	StudentCount  int
}

func (c Course) Display() {
	fmt.Printf("Название курса: %s, Преподаватель: %s, Длительность: %d часов, Стоимость курса: %.1f, Количество студентов: %d\n",
		c.Name, c.Instructor, c.DurationHours, c.Price, c.StudentCount)
}

func addCourse(courses []Course, course Course) []Course {
	return append(courses, course)
}

func FilterByPrice(courses []Course, minPrice float64) []Course {
	var filteredCourse []Course
	for _, course := range courses {
		if course.Price >= minPrice {
			filteredCourse = append(filteredCourse, course)
		}
	}
	return filteredCourse
}

func FilterByDuration(courses []Course, minDuration int) []Course {
	var filteredCourse []Course
	for _, course := range courses {
		if course.DurationHours >= minDuration {
			filteredCourse = append(filteredCourse, course)
		}
	}
	return filteredCourse
}

func printDisplayable(courses []Displayable) {
	for _, course := range courses {
		course.Display()
	}
}

func convertToDisplayable(courses []Course) []Displayable {
	var displayable []Displayable
	for _, course := range courses {
		displayable = append(displayable, course)
	}
	return displayable
}

func main() {
	courses := []Course{
		{Name: "Go для начинающих", Instructor: "Иван Петров", DurationHours: 40, Price: 15000, StudentCount: 120},
		{Name: "Продвинутый Go и конкурентность", Instructor: "Алексей Смирнов", DurationHours: 60, Price: 25000, StudentCount: 80},
		{Name: "Основы SQL и работа с базами данных", Instructor: "Мария Иванова", DurationHours: 35, Price: 12000, StudentCount: 150},
		{Name: "Системный анализ в IT", Instructor: "Дмитрий Кузнецов", DurationHours: 50, Price: 20000, StudentCount: 95},
		{Name: "Микросервисы и архитектура систем", Instructor: "Сергей Волков", DurationHours: 45, Price: 22000, StudentCount: 70},
		{Name: "Docker и Kubernetes с нуля", Instructor: "Ольга Соколова", DurationHours: 55, Price: 23000, StudentCount: 85},
		{Name: "Основы Python для анализа данных", Instructor: "Анна Морозова", DurationHours: 40, Price: 14000, StudentCount: 110},
	}

	newCourse := Course{Name: "Алгоритмы и структуры данных", Instructor: "Николай Орлов", DurationHours: 65, Price: 26000, StudentCount: 60}
	courses = addCourse(courses, newCourse)

	fmt.Println("Полный список курсов: ")
	printDisplayable(convertToDisplayable(courses))

	var minPrice float64
	var minDuration int

	fmt.Print("Введите минимальную стоимость курса: ")
	fmt.Scanln(&minPrice)
	fmt.Print("Введите минимальную продолжительность курса: ")
	fmt.Scanln(&minDuration)

	filteredPrice := FilterByPrice(courses, minPrice)
	filteredDuration := FilterByDuration(courses, minDuration)

	fmt.Printf("Курсы со стоимостью от %.1f руб.\n", minPrice)
	printDisplayable(convertToDisplayable(filteredPrice))

	fmt.Printf("Курсы с продолжительностью от %d часов\n", minDuration)
	printDisplayable(convertToDisplayable(filteredDuration))
}
