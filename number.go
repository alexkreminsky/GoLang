package main

import "fmt"

func main() {

	var num int
	fmt.Print("Введите любое трехзначное число: ")
	fmt.Scanln(&num)


	a := num / 100
	b := (num - a*100)/10
	c := num - (a*100 + b*10)

	fmt.Println(a, "сотен" , b, "десятков", c, "единиц")


}
