package pt_AO

import "github.com/go-playground/universal-translator"

var calendar = ut.Calendar{Formats: ut.CalendarFormats{DateEra: ut.DateEra{BC: ut.CalendarDateFormat{Full: "EEEE, d 'de' MMMM 'de' y", Long: "d 'de' MMMM 'de' y", Medium: "d 'de' MMM 'de' y", Short: "dd/MM/yy"}, AD: ut.CalendarDateFormat{Full: "EEEE, d 'de' MMMM 'de' y", Long: "d 'de' MMMM 'de' y", Medium: "d 'de' MMM 'de' y", Short: "dd/MM/yy"}}, Time: ut.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: ut.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"}}, FormatNames: ut.CalendarFormatNames{Months: ut.CalendarMonthFormatNames{Abbreviated: ut.CalendarMonthFormatNameValue{3: "mar", 8: "ago", 1: "jan", 4: "abr", 11: "nov", 9: "set", 12: "dez", 2: "fev", 5: "mai", 6: "jun", 7: "jul", 10: "out"}, Narrow: ut.CalendarMonthFormatNameValue{2: "F", 5: "M", 6: "J", 12: "D", 3: "M", 7: "J", 9: "S", 4: "A", 8: "A", 10: "O", 11: "N", 1: "J"}, Short: ut.CalendarMonthFormatNameValue{}, Wide: ut.CalendarMonthFormatNameValue{3: "março", 5: "maio", 6: "junho", 11: "novembro", 12: "dezembro", 1: "janeiro", 7: "julho", 2: "fevereiro", 8: "agosto", 9: "setembro", 10: "outubro", 4: "abril"}}, Days: ut.CalendarDayFormatNames{Abbreviated: ut.CalendarDayFormatNameValue{0: "dom", 1: "seg", 2: "ter", 3: "qua", 4: "qui", 5: "sex", 6: "sáb"}, Narrow: ut.CalendarDayFormatNameValue{1: "S", 2: "T", 3: "Q", 4: "Q", 5: "S", 6: "S", 0: "D"}, Short: ut.CalendarDayFormatNameValue{1: "seg", 2: "ter", 3: "qua", 4: "qui", 5: "sex", 6: "sáb", 0: "dom"}, Wide: ut.CalendarDayFormatNameValue{6: "sábado", 0: "domingo", 1: "segunda-feira", 2: "terça-feira", 3: "quarta-feira", 4: "quinta-feira", 5: "sexta-feira"}}, Periods: ut.CalendarPeriodFormatNames{Abbreviated: ut.CalendarPeriodFormatNameValue{"pm": "PM", "morning1": "manhã", "afternoon1": "tarde", "evening1": "noite", "night1": "madrugada", "midnight": "meia-noite", "am": "AM", "noon": "meio-dia"}, Narrow: ut.CalendarPeriodFormatNameValue{"am": "AM", "noon": "meio-dia", "pm": "PM", "morning1": "manhã", "afternoon1": "tarde", "evening1": "noite", "night1": "madrugada", "midnight": "meia-noite"}, Short: ut.CalendarPeriodFormatNameValue{}, Wide: ut.CalendarPeriodFormatNameValue{"morning1": "manhã", "afternoon1": "tarde", "evening1": "noite", "night1": "madrugada", "midnight": "meia-noite", "am": "AM", "noon": "meio-dia", "pm": "PM"}}, Eras: ut.Eras{AD: ut.CalendarEraFormatNames{Full: "", Abbrev: "", Narrow: ""}, BC: ut.CalendarEraFormatNames{Full: "", Abbrev: "", Narrow: ""}}}}