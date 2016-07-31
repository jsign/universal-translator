package en_AU

import "github.com/go-playground/universal-translator"

var calendar = ut.Calendar{Formats: ut.CalendarFormats{DateEra: ut.DateEra{BC: ut.CalendarDateFormat{Full: "EEEE, MMMM d, y GGGG", Long: "MMMM d, y GG", Medium: "MMM d, y", Short: "d/M/yy"}, AD: ut.CalendarDateFormat{Full: "EEEE, MMMM d, y", Long: "MMMM d, y", Medium: "MMM d, y", Short: "d/M/yy"}}, Time: ut.CalendarDateFormat{Full: "h:mm:ss a zzzz", Long: "h:mm:ss a z", Medium: "h:mm:ss a", Short: "h:mm a"}, DateTime: ut.CalendarDateFormat{Full: "{1} 'at' {0}", Long: "{1} 'at' {0}", Medium: "{1}, {0}", Short: "{1}, {0}"}}, FormatNames: ut.CalendarFormatNames{Months: ut.CalendarMonthFormatNames{Abbreviated: ut.CalendarMonthFormatNameValue{1: "Jan.", 3: "Mar.", 4: "Apr.", 7: "Jul.", 8: "Aug.", 11: "Nov.", 12: "Dec.", 2: "Feb.", 5: "May", 6: "Jun.", 9: "Sep.", 10: "Oct."}, Narrow: ut.CalendarMonthFormatNameValue{1: "Jan.", 2: "Feb.", 3: "Mar.", 5: "May", 6: "Jun.", 7: "Jul.", 9: "Sep.", 10: "Oct.", 4: "Apr.", 8: "Aug.", 11: "Nov.", 12: "Dec."}, Short: ut.CalendarMonthFormatNameValue{}, Wide: ut.CalendarMonthFormatNameValue{5: "May", 11: "November", 12: "December", 3: "March", 8: "August", 9: "September", 10: "October", 1: "January", 4: "April", 2: "February", 6: "June", 7: "July"}}, Days: ut.CalendarDayFormatNames{Abbreviated: ut.CalendarDayFormatNameValue{0: "Sun.", 1: "Mon.", 2: "Tue.", 3: "Wed.", 4: "Thu.", 5: "Fri.", 6: "Sat."}, Narrow: ut.CalendarDayFormatNameValue{6: "Sa.", 0: "Su.", 1: "M.", 2: "Tu.", 3: "W.", 4: "Th.", 5: "F."}, Short: ut.CalendarDayFormatNameValue{4: "Th.", 5: "Fri.", 6: "Sat.", 0: "Su.", 1: "Mon.", 2: "Tu.", 3: "Wed."}, Wide: ut.CalendarDayFormatNameValue{3: "Wednesday", 4: "Thursday", 5: "Friday", 6: "Saturday", 0: "Sunday", 1: "Monday", 2: "Tuesday"}}, Periods: ut.CalendarPeriodFormatNames{Abbreviated: ut.CalendarPeriodFormatNameValue{"noon": "midday", "pm": "pm", "morning1": "in the morning", "afternoon1": "in the afternoon", "evening1": "in the evening", "night1": "at night", "midnight": "midnight", "am": "am"}, Narrow: ut.CalendarPeriodFormatNameValue{"pm": "pm", "night1": "at night", "midnight": "mi", "morning1": "in the morning", "afternoon1": "in the afternoon", "evening1": "in the evening", "am": "am", "noon": "midday"}, Short: ut.CalendarPeriodFormatNameValue{}, Wide: ut.CalendarPeriodFormatNameValue{"evening1": "evening", "night1": "night", "midnight": "midnight", "morning1": "morning", "afternoon1": "afternoon", "am": "am", "noon": "midday", "pm": "pm"}}, Eras: ut.Eras{AD: ut.CalendarEraFormatNames{Full: "", Abbrev: "", Narrow: ""}, BC: ut.CalendarEraFormatNames{Full: "", Abbrev: "", Narrow: ""}}}}