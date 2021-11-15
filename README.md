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

	//
	allEmojisMap := emojis.GetAllEmojis("<Your_access_key>")

	//
	fmt.Println(allEmojisMap)

	//
	fmt.Println("\n\n\n")

	//
	anatomicalHeartEmoji := emojis.GetASingleEmoji("anatomical-heart", "<Your_access_key>")

	//
	fmt.Println("Character: " + anatomicalHeartEmoji.GetCharacter())
	fmt.Println("Slug: " + anatomicalHeartEmoji.GetSlug())
	fmt.Println("Unicode name: " + anatomicalHeartEmoji.GetUnicodeName())
	fmt.Println("Code point: " + anatomicalHeartEmoji.GetCodePoint())
	fmt.Println("Group: " + anatomicalHeartEmoji.GetGroup())
	fmt.Println("Subgroup: " + anatomicalHeartEmoji.GetSubGroup())

	//
	fmt.Println("\n\n\n")
}
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
