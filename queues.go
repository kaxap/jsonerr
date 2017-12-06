// serialization errors
// error codes start from 5000

package jsonerr

import "fmt"

// code 5000
func NewQueueCreateError(funcName string, queueName string, serverAddr string, err error, doLog bool) error {
	if err == nil {
		return nil
	}

	return newJsonError(5000, err, fmt.Sprintf("Could not connect to server \"%s\", queue \"%s\": %s",
		serverAddr, queueName, err.Error()), funcName, doLog)
}

// code 5001
func NewQueueSendError(funcName string, queueName string, err error, doLog bool) error {
	if err == nil {
		return nil
	}
	return newJsonError(5001, err, fmt.Sprintf("Could not send to queue \"%s\": %s", queueName,
		err.Error()), funcName, doLog)
}

