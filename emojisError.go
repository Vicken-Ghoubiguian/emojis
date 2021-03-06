package emojis

// Defining the type 'EmojisError' which define an occured error in this package from the Open Emoji API...
type EmojisError struct {
	code    int
	status  string
	message string
}

// Defining the EmojisError initializer...
func (ee *EmojisError) InitializeEmojisError(code int, status string, message string) {

	ee.code = code
	ee.status = status
	ee.message = message
}

// Defining the 'code' field getter...
func (ee *EmojisError) GetCode() int {

	return ee.code
}

// Defining the 'status' field getter...
func (ee *EmojisError) GetStatus() string {

	return ee.status
}

// Defining the 'message' field getter...
func (ee *EmojisError) GetMessage() string {

	return ee.message
}
