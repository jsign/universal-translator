package xog

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type xog struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'xog' locale
func New() locales.Translator {
	return &xog{
		locale:  "xog",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *xog) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'xog'
func (t *xog) Plurals() []locales.PluralRule {
	return t.plurals
}

// CardinalPluralRule returns the PluralRule given 'num' for 'xog'
func (t *xog) CardinalPluralRule(num string) (locales.PluralRule, error) {

	n, err := locales.N(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	if n == 1 {
		return locales.PluralRuleOne, nil
	}

	return locales.PluralRuleOther, nil
}