package sg

import "github.com/go-playground/universal-translator"

var calendar = ut.Calendar{Formats: ut.CalendarFormats{Date: ut.CalendarDateFormat{Full: "EEEE d MMMM y", Long: "d MMMM y", Medium: "d MMM, y", Short: "d/M/y"}, Time: ut.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: ut.CalendarDateFormat{Full: "", Long: "", Medium: "", Short: ""}}, FormatNames: ut.CalendarFormatNames{Months: ut.CalendarMonthFormatNames{Abbreviated: ut.CalendarMonthFormatNameValue{11: "Nab", 12: "Kak", 4: "Ngu", 5: "Bêl", 6: "Fön", 8: "Kük", 9: "Mvu", 10: "Ngb", 1: "Nye", 2: "Ful", 3: "Mbä", 7: "Len"}, Narrow: ut.CalendarMonthFormatNameValue{8: "K", 9: "M", 10: "N", 11: "N", 12: "K", 1: "N", 5: "B", 6: "F", 7: "L", 2: "F", 3: "M", 4: "N"}, Short: ut.CalendarMonthFormatNameValue(nil), Wide: ut.CalendarMonthFormatNameValue{1: "Nyenye", 3: "Mbängü", 5: "Bêläwü", 7: "Lengua", 10: "Ngberere", 11: "Nabändüru", 2: "Fulundïgi", 4: "Ngubùe", 6: "Föndo", 8: "Kükürü", 9: "Mvuka", 12: "Kakauka"}}, Days: ut.CalendarDayFormatNames{Abbreviated: ut.CalendarDayFormatNameValue{5: "Lâp", 6: "Lây", 0: "Bk1", 1: "Bk2", 2: "Bk3", 3: "Bk4", 4: "Bk5"}, Narrow: ut.CalendarDayFormatNameValue{5: "P", 6: "Y", 0: "K", 1: "S", 2: "T", 3: "S", 4: "K"}, Short: ut.CalendarDayFormatNameValue(nil), Wide: ut.CalendarDayFormatNameValue{2: "Bïkua-ptâ", 3: "Bïkua-usïö", 4: "Bïkua-okü", 5: "Lâpôsö", 6: "Lâyenga", 0: "Bikua-ôko", 1: "Bïkua-ûse"}}, Periods: ut.CalendarPeriodFormatNames{Abbreviated: ut.CalendarPeriodFormatNameValue(nil), Narrow: ut.CalendarPeriodFormatNameValue(nil), Short: ut.CalendarPeriodFormatNameValue(nil), Wide: ut.CalendarPeriodFormatNameValue{"am": "ND", "pm": "LK"}}}}