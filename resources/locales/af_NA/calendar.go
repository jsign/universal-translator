package af_NA

import "github.com/go-playground/universal-translator"

var calendar = ut.Calendar{Formats: ut.CalendarFormats{DateEra: ut.DateEra{BC: ut.CalendarDateFormat{Full: "EEEE d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "y-MM-dd"}, AD: ut.CalendarDateFormat{Full: "EEEE d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "y-MM-dd"}}, Time: ut.CalendarDateFormat{Full: "h:mm:ss a zzzz", Long: "h:mm:ss a z", Medium: "h:mm:ss a", Short: "h:mm a"}, DateTime: ut.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"}}, FormatNames: ut.CalendarFormatNames{Months: ut.CalendarMonthFormatNames{Abbreviated: ut.CalendarMonthFormatNameValue{12: "Des.", 3: "Mrt.", 5: "Mei", 9: "Sep.", 10: "Okt.", 11: "Nov.", 2: "Feb.", 4: "Apr.", 7: "Jul.", 8: "Aug.", 6: "Jun.", 1: "Jan."}, Narrow: ut.CalendarMonthFormatNameValue{1: "J", 3: "M", 4: "A", 5: "M", 11: "N", 6: "J", 7: "J", 8: "A", 9: "S", 12: "D", 10: "O", 2: "F"}, Short: ut.CalendarMonthFormatNameValue{}, Wide: ut.CalendarMonthFormatNameValue{2: "Februarie", 9: "September", 6: "Junie", 12: "Desember", 7: "Julie", 10: "Oktober", 11: "November", 1: "Januarie", 3: "Maart", 8: "Augustus", 4: "April", 5: "Mei"}}, Days: ut.CalendarDayFormatNames{Abbreviated: ut.CalendarDayFormatNameValue{3: "Wo.", 4: "Do.", 5: "Vr.", 6: "Sa.", 0: "So.", 1: "Ma.", 2: "Di."}, Narrow: ut.CalendarDayFormatNameValue{1: "M", 2: "D", 3: "W", 4: "D", 5: "V", 6: "S", 0: "S"}, Short: ut.CalendarDayFormatNameValue{3: "Wo.", 4: "Do.", 5: "Vr.", 6: "Sa.", 0: "So.", 1: "Ma.", 2: "Di."}, Wide: ut.CalendarDayFormatNameValue{2: "Dinsdag", 3: "Woensdag", 4: "Donderdag", 5: "Vrydag", 6: "Saterdag", 0: "Sondag", 1: "Maandag"}}, Periods: ut.CalendarPeriodFormatNames{Abbreviated: ut.CalendarPeriodFormatNameValue{"midnight": "middernag", "am": "vm.", "pm": "nm.", "morning1": "oggend", "afternoon1": "middag", "evening1": "aand", "night1": "nag"}, Narrow: ut.CalendarPeriodFormatNameValue{"midnight": "mn", "am": "v", "pm": "n", "morning1": "o", "afternoon1": "m", "evening1": "a", "night1": "n"}, Short: ut.CalendarPeriodFormatNameValue{}, Wide: ut.CalendarPeriodFormatNameValue{"am": "vm.", "pm": "nm.", "morning1": "oggend", "afternoon1": "middag", "evening1": "aand", "night1": "nag", "midnight": "middernag"}}, Eras: ut.Eras{AD: ut.CalendarEraFormatNames{Full: "", Abbrev: "", Narrow: ""}, BC: ut.CalendarEraFormatNames{Full: "", Abbrev: "", Narrow: ""}}}}
