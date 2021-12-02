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

	// Declaration and initialization of the 'returnedArray' array...
	returnedArray := make([]string, currentReflectValue.Len())

	// Initialization of the main loop of this function...
	for i := 0; i < currentReflectValue.Len(); i++ {

		// Put the element in the interface inside the 'returnedArray' []string array...
		returnedArray[i] = fmt.Sprintf("%v", currentReflectValue.Index(i).Interface())
	}

	// Returning the completed 'returnedArray' []string array...
	return returnedArray
}
