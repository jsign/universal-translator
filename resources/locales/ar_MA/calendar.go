package ar_MA

import "github.com/go-playground/universal-translator"

var calendar = ut.Calendar{Formats: ut.CalendarFormats{DateEra: ut.DateEra{BC: ut.CalendarDateFormat{Full: "EEEE، d MMMM، y", Long: "d MMMM، y", Medium: "dd\u200f/MM\u200f/y", Short: "d\u200f/M\u200f/y"}, AD: ut.CalendarDateFormat{Full: "EEEE، d MMMM، y", Long: "d MMMM، y", Medium: "dd\u200f/MM\u200f/y", Short: "d\u200f/M\u200f/y"}}, Time: ut.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: ut.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"}}, FormatNames: ut.CalendarFormatNames{Months: ut.CalendarMonthFormatNames{Abbreviated: ut.CalendarMonthFormatNameValue{2: "فبراير", 4: "أبريل", 6: "يونيو", 8: "غشت", 10: "أكتوبر", 12: "دجنبر", 1: "يناير", 3: "مارس", 5: "ماي", 7: "يوليوز", 9: "شتنبر", 11: "نونبر"}, Narrow: ut.CalendarMonthFormatNameValue{4: "أ", 3: "م", 2: "ف", 5: "م", 6: "ن", 7: "ل", 8: "غ", 9: "ش", 10: "ك", 1: "ي", 12: "د", 11: "ب"}, Short: ut.CalendarMonthFormatNameValue{}, Wide: ut.CalendarMonthFormatNameValue{9: "شتنبر", 10: "أكتوبر", 12: "دجنبر", 1: "يناير", 5: "ماي", 6: "يونيو", 7: "يوليوز", 11: "نونبر", 2: "فبراير", 3: "مارس", 4: "أبريل", 8: "غشت"}}, Days: ut.CalendarDayFormatNames{Abbreviated: ut.CalendarDayFormatNameValue{2: "الثلاثاء", 3: "الأربعاء", 4: "الخميس", 5: "الجمعة", 6: "السبت", 0: "الأحد", 1: "الاثنين"}, Narrow: ut.CalendarDayFormatNameValue{4: "خ", 5: "ج", 6: "س", 0: "ح", 1: "ن", 2: "ث", 3: "ر"}, Short: ut.CalendarDayFormatNameValue{0: "الأحد", 1: "الاثنين", 2: "الثلاثاء", 3: "الأربعاء", 4: "الخميس", 5: "الجمعة", 6: "السبت"}, Wide: ut.CalendarDayFormatNameValue{1: "الاثنين", 2: "الثلاثاء", 3: "الأربعاء", 4: "الخميس", 5: "الجمعة", 6: "السبت", 0: "الأحد"}}, Periods: ut.CalendarPeriodFormatNames{Abbreviated: ut.CalendarPeriodFormatNameValue{"morning2": "صباحًا", "afternoon1": "ظهرًا", "afternoon2": "بعد الظهر", "am": "ص", "morning1": "فجرا", "evening1": "مساءً", "pm": "م", "night1": "منتصف الليل", "night2": "ليلاً"}, Narrow: ut.CalendarPeriodFormatNameValue{"night1": "منتصف الليل", "evening1": "مساءً", "morning2": "صباحًا", "morning1": "فجرا", "afternoon1": "ظهرًا", "am": "ص", "pm": "م", "afternoon2": "بعد الظهر", "night2": "ليلاً"}, Short: ut.CalendarPeriodFormatNameValue{}, Wide: ut.CalendarPeriodFormatNameValue{"am": "صباحًا", "afternoon2": "بعد الظهر", "morning2": "صباحًا", "afternoon1": "ظهرًا", "evening1": "مساءً", "pm": "مساءً", "night1": "منتصف الليل", "night2": "ليلاً", "morning1": "فجرا"}}, Eras: ut.Eras{AD: ut.CalendarEraFormatNames{Full: "", Abbrev: "", Narrow: ""}, BC: ut.CalendarEraFormatNames{Full: "", Abbrev: "", Narrow: ""}}}}