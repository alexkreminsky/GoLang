package main		// Поиск простого числа

import "fmt"

func isPrime(N int) (primes []int) {
	b := make([]bool, N)
	for i := 2; i < N; i++ {
		if b[i] == true {
			continue
		}
		primes = append(primes, i)
		for k := i * i; k < N; k += i {
			b[k] = true
		}
	}
	return
}

func main() {

	fmt.Print("Введите любое целое число: ") //}  Этот кусок кода мой
	var input int                            //}  остальное взял
	fmt.Scanf("%d", &input)                  //}  из интеренета

	primes := isPrime(input)
	for _, p := range primes {
		fmt.Println(p)
	}
}
