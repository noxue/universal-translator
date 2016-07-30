package en_FI

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type en_FI struct {
	locale string
}

// New returns a new instance of translator for the 'en_FI' locale
func New() locales.Translator {
	return &en_FI{
		locale: "en_FI",
	}
}

// Locale returns the current translators string locale
func (l *en_FI) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *en_FI) CardinalPluralRule(num string) (locales.PluralRule, error) {

	v := locales.V(num)

	i, err := locales.I(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	if i == 1 && v == 0 {
		return locales.PluralRuleOne, nil
	}

	return locales.PluralRuleOther, nil
}