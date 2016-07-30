package mgo_CM

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type mgo_CM struct {
	locale string
}

// New returns a new instance of translator for the 'mgo_CM' locale
func New() locales.Translator {
	return &mgo_CM{
		locale: "mgo_CM",
	}
}

// Locale returns the current translators string locale
func (l *mgo_CM) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *mgo_CM) CardinalPluralRule(num string) (locales.PluralRule, error) {

	n, err := locales.N(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	if n == 1 {
		return locales.PluralRuleOne, nil
	}

	return locales.PluralRuleOther, nil
}