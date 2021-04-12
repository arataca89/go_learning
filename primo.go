// primo.go
// verifica se um número é primo
// arataca89@gmail.com
// 20210412
package main

import "fmt"

func main() {
	var nr int
	var divisor int
	isprime := true

	fmt.Print("Entre com o número: ")
	fmt.Scanf("%d", &nr)

	for divisor = 2; divisor < nr; divisor++ {
		if nr%divisor == 0 {
			isprime = false
			break
		}
	}

	if isprime {
		fmt.Println("O número", nr, "é primo")
	} else {
		fmt.Println("O número", nr, "não é primo")
	}
}
