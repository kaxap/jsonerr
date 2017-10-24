// serialization errors
// error codes start from 1000

package jsonerr

// code 1000
func NewJsonDecodeError(funcName string, err error, doLog bool) error {
	return newJsonError(1000, err, "JSON decode error: " + err.Error(), funcName, doLog)
}

// code 1001
func NewJsonEncodeError(funcName string, err error, doLog bool) error {
	return newJsonError(1001, err, "JSON encode error: " + err.Error(), funcName, doLog)
}

// code 1002
func NewFormatError(funcName string, err error, doLog bool) error {
	return newJsonError(1002, err, "", funcName, doLog)
}