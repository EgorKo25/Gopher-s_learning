package testing_theme

import (
	"math"
)

// Abs возвращает абсолютное значение.
// Например: 3.1 => 3.1, -3.14 => 3.14, -0 => 0.
// Покрыть тестами нужно эту функцию.
func Abs(value float64) float64 {
	return math.Abs(value)
}
