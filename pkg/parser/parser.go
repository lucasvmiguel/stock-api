// parser package contains functions for converting types
package parser

import (
	"strconv"

	"github.com/pkg/errors"
)

// converts number as string in uint
func StringToUint(number string) (uint, error) {
	num, err := strconv.ParseUint(number, 10, 64)
	if err != nil {
		return 0, errors.Wrap(err, "invalid number")
	}

	return uint(num), nil
}
