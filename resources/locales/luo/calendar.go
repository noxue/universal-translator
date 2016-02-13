package luo

import "github.com/go-playground/universal-translator"

var calendar = ut.Calendar{Formats: ut.CalendarFormats{Date: ut.CalendarDateFormat{Full: "EEEE, d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "dd/MM/y"}, Time: ut.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: ut.CalendarDateFormat{Full: "", Long: "", Medium: "", Short: ""}}, FormatNames: ut.CalendarFormatNames{Months: ut.CalendarMonthFormatNames{Abbreviated: ut.CalendarMonthFormatNameValue{2: "DAR", 3: "DAD", 7: "DAO", 10: "DAP", 9: "DOC", 11: "DGI", 12: "DAG", 1: "DAC", 4: "DAN", 5: "DAH", 6: "DAU", 8: "DAB"}, Narrow: ut.CalendarMonthFormatNameValue{4: "N", 5: "B", 7: "B", 8: "B", 9: "C", 10: "P", 2: "R", 3: "D", 11: "C", 12: "P", 1: "C", 6: "U"}, Short: ut.CalendarMonthFormatNameValue(nil), Wide: ut.CalendarMonthFormatNameValue{9: "Dwe mar Ochiko", 12: "Dwe mar Apar gi ariyo", 2: "Dwe mar Ariyo", 3: "Dwe mar Adek", 5: "Dwe mar Abich", 7: "Dwe mar Abiriyo", 8: "Dwe mar Aboro", 10: "Dwe mar Apar", 11: "Dwe mar gi achiel", 1: "Dwe mar Achiel", 4: "Dwe mar Ang’wen", 6: "Dwe mar Auchiel"}}, Days: ut.CalendarDayFormatNames{Abbreviated: ut.CalendarDayFormatNameValue{6: "NGS", 0: "JMP", 1: "WUT", 2: "TAR", 3: "TAD", 4: "TAN", 5: "TAB"}, Narrow: ut.CalendarDayFormatNameValue{3: "T", 4: "T", 5: "T", 6: "N", 0: "J", 1: "W", 2: "T"}, Short: ut.CalendarDayFormatNameValue(nil), Wide: ut.CalendarDayFormatNameValue{6: "Ngeso", 0: "Jumapil", 1: "Wuok Tich", 2: "Tich Ariyo", 3: "Tich Adek", 4: "Tich Ang’wen", 5: "Tich Abich"}}, Periods: ut.CalendarPeriodFormatNames{Abbreviated: ut.CalendarPeriodFormatNameValue(nil), Narrow: ut.CalendarPeriodFormatNameValue(nil), Short: ut.CalendarPeriodFormatNameValue(nil), Wide: ut.CalendarPeriodFormatNameValue{"am": "OD", "pm": "OT"}}}}