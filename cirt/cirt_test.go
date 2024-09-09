package cirt_test

import (
	"testing"

	"github.com/silvano-paulino/cirt"
)

func TestCirt(t *testing.T) {
	t.Run("Salary base after late", func(t *testing.T) {

		t.Run("Should calculate salary base after delay", func(t *testing.T) {
			// Arrange
			service := cirt.NewService()
			salaryBase := 20000.0
			days := 20
			delay := 2

			expected := 18000

			// Act
			result, _ := service.CalculateSalaryAfterDelay(salaryBase, delay, days)

			// Assert
			if result != float64(expected) {
				t.Errorf("Expeted %v, received %v", expected, result)
			}
		})

		t.Run("Should calculate salary base after late when default days is 22", func(t *testing.T) {
			// Arrange
			service := cirt.NewService()
			salaryBase := 22000.0
			delay := 2

			expected := 20000

			// Act
			result, _ := service.CalculateSalaryAfterDelay(salaryBase, delay)

			// Assert
			if result != float64(expected) {
				t.Errorf("Expeted %v, received %v", expected, result)
			}
		})

		t.Run("Should calculate salary base after late when salary base is more than one decimal places", func(t *testing.T) {
			// Arrange
			service := cirt.NewService()
			salaryBase := 22000.45
			delay := 2

			expected := 20000.41

			// Act
			result, _ := service.CalculateSalaryAfterDelay(salaryBase, delay)

			// Assert
			if result != float64(expected) {
				t.Errorf("Expeted %v, received %v", expected, result)
			}
		})

		t.Run("Should not calculate salary base after late when salary base is negative", func(t *testing.T) {
			// Arrange
			service := cirt.NewService()
			salaryBase := -22000.45
			delay := 2

			// Act
			_, err := service.CalculateSalaryAfterDelay(salaryBase, delay)

			// Assert
			if err != cirt.ErrSalaryBaseNegative {
				t.Error(err)
			}
		})

	})
	t.Run("subsidys", func(t *testing.T) {
		t.Run("Should calculate subsidy", func(t *testing.T) {
			// Arrange
			service := cirt.NewService()
			subsidy := 1000.0
			days := 20

			expected := 20000

			// Act
			result, _ := service.CalculateSubsidy(subsidy, days)

			// Assert
			if result != float64(expected) {
				t.Errorf("Expected %v, received %v", expected, result)
			}
		})
		t.Run("Should calculate subsidy when default days is 22", func(t *testing.T) {
			// Arrange
			service := cirt.NewService()
			subsidy := 1000.0

			expected := 22000

			// Act
			result, _ := service.CalculateSubsidy(subsidy)

			// Assert
			if result != float64(expected) {
				t.Errorf("Expected %v, received %v", expected, result)
			}
		})
		t.Run("Should calculate subsidy with more than one decimal places", func(t *testing.T) {
			// Arrange
			service := cirt.NewService()
			subsidy := 1000.459

			expected := 22010.10

			// Act
			result, _ := service.CalculateSubsidy(subsidy)

			// Assert
			if result != float64(expected) {
				t.Errorf("Expected %v, received %v", expected, result)
			}
		})
		t.Run("Should not calculate subsidy when subsidy is negative", func(t *testing.T) {
			// Arrange
			service := cirt.NewService()
			subsidy := -1000.0

			// Act
			_, err := service.CalculateSubsidy(subsidy)

			// Assert
			if err != cirt.ErrSubsidyNegative {
				t.Error(err)
			}
		})
	})
	t.Run("should calculate excess", func(t *testing.T) {
		// Arrange
		service := cirt.NewService()
		subsidy := 2000.0

		expected := 14000

		result, _ := service.CalculateSubsidy(subsidy)

		// Act
		excess := service.Excess(result)

		// Assert
		if excess != float64(expected) {
			t.Errorf("Expected %v, received %v", expected, result)
		}
	})
	t.Run("Should not calculate excess when subsidy is less than 30000", func(t *testing.T) {
		// Arrange
		service := cirt.NewService()
		subsidy := 1000.0

		expected := 0

		result, _ := service.CalculateSubsidy(subsidy)

		// Act
		excess := service.Excess(result)

		// Assert
		if excess != float64(expected) {
			t.Errorf("Expected %v, received %v", expected, result)
		}
	})
	t.Run("Social Segurance", func(t *testing.T) {
		t.Run("Should calculate inss", func(t *testing.T) {
			// Arrange
			service := cirt.NewService()
			salaryBase := 20000.0
			subsidyFood := 1000.0
			subsidyTransPort := 1000.0
			premeo := 5000.0

			expected := 22150

			// Act
			result := service.CalculateSocialSegurance(salaryBase, subsidyFood, subsidyTransPort, premeo)

			// assert
			if result != float64(expected) {
				t.Errorf("Expected %v, received %v", expected, result)
			}
		})
		t.Run("Should calculate inss with more than one decimal places", func(t *testing.T) {
			// Arrange
			service := cirt.NewService()
			salaryBase := 20000.256
			subsidyFood := 1000.0
			subsidyTransPort := 1000.0
			premeo := 5000.0

			expected := 22150.26

			// Act
			result := service.CalculateSocialSegurance(salaryBase, subsidyFood, subsidyTransPort, premeo)

			// assert
			if result != float64(expected) {
				t.Errorf("Expected %v, received %v", expected, result)
			}
		})

	})
	t.Run("IRT", func(t *testing.T) {
		t.Run("Should calculate colletable material", func(t *testing.T) {
			// Arrange
			service := cirt.NewService()
			salarybase := 20000.0
			sujeicaoIrt := 16000.0
			inss := 20000.0

			expected := 16000.0

			// Act
			result := service.CalculateColletableMaterial(salarybase, sujeicaoIrt, inss)

			// Assert
			if result != float64(expected) {
				t.Errorf("Expected %v, received %v", expected, result)
			}
		})
		t.Run("Should calculate colletable material with more than one decimal places", func(t *testing.T) {
			// Arrange
			service := cirt.NewService()
			salarybase := 20000.564
			sujeicaoIrt := 16000.0
			inss := 20000.0

			expected := 16000.56

			// Act
			result := service.CalculateColletableMaterial(salarybase, sujeicaoIrt, inss)

			// Assert
			if result != float64(expected) {
				t.Errorf("Expected %v, received %v", expected, result)
			}
		})
		t.Run("Should calculate IRT", func(t *testing.T) {
			// Arrange
			service := cirt.NewService()
			salarybase := 20000.564
			sujeicaoIrt := 16000.0
			inss := 20000.0

			colletableMaterial := service.CalculateColletableMaterial(salarybase, sujeicaoIrt, inss)

			expected := 16000.56

			// Act
			result, err := service.CalculateIRT(colletableMaterial)

			// Assert
			if result != float64(expected) {
				t.Errorf("Expected %v, received %v", expected, result)
			}

			if err != nil {
				t.Error(err)
			}
		})
		t.Run("Should calculate discount total", func(t *testing.T) {
			// Arrange
			service := cirt.NewService()
			salaryBase := 20000.256
			subsidyFood := 1000.0
			subsidyTransPort := 1000.0
			premeo := 5000.0
			sujeicaoIrt := 16000.0

			inss := service.CalculateSocialSegurance(salaryBase, subsidyFood, subsidyTransPort, premeo)
			mc := service.CalculateColletableMaterial(salaryBase, sujeicaoIrt, inss)
			irt, _ := service.CalculateIRT(mc)

			expected := 36000.26

			// Act
			result := service.DiscountTotal(inss, irt)

			// Assert
			if result != float64(expected) {
				t.Errorf("Expected %v, received %v", expected, result)
			}
		})
		t.Run("Should calculate salary liquid", func(t *testing.T) {
			// Arrange
			service := cirt.NewService()
			salaryBase := 20000.256
			subsidyTransport := 1000.0
			subsidyFood := 1000.0
			premeo := 5000.0
			sujeicaoIrt := 0.0
			inss := service.CalculateSocialSegurance(salaryBase, subsidyFood, subsidyTransport, premeo)
			mc := service.CalculateColletableMaterial(salaryBase, sujeicaoIrt, inss)
			irt, _ := service.CalculateIRT(mc)
			discount := service.DiscountTotal(inss, irt)

			expected := 164752.25

			// Act
			result := service.CalculateSalaryLiquido(salaryBase, subsidyTransport, subsidyFood, premeo, discount)

			// Assert
			if result != float64(expected) {
				t.Errorf("Expected %v, received %v", expected, result)
			}
		})
	})
}
