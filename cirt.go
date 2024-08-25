package cirt

type Service struct {}

func NewService() Service {
	return Service{}
}

func (s Service) SalaryPerDay(salaryBase float64, days int) float64 {
	return salaryBase / float64(days)
}
