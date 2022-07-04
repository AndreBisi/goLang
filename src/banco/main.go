package main

import (
	"banco/clientes"
	"banco/contas"
	"fmt"
)

func main() {
	//contaDoAndre := ContaCorrente{}

	// //ou
	// var contaDoAndre2 *ContaCorrente
	// contaDoAndre2 = new(ContaCorrente)
	// contaDoAndre2.titular = "Andre 2"
	// fmt.Println(*contaDoAndre2)

	// contaDoAndre.titular = "André"
	// contaDoAndre.numeroAgencia = 123
	// contaDoAndre.numeroConta = 456
	// contaDoAndre.saldo = 125.62

	// contaDoGuilherme := ContaCorrente{titular: "Guilherme", numeroAgencia: 135, numeroConta: 246, saldo: 158.89}

	// contaDaAna := ContaCorrente{"Ana", 486, 587, 55.58}

	// fmt.Println(contaDoAndre)
	// fmt.Println(contaDoGuilherme)
	// fmt.Println(contaDaAna)

	// contaDaSilvia := contas.ContaCorrente{}
	// contaDaSilvia.Titular = "Silvia"
	// contaDaSilvia.Saldo = 500

	// contaDoGustavo := contas.ContaCorrente{}
	// contaDoGustavo.Titular = "Gustavo"
	// contaDoGustavo.Saldo = 300

	// fmt.Println(contaDaSilvia)

	// fmt.Println(contaDaSilvia.Sacar(300))
	// fmt.Println(contaDaSilvia)

	// fmt.Println(contaDaSilvia.Depositar(800))
	// fmt.Println(contaDaSilvia)

	// 	status := contaDaSilvia.Transferir(600, &contaDoGustavo)
	// 	fmt.Println(status)
	// 	fmt.Println(contaDaSilvia)
	// 	fmt.Println(contaDoGustavo)

	contaDoBruno := contas.ContaCorrente{Titular: clientes.Titular{
		Nome:      "Bruno",
		CPF:       "123.111.123.12",
		Profissao: "T.I."},
		NumeroAgencia: 123,
		NumeroConta:   123456}
	fmt.Println(contaDoBruno)

	PagarBoleto(&contaDoBruno, 30)

}

func PagarBoleto(conta VerificaConta, valorDoBoleto float64) {
	conta.Sacar(valorDoBoleto)
}

type VerificaConta interface {
	Sacar(valor float64) string
}
