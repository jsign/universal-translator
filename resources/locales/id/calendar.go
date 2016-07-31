package id

import "github.com/go-playground/universal-translator"

var calendar = ut.Calendar{Formats: ut.CalendarFormats{DateEra: ut.DateEra{BC: ut.CalendarDateFormat{Full: "EEEE, dd MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "dd/MM/yy"}, AD: ut.CalendarDateFormat{Full: "EEEE, dd MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "dd/MM/yy"}}, Time: ut.CalendarDateFormat{Full: "HH.mm.ss zzzz", Long: "HH.mm.ss z", Medium: "HH.mm.ss", Short: "HH.mm"}, DateTime: ut.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"}}, FormatNames: ut.CalendarFormatNames{Months: ut.CalendarMonthFormatNames{Abbreviated: ut.CalendarMonthFormatNameValue{4: "Apr", 5: "Mei", 7: "Jul", 2: "Feb", 3: "Mar", 6: "Jun", 8: "Agt", 9: "Sep", 10: "Okt", 11: "Nov", 12: "Des", 1: "Jan"}, Narrow: ut.CalendarMonthFormatNameValue{6: "J", 9: "S", 12: "D", 2: "F", 4: "A", 5: "M", 8: "A", 10: "O", 11: "N", 1: "J", 3: "M", 7: "J"}, Short: ut.CalendarMonthFormatNameValue(nil), Wide: ut.CalendarMonthFormatNameValue{2: "Februari", 4: "April", 5: "Mei", 7: "Juli", 9: "September", 11: "November", 1: "Januari", 3: "Maret", 6: "Juni", 8: "Agustus", 10: "Oktober", 12: "Desember"}}, Days: ut.CalendarDayFormatNames{Abbreviated: ut.CalendarDayFormatNameValue{0: "Min", 1: "Sen", 2: "Sel", 3: "Rab", 4: "Kam", 5: "Jum", 6: "Sab"}, Narrow: ut.CalendarDayFormatNameValue{0: "M", 1: "S", 2: "S", 3: "R", 4: "K", 5: "J", 6: "S"}, Short: ut.CalendarDayFormatNameValue{0: "Min", 1: "Sen", 2: "Sel", 3: "Rab", 4: "Kam", 5: "Jum", 6: "Sab"}, Wide: ut.CalendarDayFormatNameValue{1: "Senin", 2: "Selasa", 3: "Rabu", 4: "Kamis", 5: "Jumat", 6: "Sabtu", 0: "Minggu"}}, Periods: ut.CalendarPeriodFormatNames{Abbreviated: ut.CalendarPeriodFormatNameValue{"noon": "tengah hari", "pm": "PM", "morning1": "pagi", "afternoon1": "siang", "evening1": "sore", "night1": "malam", "midnight": "tengah malam", "am": "AM"}, Narrow: ut.CalendarPeriodFormatNameValue{"morning1": "pagi", "afternoon1": "siang", "evening1": "sore", "night1": "malam", "midnight": "tengah malam", "am": "AM", "noon": "tengah hari", "pm": "PM"}, Short: ut.CalendarPeriodFormatNameValue(nil), Wide: ut.CalendarPeriodFormatNameValue{"night1": "malam", "midnight": "tengah malam", "am": "AM", "noon": "tengah hari", "pm": "PM", "morning1": "pagi", "afternoon1": "siang", "evening1": "sore"}}, Eras: ut.Eras{AD: ut.CalendarEraFormatNames{Full: "M", Abbrev: "M", Narrow: ""}, BC: ut.CalendarEraFormatNames{Full: "Sebelum Masehi", Abbrev: "SM", Narrow: "SM"}}}}
