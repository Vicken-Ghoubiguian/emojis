package emojis

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
 * NEW 'EMOJI' STRUCTURE TO DEFINE AN EMOJI WITH ALL DATAS FROM THE OPEN EMOJI API WITH GETTERS (ALL ARE PUBLIC)...
 */

// Defining the type 'Emoji' which define an emoji
type Emoji struct {
	slug        string
	character   string
	unicodeName string
	codePoint   string
	group       string
	subGroup    string
}

// Defining the Emoji initializer
func (ej *Emoji) InitializeWeatherModule(slug string, character string, unicodeName string, codePoint string, group string, subGroup string) {

	ej.slug = slug
	ej.character = character
	ej.unicodeName = unicodeName
	ej.codePoint = codePoint
	ej.group = group
	ej.subGroup = subGroup
}

//
func (ej *Emoji) GetSlug() string {

	return ej.slug
}

//
func (ej *Emoji) GetCharacter() string {

	return ej.character
}

//
func (ej *Emoji) GetUnicodeName() string {

	return ej.unicodeName
}

//
func (ej *Emoji) GetCodePoint() string {

	return ej.codePoint
}

//
func (ej *Emoji) GetGroup() string {

	return ej.group
}

//
func (ej *Emoji) GetSubGroup() string {

	return ej.subGroup
}

//
func (ej *Emoji) GetEmoji() string {

	resultAsRune, _ := strconv.ParseInt(strings.TrimPrefix(ej.codePoint, "\\U"), 16, 32)

	return string(rune(resultAsRune))
}

/*
 * INTERNAL FUNCTIONS OF THE 'EMOJI' PACKAGE TO USE TO MAKE RUN THIS GO PACKAGE...
 */

// Function which display errors when they occurs
func errorHandlerFunction(err error) {

	// If the error is null (in this case, there is no error)
	if err != nil {

		// To display the error in red
		fmt.Println("\033[31m" + err.Error() + "\033[0m")

		// Exit the program (with exit code 1 (with error))
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

	var emojisWithAll map[string]Emoji

	return emojisWithAll
}

//
func GetASingleEmoji(unicodeName string, accessKey string) Emoji {

	var returnedEmoji Emoji

	return returnedEmoji
}

//
func GetSearchedForEmojis(searchedFor string, accessKey string) map[string]Emoji {

	var searchedForEmojis map[string]Emoji

	return searchedForEmojis
}
