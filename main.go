package main

import (
	"flag"
	"fmt"
)

func sujeicao(a, t float64) float64 {
	return a + t
}

func main() {
	var (
		salaryBase,
		subSidioA,
		subSidioT,
		premeo float64
		days,
		late int
	)

	if v := validate(salaryBase, days, late, subSidioA, subSidioT, premeo); !v {
		panic("invalid input")
	}

	service := Service{}

	salary, err := service.CalculateSalaryAfterLate(salaryBase, late, days)
	if err != nil {
		panic(err)
	}

	subSidio_a, err := service.CalculoSubsiodio(subSidioA, days)
	if err != nil {
		panic(err)
	}

	subSidio_t, err := service.CalculoSubsiodio(subSidioT, days)
	if err != nil {
		panic(err)
	}

	sujeicao_a := service.Excesso(subSidioA)
	sujeicao_t := service.Excesso(subSidioT)

	sujeicaoIrt := sujeicao(sujeicao_a, sujeicao_t)

	inss := service.CalculateSocialSegurance(salary, subSidio_a, subSidio_t, premeo)
	mc := service.CalculateMateriaColectavel(salaryBase, sujeicaoIrt, inss)

	irt, err := service.CalculateIRT(mc)
	if err != nil {
		panic(err)
	}

	discount := service.DiscountTotal(inss, irt)
	salaryLiquido := service.CalculateSalaryLiquido(salaryBase, subSidioT, subSidioA, premeo, discount)

	print(salaryBase, salary, subSidio_a, subSidio_a, premeo, inss, irt, salaryLiquido)

	flag.Float64Var(&salaryBase, "s", 0, "salary base")
	flag.Float64Var(&subSidioA, "a", 0, "subsidio alimentacao")
	flag.Float64Var(&subSidioT, "t", 0, "subsidio transporte")
	flag.Float64Var(&premeo, "p", 0, "premeo")
	flag.IntVar(&late, "f", 0, "lates")
	flag.IntVar(&days, "d", 22, "days utils")
	flag.Parse()

}

func print(salaryBase, salaryBaseAfterLate, subsidioA, subsidioT, premeo, inss, irt, salaryLiquido float64) {
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

func validate(salaryBase float64, days, late int, subsidioA, subsidioT, premeo float64) bool {
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
