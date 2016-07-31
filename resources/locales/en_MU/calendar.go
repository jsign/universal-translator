package en_MU

import "github.com/go-playground/universal-translator"

var calendar = ut.Calendar{Formats: ut.CalendarFormats{DateEra: ut.DateEra{BC: ut.CalendarDateFormat{Full: "EEEE, MMMM d, y GGGG", Long: "MMMM d, y GG", Medium: "MMM d, y", Short: "M/d/yy"}, AD: ut.CalendarDateFormat{Full: "EEEE, MMMM d, y", Long: "MMMM d, y", Medium: "MMM d, y", Short: "M/d/yy"}}, Time: ut.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: ut.CalendarDateFormat{Full: "{1} 'at' {0}", Long: "{1} 'at' {0}", Medium: "{1}, {0}", Short: "{1}, {0}"}}, FormatNames: ut.CalendarFormatNames{Months: ut.CalendarMonthFormatNames{Abbreviated: ut.CalendarMonthFormatNameValue{10: "Oct", 11: "Nov", 7: "Jul", 8: "Aug", 1: "Jan", 4: "Apr", 5: "May", 12: "Dec", 2: "Feb", 6: "Jun", 9: "Sep", 3: "Mar"}, Narrow: ut.CalendarMonthFormatNameValue{2: "F", 8: "A", 3: "M", 5: "M", 10: "O", 12: "D", 1: "J", 9: "S", 11: "N", 6: "J", 7: "J", 4: "A"}, Short: ut.CalendarMonthFormatNameValue{}, Wide: ut.CalendarMonthFormatNameValue{4: "April", 11: "November", 10: "October", 2: "February", 6: "June", 7: "July", 8: "August", 9: "September", 5: "May", 12: "December", 1: "January", 3: "March"}}, Days: ut.CalendarDayFormatNames{Abbreviated: ut.CalendarDayFormatNameValue{6: "Sat", 0: "Sun", 1: "Mon", 2: "Tue", 3: "Wed", 4: "Thu", 5: "Fri"}, Narrow: ut.CalendarDayFormatNameValue{3: "W", 4: "T", 5: "F", 6: "S", 0: "S", 1: "M", 2: "T"}, Short: ut.CalendarDayFormatNameValue{0: "Su", 1: "Mo", 2: "Tu", 3: "We", 4: "Th", 5: "Fr", 6: "Sa"}, Wide: ut.CalendarDayFormatNameValue{5: "Friday", 6: "Saturday", 0: "Sunday", 1: "Monday", 2: "Tuesday", 3: "Wednesday", 4: "Thursday"}}, Periods: ut.CalendarPeriodFormatNames{Abbreviated: ut.CalendarPeriodFormatNameValue{"evening1": "in the evening", "night1": "at night", "midnight": "midnight", "am": "AM", "noon": "noon", "pm": "PM", "morning1": "in the morning", "afternoon1": "in the afternoon"}, Narrow: ut.CalendarPeriodFormatNameValue{"pm": "p", "morning1": "in the morning", "afternoon1": "in the afternoon", "evening1": "in the evening", "night1": "at night", "midnight": "mi", "am": "a", "noon": "n"}, Short: ut.CalendarPeriodFormatNameValue{}, Wide: ut.CalendarPeriodFormatNameValue{"am": "AM", "noon": "noon", "pm": "PM", "morning1": "morning", "afternoon1": "afternoon", "evening1": "evening", "night1": "night", "midnight": "midnight"}}, Eras: ut.Eras{AD: ut.CalendarEraFormatNames{Full: "", Abbrev: "", Narrow: ""}, BC: ut.CalendarEraFormatNames{Full: "", Abbrev: "", Narrow: ""}}}}