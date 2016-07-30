package lv_LV

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type lv_LV struct {
	locale string
}

// New returns a new instance of translator for the 'lv_LV' locale
func New() locales.Translator {
	return &lv_LV{
		locale: "lv_LV",
	}
}

// Locale returns the current translators string locale
func (l *lv_LV) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *lv_LV) CardinalPluralRule(num string) (locales.PluralRule, error) {

	n, err := locales.N(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	v := locales.V(num)

	f, err := locales.F(num)
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