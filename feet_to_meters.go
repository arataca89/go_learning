// feet_to_meter.go
// converte pés em metros
// arataca89@gmail.com
// 20210409

package main

import "fmt"

func main() {
	var feet float64

	fmt.Print("Entre com a medida em pés:")
	fmt.Scanf("%f", &feet)

	meters := feet * 0.3048

	fmt.Println("A medida em metros é", meters)
}
