package cirt

type Service struct{}

func NewService() Service {
	return Service{}
}

func (s Service) SalaryPerDay(salaryBase float64, days int) float64 {
	return salaryBase / float64(days)
}

func (s Service) CalculateSalaryAfterLate(salaryBase float64, days, late int) float64 {
	salaryPerDay := s.SalaryPerDay(salaryBase, days)
	return salaryBase - salaryPerDay * float64(late)
}
