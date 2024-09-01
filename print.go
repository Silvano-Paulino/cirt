package main

import "fmt"

func Print(salaryBase, salaryBaseAfterLate, subsidioA, subsidioT, premeo, inss, irt, salaryLiquido float64) {
	fmt.Println()
	fmt.Println("Salário Base: ", salaryBase, "kz")
	fmt.Println()
	fmt.Println("========= Incrementos ========")
	fmt.Println("Subsídio de Alimentação: ", subsidioA, "kz")
	fmt.Println("Subsídio de Transporte: ", subsidioT, "kz")
	fmt.Println("Prémios: ", premeo, "kz")
	fmt.Println()
	fmt.Println("========= Descontos ========")
	fmt.Println("Salário Base Após Falta: ", salaryBaseAfterLate, "kz")
	fmt.Println("Segurança Social: ", inss, "kz")
	fmt.Println("IRT: ", irt, "kz")
	fmt.Println("Salário Líquido: ", salaryLiquido, "kz")
	fmt.Println()
}
