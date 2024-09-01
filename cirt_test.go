package main_test

import (
	"testing"

	cirt "github.com/silvano-paulino"
)

func TestCirt(t *testing.T) {
	t.Run("Salary base after late", func(t *testing.T) {

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
			result, _ := service.CalculoSubsiodio(subsidio, days)

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
			result, _ := service.CalculoSubsiodio(subsidio)

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
			result, _ := service.CalculoSubsiodio(subsidio)

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
			_, err := service.CalculoSubsiodio(subsidio)

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

		result, _ := service.CalculoSubsiodio(subsidio)

		// Act
		execesso := service.Excesso(result)

		// Assert
		if execesso != float64(expected) {
			t.Errorf("Expected %v, received %v", expected, result)
		}
	})

	t.Run("Should not calculate excesso when subsidio is less than 30000", func(t *testing.T) {
		// Arrange
		service := cirt.NewService()
		subsidio := 1000.0

		expected := 0

		result, _ := service.CalculoSubsiodio(subsidio)

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

	t.Run("Imposto sobre Rendimento de Trabalho (IRT)", func(t *testing.T) {
		t.Run("Should calculate materia colectavel", func(t *testing.T) {
			// Arrange
			service := cirt.NewService()
			salarybase := 20000.0
			sujeicaoIrt := 16000.0
			inss := 20000.0

			expected := 16000.0

			// Act
			result := service.CalculateMateriaColectavel(salarybase, sujeicaoIrt, inss)

			// Assert
			if result != float64(expected) {
				t.Errorf("Expected %v, received %v", expected, result)
			}
		})

		t.Run("Should calculate materia colectavel with more than one decimal places", func(t *testing.T) {
			// Arrange
			service := cirt.NewService()
			salarybase := 20000.564
			sujeicaoIrt := 16000.0
			inss := 20000.0

			expected := 16000.56

			// Act
			result := service.CalculateMateriaColectavel(salarybase, sujeicaoIrt, inss)

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

			mc := service.CalculateMateriaColectavel(salarybase, sujeicaoIrt, inss)

			expected := 16000.56

			// Act
			result, err := service.CalculateIRT(mc)

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
			subsidioAlimentacao := 1000.0
			subsidioTransPorte := 1000.0
			premeo := 5000.0
			sujeicaoIrt := 16000.0

			inss := service.CalculateSocialSegurance(salaryBase, subsidioAlimentacao, subsidioTransPorte, premeo)
			mc := service.CalculateMateriaColectavel(salaryBase, sujeicaoIrt, inss)
			irt, _ := service.CalculateIRT(mc)

			expected := 36000.26

			// Act
			result := service.DiscountTotal(inss, irt)

			// Assert
			if result != float64(expected) {
				t.Errorf("Expected %v, received %v", expected, result)
			}
		})

		t.Run("Should calculate salary liquido", func(t *testing.T) {
			// Arrange
			service := cirt.NewService()
			salaryBase := 20000.256
			subSidiT := 1000.0
			subSidiA := 1000.0
			premeo := 5000.0
			sujeicaoIrt := 0.0
			inss := service.CalculateSocialSegurance(salaryBase, subSidiA, subSidiT, premeo)
			mc := service.CalculateMateriaColectavel(salaryBase, sujeicaoIrt, inss)
			irt, _ := service.CalculateIRT(mc)
			discount := service.DiscountTotal(inss, irt)

			expected := 164752.25

			// Act
			result := service.CalculateSalaryLiquido(salaryBase, subSidiT, subSidiA, premeo, discount)

			// Assert
			if result != float64(expected) {
				t.Errorf("Expected %v, received %v", expected, result)
			}
		})
	})
}
