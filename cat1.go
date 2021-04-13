// Implementação do comando cat versão 1
// Lê arquivos da linha de comando e os exibe.
// arataca89@gmail.com
// 20210413

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	files := os.Args[1:]

	if len(files) == 0 {
		fmt.Println("cat1 exibe o conteúdo de arquivos no console.")
		fmt.Println("USO: cat1 arquivo1 arquvo2 arquivo3 ... arquivo n")
		os.Exit(0)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "ERRO: %v\n", err)
				continue
			}
			printFile(f)
			f.Close()
		}
	}
}

func printFile(f *os.File) {
	fmt.Println("")
	input := bufio.NewScanner(f)
	for input.Scan() {
		fmt.Println(input.Text())
	}
}

/////////////////////////////////////////////////////////////////////
//
// func NewScanner(r io.Reader) *Scanner
//
// NewScanner retorna um Scanner a ser lido.A função split padrão
// é ScanLines()
//
// Scanner fornece uma interface para ler dados de um arquivo com
// linhas de texto terminadas pelo caractere de nova linha.
// Sucessivas chamadas ao método Scan() irão percorrer os “tokens”
// do arquivo, pulando os bytes entre os tokens. A especificação de
// um token é definida por uma função do tipo SplitFunc() cujo padrão
// quebra a entrada em linhas com o caractere de nova linha retirado.
// Funções Split são definidas neste pacote para scanear um arquivo
// em linhas, bytes, UTF-8-encoded runes e palavras delimitadas por
// espaço. Apesar disso o cliente pode especificar uma função split
// customizada.
// A operação de varredura do arquivo termina  no fim do arquivo
// (EOF), se houver um erro de entrada/saída, ou se o token sendo
// lido for maior que o buffer especificado para armazená-lo.
// Quando a varredura é encerrada a leitura pode ter avançado além do
// último token. Em programas que necessitam mais controle sobre a
// manipulação de erros ou tokens grandes, ou devem usar varreduras
// sequenciais, deve-se usar bufio.Reader() em vez de Scan().
//
//
// func (s *Scanner) Scan() bool
//
// Scan() avança o Scanner para o próximo token. A função retorna
// false quando a varredura termina ao atingir o fim da entrada ou
// ao atingir um error. Após retornar false, o mérodo Err() retornará
// o erro ocorrido durante a varredura, exceto se foi um io.EOF,
// quando Err() retornará nil. Scan() entra em pânico se a função
// split retornar muitos tokens vazios. Isto é um erro comum para
// scanners.
//
//
// func (s *Scanner) Text() string
//
// Text() retorna o último token gerado por uma chamada a Scan()
//
// Referências:
// https://golang.org/pkg/bufio/#NewScanner
// https://golang.org/pkg/bufio/#Scanner
// https://golang.org/pkg/bufio/#Scanner.Scan
// https://golang.org/pkg/bufio/#Scanner.Text
//
