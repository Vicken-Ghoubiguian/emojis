package emojis

// Defining the type 'ListOfCategories' which define a list of emojis categories...
type ListOfCategories struct {
	mapOfCategories   map[string]Category
	categoriesAccount int
}

// Defining the ListOfCategories initializer...
func (loc *ListOfCategories) InitializeListOfCategories(listOfCategories map[string]Category) {

	loc.mapOfCategories = listOfCategories
	loc.categoriesAccount = len(listOfCategories)
}

// Defining the 'mapOfCategories' field getter...
func (loc *ListOfCategories) GetMapOfCategories() map[string]Category {

	return loc.mapOfCategories
}

// Defining the 'categoriesAccount' field getter...
func (loc *ListOfCategories) GetCategoriesAccount() int {

	return loc.categoriesAccount
}

//
func (loc *ListOfCategories) GetPercentageInListOfEmojis() int {

	//

	//
	return 0
}

// => YOU MUST DEVELOP SOME STATISTICAL CALCULATION FUNCTIONS FOR THE LISTOF'CATEGORIES' TYPE...
