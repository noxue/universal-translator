package rn

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type rn struct {
	locale string
}

// New returns a new instance of translator for the 'rn' locale
func New() locales.Translator {
	return &rn{
		locale: "rn",
	}
}

// Locale returns the current translators string locale
func (l *rn) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *rn) CardinalPluralRule(num string) (locales.PluralRule, error) {

}