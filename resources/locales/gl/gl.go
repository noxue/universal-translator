package gl

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type gl struct {
	locale string
}

// New returns a new instance of translator for the 'gl' locale
func New() locales.Translator {
	return &gl{
		locale: "gl",
	}
}

// Locale returns the current translators string locale
func (l *gl) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *gl) CardinalPluralRule(num string) (locales.PluralRule, error) {

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