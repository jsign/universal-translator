package ur

import "github.com/go-playground/universal-translator"

var calendar = ut.Calendar{Formats: ut.CalendarFormats{DateEra: ut.DateEra{BC: ut.CalendarDateFormat{Full: "EEEE، d MMMM، y", Long: "d MMMM، y", Medium: "d MMM، y", Short: "d/M/yy"}, AD: ut.CalendarDateFormat{Full: "EEEE، d MMMM، y", Long: "d MMMM، y", Medium: "d MMM، y", Short: "d/M/yy"}}, Time: ut.CalendarDateFormat{Full: "h:mm:ss a zzzz", Long: "h:mm:ss a z", Medium: "h:mm:ss a", Short: "h:mm a"}, DateTime: ut.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"}}, FormatNames: ut.CalendarFormatNames{Months: ut.CalendarMonthFormatNames{Abbreviated: ut.CalendarMonthFormatNameValue{12: "دسمبر", 1: "جنوری", 2: "فروری", 4: "اپریل", 8: "اگست", 11: "نومبر", 10: "اکتوبر", 3: "مارچ", 5: "مئی", 6: "جون", 7: "جولائی", 9: "ستمبر"}, Narrow: ut.CalendarMonthFormatNameValue{8: "A", 3: "M", 4: "A", 5: "M", 6: "J", 7: "J", 9: "S", 10: "O", 11: "N", 1: "J", 2: "F", 12: "D"}, Short: ut.CalendarMonthFormatNameValue(nil), Wide: ut.CalendarMonthFormatNameValue{4: "اپریل", 6: "جون", 9: "ستمبر", 12: "دسمبر", 1: "جنوری", 2: "فروری", 3: "مارچ", 10: "اکتوبر", 11: "نومبر", 5: "مئی", 7: "جولائی", 8: "اگست"}}, Days: ut.CalendarDayFormatNames{Abbreviated: ut.CalendarDayFormatNameValue{1: "سوموار", 2: "منگل", 3: "بدھ", 4: "جمعرات", 5: "جمعہ", 6: "ہفتہ", 0: "اتوار"}, Narrow: ut.CalendarDayFormatNameValue{2: "T", 3: "W", 4: "T", 5: "F", 6: "S", 0: "S", 1: "M"}, Short: ut.CalendarDayFormatNameValue{3: "بدھ", 4: "جمعرات", 5: "جمعہ", 6: "ہفتہ", 0: "اتوار", 1: "سوموار", 2: "منگل"}, Wide: ut.CalendarDayFormatNameValue{1: "سوموار", 2: "منگل", 3: "بدھ", 4: "جمعرات", 5: "جمعہ", 6: "ہفتہ", 0: "اتوار"}}, Periods: ut.CalendarPeriodFormatNames{Abbreviated: ut.CalendarPeriodFormatNameValue{"afternoon2": "سہ پہر", "evening1": "شام", "night1": "رات", "midnight": "آدھی رات", "am": "AM", "pm": "PM", "morning1": "صبح", "afternoon1": "دوپہر"}, Narrow: ut.CalendarPeriodFormatNameValue{"pm": "p", "morning1": "صبح", "afternoon1": "دوپہر", "afternoon2": "سہ پہر", "evening1": "شام", "night1": "رات", "midnight": "آدھی رات", "am": "a"}, Short: ut.CalendarPeriodFormatNameValue(nil), Wide: ut.CalendarPeriodFormatNameValue{"morning1": "صبح", "afternoon1": "دوپہر", "afternoon2": "سہ پہر", "evening1": "شام", "night1": "رات", "midnight": "آدھی رات", "am": "AM", "pm": "PM"}}, Eras: ut.Eras{AD: ut.CalendarEraFormatNames{Full: "", Abbrev: "", Narrow: ""}, BC: ut.CalendarEraFormatNames{Full: "قبل مسیح", Abbrev: "قبل مسیح", Narrow: ""}}}}
