package main

import (
	"fmt"
	"math"
)

func main() {

	var s float64
	pi := math.Pi

	fmt.Print("Введите площадь окружности: ")
	fmt.Scanf("%f", &s)

	radius := math.Sqrt(s) / pi
	len := 2 * pi * radius

	fmt.Println("Диаметр окружности площадью ", s, "равен: ", radius*2)
	fmt.Println("Длина этой окружности равна: ", len)

}
