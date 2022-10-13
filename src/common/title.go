package common

import (
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

var caser = cases.Title(language.English, cases.NoLower)

func Title(original string) string {
	return caser.String(original)
}
