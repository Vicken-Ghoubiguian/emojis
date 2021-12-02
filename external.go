package emojis

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
	"strconv"
	"strings"
)

//
func GetAllCategories(accessKey string) ListOfCategories {

	//
	var categoriesInterface []interface{}

	//
	var allCategoriesMap map[string]Category

	//
	var allCategories ListOfCategories

	//
	var categoriesInterfaceLen int

	// Declaration of the 'currentCategory' Category...
	var currentCategory Category

	var subCategoriesArray []string

	//
	getEmojisCategoriesFromTheOpenEmojiAPIRequest := "https://emoji-api.com/categories?access_key=" + accessKey

	//
	getEmojisCategoriesFromTheOpenEmojiAPIResp, err := http.Get(getEmojisCategoriesFromTheOpenEmojiAPIRequest)

	//
	errorHandlerFunction(err)

	//
	getEmojisCategoriesFromTheOpenEmojiAPIJsonString, err := ioutil.ReadAll(getEmojisCategoriesFromTheOpenEmojiAPIResp.Body)

	//
	errorHandlerFunction(err)

	//
	err = json.Unmarshal(getEmojisCategoriesFromTheOpenEmojiAPIJsonString, &categoriesInterface)

	//
	errorHandlerFunction(err)

	//
	categoriesInterfaceLen = len(categoriesInterface)

	// Allocation of all necessary memory for the 'allCategoriesMap' map[string]Category...
	allCategoriesMap = make(map[string]Category)

	//
	for i := 0; i < categoriesInterfaceLen; i++ {

		//
		currentCategoryAsMap := categoriesInterface[i].(map[string]interface{})

		lines := reflect.ValueOf(currentCategoryAsMap["subCategories"])

		fmt.Println("\n")

		fmt.Println(lines)

		lines2 := lines.Index(0)

		fmt.Println("\n")

		fmt.Println(lines2)

		fmt.Println("\n")

		fmt.Println(reflectValueToStringArrayFunction(lines))

		fmt.Println("\n")

		//
		currentCategory.InitializeCategory(fmt.Sprintf("%v", currentCategoryAsMap["slug"]), subCategoriesArray)

		//
		allCategoriesMap[fmt.Sprintf("%v", currentCategoryAsMap["slug"])] = currentCategory
	}

	//
	allCategories.InitializeListOfCategories(allCategoriesMap)

	//
	return allCategories
}

// Function to return the emoji calculated from the 'codePointOfEmoji' codePoint...
func GetEmojiFromCodePoint(codePointOfEmoji string) string {

	// Calculating and returning the emoji from the 'codePointOfEmoji' codePoint...
	resultAsRune, error := strconv.ParseInt(strings.TrimPrefix(codePointOfEmoji, "\\U"), 16, 32)

	// Manage the possible occured error...
	errorHandlerFunction(error)

	// Returning finally the emoticon with some necessary operations...
	return string(rune(resultAsRune))
}

