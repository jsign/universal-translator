package ur_IN

import "github.com/go-playground/universal-translator"

var calendar = ut.Calendar{Formats: ut.CalendarFormats{DateEra: ut.DateEra{BC: ut.CalendarDateFormat{Full: "EEEE، d MMMM، y", Long: "d MMMM، y", Medium: "d MMM، y", Short: "d/M/yy GGGGG"}, AD: ut.CalendarDateFormat{Full: "EEEE، d MMMM، y", Long: "d MMMM، y", Medium: "d MMM، y", Short: "d/M/yy GGGGG"}}, Time: ut.CalendarDateFormat{Full: "h:mm:ss a zzzz", Long: "h:mm:ss a z", Medium: "h:mm:ss a", Short: "h:mm a"}, DateTime: ut.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"}}, FormatNames: ut.CalendarFormatNames{Months: ut.CalendarMonthFormatNames{Abbreviated: ut.CalendarMonthFormatNameValue{11: "نومبر", 12: "دسمبر", 9: "ستمبر", 5: "مئی", 4: "اپریل", 3: "مارچ", 8: "اگست", 1: "جنوری", 2: "فروری", 6: "جون", 7: "جولائی", 10: "اکتوبر"}, Narrow: ut.CalendarMonthFormatNameValue{5: "M", 6: "J", 9: "S", 11: "N", 4: "A", 7: "J", 10: "O", 1: "J", 2: "F", 12: "D", 8: "A", 3: "M"}, Short: ut.CalendarMonthFormatNameValue{}, Wide: ut.CalendarMonthFormatNameValue{5: "مئی", 7: "جولائی", 10: "اکتوبر", 12: "دسمبر", 1: "جنوری", 2: "فروری", 8: "اگست", 11: "نومبر", 9: "ستمبر", 3: "مارچ", 4: "اپریل", 6: "جون"}}, Days: ut.CalendarDayFormatNames{Abbreviated: ut.CalendarDayFormatNameValue{3: "بدھ", 4: "جمعرات", 5: "جمعہ", 6: "ہفتہ", 0: "اتوار", 1: "سوموار", 2: "منگل"}, Narrow: ut.CalendarDayFormatNameValue{2: "T", 3: "W", 4: "T", 5: "F", 6: "S", 0: "S", 1: "M"}, Short: ut.CalendarDayFormatNameValue{0: "اتوار", 1: "سوموار", 2: "منگل", 3: "بدھ", 4: "جمعرات", 5: "جمعہ", 6: "ہفتہ"}, Wide: ut.CalendarDayFormatNameValue{1: "سوموار", 2: "منگل", 3: "بدھ", 4: "جمعرات", 5: "جمعہ", 6: "ہفتہ", 0: "اتوار"}}, Periods: ut.CalendarPeriodFormatNames{Abbreviated: ut.CalendarPeriodFormatNameValue{"afternoon1": "دوپہر", "afternoon2": "سہ پہر", "evening1": "شام", "night1": "رات", "midnight": "آدھی رات", "am": "AM", "pm": "PM", "morning1": "صبح"}, Narrow: ut.CalendarPeriodFormatNameValue{"evening1": "شام", "night1": "رات", "midnight": "آدھی رات", "am": "a", "pm": "p", "morning1": "صبح", "afternoon1": "دوپہر", "afternoon2": "سہ پہر"}, Short: ut.CalendarPeriodFormatNameValue{}, Wide: ut.CalendarPeriodFormatNameValue{"afternoon1": "دوپہر", "afternoon2": "سہ پہر", "evening1": "شام", "night1": "رات", "midnight": "آدھی رات", "am": "AM", "pm": "PM", "morning1": "صبح"}}, Eras: ut.Eras{AD: ut.CalendarEraFormatNames{Full: "", Abbrev: "", Narrow: ""}, BC: ut.CalendarEraFormatNames{Full: "", Abbrev: "", Narrow: ""}}}}