package ko_KR

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type ko_KR struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'ko_KR' locale
func New() locales.Translator {
	return &ko_KR{
		locale:  "ko_KR",
		plurals: []locales.PluralRule{6},
	}
}

// Locale returns the current translators string locale
func (t *ko_KR) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'ko_KR'
func (t *ko_KR) Plurals() []locales.PluralRule {
	return t.plurals
}

// CardinalPluralRule returns the PluralRule given 'num' for 'ko_KR'
func (t *ko_KR) CardinalPluralRule(num string) (locales.PluralRule, error) {

	return locales.PluralRuleOther, nil
}