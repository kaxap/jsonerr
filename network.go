// network errors
// error codes start from 4000

package jsonerr

//code 4000
func NewGeneralNetworkError(funcName string, err error, doLog bool) error {
    if err == nil {
        return nil
    }
    return newJsonError(4000, err, "General DAO error: " + err.Error(), funcName, doLog)
}
