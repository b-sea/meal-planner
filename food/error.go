package food

import (
	"errors"
	"fmt"
)

// ErrUnitConversion is raised when a unit conversion fails.
var ErrUnitConversion = errors.New("unit conversion error")

func unitConversionError(v any) error {
	return fmt.Errorf("%w: %v", ErrUnitConversion, v)
}
