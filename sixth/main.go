package main

import (
	"fmt"
	// "strings"
	// "math"
)

// 1. Создайте структуру Person с полями Name и Age. Добавьте метод SayHello, который выводит:
// Привет, меня зовут Name, мне Age лет.

type Person struct {
	name string
	age  int
}

func SayHello(p Person) {
	fmt.Printf("Привет, меня зовут %s, мне %d лет.\n", p.name, p.age)
}

// 2. Определите структуру Product с названием и ценой. Напишите метод ChangePrice(newPrice
// float64), который обновляет цену.

type Product struct {
	name  string
	price float64
}

func ChangePrice(p Product, newPrice float64) {
	p.price = newPrice

	fmt.Println(p.name, p.price)
}

type Student struct {
	name  string
	grade int
}

func StudentActivity(s Student) {
	if s.grade >= 60 {
		fmt.Println(true)
	} else {
		fmt.Println(false)
	}
}

type Account struct {
	name  string
	balance float64
}

func Deposit(a Account, newSum float64) {
	a.balance += newSum

	fmt.Println(a.balance)
}

func Withdraw(a Account, newSum float64) {
	a.balance -= newSum

	fmt.Println(a.balance)
}

func main() {

	person := Person{
		name: "Agata",
		age:  11,
	}
	SayHello(person)

	product := Product{
		name:  "TV",
		price: 50000.0,
	}
	ChangePrice(product, 45000.0)

	// 3. Создайте структуру Student, содержащую имя студента и его оценку (число от 0 до 100).
	// Напишите функцию, которая принимает объект Student и возвращает true, если студент сдал
	// экзамен (оценка 60 и выше).
	student := Student{
		name: "Agata",
		grade:  77,
	}
	StudentActivity(student)
	// 4. Создайте структуру Account, которая содержит имя владельца и текущий баланс. Реализуйте
	// методы Deposit(sum float64) — для пополнения счёта и Withdraw(sum float64) — для снятия со
	// счёта.

	account := Account{
		name:  "Andrey",
		balance: 50000.0,
	}
	Deposit(account, 45000.0)
	Withdraw(account, 15000.0)

	// 5. Создайте структуру Movie с полями Название и Рейтинг. Добавьте метод IsHit(), который
	// возвращает true, если рейтинг фильма 8.0 или выше.
}
