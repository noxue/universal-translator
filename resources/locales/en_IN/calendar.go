package en_IN

import "github.com/go-playground/universal-translator"

var calendar = ut.Calendar{Formats: ut.CalendarFormats{Date: ut.CalendarDateFormat{Full: "EEEE d MMMM y", Long: "MMMM d, y", Medium: "dd-MMM-y", Short: "dd/MM/yy"}, Time: ut.CalendarDateFormat{Full: "h:mm:ss a zzzz", Long: "h:mm:ss a z", Medium: "h:mm:ss a", Short: "h:mm a"}, DateTime: ut.CalendarDateFormat{Full: "{1} 'at' {0}", Long: "{1} 'at' {0}", Medium: "{1}, {0}", Short: "{1}, {0}"}}, FormatNames: ut.CalendarFormatNames{Months: ut.CalendarMonthFormatNames{Abbreviated: ut.CalendarMonthFormatNameValue{9: "Sep", 11: "Nov", 10: "Oct", 4: "Apr", 6: "Jun", 1: "Jan", 5: "May", 2: "Feb", 3: "Mar", 7: "Jul", 8: "Aug", 12: "Dec"}, Narrow: ut.CalendarMonthFormatNameValue{3: "M", 5: "M", 10: "O", 2: "F", 8: "A", 9: "S", 12: "D", 1: "J", 6: "J", 7: "J", 11: "N", 4: "A"}, Short: ut.CalendarMonthFormatNameValue{}, Wide: ut.CalendarMonthFormatNameValue{10: "October", 3: "March", 1: "January", 4: "April", 5: "May", 9: "September", 2: "February", 12: "December", 6: "June", 7: "July", 8: "August", 11: "November"}}, Days: ut.CalendarDayFormatNames{Abbreviated: ut.CalendarDayFormatNameValue{4: "Thu", 5: "Fri", 6: "Sat", 0: "Sun", 1: "Mon", 2: "Tue", 3: "Wed"}, Narrow: ut.CalendarDayFormatNameValue{1: "M", 2: "T", 3: "W", 4: "T", 5: "F", 6: "S", 0: "S"}, Short: ut.CalendarDayFormatNameValue{2: "Tu", 3: "We", 4: "Th", 5: "Fr", 6: "Sa", 0: "Su", 1: "Mo"}, Wide: ut.CalendarDayFormatNameValue{2: "Tuesday", 3: "Wednesday", 4: "Thursday", 5: "Friday", 6: "Saturday", 0: "Sunday", 1: "Monday"}}, Periods: ut.CalendarPeriodFormatNames{Abbreviated: ut.CalendarPeriodFormatNameValue{"noon": "noon", "pm": "PM", "morning1": "in the morning", "afternoon1": "in the afternoon", "evening1": "in the evening", "night1": "at night", "midnight": "midnight", "am": "AM"}, Narrow: ut.CalendarPeriodFormatNameValue{"noon": "n", "pm": "p", "morning1": "in the morning", "afternoon1": "in the afternoon", "evening1": "in the evening", "night1": "at night", "midnight": "mi", "am": "a"}, Short: ut.CalendarPeriodFormatNameValue{}, Wide: ut.CalendarPeriodFormatNameValue{"night1": "night", "midnight": "midnight", "am": "AM", "noon": "noon", "pm": "PM", "morning1": "morning", "afternoon1": "afternoon", "evening1": "evening"}}}}