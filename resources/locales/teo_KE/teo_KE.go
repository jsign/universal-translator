package teo_KE

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type teo_KE struct {
	locale   string
	plurals  []locales.PluralRule
	decimal  []byte
	group    []byte
	minus    []byte
	percent  []byte
	perMille []byte
	symbol   []byte
}

// New returns a new instance of translator for the 'teo_KE' locale
func New() locales.Translator {
	return &teo_KE{
		locale:   "teo_KE",
		plurals:  []locales.PluralRule{2, 6},
		decimal:  []byte{},
		group:    []byte{},
		minus:    []byte{},
		percent:  []byte{},
		perMille: []byte{},
		symbol:   []byte{},
	}
}

// Locale returns the current translators string locale
func (t *teo_KE) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'teo_KE'
func (t *teo_KE) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'teo_KE'
func (t *teo_KE) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}