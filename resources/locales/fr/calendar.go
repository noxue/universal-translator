package fr

import "github.com/go-playground/universal-translator"

var calendar = ut.Calendar{Formats: ut.CalendarFormats{Date: ut.CalendarDateFormat{Full: "EEEE d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "dd/MM/y"}, Time: ut.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: ut.CalendarDateFormat{Full: "{1} 'à' {0}", Long: "{1} 'à' {0}", Medium: "{1} 'à' {0}", Short: "{1} {0}"}}, FormatNames: ut.CalendarFormatNames{Months: ut.CalendarMonthFormatNames{Abbreviated: ut.CalendarMonthFormatNameValue{11: "nov.", 2: "févr.", 3: "mars", 5: "mai", 8: "août", 9: "sept.", 10: "oct.", 1: "janv.", 4: "avr.", 6: "juin", 7: "juil.", 12: "déc."}, Narrow: ut.CalendarMonthFormatNameValue{11: "N", 12: "D", 3: "M", 4: "A", 6: "J", 8: "A", 10: "O", 1: "J", 2: "F", 5: "M", 7: "J", 9: "S"}, Short: ut.CalendarMonthFormatNameValue(nil), Wide: ut.CalendarMonthFormatNameValue{6: "juin", 10: "octobre", 11: "novembre", 12: "décembre", 1: "janvier", 4: "avril", 5: "mai", 7: "juillet", 8: "août", 9: "septembre", 2: "février", 3: "mars"}}, Days: ut.CalendarDayFormatNames{Abbreviated: ut.CalendarDayFormatNameValue{0: "dim.", 1: "lun.", 2: "mar.", 3: "mer.", 4: "jeu.", 5: "ven.", 6: "sam."}, Narrow: ut.CalendarDayFormatNameValue{1: "L", 2: "M", 3: "M", 4: "J", 5: "V", 6: "S", 0: "D"}, Short: ut.CalendarDayFormatNameValue{5: "ve", 6: "sa", 0: "di", 1: "lu", 2: "ma", 3: "me", 4: "je"}, Wide: ut.CalendarDayFormatNameValue{5: "vendredi", 6: "samedi", 0: "dimanche", 1: "lundi", 2: "mardi", 3: "mercredi", 4: "jeudi"}}, Periods: ut.CalendarPeriodFormatNames{Abbreviated: ut.CalendarPeriodFormatNameValue{"afternoon1": "ap.m.", "evening1": "soir", "night1": "nuit", "midnight": "min.", "am": "AM", "noon": "midi", "pm": "PM", "morning1": "mat."}, Narrow: ut.CalendarPeriodFormatNameValue{"noon": "midi", "pm": "PM", "morning1": "mat.", "afternoon1": "ap.m.", "evening1": "soir", "night1": "nuit", "midnight": "min.", "am": "AM"}, Short: ut.CalendarPeriodFormatNameValue(nil), Wide: ut.CalendarPeriodFormatNameValue{"pm": "PM", "morning1": "matin", "afternoon1": "après-midi", "evening1": "soir", "night1": "nuit", "midnight": "minuit", "am": "AM", "noon": "midi"}}}}