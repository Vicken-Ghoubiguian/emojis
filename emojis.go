package emojis

import (
	"strconv"
	"strings"
)

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
func fromCodePointToEmoji(codePointOfEmoji string) (string, error) {

	resultAsRune, err := strconv.ParseInt(strings.TrimPrefix(codePointOfEmoji, "\\U"), 16, 32)

	return string(rune(resultAsRune)), err
}

//
func GetAllDatasFromEmojis(accessKey string) map[string]Emoji {

	var emojisWithAll map[string]Emoji

	return emojisWithAll
}

//
func GetAllEmojis(accessKey string) map[string]string {

	//var brutEmojis map[string]Emoji

	var emojis map[string]string

	return emojis
}

//
func GetDatasFromOneParticularEmoji(unicodeName string, accessKey string) Emoji {

	var returnedEmoji Emoji

	return returnedEmoji
}

//
func GetOneParticularEmoji(unicodeName string, accessKey string) string {

	var brutEmoji Emoji

	emoji, _ := fromCodePointToEmoji(brutEmoji.codePoint)

	return emoji
}
