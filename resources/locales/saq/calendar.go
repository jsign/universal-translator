package saq

import "github.com/go-playground/universal-translator"

var calendar = ut.Calendar{Formats: ut.CalendarFormats{DateEra: ut.DateEra{BC: ut.CalendarDateFormat{Full: "EEEE, d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "dd/MM/y"}, AD: ut.CalendarDateFormat{Full: "EEEE, d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "dd/MM/y"}}, Time: ut.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: ut.CalendarDateFormat{Full: "", Long: "", Medium: "", Short: ""}}, FormatNames: ut.CalendarFormatNames{Months: ut.CalendarMonthFormatNames{Abbreviated: ut.CalendarMonthFormatNameValue{12: "Tow", 1: "Obo", 2: "Waa", 7: "Sap", 8: "Isi", 10: "Tom", 11: "Tob", 3: "Oku", 4: "Ong", 5: "Ime", 6: "Ile", 9: "Saa"}, Narrow: ut.CalendarMonthFormatNameValue{6: "I", 9: "S", 10: "T", 11: "T", 1: "O", 3: "O", 5: "I", 7: "S", 8: "I", 12: "T", 2: "W", 4: "O"}, Short: ut.CalendarMonthFormatNameValue(nil), Wide: ut.CalendarMonthFormatNameValue{5: "Lapa le imet", 8: "Lapa le isiet", 10: "Lapa le tomon", 11: "Lapa le tomon obo", 1: "Lapa le obo", 2: "Lapa le waare", 4: "Lapa le ong’wan", 9: "Lapa le saal", 12: "Lapa le tomon waare", 3: "Lapa le okuni", 6: "Lapa le ile", 7: "Lapa le sapa"}}, Days: ut.CalendarDayFormatNames{Abbreviated: ut.CalendarDayFormatNameValue{6: "Kwe", 0: "Are", 1: "Kun", 2: "Ong", 3: "Ine", 4: "Ile", 5: "Sap"}, Narrow: ut.CalendarDayFormatNameValue{3: "I", 4: "I", 5: "S", 6: "K", 0: "A", 1: "K", 2: "O"}, Short: ut.CalendarDayFormatNameValue(nil), Wide: ut.CalendarDayFormatNameValue{6: "Mderot ee kwe", 0: "Mderot ee are", 1: "Mderot ee kuni", 2: "Mderot ee ong’wan", 3: "Mderot ee inet", 4: "Mderot ee ile", 5: "Mderot ee sapa"}}, Periods: ut.CalendarPeriodFormatNames{Abbreviated: ut.CalendarPeriodFormatNameValue{"am": "Tesiran", "pm": "Teipa"}, Narrow: ut.CalendarPeriodFormatNameValue(nil), Short: ut.CalendarPeriodFormatNameValue(nil), Wide: ut.CalendarPeriodFormatNameValue{"am": "Tesiran", "pm": "Teipa"}}, Eras: ut.Eras{AD: ut.CalendarEraFormatNames{Full: "", Abbrev: "", Narrow: ""}, BC: ut.CalendarEraFormatNames{Full: "Kabla ya Christo", Abbrev: "KK", Narrow: ""}}}}
