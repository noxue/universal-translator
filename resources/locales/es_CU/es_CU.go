package es_CU

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type es_CU struct {
	locale string
}

// New returns a new instance of translator for the 'es_CU' locale
func New() locales.Translator {
	return &es_CU{
		locale: "es_CU",
	}
}

// Locale returns the current translators string locale
func (l *es_CU) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *es_CU) CardinalPluralRule(num string) (locales.PluralRule, error) {

	n, err := locales.N(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	if n == 1 {
		return locales.PluralRuleOne, nil
	}

	return locales.PluralRuleOther, nil
}