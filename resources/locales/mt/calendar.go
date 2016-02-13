package mt

import "github.com/go-playground/universal-translator"

var calendar = ut.Calendar{Formats: ut.CalendarFormats{Date: ut.CalendarDateFormat{Full: "EEEE, d 'ta'’ MMMM y", Long: "d 'ta'’ MMMM y", Medium: "dd MMM y", Short: "dd/MM/y"}, Time: ut.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: ut.CalendarDateFormat{Full: "", Long: "", Medium: "", Short: ""}}, FormatNames: ut.CalendarFormatNames{Months: ut.CalendarMonthFormatNames{Abbreviated: ut.CalendarMonthFormatNameValue{11: "Nov", 9: "Set", 10: "Ott", 3: "Mar", 4: "Apr", 5: "Mej", 6: "Ġun", 7: "Lul", 8: "Aww", 1: "Jan", 2: "Fra", 12: "Diċ"}, Narrow: ut.CalendarMonthFormatNameValue{5: "Mj", 8: "Aw", 3: "Mz", 4: "Ap", 6: "Ġn", 7: "Lj", 9: "St", 10: "Ob", 1: "Jn", 2: "Fr", 11: "Nv", 12: "Dċ"}, Short: ut.CalendarMonthFormatNameValue(nil), Wide: ut.CalendarMonthFormatNameValue{4: "April", 8: "Awwissu", 9: "Settembru", 2: "Frar", 3: "Marzu", 6: "Ġunju", 7: "Lulju", 10: "Ottubru", 11: "Novembru", 12: "Diċembru", 1: "Jannar", 5: "Mejju"}}, Days: ut.CalendarDayFormatNames{Abbreviated: ut.CalendarDayFormatNameValue{2: "Tli", 3: "Erb", 4: "Ħam", 5: "Ġim", 6: "Sib", 0: "Ħad", 1: "Tne"}, Narrow: ut.CalendarDayFormatNameValue{3: "Er", 4: "Ħm", 5: "Ġm", 6: "Sb", 0: "Ħd", 1: "Tn", 2: "Tl"}, Short: ut.CalendarDayFormatNameValue(nil), Wide: ut.CalendarDayFormatNameValue{1: "It-Tnejn", 2: "It-Tlieta", 3: "L-Erbgħa", 4: "Il-Ħamis", 5: "Il-Ġimgħa", 6: "Is-Sibt", 0: "Il-Ħadd"}}, Periods: ut.CalendarPeriodFormatNames{Abbreviated: ut.CalendarPeriodFormatNameValue(nil), Narrow: ut.CalendarPeriodFormatNameValue(nil), Short: ut.CalendarPeriodFormatNameValue(nil), Wide: ut.CalendarPeriodFormatNameValue{"am": "AM", "pm": "PM"}}}}