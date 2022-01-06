package emojis

import (
	"strings"
)

// Defining the type 'ListOfEmojis' which define a list of emojis...
type ListOfEmojis struct {
	mapOfEmojis   map[string]Emoji
	emojisAccount int
}

// Defining the ListOfEmojis initializer...
func (loe *ListOfEmojis) InitializeListOfEmojis(listOfEmojis map[string]Emoji) {

	loe.mapOfEmojis = listOfEmojis
	loe.emojisAccount = len(listOfEmojis)
}

// Defining the 'mapOfEmojis' field getter...
func (loe *ListOfEmojis) GetListOfEmojis() map[string]Emoji {

	return loe.mapOfEmojis
}

// Defining the 'emojisAccount' field getter...
func (loe *ListOfEmojis) GetEmojisAccount() int {

	return loe.emojisAccount
}

// Defining the 'GetEmojisFromGroup' function to get and return all emojis from a wished 'group' group in a 'ListOfEmojis' variable...
func (loe *ListOfEmojis) GetEmojisFromGroup(group string) ListOfEmojis {

	// Declaration of the 'returnedListOfEmojis' ListOfEmojis...
	var returnedListOfEmojis ListOfEmojis

	// Allocation of all necessary memory for the 'returnedEmojisMap' map[string]Emoji...
	returnedEmojisMap := make(map[string]Emoji)

	// Initialization of the main loop of this function...
	for slug, emoji := range loe.mapOfEmojis {

		// If the current emoji's group is equal to the wished 'group' group so...
		if emoji.GetGroup() == group {

			// Adding the emoji to the 'returnedEmojisMap' map[string]Emoji...
			returnedEmojisMap[slug] = emoji
		}
	}

	// Initialization of the 'returnedListOfEmojis' ListOfEmojis which will contain all existing emojis in the wished group...
	returnedListOfEmojis.InitializeListOfEmojis(returnedEmojisMap)

	// Returning the completed 'returnedListOfEmojis' ListOfEmojis which now contains all existing emojis in the wished group...
	return returnedListOfEmojis
}

// Defining the 'GetEmojisFromSubGroup' function to get and return all emojis from a wished 'subGroup' subgroup in a 'ListOfEmojis' variable...
func (loe *ListOfEmojis) GetEmojisFromSubGroup(subGroup string) ListOfEmojis {

	// Declaration of the 'returnedListOfEmojis' ListOfEmojis...
	var returnedListOfEmojis ListOfEmojis

	// Allocation of all necessary memory for the 'returnedEmojisMap' map[string]Emoji...
	returnedEmojisMap := make(map[string]Emoji)

	// Initialization of the main loop of this function...
	for slug, emoji := range loe.mapOfEmojis {

		// If the current emoji's subgroup is equal to the wished 'subGroup' subgroup so...
		if emoji.GetSubGroup() == subGroup {

			// Adding the emoji to the 'returnedEmojisMap' map[string]Emoji...
			returnedEmojisMap[slug] = emoji
		}
	}

	// Initialization of the 'returnedListOfEmojis' ListOfEmojis which will contain all existing emojis in the wished subgroup...
	returnedListOfEmojis.InitializeListOfEmojis(returnedEmojisMap)

	// Returning the completed 'returnedListOfEmojis' ListOfEmojis which now contains all existing emojis in the wished subgroup...
	return returnedListOfEmojis
}

// Defining the 'GetAllCategories' function to get all categories from the current 'loe' ListOfEmojis...
func (loe *ListOfEmojis) GetAllCategories() ListOfCategories {

	//
	var currentListOfCategories ListOfCategories

	//
	return currentListOfCategories
}

// Defining the 'GetASingleEmoji' function to return the emoji specified by its slug of the wished emoji of the current map[string]Emoji of 'loe' ListOfEmojis...
func (loe *ListOfEmojis) GetASingleEmoji(slug string) Emoji {

	// Declaration of the 'returnedEmoji' emoji...
	var returnedEmoji Emoji

	// In the case where the 'slug' key exists and have an Emoji...
	if foundedEmoji, currentEmoji := loe.mapOfEmojis[slug]; currentEmoji {

		// Initialization of the 'returnedEmoji' Emoji with the 'foundedEmoji' Emoji...
		returnedEmoji = foundedEmoji
	}

	// Returning the 'returnedEmoji' emoji...
	return returnedEmoji
}

