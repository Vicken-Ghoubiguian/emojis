package emojis

// =============================================================== For the 'CartesianProductEmojis' type ===============================================================

//
type orderedPairAccordingKuratowski struct {
	firstEmoji  Emoji
	secondEmoji Emoji
}

//
type CartesianProductEmojis struct {
}

// =============================================================== For all the set theory functions ===============================================================

// Function to determine the intersection between the 'loe' ListOfEmojis and the 'currentListOfEmojis' ListOfEmojis...
func (loe *ListOfEmojis) Intersection(currentListOfEmojis ListOfEmojis) ListOfEmojis {

	// Declaration of the 'intersectionEmojisMap' map[string]Emoji...
	var intersectionEmojisMap map[string]Emoji

	// Declaration of the 'intersectionEmojis' ListOfEmojis...
	var intersectionEmojis ListOfEmojis

	// Definition of the 'intersectionEmojisMap' map[string]Emoji which will contain all emojis from the intersection of the 'loe' ListOfEmojis and the 'currentListOfEmojis' ListOfEmojis...
	intersectionEmojisMap = make(map[string]Emoji)

	// Initialization of the main loop of this function...
	for _, emoji := range loe.mapOfEmojis {

		// If the 'currentListOfEmojis' ListOfEmojis does not contains the 'emoji' Emoji...
		if currentListOfEmojis.Contains(emoji) {

			// Put the emoji inside the 'intersectionEmojisMap' map[string]Emoji...
			intersectionEmojisMap[emoji.GetSlug()] = emoji
		}
	}

	// Initialization of the 'intersectionEmojis' ListOfEmojis to make it contain all emojis from the intersection of the 'loe' ListOfEmojis and the 'currentListOfEmojis' ListOfEmojis...
	intersectionEmojis.InitializeListOfEmojis(intersectionEmojisMap)

	// Returning the 'intersectionEmojis' ListOfEmojis...
	return intersectionEmojis
}

// Function to determine the union of the 'loe' ListOfEmojis with the 'currentListOfEmojis' ListOfEmojis...
func (loe *ListOfEmojis) Union(currentListOfEmojis ListOfEmojis) ListOfEmojis {

	// Declaration of the 'unionEmojisMap' map[string]Emoji...
	var unionEmojisMap map[string]Emoji

	// Declaration of the 'unionEmojis' ListOfEmojis...
	var unionEmojis ListOfEmojis

	// Definition of the 'intersectionEmojisMap' map[string]Emoji which will contain all emojis in the 'loe' ListOfEmojis with all emojis in the 'currentListOfEmojis' ListOfEmojis...
	unionEmojisMap = make(map[string]Emoji)

	// Initialization of the first main loop of this function...
	for _, emoji := range loe.mapOfEmojis {

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

	// Initialization of the 'unionEmojis' ListOfEmojis to make it contain all emojis in the 'loe' ListOfEmojis with all emojis in the 'currentListOfEmojis' ListOfEmojis...
	unionEmojis.InitializeListOfEmojis(unionEmojisMap)

	// Returning the 'unionEmojis' ListOfEmojis...
	return unionEmojis
}

//
func (loe *ListOfEmojis) Complement(currentListOfEmojis ListOfEmojis) ListOfEmojis {

	var complementEmojis ListOfEmojis

	return complementEmojis
}

//
func (loe *ListOfEmojis) CartesianProduct(currentListOfEmojis ListOfEmojis) ListOfEmojis {

	//
	var cartesianProductEmojisMap map[string]Emoji

	//
	var cartesianProductEmojis ListOfEmojis

	/**/

	//
	cartesianProductEmojis.InitializeListOfEmojis(cartesianProductEmojisMap)

	//
	return cartesianProductEmojis
}

// => YOU MUST DEFINE AND DEVELOP SOME FUNCTIONS TO MAKE OPERATIONS OF SET THEORY IN LIST OF EMOJIS...
