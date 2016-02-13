package hy

import "github.com/go-playground/universal-translator"

var calendar = ut.Calendar{Formats: ut.CalendarFormats{Date: ut.CalendarDateFormat{Full: "yթ. MMMM d, EEEE", Long: "dd MMMM, yթ.", Medium: "dd MMM, yթ.", Short: "dd.MM.yy"}, Time: ut.CalendarDateFormat{Full: "H:mm:ss, zzzz", Long: "H:mm:ss, z", Medium: "H:mm:ss", Short: "H:mm"}, DateTime: ut.CalendarDateFormat{Full: "{1}, {0}", Long: "{1}, {0}", Medium: "{1}, {0}", Short: "{1}, {0}"}}, FormatNames: ut.CalendarFormatNames{Months: ut.CalendarMonthFormatNames{Abbreviated: ut.CalendarMonthFormatNameValue{1: "հնվ", 2: "փտվ", 3: "մրտ", 4: "ապր", 5: "մյս", 6: "հնս", 7: "հլս", 8: "օգս", 11: "նոյ", 12: "դեկ", 9: "սեպ", 10: "հոկ"}, Narrow: ut.CalendarMonthFormatNameValue{3: "Մ", 4: "Ա", 5: "Մ", 8: "Օ", 10: "Հ", 12: "Դ", 1: "Հ", 2: "Փ", 6: "Հ", 7: "Հ", 9: "Ս", 11: "Ն"}, Short: ut.CalendarMonthFormatNameValue(nil), Wide: ut.CalendarMonthFormatNameValue{11: "նոյեմբեր", 1: "հունվար", 2: "փետրվար", 6: "հունիս", 7: "հուլիս", 10: "հոկտեմբեր", 12: "դեկտեմբեր", 3: "մարտ", 4: "ապրիլ", 5: "մայիս", 8: "օգոստոս", 9: "սեպտեմբեր"}}, Days: ut.CalendarDayFormatNames{Abbreviated: ut.CalendarDayFormatNameValue{4: "հնգ", 5: "ուր", 6: "շբթ", 0: "կիր", 1: "երկ", 2: "երք", 3: "չրք"}, Narrow: ut.CalendarDayFormatNameValue{0: "Կր", 1: "Եկ", 2: "Եր", 3: "Չր", 4: "Հգ", 5: "Ու", 6: "Շբ"}, Short: ut.CalendarDayFormatNameValue{6: "շբթ", 0: "կիր", 1: "երկ", 2: "երք", 3: "չրք", 4: "հնգ", 5: "ուր"}, Wide: ut.CalendarDayFormatNameValue{3: "չորեքշաբթի", 4: "հինգշաբթի", 5: "ուրբաթ", 6: "շաբաթ", 0: "կիրակի", 1: "երկուշաբթի", 2: "երեքշաբթի"}}, Periods: ut.CalendarPeriodFormatNames{Abbreviated: ut.CalendarPeriodFormatNameValue{"pm": "PM", "am": "AM"}, Narrow: ut.CalendarPeriodFormatNameValue{"noon": "կեսօր", "pm": "PM", "morning1": "առավոտ", "afternoon1": "ցերեկ", "evening1": "երեկո", "night1": "գիշեր", "midnight": "կեսգիշեր", "am": "AM"}, Short: ut.CalendarPeriodFormatNameValue(nil), Wide: ut.CalendarPeriodFormatNameValue{"pm": "PM", "morning1": "առավոտ", "afternoon1": "ցերեկ", "evening1": "երեկո", "night1": "գիշեր", "midnight": "կեսգիշեր", "am": "AM", "noon": "կեսօր"}}}}