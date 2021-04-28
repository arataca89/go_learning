// Gera números aleatórios entre um um valor mínimo e um valor máximo
// arataca89@gmail.com
// 20210428

package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Retorna um número aleatório entre min e max
// OBS: inserir rand.Seed(time.Now().UnixNano()) na função main()
func randInt(min, max int) int {
	if min < 0 || max <= 0 || min > max {
		return 0
	}
	if min == max {
		return min
	}
	return rand.Intn(max-min) + min
}

func main() {
	rand.Seed(time.Now().UnixNano())
	for i := 1; i <= 100; i++ {
		fmt.Printf("%d ", randInt(0, 9))
		if i%10 == 0 {
			fmt.Printf("\n")
		}
	}
}

/////////////////////////////////////////////////////////////////////
//
// TESTE
// C:\Users\nerd\golang\go_estudo>go run randInt.go
// 8 6 5 0 3 1 6 3 0 0
// 4 8 3 2 3 4 5 4 8 1
// 2 2 1 1 6 5 7 8 3 5
// 2 3 1 4 0 5 4 1 5 4
// 0 1 3 8 3 2 3 0 3 0
// 4 6 2 2 0 8 8 3 5 0
// 0 7 0 5 2 2 6 0 7 4
// 3 0 8 5 4 5 2 4 1 2
// 5 8 5 7 2 4 2 0 4 6
// 0 2 0 0 4 5 0 0 3 2

// C:\Users\nerd\golang\go_estudo>go run randInt.go
// 2 8 8 7 1 6 3 2 6 3
// 3 7 3 1 1 1 1 1 4 2
// 3 8 2 3 1 8 1 3 7 6
// 6 6 2 3 2 3 0 5 1 7
// 4 5 5 7 8 1 6 0 4 7
// 3 5 7 7 6 4 4 0 4 1
// 4 7 4 5 5 5 2 4 8 7
// 7 8 4 5 6 7 6 5 4 2
// 6 8 7 8 6 6 3 6 3 7
// 2 5 0 2 0 0 6 0 7 4

// C:\Users\nerd\golang\go_estudo>go run randInt.go
// 3 0 6 8 3 4 0 3 4 4
// 8 4 7 4 0 3 1 8 5 2
// 1 4 3 2 5 7 2 2 8 1
// 4 5 2 1 6 1 5 8 4 8
// 2 7 0 3 4 0 0 1 5 5
// 8 5 1 3 7 4 8 1 3 6
// 8 5 3 0 2 5 0 6 1 1
// 0 3 2 5 0 5 8 2 3 8
// 5 6 6 0 2 5 0 3 8 8
// 1 5 2 7 1 6 5 2 4 8

// C:\Users\nerd\golang\go_estudo>
