package af

import "github.com/go-playground/universal-translator"

var calendar = ut.Calendar{Formats: ut.CalendarFormats{DateEra: ut.DateEra{BC: ut.CalendarDateFormat{Full: "EEEE, dd MMMM y", Long: "dd MMMM y", Medium: "dd MMM y", Short: "y-MM-dd"}, AD: ut.CalendarDateFormat{Full: "EEEE, dd MMMM y", Long: "dd MMMM y", Medium: "dd MMM y", Short: "y-MM-dd"}}, Time: ut.CalendarDateFormat{Full: "h:mm:ss a zzzz", Long: "h:mm:ss a z", Medium: "h:mm:ss a", Short: "h:mm a"}, DateTime: ut.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"}}, FormatNames: ut.CalendarFormatNames{Months: ut.CalendarMonthFormatNames{Abbreviated: ut.CalendarMonthFormatNameValue{9: "Sep.", 10: "Okt.", 11: "Nov.", 12: "Des.", 3: "Mrt.", 5: "Mei", 6: "Jun.", 7: "Jul.", 8: "Aug.", 1: "Jan.", 2: "Feb.", 4: "Apr."}, Narrow: ut.CalendarMonthFormatNameValue{1: "J", 2: "F", 3: "M", 4: "A", 5: "M", 6: "J", 7: "J", 10: "O", 8: "A", 9: "S", 11: "N", 12: "D"}, Short: ut.CalendarMonthFormatNameValue(nil), Wide: ut.CalendarMonthFormatNameValue{12: "Desember", 1: "Januarie", 2: "Februarie", 3: "Maart", 7: "Julie", 10: "Oktober", 11: "November", 4: "April", 5: "Mei", 6: "Junie", 8: "Augustus", 9: "September"}}, Days: ut.CalendarDayFormatNames{Abbreviated: ut.CalendarDayFormatNameValue{2: "Di.", 3: "Wo.", 4: "Do.", 5: "Vr.", 6: "Sa.", 0: "So.", 1: "Ma."}, Narrow: ut.CalendarDayFormatNameValue{3: "W", 4: "D", 5: "V", 6: "S", 0: "S", 1: "M", 2: "D"}, Short: ut.CalendarDayFormatNameValue{1: "Ma.", 2: "Di.", 3: "Wo.", 4: "Do.", 5: "Vr.", 6: "Sa.", 0: "So."}, Wide: ut.CalendarDayFormatNameValue{5: "Vrydag", 6: "Saterdag", 0: "Sondag", 1: "Maandag", 2: "Dinsdag", 3: "Woensdag", 4: "Donderdag"}}, Periods: ut.CalendarPeriodFormatNames{Abbreviated: ut.CalendarPeriodFormatNameValue{"pm": "nm.", "morning1": "oggend", "afternoon1": "middag", "evening1": "aand", "night1": "nag", "midnight": "middernag", "am": "vm."}, Narrow: ut.CalendarPeriodFormatNameValue{"afternoon1": "m", "evening1": "a", "night1": "n", "midnight": "mn", "am": "v", "pm": "n", "morning1": "o"}, Short: ut.CalendarPeriodFormatNameValue(nil), Wide: ut.CalendarPeriodFormatNameValue{"midnight": "middernag", "am": "vm.", "pm": "nm.", "morning1": "oggend", "afternoon1": "middag", "evening1": "aand", "night1": "nag"}}, Eras: ut.Eras{AD: ut.CalendarEraFormatNames{Full: "na Christus", Abbrev: "n.C.", Narrow: "n.C."}, BC: ut.CalendarEraFormatNames{Full: "voor Christus", Abbrev: "v.C.", Narrow: "v.C."}}}}