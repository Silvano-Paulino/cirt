package cmd

import "fmt"

func Print(salaryBase, salaryFinal, subSidioA, subSidioT, premeo, inss, irt, salaryLiquido float64) {
	fmt.Println()
	fmt.Println("Salário Base: ", salaryBase, "kz")
	fmt.Println()
	fmt.Println("========= Incrementos ========")
	fmt.Println("Subsídio de Alimentação: ", subSidioA, "kz")
	fmt.Println("Subsídio de Transporte: ", subSidioT, "kz")
	fmt.Println("Prémeos: ", premeo, "kz")
	fmt.Println()
	fmt.Println("========= Descontos ========")
	fmt.Println("Salário Base Após Falta: ", salaryFinal, "kz")
	fmt.Println("Segurança Social: ", inss, "kz")
	fmt.Println("IRT: ", irt, "kz")
	fmt.Println("Salário Líquido: ", salaryLiquido, "kz")
	fmt.Println()
}
