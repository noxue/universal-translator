package ne

import "github.com/go-playground/universal-translator"

var calendar = ut.Calendar{Formats: ut.CalendarFormats{Date: ut.CalendarDateFormat{Full: "y MMMM d, EEEE", Long: "y MMMM d", Medium: "y MMM d", Short: "y-MM-dd"}, Time: ut.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: ut.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1}, {0}", Short: "{1}, {0}"}}, FormatNames: ut.CalendarFormatNames{Months: ut.CalendarMonthFormatNames{Abbreviated: ut.CalendarMonthFormatNameValue{6: "जुन", 7: "जुलाई", 8: "अगस्ट", 11: "नोभेम्बर", 12: "डिसेम्बर", 1: "जनवरी", 2: "फेब्रुअरी", 3: "मार्च", 10: "अक्टोबर", 4: "अप्रिल", 5: "मे", 9: "सेप्टेम्बर"}, Narrow: ut.CalendarMonthFormatNameValue{7: "७", 8: "८", 10: "१०", 11: "११", 12: "१२", 1: "१", 2: "२", 3: "३", 4: "४", 5: "५", 6: "६", 9: "९"}, Short: ut.CalendarMonthFormatNameValue(nil), Wide: ut.CalendarMonthFormatNameValue{1: "जनवरी", 3: "मार्च", 5: "मे", 7: "जुलाई", 8: "अगस्ट", 12: "डिसेम्बर", 2: "फेब्रुअरी", 4: "अप्रिल", 6: "जुन", 9: "सेप्टेम्बर", 10: "अक्टोबर", 11: "नोभेम्बर"}}, Days: ut.CalendarDayFormatNames{Abbreviated: ut.CalendarDayFormatNameValue{5: "शुक्र", 6: "शनि", 0: "आइत", 1: "सोम", 2: "मङ्गल", 3: "बुध", 4: "बिही"}, Narrow: ut.CalendarDayFormatNameValue{2: "म", 3: "बु", 4: "बि", 5: "शु", 6: "श", 0: "आ", 1: "सो"}, Short: ut.CalendarDayFormatNameValue{3: "बुध", 4: "बिही", 5: "शुक्र", 6: "शनि", 0: "आइत", 1: "सोम", 2: "मङ्गल"}, Wide: ut.CalendarDayFormatNameValue{6: "शनिबार", 0: "आइतबार", 1: "सोमबार", 2: "मङ्गलबार", 3: "बुधबार", 4: "बिहिबार", 5: "शुक्रबार"}}, Periods: ut.CalendarPeriodFormatNames{Abbreviated: ut.CalendarPeriodFormatNameValue{"afternoon1": "अपरान्ह", "afternoon2": "बेलुका", "evening1": "साँझ", "night1": "रात", "midnight": "मध्यरात", "noon": "मध्यान्ह", "morning1": "बिहान"}, Narrow: ut.CalendarPeriodFormatNameValue{"evening1": "बेलुका", "night1": "रात", "am": "पूर्वाह्न", "pm": "अपराह्न", "morning1": "बिहान", "afternoon2": "साँझ", "midnight": "मध्यरात", "noon": "मध्यान्ह", "afternoon1": "अपरान्ह"}, Short: ut.CalendarPeriodFormatNameValue(nil), Wide: ut.CalendarPeriodFormatNameValue{"midnight": "मध्यरात", "pm": "अपराह्न", "morning1": "विहान", "am": "पूर्वाह्न", "noon": "मध्यान्ह", "afternoon1": "अपरान्ह", "afternoon2": "साँझ", "evening1": "बेलुका", "night1": "रात"}}}}