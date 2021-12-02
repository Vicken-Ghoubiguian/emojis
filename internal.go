package emojis

import (
	"fmt"
	"os"
	"reflect"
)

// Function which display errors when they occurs...
func errorHandlerFunction(err error) {

	// If the error is null (in this case, there is no error)...
	if err != nil {

		// To display the error in red...
		fmt.Println("\033[31m" + err.Error() + "\033[0m")

		// Exit the program (with exit code 1 (with error))...
		os.Exit(1)
	}
}

// Function which convert a 'reflect.Value' variable to a '[]string' variable and which return it...
func reflectValueToStringArrayFunction(currentReflectValue reflect.Value) []string {

	//
	returnedArray := make([]string, currentReflectValue.Len())

	//
	for i := 0; i < currentReflectValue.Len(); i++ {

		//
		returnedArray[i] = fmt.Sprintf("%v", currentReflectValue.Index(i).Interface())
	}

	//
	return returnedArray
}