// Function to return all emojis by taken your personal 'accessKey' of the Open Emoji API as argument...
func GetAllEmojis(accessKey string) ListOfEmojis {

	// Declaration of the 'emojiSInterface' interface...
	var emojisInterface []interface{}

	// Declaration of the 'allEmojisMap' map[string]Emoji...
	var allEmojisMap map[string]Emoji

	// Declaration of the 'allEmojis' list of emojis...
	var allEmojis ListOfEmojis

	// Declaration of the 'emojisInterfaceLen' variable which contains the length of the 'emojisInterface' interface...
	var emojisInterfaceLen int

	// Declaration of the 'currentEmoji' emoji...
	var currentEmoji Emoji

	// Definition of the HTTPS request's URL to get all emojis from the Open Emoji API...
	getEmojiFromTheOpenEmojiAPIRequest := "https://emoji-api.com/emojis?access_key=" + accessKey

	// Execution of the Get HTTPS request to get all existing emojis from the Open Emoji API...
	getEmojiFromTheOpenEmojiAPIResp, err := http.Get(getEmojiFromTheOpenEmojiAPIRequest)

	// Manage the possible occured error...
	errorHandlerFunction(err)

	// Read the HTTP response's body in the 'getEmojiFromTheOpenEmojiAPIJsonString' string variable...
	getEmojiFromTheOpenEmojiAPIJsonString, err := ioutil.ReadAll(getEmojiFromTheOpenEmojiAPIResp.Body)

	// Manage the possible occured error...
	errorHandlerFunction(err)

	// Unmarshall all of received datas of all received emojis in the 'emojisInterface' interface...
	err = json.Unmarshal(getEmojiFromTheOpenEmojiAPIJsonString, &emojisInterface)

	// Manage the possible occured error...
	errorHandlerFunction(err)

	// Initialization of the 'emojisInterfaceLen' variable with the 'emojisInterface' interface length...
	emojisInterfaceLen = len(emojisInterface)

	// Allocation of all necessary memory for the 'allEmojisMap' map[string]Emoji...
	allEmojisMap = make(map[string]Emoji)

	// Initialization of the main loop of this function...
	for i := 0; i < emojisInterfaceLen; i++ {

		// Conversion of the 'emojisInterface' interface to a map[string]interface{}...
		currentEmojiAsMap := emojisInterface[i].(map[string]interface{})

		// Initialization of the 'currentEmoji' emoji with the corresponding collected datas...
		currentEmoji.InitializeEmoji(fmt.Sprintf("%v", currentEmojiAsMap["slug"]), fmt.Sprintf("%v", currentEmojiAsMap["character"]), fmt.Sprintf("%v", currentEmojiAsMap["unicodeName"]), fmt.Sprintf("%v", currentEmojiAsMap["codePoint"]), fmt.Sprintf("%v", currentEmojiAsMap["group"]), fmt.Sprintf("%v", currentEmojiAsMap["subGroup"]))

		// Adding the 'currentEmoji' emoji in the 'allEmojisMap' map[string]Emoji...
		allEmojisMap[fmt.Sprintf("%v", currentEmojiAsMap["slug"])] = currentEmoji
	}

	// Initialization of the 'allEmojis' ListOfEmojis to make it contain all existing emojis...
	allEmojis.InitializeListOfEmojis(allEmojisMap)

	// Returning the completed 'allEmojisMap' ListOfEmojis which now contains all existing emojis...
	return allEmojis
}

// Function to return the emoji specified by its slug by taken the 'unicodeName' of the wished emoji and your personal 'accessKey' of the Open Emoji API as arguments...
func GetASingleEmoji(slug string, accessKey string) Emoji {

	// Declaration of the 'emojiSInterface' interface...
	var emojiSInterface []interface{}

	// Declaration of the 'emojiSMap' map[string]interface{}...
	var emojiSMap map[string]interface{}

	// Declaration of the 'returnedEmoji' emoji...
	var returnedEmoji Emoji

	// Definition of the HTTPS request's URL to get the wished emoji from the Open Emoji API...
	getEmojiFromTheOpenEmojiAPIRequest := "https://emoji-api.com/emojis/" + slug + "?access_key=" + accessKey

	// Execution of the Get HTTPS request to get the wished emoji from the Open Emoji API...
	getEmojiFromTheOpenEmojiAPIResp, err := http.Get(getEmojiFromTheOpenEmojiAPIRequest)

	// Manage the possible occured error...
	errorHandlerFunction(err)

	// Take the body of the previous Get HTTPS request's response in the 'getEmojiFromTheOpenEmojiAPIResp' variable...
	getEmojiFromTheOpenEmojiAPIJsonString, err := ioutil.ReadAll(getEmojiFromTheOpenEmojiAPIResp.Body)

	// Manage the possible occured error...
	errorHandlerFunction(err)

	// Unmarshall all of received datas in the 'emojiSInterface' interface...
	err = json.Unmarshal(getEmojiFromTheOpenEmojiAPIJsonString, &emojiSInterface)

	// Manage the possible occured error...
	errorHandlerFunction(err)

	// Conversion of the first element of the 'emojiSInterface' interface to a map[string]interface{}...
	emojiSMap = emojiSInterface[0].(map[string]interface{})

	// Initialization of the 'returnedEmoji' current emoji...
	returnedEmoji.InitializeEmoji(fmt.Sprintf("%v", emojiSMap["slug"]), fmt.Sprintf("%v", emojiSMap["character"]), fmt.Sprintf("%v", emojiSMap["unicodeName"]), fmt.Sprintf("%v", emojiSMap["codePoint"]), fmt.Sprintf("%v", emojiSMap["group"]), fmt.Sprintf("%v", emojiSMap["subGroup"]))

	// Returning the 'returnedEmoji' initialized emoji...
	return returnedEmoji
}

