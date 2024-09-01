package main

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

	if v := Validate(salaryBase, days, late, subSidioA, subSidioT, premeo); !v {
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

	cli(salaryBase, subSidioA, subSidioT, premeo, late, days)
}
