package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Displayable interface {
	Display()
	GetAge() int
	GetStatus() string
}

type Animal struct {
	Name    string
	Species string
	Age     int
	Weight  float64
	Status  string
}

type Dog struct {
	Animal
}

type Cat struct {
	Animal
}

func (a Animal) Display() {
	fmt.Printf("Кличка: %s, Вид: %s, Возраст: %d, Вес: %.2f, Статус: %s\n",
		a.Name, a.Species, a.Age, a.Weight, a.Status)
}

func (a Animal) GetAge() int {
	return a.Age
}

func (a Animal) GetStatus() string {
	return a.Status
}

func addAnimal(animals []Displayable, animal Displayable) []Displayable {
	return append(animals, animal)
}

func filterAnimalsByAge(animals []Displayable, filteredAge int) []Displayable {
	var filteredAnimalsByAge []Displayable
	for _, a := range animals {
		if a.GetAge() >= filteredAge {
			filteredAnimalsByAge = append(filteredAnimalsByAge, a)
		}
	}
	return filteredAnimalsByAge
}

func filterAnimalsByStatus(animals []Displayable, filteredStatus string) []Displayable {
	var filteredAnimalsByStatus []Displayable
	for _, a := range animals {
		if a.GetStatus() == filteredStatus {
			filteredAnimalsByStatus = append(filteredAnimalsByStatus, a)
		}
	}
	return filteredAnimalsByStatus
}

func printDisplayable(animals []Displayable) {
	for _, a := range animals {
		a.Display()
	}
}

func main() {
	animals := []Displayable{
		Cat{
			Animal: Animal{Name: "Соня", Species: "кошка", Age: 1, Weight: 3.1, Status: "ищет дом"},
		},
		Dog{
			Animal: Animal{Name: "Бим", Species: "собака", Age: 4, Weight: 18.5, Status: "ищет дом"},
		},
		Cat{
			Animal: Animal{Name: "Мурка", Species: "кошка", Age: 2, Weight: 4.2, Status: "лечится"},
		},
		Dog{
			Animal: Animal{Name: "Рекс", Species: "собака", Age: 7, Weight: 25.0, Status: "забронирован"},
		},
	}

	newAnimal := Dog{
		Animal: Animal{Name: "Шарик", Species: "собака", Age: 5, Weight: 14.8, Status: "лечится"},
	}
	animals = addAnimal(animals, newAnimal)

	fmt.Println("Полный список животных:")
	printDisplayable(animals)

	var minAge int

	fmt.Print("Введите возраст, для фильтрации: ")
	fmt.Scanln(&minAge)

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Введите статус для фильтрации: ")
	status, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Ошибка чтения статуса", err)
		return
	}
	status = strings.TrimSpace(status)

	filteredAges := filterAnimalsByAge(animals, minAge)
	filteredStatus := filterAnimalsByStatus(animals, status)

	fmt.Printf("Животные старше %d лет:\n", minAge)
	printDisplayable(filteredAges)

	fmt.Printf("Животные со статусом - %s:\n", status)
	printDisplayable(filteredStatus)
}
