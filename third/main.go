package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	// 	1. Вывести каждый символ строки “Привет” 2 способами:, используя for и for range.
	str := "Привет"
	runes := []rune(str)

	for i := 0; i < len(runes); i++ {
		fmt.Printf("%c", runes[i])
	}
	fmt.Println()

	for _, r := range str {
		fmt.Printf("%c", r)
	}
	fmt.Println()

	// 2. Вывести строку GoLang в обратном порядке

	str2 := "GoLang"
	runes2 := []rune(str2)

	for i, j := 0, len(runes2)-1; i < j; i, j = i+1, j-1 {
		runes2[i], runes2[j] = runes2[j], runes2[i]
	}
	fmt.Println(string(runes2))

	// 3. Проверить, содержит ли строка hello world подстроку world.
	str3 := "hello world"
	substr3 := "world"

	if strings.Contains(str3, substr3) {
		fmt.Printf("Строка \"%s\" содержит подстроку \"%s\"\n", str3, substr3)
	} else {
		fmt.Printf("Строка \"%s\" не содержит подстроку \"%s\"\n", str3, substr3)
	}

	// 4. Подсчитать, сколько раз подстрока go встречается в gogogopher.
	str_4 := "gogogopher"
	substr_4 := "go"

	count := strings.Count(str_4, substr_4)
	fmt.Printf("Подстрока \"%s\" встречается в строке \"%s\" %d раз(а)\n", substr_4, str_4, count)

	// 5. Заменить все cat на dog в строке cat-cat-dog.

	str5 := "cat-cat-dog"

	result5 := strings.ReplaceAll(str5, "cat", "dog")
	fmt.Printf("Результат замены: %s\n", result5)

	// 6. Заменить первое вхождение go на GO в строке go go go.

	str6 := "go go go"

	result6 := strings.Replace(str6, "go", "GO", 1)
	fmt.Printf("Результат замены: %q\n", result6)

	// 7. Сделать все буквы строки hello заглавными.

	str7 := "hello"
	fmt.Printf(strings.ToUpper(str7))
	fmt.Println()

	// 8. Убрать пробелы из строки “ hello world “.

	str_8 := " hello world "
	fmt.Println(strings.ReplaceAll(str_8, " ", ""))

	// 9. Разбить строку a,b,c,d на срез строк.

	str_9 := "a,b,c,d"
	fmt.Println(strings.Split(str_9, ","))

	// 10. Объединить []string{"go", "lang"} в строку с пробелом между ними.

	words := []string{"go", "lang"}
	fmt.Println(strings.Join(words, " "))

	// 11. Вывести ha 5 раз подряд.

	fmt.Println(strings.Repeat("ha", 5))

	// 12. Проверить, начинается ли строка golang с go.
	str_12 := "golang"
	prefix := "go"

	if strings.HasPrefix(strings.ToLower(str_12), strings.ToLower(prefix)) {
		fmt.Printf("Строка '%s' начинается с '%s'\n", str_12, prefix)
	} else {
		fmt.Printf("Строка '%s' НЕ начинается с '%s'\n", str_12, prefix)
	}

	// 13. Проверить, заканчивается ли строка index.html на .html.

	str_13 := "index.html"
	end_str := ".html"

	if strings.HasSuffix(strings.ToLower(str_13), strings.ToLower(end_str)) {
		fmt.Printf("Строка '%s' заканчивается на '%s'\n", str_13, end_str)
	} else {
		fmt.Printf("Строка '%s' НЕ заканчивается на '%s'\n", str_13, end_str)
	}

	// 14. Из строки h e l l o удалить все пробелы.

	fmt.Println(strings.ReplaceAll("h e l l o", " ", ""))

	// 15. Для строки GoLang напечатать каждый символ и его код, например: G - 71, o - 111, L - 76 и т.д.

	str_15 := "GoLang"

	for _, char := range str_15 {
		fmt.Printf("%c - %d, ", char, char)
	}
	fmt.Println()

	// 16. Подсчитать количество слов в строке go is awesome.

	words_16 := strings.Fields("go is awesome")
	fmt.Println("Количество слов:", len(words_16))

	// 17. Подсчитать количество заглавных букв в строке GoLang.

	str_17 := "GoLang"
	count_17 := 0

	for _, char := range str_17 {
		if unicode.IsUpper(char) {
			count_17++
		}
	}
	fmt.Printf("Количество заглавных букв: %d\n", count)

	// 18. Если строка заканчивается на ., удалить её.

	str_18 := "Текст с точкой в конце."
	if strings.HasSuffix(str_18, ".") {
		str_18 = str_18[:len(str_18)-1]
	}
	fmt.Println(str_18)

	// 19. В предложении I love apples, заменить apples на oranges.

	fmt.Println(strings.ReplaceAll("I love apples", "apples", "oranges"))

	// 20. Учитывать только латинские гласные a, e, i, o, u.
	str_20 := "Hello, World! Go is awesome!"
	vowels := "aeiouAEIOU"
	count_20 := 0

	for _, char := range str_20 {
		if strings.ContainsRune(vowels, char) {
			count_20++
		}
	}

	fmt.Printf("Количество латинских гласных: %d\n", count_20)

	// 21. Инверсия регистра (если буква — заглавная, сделать строчной и наоборот) GoLang → gOlANG

	result_21 := []rune("GoLang")
	for i, char := range result_21 {
		if unicode.IsUpper(char) {
			result_21[i] = unicode.ToLower(char)
		} else if unicode.IsLower(char) {
			result_21[i] = unicode.ToUpper(char)
		}
	}
	fmt.Println(string(result_21))

	// 22. Проверка, является ли строка палиндромом. казак, шалаш → true



	// 23. Поиск самого длинного слова в строке go is an expressive and concise language → expressive



	// 24. Найти все уникальные символы в строке aabccdee → a b c d e
	
	
	
	// 	25. Подсчитать количество цифр в строке Пример: в abc123def456 — 6 цифр.



	// 	26. Сделать первую букву каждого слова заглавной Пример: go is fun → Go Is Fun.



	// 	27. Подсчитать количество символов пунктуации Для строки Hello, world! How are you? должно
	// 	быть 3 (,, !, ?).



	// 	28. Найти подстроки, начинающиеся с заглавной буквы Пример: Welcome To Go Language →
	// 	Welcome, To, Go, Language



	// 	29. Проверить, является ли строка числом Примеры: 123 → true, 12a3 → false
	// 	Метод: strconv.Atoi



	// 	30. Удалить из строки все согласные Пример: banana → aaa
	//
}
