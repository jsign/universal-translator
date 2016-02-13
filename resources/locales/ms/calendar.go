package ms

import "github.com/go-playground/universal-translator"

var calendar = ut.Calendar{Formats: ut.CalendarFormats{Date: ut.CalendarDateFormat{Full: "EEEE, d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "d/MM/yy"}, Time: ut.CalendarDateFormat{Full: "h:mm:ss a zzzz", Long: "h:mm:ss a z", Medium: "h:mm:ss a", Short: "h:mm a"}, DateTime: ut.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"}}, FormatNames: ut.CalendarFormatNames{Months: ut.CalendarMonthFormatNames{Abbreviated: ut.CalendarMonthFormatNameValue{10: "Okt", 1: "Jan", 2: "Feb", 3: "Mac", 7: "Jul", 8: "Ogo", 12: "Dis", 4: "Apr", 5: "Mei", 6: "Jun", 9: "Sep", 11: "Nov"}, Narrow: ut.CalendarMonthFormatNameValue{8: "O", 9: "S", 11: "N", 12: "D", 1: "J", 3: "M", 4: "A", 7: "J", 10: "O", 2: "F", 5: "M", 6: "J"}, Short: ut.CalendarMonthFormatNameValue(nil), Wide: ut.CalendarMonthFormatNameValue{2: "Februari", 4: "April", 7: "Julai", 9: "September", 12: "Disember", 1: "Januari", 5: "Mei", 6: "Jun", 8: "Ogos", 10: "Oktober", 11: "November", 3: "Mac"}}, Days: ut.CalendarDayFormatNames{Abbreviated: ut.CalendarDayFormatNameValue{6: "Sab", 0: "Ahd", 1: "Isn", 2: "Sel", 3: "Rab", 4: "Kha", 5: "Jum"}, Narrow: ut.CalendarDayFormatNameValue{3: "R", 4: "K", 5: "J", 6: "S", 0: "A", 1: "I", 2: "S"}, Short: ut.CalendarDayFormatNameValue{1: "Is", 2: "Se", 3: "Ra", 4: "Kh", 5: "Ju", 6: "Sa", 0: "Ah"}, Wide: ut.CalendarDayFormatNameValue{2: "Selasa", 3: "Rabu", 4: "Khamis", 5: "Jumaat", 6: "Sabtu", 0: "Ahad", 1: "Isnin"}}, Periods: ut.CalendarPeriodFormatNames{Abbreviated: ut.CalendarPeriodFormatNameValue{"morning2": "pagi", "afternoon1": "tengah hari", "evening1": "petang", "night1": "malam", "am": "PG", "pm": "PTG", "morning1": "tengah malam"}, Narrow: ut.CalendarPeriodFormatNameValue{"night1": "malam", "am": "a", "pm": "p", "morning1": "tengah malam", "morning2": "pagi", "afternoon1": "tengah hari", "evening1": "petang"}, Short: ut.CalendarPeriodFormatNameValue(nil), Wide: ut.CalendarPeriodFormatNameValue{"afternoon1": "tengah hari", "evening1": "petang", "night1": "malam", "am": "PG", "pm": "PTG", "morning1": "tengah malam", "morning2": "pagi"}}}}