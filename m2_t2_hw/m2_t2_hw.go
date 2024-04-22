/*
Задача № 4. Проверить, является ли четырехзначное число палиндромом
Пример:
Вход: 1221  Выход: 1221 - палиндром
Вход: 1234  Выход: 1234 - не палиндром
*/
package main

import "fmt"

func main() {
	var number int

	fmt.Print("Enter a four-digit number: ")
	fmt.Scan(&number)

	if number < 1000 || number > 9999 {
		fmt.Println("Error: The number must be a four-digit number.")
		return
	}

	firstDigit := number / 1000
	secondDigit := (number / 100) % 10
	thirdDigit := (number / 10) % 10
	fourthDigit := number % 10

	reversed := fourthDigit*1000 + thirdDigit*100 + secondDigit*10 + firstDigit

	if number == reversed {
		fmt.Printf("%d is a palindrome\n", number)
	} else {
		fmt.Printf("%d is not a palindrome\n", number)
	}
}
