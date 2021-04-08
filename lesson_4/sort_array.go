package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	arr := [5]int64{}

	scaner := bufio.NewScanner(os.Stdin)
	for i := len(arr) - 1; i >= 0; i-- {
		if scaner.Scan() {
			n, err := strconv.ParseInt(scaner.Text(), 10, 64)
			if err != nil {
				fmt.Println("Ошибка, нужно было ввести числа")
				return
			}
			arr[i] = n

		}
	}
	fmt.Println(BubbleSort(arr))

}

func BubbleSort(arr [5]int64) [5]int64 {
	for i := 0; i < len(arr); i++ {
		for j := len(arr) - 1; j > i; j-- {
			if arr[j-1] > arr[j] {
				arr[j-1], arr[j] = arr[j], arr[j-1]

			}
		}
	}
	return arr

}
