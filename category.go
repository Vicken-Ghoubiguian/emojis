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
func (catg *Category) GetCategoriesArray() []string {

	return catg.subCategoriesArray
}

// Defining the 'subCategoriesAccount' field getter...
func (catg *Category) GetSubCategoriesAccount() int {

	return catg.subCategoriesAccount
}

// => YOU MUST DEVELOP SOME STATISTICAL CALCULATION FUNCTIONS FOR THE 'CATEGORY' TYPE...
