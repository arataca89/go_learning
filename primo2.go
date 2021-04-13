// primo2.go
// verifica se um número é primo
// arataca89@gmail.com
// 20210413
package main

import "fmt"

func isprime(nr int) bool {

	for divisor := 2; divisor < nr; divisor++ {
		if nr%divisor == 0 {
			return false
		}
	}

	return true

}

func main() {
	var nr int

	fmt.Print("Entre com o número: ")
	fmt.Scanf("%d", &nr)

	if isprime(nr) {
		fmt.Println("O número", nr, "é primo")
	} else {
		fmt.Println("O número", nr, "não é primo")
	}

}
