package emojis

// =============================================================== For the 'OrderedPairAccordingKuratowski' type ===============================================================

// Defining the type 'OrderedPairAccordingKuratowski' which will define an ordered pair according Kuratowski made up of 2 emojis...
type OrderedPairAccordingKuratowski struct {
	firstEmoji  Emoji
	secondEmoji Emoji
}

// Defining the orderedPairAccordingKuratowski initializer...
func (opak *OrderedPairAccordingKuratowski) InitializeOrderedPairAccordingKuratowski(firstEmoji Emoji, secondEmoji Emoji) {

	opak.firstEmoji = firstEmoji
	opak.secondEmoji = secondEmoji
}

// Defining the 'firstEmoji' field getter...
func (opak *OrderedPairAccordingKuratowski) GetFirstEmoji() Emoji {

	return opak.firstEmoji
}

// Defining the 'secondEmoji' field getter...
func (opak *OrderedPairAccordingKuratowski) GetSecondEmoji() Emoji {

	return opak.secondEmoji
}

// Defining the 'ToString' function to return the opak 'OrderedPairAccordingKuratowski' as a string...
func (opak *OrderedPairAccordingKuratowski) ToString() string {

	return "(" + opak.firstEmoji.GetCharacter() + ", " + opak.secondEmoji.GetCharacter() + ")"
}

// =============================================================== For the 'CartesianProductEmojis' type ===============================================================

// Defining the type 'CartesianProductEmojis' which will define a cartesian product between two emojis set (formally 'ListOfEmojis')...
type CartesianProductEmojis struct {
	cartesianProductEmojisResult  []OrderedPairAccordingKuratowski
	cartesianProductEmojisAccount int
}

// Defining the 'cartesianProductEmojisAccount' field getter...
func (cpe *CartesianProductEmojis) GetCartesianProductEmojisAccount() int {

	//
	return cpe.cartesianProductEmojisAccount
}

//
func (cpe *CartesianProductEmojis) InitializeOrderedCartesianProductEmojis(firstLoe ListOfEmojis, secondLoe ListOfEmojis) {

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
