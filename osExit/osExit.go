/*
A função os.Exit() é usada para encerrar o programa com um
código de status fornecido. Normalmente zero indica sucesso
e um status diferente de zero indica erro. O programa termina
imediatamente. Funções defer não serão executadas.
Para portabilidade os códigos de status devem estar entre
[0, 125].

Segundo o Advanced Bash-Scripting Guide alguns códigos de status
têm significado especial. Entre eles temos:

1		Usado para erros em geral, como "dividir por zero"
		e outras operações inadmissíveis.

2		Palavra-chave ou comando ausente; ou problema de permissões.

Assim, convém que tais códigos sejam usados neste contexto.

Fora isso o uso de código de status deve ser documentado para
uma melhor compreensão do código.

Referências:
https://golang.org/pkg/os/#Exit
https://tldp.org/LDP/abs/html/exitcodes.html
https://golangcode.com/exit-application-with-error-code/

arataca89@gmail.com
20210608
*/
package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("usando os.Exit()")
	os.Exit(0)
	fmt.Println("isto não é exibido")

}
