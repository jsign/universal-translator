package ti_ER

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type ti_ER struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'ti_ER' locale
func New() locales.Translator {
	return &ti_ER{
		locale:  "ti_ER",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *ti_ER) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'ti_ER'
func (t *ti_ER) Plurals() []locales.PluralRule {
	return t.plurals
}

// CardinalPluralRule returns the PluralRule given 'num' for 'ti_ER'
func (t *ti_ER) CardinalPluralRule(num string) (locales.PluralRule, error) {

	n, err := locales.N(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	if n >= 0 && n <= 1 {
		return locales.PluralRuleOne, nil
	}

	return locales.PluralRuleOther, nil
}