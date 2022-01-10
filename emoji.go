package emojis

import (
	"strconv"
	"strings"
)

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

	// Calculating and returning the emoji from the current emoji's codePoint...
	resultAsRune, _ := strconv.ParseInt(strings.TrimPrefix(ej.codePoint, "\\U"), 16, 32)

	// Returning finally the emoticon with some necessary operations...
	return string(rune(resultAsRune))
}

//
func (ej *Emoji) Equal(secondEmoji Emoji) bool {

	//
	return false
}

// Defining the function 'ToString' to display an emoji with all of its datas...
func (ej *Emoji) ToString() string {

	// Return the string with all datas of the current emojis...
	return "Character: " + ej.character + "\n" +
		"Slug: " + ej.slug + "\n" +
		"Unicode name: " + ej.unicodeName + "\n" +
		"Code point: " + ej.codePoint + "\n" +
		"Group: " + ej.group + "\n" +
		"Subgroup: " + ej.subGroup
}
