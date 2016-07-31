package pa_Arab_PK

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type pa_Arab_PK struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'pa_Arab_PK' locale
func New() locales.Translator {
	return &pa_Arab_PK{
		locale:  "pa_Arab_PK",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *pa_Arab_PK) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'pa_Arab_PK'
func (t *pa_Arab_PK) Plurals() []locales.PluralRule {
	return t.plurals
}

// CardinalPluralRule returns the PluralRule given 'num' for 'pa_Arab_PK'
func (t *pa_Arab_PK) CardinalPluralRule(num string) (locales.PluralRule, error) {

	n, err := locales.N(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	if n >= 0 && n <= 1 {
		return locales.PluralRuleOne, nil
	}

	return locales.PluralRuleOther, nil
}