package cirt

type Service struct{}

func NewService() Service {
	return Service{}
}

func (s Service) SalaryPerDay(salaryBase float64, days int) float64 {
	return salaryBase / float64(days)
}

func (s Service) CalculateSalaryAfterLate(salaryBase float64, late int, days ...int) float64 {
	var defaultDays int = 22

	if len(days) > 0 && days[0] > 0 {
		defaultDays = days[0]
	}

	salaryPerDay := s.SalaryPerDay(salaryBase, defaultDays)

	return salaryBase - salaryPerDay*float64(late)
}
