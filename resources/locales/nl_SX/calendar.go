package nl_SX

import "github.com/go-playground/universal-translator"

var calendar = ut.Calendar{Formats: ut.CalendarFormats{DateEra: ut.DateEra{BC: ut.CalendarDateFormat{Full: "EEEE d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "dd-MM-yy"}, AD: ut.CalendarDateFormat{Full: "EEEE d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "dd-MM-yy"}}, Time: ut.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: ut.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"}}, FormatNames: ut.CalendarFormatNames{Months: ut.CalendarMonthFormatNames{Abbreviated: ut.CalendarMonthFormatNameValue{1: "jan.", 12: "dec.", 3: "mrt.", 4: "apr.", 5: "mei", 2: "feb.", 6: "jun.", 7: "jul.", 11: "nov.", 8: "aug.", 9: "sep.", 10: "okt."}, Narrow: ut.CalendarMonthFormatNameValue{8: "A", 10: "O", 2: "F", 1: "J", 11: "N", 7: "J", 9: "S", 12: "D", 4: "A", 5: "M", 6: "J", 3: "M"}, Short: ut.CalendarMonthFormatNameValue{}, Wide: ut.CalendarMonthFormatNameValue{3: "maart", 6: "juni", 2: "februari", 10: "oktober", 1: "januari", 5: "mei", 7: "juli", 8: "augustus", 11: "november", 4: "april", 9: "september", 12: "december"}}, Days: ut.CalendarDayFormatNames{Abbreviated: ut.CalendarDayFormatNameValue{1: "ma", 2: "di", 3: "wo", 4: "do", 5: "vr", 6: "za", 0: "zo"}, Narrow: ut.CalendarDayFormatNameValue{6: "Z", 0: "Z", 1: "M", 2: "D", 3: "W", 4: "D", 5: "V"}, Short: ut.CalendarDayFormatNameValue{0: "zo", 1: "ma", 2: "di", 3: "wo", 4: "do", 5: "vr", 6: "za"}, Wide: ut.CalendarDayFormatNameValue{5: "vrijdag", 6: "zaterdag", 0: "zondag", 1: "maandag", 2: "dinsdag", 3: "woensdag", 4: "donderdag"}}, Periods: ut.CalendarPeriodFormatNames{Abbreviated: ut.CalendarPeriodFormatNameValue{"morning1": "ochtend", "afternoon1": "middag", "evening1": "avond", "night1": "nacht", "midnight": "middernacht", "am": "a.m.", "pm": "p.m."}, Narrow: ut.CalendarPeriodFormatNameValue{"morning1": "ochtend", "afternoon1": "middag", "evening1": "avond", "night1": "nacht", "midnight": "middernacht", "am": "a.m.", "pm": "p.m."}, Short: ut.CalendarPeriodFormatNameValue{}, Wide: ut.CalendarPeriodFormatNameValue{"evening1": "avond", "night1": "nacht", "midnight": "middernacht", "am": "a.m.", "pm": "p.m.", "morning1": "ochtend", "afternoon1": "middag"}}, Eras: ut.Eras{AD: ut.CalendarEraFormatNames{Full: "", Abbrev: "", Narrow: ""}, BC: ut.CalendarEraFormatNames{Full: "", Abbrev: "", Narrow: ""}}}}