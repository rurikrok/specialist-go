/*
Задача № 2. Получить реверсную запись трехзначного числа
Пример:
вход: 346, выход: 643
вход: 100, выход: 001
*/
package main

import "fmt"

func main() {
	var number int

	fmt.Print("Enter a three-digit number: ")
	fmt.Scan(&number)

	if number < 100 || number > 999 {
		fmt.Println("Error: The number must be a three-digit number.")
		return
	}

	firstDigit := number / 100
	secondDigit := (number / 10) % 10
	thirdDigit := number % 10

	reversed := thirdDigit*100 + secondDigit*10 + firstDigit

	fmt.Printf("The reversed number is: %03d\n", reversed)
}
