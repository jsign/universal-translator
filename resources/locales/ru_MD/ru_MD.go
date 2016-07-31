package ru_MD

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type ru_MD struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'ru_MD' locale
func New() locales.Translator {
	return &ru_MD{
		locale:  "ru_MD",
		plurals: []locales.PluralRule{2, 4, 5, 6},
	}
}

// Locale returns the current translators string locale
func (t *ru_MD) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'ru_MD'
func (t *ru_MD) Plurals() []locales.PluralRule {
	return t.plurals
}

// CardinalPluralRule returns the PluralRule given 'num' for 'ru_MD'
func (t *ru_MD) CardinalPluralRule(num string) (locales.PluralRule, error) {

	i, err := locales.I(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	v := locales.V(num)

	if v == 0 && i%10 == 1 && i%100 != 11 {
		return locales.PluralRuleOne, nil
	} else if v == 0 && i%10 >= 2 && i%10 <= 4 && i%100 < 12 && i%100 > 14 {
		return locales.PluralRuleFew, nil
	} else if (v == 0 && i%10 == 0) || (v == 0 && i%10 >= 5 && i%10 <= 9) || (v == 0 && i%100 >= 11 && i%100 <= 14) {
		return locales.PluralRuleMany, nil
	}

	return locales.PluralRuleOther, nil
}