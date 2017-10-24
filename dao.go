// DAO errors
// error codes start from 3000

package jsonerr

//code 3000
func NewGeneralDAOError(funcName string, err error, doLog bool) error {
	return newJsonError(3000, err, "General DAO error: " + err.Error(), funcName, doLog)
}

//code 3001
func NewUpdateDAOError(funcName string, err error, doLog bool) error {
	return newJsonError(3001, err, "Update DAO error: " + err.Error(), funcName, doLog)
}

//code 3002
func NewCreateDAOError(funcName string, err error, doLog bool) error {
	return newJsonError(3002, err, "Create DAO error: " + err.Error(), funcName, doLog)
}

//code 3003
func NewDeleteDAOError(funcName string, err error, doLog bool) error {
	return newJsonError(3003, err, "Delete DAO error: " + err.Error(), funcName, doLog)
}
