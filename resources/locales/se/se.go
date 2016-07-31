package se

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type se struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'se' locale
func New() locales.Translator {
	return &se{
		locale:  "se",
		plurals: []locales.PluralRule{2, 3, 6},
	}
}

// Locale returns the current translators string locale
func (t *se) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'se'
func (t *se) Plurals() []locales.PluralRule {
	return t.plurals
}

// CardinalPluralRule returns the PluralRule given 'num' for 'se'
func (t *se) CardinalPluralRule(num string) (locales.PluralRule, error) {

	n, err := locales.N(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	if n == 1 {
		return locales.PluralRuleOne, nil
	} else if n == 2 {
		return locales.PluralRuleTwo, nil
	}

	return locales.PluralRuleOther, nil
}