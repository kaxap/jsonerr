// serialization errors
// error codes start from 2000

package jsonerr

import (
	"github.com/pkg/errors"
)

// code 2000
func NewTypeCastError(funcName, castTo string, doLog bool) error {
	err := errors.New("Error while casting to " + castTo)
	return newJsonError(2000, err, "", funcName, doLog)
}