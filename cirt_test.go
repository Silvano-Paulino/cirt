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

			expected := 18000

			// Act
			result, _ := service.CalculateSalaryAfterLate(salaryBase, late, days)

			// Assert
			if result != float64(expected) {
				t.Errorf("Expeted %v, received %v", expected, result)
			}
		})

		t.Run("Should calculate salary base after late when default days is 22", func(t *testing.T) {
			// Arrange
			service := cirt.NewService()
			salaryBase := 22000.0
			late := 2

			expected := 20000

			// Act
			result, _ := service.CalculateSalaryAfterLate(salaryBase, late)

			// Assert
			if result != float64(expected) {
				t.Errorf("Expeted %v, received %v", expected, result)
			}
		})

		t.Run("Should calculate salary base after late when salary base is more than one decimal places", func(t *testing.T) {
			// Arrange
			service := cirt.NewService()
			salaryBase := 22000.45
			late := 2

			expected := 20000.41

			// Act
			result, _ := service.CalculateSalaryAfterLate(salaryBase, late)

			// Assert
			if result != float64(expected) {
				t.Errorf("Expeted %v, received %v", expected, result)
			}
		})

		t.Run("Should not calculate salary base after late when salary base is negative", func(t *testing.T) {
			// Arrange
			service := cirt.NewService()
			salaryBase := -22000.45
			late := 2

			// Act
			_, err := service.CalculateSalaryAfterLate(salaryBase, late)

			// Assert
			if err != cirt.ErrSalaryBaseNegative {
				t.Error(err)
			}
		})

	})

	t.Run("Subsidio de Alimentação", func(t *testing.T) {
		t.Run("Should calculate subsidio de alimentação", func(t *testing.T) {
			// Arrange
			service := cirt.NewService()
			subsidio := 1000.0
			days := 20

			expected := 20000

			// Act
			result := service.CalculoSubisiodio(subsidio, days)

			// Assert
			if result != float64(expected) {
				t.Errorf("Expected %v, received %v", expected, result)
			}
		})

		t.Run("Should calculate subsidio de alimentação when default days is 22", func(t *testing.T) {
			// Arrange
			service := cirt.NewService()
			subsidio := 1000.0

			expected := 22000

			// Act
			result := service.CalculoSubisiodio(subsidio)

			// Assert
			if result != float64(expected) {
				t.Errorf("Expected %v, received %v", expected, result)
			}
		})
	})
}
