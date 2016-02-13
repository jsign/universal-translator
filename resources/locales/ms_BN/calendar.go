package ms_BN

import "github.com/go-playground/universal-translator"

var calendar = ut.Calendar{Formats: ut.CalendarFormats{Date: ut.CalendarDateFormat{Full: "dd MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "d/MM/yy"}, Time: ut.CalendarDateFormat{Full: "h:mm:ss a zzzz", Long: "h:mm:ss a z", Medium: "h:mm:ss a", Short: "h:mm a"}, DateTime: ut.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"}}, FormatNames: ut.CalendarFormatNames{Months: ut.CalendarMonthFormatNames{Abbreviated: ut.CalendarMonthFormatNameValue{7: "Jul", 10: "Okt", 2: "Feb", 9: "Sep", 11: "Nov", 5: "Mei", 3: "Mac", 1: "Jan", 6: "Jun", 12: "Dis", 4: "Apr", 8: "Ogo"}, Narrow: ut.CalendarMonthFormatNameValue{5: "M", 6: "J", 11: "N", 12: "D", 4: "A", 9: "S", 2: "F", 1: "J", 3: "M", 8: "O", 7: "J", 10: "O"}, Short: ut.CalendarMonthFormatNameValue{}, Wide: ut.CalendarMonthFormatNameValue{8: "Ogos", 10: "Oktober", 5: "Mei", 4: "April", 7: "Julai", 9: "September", 12: "Disember", 1: "Januari", 6: "Jun", 11: "November", 3: "Mac", 2: "Februari"}}, Days: ut.CalendarDayFormatNames{Abbreviated: ut.CalendarDayFormatNameValue{4: "Kha", 5: "Jum", 6: "Sab", 0: "Ahd", 1: "Isn", 2: "Sel", 3: "Rab"}, Narrow: ut.CalendarDayFormatNameValue{0: "A", 1: "I", 2: "S", 3: "R", 4: "K", 5: "J", 6: "S"}, Short: ut.CalendarDayFormatNameValue{0: "Ah", 1: "Is", 2: "Se", 3: "Ra", 4: "Kh", 5: "Ju", 6: "Sa"}, Wide: ut.CalendarDayFormatNameValue{1: "Isnin", 2: "Selasa", 3: "Rabu", 4: "Khamis", 5: "Jumaat", 6: "Sabtu", 0: "Ahad"}}, Periods: ut.CalendarPeriodFormatNames{Abbreviated: ut.CalendarPeriodFormatNameValue{"morning2": "pagi", "afternoon1": "tengah hari", "evening1": "petang", "night1": "malam", "am": "PG", "pm": "PTG", "morning1": "tengah malam"}, Narrow: ut.CalendarPeriodFormatNameValue{"am": "a", "pm": "p", "morning1": "tengah malam", "morning2": "pagi", "afternoon1": "tengah hari", "evening1": "petang", "night1": "malam"}, Short: ut.CalendarPeriodFormatNameValue{}, Wide: ut.CalendarPeriodFormatNameValue{"pm": "PTG", "morning1": "tengah malam", "morning2": "pagi", "afternoon1": "tengah hari", "evening1": "petang", "night1": "malam", "am": "PG"}}}}