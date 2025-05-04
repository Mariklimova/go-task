package main

import (
	"fmt"
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
	name    string
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

// 5. Создайте структуру Movie с полями Название и Рейтинг. Добавьте метод IsHit(), который
// возвращает true, если рейтинг фильма 8.0 или выше.
type Movie struct {
	name   string
	rating float32
}

func IsHit(m Movie) {
	if m.rating >= 8.0 {
		fmt.Println(true)
	} else {
		fmt.Println(false)
	}
}

// 6. Создайте структуру User с полями Имя и Пароль. Добавьте метод ChangePassword(oldPass,
// 	newPass string), который меняет пароль, только если введён правильный старый пароль.

type User struct {
	name     string
	password string
}

func ChangePassword(u User, oldPass, newPass string) {
	if u.password == oldPass {
		u.password = newPass
		fmt.Printf("Пароль для пользователя %s успешно изменен\n", u.name)
	} else {
		fmt.Println("Ошибка: неверный старый пароль")
	}
}

// 7. Создайте структуру Item с полями Название и Цена. Создайте структуру Store, которая
// 	содержит имя магазина и слайс из товаров []Item. Добавьте метод TotalPrice(), который
// 	возвращает суммарную стоимость всех товаров в магазине.

type Item struct {
	name  string
	price float64
}

type Store struct {
	name  string
	items []Item
}

func TotalPrice(s Store) float64 {
	total := 0.0
	for _, item := range s.items {
		total += item.price
	}
	return total
}

// 8. Создайте структуру Client с полями Имя и Баланс. Создайте слайс клиентов и функцию
// 	FindRichestClient(clients []Client), которая находит и возвращает клиента с самым большим
// 	балансом.

type Client struct {
	Name    string
	Balance float64
}

func FindRichestClient(clients []Client) Client {
	if len(clients) == 0 {
		return Client{}
	}

	richest := clients[0]

	for _, client := range clients {
		if client.Balance > richest.Balance {
			richest = client
		}
	}

	return richest
}

// 9. Создайте структуру User с полями Логин и Пароль. Создайте функцию Login(users []User, login,
// 	password string), которая проверяет, есть ли пользователь с таким логином и паролем. Если
// 	// пользователь найден — верните true, иначе — false

type User_9 struct {
	Login    string
	Password string
}

func Login(users []User_9, login, password string) bool {
	for _, user := range users {
		if user.Login == login && user.Password == password {
			return true
		}
	}
	return false
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
		name:  "Agata",
		grade: 77,
	}
	StudentActivity(student)
	// 4. Создайте структуру Account, которая содержит имя владельца и текущий баланс. Реализуйте
	// методы Deposit(sum float64) — для пополнения счёта и Withdraw(sum float64) — для снятия со
	// счёта.

	account := Account{
		name:    "Andrey",
		balance: 50000.0,
	}
	Deposit(account, 45000.0)
	Withdraw(account, 15000.0)

	// 5. Создайте структуру Movie с полями Название и Рейтинг. Добавьте метод IsHit(), который
	// возвращает true, если рейтинг фильма 8.0 или выше.
	movie := Movie{
		name:   "The Vampire diaries",
		rating: 8.8,
	}
	IsHit(movie)

	// 6. Создайте структуру User с полями Имя и Пароль. Добавьте метод ChangePassword(oldPass,
	// 	newPass string), который меняет пароль, только если введён правильный старый пароль.

	user := User{
		name:     "Иван",
		password: "qwerty123",
	}
	ChangePassword(user, "qwerty123", "newSecurePassword")

	// 7. Создайте структуру Item с полями Название и Цена. Создайте структуру Store, которая
	// 	содержит имя магазина и слайс из товаров []Item. Добавьте метод TotalPrice(), который
	// 	возвращает суммарную стоимость всех товаров в магазине.

	store := Store{
		name: "Продовольственный магазин",
		items: []Item{
			{name: "Хлеб", price: 50.5},
			{name: "Молоко", price: 80.0},
			{name: "Яйца", price: 120.75},
		},
	}

	total := TotalPrice(store)
	fmt.Printf("Общая стоимость товаров в магазине '%s': %.2f руб.\n", store.name, total)

	store.items = append(store.items, Item{name: "Сыр", price: 350.25})

	fmt.Printf("Общая стоимость после добавления сыра: %.2f руб.\n", TotalPrice(store))

	// 8. Создайте структуру Client с полями Имя и Баланс. Создайте слайс клиентов и функцию
	// 	FindRichestClient(clients []Client), которая находит и возвращает клиента с самым большим
	// 	балансом.

	clients := []Client{
		{Name: "Алексей", Balance: 1500.50},
		{Name: "Мария", Balance: 3200.75},
		{Name: "Иван", Balance: 800.25},
		{Name: "Ольга", Balance: 5000.00},
	}

	richest := FindRichestClient(clients)
	fmt.Printf("Самый богатый клиент: %s (Баланс: %.2f)\n", richest.Name, richest.Balance)

	// 9. Создайте структуру User с полями Логин и Пароль. Создайте функцию Login(users []User, login,
	// 	password string), которая проверяет, есть ли пользователь с таким логином и паролем. Если
	// 	// пользователь найден — верните true, иначе — false

	users := []User_9{
		{Login: "admin", Password: "qwerty123"},
		{Login: "user1", Password: "password1"},
		{Login: "user2", Password: "myp@ssw0rd"},
	}

	fmt.Println(Login(users, "admin", "qwerty123")) 
	fmt.Println(Login(users, "user1", "wrongpass")) 
	fmt.Println(Login(users, "unknown", "password")) 

	// 10. Найти числа находящиеся сразу в двух слайсах Пример: [1,2,3] и [2,3,4] → [2,3]

	// 11. Найти объединение двух слайсов без повторов Пример: [1,2,3] и [2,3,4] → [1,2,3,4]

}
