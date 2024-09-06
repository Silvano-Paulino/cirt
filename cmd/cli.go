package cmd

import (
	"github.com/silvano-paulino/cirt"
	"github.com/spf13/cobra"
)

func NewCirt() *cobra.Command {

	var (
		salaryBase, subSidioA, subSidioT, premeo float64
		days, late                               int
	)

	cirt := cobra.Command{
		Use:       "cirt",
		Short:     "calculate IRT",
		ValidArgs: []string{"s", "d", "f", "p", "a", "t"},
		Run: func(cmd *cobra.Command, args []string) {

			if invalid := Validate(salaryBase, days, late, subSidioA, subSidioT, premeo); !invalid {
				panic("invalid input")
			}

			service := cirt.NewService()

			salaryFinal, err := service.CalculateSalaryAfterLate(salaryBase, late, days)
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

			sujeicao_a := service.Excesso(subSidio_a)
			sujeicao_t := service.Excesso(subSidio_t)

			sujeicaoIrt := sujeicao(sujeicao_a, sujeicao_t)

			inss := service.CalculateSocialSegurance(salaryBase, subSidio_a, subSidio_t, premeo)
			mc := service.CalculateMateriaColectavel(salaryBase, sujeicaoIrt, inss)

			irt, err := service.CalculateIRT(mc)
			if err != nil {
				panic(err)
			}

			discount := service.DiscountTotal(inss, irt)
			salaryLiquido := service.CalculateSalaryLiquido(salaryBase, subSidioT, subSidioA, premeo, discount)

			Print(salaryBase, salaryFinal, subSidioA, subSidioT, premeo, inss, irt, salaryLiquido)
		},
	}

	cirt.Flags().IntVarP(&days, "diasuteis", "d", 22, "dias úteis de trabalho")
	cirt.Flags().IntVarP(&late, "falta", "f", 0, "faltas")
	cirt.Flags().Float64VarP(&subSidioA, "subalimentacao", "a", 0, "subsídio de alimentação")
	cirt.Flags().Float64VarP(&subSidioT, "subtransporte", "t", 0, "subsídio de transporte")
	cirt.Flags().Float64VarP(&premeo, "premios", "p", 0, "prémios")
	cirt.Flags().Float64VarP(&salaryBase, "salariobase", "s", 0, "salário base")

	return &cirt
}

func sujeicao(a, t float64) float64 {
	return a + t
}
