package ee_TG

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type ee_TG struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'ee_TG' locale
func New() locales.Translator {
	return &ee_TG{
		locale:  "ee_TG",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *ee_TG) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'ee_TG'
func (t *ee_TG) Plurals() []locales.PluralRule {
	return t.plurals
}

// CardinalPluralRule returns the PluralRule given 'num' for 'ee_TG'
func (t *ee_TG) CardinalPluralRule(num string) (locales.PluralRule, error) {

	n, err := locales.N(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	if n == 1 {
		return locales.PluralRuleOne, nil
	}

	return locales.PluralRuleOther, nil
}