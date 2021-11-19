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

	// Display the whole 'allEmojisListOfEmojis' list of all existing emojis...
	//fmt.Println(allEmojisListOfEmojis.GetListOfEmojis())

	// Display a separator...
	fmt.Println("\n==================================\n")

	// Defining the loop to display datas of all existing emojis...
	for slug, emoji := range allEmojisListOfEmojis.GetListOfEmojis() {

		// Display datas of the current emojis...
		fmt.Println("Character: " + allEmojisListOfEmojis[slug].GetCharacter())
		fmt.Println("Slug: " + allEmojisListOfEmojis[slug].GetSlug())
		fmt.Println("Unicode name: " + allEmojisListOfEmojis[slug].GetUnicodeName())
		fmt.Println("Code point: " + allEmojisListOfEmojis[slug].GetCodePoint())
		fmt.Println("Group: " + allEmojisListOfEmojis[slug].GetGroup())
		fmt.Println("Subgroup: " + allEmojisListOfEmojis[slug].GetSubGroup())

		// Break 3 lines...
		fmt.Println("\n\n\n")
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
