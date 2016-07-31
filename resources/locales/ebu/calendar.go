package ebu

import "github.com/go-playground/universal-translator"

var calendar = ut.Calendar{Formats: ut.CalendarFormats{DateEra: ut.DateEra{BC: ut.CalendarDateFormat{Full: "EEEE, d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "dd/MM/y"}, AD: ut.CalendarDateFormat{Full: "EEEE, d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "dd/MM/y"}}, Time: ut.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: ut.CalendarDateFormat{Full: "", Long: "", Medium: "", Short: ""}}, FormatNames: ut.CalendarFormatNames{Months: ut.CalendarMonthFormatNames{Abbreviated: ut.CalendarMonthFormatNameValue{9: "Ken", 10: "Iku", 11: "Imw", 1: "Mbe", 3: "Kat", 4: "Kan", 6: "Gan", 12: "Igi", 2: "Kai", 5: "Gat", 7: "Mug", 8: "Knn"}, Narrow: ut.CalendarMonthFormatNameValue{5: "G", 6: "G", 7: "M", 8: "K", 10: "I", 11: "I", 1: "M", 2: "K", 3: "K", 4: "K", 9: "K", 12: "I"}, Short: ut.CalendarMonthFormatNameValue(nil), Wide: ut.CalendarMonthFormatNameValue{1: "Mweri wa mbere", 2: "Mweri wa kaĩri", 3: "Mweri wa kathatũ", 5: "Mweri wa gatano", 6: "Mweri wa gatantatũ", 10: "Mweri wa ikũmi", 4: "Mweri wa kana", 7: "Mweri wa mũgwanja", 8: "Mweri wa kanana", 9: "Mweri wa kenda", 11: "Mweri wa ikũmi na ũmwe", 12: "Mweri wa ikũmi na Kaĩrĩ"}}, Days: ut.CalendarDayFormatNames{Abbreviated: ut.CalendarDayFormatNameValue{2: "Ine", 3: "Tan", 4: "Arm", 5: "Maa", 6: "NMM", 0: "Kma", 1: "Tat"}, Narrow: ut.CalendarDayFormatNameValue{3: "N", 4: "A", 5: "M", 6: "N", 0: "K", 1: "N", 2: "N"}, Short: ut.CalendarDayFormatNameValue(nil), Wide: ut.CalendarDayFormatNameValue{1: "Njumatatu", 2: "Njumaine", 3: "Njumatano", 4: "Aramithi", 5: "Njumaa", 6: "NJumamothii", 0: "Kiumia"}}, Periods: ut.CalendarPeriodFormatNames{Abbreviated: ut.CalendarPeriodFormatNameValue{"am": "KI", "pm": "UT"}, Narrow: ut.CalendarPeriodFormatNameValue(nil), Short: ut.CalendarPeriodFormatNameValue(nil), Wide: ut.CalendarPeriodFormatNameValue{"am": "KI", "pm": "UT"}}, Eras: ut.Eras{AD: ut.CalendarEraFormatNames{Full: "", Abbrev: "", Narrow: ""}, BC: ut.CalendarEraFormatNames{Full: "Mbere ya Kristo", Abbrev: "MK", Narrow: ""}}}}
