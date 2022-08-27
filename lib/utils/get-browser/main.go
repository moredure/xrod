// Package main ...
package main

import (
	"fmt"

	"github.com/moredure/xrod/lib/launcher"
	"github.com/moredure/xrod/lib/utils"
)

func main() {
	p, err := launcher.NewBrowser().Get()
	utils.E(err)

	fmt.Println(p)
}
