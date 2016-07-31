package ar_SS

import "github.com/go-playground/universal-translator"

var calendar = ut.Calendar{Formats: ut.CalendarFormats{DateEra: ut.DateEra{BC: ut.CalendarDateFormat{Full: "EEEE، d MMMM، y", Long: "d MMMM، y", Medium: "dd\u200f/MM\u200f/y", Short: "d\u200f/M\u200f/y"}, AD: ut.CalendarDateFormat{Full: "EEEE، d MMMM، y", Long: "d MMMM، y", Medium: "dd\u200f/MM\u200f/y", Short: "d\u200f/M\u200f/y"}}, Time: ut.CalendarDateFormat{Full: "h:mm:ss a zzzz", Long: "h:mm:ss a z", Medium: "h:mm:ss a", Short: "h:mm a"}, DateTime: ut.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"}}, FormatNames: ut.CalendarFormatNames{Months: ut.CalendarMonthFormatNames{Abbreviated: ut.CalendarMonthFormatNameValue{11: "نوفمبر", 4: "أبريل", 6: "يونيو", 12: "ديسمبر", 9: "سبتمبر", 3: "مارس", 5: "مايو", 7: "يوليو", 8: "أغسطس", 1: "يناير", 2: "فبراير", 10: "أكتوبر"}, Narrow: ut.CalendarMonthFormatNameValue{4: "أ", 12: "د", 1: "ي", 8: "غ", 6: "ن", 11: "ب", 2: "ف", 9: "س", 5: "و", 7: "ل", 10: "ك", 3: "م"}, Short: ut.CalendarMonthFormatNameValue{}, Wide: ut.CalendarMonthFormatNameValue{11: "نوفمبر", 2: "فبراير", 5: "مايو", 7: "يوليو", 10: "أكتوبر", 1: "يناير", 12: "ديسمبر", 8: "أغسطس", 3: "مارس", 4: "أبريل", 6: "يونيو", 9: "سبتمبر"}}, Days: ut.CalendarDayFormatNames{Abbreviated: ut.CalendarDayFormatNameValue{5: "الجمعة", 6: "السبت", 0: "الأحد", 1: "الاثنين", 2: "الثلاثاء", 3: "الأربعاء", 4: "الخميس"}, Narrow: ut.CalendarDayFormatNameValue{6: "س", 0: "ح", 1: "ن", 2: "ث", 3: "ر", 4: "خ", 5: "ج"}, Short: ut.CalendarDayFormatNameValue{2: "الثلاثاء", 3: "الأربعاء", 4: "الخميس", 5: "الجمعة", 6: "السبت", 0: "الأحد", 1: "الاثنين"}, Wide: ut.CalendarDayFormatNameValue{3: "الأربعاء", 4: "الخميس", 5: "الجمعة", 6: "السبت", 0: "الأحد", 1: "الاثنين", 2: "الثلاثاء"}}, Periods: ut.CalendarPeriodFormatNames{Abbreviated: ut.CalendarPeriodFormatNameValue{"afternoon2": "بعد الظهر", "night2": "ليلاً", "pm": "م", "night1": "منتصف الليل", "am": "ص", "morning2": "صباحًا", "afternoon1": "ظهرًا", "evening1": "مساءً", "morning1": "فجرا"}, Narrow: ut.CalendarPeriodFormatNameValue{"am": "ص", "morning2": "صباحًا", "afternoon2": "بعد الظهر", "afternoon1": "ظهرًا", "night1": "منتصف الليل", "evening1": "مساءً", "night2": "ليلاً", "pm": "م", "morning1": "فجرا"}, Short: ut.CalendarPeriodFormatNameValue{}, Wide: ut.CalendarPeriodFormatNameValue{"afternoon1": "ظهرًا", "morning1": "فجرا", "pm": "مساءً", "night1": "منتصف الليل", "morning2": "صباحًا", "evening1": "مساءً", "afternoon2": "بعد الظهر", "night2": "ليلاً", "am": "صباحًا"}}, Eras: ut.Eras{AD: ut.CalendarEraFormatNames{Full: "", Abbrev: "", Narrow: ""}, BC: ut.CalendarEraFormatNames{Full: "", Abbrev: "", Narrow: ""}}}}