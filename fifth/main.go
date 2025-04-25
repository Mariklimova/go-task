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

func main() {
	fmt.Println(verific(5))
	fmt.Println(sum_el([]int{1, 2, 3, 4, 5}))
	fmt.Println(max_min_el([]float64{1, 2, 3, 4, 5}))
	fmt.Println(days(5))
	fmt.Println(divide_el([]int{1, 2, 3, 4, 5}))
	fmt.Println(without_ou([]int{1, 0, 2, 3, 0, 4, 5}))
	fmt.Println(incl_el([]int{1, 0, 2, 3, 0, 4, 5}, 7))
}
