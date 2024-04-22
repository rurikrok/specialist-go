/*
Задача № 3. Вывести на экран в порядке возрастания три введенных числа
Пример:
Вход: 1, 9, 2
Выход: 1, 2, 9
*/
package main

import "fmt"

func main() {
	var a, b, c int

	fmt.Print("Enter three numbers: ")
	fmt.Scan(&a, &b, &c)

	if a > b {
		a, b = b, a
	}
	if a > c {
		a, c = c, a
	}
	if b > c {
		b, c = c, b
	}

	fmt.Printf("The numbers in ascending order are: %d, %d, %d\n", a, b, c)
}
