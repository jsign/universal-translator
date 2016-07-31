package seh

import "github.com/go-playground/universal-translator"

var calendar = ut.Calendar{Formats: ut.CalendarFormats{DateEra: ut.DateEra{BC: ut.CalendarDateFormat{Full: "EEEE, d 'de' MMMM 'de' y", Long: "d 'de' MMMM 'de' y", Medium: "d 'de' MMM 'de' y", Short: "d/M/y"}, AD: ut.CalendarDateFormat{Full: "EEEE, d 'de' MMMM 'de' y", Long: "d 'de' MMMM 'de' y", Medium: "d 'de' MMM 'de' y", Short: "d/M/y"}}, Time: ut.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: ut.CalendarDateFormat{Full: "", Long: "", Medium: "", Short: ""}}, FormatNames: ut.CalendarFormatNames{Months: ut.CalendarMonthFormatNames{Abbreviated: ut.CalendarMonthFormatNameValue{4: "Abr", 5: "Mai", 7: "Jul", 8: "Aug", 2: "Fev", 3: "Mar", 6: "Jun", 9: "Set", 10: "Otu", 11: "Nov", 12: "Dec", 1: "Jan"}, Narrow: ut.CalendarMonthFormatNameValue{9: "S", 1: "J", 2: "F", 4: "A", 7: "J", 10: "O", 11: "N", 12: "D", 3: "M", 5: "M", 6: "J", 8: "A"}, Short: ut.CalendarMonthFormatNameValue(nil), Wide: ut.CalendarMonthFormatNameValue{8: "Augusto", 9: "Setembro", 12: "Decembro", 2: "Fevreiro", 3: "Marco", 4: "Abril", 5: "Maio", 6: "Junho", 7: "Julho", 10: "Otubro", 11: "Novembro", 1: "Janeiro"}}, Days: ut.CalendarDayFormatNames{Abbreviated: ut.CalendarDayFormatNameValue{4: "Nai", 5: "Sha", 6: "Sab", 0: "Dim", 1: "Pos", 2: "Pir", 3: "Tat"}, Narrow: ut.CalendarDayFormatNameValue{0: "D", 1: "P", 2: "C", 3: "T", 4: "N", 5: "S", 6: "S"}, Short: ut.CalendarDayFormatNameValue(nil), Wide: ut.CalendarDayFormatNameValue{0: "Dimingu", 1: "Chiposi", 2: "Chipiri", 3: "Chitatu", 4: "Chinai", 5: "Chishanu", 6: "Sabudu"}}, Periods: ut.CalendarPeriodFormatNames{Abbreviated: ut.CalendarPeriodFormatNameValue(nil), Narrow: ut.CalendarPeriodFormatNameValue(nil), Short: ut.CalendarPeriodFormatNameValue(nil), Wide: ut.CalendarPeriodFormatNameValue(nil)}, Eras: ut.Eras{AD: ut.CalendarEraFormatNames{Full: "", Abbrev: "", Narrow: ""}, BC: ut.CalendarEraFormatNames{Full: "Antes de Cristo", Abbrev: "AC", Narrow: ""}}}}
