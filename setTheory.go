package emojis

//
func (loe *ListOfEmojis) Intersection(currentListOfEmojis ListOfEmojis) ListOfEmojis {

	//
	var intersectionEmojisMap map[string]Emoji

	//
	var intersectionEmojis ListOfEmojis

	//
	intersectionEmojisMap = make(map[string]Emoji)

	//
	for _, emoji := range loe.GetListOfEmojis() {

		//
		if currentListOfEmojis.Contains(emoji) {

			//
			intersectionEmojisMap[emoji.GetSlug()] = emoji
		}
	}

	//
	intersectionEmojis.InitializeListOfEmojis(intersectionEmojisMap)

	//
	return intersectionEmojis
}

//
/*func (loe *ListOfEmojis) Union(currentListOfEmojis ListOfEmojis) ListOfEmojis {

	//
	var unionEmojisMap map[string]Emoji

	//
	var unionEmojis ListOfEmojis

	//
	unionEmojisMap = make(map[string]Emoji)

	//
	for _, emoji := range loe.GetListOfEmojis() {

		//
		if currentListOfEmojis.Contains(emoji) && {


		}
	}

	return unionEmojis
}*/

//
func (loe *ListOfEmojis) Complement(currentListOfEmojis ListOfEmojis) ListOfEmojis {

	var complementEmojis ListOfEmojis

	return complementEmojis
}

//
/*func (loe *ListOfEmojis) CartesianProduct() ListOfEmojis {

	return true
}*/

// => YOU MUST DEFINE AND DEVELOP SOME FUNCTIONS TO MAKE OPERATIONS OF SET THEORY IN LIST OF EMOJIS...
