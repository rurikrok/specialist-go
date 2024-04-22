package main

import (
	"fmt"

	"day2.com/bottles"
	"day2.com/chess"
	"day2.com/triangle"
)

func main() {
	var bottlesCount int

	fmt.Print("Enter the number of bottles: ")
	fmt.Scan(&bottlesCount)

	if bottlesCount < 0 || bottlesCount > 200 {
		fmt.Println("Error: The number of bottles must be between 0 and 200.")
	} else {
		fmt.Printf("%d %s\n", bottlesCount, bottles.BottleWord(bottlesCount))
	}

	var x1, y1, x2, y2, x3, y3 float64

	fmt.Print("Enter the coordinates of three points: ")
	fmt.Scan(&x1, &y1, &x2, &y2, &x3, &y3)

	if !triangle.СanFormTriangle(x1, y1, x2, y2, x3, y3) {
		fmt.Println("The points do not form a triangle.")
	} else {
		fmt.Printf("The area of the triangle is: %.2f\n", triangle.TriangleArea(x1, y1, x2, y2, x3, y3))
		if triangle.IsRightTriangle(x1, y1, x2, y2, x3, y3) {
			fmt.Println("The triangle is a right triangle.")
		} else {
			fmt.Println("The triangle is not a right triangle.")
		}
	}

	var size int

	fmt.Print("Enter the size of the chessboard: ")
	fmt.Scan(&size)

	if size < 0 || size > 20 {
		fmt.Println("Error: The size must be between 0 and 20.")
	} else {
		chess.PrintChessboard(size)
	}
}
