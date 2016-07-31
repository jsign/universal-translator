package yi_001

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type yi_001 struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'yi_001' locale
func New() locales.Translator {
	return &yi_001{
		locale:  "yi_001",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *yi_001) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'yi_001'
func (t *yi_001) Plurals() []locales.PluralRule {
	return t.plurals
}

// CardinalPluralRule returns the PluralRule given 'num' for 'yi_001'
func (t *yi_001) CardinalPluralRule(num string) (locales.PluralRule, error) {

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