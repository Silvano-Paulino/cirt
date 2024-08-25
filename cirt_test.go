package cirt_test

import (
	"testing"

	cirt "github.com/silvano-paulino"
)

func TestCirt(t *testing.T) {
	t.Run("Before calculate irt__should calculate salary base aflter late", func(t *testing.T) {

		t.Run("should calculate salary after late", func(t *testing.T) {
			// Arrange
			service := cirt.NewService()
			salaryBase := 20000.0
			days := 2

			// Act
			result := service.Cirt(salaryBase, days)

			// Assert
			if result == 0 {
				t.Errorf("Error: salary was not calculated")
			}
		})

		t.Run("When calaulate salary base after late__should calculate salary base per day", func(t *testing.T) {
			// Arrange
			service := cirt.NewService()
			salaryBase := 20000.0
			days := 2

			expeted := 10000

			// Act
			result := service.Cirt(salaryBase, days)

			// Assert
			if result != float64(expeted) {
				t.Errorf("Error: Expeted %v, received %v", expeted, result)
			}
		})

	})
}
