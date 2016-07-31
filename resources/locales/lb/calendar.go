package lb

import "github.com/go-playground/universal-translator"

var calendar = ut.Calendar{Formats: ut.CalendarFormats{DateEra: ut.DateEra{BC: ut.CalendarDateFormat{Full: "EEEE, d. MMMM y", Long: "d. MMMM y", Medium: "d. MMM y", Short: "dd.MM.yy"}, AD: ut.CalendarDateFormat{Full: "EEEE, d. MMMM y", Long: "d. MMMM y", Medium: "d. MMM y", Short: "dd.MM.yy"}}, Time: ut.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: ut.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"}}, FormatNames: ut.CalendarFormatNames{Months: ut.CalendarMonthFormatNames{Abbreviated: ut.CalendarMonthFormatNameValue{10: "Okt", 11: "Nov", 4: "Abr", 6: "Jun", 7: "Jul", 8: "Aug", 9: "Sep", 12: "Dez", 1: "Jan", 2: "Feb", 3: "Mäe", 5: "Mee"}, Narrow: ut.CalendarMonthFormatNameValue{1: "J", 5: "M", 7: "J", 9: "S", 10: "O", 11: "N", 2: "F", 3: "M", 4: "A", 6: "J", 8: "A", 12: "D"}, Short: ut.CalendarMonthFormatNameValue(nil), Wide: ut.CalendarMonthFormatNameValue{1: "Januar", 2: "Februar", 3: "Mäerz", 9: "September", 10: "Oktober", 11: "November", 4: "Abrëll", 5: "Mee", 6: "Juni", 7: "Juli", 8: "August", 12: "Dezember"}}, Days: ut.CalendarDayFormatNames{Abbreviated: ut.CalendarDayFormatNameValue{4: "Don", 5: "Fre", 6: "Sam", 0: "Son", 1: "Méi", 2: "Dën", 3: "Mët"}, Narrow: ut.CalendarDayFormatNameValue{0: "S", 1: "M", 2: "D", 3: "M", 4: "D", 5: "F", 6: "S"}, Short: ut.CalendarDayFormatNameValue{6: "Sa.", 0: "So.", 1: "Mé.", 2: "Dë.", 3: "Më.", 4: "Do.", 5: "Fr."}, Wide: ut.CalendarDayFormatNameValue{4: "Donneschdeg", 5: "Freideg", 6: "Samschdeg", 0: "Sonndeg", 1: "Méindeg", 2: "Dënschdeg", 3: "Mëttwoch"}}, Periods: ut.CalendarPeriodFormatNames{Abbreviated: ut.CalendarPeriodFormatNameValue{"am": "moies", "pm": "nomëttes"}, Narrow: ut.CalendarPeriodFormatNameValue{"am": "mo.", "pm": "nomë."}, Short: ut.CalendarPeriodFormatNameValue(nil), Wide: ut.CalendarPeriodFormatNameValue{"am": "moies", "pm": "nomëttes"}}, Eras: ut.Eras{AD: ut.CalendarEraFormatNames{Full: "", Abbrev: "n. Chr.", Narrow: ""}, BC: ut.CalendarEraFormatNames{Full: "v. Chr.", Abbrev: "v. Chr.", Narrow: ""}}}}
