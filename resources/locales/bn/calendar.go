package bn

import "github.com/go-playground/universal-translator"

var calendar = ut.Calendar{Formats: ut.CalendarFormats{Date: ut.CalendarDateFormat{Full: "EEEE, d MMMM, y", Long: "d MMMM, y", Medium: "d MMM, y", Short: "d/M/yy"}, Time: ut.CalendarDateFormat{Full: "h:mm:ss a zzzz", Long: "h:mm:ss a z", Medium: "h:mm:ss a", Short: "h:mm a"}, DateTime: ut.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"}}, FormatNames: ut.CalendarFormatNames{Months: ut.CalendarMonthFormatNames{Abbreviated: ut.CalendarMonthFormatNameValue{11: "নভেম্বর", 12: "ডিসেম্বর", 1: "জানুয়ারী", 2: "ফেব্রুয়ারী", 4: "এপ্রিল", 7: "জুলাই", 8: "আগস্ট", 9: "সেপ্টেম্বর", 3: "মার্চ", 5: "মে", 6: "জুন", 10: "অক্টোবর"}, Narrow: ut.CalendarMonthFormatNameValue{7: "জু", 10: "অ", 12: "ডি", 2: "ফে", 3: "মা", 4: "এ", 5: "মে", 6: "জুন", 1: "জা", 8: "আ", 9: "সে", 11: "ন"}, Short: ut.CalendarMonthFormatNameValue(nil), Wide: ut.CalendarMonthFormatNameValue{10: "অক্টোবর", 1: "জানুয়ারী", 3: "মার্চ", 4: "এপ্রিল", 6: "জুন", 7: "জুলাই", 12: "ডিসেম্বর", 2: "ফেব্রুয়ারী", 5: "মে", 8: "আগস্ট", 9: "সেপ্টেম্বর", 11: "নভেম্বর"}}, Days: ut.CalendarDayFormatNames{Abbreviated: ut.CalendarDayFormatNameValue{1: "সোম", 2: "মঙ্গল", 3: "বুধ", 4: "বৃহস্পতি", 5: "শুক্র", 6: "শনি", 0: "রবি"}, Narrow: ut.CalendarDayFormatNameValue{5: "শু", 6: "শ", 0: "র", 1: "সো", 2: "ম", 3: "বু", 4: "বৃ"}, Short: ut.CalendarDayFormatNameValue{5: "শুঃ", 6: "শনি", 0: "রঃ", 1: "সোঃ", 2: "মঃ", 3: "বুঃ", 4: "বৃঃ"}, Wide: ut.CalendarDayFormatNameValue{5: "শুক্রবার", 6: "শনিবার", 0: "রবিবার", 1: "সোমবার", 2: "মঙ্গলবার", 3: "বুধবার", 4: "বৃহষ্পতিবার"}}, Periods: ut.CalendarPeriodFormatNames{Abbreviated: ut.CalendarPeriodFormatNameValue{"afternoon2": "বিকাল", "evening1": "সন্ধ্যা", "night1": "রাত্রি", "am": "AM", "pm": "PM", "morning1": "ভোর", "morning2": "সকাল", "afternoon1": "দুপুর"}, Narrow: ut.CalendarPeriodFormatNameValue{"morning1": "ভোর", "morning2": "সকাল", "afternoon1": "দুপুর", "afternoon2": "বিকাল", "evening1": "সন্ধ্যা", "night1": "রাত্রি", "am": "AM", "pm": "PM"}, Short: ut.CalendarPeriodFormatNameValue(nil), Wide: ut.CalendarPeriodFormatNameValue{"afternoon1": "দুপুর", "afternoon2": "বিকাল", "evening1": "সন্ধ্যা", "night1": "রাত্রি", "am": "AM", "pm": "PM", "morning1": "ভোর", "morning2": "সকাল"}}}}