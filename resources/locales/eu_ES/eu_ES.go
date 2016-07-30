package eu_ES

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type eu_ES struct {
	locale string
}

// New returns a new instance of translator for the 'eu_ES' locale
func New() locales.Translator {
	return &eu_ES{
		locale: "eu_ES",
	}
}

// Locale returns the current translators string locale
func (l *eu_ES) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *eu_ES) CardinalPluralRule(num string) (locales.PluralRule, error) {

	n, err := locales.N(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	if n == 1 {
		return locales.PluralRuleOne, nil
	}

	return locales.PluralRuleOther, nil
}