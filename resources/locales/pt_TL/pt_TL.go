package pt_TL

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type pt_TL struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'pt_TL' locale
func New() locales.Translator {
	return &pt_TL{
		locale:  "pt_TL",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *pt_TL) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'pt_TL'
func (t *pt_TL) Plurals() []locales.PluralRule {
	return t.plurals
}

// CardinalPluralRule returns the PluralRule given 'num' for 'pt_TL'
func (t *pt_TL) CardinalPluralRule(num string) (locales.PluralRule, error) {

	n, err := locales.N(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	if n >= 0 && n <= 2 && n != 2 {
		return locales.PluralRuleOne, nil
	}

	return locales.PluralRuleOther, nil
}