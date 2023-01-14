package main

import (
	"errors"
	"fmt"
)

// See types.go for the types defined for this exercise.

// TODO: Define the SillyNephewError type here.
type SillyNephewError struct {
	cows int
}

func (s *SillyNephewError) Error() string {
	return fmt.Sprintf("silly nephew, there cannot be %v cows", s.cows)
}

// DivideFood computes the fodder amount per cow for the given cows.
func DivideFood(weightFodder WeightFodder, cows int) (float64, error) {
	fodder, err := weightFodder.FodderAmount()
	switch {
	case fodder < 0 && (err == nil || err == ErrScaleMalfunction):
		err = errors.New("negative fodder")
	case cows == 0:
		err = errors.New("division by zero")
	case cows < 0:
		err = &SillyNephewError{cows: cows}
	case err == ErrScaleMalfunction:
		fodder, err = fodder*2, nil
	}
	if err != nil {
		return 0, err
	}
	return fodder / float64(cows), err
}
