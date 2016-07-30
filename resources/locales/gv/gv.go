package gv

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type gv struct {
	locale string
}

// New returns a new instance of translator for the 'gv' locale
func New() locales.Translator {
	return &gv{
		locale: "gv",
	}
}

// Locale returns the current translators string locale
func (l *gv) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *gv) CardinalPluralRule(num string) (locales.PluralRule, error) {

	i, err := locales.I(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	v := locales.V(num)

	if v == 0 && i%10 == 1 {
		return locales.PluralRuleOne, nil
	} else if v == 0 && i%10 == 2 {
		return locales.PluralRuleTwo, nil
	} else if v == 0 && (i%100 == 0 || i%100 == 20 || i%100 == 40 || i%100 == 60 || i%100 == 80) {
		return locales.PluralRuleFew, nil
	} else if v != 0 {
		return locales.PluralRuleMany, nil
	}

	return locales.PluralRuleOther, nil
}