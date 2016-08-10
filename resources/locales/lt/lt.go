package lt

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type lt struct {
	locale   string
	plurals  []locales.PluralRule
	decimal  []byte
	group    []byte
	minus    []byte
	percent  []byte
	perMille []byte
	symbol   []byte
}

// New returns a new instance of translator for the 'lt' locale
func New() locales.Translator {
	return &lt{
		locale:   "lt",
		plurals:  []locales.PluralRule{2, 4, 5, 6},
		decimal:  []byte{0x2c},
		group:    []byte{0xc2, 0xa0},
		minus:    []byte{0xe2, 0x88, 0x92},
		percent:  []byte{0x25},
		perMille: []byte{0xe2, 0x80, 0xb0},
		symbol:   []byte{},
	}
}

// Locale returns the current translators string locale
func (t *lt) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'lt'
func (t *lt) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'lt'
func (t *lt) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)
	f := locales.F(n, v)

	if n%10 == 1 && n%100 < 11 && n%100 > 19 {
		return locales.PluralRuleOne
	} else if n%10 >= 2 && n%10 <= 9 && n%100 < 11 && n%100 > 19 {
		return locales.PluralRuleFew
	} else if f != 0 {
		return locales.PluralRuleMany
	}

	return locales.PluralRuleOther
}