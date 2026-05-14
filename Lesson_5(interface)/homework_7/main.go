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

type Order struct {
	Number   int
	Name     string
	Sum      float64
	Status   string
	Quantity int
}

func (o Order) Display() {
	fmt.Printf("Номер заказа: %d, Имя клиента: %s, Сумма заказа: %.2f, Статус: %s, Количество товаров в заказе: %d\n ",
		o.Number, o.Name, o.Sum, o.Status, o.Quantity)
}

func addOrder(orders []Order, order Order) []Order {
	return append(orders, order)
}

func FilterBySum(orders []Order, minSum float64) []Order {
	var filteredOrders []Order
	for _, order := range orders {
		if order.Sum >= minSum {
			filteredOrders = append(filteredOrders, order)
		}
	}
	return filteredOrders
}

func FilterByStatus(orders []Order, status string) []Order {
	var filteredOrders []Order
	for _, order := range orders {
		if order.Status == status {
			filteredOrders = append(filteredOrders, order)
		}
	}
	return filteredOrders
}

func printDisplayer(orders []Displayer) {
	for _, order := range orders {
		order.Display()
	}
}

func convertToDisplayer(orders []Order) []Displayer {
	var displayer []Displayer
	for _, order := range orders {
		displayer = append(displayer, order)
	}
	return displayer
}

func main() {
	orders := []Order{
		{Number: 1, Name: "Иван", Sum: 89999.99, Status: "Новый", Quantity: 1},
		{Number: 2, Name: "Александр", Sum: 59999.50, Status: "В обработке", Quantity: 2},
		{Number: 3, Name: "Василий", Sum: 12999.00, Status: "Доставлен", Quantity: 3},
		{Number: 4, Name: "Дина", Sum: 24999.99, Status: "Отменён", Quantity: 1},
		{Number: 5, Name: "Оля", Sum: 4999.90, Status: "Доставлен", Quantity: 2},
		{Number: 6, Name: "Мышь-Матвей", Sum: 2999.50, Status: "В обработке", Quantity: 4},
		{Number: 7, Name: "Андрей", Sum: 17999.00, Status: "Новый", Quantity: 1},
	}

	newOrder := Order{Number: 8, Name: "Дмитрий", Sum: 6999.99, Status: "Доставлен", Quantity: 2}
	orders = addOrder(orders, newOrder)

	fmt.Println("Полный список заказов: ")
	printDisplayer(convertToDisplayer(orders))

	var minSum float64
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Введите минимальную сумму заказа, для фильтрации: ")
	fmt.Scanln(&minSum)

	fmt.Print("Введите статус заказа, для фильтрации: ")
	status, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Ошибка чтения статуса", err)
		return
	}
	status = strings.TrimSpace(status)

	filteredSums := FilterBySum(orders, minSum)
	filteredStats := FilterByStatus(orders, status)

	fmt.Printf("Заказы на сумму выше %.2f: ", minSum)
	printDisplayer(convertToDisplayer(filteredSums))

	fmt.Printf("Заказы со статусом %s: ", status)
	printDisplayer(convertToDisplayer(filteredStats))
}
