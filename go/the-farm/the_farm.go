package thefarm

import (
	"errors"
	"fmt"
)

// See types.go for the types defined for this exercise.

// TODO: Define the SillyNephewError type here.
type SillyNephewError struct {
	message string
	cows    int
}

func (e *SillyNephewError) Error() string {
	return fmt.Sprintf("%s %d cows", e.message, e.cows)
}

// DivideFood computes the fodder amount per cow for the given cows.
func DivideFood(weightFodder WeightFodder, cows int) (float64, error) {

	totalFodder, err := weightFodder.FodderAmount()

	if err != nil {
		if err == ErrScaleMalfunction {
			totalFodder *= 2
			if totalFodder < 0 {
				return 0, errors.New("negative fodder")
			}
			return totalFodder / float64(cows), nil
		}
		return 0, err
	}
	if totalFodder < 0 {
		return 0, errors.New("negative fodder")
	}
	if cows == 0 {
		return 0, errors.New("division by zero")
	}
	if cows <= 0 {
		err := &SillyNephewError{message: "silly nephew, there cannot be", cows: cows}
		return 0, err
	}

	fodder := totalFodder / float64(cows)
	return fodder, err
}
