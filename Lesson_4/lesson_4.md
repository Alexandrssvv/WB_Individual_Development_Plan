# Модуль 4. Структуры и методы

1. Определение struct
struct — это составной тип данных, который позволяет объединять несколько полей разных типов в одну сущность.
Пример:

        package main
            
        import "fmt"
        
        type User struct {
        ID    int
        Name  string
        Email string
        }
        
        func main() {
        u := User{ID: 1, Name: "Alex", Email: "alex@example.com"}
        fmt.Println(u)
        }

   2. Методы структур
   Методы — это функции, "привязанные" к типу (например, к структуре).
   Метод отличается от функции наличием получателя (receiver).

           package main
        
           import "fmt"
        
           type User struct {
           Name string
           Age  int
          }

// Метод для структуры User

        func (u User) SayHello() {
        fmt.Printf("Привет, меня зовут %s, мне %d лет\n", u.Name, u.Age)
        }

// Метод с указателем (может изменять данные)

        func (u *User) Birthday() {
        u.Age++
        }
        
        func main() {
        u := User{Name: "Alex", Age: 25}
        u.SayHello()
        u.Birthday()
        u.SayHello()
        }

3. Встраивание структур (аналог наследования)
В Go нет наследования как в ООП-языках, но есть встраивание — одна структура может содержать другую как "базовую".

        type Person struct {
        Name string
        Age  int
        }
        
        type Employee struct {
        Person     // встраивание
        Position   string
        Department string
        }
        
        func main() {
        e := Employee{
        Person:   Person{Name: "Alex", Age: 30},
        Position: "Developer",
        }
        fmt.Println(e.Name, e.Age, e.Position) // доступ напрямую к полям Person
        }

4. Композиция вместо наследования
Композиция означает использование структур как полей других структур, но без автоматического "подъёма" полей, как при встраивании.
Это даёт более гибкий контроль.

       type Address struct {
       City  string
       Street string
       }
    
       type Company struct {
       Name    string
       Address Address // композиция
       }

5. Конструкторы (фабричные функции)
В Go нет ключевого слова constructor, вместо этого обычно создают фабричную функцию NewXxx, которая возвращает указатель на структуру.

       type User struct {
       Name string
       Age  int
       }

// Фабричная функция

    func NewUser(name string, age int) *User {
    return &User{Name: name, Age: age}
    }
    
    func main() {
    u := NewUser("Alex", 25)
    fmt.Println(u)
    }

6. Пакеты (package, import)
Go-код организуется в пакеты.
package main — точка входа программы.
import — подключение внешних и стандартных пакетов.
Пример:

        project/
        main.go
        user/
        user.go
        user/user.go
        package user
        
        type User struct {
        Name string
        Age  int
        }
        main.go
        package main
        
        import (
        "fmt"
        "project/user"
        )
        
        func main() {
        u := user.User{Name: "Alex", Age: 25}
        fmt.Println(u)
        }
