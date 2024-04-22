/*
Написать 3 функции.
Даны координаты трех точек(x1, y1, x2, y2, x3, y3), значения(целые) которых >= 0.
Первая функция проверяет, что можно построить треугольник по заданным точкам
Вторая функция вычисляет площадь треугольника.
Третья функция должна определить, является ли треугольник прямоугольным.
*/
package triangle

import (
	"math"
)

func СanFormTriangle(x1, y1, x2, y2, x3, y3 float64) bool {
	a := math.Sqrt((x2-x1)*(x2-x1) + (y2-y1)*(y2-y1))
	b := math.Sqrt((x3-x2)*(x3-x2) + (y3-y2)*(y3-y2))
	c := math.Sqrt((x1-x3)*(x1-x3) + (y1-y3)*(y1-y3))

	return a+b > c && a+c > b && b+c > a
}

func TriangleArea(x1, y1, x2, y2, x3, y3 float64) float64 {
	return math.Abs((x1*(y2-y3) + x2*(y3-y1) + x3*(y1-y2)) / 2.0)
}

func IsRightTriangle(x1, y1, x2, y2, x3, y3 float64) bool {
	a := math.Sqrt((x2-x1)*(x2-x1) + (y2-y1)*(y2-y1))
	b := math.Sqrt((x3-x2)*(x3-x2) + (y3-y2)*(y3-y2))
	c := math.Sqrt((x1-x3)*(x1-x3) + (y1-y3)*(y1-y3))

	return math.Abs(a*a+b*b-c*c) < 1e-10 || math.Abs(a*a+c*c-b*b) < 1e-10 || math.Abs(b*b+c*c-a*a) < 1e-10
}
