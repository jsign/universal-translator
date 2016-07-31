package ar_ER

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type ar_ER struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'ar_ER' locale
func New() locales.Translator {
	return &ar_ER{
		locale:  "ar_ER",
		plurals: []locales.PluralRule{1, 2, 3, 4, 5, 6},
	}
}

// Locale returns the current translators string locale
func (t *ar_ER) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'ar_ER'
func (t *ar_ER) Plurals() []locales.PluralRule {
	return t.plurals
}

// CardinalPluralRule returns the PluralRule given 'num' for 'ar_ER'
func (t *ar_ER) CardinalPluralRule(num string) (locales.PluralRule, error) {

	n, err := locales.N(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	if n == 0 {
		return locales.PluralRuleZero, nil
	} else if n == 1 {
		return locales.PluralRuleOne, nil
	} else if n == 2 {
		return locales.PluralRuleTwo, nil
	} else if n%100 >= 3 && n%100 <= 10 {
		return locales.PluralRuleFew, nil
	} else if n%100 >= 11 && n%100 <= 99 {
		return locales.PluralRuleMany, nil
	}

	return locales.PluralRuleOther, nil
}