// Function to return all emojis from a wished 'searchedFor' subGroup of emojis as a ListOfEmojis...
func GetSearchedForEmojis(searchedFor string, accessKey string) ListOfEmojis {

	// Declaration of the 'searchedForEmojisInterface' interface...
	var searchedForEmojisInterface []interface{}

	// Declaration of the 'searchedForEmojisMap' map[string]Emoji...
	var searchedForEmojisMap map[string]Emoji

	// Declaration of the 'searchedForEmojis' list of emojis...
	var searchedForEmojis ListOfEmojis

	// Declaration of the 'searchedForEmojisInterfaceLen' variable which contains the length of the 'searchedForEmojisInterface' interface...
	var searchedForEmojisInterfaceLen int

	// Declaration of the 'currentEmoji' emoji...
	var currentEmoji Emoji

	// Definition of the HTTPS request's URL to get all interesting emojis from the Open Emoji API...
	getSearchedForEmojisFromTheOpenEmojiAPIRequest := "https://emoji-api.com/emojis?search=" + searchedFor + "&access_key=" + accessKey

	// Execution of the Get HTTPS request to get all interesting emojis from the Open Emoji API...
	getSearchedForEmojisFromTheOpenEmojiAPIResp, err := http.Get(getSearchedForEmojisFromTheOpenEmojiAPIRequest)

	// Manage the possible occured error...
	errorHandlerFunction(err)

	// Take the body of the previous Get HTTPS request's response in the 'getSearchedForEmojisFromTheOpenEmojiAPIResp' variable...
	getSearchedForEmojisFromTheOpenEmojiAPIJsonString, err := ioutil.ReadAll(getSearchedForEmojisFromTheOpenEmojiAPIResp.Body)

	// Manage the possible occured error...
	errorHandlerFunction(err)

	// Unmarshall all of received datas in the 'searchedForEmojisInterface' interface...
	err = json.Unmarshal(getSearchedForEmojisFromTheOpenEmojiAPIJsonString, &searchedForEmojisInterface)

	// Manage the possible occured error...
	errorHandlerFunction(err)

	// Initialization of the 'searchedForEmojisInterfaceLen' variable with the 'searchedForEmojisInterface' interface length...
	searchedForEmojisInterfaceLen = len(searchedForEmojisInterface)

	// Allocation of all necessary memory for the 'searchedForEmojisMap' map[string]Emoji...
	searchedForEmojisMap = make(map[string]Emoji)

	// Initialization of the main loop of this function...
	for i := 0; i < searchedForEmojisInterfaceLen; i++ {

		// Conversion of the 'searchedForEmojisInterface' interface to a map[string]interface{}...
		currentEmojiAsMap := searchedForEmojisInterface[i].(map[string]interface{})

		// Initialization of the 'currentEmoji' emoji with the corresponding collected datas...
		currentEmoji.InitializeEmoji(fmt.Sprintf("%v", currentEmojiAsMap["slug"]), fmt.Sprintf("%v", currentEmojiAsMap["character"]), fmt.Sprintf("%v", currentEmojiAsMap["unicodeName"]), fmt.Sprintf("%v", currentEmojiAsMap["codePoint"]), fmt.Sprintf("%v", currentEmojiAsMap["group"]), fmt.Sprintf("%v", currentEmojiAsMap["subGroup"]))

		//
		searchedForEmojisMap[fmt.Sprintf("%v", currentEmojiAsMap["slug"])] = currentEmoji
	}

	// Initialization of the 'searchedForEmojis' current ListOfEmojis...
	searchedForEmojis.InitializeListOfEmojis(searchedForEmojisMap)

	// Returning the completed 'searchedForEmojis' ListOfEmojis which now contains all interesting emojis...
	return searchedForEmojis
}

