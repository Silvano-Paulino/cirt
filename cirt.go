package main

import (
	"errors"
	"math"
)

var (
	ErrSalaryBaseNegative = errors.New("salary base cannot be negative")
	ErrSubsidioNegative   = errors.New("subsidio cannot be negative")
)

type Service struct{}

func NewService() Service {
	return Service{}
}

func (s Service) Excesso(subsidio float64) float64 {
	
	if subsidio < 30000 {
		return 0
	}

	return subsidio - 30000
}

func (s Service) SalaryPerDay(salaryBase float64, days int) float64 {
	return salaryBase / float64(days)
}

func (s Service) round(salary float64) float64 {
	return math.Round(salary*100) / 100
}

func (s Service) CalculateSalaryAfterLate(salaryBase float64, late int, days ...int) (float64, error) {
	var defaultDays int = 22

	if len(days) > 0 && days[0] > 0 {
		defaultDays = days[0]
	}

	if salaryBase < 0 {
		return 0, ErrSalaryBaseNegative
	}

	salaryPerDay := s.SalaryPerDay(salaryBase, defaultDays)

	return s.round(salaryBase - salaryPerDay*float64(late)), nil
}

func (s Service) CalculoSubsiodio(subsidio float64, days ...int) (float64, error) {
	var defaultDays int = 22

	if len(days) > 0 && days[0] > 0 {
		defaultDays = days[0]
	}

	if subsidio < 0 {
		return 0, ErrSubsidioNegative
	}

	return s.round(subsidio * float64(defaultDays)), nil
}

func (s Service) CalculateSocialSegurance(salaryBase, subsidioAlimentacao, subsidioTransPorte, premeo float64) float64 {
	result := salaryBase + subsidioAlimentacao + subsidioTransPorte + premeo*0.03
	return s.round(result)
}

func (s Service) CalculateMateriaColectavel(salaryBase, sujeicaoIrt, inss float64) float64 {
	return s.round(salaryBase + sujeicaoIrt - inss)
}

func (s Service) CalculateIRT(mc float64) (float64, error) {
	data, err := Get(mc)
	if err != nil {
		return 0, err
	}

	return data.ParcelaFixa + mc - data.Excesso*data.Taxa, nil
}

func (s Service) DiscountTotal(inss, irt float64) float64 {
	return s.round(inss + irt)
}

func (s Service) CalculateSalaryLiquido(salaryBase, subSidiT, subSidiA, premeo, discount float64) float64 {
	return s.round(salaryBase + subSidiA + subSidiT + premeo - discount)
}
