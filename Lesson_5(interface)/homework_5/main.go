package main

import "fmt"

type Displayer interface {
	Display()
}

type Film struct {
	Name         string
	FilmDirector string
	Year         int
	Rating       float64
	Duration     int
}

func (film Film) Display() {
	fmt.Printf("Название фильма: %s, Режисер фильма: %s, Год выпуска: %d, Ретинг фильма: %.1f, Длительность фильма, %d мин\n",
		film.Name, film.FilmDirector, film.Year, film.Rating, film.Duration)
}

func addFilm(films []Film, film Film) []Film {
	return append(films, film)
}

func filterByRating(films []Film, minrating float64) []Film {
	var filteredFilms []Film
	for _, film := range films {
		if film.Rating > minrating {
			filteredFilms = append(filteredFilms, film)
		}
	}
	return filteredFilms
}

func filterByYear(films []Film, year int) []Film {
	var filteredFilms []Film
	for _, film := range films {
		if film.Year == year {
			filteredFilms = append(filteredFilms, film)
		}
	}
	return filteredFilms
}

func printFilms(items []Displayer) {
	for _, item := range items {
		item.Display()
	}
}

func convertorToDisplayers(films []Film) []Displayer {
	var items []Displayer
	for _, film := range films {
		items = append(items, film)
	}
	return items
}

func main() {
	films := []Film{
		{
			Name:         "Inception",
			FilmDirector: "Christopher Nolan",
			Year:         2010,
			Rating:       8.8,
			Duration:     148,
		},
		{
			Name:         "The Shawshank Redemption",
			FilmDirector: "Frank Darabont",
			Year:         1994,
			Rating:       9.3,
			Duration:     142,
		},
		{
			Name:         "The Dark Knight",
			FilmDirector: "Christopher Nolan",
			Year:         2008,
			Rating:       9.0,
			Duration:     152,
		},
		{
			Name:         "Interstellar",
			FilmDirector: "Christopher Nolan",
			Year:         2014,
			Rating:       8.7,
			Duration:     169,
		},
		{
			Name:         "Fight Club",
			FilmDirector: "David Fincher",
			Year:         1999,
			Rating:       8.8,
			Duration:     139,
		},
		{
			Name:         "Forrest Gump",
			FilmDirector: "Robert Zemeckis",
			Year:         1994,
			Rating:       8.8,
			Duration:     142,
		},
	}

	newFilm := Film{
		Name:         "The Matrix",
		FilmDirector: "Lana Wachowski & Lilly Wachowski",
		Year:         1999,
		Rating:       8.7,
		Duration:     136,
	}
	films = addFilm(films, newFilm)

	fmt.Println("Полный список фильмов:")
	printFilms(convertorToDisplayers(films))

	var minrating float64
	var year int

	fmt.Print("Минимальный рейтинг: ")
	fmt.Scanln(&minrating)
	fmt.Print("Введите год выпуска для фильтрации: ")
	fmt.Scanln(&year)

	filteredByRating := filterByRating(films, minrating)
	filteredByYear := filterByYear(films, year)

	fmt.Printf("Фильмы с минимальным рейтингом %.1f\n", minrating)
	printFilms(convertorToDisplayers(filteredByRating))

	fmt.Printf("Фильмы выпущенные в %d году\n", year)
	printFilms(convertorToDisplayers(filteredByYear))

}
