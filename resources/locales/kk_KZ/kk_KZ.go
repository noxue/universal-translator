package kk_KZ

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type kk_KZ struct {
	locale string
}

// New returns a new instance of translator for the 'kk_KZ' locale
func New() locales.Translator {
	return &kk_KZ{
		locale: "kk_KZ",
	}
}

// Locale returns the current translators string locale
func (l *kk_KZ) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *kk_KZ) CardinalPluralRule(num string) (locales.PluralRule, error) {

	n, err := locales.N(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	if n == 1 {
		return locales.PluralRuleOne, nil
	}

	return locales.PluralRuleOther, nil
}