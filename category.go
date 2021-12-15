package emojis

// Defining the type 'Category' which define a category of emojis...
type Category struct {
	slug                 string
	subCategoriesArray   []string
	subCategoriesAccount int
}

// Defining the Category initializer...
func (catg *Category) InitializeCategory(slug string, subCategoriesArray []string) {

	catg.slug = slug
	catg.subCategoriesArray = subCategoriesArray
	catg.subCategoriesAccount = len(subCategoriesArray)
}

// Defining the 'slug' field getter...
func (catg *Category) GetSlug() string {

	return catg.slug
}

// Defining the 'subCategoriesArray' field getter...
func (catg *Category) GetSubCategoriesArray() []string {

	return catg.subCategoriesArray
}

// Defining the 'subCategoriesAccount' field getter...
func (catg *Category) GetSubCategoriesAccount() int {

	return catg.subCategoriesAccount
}

// Defining the 'GetPercentageInListOfEmojis' function to return the presence percentage of 'catg' emoji category in the 'loe' ListOfEmojis...
func (catg *Category) GetPercentageInListOfEmojis(loe ListOfEmojis) float64 {

	// Declaration of the 'emojisTotal' variable which will contain the total count of emojis in the 'loe' ListOfEmojis...
	var emojisTotal int

	// Declaration of the 'categorySCount' variable which will contain the count of emojis in the 'loe' ListOfEmojis of the specific category...
	var categorySCount int

	// Affectation of the 0 value to the 'categorySCount' variable...
	categorySCount = 0

	// Affectation of the total count of emojis in the 'loe' ListOfEmojis to the 'emojisTotal' variable...
	emojisTotal = catg.subCategoriesAccount

	// Initialization of the main loop of this function...
	for _, emoji := range loe.GetListOfEmojis() {

		//
		if catg.slug == emoji.GetGroup() {

			//
			categorySCount = categorySCount + 1
		}
	}

	//
	return (float64(categorySCount) * 100) / float64(emojisTotal)
}

// => YOU MUST DEVELOP SOME STATISTICAL CALCULATION FUNCTIONS FOR THE 'CATEGORY' TYPE...
