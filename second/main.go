package main

import (
	"fmt"
	"math"
	"strings"
)

func isPrime(num int) bool {
	if num <= 1 {
		return false
	}
	if num == 2 {
		return true
	}
	if num%2 == 0 {
		return false
	}

	maxDivisor := int(math.Sqrt(float64(num))) + 1
	for i := 3; i < maxDivisor; i += 2 {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	// 	1. Напишите программу, которая принимает два числа и выводит наибольшее из них.

	var a, b float64
	fmt.Print("Введите два числа: ")
	fmt.Scan(&a, &b)
	fmt.Println("Наибольшее число:", max(a, b))

	// 2. Напишите программу, которая выводит сумму чисел от 1 до N, где N вводится пользователем.

	var N int
	sum := 0
	fmt.Print("Введите число N: ")
	fmt.Scan(&N)

	for i := 1; i <= N; i++ {
		sum += i
	}

	fmt.Printf("Сумма чисел от 1 до %d ровна: %d\n", N, sum)

	// 3. Напишите программу, которая по номеру дня недели (от 1 до 7) выводит название дня
	// (например, для 1 — "Понедельник").
	// Ввод: 2 → Вывод: Вторник
	// Ввод: 5 → Вывод: Пятница

	var dayNumber int

	fmt.Print("Введите номер дня недели (1-7): ")
	fmt.Scan(&dayNumber)

	switch dayNumber {
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
	default:
		fmt.Println("Ошибка: введите число от 1 до 7")
	}

	// 4. Напишите программу, которая принимает число N и выводит все числа от 1 до N, делящиеся
	// на 3
	// Ввод: 10 → Вывод: 3 6 9
	// Ввод: 20 → Вывод: 3 6 9 12 15 18

	var n int = 10

	if n < 2 {
		fmt.Println("Нет простых чисел до", n)
		return
	}

	fmt.Printf("Простые числа до %d: ", n)
	for i := 2; i <= n; i++ {
		if isPrime(i) {
			fmt.Print(i, " ")
		}
	}
	fmt.Println()

	// 	5. Напишите программу, которая проверяет, делится ли введенное число на 5 и 10.
	// Ввод: 25 → Вывод: Делится на 5, но не на 10
	// Ввод: 30 → Вывод: Делится на 5 и на 10

	var num int

	fmt.Print("Введите число: ")
	fmt.Scan(&num)

	switch {
	case num%5 == 0 && num%10 == 0:
		fmt.Printf("%d делится на 5 и на 10\n", num)
	case num%5 == 0:
		fmt.Printf("%d делится на 5, но не на 10\n", num)
	default:
		fmt.Printf("%d не делится ни на 5, ни на 10\n", num)
	}

	// 6. Напишите программу, которая принимает 3 числа и выводит максимальное из них.
	// Ввод: 3, 5, 2 → Вывод: 5
	// Ввод: 9, 7, 8 → Вывод: 9

	num_1 := 3
	num_2 := 5
	num_3 := 2

	fmt.Println(max(num_1, num_2, num_3))

	// 7. Напишите программу, которая вычисляет факториал числа.
	// Ввод: 5 → Вывод: 120
	// Ввод: 7 → Вывод: 5040

	number := 5
	result := 1
	for i := 1; i <= number; i++ {
		result *= i
	}
	fmt.Println(result)

	// 	8. Напишите программу, которая выводит все нечетные числа от 1 до N.
	// Ввод: 10 → Вывод: 1 3 5 7 9
	// Ввод: 6 → Вывод: 1 3 5

	x := 10
	var result_8 []int
	for i := 1; i <= x; i += 2 {
		result_8 = append(result_8, i)
	}
	fmt.Println(result_8)

	// 9. Напишите программу, которая принимает строку и подсчитывает количество гласных в ней.
	// Ввод: Hello World → Вывод: 3
	// Ввод: Go Programming → Вывод: 4

	var str_9 string = "Hello World"
	vowels := "aeiouAEIOU"
	count := 0

	for _, char := range str_9 {
		if strings.ContainsRune(vowels, char) {
			count++
		}
	}
	fmt.Printf("Количество гласных: %d\n", count)

	// 10. Напишите программу, которая проверяет, является ли строка палиндромом.
	// Ввод: madam → Вывод: Да, это палиндром
	// Ввод: hello → Вывод: Нет, это не палиндром

	var str_10 string = "madam"

	str_10 = strings.ToLower(strings.ReplaceAll(str_10, " ", ""))

	for i := 0; i < len(str_10)/2; i++ {
		if str_10[i] != str_10[len(str_10)-1-i] {
			fmt.Println("Нет, это не палиндром")
			break
		}
	}
	fmt.Println("Да, это палиндром")

	// 11. Напишите программу, которая находит сумму чисел от X до Y.
	// Ввод: 3, 5 → Вывод: 12
	// Ввод: 1, 4 → Вывод: 10

	var q, y int
	fmt.Print("Введите первое число (X): ")
	fmt.Scan(&q)

	fmt.Print("Введите второе число (Y): ")
	fmt.Scan(&y)

	if q > y {
		q, y = y, q
	}
	sum_11 := 0
	for i := q; i <= y; i++ {
		sum_11 += i
	}
	fmt.Printf("Сумма чисел от %d до %d: %d\n", q, y, sum_11)

	// 12. Напишите программу, которая выводит все простые числа до числа N.
	// Ввод: 10 → Вывод: 2 3 5 7
	// Ввод: 20 → Вывод: 2 3 5 7 11 13 17 19

	num_12 := 10

	// Проверка, является ли num_12 простым числом
	if num_12 <= 1 {
		fmt.Println(num_12, "не простое число")
	} else if num_12 == 2 {
		fmt.Println(num_12, "простое число")
	} else if num_12%2 == 0 {
		fmt.Println(num_12, "не простое число (делится на 2)")
	} else {
		isPrime := true
		maxDivisor := int(math.Sqrt(float64(num_12))) + 1
		for i := 3; i < maxDivisor; i += 2 {
			if num_12%i == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			fmt.Println(num_12, "простое число")
		} else {
			fmt.Println(num_12, "не простое число")
		}
	}

	// Вывод всех простых чисел до num_12
	if num_12 < 2 {
		fmt.Println("Нет простых чисел до", num_12)
	} else {
		fmt.Printf("Простые числа до %d: ", num_12)
		for i := 2; i <= num_12; i++ {
			isPrime := true
			if i <= 1 {
				isPrime = false
			} else if i == 2 {
				isPrime = true
			} else if i%2 == 0 {
				isPrime = false
			} else {
				maxDiv := int(math.Sqrt(float64(i))) + 1
				for j := 3; j < maxDiv; j += 2 {
					if i%j == 0 {
						isPrime = false
						break
					}
				}
			}
			if isPrime {
				fmt.Print(i, " ")
			}
		}
		fmt.Println()
	}
	
	// 13. Напишите программу, которая находит индекс первого вхождения символа в строку.
	// Ввод: hello, e → Вывод: 1
	// Ввод: world, o → Вывод: 1

	str_13 := "hello"
	const char rune = 'e'
	index := strings.IndexRune(str_13, char)
	fmt.Println(index)
	
	// 	14. Напишите программу, которая выводит все числа от 1 до N с шагом 2.
	// Ввод: 10 → Вывод: 1 3 5 7 9
	// Ввод: 6 → Вывод: 1 3 5

	// 15. Напишите программу, которая проверяет, является ли число отрицательным.
	// Ввод: -5 → Вывод: Отрицательное число
	// Ввод: 3 → Вывод: Не отрицательное число

	// 16. Напишите программу, которая принимает 2 строки и выводит их длину.
	// Ввод: Hello, World → Вывод: 5 5
	// Ввод: Go, Programming → Вывод: 2 11

}
