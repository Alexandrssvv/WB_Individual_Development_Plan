# Теория:
**Модуль 3. Функции**
1. Определение функций
 
Функции в Go объявляются через ключевое слово func:

    func add(a int, b int) int {
    return a + b
    }

Вызов:

    result := add(3, 5) // result = 8

2. Множественные возвращаемые значения

Функция может возвращать сразу несколько значений:

    func divide(a, b int) (int, error) {
    if b == 0 {
    return 0, fmt.Errorf("деление на ноль")
    }
    return a / b, nil
    }

Вызов:

    res, err := divide(10, 2)
    if err != nil {
    fmt.Println("Ошибка:", err)
    } else {
    fmt.Println("Результат:", res)
    }

3. Именованные возвращаемые значения
Можно заранее объявить имена возвращаемых переменных:
    
    func rectangle(width, height int) (area int, perimeter int) {
    area = width * height
    perimeter = 2 * (width + height)
    return // возвращает area, perimeter автоматически
    }

4. Передача по значению и по указателю

По умолчанию Go передаёт копию значения.
Чтобы изменить переменную в функции — используем указатель (* и &):
    
    func increment(x int) {
    x = x + 1
    }

    func incrementPtr(x *int) {
    *x = *x + 1
    }

    func main() {
    a := 10
    increment(a)
    fmt.Println(a) // 10 (без изменений)

    incrementPtr(&a)
    fmt.Println(a) // 11 (изменено)
    }
5. Замыкания
Функция может возвращать функцию и хранить её окружение:


    func counter() func() int {
    count := 0
    return func() int {
    count++
    return count
    }
    }
    
    func main() {
    c := counter()
    fmt.Println(c()) // 1
    fmt.Println(c()) // 2
    fmt.Println(c()) // 3
    }

6. Рекурсия
Функция может вызывать саму себя:


    func factorial(n int) int {
    if n == 0 {
    return 1
    }
    return n * factorial(n-1)
    }
    
    func main() {
    fmt.Println(factorial(5)) // 120
    }
7. Встроенные функции
В Go есть ряд встроенных функций:
len — длина среза, массива, строки, карты, канала
append — добавление элементов в срез
copy — копирование элементов из одного среза в другой
make — создание среза, карты или канала
new — выделяет память под переменную и возвращает указатель
Примеры:


    func main() {
    nums := []int{1, 2, 3}
    nums = append(nums, 4) // [1 2 3 4]

    copyNums := make([]int, len(nums))
    copy(copyNums, nums) // копирование

    fmt.Println(len(nums))     // 4
    fmt.Println(copyNums)      // [1 2 3 4]

    p := new(int)  // *int
    *p = 42
    fmt.Println(*p) // 42
    }
