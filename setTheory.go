package emojis

// Function to determine the intersection between the 'loe' ListOfEmojis and the 'currentListOfEmojis' ListOfEmojis...
func (loe *ListOfEmojis) Intersection(currentListOfEmojis ListOfEmojis) ListOfEmojis {

	//
	var intersectionEmojisMap map[string]Emoji

	//
	var intersectionEmojis ListOfEmojis

	// Definition of the 'intersectionEmojisMap' map[string]Emoji which will contain all Emojis from the intersection of the 'loe' ListOfEmojis and the 'currentListOfEmojis' ListOfEmojis...
	intersectionEmojisMap = make(map[string]Emoji)

	// Initialization of the main loop of this function...
	for _, emoji := range loe.GetListOfEmojis() {

		// If the 'currentListOfEmojis' ListOfEmojis does not contains the 'emoji' Emoji...
		if currentListOfEmojis.Contains(emoji) {

			// Put the emoji inside the 'intersectionEmojisMap' map[string]Emoji...
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

	// Initialization of the first main loop of this function...
	for _, emoji := range loe.GetListOfEmojis() {

		// Put the emoji inside the 'unionEmojisMap' map[string]Emoji...
		unionEmojisMap[emoji.GetSlug()] = emoji
	}

	// Initialization of the second main loop of this function...
	for _, emoji := range currentListOfEmojis.GetListOfEmojis() {

		// If the 'loe' ListOfEmojis does not contains the 'emoji' Emoji...
		if !loe.Contains(emoji) {

			// Put the emoji inside the 'unionEmojisMap' map[string]Emoji...
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
