package sv

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type sv struct {
	locale string
}

// New returns a new instance of translator for the 'sv' locale
func New() locales.Translator {
	return &sv{
		locale: "sv",
	}
}

// Locale returns the current translators string locale
func (l *sv) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *sv) CardinalPluralRule(num string) (locales.PluralRule, error) {

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