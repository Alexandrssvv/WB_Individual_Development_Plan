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

type Book struct {
	Name   string
	Author string
	Year   int
	Pages  int
	Genre  string
}

func (b Book) Display() {
	fmt.Printf("Название книги: %s, Автор: %s, Год издания: %d, Количество страниц: %d, Жанр: %s\n",
		b.Name, b.Author, b.Year, b.Pages, b.Genre)
}

func addBook(books []Book, book Book) []Book {
	return append(books, book)
}

func filterByYear(books []Book, year int) []Book {
	var filteredBook []Book
	for _, book := range books {
		if book.Year == year {
			filteredBook = append(filteredBook, book)
		}
	}
	return filteredBook
}

func filterByGenre(books []Book, genre string) []Book {
	var filteredBook []Book
	for _, book := range books {
		if book.Genre == genre {
			filteredBook = append(filteredBook, book)
		}
	}
	return filteredBook
}

func printDisplayer(items []Displayer) {
	for _, item := range items {
		item.Display()
	}
}

func convertToDisplayers(books []Book) []Displayer {
	var items []Displayer
	for _, book := range books {
		items = append(items, book)
	}
	return items
}

func main() {
	books := []Book{
		{"Тень над городом", "Алексей Морозов", 2018, 352, "Детектив"},
		{"Предел разума", "Ирина Соколова", 2021, 416, "Научная фантастика"},
		{"Сердце океана", "Мария Лебедева", 2015, 288, "Роман"},
		{"Код хаоса", "Дмитрий Власов", 2020, 384, "Триллер"},
		{"Путь в никуда", "Олег Кравцов", 2017, 320, "Драма"},
		{"Империя пепла", "Анна Белова", 2022, 540, "Фэнтези"},
	}

	newBook := Book{"Тишина внутри", "Екатерина Орлова", 2019, 260, "Психологический роман"}
	books = addBook(books, newBook)

	reader := bufio.NewReader(os.Stdin)

	var year int
	fmt.Print("Введите год для фильтрации: ")
	fmt.Scanln(&year)
	fmt.Print("Введите жанр для фильтрации: ")
	genre, _ := reader.ReadString('\n')
	genre = strings.TrimSpace(genre)

	filteredByYear := filterByYear(books, year)
	filteredByGenre := filterByGenre(books, genre)

	fmt.Println("Все книги:")
	printDisplayer(convertToDisplayers(books))

	fmt.Println("Фильтрация по году:")
	printDisplayer(convertToDisplayers(filteredByYear))

	fmt.Println("Фильтрация по жанру:")
	printDisplayer(convertToDisplayers(filteredByGenre))
}
