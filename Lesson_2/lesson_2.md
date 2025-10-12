# Теория:
Модуль 2. Основы Go
1. Переменные и константы
   В Go есть три способа объявления переменных:
   var x int = 10         // явное указание типа
   var y = 20             // тип выводится автоматически
   z := 30                // короткая форма (только внутри функций)
   Константы объявляются через const:
   const Pi = 3.14
2. Типы данных
   Целые: int, int8, int16, int32, int64
   Беззнаковые: uint, uint8 (aka byte), uint16, uint32, uint64
   Числа с плавающей точкой: float32, float64
   Строки: string
   Булевы: bool (true, false)
   Пример:
   var age int = 25
   var name string = "Alex"
   var isOk bool = true
3. Коллекции
   **Массивы**
   Фиксированного размера:
   var arr [3]int = [3]int{1, 2, 3}
   fmt.Println(arr[0]) // 1
   **Срезы (slices)**
   Гибкие, динамические массивы:
   nums := []int{1, 2, 3}
   nums = append(nums, 4) // добавляем элемент
   fmt.Println(nums)      // [1 2 3 4]
   **Карты (map)**
   Ассоциативные массивы («словарь»):
   m := make(map[string]int)
   m["apple"] = 5
   m["banana"] = 7
   fmt.Println(m["apple"]) // 5
4. Управляющие конструкции
   **if**
   if age >= 18 {
   fmt.Println("Совершеннолетний")
   } else {
   fmt.Println("Несовершеннолетний")
   }
   **for**
   Единственный цикл в Go:
   for i := 0; i < 5; i++ {
   fmt.Println(i)
   }
   **switch**
   day := 3
   switch day {
   case 1:
   fmt.Println("Понедельник")
   case 2:
   fmt.Println("Вторник")
   default:
   fmt.Println("Другой день")
   }
