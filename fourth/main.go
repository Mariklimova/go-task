package main

import (
	"fmt"
)

func main() {
	// 	1. Создать массив из 10 чисел и вывести их сумму.

	numbers := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	sum := 0

	for _, num := range numbers {
		sum += num
	}
	fmt.Println(sum)

	// 2. Найти и вывести наибольшее число в []int{3, 7, 2, 9, 4}.

	nums_2 := []int{3, 7, 2, 9, 4}

	max := nums_2[0]
	for _, num_2 := range nums_2 {
		if num_2 > max {
			max = num_2
		}
	}
	fmt.Println("Наиболшее число: ", max)

	// 3. Вывести минимальный элемент из массива из 6 чисел.

	nums_3 := [5]int{3, 7, 2, 9, 4}

	min := nums_2[0]
	for _, num_3 := range nums_3 {
		if num_3 < min {
			min = num_3
		}
	}
	fmt.Println("Наименьшее число: ", min)

	// 4. Подсчитать количество чётных чисел в слайсе Пример: []int{1, 2, 3, 4, 5, 6} → 3 четных.

	nums_4 := []int{1, 2, 3, 4, 5, 6}
	count := 0
	for _, item := range nums_4 {
		if item%2 == 0 {
			count++
		}
	}
	fmt.Println(count)

	// 5. Разделить слайс на два: чётные и нечётные элементы Пример: []int{1,2,3,4} → [2,4], [1,3]
	nums_5 := []int{1, 2, 3, 4}
	var even []int
	var odd []int

	for _, item := range nums_5 {
		if item%2 == 0 {
			even = append(even, item)
		} else {
			odd = append(odd, item)
		}
	}
	fmt.Println(even, odd)

	// 6. Вывести массив в обратном порядке.

	arr := [6]int{1, 2, 3, 4, 5, 6}
	var newArr []int

	for i := len(arr) - 1; i >= 0; i-- {
		newArr = append(newArr, i)
	}
	fmt.Println(newArr)

	// 7. Подсчитать количество положительных и отрицательных чисел Пример: []int{-1, 2, -3, 4} →
	// положительные: 2, отрицательные: 2
	nums_7 := []int{-1, 2, -3, 4, -6}
	count_7 := 0
	count_77 := 0
	for _, item := range nums_7 {
		if item > 0 {
			count_7++

		} else {
			count_77++
		}
	}
	fmt.Println("положительные: ", count_7, "отрицательные: ", count_77)

	// 8. Удалить все нули из слайса Пример: []int{0, 1, 0, 2, 3} → [1, 2, 3]
	arr_8 := []int{0, 1, 0, 2, 3}
	var newArr_8  []int
	for _, item := range arr_8 {
		if item != 0 {
			newArr_8=append(newArr_8, item)

		}
	}
	fmt.Println(newArr_8)

	// 9. Найти среднее арифметическое всех чисел Реализовать подсчет среднего значения слайса
	// чисел.

	
	// 10. Умножить каждый элемент на 2 Пример: []int{1, 2, 3} → [2, 4, 6]

}
