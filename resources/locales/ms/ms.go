package ms

import "github.com/go-playground/universal-translator/resources/locales"

type ms struct {
	locale   string
	plurals  []locales.PluralRule
	decimal  []byte
	group    []byte
	minus    []byte
	percent  []byte
	perMille []byte
	symbol   []byte
}

// New returns a new instance of translator for the 'ms' locale
func New() locales.Translator {
	return &ms{
		locale:   "ms",
		plurals:  []locales.PluralRule{6},
		decimal:  []byte{0x2e},
		group:    []byte{0x2c},
		minus:    []byte{0x2d},
		percent:  []byte{0x25},
		perMille: []byte{0xe2, 0x80, 0xb0},
		symbol:   []byte{},
	}
}

// Locale returns the current translators string locale
func (t *ms) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'ms'
func (t *ms) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'ms'
func (t *ms) cardinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleOther
}