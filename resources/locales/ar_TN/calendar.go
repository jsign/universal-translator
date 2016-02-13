package ar_TN

import "github.com/go-playground/universal-translator"

var calendar = ut.Calendar{Formats: ut.CalendarFormats{Date: ut.CalendarDateFormat{Full: "EEEE، d MMMM، y", Long: "d MMMM، y", Medium: "dd\u200f/MM\u200f/y", Short: "d\u200f/M\u200f/y"}, Time: ut.CalendarDateFormat{Full: "h:mm:ss a zzzz", Long: "h:mm:ss a z", Medium: "h:mm:ss a", Short: "h:mm a"}, DateTime: ut.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"}}, FormatNames: ut.CalendarFormatNames{Months: ut.CalendarMonthFormatNames{Abbreviated: ut.CalendarMonthFormatNameValue{12: "ديسمبر", 1: "جانفي", 2: "فيفري", 3: "مارس", 9: "سبتمبر", 10: "أكتوبر", 11: "نوفمبر", 4: "أفريل", 5: "ماي", 6: "جوان", 7: "جويلية", 8: "أوت"}, Narrow: ut.CalendarMonthFormatNameValue{5: "م", 7: "ج", 9: "س", 10: "أ", 11: "ن", 2: "ف", 3: "م", 4: "أ", 12: "د", 1: "ج", 6: "ج", 8: "أ"}, Short: ut.CalendarMonthFormatNameValue{}, Wide: ut.CalendarMonthFormatNameValue{12: "ديسمبر", 3: "مارس", 5: "ماي", 6: "جوان", 7: "جويلية", 8: "أوت", 9: "سبتمبر", 11: "نوفمبر", 1: "جانفي", 2: "فيفري", 4: "أفريل", 10: "أكتوبر"}}, Days: ut.CalendarDayFormatNames{Abbreviated: ut.CalendarDayFormatNameValue{3: "الأربعاء", 4: "الخميس", 5: "الجمعة", 6: "السبت", 0: "الأحد", 1: "الاثنين", 2: "الثلاثاء"}, Narrow: ut.CalendarDayFormatNameValue{5: "ج", 6: "س", 0: "ح", 1: "ن", 2: "ث", 3: "ر", 4: "خ"}, Short: ut.CalendarDayFormatNameValue{0: "الأحد", 1: "الاثنين", 2: "الثلاثاء", 3: "الأربعاء", 4: "الخميس", 5: "الجمعة", 6: "السبت"}, Wide: ut.CalendarDayFormatNameValue{6: "السبت", 0: "الأحد", 1: "الاثنين", 2: "الثلاثاء", 3: "الأربعاء", 4: "الخميس", 5: "الجمعة"}}, Periods: ut.CalendarPeriodFormatNames{Abbreviated: ut.CalendarPeriodFormatNameValue{"pm": "م", "night1": "منتصف الليل", "afternoon2": "بعد الظهر", "am": "ص", "afternoon1": "ظهرًا", "night2": "ليلاً", "morning1": "فجرا", "morning2": "صباحًا", "evening1": "مساءً"}, Narrow: ut.CalendarPeriodFormatNameValue{"morning1": "فجرا", "morning2": "صباحًا", "afternoon1": "ظهرًا", "night2": "ليلاً", "afternoon2": "بعد الظهر", "am": "ص", "pm": "م", "evening1": "مساءً", "night1": "منتصف الليل"}, Short: ut.CalendarPeriodFormatNameValue{}, Wide: ut.CalendarPeriodFormatNameValue{"afternoon1": "ظهرًا", "evening1": "مساءً", "night2": "ليلاً", "am": "صباحًا", "night1": "منتصف الليل", "pm": "مساءً", "morning2": "صباحًا", "afternoon2": "بعد الظهر", "morning1": "فجرا"}}}}