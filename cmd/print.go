package cmd

import "fmt"

func Print(salaryBase, salaryFinal, subsidyFood, subsidyTransport, premeo, inss, irt, salaryLiquid float64) {
	fmt.Println()
	fmt.Println("Salary Base: ", salaryBase, "kz")
	fmt.Println()
	fmt.Println("========= Increments ========")
	fmt.Println("Subsidy Food: ", subsidyFood, "kz")
	fmt.Println("SubsidyTransport: ", subsidyTransport, "kz")
	fmt.Println("Premeo: ", premeo, "kz")
	fmt.Println()
	fmt.Println("========= Discounts ========")
	fmt.Println("Salary Base after Delay: ", salaryFinal, "kz")
	fmt.Println("Social Segurance: ", inss, "kz")
	fmt.Println("IRT: ", irt, "kz")
	fmt.Println("Salary Liquid: ", salaryLiquid, "kz")
	fmt.Println()
}
