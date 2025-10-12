package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println("Hello world")
	nextWeek := now.Add(7 * 24 * time.Hour)
	newYear := time.Date(now.Year()+1, time.January, 1, 0, 0, 0, 0, now.Location())
	dayForNewYear := int(newYear.Sub(now).Hours() / 24)
	fmt.Println("Текущая дата:", now.Format("02.01.2006"))
	fmt.Println("Текущее время:", now.Format("15:04"))
	fmt.Println("Через неделю будет:", nextWeek.Format("02.01.2006"))
	fmt.Println("До нового года осталось:", dayForNewYear, "дн.")
	elapsed := time.Since(now)
	fmt.Println("Время выполнения:", elapsed)
}
