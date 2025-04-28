package main

import (
	"fmt"
	"math"
)

//1. Напишите функцию, которая принимает целое число и возвращает строку "Четное" или "Нечетное", в зависимости от того, является ли число чётным или нечётным.

func verific(a int) string {
	if a%2 == 0 {
		return "Четное"
	} else {
		return "Нечетное"
	}
}

//2. Напишите функцию, которая принимает слайс целых чисел и возвращает сумму всех его элементов.

func sum_el(sl []int) int {
	sum := 0
	for _, item := range sl {
		sum += item
	}
	return sum
}

//3. Напишите функцию, которая принимает слайс целых чисел и возвращает максимальное и минимальное число в этом слайсе.

func max_min_el(sl []float64) (float64, float64) {
	max := math.Inf(-1)
	min := math.Inf(1)
	for _, item := range sl {
		if item > max {
			max = item
		}
		if item < min {
			min = item
		}
	}
	return max, min
}

//4. Напишите функцию, которая принимает число от 1 до 7 (представляющее день недели) и возвращает строку, например "Понедельник", "Вторник" и т. д. Используйте оператор switch.

func days(day int) string {
	switch day {
	case 1:
		return "Понедельник"

	case 2:
		return "Вторник"

	case 3:
		return "Среда"

	case 4:
		return "Четверг"

	case 5:
		return "Пятница"

	case 6:
		return "Суббота"

	default:
		return "Воскресенье"

	}

}

//5. Напишите функцию, которая принимает слайс целых чисел и разделяет его на два слайса: один с чётными числами, другой с нечётными.

func divide_el(sl []int) ([]int, []int) {
	var odd []int
	var even []int

	for _, item := range sl {
		if item%2 == 0 {
			odd = append(odd, item)

		} else {
			even = append(even, item)
		}
	}
	return odd, even
}

//6. Напишите функцию, которая принимает слайс целых чисел и удаляет все нули из него.

func without_ou(sl []int) []int {
	var new_slice []int

	for _, item := range sl {
		if item != 0 {
			new_slice = append(new_slice, item)

		}
	}
	return new_slice
}

//7. Напишите функцию, которая принимает слайс целых чисел и число, и проверяет, содержится ли это число в слайсе.

func incl_el(sl []int, num int) bool {
	var flag bool
	for _, item := range sl {
		if item == num {
			flag = true
		} else {
			flag = false
		}
	}
	return flag
}

// 8. Напишите функцию, которая принимает два числа и операцию (например, +, -, *, /) и выполняет
// соответствующую арифметическую операцию.
func calculate(a float64, b float64, operation string) (float64, error) {
	switch operation {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		if b == 0 {
			return 0, fmt.Errorf("деление на ноль")
		}
		return a / b, nil
	default:
		return 0, fmt.Errorf("неизвестная операция: %s", operation)
	}
}

// 9. Напишите функцию, которая принимает слайс целых чисел и возвращает среднее
// арифметическое этих чисел.

func average_value(sl []int) int {
	var sum int

	for _, item := range sl {
		sum += item
	}
	return sum / len(sl)
}

// 10. Создайте функцию, которая принимает массив чисел, находит все повторяющиеся элементы и
// удаляет их. После этого функция должна проверить, не является ли массив пустым, и вывести
// соответствующее сообщение.

func removeDuplicates(nums []int) []int {

	seen := make(map[int]bool)
	result := []int{}

	for _, num := range nums {
		if !seen[num] {
			seen[num] = true
			result = append(result, num)
		}
	}
	return result
}
func processArray(numbers []int) []int {
	uniqueNumbers := removeDuplicates(numbers)
	if len(uniqueNumbers) == 0 {
		fmt.Println("Массив пуст после удаления дубликатов")
	}
	return uniqueNumbers
}

// 11. Разработайте программу, которая находит сумму всех нечётных чисел в слайсе и выводит их
// индексы.

func sumOddNumbersWithIndices(slice []int) (int, []int) {
	sum := 0
	var indices []int

	for index, value := range slice {
		if value%2 != 0 {
			sum += value
			indices = append(indices, index)
		}
	}

	return sum, indices
}

// 12. Напишите программу, которая проверяет, является ли слайс чисел палиндромом, то есть
// читается ли слайс одинаково в обоих направлениях. В случае палиндрома программа должна
// вывести true, иначе false.

func isPalindrome(slice []int) bool {
	for i := 0; i < len(slice)/2; i++ {
		if slice[i] != slice[len(slice)-1-i] {
			return false
		}
	}
	return true
}

// 13. Напишите программу, которая находит сумму всех чисел в слайсе, которые больше среднего
// значения



// 14. Напишите программу, которая генерирует два случайных слайса чисел от 1 до 100 и находит
// пересечение этих слайсов (элементы, которые встречаются в обоих слайсах).



// 15. Напишите программу, которая генерирует слайс из N случайных чисел от 1 до 100, затем
// создает два слайса: один с чисел, делящихся на 5, а другой — на 7.




func main() {
	fmt.Println(verific(5))
	fmt.Println(sum_el([]int{1, 2, 3, 4, 5}))
	fmt.Println(max_min_el([]float64{1, 2, 3, 4, 5}))
	fmt.Println(days(5))
	fmt.Println(divide_el([]int{1, 2, 3, 4, 5}))
	fmt.Println(without_ou([]int{1, 0, 2, 3, 0, 4, 5}))
	fmt.Println(incl_el([]int{1, 0, 2, 3, 0, 4, 5}, 7))
	fmt.Println(calculate(10, 5, "+"))
	fmt.Println(average_value([]int{1, 2, 3, 4, 5}))
	fmt.Println("Итоговый результат:", processArray([]int{1, 2, 2, 3, 4, 4, 5}))
	sum, oddIndices := sumOddNumbersWithIndices([]int{2, 5, 8, 3, 6, 1, 7, 4})
	fmt.Printf("Сумма нечётных чисел: %d\nИндексы нечётных чисел: %v\n", sum, oddIndices)
	fmt.Println(isPalindrome([]int{1, 2, 3, 2, 1}))

}
