package sk

import "github.com/go-playground/universal-translator"

var calendar = ut.Calendar{Formats: ut.CalendarFormats{DateEra: ut.DateEra{BC: ut.CalendarDateFormat{Full: "EEEE, d. MMMM y", Long: "d. MMMM y", Medium: "d. M. y", Short: "d.M.yy"}, AD: ut.CalendarDateFormat{Full: "EEEE, d. MMMM y", Long: "d. MMMM y", Medium: "d. M. y", Short: "d.M.yy"}}, Time: ut.CalendarDateFormat{Full: "H:mm:ss zzzz", Long: "H:mm:ss z", Medium: "H:mm:ss", Short: "H:mm"}, DateTime: ut.CalendarDateFormat{Full: "{1}, {0}", Long: "{1}, {0}", Medium: "{1}, {0}", Short: "{1} {0}"}}, FormatNames: ut.CalendarFormatNames{Months: ut.CalendarMonthFormatNames{Abbreviated: ut.CalendarMonthFormatNameValue{1: "jan", 2: "feb", 7: "júl", 10: "okt", 12: "dec", 9: "sep", 11: "nov", 3: "mar", 4: "apr", 5: "máj", 6: "jún", 8: "aug"}, Narrow: ut.CalendarMonthFormatNameValue{9: "s", 11: "n", 3: "m", 4: "a", 5: "m", 6: "j", 7: "j", 8: "a", 10: "o", 12: "d", 1: "j", 2: "f"}, Short: ut.CalendarMonthFormatNameValue(nil), Wide: ut.CalendarMonthFormatNameValue{1: "január", 2: "február", 3: "marec", 4: "apríl", 9: "september", 10: "október", 11: "november", 12: "december", 5: "máj", 6: "jún", 7: "júl", 8: "august"}}, Days: ut.CalendarDayFormatNames{Abbreviated: ut.CalendarDayFormatNameValue{0: "ne", 1: "po", 2: "ut", 3: "st", 4: "št", 5: "pi", 6: "so"}, Narrow: ut.CalendarDayFormatNameValue{0: "n", 1: "p", 2: "u", 3: "s", 4: "š", 5: "p", 6: "s"}, Short: ut.CalendarDayFormatNameValue{0: "ne", 1: "po", 2: "ut", 3: "st", 4: "št", 5: "pi", 6: "so"}, Wide: ut.CalendarDayFormatNameValue{6: "sobota", 0: "nedeľa", 1: "pondelok", 2: "utorok", 3: "streda", 4: "štvrtok", 5: "piatok"}}, Periods: ut.CalendarPeriodFormatNames{Abbreviated: ut.CalendarPeriodFormatNameValue{"morning1": "ráno", "morning2": "dopol.", "afternoon1": "popol.", "midnight": "poln.", "am": "AM", "noon": "pol.", "pm": "PM", "evening1": "večer", "night1": "noc"}, Narrow: ut.CalendarPeriodFormatNameValue{"morning1": "ráno", "morning2": "dop.", "afternoon1": "pop.", "evening1": "več.", "midnight": "poln.", "noon": "pol.", "pm": "PM", "am": "AM", "night1": "noc"}, Short: ut.CalendarPeriodFormatNameValue(nil), Wide: ut.CalendarPeriodFormatNameValue{"midnight": "polnoc", "morning2": "dopoludnie", "afternoon1": "popoludnie", "night1": "noc", "am": "AM", "noon": "poludnie", "pm": "PM", "morning1": "ráno", "evening1": "večer"}}, Eras: ut.Eras{AD: ut.CalendarEraFormatNames{Full: "po Kristovi", Abbrev: "po Kr.", Narrow: ""}, BC: ut.CalendarEraFormatNames{Full: "pred Kristom", Abbrev: "pred Kr.", Narrow: ""}}}}
