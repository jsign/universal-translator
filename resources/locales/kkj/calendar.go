package kkj

import "github.com/go-playground/universal-translator"

var calendar = ut.Calendar{Formats: ut.CalendarFormats{DateEra: ut.DateEra{BC: ut.CalendarDateFormat{Full: "EEEE dd MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "dd/MM y"}, AD: ut.CalendarDateFormat{Full: "EEEE dd MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "dd/MM y"}}, Time: ut.CalendarDateFormat{Full: "", Long: "", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: ut.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"}}, FormatNames: ut.CalendarFormatNames{Months: ut.CalendarMonthFormatNames{Abbreviated: ut.CalendarMonthFormatNameValue(nil), Narrow: ut.CalendarMonthFormatNameValue(nil), Short: ut.CalendarMonthFormatNameValue(nil), Wide: ut.CalendarMonthFormatNameValue{9: "njapi", 10: "nyukul", 11: "11", 12: "ɓulɓusɛ", 1: "pamba", 6: "Nyaŋgwɛ ŋgbanja", 8: "fɛ", 5: "Mɔnɔ ŋgbanja", 7: "kuŋgwɛ", 2: "wanja", 3: "mbiyɔ mɛndoŋgɔ", 4: "Nyɔlɔmbɔŋgɔ"}}, Days: ut.CalendarDayFormatNames{Abbreviated: ut.CalendarDayFormatNameValue{6: "mɔnɔ sɔndi", 0: "sɔndi", 1: "lundi", 2: "mardi", 3: "mɛrkɛrɛdi", 4: "yedi", 5: "vaŋdɛrɛdi"}, Narrow: ut.CalendarDayFormatNameValue{3: "mɛ", 4: "ye", 5: "va", 6: "ms", 0: "so", 1: "lu", 2: "ma"}, Short: ut.CalendarDayFormatNameValue{6: "ms", 0: "so", 1: "lu", 2: "ma", 3: "mɛ", 4: "ye", 5: "va"}, Wide: ut.CalendarDayFormatNameValue{4: "yedi", 5: "vaŋdɛrɛdi", 6: "mɔnɔ sɔndi", 0: "sɔndi", 1: "lundi", 2: "mardi", 3: "mɛrkɛrɛdi"}}, Periods: ut.CalendarPeriodFormatNames{Abbreviated: ut.CalendarPeriodFormatNameValue(nil), Narrow: ut.CalendarPeriodFormatNameValue(nil), Short: ut.CalendarPeriodFormatNameValue(nil), Wide: ut.CalendarPeriodFormatNameValue(nil)}, Eras: ut.Eras{AD: ut.CalendarEraFormatNames{Full: "", Abbrev: "", Narrow: ""}, BC: ut.CalendarEraFormatNames{Full: "", Abbrev: "", Narrow: ""}}}}
