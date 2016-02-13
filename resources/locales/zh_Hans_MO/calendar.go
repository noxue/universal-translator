package zh_Hans_MO

import "github.com/go-playground/universal-translator"

var calendar = ut.Calendar{Formats: ut.CalendarFormats{Date: ut.CalendarDateFormat{Full: "y年M月d日EEEE", Long: "y年M月d日", Medium: "y年M月d日", Short: "d/M/yy"}, Time: ut.CalendarDateFormat{Full: "zzzz ah:mm:ss", Long: "z ah:mm:ss", Medium: "ah:mm:ss", Short: "ah:mm"}, DateTime: ut.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"}}, FormatNames: ut.CalendarFormatNames{Months: ut.CalendarMonthFormatNames{Abbreviated: ut.CalendarMonthFormatNameValue{10: "10月", 5: "5月", 12: "12月", 1: "1月", 4: "4月", 7: "7月", 8: "8月", 9: "9月", 11: "11月", 2: "2月", 3: "3月", 6: "6月"}, Narrow: ut.CalendarMonthFormatNameValue{12: "12", 2: "2", 5: "5", 7: "7", 10: "10", 3: "3", 9: "9", 11: "11", 6: "6", 8: "8", 1: "1", 4: "4"}, Short: ut.CalendarMonthFormatNameValue{}, Wide: ut.CalendarMonthFormatNameValue{10: "十月", 12: "十二月", 7: "七月", 8: "八月", 1: "一月", 4: "四月", 5: "五月", 9: "九月", 2: "二月", 3: "三月", 6: "六月", 11: "十一月"}}, Days: ut.CalendarDayFormatNames{Abbreviated: ut.CalendarDayFormatNameValue{0: "周日", 1: "周一", 2: "周二", 3: "周三", 4: "周四", 5: "周五", 6: "周六"}, Narrow: ut.CalendarDayFormatNameValue{1: "一", 2: "二", 3: "三", 4: "四", 5: "五", 6: "六", 0: "日"}, Short: ut.CalendarDayFormatNameValue{5: "周五", 6: "周六", 0: "周日", 1: "周一", 2: "周二", 3: "周三", 4: "周四"}, Wide: ut.CalendarDayFormatNameValue{4: "星期四", 5: "星期五", 6: "星期六", 0: "星期日", 1: "星期一", 2: "星期二", 3: "星期三"}}, Periods: ut.CalendarPeriodFormatNames{Abbreviated: ut.CalendarPeriodFormatNameValue{"morning2": "上午", "morning1": "早上", "evening1": "晚上", "am": "上午", "night1": "凌晨", "afternoon1": "中午", "afternoon2": "下午", "midnight": "午夜", "pm": "下午"}, Narrow: ut.CalendarPeriodFormatNameValue{"pm": "下午", "morning2": "上午", "afternoon1": "中午", "afternoon2": "下午", "evening1": "晚上", "am": "上午", "morning1": "早上", "night1": "凌晨", "midnight": "午夜"}, Short: ut.CalendarPeriodFormatNameValue{}, Wide: ut.CalendarPeriodFormatNameValue{"afternoon2": "下午", "midnight": "午夜", "evening1": "晚上", "night1": "凌晨", "morning1": "早上", "morning2": "上午", "afternoon1": "中午", "am": "上午", "pm": "下午"}}}}