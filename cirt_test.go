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

	t.Run("Subsidio de Alimentação ou Transport", func(t *testing.T) {
		t.Run("Should calculate subsidio", func(t *testing.T) {
			// Arrange
			service := cirt.NewService()
			subsidio := 1000.0
			days := 20

			expected := 20000

			// Act
			result, _ := service.CalculoSubisiodioAlimentacaoOuTransporte(subsidio, days)

			// Assert
			if result != float64(expected) {
				t.Errorf("Expected %v, received %v", expected, result)
			}
		})

		t.Run("Should calculate subsidio when default days is 22", func(t *testing.T) {
			// Arrange
			service := cirt.NewService()
			subsidio := 1000.0

			expected := 22000

			// Act
			result, _ := service.CalculoSubisiodioAlimentacaoOuTransporte(subsidio)

			// Assert
			if result != float64(expected) {
				t.Errorf("Expected %v, received %v", expected, result)
			}
		})

		t.Run("Should calculate subsidio with more than one decimal places", func(t *testing.T) {
			// Arrange
			service := cirt.NewService()
			subsidio := 1000.459

			expected := 22010.10

			// Act
			result, _ := service.CalculoSubisiodioAlimentacaoOuTransporte(subsidio)

			// Assert
			if result != float64(expected) {
				t.Errorf("Expected %v, received %v", expected, result)
			}
		})

		t.Run("Should not calculate subsidio when subsidio is negative", func(t *testing.T) {
			// Arrange
			service := cirt.NewService()
			subsidio := -1000.0

			// Act
			_, err := service.CalculoSubisiodioAlimentacaoOuTransporte(subsidio)

			// Assert
			if err != cirt.ErrSubsidioNegative {
				t.Error(err)
			}
		})
	})

	t.Run("should calculate excesso", func(t *testing.T) {
		// Arrange
		service := cirt.NewService()
		subsidio := 2000.0

		expected := 14000

		result, _ := service.CalculoSubisiodioAlimentacaoOuTransporte(subsidio)

		// Act
		execesso := service.Excesso(result)

		// Assert
		if execesso != float64(expected) {
			t.Errorf("Expected %v, received %v", expected, result)
		}
	})

	t.Run("Social Segurance", func(t *testing.T) {
		t.Run("Should calculate inss", func(t *testing.T) {
			// Arrange
			service := cirt.NewService()
			salaryBase := 20000.0
			subsidioAlimentacao := 1000.0
			subsidioTransPorte := 1000.0
			premeo := 5000.0

			expected := 22150

			// Act
			result := service.CalculateSocialSegurance(salaryBase, subsidioAlimentacao, subsidioTransPorte, premeo)

			// assert
			if result != float64(expected) {
				t.Errorf("Expected %v, received %v", expected, result)
			}
		})

		t.Run("Should calculate inss with more than one decimal places", func(t *testing.T) {
			// Arrange
			service := cirt.NewService()
			salaryBase := 20000.256
			subsidioAlimentacao := 1000.0
			subsidioTransPorte := 1000.0
			premeo := 5000.0

			expected := 22150.26

			// Act
			result := service.CalculateSocialSegurance(salaryBase, subsidioAlimentacao, subsidioTransPorte, premeo)

			// assert
			if result != float64(expected) {
				t.Errorf("Expected %v, received %v", expected, result)
			}
		})
	})
}
