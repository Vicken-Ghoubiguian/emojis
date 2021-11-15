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
func (ej *Emoji) InitializeWeatherModule(slug string, character string, unicodeName string, codePoint string, group string, subGroup string) {

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

	resultAsRune, _ := strconv.ParseInt(strings.TrimPrefix(ej.codePoint, "\\U"), 16, 32)

	return string(rune(resultAsRune))
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

	resultAsRune, error := strconv.ParseInt(strings.TrimPrefix(codePointOfEmoji, "\\U"), 16, 32)

	errorHandlerFunction(error)

	return string(rune(resultAsRune))
}

//
func GetAllEmojis(accessKey string) map[string]Emoji {

	// Declaration of the 'emojiSInterface' interface...
	var emojisInterface []interface{}

	//
	var allEmojis map[string]Emoji

	//
	getEmojiFromTheOpenEmojiAPIRequest := "https://emoji-api.com/emojis?access_key=" + accessKey

	//
	getEmojiFromTheOpenEmojiAPIAPIResp, err := http.Get(getEmojiFromTheOpenEmojiAPIRequest)

	// Manage the possible occured error...
	errorHandlerFunction(err)

	//
	getEmojiFromTheOpenEmojiAPIJsonString, err := ioutil.ReadAll(getEmojiFromTheOpenEmojiAPIAPIResp.Body)

	// Manage the possible occured error...
	errorHandlerFunction(err)

	//
	err = json.Unmarshal(getEmojiFromTheOpenEmojiAPIJsonString, &emojisInterface)

	// Manage the possible occured error...
	errorHandlerFunction(err)

	//
	fmt.Println(emojisInterface[25])

	//
	return allEmojis
}

// Function to take the 'unicodeName' of the wished emoji and your personal 'accessKey' of the Open Emoji API as arguments...
func GetASingleEmoji(unicodeName string, accessKey string) Emoji {

	// Declaration of the 'emojiSInterface' interface...
	var emojiSInterface []interface{}

	// Declaration of the 'emojiSMap' map[string]interface{}...
	var emojiSMap map[string]interface{}

	// Declaration of the 'returnedEmoji' emoji...
	var returnedEmoji Emoji

	// Definition of the HTTPS request's URL to get the wished emoji from the Open Emoji API...
	getEmojiFromTheOpenEmojiAPIRequest := "https://emoji-api.com/emojis/" + unicodeName + "?access_key=" + accessKey

	// Execution of the Get HTTPS request to get all emojis from the Open Emoji API...
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
	returnedEmoji.InitializeWeatherModule(fmt.Sprintf("%v", emojiSMap["slug"]), fmt.Sprintf("%v", emojiSMap["character"]), fmt.Sprintf("%v", emojiSMap["unicodeName"]), fmt.Sprintf("%v", emojiSMap["codePoint"]), fmt.Sprintf("%v", emojiSMap["group"]), fmt.Sprintf("%v", emojiSMap["subGroup"]))

	// Returning the 'returnedEmoji' initialized emoji...
	return returnedEmoji
}

//
func GetSearchedForEmojis(searchedFor string, accessKey string) map[string]Emoji {

	var searchedForEmojis map[string]Emoji

	return searchedForEmojis
}

// Function to retrieve all emojis by a given category...
func GetInCategoryEmojis(category string, accessKey string) map[string]Emoji {

	var inCategoryEmojis map[string]Emoji

	return inCategoryEmojis
}
