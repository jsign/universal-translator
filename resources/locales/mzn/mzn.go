package mzn

import "github.com/go-playground/universal-translator/resources/locales"

type mzn struct {
	locale   string
	plurals  []locales.PluralRule
	decimal  []byte
	group    []byte
	minus    []byte
	percent  []byte
	perMille []byte
	symbol   []byte
}

// New returns a new instance of translator for the 'mzn' locale
func New() locales.Translator {
	return &mzn{
		locale:   "mzn",
		plurals:  nil,
		decimal:  []byte{},
		group:    []byte{},
		minus:    []byte{},
		percent:  []byte{},
		perMille: []byte{},
		symbol:   []byte{},
	}
}

// Locale returns the current translators string locale
func (t *mzn) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'mzn'
func (t *mzn) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'mzn'
func (t *mzn) cardinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}