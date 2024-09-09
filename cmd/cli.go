package cmd

import (
	"github.com/silvano-paulino/cirt"
	"github.com/spf13/cobra"
)

func NewCirt() *cobra.Command {

	var (
		salaryBase, subsidyFood, subsidyTransport, premeo float64
		days, delay                                       int
	)

	cirt := cobra.Command{
		Use:       "cirt",
		Short:     "calculate IRT",
		ValidArgs: []string{"s", "d", "f", "p", "a", "t"},
		Run: func(cmd *cobra.Command, args []string) {

			if invalid := Validate(salaryBase, days, delay, subsidyFood, subsidyTransport, premeo); !invalid {
				panic(invalid)
			}

			service := cirt.NewService()

			salaryFinal, err := service.CalculateSalaryAfterDelay(salaryBase, delay, days)
			if err != nil {
				panic(err)
			}

			subsidy_food, err := service.CalculateSubsidy(subsidyFood, days)
			if err != nil {
				panic(err)
			}

			subsidy_transport, err := service.CalculateSubsidy(subsidyTransport, days)
			if err != nil {
				panic(err)
			}

			excess_food := service.Excess(subsidy_food)
			excess_transport := service.Excess(subsidy_transport)

			sujeicaoIrt := sujeicao(excess_food, excess_transport)

			inss := service.CalculateSocialSegurance(salaryBase, subsidy_food, subsidy_transport, premeo)
			mc := service.CalculateColletableMaterial(salaryBase, sujeicaoIrt, inss)

			irt, err := service.CalculateIRT(mc)
			if err != nil {
				panic(err)
			}

			discount := service.DiscountTotal(inss, irt)
			salaryLiquido := service.CalculateSalaryLiquido(salaryBase, subsidyTransport, subsidyFood, premeo, discount)

			Print(salaryBase, salaryFinal, subsidyFood, subsidyTransport, premeo, inss, irt, salaryLiquido)
		},
	}

	cirt.Flags().IntVarP(&days, "diasuteis", "d", 22, "dias úteis de trabalho")
	cirt.Flags().IntVarP(&delay, "delay", "f", 0, "faltas")
	cirt.Flags().Float64VarP(&subsidyFood, "subsidy-food", "a", 0, "subsidy food")
	cirt.Flags().Float64VarP(&subsidyTransport, "subsidy-transport", "t", 0, "subsidy transport")
	cirt.Flags().Float64VarP(&premeo, "premeo", "p", 0, "premeo")
	cirt.Flags().Float64VarP(&salaryBase, "salarybase", "s", 0, "salary base")

	return &cirt
}

func sujeicao(a, t float64) float64 {
	return a + t
}
