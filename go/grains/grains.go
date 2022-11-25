package grains

import (
	"errors"
	"math"
)

func Square(number int) (uint64, error) {
	if number <= 0 || number > 64 {
		return 0, errors.New("must return cannot use zero or negative values")
	}
	return uint64(math.Pow(2, float64(number-1))), nil
}

func Total() uint64 {
	var sum uint64 = 0
	for i := 1; i <= 64; i++ {
		val, _ := Square(i)
		sum += val
	}
	return sum
}
