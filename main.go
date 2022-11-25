package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func main() {
	contaDoGuilherme := ContaCorrente{titular: "Guilherme"}
	fmt.Println(contaDoGuilherme)
	contaDaSabrina := ContaCorrente{"Sabrina", 12555, 565656, 156.30}
	fmt.Println(contaDaSabrina)

}
