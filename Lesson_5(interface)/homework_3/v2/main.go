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
	fmt.Printf("Автомобиль: %s %s, Год: %d, Пробег: %d тыс.км, Цена: %d тыс.руб.\n",
		c.Brand, c.Model, c.Year, c.Mileage, c.Price)
}

type Motorcycle struct {
	Brand   string
	Model   string
	Year    int
	Mileage int
	Price   int
}

func (m Motorcycle) Print() {
	fmt.Printf("Мотоцикл: %s %s, Год: %d, Пробег: %d тыс.км, Цена: %d тыс.руб.\n",
		m.Brand, m.Model, m.Year, m.Mileage, m.Price)
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

func filterByMileage(cars []Car, maxMileage int) []Car {
	var filteredCars []Car
	for _, car := range cars {
		if car.Mileage <= maxMileage {
			filteredCars = append(filteredCars, car)
		}
	}
	return filteredCars
}

func printVehicles(items []Printable) {
	for _, item := range items {
		item.Print()
	}
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

	minYear := 2015
	maxMileage := 121

	filteredByYear := filterByYear(cars, minYear)
	filteredByMileage := filterByMileage(cars, maxMileage)

	var vehicles []Printable

	for _, car := range cars {
		vehicles = append(vehicles, car)
	}

	bike := Motorcycle{"Yamaha", "R1", 2022, 8, 1500}
	vehicles = append(vehicles, bike)

	fmt.Println("Весь транспорт:")
	printVehicles(vehicles)

	fmt.Printf("\nАвтомобили новее %d:\n", minYear)
	var yearVehicles []Printable
	for _, car := range filteredByYear {
		yearVehicles = append(yearVehicles, car)
	}
	printVehicles(yearVehicles)

	fmt.Printf("\nАвтомобили с пробегом меньше %d тыс.км:\n", maxMileage)
	var mileageVehicles []Printable
	for _, car := range filteredByMileage {
		mileageVehicles = append(mileageVehicles, car)
	}
	printVehicles(mileageVehicles)
}
