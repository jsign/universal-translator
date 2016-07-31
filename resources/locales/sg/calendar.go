package sg

import "github.com/go-playground/universal-translator"

var calendar = ut.Calendar{Formats: ut.CalendarFormats{DateEra: ut.DateEra{BC: ut.CalendarDateFormat{Full: "EEEE d MMMM y", Long: "d MMMM y", Medium: "d MMM, y", Short: "d/M/y"}, AD: ut.CalendarDateFormat{Full: "EEEE d MMMM y", Long: "d MMMM y", Medium: "d MMM, y", Short: "d/M/y"}}, Time: ut.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: ut.CalendarDateFormat{Full: "", Long: "", Medium: "", Short: ""}}, FormatNames: ut.CalendarFormatNames{Months: ut.CalendarMonthFormatNames{Abbreviated: ut.CalendarMonthFormatNameValue{3: "Mbä", 4: "Ngu", 1: "Nye", 2: "Ful", 5: "Bêl", 6: "Fön", 7: "Len", 8: "Kük", 9: "Mvu", 10: "Ngb", 11: "Nab", 12: "Kak"}, Narrow: ut.CalendarMonthFormatNameValue{3: "M", 4: "N", 6: "F", 7: "L", 8: "K", 9: "M", 11: "N", 1: "N", 2: "F", 5: "B", 10: "N", 12: "K"}, Short: ut.CalendarMonthFormatNameValue(nil), Wide: ut.CalendarMonthFormatNameValue{2: "Fulundïgi", 7: "Lengua", 8: "Kükürü", 12: "Kakauka", 1: "Nyenye", 3: "Mbängü", 4: "Ngubùe", 5: "Bêläwü", 6: "Föndo", 9: "Mvuka", 10: "Ngberere", 11: "Nabändüru"}}, Days: ut.CalendarDayFormatNames{Abbreviated: ut.CalendarDayFormatNameValue{5: "Lâp", 6: "Lây", 0: "Bk1", 1: "Bk2", 2: "Bk3", 3: "Bk4", 4: "Bk5"}, Narrow: ut.CalendarDayFormatNameValue{3: "S", 4: "K", 5: "P", 6: "Y", 0: "K", 1: "S", 2: "T"}, Short: ut.CalendarDayFormatNameValue(nil), Wide: ut.CalendarDayFormatNameValue{1: "Bïkua-ûse", 2: "Bïkua-ptâ", 3: "Bïkua-usïö", 4: "Bïkua-okü", 5: "Lâpôsö", 6: "Lâyenga", 0: "Bikua-ôko"}}, Periods: ut.CalendarPeriodFormatNames{Abbreviated: ut.CalendarPeriodFormatNameValue{"am": "ND", "pm": "LK"}, Narrow: ut.CalendarPeriodFormatNameValue(nil), Short: ut.CalendarPeriodFormatNameValue(nil), Wide: ut.CalendarPeriodFormatNameValue{"am": "ND", "pm": "LK"}}, Eras: ut.Eras{AD: ut.CalendarEraFormatNames{Full: "", Abbrev: "", Narrow: ""}, BC: ut.CalendarEraFormatNames{Full: "Kôzo na Krîstu", Abbrev: "KnK", Narrow: ""}}}}