package tr

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type tr struct {
	locale string
}

// New returns a new instance of translator for the 'tr' locale
func New() locales.Translator {
	return &tr{
		locale: "tr",
	}
}

// Locale returns the current translators string locale
func (l *tr) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *tr) CardinalPluralRule(num string) (locales.PluralRule, error) {

	n, err := locales.N(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	if n == 1 {
		return locales.PluralRuleOne, nil
	}

	return locales.PluralRuleOther, nil
}