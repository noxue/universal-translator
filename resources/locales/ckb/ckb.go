package ckb

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type ckb struct {
	locale string
}

// New returns a new instance of translator for the 'ckb' locale
func New() locales.Translator {
	return &ckb{
		locale: "ckb",
	}
}

// Locale returns the current translators string locale
func (l *ckb) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *ckb) CardinalPluralRule(num string) (locales.PluralRule, error) {

	n, err := locales.N(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	if n == 1 {
		return locales.PluralRuleOne, nil
	}

	return locales.PluralRuleOther, nil
}