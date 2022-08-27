// Package main ...
package main

import (
	"fmt"

	"github.com/moredure/xrod"
	"github.com/moredure/xrod/lib/launcher"
)

func main() {
	l := launcher.New()

	// For more info: https://pkg.go.dev/github.com/moredure/xrod/lib/launcher
	u := l.MustLaunch()

	browser := rod.New().ControlURL(u).MustConnect()

	page := browser.MustPage("http://example.com")

	fmt.Println(page.MustInfo().Title)
}
