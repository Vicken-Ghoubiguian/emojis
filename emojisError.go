package emojis

//
type EmojisError struct {
	code    int
	status  string
	message string
}

//
func (ee *EmojisError) InitializeEmojisError(code int, status string, message string) {

	ee.code = code
	ee.status = status
	ee.message = message
}

// Defining the 'code' field getter...
func (ee *EmojisError) getCode() int {

	return ee.code
}

// Defining the 'status' field getter...
func (ee *EmojisError) getStatus() string {

	return ee.status
}

// Defining the 'message' field getter...
func (ee *EmojisError) getMessage() string {

	return ee.message
}
