package emojis

import (
	"fmt"
	"sort"
)

// Defining a function named 'SortAlphabetically' to sort 'loe' emojis list in alphabetical order...
func (loe *ListOfEmojis) SortAlphabetically() {

	// Allocation of all necessary memory for the 'slugs' []string...
	slugs := make([]string, 0, len(loe.mapOfEmojis))

	// Initialization of the first loop of this function...
	for s := range loe.mapOfEmojis {

		// Put the 's' string (current slug) in the 'slugs' []string...
		slugs = append(slugs, s)
	}

	// Sort the slice containing the slugs in alphabetical order...
	sort.Strings(slugs)

	// Initialization of the second loop of this function...
	for s := range slugs {

		// Display the concerned emoji...
		fmt.Println(loe.mapOfEmojis[slugs[s]])
	}
}

// Defining a function named 'SortReverseAlphabetically' to sort 'loe' emojis list in reverse alphabetical order...
func (loe *ListOfEmojis) SortReverseAlphabetically() {

	// Allocation of all necessary memory for the 'slugs' []string...
	slugs := make([]string, 0, len(loe.mapOfEmojis))

	// Initialization of the first loop of this function...
	for s := range loe.mapOfEmojis {

		// Put the 's' string (current slug) in the 'slugs' []string...
		slugs = append(slugs, s)
	}

	// Sort the slice containing the slugs in reverse alphabetical order...
	sort.Sort(sort.Reverse(sort.StringSlice(slugs)))

	// Initialization of the second loop of this function...
	for s := range slugs {

		// Display the concerned emoji...
		fmt.Println(loe.mapOfEmojis[slugs[s]])
	}
}

// => YOU MUST DEFINE AND DEVELOP SOME SORT FUNCTIONS FOR THE 'LISTOFEMOJIS' TYPE...

// => YOU MUST DEVELOP SOME STATISTICAL CALCULATION FUNCTIONS FOR THE LISTOF'CATEGORIES' TYPE...
