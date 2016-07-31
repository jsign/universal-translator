package az

import "github.com/go-playground/universal-translator"

var calendar = ut.Calendar{Formats: ut.CalendarFormats{DateEra: ut.DateEra{BC: ut.CalendarDateFormat{Full: "d MMMM y, EEEE", Long: "d MMMM y", Medium: "d MMM y", Short: "dd.MM.yy"}, AD: ut.CalendarDateFormat{Full: "d MMMM y, EEEE", Long: "d MMMM y", Medium: "d MMM y", Short: "dd.MM.yy"}}, Time: ut.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: ut.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"}}, FormatNames: ut.CalendarFormatNames{Months: ut.CalendarMonthFormatNames{Abbreviated: ut.CalendarMonthFormatNameValue{5: "may", 7: "iyl", 9: "sen", 1: "yan", 4: "apr", 6: "iyn", 8: "avq", 10: "okt", 11: "noy", 12: "dek", 2: "fev", 3: "mar"}, Narrow: ut.CalendarMonthFormatNameValue{10: "10", 11: "11", 5: "5", 6: "6", 8: "8", 4: "4", 7: "7", 9: "9", 12: "12", 1: "1", 2: "2", 3: "3"}, Short: ut.CalendarMonthFormatNameValue(nil), Wide: ut.CalendarMonthFormatNameValue{4: "Aprel", 5: "May", 6: "İyun", 8: "Avqust", 9: "Sentyabr", 11: "Noyabr", 1: "Yanvar", 2: "Fevral", 3: "Mart", 7: "İyul", 10: "Oktyabr", 12: "Dekabr"}}, Days: ut.CalendarDayFormatNames{Abbreviated: ut.CalendarDayFormatNameValue{2: "Ç.A.", 3: "Ç.", 4: "C.A.", 5: "C.", 6: "Ş.", 0: "B.", 1: "B.E."}, Narrow: ut.CalendarDayFormatNameValue{4: "4", 5: "5", 6: "6", 0: "7", 1: "1", 2: "2", 3: "3"}, Short: ut.CalendarDayFormatNameValue{0: "B.", 1: "B.E.", 2: "Ç.A.", 3: "Ç.", 4: "C.A.", 5: "C.", 6: "Ş."}, Wide: ut.CalendarDayFormatNameValue{6: "şənbə", 0: "bazar", 1: "bazar ertəsi", 2: "çərşənbə axşamı", 3: "çərşənbə", 4: "cümə axşamı", 5: "cümə"}}, Periods: ut.CalendarPeriodFormatNames{Abbreviated: ut.CalendarPeriodFormatNameValue{"morning2": "səhər", "afternoon1": "gündüz", "evening1": "axşamüstü", "morning1": "sübh", "am": "AM", "noon": "günorta", "pm": "PM", "night1": "axşam", "night2": "gecə", "midnight": "gecəyarı"}, Narrow: ut.CalendarPeriodFormatNameValue{"morning2": "səhər", "afternoon1": "gündüz", "evening1": "axşamüstü", "night1": "axşam", "am": "a", "noon": "g", "morning1": "sübh", "midnight": "gecəyarı", "pm": "p", "night2": "gecə"}, Short: ut.CalendarPeriodFormatNameValue(nil), Wide: ut.CalendarPeriodFormatNameValue{"midnight": "gecəyarı", "am": "AM", "morning2": "səhər", "afternoon1": "gündüz", "evening1": "axşamüstü", "noon": "günorta", "pm": "PM", "morning1": "sübh", "night1": "axşam", "night2": "gecə"}}, Eras: ut.Eras{AD: ut.CalendarEraFormatNames{Full: "eramız", Abbrev: "b.e.", Narrow: ""}, BC: ut.CalendarEraFormatNames{Full: "eramızdan əvvəl", Abbrev: "e.ə.", Narrow: ""}}}}