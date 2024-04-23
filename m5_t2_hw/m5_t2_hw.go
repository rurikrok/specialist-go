/*
Задача №1
Написать функцию, которая расшифрует строку.
code = "220411112603141304"
Каждые две цифры - это либо буква латинского алфавита в нижнем регистре либо пробел.
Отчет с 00 -> 'a' и до 25 -> 'z', 26 -> ' '(пробел).
Вход: строка из цифр. Выход: Текст.
Проверка работы функции выполняется через вторую строку.

codeToString(code) -> "???????'
*/
package main

import (
	"fmt"
	"strconv"
)

func codeToString(code string) string {
	result := ""

	for i := 0; i < len(code); i += 2 {
		num, _ := strconv.Atoi(code[i : i+2])
		if num == 26 {
			result += " "
		} else {
			result += string('a' + num)
		}
	}

	return result
}

// var decoder = map[string]string{
// 	"00": "a", "01": "b", "02": "c", "03": "d", "04": "e", "05": "f",
// 	"06": "g", "07": "h", "08": "i", "09": "j", "10": "k", "11": "l",
// 	"12": "m", "13": "n", "14": "o", "15": "p", "16": "q", "17": "r",
// 	"18": "s", "19": "t", "20": "u", "21": "v", "22": "w", "23": "x",
// 	"24": "y", "25": "z", "26": " ",
// }

var decoder = make(map[string]string)

func init() {
	for i := 0; i <= 26; i++ {
		key := fmt.Sprintf("%02d", i)
		if i == 26 {
			decoder[key] = " "
		} else {
			decoder[key] = string('a' + i)
		}
	}
}

func codeToStringMap(code string) string {
	result := ""

	for i := 0; i < len(code); i += 2 {
		result += decoder[code[i:i+2]]
	}

	return result
}

func main() {
	code := "220411112603141304"
	fmt.Println(codeToString(code))
	fmt.Println(codeToStringMap(code))
}
