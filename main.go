package main

import (
	"GOLANGOO/contas"
	"fmt"
)

func main() {

	contaDaSilvia := contas.ContaCorrente{}
	contaDaSilvia.Titular = "Silvia"
	contaDaSilvia.Saldo = 500

	contaAdriano := contas.ContaCorrente{}
	contaAdriano.Titular = "Adriano"
	contaAdriano.Saldo = 200

	contaDaSilvia.Transferir(-200, &contaAdriano)
	fmt.Println(contaDaSilvia)
	fmt.Println(contaAdriano)

}
