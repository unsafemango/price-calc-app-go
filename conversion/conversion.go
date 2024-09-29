package conversion

import (
	"errors"
	"strconv"
)

func StringsToFloats(strings []string) ([]float64, error) {

	var floats []float64

	for _, stringVal := range strings {
		floatPrice, err := strconv.ParseFloat(stringVal, 64) // converts text to other value types

		if err != nil {
			return nil, errors.New("failed to conver the string to float")
		}

		floats = append(floats, floatPrice)

	}

	return floats, nil
}
