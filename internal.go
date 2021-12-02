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

//
func reflectValueToStringArrayFunction(currentReflectValue reflect.Value) []string {

	//
	returnedArray := make([]string, currentReflectValue.Len())

	//
	for i := 0; i < currentReflectValue.Len(); i++ {

		fmt.Println("\n")

		fmt.Println(currentReflectValue.Index(i))

		test := currentReflectValue.Index(i)

		fmt.Println("\n")

		fmt.Println(fmt.Sprintf("%v", test.Interface()))

		//
		returnedArray[i] = fmt.Sprintf("%v", test.Interface())
	}

	//
	return returnedArray
}
