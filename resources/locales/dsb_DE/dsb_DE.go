package dsb_DE

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type dsb_DE struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'dsb_DE' locale
func New() locales.Translator {
	return &dsb_DE{
		locale:  "dsb_DE",
		plurals: []locales.PluralRule{2, 3, 4, 6},
	}
}

// Locale returns the current translators string locale
func (t *dsb_DE) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'dsb_DE'
func (t *dsb_DE) Plurals() []locales.PluralRule {
	return t.plurals
}

// CardinalPluralRule returns the PluralRule given 'num' for 'dsb_DE'
func (t *dsb_DE) CardinalPluralRule(num string) (locales.PluralRule, error) {

	i, err := locales.I(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	v := locales.V(num)

	f, err := locales.F(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	if (v == 0 && i%100 == 1) || (f%100 == 1) {
		return locales.PluralRuleOne, nil
	} else if (v == 0 && i%100 == 2) || (f%100 == 2) {
		return locales.PluralRuleTwo, nil
	} else if (v == 0 && i%100 >= 3 && i%100 <= 4) || (f%100 >= 3 && f%100 <= 4) {
		return locales.PluralRuleFew, nil
	}

	return locales.PluralRuleOther, nil
}