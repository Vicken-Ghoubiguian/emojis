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

//
func (ee *EmojisError) getCode() int {

	return ee.code
}

//
func (ee *EmojisError) getStatus() string {

	return ee.status
}

//
func (ee *EmojisError) getMessage() string {

	return ee.message
}
