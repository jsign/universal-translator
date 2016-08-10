package rm

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type rm struct {
	locale   string
	plurals  []locales.PluralRule
	decimal  []byte
	group    []byte
	minus    []byte
	percent  []byte
	perMille []byte
	symbol   []byte
}

// New returns a new instance of translator for the 'rm' locale
func New() locales.Translator {
	return &rm{
		locale:   "rm",
		plurals:  []locales.PluralRule{2, 6},
		decimal:  []byte{0x2e},
		group:    []byte{0xe2, 0x80, 0x99},
		minus:    []byte{0xe2, 0x88, 0x92},
		percent:  []byte{0x25},
		perMille: []byte{0xe2, 0x80, 0xb0},
		symbol:   []byte{},
	}
}

// Locale returns the current translators string locale
func (t *rm) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'rm'
func (t *rm) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'rm'
func (t *rm) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}