// Function to return all emojis from a wished 'category' category of emojis as a ListOfEmojis...
func GetInCategoryEmojis(category string, accessKey string) ListOfEmojis {

	// Declaration of the 'inCategoryEmojisInterface' interface...
	var inCategoryEmojisInterface []interface{}

	// Declaration of the 'inCategoryEmojisMap' map[string]Emoji...
	var inCategoryEmojisMap map[string]Emoji

	// Declaration of the 'inCategoryEmojis' list of emojis...
	var inCategoryEmojis ListOfEmojis

	// Declaration of the 'inCategoryEmojisInterfaceLen' variable which contains the length of the 'inCategoryEmojisInterface' interface...
	var inCategoryEmojisInterfaceLen int

	// Declaration of the 'currentEmoji' emoji...
	var currentEmoji Emoji

	// Definition of the HTTPS request's URL to get the wished emoji in the wished category from the Open Emoji API...
	getEmojisFromCategoryFromTheOpenEmojiAPIRequest := "https://emoji-api.com/categories/" + category + "?access_key=" + accessKey

	// Execution of the Get HTTPS request to get all existing emojis in the wished category from the Open Emoji API...
	getEmojisFromCategoryFromTheOpenEmojiAPIResp, err := http.Get(getEmojisFromCategoryFromTheOpenEmojiAPIRequest)

	// Manage the possible occured error...
	errorHandlerFunction(err)

	// Read the HTTP response's body in the 'getEmojisFromCategoryFromTheOpenEmojiAPIJsonString' string variable...
	getEmojisFromCategoryFromTheOpenEmojiAPIJsonString, err := ioutil.ReadAll(getEmojisFromCategoryFromTheOpenEmojiAPIResp.Body)

	// Manage the possible occured error...
	errorHandlerFunction(err)

	// Unmarshall all of received datas of all received emojis belonging to the wished category in the 'inCategoryEmojisInterface' interface...
	err = json.Unmarshal(getEmojisFromCategoryFromTheOpenEmojiAPIJsonString, &inCategoryEmojisInterface)

	// Manage the possible occured error...
	errorHandlerFunction(err)

	// Initialization of the 'inCategoryEmojisInterfaceLen' variable with the 'inCategoryEmojisInterface' interface length...
	inCategoryEmojisInterfaceLen = len(inCategoryEmojisInterface)

	// Allocation of all necessary memory for the 'inCategoryEmojisMap' map[string]Emoji...
	inCategoryEmojisMap = make(map[string]Emoji)

	// Initialization of the main loop of this function...
	for i := 0; i < inCategoryEmojisInterfaceLen; i++ {

		// Conversion of the 'inCategoryEmojisInterface' interface to a map[string]interface{}...
		currentEmojiAsMap := inCategoryEmojisInterface[i].(map[string]interface{})

		// Initialization of the 'currentEmoji' emoji with the corresponding collected datas...
		currentEmoji.InitializeEmoji(fmt.Sprintf("%v", currentEmojiAsMap["slug"]), fmt.Sprintf("%v", currentEmojiAsMap["character"]), fmt.Sprintf("%v", currentEmojiAsMap["unicodeName"]), fmt.Sprintf("%v", currentEmojiAsMap["codePoint"]), fmt.Sprintf("%v", currentEmojiAsMap["group"]), fmt.Sprintf("%v", currentEmojiAsMap["subGroup"]))

		// Adding the 'currentEmoji' emoji in the 'inCategoryEmojisMap' map[string]Emoji...
		inCategoryEmojisMap[fmt.Sprintf("%v", currentEmojiAsMap["slug"])] = currentEmoji
	}

	// Initialization of the 'inCategoryEmojis' ListOfEmojis to make it contain all existing emojis belonging to the wished category...
	inCategoryEmojis.InitializeListOfEmojis(inCategoryEmojisMap)

	// Returning the completed 'inCategoryEmojis' ListOfEmojis which now contains all existing emojis in the wished category...
	return inCategoryEmojis
}

