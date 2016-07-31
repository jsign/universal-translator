package ru_MD

import "github.com/go-playground/universal-translator"

var calendar = ut.Calendar{Formats: ut.CalendarFormats{DateEra: ut.DateEra{BC: ut.CalendarDateFormat{Full: "EEEE, d MMMM y 'г'.", Long: "d MMMM y 'г'.", Medium: "d MMM y 'г'.", Short: "dd.MM.yy"}, AD: ut.CalendarDateFormat{Full: "EEEE, d MMMM y 'г'.", Long: "d MMMM y 'г'.", Medium: "d MMM y 'г'.", Short: "dd.MM.yy"}}, Time: ut.CalendarDateFormat{Full: "H:mm:ss zzzz", Long: "H:mm:ss z", Medium: "H:mm:ss", Short: "H:mm"}, DateTime: ut.CalendarDateFormat{Full: "{1}, {0}", Long: "{1}, {0}", Medium: "{1}, {0}", Short: "{1}, {0}"}}, FormatNames: ut.CalendarFormatNames{Months: ut.CalendarMonthFormatNames{Abbreviated: ut.CalendarMonthFormatNameValue{10: "окт.", 3: "март", 9: "сент.", 8: "авг.", 1: "янв.", 4: "апр.", 6: "июнь", 2: "февр.", 5: "май", 7: "июль", 11: "нояб.", 12: "дек."}, Narrow: ut.CalendarMonthFormatNameValue{6: "И", 1: "Я", 5: "М", 8: "А", 11: "Н", 12: "Д", 9: "С", 7: "И", 10: "О", 2: "Ф", 3: "М", 4: "А"}, Short: ut.CalendarMonthFormatNameValue{}, Wide: ut.CalendarMonthFormatNameValue{6: "июнь", 7: "июль", 9: "сентябрь", 11: "ноябрь", 4: "апрель", 1: "январь", 5: "май", 10: "октябрь", 8: "август", 12: "декабрь", 2: "февраль", 3: "март"}}, Days: ut.CalendarDayFormatNames{Abbreviated: ut.CalendarDayFormatNameValue{6: "сб", 0: "вс", 1: "пн", 2: "вт", 3: "ср", 4: "чт", 5: "пт"}, Narrow: ut.CalendarDayFormatNameValue{4: "Ч", 5: "П", 6: "С", 0: "В", 1: "П", 2: "В", 3: "С"}, Short: ut.CalendarDayFormatNameValue{0: "вс", 1: "пн", 2: "вт", 3: "ср", 4: "чт", 5: "пт", 6: "сб"}, Wide: ut.CalendarDayFormatNameValue{6: "суббота", 0: "воскресенье", 1: "понедельник", 2: "вторник", 3: "среда", 4: "четверг", 5: "пятница"}}, Periods: ut.CalendarPeriodFormatNames{Abbreviated: ut.CalendarPeriodFormatNameValue{"night1": "ночь", "midnight": "полн.", "am": "ДП", "noon": "полд.", "pm": "ПП", "morning1": "утро", "afternoon1": "день", "evening1": "веч."}, Narrow: ut.CalendarPeriodFormatNameValue{"am": "ДП", "noon": "полд.", "pm": "ПП", "morning1": "утро", "afternoon1": "день", "evening1": "веч.", "night1": "ночь", "midnight": "полн."}, Short: ut.CalendarPeriodFormatNameValue{}, Wide: ut.CalendarPeriodFormatNameValue{"afternoon1": "день", "evening1": "вечер", "night1": "ночь", "midnight": "полночь", "am": "ДП", "noon": "полдень", "pm": "ПП", "morning1": "утро"}}, Eras: ut.Eras{AD: ut.CalendarEraFormatNames{Full: "", Abbrev: "", Narrow: ""}, BC: ut.CalendarEraFormatNames{Full: "", Abbrev: "", Narrow: ""}}}}