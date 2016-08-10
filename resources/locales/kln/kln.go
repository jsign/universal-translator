package kln

import "github.com/go-playground/universal-translator/resources/locales"

type kln struct {
	locale   string
	plurals  []locales.PluralRule
	decimal  []byte
	group    []byte
	minus    []byte
	percent  []byte
	perMille []byte
	symbol   []byte
}

// New returns a new instance of translator for the 'kln' locale
func New() locales.Translator {
	return &kln{
		locale:   "kln",
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
func (t *kln) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'kln'
func (t *kln) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'kln'
func (t *kln) cardinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}