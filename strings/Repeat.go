/////////////////////////////////////////////////////////////////////
// arataca89@gmail.com
// 20210417
//
// func Repeat(s string, count int) string
//
// Retorna uma nova string com s repetida count vezes.
//
// Se count é negativo ou se len(s) * count provoca um overflow
// Repeat() entra em pânico (panic)
//

package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("ba" + strings.Repeat("na", 2)) // banana
}
