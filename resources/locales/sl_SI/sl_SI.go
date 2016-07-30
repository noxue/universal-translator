package sl_SI

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type sl_SI struct {
	locale string
}

// New returns a new instance of translator for the 'sl_SI' locale
func New() locales.Translator {
	return &sl_SI{
		locale: "sl_SI",
	}
}

// Locale returns the current translators string locale
func (l *sl_SI) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *sl_SI) CardinalPluralRule(num string) (locales.PluralRule, error) {

	i, err := locales.I(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	v := locales.V(num)

	if v == 0 && i%100 == 1 {
		return locales.PluralRuleOne, nil
	} else if v == 0 && i%100 == 2 {
		return locales.PluralRuleTwo, nil
	} else if (v == 0 && i%100 >= 3 && i%100 <= 4) || (v != 0) {
		return locales.PluralRuleFew, nil
	}

	return locales.PluralRuleOther, nil
}