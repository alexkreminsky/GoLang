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
	l := 2 * pi * radius

	fmt.Println("Радиус окружности площадью ", s, "равен: ", radius)
	fmt.Println("Длина этой окружности равна: ", l)

}
