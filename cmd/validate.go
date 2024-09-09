package cmd

import "fmt"

func Validate(salaryBase float64, days, delay int, subsidyFood, subsidyTransport, premeo float64) bool {
	if salaryBase <= 0 {
		fmt.Println("salary bse cannot be negative or equals zero.")
		return false
	}
	if days <= 0 || days > 31 {
		fmt.Println("os dias úteis de trabalho não podem ser negativos ou maiores que 31 dias.")
		return false
	}
	if delay < 0 || delay > days {
		fmt.Println("delay cannot be negative or up than days.")
		return false
	}
	if subsidyFood < 0 || subsidyFood > salaryBase {
		fmt.Println("subsidy food cannot be neagtive or up than salary base.")
		return false
	}
	if subsidyTransport < 0 || subsidyTransport > salaryBase {
		fmt.Println("subsidy transport cannot be neagtive or up than salary base.")
		return false
	}
	if premeo < 0 {
		fmt.Println("premeo cannot be negative")
		return false
	}
	return true
}
