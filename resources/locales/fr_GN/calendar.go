package fr_GN

import "github.com/go-playground/universal-translator"

var calendar = ut.Calendar{Formats: ut.CalendarFormats{Date: ut.CalendarDateFormat{Full: "EEEE d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "dd/MM/y"}, Time: ut.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: ut.CalendarDateFormat{Full: "{1} 'à' {0}", Long: "{1} 'à' {0}", Medium: "{1} 'à' {0}", Short: "{1} {0}"}}, FormatNames: ut.CalendarFormatNames{Months: ut.CalendarMonthFormatNames{Abbreviated: ut.CalendarMonthFormatNameValue{9: "sept.", 10: "oct.", 11: "nov.", 6: "juin", 7: "juil.", 12: "déc.", 1: "janv.", 5: "mai", 8: "août", 2: "févr.", 4: "avr.", 3: "mars"}, Narrow: ut.CalendarMonthFormatNameValue{4: "A", 6: "J", 9: "S", 1: "J", 5: "M", 8: "A", 12: "D", 3: "M", 2: "F", 10: "O", 11: "N", 7: "J"}, Short: ut.CalendarMonthFormatNameValue{}, Wide: ut.CalendarMonthFormatNameValue{1: "janvier", 4: "avril", 10: "octobre", 3: "mars", 8: "août", 9: "septembre", 6: "juin", 11: "novembre", 12: "décembre", 2: "février", 5: "mai", 7: "juillet"}}, Days: ut.CalendarDayFormatNames{Abbreviated: ut.CalendarDayFormatNameValue{4: "jeu.", 5: "ven.", 6: "sam.", 0: "dim.", 1: "lun.", 2: "mar.", 3: "mer."}, Narrow: ut.CalendarDayFormatNameValue{0: "D", 1: "L", 2: "M", 3: "M", 4: "J", 5: "V", 6: "S"}, Short: ut.CalendarDayFormatNameValue{6: "sa", 0: "di", 1: "lu", 2: "ma", 3: "me", 4: "je", 5: "ve"}, Wide: ut.CalendarDayFormatNameValue{3: "mercredi", 4: "jeudi", 5: "vendredi", 6: "samedi", 0: "dimanche", 1: "lundi", 2: "mardi"}}, Periods: ut.CalendarPeriodFormatNames{Abbreviated: ut.CalendarPeriodFormatNameValue{"afternoon1": "ap.m.", "evening1": "soir", "night1": "nuit", "midnight": "min.", "am": "AM", "noon": "midi", "pm": "PM", "morning1": "mat."}, Narrow: ut.CalendarPeriodFormatNameValue{"noon": "midi", "pm": "PM", "morning1": "mat.", "afternoon1": "ap.m.", "evening1": "soir", "night1": "nuit", "midnight": "min.", "am": "AM"}, Short: ut.CalendarPeriodFormatNameValue{}, Wide: ut.CalendarPeriodFormatNameValue{"pm": "PM", "morning1": "matin", "afternoon1": "après-midi", "evening1": "soir", "night1": "nuit", "midnight": "minuit", "am": "AM", "noon": "midi"}}}}