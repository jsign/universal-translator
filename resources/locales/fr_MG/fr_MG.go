package fr_MG

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type fr_MG struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'fr_MG' locale
func New() locales.Translator {
	return &fr_MG{
		locale:  "fr_MG",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *fr_MG) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'fr_MG'
func (t *fr_MG) Plurals() []locales.PluralRule {
	return t.plurals
}

// CardinalPluralRule returns the PluralRule given 'num' for 'fr_MG'
func (t *fr_MG) CardinalPluralRule(num string) (locales.PluralRule, error) {

	i, err := locales.I(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	if i == 0 || i == 1 {
		return locales.PluralRuleOne, nil
	}

	return locales.PluralRuleOther, nil
}