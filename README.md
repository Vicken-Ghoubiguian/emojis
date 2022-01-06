# emojis

Package written in Go to get, return, treat and manage emojis ⛑ 🪱 💥 🏵️ 🇳🇺 💾 🏟️ 🫁 💦 🦭🐽 🔱 🦋 📌 👹 ♨️ 🇳🇷 📐 🤬 🐍 🐈 ⚜️ 🔫 🚄 ❤️‍🩹 🌋 🎓 💈 🎃 🥑 😈 🥥 😸 🐾 🥵 ♻️ 🦾 🙀 🎗️ ☠️ 🫀 👻 🇰🇵 🫐 🗾 🏗️ 🤢 🔕 🏧 👽 🪔 🌹 🪙 ⚕️🛎️ 💩 🐡 🐙 😻 🧳 🎥 💻 🥶 🔘 🖥️ 🔮 🐳 🕳️ 💯 💫 🪐 🧨 🗺️ 🇨🇺 🦣 🦈 🐲 🧱 🍳 🇨🇼 ⛩️ 🤯 🌡️ 🥩 👺 📎 💳 🦧 🎳 🧶 🏛️ 🧠 🥀 🛑 ✔️ 🐬 🇮🇲 🔥 from the [Open Emoji API](https://emoji-api.com) in every Go project...

## Contents

* [Introduction](#introduction)
* [Presentation](#presentation)
* [Project's structure](#project_s_structure)
* [How was this project developed ?](#how_was_this_project_developed)
* [How to use this package ?](#how_to_use_this_package)
* [Little examples](#little_examples)
	* [Example 1 - Get and display the data of all emojis](#example_1)
	* [Example 2 - Get and display the data of one specific emoji](#example_2)
	* [Example 3 - Get and display the data of all emojis from a specific category (1)](#example_3)
	* [Example 4 - Get and display the data of all emojis from a specific sub group](#example_4)
	* [Example 5 - Get and display the data of all emojis from a specific category (2)](#example_5)
	* [Example 6 - Get and display the data of all emojis attached to a particular string (1)](#example_6)
	* [Example 7 - ](#example_7)
	* [Example 8 - Get and display the data of all emojis attached to a particular string (2)](#example_8)
	* [Example 9](#example_9)
	* [Example 10](#example_10)
	* [Example 11](#example_11)
	* [Example 12](#example_12)
	* [Example 13 - Get a list of all emojis and display the corresponding percentage for each category](#example_13)
	* [Example 14](#example_14)
	* [Example 15 - Get all datas of all emojis and use them in a displayed text](#example_15)
	* [Example 16 - Delete an emoji from a list of emojis](#example_16)
	* [Example 17 - Add an emoji to a list of emojis](#example_17)
	* [Example 18](#example_18)
	* [Example 19](#example_19)
	* [Example 20 - Check if an emoji belong to a list of emojis](#example_20)
* [Where to use it ?](#where_to_use_it)
* [Useful links](#useful_links)
* [Conclusion](#conclusion)

<a name="introduction"></a>
## Introduction

<a name="presentation"></a>
## Presentation

So this golang package allows you to do the following things with emojis:

* collect all the existing emojis,
* collect all emojis belonging to a specific category,
* collect all emojis belonging to a specific sub-category,
* get the emoji from its code point (),
* get all the existing emoji categories with the corresponding sub-categories,
* do different kinds of sorting on emojis (),
* do various statistical calculations on emojis ().

<a name="project_s_structure"></a>
## Project's structure

<a name="how_was_this_project_developed"></a>
## How was this project developed ?

<a name="how_to_use_this_package"></a>
## How to use this package ?

The `emojis` Go package can be used in every Go project you develop. However, you can use it only in Go projects.
To do it, you must execute this command to install the `emojis` golang package:

```bash
go get -u github.com/Vicken-Ghoubiguian/emojis
```

This command will download, from its GitHub repository, and install, on your local machine, the `emojis` Go package. You can use it in any directory and Go file in your file system tree...
To know how to use this Go package, please take a look at the examples below...

<a name="little_examples"></a>
## Little examples

Below are the examples that show how to use, in details, all the functions of this package:

<a name="example_1"></a>
### Example 1 - Get and display the data of all emojis...

```go
package main

import (
	
	"fmt"
	"github.com/Vicken-Ghoubiguian/emojis"
)

// Definition of the main function...
func main() {

	// Declaration and initialization of the 'allEmojisListOfEmojis' variable which will contain all existing emojis...
	allEmojisListOfEmojis := emojis.GetAllEmojis("<Your_access_key>")

	// Display a separator...
	fmt.Println("\n==================================\n")

	// Defining the loop to display datas of all existing emojis...
	for _, emoji := range allEmojisListOfEmojis.GetListOfEmojis() {

		// Display the current emoji...
		fmt.Println(emoji.ToString())

		// Break one line...
		fmt.Println("\n")
	}

	// Display a separator...
	fmt.Println("\n==================================\n")
}
```

<a name="example_2"></a>
### Example 2 - Get and display the data of one specific emoji...

```go
package main

import (
	
	"fmt"
	"github.com/Vicken-Ghoubiguian/emojis"
)

// Definition of the main function...
func main() {

	// Declaration and initialization of the 'anatomicalHeartEmoji' variable which will contain the 'anatomical heart' emoji...
	anatomicalHeartEmoji := emojis.GetASingleEmoji("anatomical-heart", "<Your_access_key>")

	// Display a separator...
    fmt.Println("\n==================================\n")

	// Display the 'anatomical heart' emoji...
	fmt.Println(anatomicalHeartEmoji.ToString())

	// Display a separator...
    fmt.Println("\n==================================\n")
}
```

<a name="example_3"></a>
### Example 3 - Get and display the data of all emojis from a specific category (1)...

```go
package main

import (

        "fmt"
        "github.com/Vicken-Ghoubiguian/emojis"
)

// Definition of the main function...
func main() {

        // Declaration and initialization of the 'allEmojisListOfEmojis' variable which will contain all existing emojis...
        allEmojisListOfEmojis := emojis.GetAllEmojis("<Your_access_key>")

        // Declaration and initialization of the 'allSmileysEmojisListOfEmojis' variable which will contain all existing emojis of the 'smileys-emotion' category...
        allSmileysEmojisListOfEmojis := allEmojisListOfEmojis.GetEmojisFromGroup("smileys-emotion")

		// Display a separator...
		fmt.Println("\n==================================\n")

		// Defining the loop to display datas of all existing emojis of the 'smileys-emotion' category...
		for _, emoji := range allSmileysEmojisListOfEmojis.GetListOfEmojis() {

			// Display the current emoji...
			fmt.Println(emoji.ToString())

			// Break one line...
			fmt.Println("\n")
		}

		// Display a separator...
		fmt.Println("\n==================================\n")
}
```

<a name="example_4"></a>
### Example 4 - Get and display the data of all emojis from a specific sub group...

```go
package main

import (

        "fmt"
        "github.com/Vicken-Ghoubiguian/emojis"
)

// Definition of the main function...
func main() {

	// Declaration and initialization of the 'allEmojisListOfEmojis' variable which will contain all existing emojis...
    allEmojisListOfEmojis := emojis.GetAllEmojis("<Your_access_key>")

    // Declaration and initialization of the 'allCatFacesEmojisListOfEmojis' variable which will contain all existing emojis of the 'cat-face' sub category...
    allCatFacesEmojisListOfEmojis := allEmojisListOfEmojis.GetEmojisFromSubGroup("cat-face")

	// Display a separator...
	fmt.Println("\n==================================\n")

	// Defining the loop to display datas of all existing emojis of the 'cat-face' sub category...
	for _, emoji := range allCatFacesEmojisListOfEmojis.GetListOfEmojis() {

		// Display the current emoji...
		fmt.Println(emoji.ToString())

		// Break one line...
		fmt.Println("\n")
	}

	// Display a separator...
	fmt.Println("\n==================================\n")
}
```

<a name="example_5"></a>
### Example 5 - Get and display the data of all emojis from a specific category (2)...

```go
package main

import (

        "fmt"
        "github.com/Vicken-Ghoubiguian/emojis"
)

// Definition of the main function...
func main() {

    // Declaration and initialization of the 'allSmileysEmojisListOfEmojis' variable which will contain all existing emojis of the 'smileys-emotion' category...
    allSmileysEmojisListOfEmojis := emojis.GetInCategoryEmojis("smileys-emotion", "<Your_access_key>")

	// Display a separator...
	fmt.Println("\n==================================\n")

	// Defining the loop to display datas of all existing emojis of the 'smileys-emotion' category...
	for _, emoji := range allSmileysEmojisListOfEmojis.GetListOfEmojis() {

			// Display the current emoji...
			fmt.Println(emoji.ToString())

			// Break one line...
			fmt.Println("\n")
		}

		// Display a separator...
		fmt.Println("\n==================================\n")
}
```

<a name="example_6"></a>
### Example 6 - Get and display the data of all emojis attached to a particular string...

```go
package main

import (

        "fmt"
        "github.com/Vicken-Ghoubiguian/emojis"
)

// Definition of the main function...
func main() {

	// Declaration and initialization of the 'allSearchForCatListOfEmojis' variable which will contain all existing emojis all existing emojis attached to the string 'cat'...
    allSearchForCatListOfEmojis := emojis.GetSearchedForEmojis("cat", "<Your_access_key>")

	// Display a separator...
	fmt.Println("\n==================================\n")

	// Defining the loop to display datas of all existing emojis attached to the string 'cat'...
	for _, emoji := range allSearchForCatListOfEmojis.GetListOfEmojis() {

		// Display the current emoji...
		fmt.Println(emoji.ToString())

		// Break one line...
		fmt.Println("\n")
	}

	// Display a separator...
	fmt.Println("\n==================================\n")
}
```

<a name="example_7"></a>
### Example 7

```go
package main

import (

        "fmt"
        "github.com/Vicken-Ghoubiguian/emojis"
)

// Definition of the main function...
func main() {

        // Declaration and initialization of the 'allEmojisListOfEmojis' variable which will contain all existing emojis...
        allEmojisListOfEmojis := emojis.GetAllEmojis("<Your_access_key>")

        // Initialization of the 'anatomicalHeartEmoji' variable which will contain the 'anatomical heart' emoji...
        anatomicalHeartEmoji := allEmojisListOfEmojis.GetASingleEmoji("anatomical-heart")

        // Display a separator...
        fmt.Println("\n==================================\n")

        // Display the 'anatomical heart' emoji...
        fmt.Println(anatomicalHeartEmoji.ToString())

        // Display a separator...
        fmt.Println("\n==================================\n")
}
```
<a name="example_8"></a>
### Example 8 - Get and display the data of all emojis attached to a particular string...

```go
package main

import (

        "fmt"
        "github.com/Vicken-Ghoubiguian/emojis"
)

// Definition of the main function...
func main() {

        // Declaration and initialization of the 'allEmojisListOfEmojis' variable which will contain all existing emojis...
        allEmojisListOfEmojis := emojis.GetAllEmojis("<Your_access_key>")

		// Declaration and initialization of the 'allSearchForCatListOfEmojis' variable which will contain all existing emojis all existing emojis attached to the string 'cat'...
		allSearchForCatListOfEmojis := allEmojisListOfEmojis.GetSearchedForEmojis("cat")

		// Display a separator...
		fmt.Println("\n==================================\n")

		// Defining the loop to display datas of all existing emojis attached to the string 'cat'...
		for _, emoji := range allSearchForCatListOfEmojis.GetListOfEmojis() {

			// Display the current emoji...
			fmt.Println(emoji.ToString())

			// Break one line...
			fmt.Println("\n")
		}

		// Display a separator...
		fmt.Println("\n==================================\n")
}
```
<a name="example_9"></a>
### Example 9

```go
package main

import (

        "fmt"
        "github.com/Vicken-Ghoubiguian/emojis"
)

// Definition of the main function...
func main() {

    // Declaration and initialization of the 'allCatFacesEmojisListOfEmojis' variable which will contain all existing emojis of the 'cat-face' sub category...
    allCatFacesEmojisListOfEmojis := emojis.GetEmojisFromSubGroup("cat-face", "<Your_access_key>")

	// Display a separator...
	fmt.Println("\n==================================\n")

	// Defining the loop to display datas of all existing emojis of the 'cat-face' sub category...
	for _, emoji := range allCatFacesEmojisListOfEmojis.GetListOfEmojis() {

		// Display the current emoji...
		fmt.Println(emoji.ToString())

		// Break one line...
		fmt.Println("\n")
	}

	// Display a separator...
	fmt.Println("\n==================================\n")
}
```

<a name="example_10"></a>
### Example 10

```go
package main

import (

        "fmt"
		"strings"
		"strconv"
        "github.com/Vicken-Ghoubiguian/emojis"
)

// Definition of the main function...
func main() {

	// Declaration and initialization of the 'allEmojisCategories' variable which will contain all existing emojis categories...
	allEmojisCategories := emojis.GetAllCategories("<Your_access_key>")

	// Display a separator...
	fmt.Println("\n==================================\n")

	// Defining the loop to display datas of all existing emojis categories...
	for _, category := range allEmojisCategories.GetMapOfCategories() {

		// Display datas of the current emojis categories...
		fmt.Println("Slug: " + category.GetSlug())
		fmt.Println("Sub categories array: " + "[" + strings.Join(category.GetSubCategoriesArray(), ", ") + "]")
		fmt.Println("Sub categories account: ", strconv.Itoa(category.GetSubCategoriesAccount()))

		// Break one line...
		fmt.Println("\n")
	}

	// Display a separator...
	fmt.Println("\n==================================\n")
}
```

<a name="example_11"></a>
### Example 11

```go
package main

import (

        "fmt"
        "github.com/Vicken-Ghoubiguian/emojis"
)

// Definition of the main function...
func main() {

        // Declaration and initialization of the 'allEmojisListOfEmojis' variable which will contain all existing emojis...
        allEmojisListOfEmojis := emojis.GetAllEmojis("<Your_access_key>")

        // Display a separator...
        fmt.Println("\n==================================\n")

        // Display all of the emojis in the 'allEmojisListOfEmojis' ListOfEmojis in an alphabetical order...
        allEmojisListOfEmojis.SortAlphabetically()

        // Display a separator...
        fmt.Println("\n==================================\n")
}
```

<a name="example_12"></a>
### Example 12

```go
package main

import (

        "fmt"
        "github.com/Vicken-Ghoubiguian/emojis"
)

// Definition of the main function...
func main() {

        // Declaration and initialization of the 'allEmojisListOfEmojis' variable which will contain all existing emojis...
        allEmojisListOfEmojis := emojis.GetAllEmojis("<Your_access_key>")

        // Display a separator...
        fmt.Println("\n==================================\n")

        // Display all of the emojis in the 'allEmojisListOfEmojis' ListOfEmojis in a reverse alphabetical order...
        allEmojisListOfEmojis.SortReverseAlphabetically()

        // Display a separator...
        fmt.Println("\n==================================\n")
}
```

<a name="example_13"></a>
### Example 13 - Get a list of all emojis and display the corresponding percentage for each category...

```go
package main

import (

        "fmt"
        "github.com/Vicken-Ghoubiguian/emojis"
)

// Definition of the main function...
func main() {

        // Declaration and initialization of the 'allEmojisListOfEmojis' variable which will contain all existing emojis...
        allEmojisListOfEmojis := emojis.GetAllEmojis("<Your_access_key>")

        // Declaration and initialization of the 'allEmojisCategories' variable which will contain all existing emojis categories...
        allEmojisCategories := emojis.GetAllCategories("<Your_access_key>")

        // Declaration and initialization of the 'allEmojisCategories' variable which will contain all existing emojis categories...
        allEmojisCategoriesMap := allEmojisCategories.GetMapOfCategories()

        // Get all emojis categories in each specific variable...
        catg_1 := allEmojisCategoriesMap["smileys-emotion"]
        catg_2 := allEmojisCategoriesMap["people-body"]
        catg_3 := allEmojisCategoriesMap["component"]
        catg_4 := allEmojisCategoriesMap["animals-nature"]
        catg_5 := allEmojisCategoriesMap["food-drink"]
        catg_6 := allEmojisCategoriesMap["travel-places"]
        catg_7 := allEmojisCategoriesMap["activities"]
        catg_8 := allEmojisCategoriesMap["objects"]
        catg_9 := allEmojisCategoriesMap["symbols"]
        catg_10 := allEmojisCategoriesMap["flags"]

		// Display a separator...
        fmt.Println("\n==================================\n")

        // Display all calculated percentage for each emojis category...
        fmt.Printf("'smileys-emotion' emojis: %f %%\n", catg_1.GetPercentageInListOfEmojis(allEmojisListOfEmojis))
		fmt.Printf("'people-body' emojis: %f %%\n", catg_2.GetPercentageInListOfEmojis(allEmojisListOfEmojis))
		fmt.Printf("'component' emojis: %f %%\n", catg_3.GetPercentageInListOfEmojis(allEmojisListOfEmojis))
		fmt.Printf("'animals-nature' emojis: %f %%\n", catg_4.GetPercentageInListOfEmojis(allEmojisListOfEmojis))
		fmt.Printf("'food-drink' emojis: %f %%\n", catg_5.GetPercentageInListOfEmojis(allEmojisListOfEmojis))
		fmt.Printf("'travel-places' emojis: %f %%\n", catg_6.GetPercentageInListOfEmojis(allEmojisListOfEmojis))
		fmt.Printf("'activities' emojis: %f %%\n", catg_7.GetPercentageInListOfEmojis(allEmojisListOfEmojis))
		fmt.Printf("'objects' emojis: %f %%\n", catg_8.GetPercentageInListOfEmojis(allEmojisListOfEmojis))
		fmt.Printf("'symbols' emojis: %f %%\n", catg_9.GetPercentageInListOfEmojis(allEmojisListOfEmojis))
		fmt.Printf("'flags' emojis: %f %%\n", catg_10.GetPercentageInListOfEmojis(allEmojisListOfEmojis))

		// Display a separator...
        fmt.Println("\n==================================\n")
}
```

<a name="example_14"></a>
### Example 14 - Check if an emoji is in a list of emojis...

```go
package main

import (

        "fmt"
        "strconv"
        "github.com/Vicken-Ghoubiguian/emojis"
)

// Definition of the main function...
func main() {

        // Declaration and initialization of the 'allSearchForCatListOfEmojis' variable which will contain all existing emojis attached to the string 'cat'...
        allSearchForCatListOfEmojis := emojis.GetSearchedForEmojis("cat", "<Your_access_key>")

		// Declaration and initialization of the 'allEmojisListOfEmojis' variable which will contain all existing emojis...
        allEmojisListOfEmojis := emojis.GetAllEmojis("<Your_access_key>")

        // Get the 'japanese-application-button' emoji...
        japaneseApplicationButtonEmoji := allSearchForCatListOfEmojis.GetASingleEmoji("japanese-application-button")

        // Initialization of the 'ZzzEmoji' variable which will contain the 'zzz' emoji...
        ZzzEmoji := allEmojisListOfEmojis.GetASingleEmoji("zzz")

        // Display a separator...
        fmt.Println("\n==================================\n")

        // Display if the 'japanese-application-button' emoji is contained in the 'allSearchForCatListOfEmojis' list of emojis...
        fmt.Println("Is the 'japanese-application-button' emoji contained in the cat's list of emojis ? " + strconv.FormatBool(allSearchForCatListOfEmojis.Contains(japaneseApplicationButtonEmoji)) + "\n")

        // Display if the 'zzz' emoji is contained in the 'allSearchForCatListOfEmojis' list of emojis...
        fmt.Println("Is the 'zzz' emoji contained in the cat's list of emojis ? " + strconv.FormatBool(allSearchForCatListOfEmojis.Contains(ZzzEmoji)))

        // Display a separator...
        fmt.Println("\n==================================\n")
}
```

<a name="example_15"></a>
### Example 15 - Get all datas of all emojis and use them in a displayed text...

```go
package main

import (
	
	"fmt"
	"github.com/Vicken-Ghoubiguian/emojis"
)

// Definition of the main function...
func main() {

	// Declaration and initialization of the 'allEmojisListOfEmojis' variable which will contain all existing emojis...
	allEmojisListOfEmojis := emojis.GetAllEmojis("<Your_access_key>")

	// Get a list of all emojis as a map[string]Emoji...
	allEmojisMapOfEmojis := allEmojisListOfEmojis.GetListOfEmojis()

	// Get all necessary emojis...
	grinningCatEmoji := allEmojisMapOfEmojis["grinning-cat"]
	foldedHandsEmoji := allEmojisMapOfEmojis["folded-hands"]
	smilingCatWithHeartEyesEmoji := allEmojisMapOfEmojis["smiling-cat-with-heart-eyes"]

	// Display a separator...
    fmt.Println("\n==================================\n")
	
	// Display the text with emojis...
	fmt.Println("Hello world " + grinningCatEmoji.GetCharacter() + "! Welcome to my awesome go package " + foldedHandsEmoji.GetCharacter() + smilingCatWithHeartEyesEmoji.GetCharacter() + "...")

	// Display a separator...
    fmt.Println("\n==================================\n")
}
```

<a name="example_16"></a>
### Example 16 - Delete an emoji from a list of emojis...

```go
package main

import (

        "fmt"
        "github.com/Vicken-Ghoubiguian/emojis"
)

// Definition of the main function...
func main() {

        // Declaration and initialization of the 'allSearchForCatListOfEmojis' variable which will contain all existing emojis attached to the string 'cat'...
        allSearchForCatListOfEmojis := emojis.GetSearchedForEmojis("cat", "<Your_access_key>")

        // Get the 'japanese-application-button' emoji which will be deleted from the 'allSearchForCatListOfEmojis' because it has nothing to do with cats...
        japaneseApplicationButtonEmoji := allSearchForCatListOfEmojis.GetASingleEmoji("japanese-application-button")

        // Delete the 'japaneseApplicationButtonEmoji' emoji from the 'allSearchForCatListOfEmojis' list of emojis...
        allSearchForCatListOfEmojis.Delete(japaneseApplicationButtonEmoji)

        // Display a separator...
        fmt.Println("\n==================================\n")

        // Defining the loop to display datas of all existing emojis attached to the string 'cat'...
        for _, emoji := range allSearchForCatListOfEmojis.GetListOfEmojis() {

                // Display datas of the current emojis...
                fmt.Println(emoji.ToString())

                // Break one line...
                fmt.Println("\n")
        }

        // Display a separator...
        fmt.Println("\n==================================\n")
}
```

<a name="example_17"></a>
### Example 17 - Add an emoji to a list of emojis...

```go
package main

import (

        "fmt"
        "github.com/Vicken-Ghoubiguian/emojis"
)

// Definition of the main function...
func main() {

        // Declaration and initialization of the 'allEmojisListOfEmojis' variable which will contain all existing emojis...
        allEmojisListOfEmojis := emojis.GetAllEmojis("<Your_access_key>")

        // Declaration and initialization of the 'allSmileysEmojisListOfEmojis' variable which will contain all existing emojis of the 'smileys-emotion' category...
        allSmileysEmojisListOfEmojis := allEmojisListOfEmojis.GetEmojisFromGroup("smileys-emotion")

		// Initialization of the 'anatomicalHeartEmoji' variable which will contain the 'anatomical heart' emoji...
        anatomicalHeartEmoji := allEmojisListOfEmojis.GetASingleEmoji("anatomical-heart")

		// Add the 'anatomicalHeartEmoji' to the 'allSmileysEmojisListOfEmojis' list of emojis...
        allSmileysEmojisListOfEmojis.Add(anatomicalHeartEmoji)

		// Display a separator...
        fmt.Println("\n==================================\n")

        // Defining the loop to display datas of all existing emojis of the 'smileys-emotion' category...
        for _, emoji := range allSmileysEmojisListOfEmojis.GetListOfEmojis() {

                // Display datas of the current emojis...
                fmt.Println(emoji.ToString())

                // Break one line...
                fmt.Println("\n")
        }

        // Display a separator...
        fmt.Println("\n==================================\n")
}
```

<a name="example_18"></a>
### Example 18

```go
package main

import (

        "fmt"
        "github.com/Vicken-Ghoubiguian/emojis"
)

// Definition of the main function...
func main() {

        // Declaration and initialization of the 'allEmojisListOfEmojis' variable which will contain all existing emojis...
        allEmojisListOfEmojis := emojis.GetAllEmojis("<Your_access_key>")

        // Declaration and initialization of the 'allSmileysEmojisListOfEmojis' variable which will contain all existing emojis of the 'smileys-emotion' category...
        allSmileysEmojisListOfEmojis := allEmojisListOfEmojis.GetEmojisFromGroup("smileys-emotion")

        // Declaration and initialization of the 'allComputersEmojisListOfEmojis' variable which will contain all existing emojis of the 'computer' category...
        allComputersEmojisListOfEmojis := allEmojisListOfEmojis.GetEmojisFromGroup("computer")

        // Initialization of the 'anatomicalHeartEmoji' variable which will contain the 'anatomical heart' emoji...
        anatomicalHeartEmoji := allEmojisListOfEmojis.GetASingleEmoji("anatomical-heart")

        // Initialization of the 'ZzzEmoji' variable which will contain the 'zzz' emoji...
        ZzzEmoji := allEmojisListOfEmojis.GetASingleEmoji("zzz")

        // Add the 'anatomicalHeartEmoji' to the 'allSmileysEmojisListOfEmojis' list of emojis...
        allSmileysEmojisListOfEmojis.Add(anatomicalHeartEmoji)

        // Add the 'ZzzEmoji' to the 'allSmileysEmojisListOfEmojis' list of emojis...
        allSmileysEmojisListOfEmojis.Add(ZzzEmoji)

        // Add the 'anatomicalHeartEmoji' to the 'allComputersEmojisListOfEmojis' list of emojis...
        allComputersEmojisListOfEmojis.Add(anatomicalHeartEmoji)

        // Add the 'ZzzEmoji' to the 'allComputersEmojisListOfEmojis' list of emojis...
        allComputersEmojisListOfEmojis.Add(ZzzEmoji)

        // Determine the intersection between the 'allFacesEmojisListOfEmojis' set and the 'allEmojisListOfEmojis' set...
        intersectionListOfEmojis := allSmileysEmojisListOfEmojis.Intersection(allComputersEmojisListOfEmojis)

        // Display a separator...
        fmt.Println("\n==================================\n")

		// Display some explanations...
		fmt.Println("All_smileys_emojis_set ⊂ All_computers_emojis_set = ")

        // Display the intersection as the 'intersectionListOfEmojis' list of emojis...
		fmt.Println(intersectionListOfEmojis)

        // Display a separator...
        fmt.Println("\n==================================\n")
}
```
<a name="example_19"></a>
### Example 19

```go
package main

import (
	
	"fmt"
	"github.com/Vicken-Ghoubiguian/emojis"
)

// Definition of the main function...
func main() {

	// Declaration and initialization of the 'allEmojisListOfEmojis' variable which will contain all existing emojis...
	allEmojisListOfEmojis := emojis.GetAllEmojis("<Your_access_key>")

	// Declaration and initialization of the 'allFacesEmojisListOfEmojis' variable which will contain all existing emojis all existing emojis attached to the string 'face'...
    allFacesEmojisListOfEmojis := allEmojisListOfEmojis.GetSearchedForEmojis("face")

	// Determine the intersection between the 'allFacesEmojisListOfEmojis' set and the 'allEmojisListOfEmojis' set...
	intersectionListOfEmojis := allFacesEmojisListOfEmojis.Intersection(allEmojisListOfEmojis)

	// Display a separator...
    fmt.Println("\n==================================\n")

	// Display some explanations...
	fmt.Println("All_faces_emojis_set ⊂ All_emojis_set = ")

	// Display the intersection as the 'intersectionListOfEmojis' list of emojis...
	fmt.Println(intersectionListOfEmojis)

	// Display a separator...
    fmt.Println("\n==================================\n")
}
```
<a name="example_20"></a>
### Example 20 - Check if an emoji belong to a list of emojis...

```go
```

<a name="where_to_use_it"></a>
## Where to use it ?

<a name="useful_links"></a>
## Useful links

* [Open Emoji API](https://emoji-api.com)
* [golang convert interface to map - GitHub Gist](https://gist.github.com/nevzatalkan/f5c5ef66e88dd446976401967b6731e8)
* [Convert interface to string - yourbasic.org](https://yourbasic.org/golang/interface-to-string/)
* [Structure Your Go Project Into Multiple Directories => Go module](https://www.jodylecompte.com/posts/go-structure-your-go-project/#go-modules)

<a name="conclusion"></a>
## Conclusion

Comming soon...