// Defining the 'GetEmojisFromSubGroup' function to get and return all emojis from a wished 'subGroup' subgroup in a 'ListOfEmojis' variable...
func GetEmojisFromSubGroup(subGroup string, accessKey string) ListOfEmojis {

	// Declaration of the 'emojiSInterface' interface...
	var emojisInterface []interface{}

	// Declaration of the 'fromSubGroupEmojisMap' map[string]Emoji...
	var fromSubGroupEmojisMap map[string]Emoji

	// Declaration of the 'inCategoryEmojis' list of emojis...
	var fromSubGroupEmojis ListOfEmojis

	// Declaration of the 'emojisInterfaceLen' variable which contains the length of the 'emojisInterface' interface...
	var emojisInterfaceLen int

	// Declaration of the 'currentEmoji' emoji...
	var currentEmoji Emoji

	// Definition of the HTTPS request's URL to get all emojis from the Open Emoji API...
	getEmojiFromTheOpenEmojiAPIRequest := "https://emoji-api.com/emojis?access_key=" + accessKey

	// Execution of the Get HTTPS request to get all existing emojis from the Open Emoji API...
	getEmojiFromTheOpenEmojiAPIResp, err := http.Get(getEmojiFromTheOpenEmojiAPIRequest)

	// Manage the possible occured error...
	errorHandlerFunction(err)

	// Read the HTTP response's body in the 'getEmojiFromTheOpenEmojiAPIJsonString' string variable...
	getEmojiFromTheOpenEmojiAPIJsonString, err := ioutil.ReadAll(getEmojiFromTheOpenEmojiAPIResp.Body)

	// Manage the possible occured error...
	errorHandlerFunction(err)

	// Unmarshall all of received datas of all received emojis in the 'emojisInterface' interface...
	err = json.Unmarshal(getEmojiFromTheOpenEmojiAPIJsonString, &emojisInterface)

	// Manage the possible occured error...
	errorHandlerFunction(err)

	// Initialization of the 'emojisInterfaceLen' variable with the 'emojisInterface' interface length...
	emojisInterfaceLen = len(emojisInterface)

	// Allocation of all necessary memory for the 'fromSubGroupEmojisMap' map[string]Emoji...
	fromSubGroupEmojisMap = make(map[string]Emoji)

	// Initialization of the main loop of this function...
	for i := 0; i < emojisInterfaceLen; i++ {

		// Conversion of the 'emojisInterface' interface to a map[string]interface{}...
		currentEmojiAsMap := emojisInterface[i].(map[string]interface{})

		// If the current emoji's sub group is equal to the wished 'subGroup' sub group so...
		if fmt.Sprintf("%v", currentEmojiAsMap["subGroup"]) == subGroup {

			// Initialization of the 'currentEmoji' emoji with the corresponding collected datas...
			currentEmoji.InitializeEmoji(fmt.Sprintf("%v", currentEmojiAsMap["slug"]), fmt.Sprintf("%v", currentEmojiAsMap["character"]), fmt.Sprintf("%v", currentEmojiAsMap["unicodeName"]), fmt.Sprintf("%v", currentEmojiAsMap["codePoint"]), fmt.Sprintf("%v", currentEmojiAsMap["group"]), fmt.Sprintf("%v", currentEmojiAsMap["subGroup"]))

			// Adding the 'currentEmoji' emoji in the 'fromSubGroupEmojisMap' map[string]Emoji...
			fromSubGroupEmojisMap[fmt.Sprintf("%v", currentEmojiAsMap["slug"])] = currentEmoji
		}
	}

	// Initialization of the 'fromSubGroupEmojis' ListOfEmojis to make it contain all emojis from the wished sub group...
	fromSubGroupEmojis.InitializeListOfEmojis(fromSubGroupEmojisMap)

	// Returning the completed 'fromSubGroupEmojis' ListOfEmojis which now contains all emojis from the wished sub group...
	return fromSubGroupEmojis
}
