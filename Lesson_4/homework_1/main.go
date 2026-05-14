package main

import "fmt"

type Book struct {
	Title  string
	Author string
	Year   int
}

func main() {
	books := []Book{
		{
			Title:  "Мастер и Маргарита",
			Author: "Михаил Булгаков",
			Year:   1967,
		},
		{
			Title:  "1984",
			Author: "Джордж Оруэлл",
			Year:   1949,
		},
		{
			Title:  "Преступление и наказание",
			Author: "Фёдор Достоевский",
			Year:   1866,
		},
		{
			Title:  "Три товарища",
			Author: "Эрих Мария Ремарк",
			Year:   1936,
		},
		{
			Title:  "Властелин колец",
			Author: "Джон Рональд Руэл Толкин",
			Year:   1954,
		},
	}

	newBook := Book{
		Title:  "Портрет Дориана Грея",
		Author: "Оскар Уайльд",
		Year:   1890,
	}

	books = append(books, newBook)

	fmt.Println("Полный список книг: ")
	for _, book := range books {
		fmt.Printf("Название книги: %s, Автор: %s, Год издания: %d\n", book.Title, book.Author, book.Year)
	}
}
