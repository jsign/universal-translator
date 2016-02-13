package en_ZA

import "github.com/go-playground/universal-translator"

var calendar = ut.Calendar{Formats: ut.CalendarFormats{Date: ut.CalendarDateFormat{Full: "EEEE, dd MMMM y", Long: "dd MMMM y", Medium: "dd MMM y", Short: "y/MM/dd"}, Time: ut.CalendarDateFormat{Full: "h:mm:ss a zzzz", Long: "h:mm:ss a z", Medium: "h:mm:ss a", Short: "h:mm a"}, DateTime: ut.CalendarDateFormat{Full: "{1} 'at' {0}", Long: "{1} 'at' {0}", Medium: "{1}, {0}", Short: "{1}, {0}"}}, FormatNames: ut.CalendarFormatNames{Months: ut.CalendarMonthFormatNames{Abbreviated: ut.CalendarMonthFormatNameValue{10: "Oct", 12: "Dec", 1: "Jan", 5: "May", 7: "Jul", 8: "Aug", 9: "Sep", 2: "Feb", 4: "Apr", 11: "Nov", 3: "Mar", 6: "Jun"}, Narrow: ut.CalendarMonthFormatNameValue{2: "F", 6: "J", 7: "J", 5: "M", 11: "N", 8: "A", 10: "O", 1: "J", 3: "M", 4: "A", 9: "S", 12: "D"}, Short: ut.CalendarMonthFormatNameValue{}, Wide: ut.CalendarMonthFormatNameValue{4: "April", 5: "May", 9: "September", 10: "October", 2: "February", 3: "March", 1: "January", 6: "June", 7: "July", 8: "August", 11: "November", 12: "December"}}, Days: ut.CalendarDayFormatNames{Abbreviated: ut.CalendarDayFormatNameValue{6: "Sat", 0: "Sun", 1: "Mon", 2: "Tue", 3: "Wed", 4: "Thu", 5: "Fri"}, Narrow: ut.CalendarDayFormatNameValue{4: "T", 5: "F", 6: "S", 0: "S", 1: "M", 2: "T", 3: "W"}, Short: ut.CalendarDayFormatNameValue{0: "Su", 1: "Mo", 2: "Tu", 3: "We", 4: "Th", 5: "Fr", 6: "Sa"}, Wide: ut.CalendarDayFormatNameValue{2: "Tuesday", 3: "Wednesday", 4: "Thursday", 5: "Friday", 6: "Saturday", 0: "Sunday", 1: "Monday"}}, Periods: ut.CalendarPeriodFormatNames{Abbreviated: ut.CalendarPeriodFormatNameValue{"noon": "noon", "pm": "PM", "morning1": "in the morning", "afternoon1": "in the afternoon", "evening1": "in the evening", "night1": "at night", "midnight": "midnight", "am": "AM"}, Narrow: ut.CalendarPeriodFormatNameValue{"night1": "at night", "midnight": "mi", "am": "a", "noon": "n", "pm": "p", "morning1": "in the morning", "afternoon1": "in the afternoon", "evening1": "in the evening"}, Short: ut.CalendarPeriodFormatNameValue{}, Wide: ut.CalendarPeriodFormatNameValue{"midnight": "midnight", "am": "AM", "noon": "noon", "pm": "PM", "morning1": "morning", "afternoon1": "afternoon", "evening1": "evening", "night1": "night"}}}}