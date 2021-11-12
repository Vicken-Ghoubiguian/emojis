package emoji

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
func (ej *Emoji) getSlug() string {

	return ej.slug
}

//
func (ej *Emoji) getCharacter() string {

	return ej.character
}

//
func (ej *Emoji) getUnicodeName() string {

	return ej.unicodeName
}

//
func (ej *Emoji) getCodePoint() string {

	return ej.codePoint
}

//
func (ej *Emoji) getGroup() string {

	return ej.group
}

//
func (ej *Emoji) getSubGroup() string {

	return ej.subGroup
}
