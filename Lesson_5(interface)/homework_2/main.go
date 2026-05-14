package main

import "fmt"

type Displayer interface {
	Display()
}

type Product struct {
	Name     string
	Category string
	Price    float64
	Quantity int
}

func (p Product) Display() {
	switch p.Category {
	case "Овощи", "Фрукты":
		fmt.Printf("Наименование: %s, Категория: %s, Цена: %.1f руб., Количество: %d кг\n",
			p.Name, p.Category, p.Price, p.Quantity)
	default:
		fmt.Printf("Наименование: %s, Категория: %s, Цена: %.1f руб., Количество: %d шт\n",
			p.Name, p.Category, p.Price, p.Quantity)
	}
}

func addProduct(products []Product, product Product) []Product {
	return append(products, product)
}

func filterByMinPrice(products []Product, minPrice float64) []Product {
	var filtered []Product

	for _, product := range products {
		if product.Price >= minPrice {
			filtered = append(filtered, product)
		}
	}

	return filtered
}

func filterByMinQuantity(products []Product, minQuantity int) []Product {
	var filtered []Product

	for _, product := range products {
		if product.Quantity >= minQuantity {
			filtered = append(filtered, product)
		}
	}

	return filtered
}

func printDisplayers(items []Displayer) {
	for _, item := range items {
		item.Display()
	}
}

func toDisplayers(products []Product) []Displayer {
	var items []Displayer
	for _, product := range products {
		items = append(items, product)
	}
	return items
}

func main() {
	products := []Product{
		{"Яблоко", "Фрукты", 294.2, 10},
		{"Апельсин", "Фрукты", 305.6, 15},
		{"Картофель", "Овощи", 105.3, 45},
		{"Морковь", "Овощи", 165.5, 30},
	}

	newProducts := Product{"Курица", "Мясо", 405, 20}
	products = addProduct(products, newProducts)

	fmt.Println("Все продукты")
	printDisplayers(toDisplayers(products))

	minPrice := 200.0
	filterPrice := filterByMinPrice(products, minPrice)
	fmt.Printf("Продукты с ценой не ниже %.1f\n", minPrice)
	printDisplayers(toDisplayers(filterPrice))

	minQuantity := 20
	filterQuantity := filterByMinQuantity(products, minQuantity)
	fmt.Printf("Продукты с остатком не менее %d\n", minQuantity)
	printDisplayers(toDisplayers(filterQuantity))
}
