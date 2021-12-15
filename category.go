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

//
func (catg *Category) GetPercentageInListOfEmojis(loe ListOfEmojis) float64 {

	//
	var emojisTotal int

	//
	var categorySCount int

	//
	categorySCount = 0

	//
	emojisTotal = len(loe.GetListOfEmojis())

	//
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
