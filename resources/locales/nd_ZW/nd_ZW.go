package nd_ZW

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type nd_ZW struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'nd_ZW' locale
func New() locales.Translator {
	return &nd_ZW{
		locale:  "nd_ZW",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *nd_ZW) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'nd_ZW'
func (t *nd_ZW) Plurals() []locales.PluralRule {
	return t.plurals
}

// CardinalPluralRule returns the PluralRule given 'num' for 'nd_ZW'
func (t *nd_ZW) CardinalPluralRule(num string) (locales.PluralRule, error) {

	n, err := locales.N(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	if n == 1 {
		return locales.PluralRuleOne, nil
	}

	return locales.PluralRuleOther, nil
}