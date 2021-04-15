// Demonstra o uso de funções do pacote "flag"
// arataca89@gmail.com
// 20210415
//
// Referências:
// (DONOVAN E KERNIGHAN, 2017)
// https://golang.org/pkg/flag/
// https://gobyexample.com/command-line-flags

package main

import (
	"flag"
	"fmt"
)

func main() {
	// flag.String(name string, value string, usage string) *string
	strFlag := flag.String("str", "flag", "flag do tipo string")

	// flag.Int(name string, value int, usage string) *int
	intFlag := flag.Int("nr", 13, "flag do tipo int")

	// flag.Bool(name string, value bool, usage string) *bool
	boolFlag := flag.Bool("bol", true, "flag do tipo boolean")

	flag.Parse()

	fmt.Println("str:", *strFlag)
	fmt.Println("nr:", *intFlag)
	fmt.Println("bol:", *boolFlag)

	// func Args() []string - retorna os argumentos de linha de
	// comando que não são flags
	fmt.Println("args:", flag.Args())
}

/////////////////////////////////////////////////////////////////////
//
// compile com o comando: go build flags.go
//
// será criado o arquivo flags.exe (no windows)
//
// TESTES:
//
// C:\Users\nerd\golang\go_estudo>flags
// str: flag
// nr: 13
// bol: true
// args: []
//
// C:\Users\nerd\golang\go_estudo>flags -str=zurg -nr=41 -bol=false arg1 arg2 arg3
// str: zurg
// nr: 41
// bol: false
// args: [arg1 arg2 arg3]
//
// C:\Users\nerd\golang\go_estudo>flags -h
// Usage of flags:
//  -bol
//        flag do tipo boolean (default true)
//  -nr int
//        flag do tipo int (default 13)
//  -str string
//        flag do tipo string (default "flag")
//