// Defining the 'GetSearchedForEmojis' function to return a ListOfEmojis which contains all Emojis belonging to a subGroup specified by the 'searchedFor' variable in argument...
func (loe *ListOfEmojis) GetSearchedForEmojis(searchedFor string) ListOfEmojis {

	// Declaration of the 'searchedForEmojis' list of emojis...
	var searchedForEmojis ListOfEmojis

	// Definition of the 'mapOfSearchedForEmojis' map[string]Emoji which will contain all Emojis belonging to the whished subGroup...
	mapOfSearchedForEmojis := make(map[string]Emoji)

	// Initialization of the main loop of this function...
	for _, currentEmoji := range loe.mapOfEmojis {

		// In the case where the 'currentEmoji' slug contains the 'searchedFor' string...
		if strings.Contains(currentEmoji.GetSlug(), searchedFor) {

			// Put the current emoji to the 'mapOfSearchedForEmojis' map[string]Emoji with the current emoji's slug as the key..
			mapOfSearchedForEmojis[currentEmoji.GetSlug()] = currentEmoji
		}
	}

	// Initialization of the 'searchedForEmojis' ListOfEmojis to make it contain all existing emojis belonging to the whished subGroup...
	searchedForEmojis.InitializeListOfEmojis(mapOfSearchedForEmojis)

	// Returning the 'searchedForEmojis' ListOfEmojis...
	return searchedForEmojis
}

//
func (loe *ListOfEmojis) ToString() string {

	// Initialization of the main loop of this function...
	/*for _, emoji := range loe.mapOfEmojis {

	}*/

	//
	return ""
}

//
func (loe *ListOfEmojis) ToSlice() []string {

	// Initialization of the main loop of this function...
	/*for _, emoji := range loe.mapOfEmojis {

	}*/

	//
	return []string{"", "", "", "", "", ""}
}

//
/*func (loe *ListOfEmojis) ToSlice() []Emoji {

	// Initialization of the main loop of this function...
	/*for _, emoji := range loe.mapOfEmojis {

	}*/

	//
	//return []string{"", "", "", "", "", ""}
//}

// Defining a function named 'Add' to add an emoji in the 'loe' emojis list...
func (loe *ListOfEmojis) Add(emoji Emoji) {

	// Adding the 'emoji' Emoji in the 'loe' emojis list map[string]Emoji...
	loe.mapOfEmojis[emoji.GetSlug()] = emoji

	// Updating the 'loe' emojis list account...
	loe.emojisAccount = len(loe.mapOfEmojis)
}

// Defining a function named 'Delete' to delete an emoji in the 'loe' emojis list...
func (loe *ListOfEmojis) Delete(emoji Emoji) {

	// Deleting the 'emoji' Emoji in the 'loe' emojis list map[string]Emoji...
	delete(loe.mapOfEmojis, emoji.GetSlug())

	// Updating the 'loe' emojis list account...
	loe.emojisAccount = len(loe.mapOfEmojis)
}

// Defining a function named 'Contains' which check if the 'emoji' Emoji is in the 'loe' emojis list...
func (loe *ListOfEmojis) Contains(emoji Emoji) bool {

	// Initialization of the main loop of this function...
	for _, currentEmoji := range loe.mapOfEmojis {

		// In the case where the 'emoji' is equal to the 'currentEmoji' emoji in the 'loe' emojis list...
		if currentEmoji.GetSlug() == emoji.GetSlug() &&
			currentEmoji.GetCharacter() == emoji.GetCharacter() &&
			currentEmoji.GetUnicodeName() == emoji.GetUnicodeName() &&
			currentEmoji.GetCodePoint() == emoji.GetCodePoint() &&
			currentEmoji.GetGroup() == emoji.GetGroup() &&
			currentEmoji.GetSubGroup() == emoji.GetSubGroup() {

			// Returning 'true' in conclusion...
			return true
		}
	}

	// Returning 'false' by default...
	return false
}
