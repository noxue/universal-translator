package lt_LT

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type lt_LT struct {
	locale string
}

// New returns a new instance of translator for the 'lt_LT' locale
func New() locales.Translator {
	return &lt_LT{
		locale: "lt_LT",
	}
}

// Locale returns the current translators string locale
func (l *lt_LT) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *lt_LT) CardinalPluralRule(num string) (locales.PluralRule, error) {

	n, err := locales.N(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	f, err := locales.F(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	if n%10 == 1 && n%100 < 11 && n%100 > 19 {
		return locales.PluralRuleOne, nil
	} else if n%10 >= 2 && n%10 <= 9 && n%100 < 11 && n%100 > 19 {
		return locales.PluralRuleFew, nil
	} else if f != 0 {
		return locales.PluralRuleMany, nil
	}

	return locales.PluralRuleOther, nil
}