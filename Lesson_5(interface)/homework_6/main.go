package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Displayer interface {
	Display()
}

type Employee struct {
	FirstName  string
	LastName   string
	Age        int
	Position   string
	Salary     int
	Department string
}

func (e Employee) Display() {
	fmt.Printf("Имя Фамилия: %s %s, Возраст: %d, Должность: %s, Зарплата: %d тыс.руб/мес., Отдел: %s\n",
		e.FirstName, e.LastName, e.Age, e.Position, e.Salary, e.Department)
}

func addEmployee(employees []Employee, employee Employee) []Employee {
	return append(employees, employee)
}

func filterByAge(employees []Employee, age int) []Employee {
	var filteredEmployees []Employee
	for _, employee := range employees {
		if age >= 40 {
			if employee.Age > age {
				filteredEmployees = append(filteredEmployees, employee)
			}
		} else {
			if employee.Age == age {
				filteredEmployees = append(filteredEmployees, employee)
			}
		}
	}
	return filteredEmployees
}

func filterBySalary(employees []Employee, minSalary int) []Employee {
	var filteredEmployees []Employee
	for _, employee := range employees {
		if employee.Salary >= minSalary {
			filteredEmployees = append(filteredEmployees, employee)
		}
	}
	return filteredEmployees
}

func filterByDepartment(employees []Employee, department string) []Employee {
	var filteredEmployees []Employee
	for _, employee := range employees {
		if department == employee.Department {
			filteredEmployees = append(filteredEmployees, employee)
		}
	}
	return filteredEmployees
}

func printDisplayers(employees []Displayer) {
	for _, employee := range employees {
		employee.Display()
	}
}

func convertToDisplayers(employees []Employee) []Displayer {
	var items []Displayer
	for _, employee := range employees {
		items = append(items, employee)
	}
	return items
}

func main() {
	employees := []Employee{
		{LastName: "Ivanov", FirstName: "Ivan", Age: 60, Position: "Backend Developer", Salary: 180, Department: "IT"},
		{LastName: "Petrova", FirstName: "Anna", Age: 32, Position: "HR Manager", Salary: 120, Department: "HR"},
		{LastName: "Sidorov", FirstName: "Pavel", Age: 35, Position: "System Analyst", Salary: 200, Department: "IT"},
		{LastName: "Smirnova", FirstName: "Elena", Age: 27, Position: "QA Engineer", Salary: 140, Department: "IT"},
		{LastName: "Kuznetsov", FirstName: "Dmitry", Age: 40, Position: "Team Lead", Salary: 250, Department: "IT"},
		{LastName: "Volkova", FirstName: "Maria", Age: 50, Position: "Accountant", Salary: 110, Department: "Finance"},
	}

	newEmployees := Employee{LastName: "Orlov", FirstName: "Sergey", Age: 45, Position: "Warehouse Manager", Salary: 160, Department: "Logistics"}
	employees = addEmployee(employees, newEmployees)

	fmt.Println("Полный список сотрудников: ")
	printDisplayers(convertToDisplayers(employees))

	reader := bufio.NewReader(os.Stdin)
	var age int
	var minSalary int
	fmt.Print("Введите возраст для фильтрации: ")
	fmt.Scanln(&age)
	fmt.Print("Введите зарплату для фильтрации: ")
	fmt.Scanln(&minSalary)
	fmt.Print("Введите название отдела для фильтрации: ")
	department, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Ошибка чтения названия отдела:", err)
	}
	department = strings.TrimSpace(department)

	filteredAge := filterByAge(employees, age)
	filteredSalary := filterBySalary(employees, minSalary)
	filteredDepartment := filterByDepartment(employees, department)

	fmt.Println("Фильтрация по возрасту:")
	printDisplayers(convertToDisplayers(filteredAge))

	fmt.Println("Фильтрация по зарплате: ")
	printDisplayers(convertToDisplayers(filteredSalary))

	fmt.Println("Фильтрация по отделу: ")
	printDisplayers(convertToDisplayers(filteredDepartment))
}
