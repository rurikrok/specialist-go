/*
Задача №4. Шахматная доска
Вход: размер шахматной доски, от 0 до 20
Выход: вывести на экран эту доску, заполняя поля нулями и единицами

Пример:
Вход: 5
Выход:

	0 1 0 1 0
	1 0 1 0 1
	0 1 0 1 0
	1 0 1 0 1
	0 1 0 1 0
*/
package chess

import "fmt"

func PrintChessboard(size int) {
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if (i+j)%2 == 0 {
				fmt.Print("0 ")
			} else {
				fmt.Print("1 ")
			}
		}
		fmt.Println()
	}
}
