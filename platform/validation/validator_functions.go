package validation

import (
	"errors"
	"fmt"
	"strconv"
)

func required(fieldName string, value interface{}, arg string) (valid bool, err error) {
	if str, ok := value.(string); ok {
		valid = str != ""
		err = fmt.Errorf("a value is required")
	} else {
		err = errors.New("the required validator is for strings")
	}

	return
}

func min(fieldName string, value interface{}, arg string) (valid bool, err error) {
	minVal, err := strconv.Atoi(arg)

	if err != nil {
		panic("Invalid arguments for validator: " + arg)
	}

	err = fmt.Errorf("the minimum value is %v", minVal)

	if iVal, iValOk := value.(int); iValOk {
		valid = iVal >= minVal
	} else if fVal, fValOk := value.(float64); fValOk {
		valid = fVal >= float64(minVal)
	} else if strVal, strValOk := value.(string); strValOk {
		err = fmt.Errorf("the minimum length is %v characters", minVal)
		valid = len(strVal) >= minVal
	} else {
		err = errors.New("the min validator is for int, float64, and str values")
	}

	return
}
