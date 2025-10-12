# Теория:
## Модуль 1. Введение
1. Установка и настройка Go
   brew install go
   Проверка версии:
   go version
   Папка для проектов: обычно ~/go или любая удобная.
   Система модулей (с 1.11+): используется go mod init для управления зависимостями.
2. Первая программа
   **Создание файла main.go:**

       package main

       import "fmt"

       func main() {
       fmt.Println("Привет, Go!")
       }
   **Запуск:**
       go run main.go

3. Структура проекта
   Минимальная:
   myproject/
   └── main.go
   С модулями:
   myproject/
   ├── go.mod
   └── main.go
4. Компиляция
   Запуск без сборки:
   go run main.go
   Сборка бинаря:
   go build -o app main.go
   ./app