package en_GH

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type en_GH struct {
	locale   string
	plurals  []locales.PluralRule
	decimal  []byte
	group    []byte
	minus    []byte
	percent  []byte
	perMille []byte
	symbol   []byte
}

// New returns a new instance of translator for the 'en_GH' locale
func New() locales.Translator {
	return &en_GH{
		locale:   "en_GH",
		plurals:  []locales.PluralRule{2, 6},
		decimal:  []byte{0x5b, 0x5d, 0x62, 0x79, 0x74, 0x65, 0x7b, 0x30, 0x78, 0x32, 0x65, 0x7d},
		group:    []byte{0x5b, 0x5d, 0x62, 0x79, 0x74, 0x65, 0x7b, 0x30, 0x78, 0x32, 0x63, 0x7d},
		minus:    []byte{0x5b, 0x5d, 0x62, 0x79, 0x74, 0x65, 0x7b, 0x30, 0x78, 0x32, 0x64, 0x7d},
		percent:  []byte{0x5b, 0x5d, 0x62, 0x79, 0x74, 0x65, 0x7b, 0x30, 0x78, 0x32, 0x35, 0x7d},
		perMille: []byte{0x5b, 0x5d, 0x62, 0x79, 0x74, 0x65, 0x7b, 0x30, 0x78, 0x65, 0x32, 0x2c, 0x20, 0x30, 0x78, 0x38, 0x30, 0x2c, 0x20, 0x30, 0x78, 0x62, 0x30, 0x7d},
		symbol:   []byte{},
	}
}

// Locale returns the current translators string locale
func (t *en_GH) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'en_GH'
func (t *en_GH) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'en_GH'
func (t *en_GH) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)
	i := int64(n)

	if i == 1 && v == 0 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}