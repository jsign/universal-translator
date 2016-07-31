package ja

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type ja struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'ja' locale
func New() locales.Translator {
	return &ja{
		locale:  "ja",
		plurals: []locales.PluralRule{6},
	}
}

// Locale returns the current translators string locale
func (t *ja) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'ja'
func (t *ja) Plurals() []locales.PluralRule {
	return t.plurals
}

// CardinalPluralRule returns the PluralRule given 'num' for 'ja'
func (t *ja) CardinalPluralRule(num string) (locales.PluralRule, error) {

	return locales.PluralRuleOther, nil
}