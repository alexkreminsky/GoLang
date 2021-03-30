package main

import "fmt"

func main() {

	var input1, input2 float64

	fmt.Print("Введит длину прямоугольника: ")
	fmt.Scanf("%f", &input1)

	if input1 < 0 {
		fmt.Println("Введенное число меньше нуля")

	}

	fmt.Print("Введите высоту прямоугольника: ")
	fmt.Scanf("%f", &input2)

	if input2 < 0 {
		fmt.Println("Введенное число меньше нуля")

	}
	result := input1 * input2
	fmt.Println("Площадь прямоугольника равна: ", result)

}
