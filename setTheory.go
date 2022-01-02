package emojis

// Function to determine the intersection between the 'loe' ListOfEmojis and the 'currentListOfEmojis' ListOfEmojis...
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

// Function to determine the union of the 'loe' ListOfEmojis with the 'currentListOfEmojis' ListOfEmojis...
func (loe *ListOfEmojis) Union(currentListOfEmojis ListOfEmojis) ListOfEmojis {

	//
	var unionEmojisMap map[string]Emoji

	//
	var unionEmojis ListOfEmojis

	//
	unionEmojisMap = make(map[string]Emoji)

	//
	for _, emoji := range loe.GetListOfEmojis() {

		//
		unionEmojisMap[emoji.GetSlug()] = emoji
	}

	//
	for _, emoji := range currentListOfEmojis.GetListOfEmojis() {

		//
		if !loe.Contains(emoji) {

			//
			unionEmojisMap[emoji.GetSlug()] = emoji
		}
	}

	//
	unionEmojis.InitializeListOfEmojis(unionEmojisMap)

	//
	return unionEmojis
}

//
func (loe *ListOfEmojis) Complement(currentListOfEmojis ListOfEmojis) ListOfEmojis {

	var complementEmojis ListOfEmojis

	return complementEmojis
}

//
func (loe *ListOfEmojis) CartesianProduct() ListOfEmojis {

	var cartesianProductEmojis ListOfEmojis

	return cartesianProductEmojis
}

// => YOU MUST DEFINE AND DEVELOP SOME FUNCTIONS TO MAKE OPERATIONS OF SET THEORY IN LIST OF EMOJIS...
