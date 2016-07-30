package fy

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type fy struct {
	locale string
}

// New returns a new instance of translator for the 'fy' locale
func New() locales.Translator {
	return &fy{
		locale: "fy",
	}
}

// Locale returns the current translators string locale
func (l *fy) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *fy) CardinalPluralRule(num string) (locales.PluralRule, error) {

	i, err := locales.I(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	v := locales.V(num)

	if i == 1 && v == 0 {
		return locales.PluralRuleOne, nil
	}

	return locales.PluralRuleOther, nil
}