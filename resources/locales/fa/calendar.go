package fa

import "github.com/go-playground/universal-translator"

var calendar = ut.Calendar{Formats: ut.CalendarFormats{DateEra: ut.DateEra{BC: ut.CalendarDateFormat{Full: "EEEE d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "y/M/d"}, AD: ut.CalendarDateFormat{Full: "EEEE d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "y/M/d"}}, Time: ut.CalendarDateFormat{Full: "H:mm:ss (zzzz)", Long: "H:mm:ss (z)", Medium: "H:mm:ss", Short: "H:mm"}, DateTime: ut.CalendarDateFormat{Full: "{1}، ساعت {0}", Long: "{1}، ساعت {0}", Medium: "{1}،\u200f {0}", Short: "{1}،\u200f {0}"}}, FormatNames: ut.CalendarFormatNames{Months: ut.CalendarMonthFormatNames{Abbreviated: ut.CalendarMonthFormatNameValue{10: "اکتبر", 11: "نوامبر", 1: "ژانویه", 2: "فوریه", 4: "آوریل", 6: "ژوئن", 7: "ژوئیه", 9: "سپتامبر", 12: "دسامبر", 3: "مارس", 5: "مه", 8: "اوت"}, Narrow: ut.CalendarMonthFormatNameValue{11: "ن", 1: "ژ", 3: "م", 4: "آ", 9: "س", 10: "ا", 12: "د", 2: "ف", 5: "م", 6: "ژ", 7: "ژ", 8: "ا"}, Short: ut.CalendarMonthFormatNameValue(nil), Wide: ut.CalendarMonthFormatNameValue{8: "اوت", 10: "اکتبر", 11: "نوامبر", 1: "ژانویه", 2: "فوریه", 4: "آوریل", 5: "مه", 6: "ژوئن", 12: "دسامبر", 3: "مارس", 7: "ژوئیه", 9: "سپتامبر"}}, Days: ut.CalendarDayFormatNames{Abbreviated: ut.CalendarDayFormatNameValue{6: "شنبه", 0: "یکشنبه", 1: "دوشنبه", 2: "سه\u200cشنبه", 3: "چهارشنبه", 4: "پنجشنبه", 5: "جمعه"}, Narrow: ut.CalendarDayFormatNameValue{3: "چ", 4: "پ", 5: "ج", 6: "ش", 0: "ی", 1: "د", 2: "س"}, Short: ut.CalendarDayFormatNameValue{6: "ش", 0: "۱ش", 1: "۲ش", 2: "۳ش", 3: "۴ش", 4: "۵ش", 5: "ج"}, Wide: ut.CalendarDayFormatNameValue{3: "چهارشنبه", 4: "پنجشنبه", 5: "جمعه", 6: "شنبه", 0: "یکشنبه", 1: "دوشنبه", 2: "سه\u200cشنبه"}}, Periods: ut.CalendarPeriodFormatNames{Abbreviated: ut.CalendarPeriodFormatNameValue{"evening1": "سر شب", "night1": "شب", "midnight": "نیمه\u200cشب", "am": "ق.ظ.", "noon": "ظ", "pm": "ب.ظ.", "morning1": "صبح", "afternoon1": "عصر"}, Narrow: ut.CalendarPeriodFormatNameValue{"pm": "ب.ظ.", "morning1": "صبح", "afternoon1": "عصر", "evening1": "سر شب", "night1": "شب", "midnight": "نیمه\u200cشب", "am": "ق.ظ.", "noon": "ظهر"}, Short: ut.CalendarPeriodFormatNameValue(nil), Wide: ut.CalendarPeriodFormatNameValue{"evening1": "سر شب", "night1": "شب", "midnight": "نیمه\u200cشب", "am": "قبل\u200cازظهر", "noon": "ظهر", "pm": "بعدازظهر", "morning1": "صبح", "afternoon1": "عصر"}}, Eras: ut.Eras{AD: ut.CalendarEraFormatNames{Full: "میلادی", Abbrev: "م.", Narrow: "م"}, BC: ut.CalendarEraFormatNames{Full: "قبل از میلاد", Abbrev: "ق.م.", Narrow: "ق"}}}}