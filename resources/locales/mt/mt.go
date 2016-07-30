package mt

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type mt struct {
	locale string
}

// New returns a new instance of translator for the 'mt' locale
func New() locales.Translator {
	return &mt{
		locale: "mt",
	}
}

// Locale returns the current translators string locale
func (l *mt) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *mt) CardinalPluralRule(num string) (locales.PluralRule, error) {

	n, err := locales.N(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	if n == 1 {
		return locales.PluralRuleOne, nil
	} else if (n == 0) || (n%100 >= 2 && n%100 <= 10) {
		return locales.PluralRuleFew, nil
	} else if n%100 >= 11 && n%100 <= 19 {
		return locales.PluralRuleMany, nil
	}

	return locales.PluralRuleOther, nil
}