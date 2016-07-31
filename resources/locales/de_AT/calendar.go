package de_AT

import "github.com/go-playground/universal-translator"

var calendar = ut.Calendar{Formats: ut.CalendarFormats{DateEra: ut.DateEra{BC: ut.CalendarDateFormat{Full: "EEEE, d. MMMM y", Long: "d. MMMM y", Medium: "dd.MM.y", Short: "dd.MM.yy"}, AD: ut.CalendarDateFormat{Full: "EEEE, d. MMMM y", Long: "d. MMMM y", Medium: "dd.MM.y", Short: "dd.MM.yy"}}, Time: ut.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: ut.CalendarDateFormat{Full: "{1} 'um' {0}", Long: "{1} 'um' {0}", Medium: "{1}, {0}", Short: "{1}, {0}"}}, FormatNames: ut.CalendarFormatNames{Months: ut.CalendarMonthFormatNames{Abbreviated: ut.CalendarMonthFormatNameValue{4: "Apr", 6: "Jun", 9: "Sep", 10: "Okt", 12: "Dez", 1: "Jän", 2: "Feb", 3: "Mär", 5: "Mai", 7: "Jul", 8: "Aug", 11: "Nov"}, Narrow: ut.CalendarMonthFormatNameValue{7: "J", 11: "N", 1: "J", 8: "A", 10: "O", 3: "M", 12: "D", 2: "F", 5: "M", 6: "J", 9: "S", 4: "A"}, Short: ut.CalendarMonthFormatNameValue{}, Wide: ut.CalendarMonthFormatNameValue{3: "März", 5: "Mai", 6: "Juni", 7: "Juli", 8: "August", 12: "Dezember", 2: "Februar", 4: "April", 9: "September", 10: "Oktober", 11: "November", 1: "Jänner"}}, Days: ut.CalendarDayFormatNames{Abbreviated: ut.CalendarDayFormatNameValue{2: "Di", 3: "Mi", 4: "Do", 5: "Fr", 6: "Sa", 0: "So", 1: "Mo"}, Narrow: ut.CalendarDayFormatNameValue{2: "D", 3: "M", 4: "D", 5: "F", 6: "S", 0: "S", 1: "M"}, Short: ut.CalendarDayFormatNameValue{4: "Do.", 5: "Fr.", 6: "Sa.", 0: "So.", 1: "Mo.", 2: "Di.", 3: "Mi."}, Wide: ut.CalendarDayFormatNameValue{1: "Montag", 2: "Dienstag", 3: "Mittwoch", 4: "Donnerstag", 5: "Freitag", 6: "Samstag", 0: "Sonntag"}}, Periods: ut.CalendarPeriodFormatNames{Abbreviated: ut.CalendarPeriodFormatNameValue{"morning2": "Vormittag", "evening1": "Abend", "midnight": "Mitternacht", "afternoon2": "Nachmittag", "am": "vorm.", "pm": "nachm.", "night1": "Nacht", "morning1": "Morgen", "afternoon1": "Mittag"}, Narrow: ut.CalendarPeriodFormatNameValue{"morning1": "Morgen", "afternoon2": "Nachmittag", "evening1": "Abend", "morning2": "Vormittag", "night1": "Nacht", "am": "vorm.", "pm": "nachm.", "afternoon1": "Mittag", "midnight": "Mitternacht"}, Short: ut.CalendarPeriodFormatNameValue{}, Wide: ut.CalendarPeriodFormatNameValue{"evening1": "Abend", "night1": "Nacht", "midnight": "Mitternacht", "pm": "nachm.", "afternoon1": "Mittag", "am": "vorm.", "morning2": "Vormittag", "afternoon2": "Nachmittag", "morning1": "Morgen"}}, Eras: ut.Eras{AD: ut.CalendarEraFormatNames{Full: "", Abbrev: "", Narrow: ""}, BC: ut.CalendarEraFormatNames{Full: "", Abbrev: "", Narrow: ""}}}}