/*
Задача №2
Вход:
Пользователь должен ввести правильный пароль, состоящий из:
цифр,
букв латинского алфавита(строчные и прописные) и
специальных символов  special = "_!@#$%^&"

Всего 4 набора различных символов.
В пароле обязательно должен быть хотя бы один символ из каждого набора.
Длина пароля от 8(мин) до 15(макс) символов.
Максимальное количество попыток ввода неправильного пароля - 5.
Каждый раз выводим номер попытки.
*Желательно выводить пояснение, почему пароль не принят и что нужно исправить.

digits = "0123456789"
lowercase = "abcdefghiklmnopqrstvxyz"
uppercase = "ABCDEFGHIKLMNOPQRSTVXYZ"
special = "_!@#$%^&"

Выход:
Написать, что ввели правильный пароль.

Пример:
хороший пароль -> o58anuahaunH!
хороший пароль -> aaaAAA111!!!
плохой пароль -> saucacAusacu8
*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	digits    = "0123456789"
	lowercase = "abcdefghiklmnopqrstvxyz"
	uppercase = "ABCDEFGHIKLMNOPQRSTVXYZ"
	special   = "_!@#$%^&"
)

func containsCharFrom(s, chars string) bool {
	for _, ch := range s {
		if strings.ContainsRune(chars, ch) {
			return true
		}
	}
	return false
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for attempt := 1; attempt <= 5; attempt++ {
		fmt.Printf("Attempt %d:\n", attempt)
		fmt.Print("Enter password: ")
		scanner.Scan()
		password := scanner.Text()

		errors := make([]string, 0)

		if len(password) < 8 || len(password) > 15 {
			errors = append(errors, "The password must be between 8 and 15 characters.")
		}
		if !containsCharFrom(password, digits) {
			errors = append(errors, "The password must contain at least one digit.")
		}
		if !containsCharFrom(password, lowercase) {
			errors = append(errors, "The password must contain at least one lowercase letter.")
		}
		if !containsCharFrom(password, uppercase) {
			errors = append(errors, "The password must contain at least one uppercase letter.")
		}
		if !containsCharFrom(password, special) {
			errors = append(errors, "The password must contain at least one special character.")
		}

		if len(errors) > 0 {
			fmt.Println("Errors:")
			for _, err := range errors {
				fmt.Println(err)
			}
			continue
		}

		fmt.Println("You have entered a valid password.")
		return
	}

	fmt.Println("You have failed to enter a valid password in 5 attempts.")
}
