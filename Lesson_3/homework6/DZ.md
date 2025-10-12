**Замыкания**
# Задание:
Написать функцию makeMultiplier(factor int) func(int) int, которая возвращает функцию-замыкание, умножающую число на factor.
Пример:

    double := makeMultiplier(2)
    fmt.Println(double(5)) // 10