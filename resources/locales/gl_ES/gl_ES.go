package gl_ES

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type gl_ES struct {
	locale string
}

// New returns a new instance of translator for the 'gl_ES' locale
func New() locales.Translator {
	return &gl_ES{
		locale: "gl_ES",
	}
}

// Locale returns the current translators string locale
func (l *gl_ES) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *gl_ES) CardinalPluralRule(num string) (locales.PluralRule, error) {

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