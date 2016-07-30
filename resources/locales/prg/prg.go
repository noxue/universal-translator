package prg

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type prg struct {
	locale string
}

// New returns a new instance of translator for the 'prg' locale
func New() locales.Translator {
	return &prg{
		locale: "prg",
	}
}

// Locale returns the current translators string locale
func (l *prg) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *prg) CardinalPluralRule(num string) (locales.PluralRule, error) {

	v := locales.V(num)

	f, err := locales.F(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	n, err := locales.N(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	if (n%10 == 0) || (n%100 >= 11 && n%100 <= 19) || (v == 2 && f%100 >= 11 && f%100 <= 19) {
		return locales.PluralRuleZero, nil
	} else if (n%10 == 1 && n%100 != 11) || (v == 2 && f%10 == 1 && f%100 != 11) || (v != 2 && f%10 == 1) {
		return locales.PluralRuleOne, nil
	}

	return locales.PluralRuleOther, nil
}