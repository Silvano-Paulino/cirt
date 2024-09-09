package cirt

import (
	"errors"
	"math"
)

var (
	ErrSalaryBaseNegative = errors.New("salary base cannot be negative")
	ErrSubsidyNegative    = errors.New("subsidio cannot be negative")
)

type Service struct{}

func NewService() Service {
	return Service{}
}

func (s Service) Excess(subsidy float64) float64 {

	if subsidy < 30000 {
		return 0
	}

	return subsidy - 30000
}

func (s Service) salaryPerDay(salaryBase float64, days int) float64 {
	return salaryBase / float64(days)
}

func (s Service) round(salary float64) float64 {
	return math.Round(salary*100) / 100
}

func (s Service) CalculateSalaryAfterDelay(salaryBase float64, delay int, days ...int) (float64, error) {
	var defaultDays int = 22

	if len(days) > 0 && days[0] > 0 {
		defaultDays = days[0]
	}

	if salaryBase < 0 {
		return 0, ErrSalaryBaseNegative
	}

	salaryPerDay := s.salaryPerDay(salaryBase, defaultDays)

	return s.round(salaryBase - salaryPerDay*float64(delay)), nil
}

func (s Service) CalculateSubsidy(subsidy float64, days ...int) (float64, error) {
	var defaultDays int = 22

	if len(days) > 0 && days[0] > 0 {
		defaultDays = days[0]
	}

	if subsidy < 0 {
		return 0, ErrSubsidyNegative
	}

	return s.round(subsidy * float64(defaultDays)), nil
}

func (s Service) CalculateSocialSegurance(salaryBase, subsidioAlimentacao, subsidioTransPorte, premeo float64) float64 {
	result := salaryBase + subsidioAlimentacao + subsidioTransPorte + premeo*0.03
	return s.round(result)
}

func (s Service) CalculateColletableMaterial(salaryBase, sujeicaoIrt, inss float64) float64 {
	return s.round(salaryBase + sujeicaoIrt - inss)
}

func (s Service) CalculateIRT(mc float64) (float64, error) {
	data, err := Get(mc)
	if err != nil {
		return 0, err
	}

	return data.ParcelaFixa + mc - data.Excess*data.Tax, nil
}

func (s Service) DiscountTotal(inss, irt float64) float64 {
	return s.round(inss + irt)
}

func (s Service) CalculateSalaryLiquido(salaryBase, subsidyTransport, subsidyFood, premeo, discount float64) float64 {
	return s.round(salaryBase + subsidyFood + subsidyTransport + premeo - discount)
}
