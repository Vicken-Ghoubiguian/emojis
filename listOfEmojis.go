package emojis

import (
	"sort"
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
	for slug, emoji := range loe.GetListOfEmojis() {

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
	for slug, emoji := range loe.GetListOfEmojis() {

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

// Defining the 'GetASingleEmoji' to return the emoji specified by its slug of the wished emoji of the current map[string]Emoji of 'loe' ListOfEmojis...
func (loe *ListOfEmojis) GetASingleEmoji(slug string) Emoji {

	// Declaration of the 'returnedEmoji' emoji...
	var returnedEmoji Emoji

	// In the case where the 'slug' key exists and have an Emoji...
	if foundedEmoji, currentEmoji := loe.GetListOfEmojis()[slug]; currentEmoji {

		// Initialization of the 'returnedEmoji' Emoji with the 'foundedEmoji' Emoji...
		returnedEmoji = foundedEmoji
	}

	// Returning the 'returnedEmoji' emoji...
	return returnedEmoji
}

// Defining the 'GetSearchedForEmojis' to return a ListOfEmojis which contains all Emojis belonging to a subGroup specified by the 'searchedFor' variable in argument...
func (loe *ListOfEmojis) GetSearchedForEmojis(searchedFor string) ListOfEmojis {

	// Declaration of the 'searchedForEmojis' list of emojis...
	var searchedForEmojis ListOfEmojis

	// Definition of the 'mapOfSearchedForEmojis' map[string]Emoji which will contain all Emojis belonging to the whished subGroup...
	mapOfSearchedForEmojis := make(map[string]Emoji)

	// Initialization of the main loop of this function...
	for _, currentEmoji := range loe.GetListOfEmojis() {

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

// Defining a function named 'SortAlphabetically' to sort 'loe' emojis list in alphabetical order ...
func (loe *ListOfEmojis) SortAlphabetically() {

	// Definition of the 'newMapOfEmojis' map[string]Emoji to contain all the emojis of 'loe' ListOfEmojis sorted in alphabetical order ...
	var newMapOfEmojis map[string]Emoji

	// Allocation of all necessary memory for the 'slugs' []string...
	slugs := make([]string, 0, len(loe.mapOfEmojis))

	// Initialization of the first loop of this function...
	for s := range loe.mapOfEmojis {

		// Put the 's' string (current slug) in the 'slugs' []string...
		slugs = append(slugs, s)
	}

	// Sort the slice containing the slugs in alphabetical order...
	sort.Strings(slugs)

	// Allocation of all necessary memory for the 'newMapOfEmojis' map[string]Emoji...
	newMapOfEmojis = make(map[string]Emoji)

	// Initialization of the second loop of this function...
	for s := range slugs {

		// Put the current emoji in the
		newMapOfEmojis[slugs[s]] = loe.mapOfEmojis[slugs[s]]
	}

	// Assignment of the 'newMapOfEmojis' new map[string]Emoji to the 'mapOfEmojis' ListOfEmojis field...
	loe.mapOfEmojis = newMapOfEmojis
}

//
func (loe *ListOfEmojis) SortReverseAlphabetically() {

	//sort.Sort(sort.Reverse(sort.StringSlice(slugs)))
}

// => YOU MUST DEFINE AND DEVELOP SOME SORT FUNCTIONS FOR THE 'LISTOFEMOJIS' TYPE...

// => YOU MUST DEFINE AND DEVELOP SOME STATISTICAL CALCULATION FUNCTIONS FOR THE 'LISTOFEMOJIS' TYPE...
