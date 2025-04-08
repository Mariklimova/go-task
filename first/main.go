package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	// 1. Проверьте возраст у входных данных :
	// Если < 18 — "Вы несовершеннолетний".
	// Если 18-60 — "Вы взрослый".
	// Если > 60 — "Вы пожилой человек".

	const age int = 11

	if age < 18 {
		fmt.Println("Вы несовершеннолетний")
	} else if age >= 18 && age <= 60 {
		fmt.Println("Вы взрослый")
	} else {
		fmt.Println("Вы пожилой человек")
	}

	// 	2. Оцените успеваемость входных данных:
	// 90-100 — "Отлично".
	// 75-89 — "Хорошо".
	// 50-74 — "Удовлетворительно".
	// < 50 — "Неудовлетворительно".

	const progress int = 95

	if progress <= 100 && progress >= 90 {
		fmt.Println("Отлично")
	} else if progress >= 75 && progress < 90 {
		fmt.Println("Хорошо")
	} else if progress >= 50 && progress < 75 {
		fmt.Println("Удовлетворительно")
	} else {
		fmt.Println("Неудовлетворительно")
	}

	// 3. Выведите день недели по номеру (1-7) с использованием switch.

	const day int = 5
	switch day {
	case 1:
		fmt.Println("Понедельник")
	case 2:
		fmt.Println("Вторник")
	case 3:
		fmt.Println("Среда")
	case 4:
		fmt.Println("Четверг")
	case 5:
		fmt.Println("Пятница")
	case 6:
		fmt.Println("Суббота")
	case 7:
		fmt.Println("Воскресенье")
	}

	// 4. Проверьте, является ли число четным или нечетным

	const num int = 44

	if num%2 == 0 {
		fmt.Println("четное")
	} else {
		fmt.Println("нечетное")
	}

	// 	5. Проверьте, имеет ли человек право на льготу (меньше 18 или старше 65 лет).
	// Вход: 15 → "Льгота есть"
	// Вход: 40 → "Льготы нет"
	// Вход: 70 → "Льгота есть«

	const agePerson int = 50

	if agePerson < 18 && agePerson > 65 {
		fmt.Println("Льгота есть")
	} else {
		fmt.Println("Льготы нет")
	}

	// 	6. Объявите переменные разных типов (int, float64, string, bool) и присвойте значения.
	// Выведите их.
	const (
		a int     = 78
		b float64 = 0.78
		c string  = "srdtgyuhi"
		d bool    = true
	)
	fmt.Println(a, b, c, d)

	// 	7. Преобразуйте целое число в строку и строку в число.
	// Вход: 123 → "123" (преобразовать в строку)
	// Вход: "456" → 456 (преобразовать в число)

	const number int = 123
	fmt.Println(strconv.Itoa(number))

	const str string = "456"
	intNum, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println("Ошибка преобразования:", err)
	} else {
		fmt.Println(intNum)
	}

	// 	8. Введите число с плавающей точкой и преобразуйте его в целое число (округлить).
	// Ввод: 3.14 → 3
	// Ввод: 7.99 → 8

	const floatNum1 float64 = 3.14
	const floatNum2 float64 = 7.99
	fmt.Println(math.Round(floatNum1))
	fmt.Println(math.Round(floatNum2))

	// 9. Выполните операции с входными числами: Сложение Вычитание Умножение Деление
	// Остаток от деления
	const (
		num_1 int = 15
		num_2 int = 5
	)
	fmt.Printf("Сложение: %d, Вычитание: %d, Умножение: %d, Деление: %d, Остаток от деления: %d\n",
		num_1+num_2,
		num_1-num_2,
		num_1*num_2,
		num_1/num_2,
		num_1%num_2,
	)

	// 10. Вычислите периметр прямоугольника по длине и ширине.
	// Вход: длина = 5, ширина = 3 → Периметр = 16

	const (
		length int = 5
		width  int = 3
	)
	fmt.Printf("Периметр: %d\n", 2*(length+width))

	// 11. Вычислите среднее арифметическое двух чисел.
	// Вход: 10 и 20 → Среднее = 15

	const (
		number_1 int = 10
		number_2 int = 20
	)
	fmt.Printf("среднее арифметическое: %d\n", (number_1+number_2)/2)

	// 	12. Сравнение двух входных чисел
	// Проверьте: Если a > b — "a больше" Если a == b — "a равно b" Если a < b — "b больше«

	var x, y int

	fmt.Print("Введите первое число: ")
	fmt.Scan(&x)

	fmt.Print("Введите второе число: ")
	fmt.Scan(&y)

	switch {
	case x > y:
		fmt.Printf("%d больше %d\n", x, y)
	case x == y:
		fmt.Printf("%d равно %d\n", x, y)
	default:
		fmt.Printf("%d больше %d\n", y, x)
	}

	// 9. Найдите максимальное из двух входных чисел.
	var n, m int

	fmt.Print("Введите первое число: ")
	fmt.Scan(&n)

	fmt.Print("Введите второе число: ")
	fmt.Scan(&m)

	max := n
	if m > n {
		max = m
	}
	fmt.Printf("Максимальное число: %d\n", max)

	// 10. Найдите минимальное из трех входных чисел.
	var s, p int

	fmt.Print("Введите первое число: ")
	fmt.Scan(&s)

	fmt.Print("Введите второе число: ")
	fmt.Scan(&p)

	min := s
	if p < s {
		max = p
	}
	fmt.Printf("Максимальное число: %d\n", min)

	// 11. Проверьте, равны ли два входных числа.

	var e, f int

	fmt.Print("Введите первое число: ")
	fmt.Scan(&e)

	fmt.Print("Введите второе число: ")
	fmt.Scan(&f)

	if e == f {
		fmt.Printf("%d равно %d\n", e, f)
	} else {
		fmt.Printf("%d неравно %d\n", e, f)
	}
}
