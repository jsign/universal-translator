package ur

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type ur struct {
	locale   string
	plurals  []locales.PluralRule
	decimal  []byte
	group    []byte
	minus    []byte
	percent  []byte
	perMille []byte
	symbol   []byte
}

// New returns a new instance of translator for the 'ur' locale
func New() locales.Translator {
	return &ur{
		locale:   "ur",
		plurals:  []locales.PluralRule{2, 6},
		decimal:  []byte{0x2e},
		group:    []byte{0x2c},
		minus:    []byte{0xe2, 0x80, 0x8e, 0x2d, 0xe2, 0x80, 0x8e},
		percent:  []byte{0x25},
		perMille: []byte{0xe2, 0x80, 0xb0},
		symbol:   []byte{},
	}
}

// Locale returns the current translators string locale
func (t *ur) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'ur'
func (t *ur) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'ur'
func (t *ur) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)
	i := int64(n)

	if i == 1 && v == 0 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}