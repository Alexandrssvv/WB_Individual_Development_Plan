package main

import "fmt"

type Printable interface {
	Print()
}

type Car struct {
	Brand   string
	Model   string
	Year    int
	Mileage int
	Price   int
}

func (c Car) Print() {
	fmt.Printf("Марка: %s, Модель: %s, Год выпуска: %d, Пробег: %d тыс.км, Цена: %d тыс.руб.\n",
		c.Brand, c.Model, c.Year, c.Mileage, c.Price)
}

func addCar(cars []Car, car Car) []Car {
	return append(cars, car)
}

func filterByYear(cars []Car, minYear int) []Car {
	var filteredCars []Car
	for _, car := range cars {
		if car.Year >= minYear {
			filteredCars = append(filteredCars, car)
		}
	}
	return filteredCars
}

func filterByMileage(cars []Car, minMileage int) []Car {
	var filteredCars []Car
	for _, car := range cars {
		if car.Mileage <= minMileage {
			filteredCars = append(filteredCars, car)
		}
	}
	return filteredCars
}

func printCars(items []Printable) {
	for _, item := range items {
		item.Print()
	}
}

func toPrintables(cars []Car) []Printable {
	var items []Printable
	for _, car := range cars {
		items = append(items, car)
	}
	return items
}

func main() {

	cars := []Car{
		{"Toyota", "Camry", 2021, 36, 1200},
		{"Lexus", "IS250", 2011, 180, 800},
		{"Toyota", "Prius", 2014, 120, 600},
		{"Ford", "Focus", 2010, 210, 500},
	}

	newCar := Car{"Ford", "Mustang", 2000, 350, 2000}
	cars = addCar(cars, newCar)

	fmt.Println("Все автомобили:")
	printCars(toPrintables(cars))

	minYear := 2015
	filterYear := filterByYear(cars, minYear)
	fmt.Printf("Автомобили новее %d\n", minYear)
	printCars(toPrintables(filterYear))

	minMileage := 121
	filterMileage := filterByMileage(cars, minMileage)
	fmt.Printf("Автомобили с пробегом меньше %d\n", minMileage)
	printCars(toPrintables(filterMileage))
}
