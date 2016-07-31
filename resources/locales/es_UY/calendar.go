package es_UY

import "github.com/go-playground/universal-translator"

var calendar = ut.Calendar{Formats: ut.CalendarFormats{DateEra: ut.DateEra{BC: ut.CalendarDateFormat{Full: "EEEE, d 'de' MMMM 'de' y", Long: "d 'de' MMMM 'de' y", Medium: "d MMM y", Short: "d/M/yy"}, AD: ut.CalendarDateFormat{Full: "EEEE, d 'de' MMMM 'de' y", Long: "d 'de' MMMM 'de' y", Medium: "d MMM y", Short: "d/M/yy"}}, Time: ut.CalendarDateFormat{Full: "H:mm:ss (zzzz)", Long: "H:mm:ss z", Medium: "H:mm:ss", Short: "H:mm"}, DateTime: ut.CalendarDateFormat{Full: "{1}, {0}", Long: "{1}, {0}", Medium: "{1} {0}", Short: "{1} {0}"}}, FormatNames: ut.CalendarFormatNames{Months: ut.CalendarMonthFormatNames{Abbreviated: ut.CalendarMonthFormatNameValue{7: "Jul.", 8: "Ago.", 9: "Set.", 12: "Dic.", 2: "Feb.", 5: "May.", 6: "Jun.", 10: "Oct.", 11: "Nov.", 1: "Ene.", 3: "Mar.", 4: "Abr."}, Narrow: ut.CalendarMonthFormatNameValue{4: "A", 5: "M", 9: "S", 2: "F", 6: "J", 11: "N", 7: "J", 10: "O", 12: "D", 1: "E", 8: "A", 3: "M"}, Short: ut.CalendarMonthFormatNameValue{}, Wide: ut.CalendarMonthFormatNameValue{5: "Mayo", 11: "Noviembre", 12: "Diciembre", 6: "Junio", 7: "Julio", 8: "Agosto", 9: "Setiembre", 1: "Enero", 2: "Febrero", 3: "Marzo", 4: "Abril", 10: "Octubre"}}, Days: ut.CalendarDayFormatNames{Abbreviated: ut.CalendarDayFormatNameValue{6: "sáb.", 0: "dom.", 1: "lun.", 2: "mar.", 3: "mié.", 4: "jue.", 5: "vie."}, Narrow: ut.CalendarDayFormatNameValue{0: "D", 1: "L", 2: "M", 3: "X", 4: "J", 5: "V", 6: "S"}, Short: ut.CalendarDayFormatNameValue{5: "VI", 6: "SA", 0: "DO", 1: "LU", 2: "MA", 3: "MI", 4: "JU"}, Wide: ut.CalendarDayFormatNameValue{4: "jueves", 5: "viernes", 6: "sábado", 0: "domingo", 1: "lunes", 2: "martes", 3: "miércoles"}}, Periods: ut.CalendarPeriodFormatNames{Abbreviated: ut.CalendarPeriodFormatNameValue{"morning2": "mañana", "evening1": "tarde", "night1": "noche", "am": "a. m.", "noon": "mediodía", "pm": "p. m.", "morning1": "madrugada"}, Narrow: ut.CalendarPeriodFormatNameValue{"evening1": "tarde", "night1": "noche", "am": "a. m.", "noon": "mediodía", "pm": "p. m.", "morning1": "madrugada", "morning2": "mañana"}, Short: ut.CalendarPeriodFormatNameValue{}, Wide: ut.CalendarPeriodFormatNameValue{"pm": "p. m.", "morning1": "madrugada", "morning2": "mañana", "evening1": "tarde", "night1": "noche", "am": "a. m.", "noon": "mediodía"}}, Eras: ut.Eras{AD: ut.CalendarEraFormatNames{Full: "", Abbrev: "", Narrow: ""}, BC: ut.CalendarEraFormatNames{Full: "", Abbrev: "", Narrow: ""}}}}
