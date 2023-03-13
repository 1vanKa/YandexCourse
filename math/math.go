package math

import "errors"

func Add(a, b int) (int, error) {
	if a == 0 || b == 0 {
		return 0, errors.New("arg is zero")
	}

	if a < 0 || b < 0 {
		return 0, errors.New("arg is negative")
	}
	return a + b, nil
}

func EstimateValue(value int) string {
	switch {
	case value < 10:
		return "small"

	case value < 100:
		return "medium"

	default:
		return "big"
	}
}
