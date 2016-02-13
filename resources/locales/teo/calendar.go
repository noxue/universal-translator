package teo

import "github.com/go-playground/universal-translator"

var calendar = ut.Calendar{Formats: ut.CalendarFormats{Date: ut.CalendarDateFormat{Full: "EEEE, d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "dd/MM/y"}, Time: ut.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: ut.CalendarDateFormat{Full: "", Long: "", Medium: "", Short: ""}}, FormatNames: ut.CalendarFormatNames{Months: ut.CalendarMonthFormatNames{Abbreviated: ut.CalendarMonthFormatNameValue{12: "Poo", 1: "Rar", 7: "Jol", 10: "Tib", 5: "Mar", 6: "Mod", 8: "Ped", 9: "Sok", 11: "Lab", 2: "Muk", 3: "Kwa", 4: "Dun"}, Narrow: ut.CalendarMonthFormatNameValue{3: "K", 5: "M", 6: "M", 8: "P", 9: "S", 10: "T", 2: "M", 4: "D", 7: "J", 11: "L", 12: "P", 1: "R"}, Short: ut.CalendarMonthFormatNameValue(nil), Wide: ut.CalendarMonthFormatNameValue{10: "Otibar", 11: "Olabor", 1: "Orara", 4: "Odung’el", 6: "Omodok’king’ol", 8: "Opedel", 9: "Osokosokoma", 12: "Opoo", 2: "Omuk", 3: "Okwamg’", 5: "Omaruk", 7: "Ojola"}}, Days: ut.CalendarDayFormatNames{Abbreviated: ut.CalendarDayFormatNameValue{5: "Kan", 6: "Sab", 0: "Jum", 1: "Bar", 2: "Aar", 3: "Uni", 4: "Ung"}, Narrow: ut.CalendarDayFormatNameValue{2: "A", 3: "U", 4: "U", 5: "K", 6: "S", 0: "J", 1: "B"}, Short: ut.CalendarDayFormatNameValue(nil), Wide: ut.CalendarDayFormatNameValue{3: "Nakauni", 4: "Nakaung’on", 5: "Nakakany", 6: "Nakasabiti", 0: "Nakaejuma", 1: "Nakaebarasa", 2: "Nakaare"}}, Periods: ut.CalendarPeriodFormatNames{Abbreviated: ut.CalendarPeriodFormatNameValue(nil), Narrow: ut.CalendarPeriodFormatNameValue(nil), Short: ut.CalendarPeriodFormatNameValue(nil), Wide: ut.CalendarPeriodFormatNameValue{"pm": "Ebongi", "am": "Taparachu"}}}}