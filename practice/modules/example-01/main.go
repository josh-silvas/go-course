// Modules Example: In this example, we will create a simple package that's dependencies are managed via go modules.
// Before you begin:
//   1. Change to the practice directory
//       - cd practice/modules/example-01
//   2. Complete the `TO DO` sections of the code
//   3. Check answers in answer/main.go

package main

import (
	"fmt"

	"github.com/labstack/gommon/color"
)

func main() {
	fmt.Println(color.Cyan("Yay! I did a thing!"))
}

// Results:
// ---------------------------------------------------------------------
// Sun, 08 Dec 2019 20:36:13 CST found 10 articles from techcrunch
// Sun, 08 Dec 2019 20:36:13 CST found 10 articles from polygon
