package ro

import "github.com/go-playground/universal-translator"

var calendar = ut.Calendar{Formats: ut.CalendarFormats{DateEra: ut.DateEra{BC: ut.CalendarDateFormat{Full: "EEEE, d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "dd.MM.y"}, AD: ut.CalendarDateFormat{Full: "EEEE, d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "dd.MM.y"}}, Time: ut.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: ut.CalendarDateFormat{Full: "{1}, {0}", Long: "{1}, {0}", Medium: "{1}, {0}", Short: "{1}, {0}"}}, FormatNames: ut.CalendarFormatNames{Months: ut.CalendarMonthFormatNames{Abbreviated: ut.CalendarMonthFormatNameValue{6: "iun.", 12: "dec.", 2: "feb.", 4: "apr.", 5: "mai", 7: "iul.", 8: "aug.", 9: "sept.", 10: "oct.", 11: "nov.", 1: "ian.", 3: "mar."}, Narrow: ut.CalendarMonthFormatNameValue{7: "I", 9: "S", 11: "N", 12: "D", 1: "I", 3: "M", 5: "M", 6: "I", 8: "A", 10: "O", 2: "F", 4: "A"}, Short: ut.CalendarMonthFormatNameValue(nil), Wide: ut.CalendarMonthFormatNameValue{6: "iunie", 7: "iulie", 8: "august", 12: "decembrie", 1: "ianuarie", 5: "mai", 4: "aprilie", 9: "septembrie", 10: "octombrie", 11: "noiembrie", 2: "februarie", 3: "martie"}}, Days: ut.CalendarDayFormatNames{Abbreviated: ut.CalendarDayFormatNameValue{5: "vin.", 6: "sâm.", 0: "dum.", 1: "lun.", 2: "mar.", 3: "mie.", 4: "joi"}, Narrow: ut.CalendarDayFormatNameValue{5: "V", 6: "S", 0: "D", 1: "L", 2: "M", 3: "M", 4: "J"}, Short: ut.CalendarDayFormatNameValue{5: "vi", 6: "sâ", 0: "du", 1: "lu", 2: "ma", 3: "mi", 4: "jo"}, Wide: ut.CalendarDayFormatNameValue{5: "vineri", 6: "sâmbătă", 0: "duminică", 1: "luni", 2: "marți", 3: "miercuri", 4: "joi"}}, Periods: ut.CalendarPeriodFormatNames{Abbreviated: ut.CalendarPeriodFormatNameValue{"midnight": "miezul nopții", "am": "a.m.", "noon": "amiază", "pm": "p.m."}, Narrow: ut.CalendarPeriodFormatNameValue{"midnight": "miezul nopții", "am": "a.m.", "pm": "p.m."}, Short: ut.CalendarPeriodFormatNameValue(nil), Wide: ut.CalendarPeriodFormatNameValue{"afternoon1": "după-amiază", "evening1": "seară", "night1": "noapte", "midnight": "la miezul nopții", "am": "a.m.", "noon": "la amiază", "pm": "p.m.", "morning1": "dimineață"}}, Eras: ut.Eras{AD: ut.CalendarEraFormatNames{Full: "după Hristos", Abbrev: "d.Hr.", Narrow: ""}, BC: ut.CalendarEraFormatNames{Full: "înainte de Hristos", Abbrev: "î.Hr.", Narrow: ""}}}}