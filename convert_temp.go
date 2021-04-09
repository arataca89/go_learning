package main

import "fmt"

func main() {
	var fahr float64
	fmt.Print("Entre com a temperatura em Fahrenheit: ")
	fmt.Scanf("%f", &fahr)
	celsius := (fahr - 32) * 5 / 9
	fmt.Println("A temperatura em Celsius é", celsius)
}
