package kn

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type kn struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'kn' locale
func New() locales.Translator {
	return &kn{
		locale:  "kn",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *kn) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'kn'
func (t *kn) Plurals() []locales.PluralRule {
	return t.plurals
}

// CardinalPluralRule returns the PluralRule given 'num' for 'kn'
func (t *kn) CardinalPluralRule(num string) (locales.PluralRule, error) {

	n, err := locales.N(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	i, err := locales.I(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	if (i == 0) || (n == 1) {
		return locales.PluralRuleOne, nil
	}

	return locales.PluralRuleOther, nil
}