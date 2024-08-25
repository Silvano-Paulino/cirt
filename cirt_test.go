package cirt_test

import (
	"testing"

	cirt "github.com/silvano-paulino"
)

func TestCirt(t *testing.T) {
	t.Run("Salary base after late", func(t *testing.T) {

		t.Run("Should calculate salary base per day", func(t *testing.T) {
			// Arrange
			service := cirt.NewService()
			salaryBase := 20000.0
			days := 20

			expeted := 1000

			// Act
			result := service.SalaryPerDay(salaryBase, days)

			// Assert
			if result != float64(expeted) {
				t.Errorf("Expeted %v, received %v", expeted, result)
			}
		})

		t.Run("Should calculate salary base after late", func(t *testing.T) {
			// Arrange
			service := cirt.NewService()
			salaryBase := 20000.0
			days := 20
			late := 2

			expected := 2000

			// Act
			result := service.CalculateSalaryAfterLate(salaryBase, days, late)

			// Assert
			if result != float64(expected) {
				t.Errorf("Expeted %v, received %v", expected, result)
			}
		})

	})

}
