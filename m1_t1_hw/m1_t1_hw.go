/*
Задача №1
Вход:

	расстояние(50 - 10000 км),
	расход в литрах (5-25 литров) на 100 км и
	стоимость бензина(константа) = 48 руб

Выход: стоимость поездки в рублях
*/
package main

import "fmt"

func main() {
	var distance, fuelConsumption float64
	const gasolineCost = 48.0

	fmt.Print("Enter distance in km: ")
	fmt.Scan(&distance)

	if distance < 50 || distance > 10000 {
		fmt.Println("Error: Distance must be between 50 and 10000 km.")
		return
	}

	fmt.Print("Enter fuel consumption in liters per 100 km: ")
	fmt.Scan(&fuelConsumption)

	if fuelConsumption < 5 || fuelConsumption > 25 {
		fmt.Println("Error: Fuel consumption must be between 5 and 25 liters per 100 km.")
		return
	}

	totalFuel := (distance * fuelConsumption) / 100
	totalCost := totalFuel * gasolineCost

	fmt.Printf("The cost of the trip is: %.2f rubles\n", totalCost)
}
