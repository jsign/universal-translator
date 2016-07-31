package ca_IT

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type ca_IT struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'ca_IT' locale
func New() locales.Translator {
	return &ca_IT{
		locale:  "ca_IT",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *ca_IT) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'ca_IT'
func (t *ca_IT) Plurals() []locales.PluralRule {
	return t.plurals
}

// CardinalPluralRule returns the PluralRule given 'num' for 'ca_IT'
func (t *ca_IT) CardinalPluralRule(num string) (locales.PluralRule, error) {

	i, err := locales.I(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	v := locales.V(num)

	if i == 1 && v == 0 {
		return locales.PluralRuleOne, nil
	}

	return locales.PluralRuleOther, nil
}