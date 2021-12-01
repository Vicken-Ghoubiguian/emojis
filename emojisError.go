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
