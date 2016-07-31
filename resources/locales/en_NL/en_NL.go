package en_NL

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type en_NL struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'en_NL' locale
func New() locales.Translator {
	return &en_NL{
		locale:  "en_NL",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *en_NL) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'en_NL'
func (t *en_NL) Plurals() []locales.PluralRule {
	return t.plurals
}

// CardinalPluralRule returns the PluralRule given 'num' for 'en_NL'
func (t *en_NL) CardinalPluralRule(num string) (locales.PluralRule, error) {

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