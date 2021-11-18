package emojis

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
)

/*
 * NEW 'EMOJI' STRUCTURE TO DEFINE AN EMOJI WITH ALL DATAS FROM THE OPEN EMOJI API WITH GETTERS (ALL ARE PUBLIC)...
 */

// Defining the type 'Emoji' which define an emoji...
type Emoji struct {
	slug        string
	character   string
	unicodeName string
	codePoint   string
	group       string
	subGroup    string
}

// Defining the Emoji initializer...
func (ej *Emoji) InitializeEmoji(slug string, character string, unicodeName string, codePoint string, group string, subGroup string) {

	ej.slug = slug
	ej.character = character
	ej.unicodeName = unicodeName
	ej.codePoint = codePoint
	ej.group = group
	ej.subGroup = subGroup
}

// Defining the 'slug' field getter...
func (ej *Emoji) GetSlug() string {

	return ej.slug
}

// Defining the 'character' field getter...
func (ej *Emoji) GetCharacter() string {

	return ej.character
}

// Defining the 'unicodeName' field getter...
func (ej *Emoji) GetUnicodeName() string {

	return ej.unicodeName
}

// Defining the 'codePoint' field getter...
func (ej *Emoji) GetCodePoint() string {

	return ej.codePoint
}

// Defining the 'group' field getter...
func (ej *Emoji) GetGroup() string {

	return ej.group
}

// Defining the 'subGroup' field getter...
func (ej *Emoji) GetSubGroup() string {

	return ej.subGroup
}

// Defining the emoticon getter...
func (ej *Emoji) GetEmoji() string {

	//
	resultAsRune, _ := strconv.ParseInt(strings.TrimPrefix(ej.codePoint, "\\U"), 16, 32)

	return string(rune(resultAsRune))
}

/*
 * NEW 'LISTOFEMOJIS' STRUCTURE TO DEFINE AN EMOJI'S MAP WITH ALL RELATED DATAS CONTAINING ALL EMOJIS WITH ALL OF THEIR DATAS FROM THE OPEN EMOJI API...
 */

// Defining the type 'ListOfEmojis' which define a list of emojis...
type ListOfEmojis struct {
	mapOfEmojis   map[string]Emoji
	emojisAccount int
}

// Defining the ListOfEmojis initializer...
func (loe *ListOfEmojis) InitializeListOfEmojis(listOfEmojis map[string]Emoji) {

	loe.mapOfEmojis = listOfEmojis
	loe.emojisAccount = len(listOfEmojis)
}

// Defining the 'mapOfEmojis' field getter...
func (loe *ListOfEmojis) GetListOfEmojis() map[string]Emoji {

	return loe.mapOfEmojis
}

// Defining the 'emojisAccount' field getter...
func (loe *ListOfEmojis) GetEmojisAccount() int {

	return loe.emojisAccount
}

//
func (loe *ListOfEmojis) GetEmojisFromGroup(group string) ListOfEmojis {

	//
	var returnedListOfEmojis ListOfEmojis

	//
	returnedEmojisMap := make(map[string]Emoji)

	//
	for slug, emoji := range loe.GetListOfEmojis() {

		//
		if emoji.GetGroup() == group {

			//
			returnedEmojisMap[slug] = emoji
		}
	}

	//
	returnedListOfEmojis.InitializeListOfEmojis(returnedEmojisMap)

	//
	return returnedListOfEmojis
}

//
func (loe *ListOfEmojis) GetEmojisFromSubGroup(subGroup string) ListOfEmojis {

	//
	var returnedListOfEmojis ListOfEmojis

	//
	returnedEmojisMap := make(map[string]Emoji)

	//
	for slug, emoji := range loe.GetListOfEmojis() {

		//
		if emoji.GetSubGroup() == subGroup {

			//
			returnedEmojisMap[slug] = emoji
		}
	}

	//
	returnedListOfEmojis.InitializeListOfEmojis(returnedEmojisMap)

	//
	return returnedListOfEmojis
}

/*
 * INTERNAL FUNCTIONS OF THE 'EMOJI' PACKAGE TO USE TO MAKE RUN THIS GO PACKAGE...
 */

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

/*
 * PUBLIC FUNCTIONS OF THE 'EMOJI' PACKAGE TO USE IN EVERY GO PROJECTS...
 */

//
func GetEmojiFromCodePoint(codePointOfEmoji string) string {

	//
	resultAsRune, error := strconv.ParseInt(strings.TrimPrefix(codePointOfEmoji, "\\U"), 16, 32)

	//
	errorHandlerFunction(error)

	//
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
	getEmojiFromTheOpenEmojiAPIAPIResp, err := http.Get(getEmojiFromTheOpenEmojiAPIRequest)

	// Manage the possible occured error...
	errorHandlerFunction(err)

	// Read the HTTP response's body in the 'getEmojiFromTheOpenEmojiAPIJsonString' string variable...
	getEmojiFromTheOpenEmojiAPIJsonString, err := ioutil.ReadAll(getEmojiFromTheOpenEmojiAPIAPIResp.Body)

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

//
func GetAllCathegories() {

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
	getEmojiFromTheOpenEmojiAPIAPIResp, err := http.Get(getEmojiFromTheOpenEmojiAPIRequest)

	// Manage the possible occured error...
	errorHandlerFunction(err)

	// Take the body of the previous Get HTTPS request's response in the 'getEmojiFromTheOpenEmojiAPIAPIResp' variable...
	getEmojiFromTheOpenEmojiAPIJsonString, err := ioutil.ReadAll(getEmojiFromTheOpenEmojiAPIAPIResp.Body)

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

//
func GetSearchedForEmojis(searchedFor string, accessKey string) ListOfEmojis {

	var searchedForEmojis ListOfEmojis

	return searchedForEmojis
}

//
func GetInCategoryEmojis(category string, accessKey string) ListOfEmojis {

	var inCategoryEmojis ListOfEmojis

	return inCategoryEmojis
}
