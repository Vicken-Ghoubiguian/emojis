# emojis

Package written in Go to get, return, treat and manage emojis ⛑ 🪱 💥 🏵️ 🇳🇺 💾 🏟️ 🫁 💦 🦭🐽 🔱 🦋 📌 👹 ♨️ 🇳🇷 📐 🤬 🐍 🐈 ⚜️ 🔫 🚄 ❤️‍🩹 🌋 🎓 💈 🎃 🥑 😈 🥥 😸 🐾 🥵 ♻️ 🦾 🙀 🎗️ ☠️ 🫀 👻 🇰🇵 🫐 🗾 🏗️ 🤢 🔕 🏧 👽 🪔 🌹 🪙 ⚕️🛎️ 💩 🐡 🐙 😻 🧳 🎥 💻 🥶 🔘 🖥️ 🔮 🐳 🕳️ 💯 💫 🪐 🧨 🗺️ 🇨🇺 🦣 🦈 🐲 🧱 🍳 🇨🇼 ⛩️ 🤯 🌡️ 🥩 👺 📎 💳 🦧 🎳 🧶 🏛️ 🧠 🥀 🛑 ✔️ 🐬 🇮🇲 🔥 from the [Open Emoji API](https://emoji-api.com) in every Go project...

## Contents

* [Introduction](#introduction)
* [Presentation](#presentation)
* [Project's structure](#project_s_structure)
* [How was this project developed ?](#how_was_this_project_developed)
* [How to use this package ?](#how_to_use_this_package)
* [Little examples](#little_examples)
	* [Example 1](#example_1)
	* [Example 2](#example_2)
	* [Example 3](#example_3)
	* [Example 4](#example_4)
	* [Example 5](#example_5)
	* [Example 6](#example_6)
	* [Example 7](#example_7)
* [Where to use it ?](#where_to_use_it)
* [Useful links](#useful_links)
* [Conclusion](#conclusion)

<a name="introduction"></a>
## Introduction

<a name="presentation"></a>
## Presentation

<a name="project_s_structure"></a>
## Project's structure

<a name="how_was_this_project_developed"></a>
## How was this project developed ?

<a name="how_to_use_this_package"></a>
## How to use this package ?

```bash
go get -u github.com/Vicken-Ghoubiguian/emojis
```

<a name="little_examples"></a>
## Little examples

<a name="example_1"></a>
### Example 1

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

		// Display datas of the current emojis...
		fmt.Println("Character: " + emoji.GetCharacter())
		fmt.Println("Slug: " + emoji.GetSlug())
		fmt.Println("Unicode name: " + emoji.GetUnicodeName())
		fmt.Println("Code point: " + emoji.GetCodePoint())
		fmt.Println("Group: " + emoji.GetGroup())
		fmt.Println("Subgroup: " + emoji.GetSubGroup())

		// Break one line...
		fmt.Println("\n")
	}

	// Display a separator...
	fmt.Println("\n==================================\n")
}
```

<a name="example_2"></a>
### Example 2

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

	// Break 3 lines...
	fmt.Println("\n\n\n")	

	// Display all attributes of the 'anatomical heart' emoji...
	fmt.Println("Character: " + anatomicalHeartEmoji.GetCharacter())
	fmt.Println("Slug: " + anatomicalHeartEmoji.GetSlug())
	fmt.Println("Unicode name: " + anatomicalHeartEmoji.GetUnicodeName())
	fmt.Println("Code point: " + anatomicalHeartEmoji.GetCodePoint())
	fmt.Println("Group: " + anatomicalHeartEmoji.GetGroup())
	fmt.Println("Subgroup: " + anatomicalHeartEmoji.GetSubGroup())

	// Break 3 lines...
	fmt.Println("\n\n\n")
}
```

<a name="example_3"></a>
### Example 3

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

			// Display datas of the current emojis...
			fmt.Println("Character: " + emoji.GetCharacter())
			fmt.Println("Slug: " + emoji.GetSlug())
			fmt.Println("Unicode name: " + emoji.GetUnicodeName())
			fmt.Println("Code point: " + emoji.GetCodePoint())
			fmt.Println("Group: " + emoji.GetGroup())
			fmt.Println("Subgroup: " + emoji.GetSubGroup())

			// Break one line...
			fmt.Println("\n")
		}

		// Display a separator...
		fmt.Println("\n==================================\n")
}
```

<a name="example_4"></a>
### Example 4

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

		// Display datas of the current emojis...
		fmt.Println("Character: " + emoji.GetCharacter())
		fmt.Println("Slug: " + emoji.GetSlug())
		fmt.Println("Unicode name: " + emoji.GetUnicodeName())
		fmt.Println("Code point: " + emoji.GetCodePoint())
		fmt.Println("Group: " + emoji.GetGroup())
		fmt.Println("Subgroup: " + emoji.GetSubGroup())

		// Break one line...
		fmt.Println("\n")
	}

	// Display a separator...
	fmt.Println("\n==================================\n")
}
```

<a name="example_5"></a>
### Example 5

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

			// Display datas of the current emojis...
			fmt.Println("Character: " + emoji.GetCharacter())
			fmt.Println("Slug: " + emoji.GetSlug())
			fmt.Println("Unicode name: " + emoji.GetUnicodeName())
			fmt.Println("Code point: " + emoji.GetCodePoint())
			fmt.Println("Group: " + emoji.GetGroup())
			fmt.Println("Subgroup: " + emoji.GetSubGroup())

			// Break one line...
			fmt.Println("\n")
		}

		// Display a separator...
		fmt.Println("\n==================================\n")
}
```

<a name="example_6"></a>
### Example 6

```go
package main

import (

        "fmt"
        "github.com/Vicken-Ghoubiguian/emojis"
)

// Definition of the main function...
func main() {

	//
    allSearchForComputerListOfEmojis := emojis.GetSearchedForEmojis("computer", "<Your_access_key>")

	// Display a separator...
	fmt.Println("\n==================================\n")

	//
	for _, emoji := range allSearchForComputerListOfEmojis.GetListOfEmojis() {

		// Display datas of the current emojis...
		fmt.Println("Character: " + emoji.GetCharacter())
		fmt.Println("Slug: " + emoji.GetSlug())
		fmt.Println("Unicode name: " + emoji.GetUnicodeName())
		fmt.Println("Code point: " + emoji.GetCodePoint())
		fmt.Println("Group: " + emoji.GetGroup())
		fmt.Println("Subgroup: " + emoji.GetSubGroup())

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
```

<a name="where_to_use_it"></a>
## Where to use it ?

<a name="useful_links"></a>
## Useful links

* [Open Emoji API](https://emoji-api.com)
* [golang convert interface to map - GitHub Gist](https://gist.github.com/nevzatalkan/f5c5ef66e88dd446976401967b6731e8)
* [Convert interface to string - yourbasic.org](https://yourbasic.org/golang/interface-to-string/)

<a name="conclusion"></a>
## Conclusion
