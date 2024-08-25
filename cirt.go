package cirt

import (
	"errors"
	"math"
)

var (
	ErrSalaryBaseNegative = errors.New("Salary base cannot be negative")
)

type Service struct{}

func NewService() Service {
	return Service{}
}

func (s Service) Excesso(subsidio float64) float64 {
	return subsidio - 30000
}

func (s Service) SalaryPerDay(salaryBase float64, days int) float64 {
	return salaryBase / float64(days)
}

func (s Service) round(salary float64, place int) float64 {
	factor := math.Pow(10, float64(place))
	return math.Round(salary*factor) / factor
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

	return s.round(salaryBase-salaryPerDay*float64(late), 2), nil
}

func (s Service) CalculoSubisiodio(subsidio float64, days ...int) float64 {
	var defaultDays int = 22

	if len(days) > 0 && days[0] > 0 {
		defaultDays = days[0]
	}

	result := subsidio * float64(defaultDays)

	return result
}
