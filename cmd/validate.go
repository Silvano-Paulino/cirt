package cmd

import "fmt"

func Validate(salaryBase float64, days, late int, subsidioA, subsidioT, premeo float64) bool {
	if salaryBase <= 0 {
		fmt.Println("O salário base não pode ser negativo ou menor ou igual a zero.")
		return false
	}
	if days <= 0 || days > 31 {
		fmt.Println("Os dias úteis de trabalho não podem ser negativos ou maiores que 31 dias.")
		return false
	}
	if late < 0 || late > days {
		fmt.Println("O número de faltas não pode ser negativo ou maior que os dias úteis de trabalho.")
		return false
	}
	if subsidioA < 0 || subsidioA > salaryBase {
		fmt.Println("O subsídio de alimentação não pode ser negativo ou maior que o salário base.")
		return false
	}
	if subsidioT < 0 || subsidioT > salaryBase {
		fmt.Println("O subsídio de transporte não pode ser negativo ou maior que o salário base.")
		return false
	}
	if premeo < 0 {
		fmt.Println("O premeo não pode ser um número negativo")
		return false
	}
	return true
}
