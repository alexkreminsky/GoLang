package main

import (
	"fmt"
	"math"
)

func main() {
// простой калькулятор
	var input1, input2 float64
	operator := " "
	text := "Результат равен: "
	fmt.Print("Операторы '+' '-' '*' '/'  и 'q2' - квадратный корень. ")
	fmt.Print("Введите первое число: ")
	fmt.Scanf("%f", &input1)

	fmt.Print("Введите оператор: ")
	fmt.Scanf("%s", &operator)
	if operator == "q2" {
		res := math.Sqrt(input1)
		fmt.Println(text, res)
		return
	}

	fmt.Print("Введите второе число: ")
	fmt.Scanf("%f", &input2)

	if operator == "*" {
		res := input1 * input2
		fmt.Println(text, res)

	} else if operator == "/" && input2 == 0 {
		fmt.Println("Ошибка: деление на ноль")

	} else if operator == "/" {
		res := input1 / input2
		fmt.Println(text, res)

	} else if operator == "-" {
		res := input1 - input2
		fmt.Println(text, res)

	} else if operator == "+" {
		res := input1 + input2
		fmt.Println(text, res)
	} else {
		fmt.Println("Ошибка. Неизвестный оператор")
		//fmt.Print("Введите оператор: ")
		//fmt.Scanf("%s", &operator)
		//func result()

		// Здесь нужно как-то вернуться в начало цикла if, пока не знаю как это сделать
		//Также нужно добавить проверку что введены числа

	}

}

/*func result(){
	if operator == "*" {
		res := input1 * input2
		fmt.Println(text, res)
	} else if operator == "/" && input2 == 0 {
		fmt.Println("Ошибка: деление на ноль")
	} else if operator == "/" {
		res := input1 / input2
		fmt.Println(text, res)
	} else if operator == "-"{
		res := input1 - input2
		fmt.Println(text, res)
	} else if operator == "+" {
		res := input1 + input2
		fmt.Println(text, res)
	} else {
		fmt.Println("Ошибка. Неизвестный оператор")
		fmt.Print("Введите оператор: ")
	    fmt.Scanf("%s", &operator)
}
} */
