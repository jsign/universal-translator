package fo_DK

import "github.com/go-playground/universal-translator"

var calendar = ut.Calendar{Formats: ut.CalendarFormats{DateEra: ut.DateEra{BC: ut.CalendarDateFormat{Full: "EEEE, d. MMMM y", Long: "d. MMMM y", Medium: "dd.MM.y", Short: "dd.MM.yy"}, AD: ut.CalendarDateFormat{Full: "EEEE, d. MMMM y", Long: "d. MMMM y", Medium: "dd.MM.y", Short: "dd.MM.yy"}}, Time: ut.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: ut.CalendarDateFormat{Full: "{1} 'kl'. {0}", Long: "{1} 'kl'. {0}", Medium: "{1}, {0}", Short: "{1}, {0}"}}, FormatNames: ut.CalendarFormatNames{Months: ut.CalendarMonthFormatNames{Abbreviated: ut.CalendarMonthFormatNameValue{7: "jul", 8: "aug", 10: "okt", 11: "nov", 12: "des", 5: "mai", 9: "sep", 4: "apr", 2: "feb", 3: "mar", 6: "jun", 1: "jan"}, Narrow: ut.CalendarMonthFormatNameValue{5: "M", 7: "J", 4: "A", 12: "D", 1: "J", 8: "A", 11: "N", 3: "M", 6: "J", 10: "O", 2: "F", 9: "S"}, Short: ut.CalendarMonthFormatNameValue{}, Wide: ut.CalendarMonthFormatNameValue{7: "juli", 12: "desember", 1: "januar", 2: "februar", 11: "november", 3: "mars", 5: "mai", 6: "juni", 4: "apríl", 8: "august", 10: "oktober", 9: "september"}}, Days: ut.CalendarDayFormatNames{Abbreviated: ut.CalendarDayFormatNameValue{2: "týs", 3: "mik", 4: "hós", 5: "frí", 6: "ley", 0: "sun", 1: "mán"}, Narrow: ut.CalendarDayFormatNameValue{3: "M", 4: "H", 5: "F", 6: "L", 0: "S", 1: "M", 2: "T"}, Short: ut.CalendarDayFormatNameValue{1: "má", 2: "tý", 3: "mi", 4: "hó", 5: "fr", 6: "le", 0: "su"}, Wide: ut.CalendarDayFormatNameValue{4: "hósdagur", 5: "fríggjadagur", 6: "leygardagur", 0: "sunnudagur", 1: "mánadagur", 2: "týsdagur", 3: "mikudagur"}}, Periods: ut.CalendarPeriodFormatNames{Abbreviated: ut.CalendarPeriodFormatNameValue{"am": "AM", "pm": "PM"}, Narrow: ut.CalendarPeriodFormatNameValue{"am": "AM", "pm": "PM"}, Short: ut.CalendarPeriodFormatNameValue{}, Wide: ut.CalendarPeriodFormatNameValue{"am": "AM", "pm": "PM"}}, Eras: ut.Eras{AD: ut.CalendarEraFormatNames{Full: "", Abbrev: "", Narrow: ""}, BC: ut.CalendarEraFormatNames{Full: "", Abbrev: "", Narrow: ""}}}}