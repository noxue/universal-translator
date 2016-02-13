package ta_MY

import "github.com/go-playground/universal-translator"

var calendar = ut.Calendar{Formats: ut.CalendarFormats{Date: ut.CalendarDateFormat{Full: "EEEE, d MMMM, y", Long: "d MMMM, y", Medium: "d MMM, y", Short: "d/M/yy"}, Time: ut.CalendarDateFormat{Full: "a h:mm:ss zzzz", Long: "a h:mm:ss z", Medium: "a h:mm:ss", Short: "a h:mm"}, DateTime: ut.CalendarDateFormat{Full: "{1} ’அன்று’ {0}", Long: "{1} ’அன்று’ {0}", Medium: "{1}, {0}", Short: "{1}, {0}"}}, FormatNames: ut.CalendarFormatNames{Months: ut.CalendarMonthFormatNames{Abbreviated: ut.CalendarMonthFormatNameValue{8: "ஆக.", 9: "செப்.", 10: "அக்.", 7: "ஜூலை", 2: "பிப்.", 1: "ஜன.", 4: "ஏப்.", 5: "மே", 6: "ஜூன்", 11: "நவ.", 12: "டிச.", 3: "மார்."}, Narrow: ut.CalendarMonthFormatNameValue{2: "பி", 6: "ஜூ", 7: "ஜூ", 9: "செ", 11: "ந", 5: "மே", 8: "ஆ", 12: "டி", 3: "மா", 4: "ஏ", 10: "அ", 1: "ஜ"}, Short: ut.CalendarMonthFormatNameValue{}, Wide: ut.CalendarMonthFormatNameValue{1: "ஜனவரி", 2: "பிப்ரவரி", 4: "ஏப்ரல்", 12: "டிசம்பர்", 3: "மார்ச்", 7: "ஜூலை", 6: "ஜூன்", 8: "ஆகஸ்டு", 9: "செப்டம்பர்", 10: "அக்டோபர்", 11: "நவம்பர்", 5: "மே"}}, Days: ut.CalendarDayFormatNames{Abbreviated: ut.CalendarDayFormatNameValue{4: "வியா.", 5: "வெள்.", 6: "சனி", 0: "ஞாயி.", 1: "திங்.", 2: "செவ்.", 3: "புத."}, Narrow: ut.CalendarDayFormatNameValue{6: "ச", 0: "ஞா", 1: "தி", 2: "செ", 3: "பு", 4: "வி", 5: "வெ"}, Short: ut.CalendarDayFormatNameValue{4: "வி", 5: "வெ", 6: "ச", 0: "ஞா", 1: "தி", 2: "செ", 3: "பு"}, Wide: ut.CalendarDayFormatNameValue{3: "புதன்", 4: "வியாழன்", 5: "வெள்ளி", 6: "சனி", 0: "ஞாயிறு", 1: "திங்கள்", 2: "செவ்வாய்"}}, Periods: ut.CalendarPeriodFormatNames{Abbreviated: ut.CalendarPeriodFormatNameValue{"midnight": "நள்ளிரவு", "am": "முற்பகல்", "afternoon1": "மதியம்", "evening1": "மாலை", "evening2": "அந்தி மாலை", "noon": "நண்பகல்", "pm": "பிற்பகல்", "night1": "இரவு", "morning1": "அதிகாலை", "morning2": "காலை", "afternoon2": "பிற்பகல்"}, Narrow: ut.CalendarPeriodFormatNameValue{"evening1": "மா.", "am": "மு.ப", "morning2": "கா.", "noon": "நண்.", "morning1": "அதி.", "afternoon2": "பிற்.", "night1": "இ.", "midnight": "நள்.", "pm": "பி.ப", "afternoon1": "மதி.", "evening2": "அந்தி மா."}, Short: ut.CalendarPeriodFormatNameValue{}, Wide: ut.CalendarPeriodFormatNameValue{"noon": "நண்பகல்", "afternoon1": "மதியம்", "am": "முற்பகல்", "pm": "பிற்பகல்", "morning1": "அதிகாலை", "evening2": "அந்தி மாலை", "night1": "இரவு", "morning2": "காலை", "afternoon2": "பிற்பகல்", "evening1": "மாலை", "midnight": "நள்ளிரவு"}}}}