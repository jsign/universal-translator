package sw_UG

import "github.com/go-playground/universal-translator"

var calendar = ut.Calendar{Formats: ut.CalendarFormats{Date: ut.CalendarDateFormat{Full: "EEEE, d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "dd/MM/y"}, Time: ut.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: ut.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"}}, FormatNames: ut.CalendarFormatNames{Months: ut.CalendarMonthFormatNames{Abbreviated: ut.CalendarMonthFormatNameValue{3: "Mac", 9: "Sep", 5: "Mei", 10: "Okt", 1: "Jan", 2: "Feb", 6: "Jun", 12: "Des", 7: "Jul", 8: "Ago", 4: "Apr", 11: "Nov"}, Narrow: ut.CalendarMonthFormatNameValue{5: "M", 6: "J", 8: "A", 4: "A", 7: "J", 9: "S", 11: "N", 1: "J", 12: "D", 2: "F", 3: "M", 10: "O"}, Short: ut.CalendarMonthFormatNameValue{}, Wide: ut.CalendarMonthFormatNameValue{1: "Januari", 7: "Julai", 11: "Novemba", 6: "Juni", 9: "Septemba", 4: "Aprili", 8: "Agosti", 5: "Mei", 10: "Oktoba", 12: "Desemba", 2: "Februari", 3: "Machi"}}, Days: ut.CalendarDayFormatNames{Abbreviated: ut.CalendarDayFormatNameValue{6: "Jumamosi", 0: "Jumapili", 1: "Jumatatu", 2: "Jumanne", 3: "Jumatano", 4: "Alhamisi", 5: "Ijumaa"}, Narrow: ut.CalendarDayFormatNameValue{3: "W", 4: "T", 5: "F", 6: "S", 0: "S", 1: "M", 2: "T"}, Short: ut.CalendarDayFormatNameValue{1: "Jumatatu", 2: "Jumanne", 3: "Jumatano", 4: "Alhamisi", 5: "Ijumaa", 6: "Jumamosi", 0: "Jumapili"}, Wide: ut.CalendarDayFormatNameValue{3: "Jumatano", 4: "Alhamisi", 5: "Ijumaa", 6: "Jumamosi", 0: "Jumapili", 1: "Jumatatu", 2: "Jumanne"}}, Periods: ut.CalendarPeriodFormatNames{Abbreviated: ut.CalendarPeriodFormatNameValue{"noon": "saa sita za mchana", "midnight": "saa sita za usiku", "morning1": "alfajiri", "afternoon1": "mchana", "evening1": "jioni", "am": "AM", "morning2": "asubuhi", "night1": "usiku", "pm": "PM"}, Narrow: ut.CalendarPeriodFormatNameValue{"afternoon1": "mchana", "night1": "usiku", "am": "am", "noon": "saa sita za mchana", "pm": "pm", "morning2": "asubuhi", "midnight": "saa sita za usiku", "morning1": "alfajiri", "evening1": "jioni"}, Short: ut.CalendarPeriodFormatNameValue{}, Wide: ut.CalendarPeriodFormatNameValue{"am": "AM", "morning1": "alfajiri", "pm": "PM", "midnight": "saa sita za usiku", "afternoon1": "mchana", "evening1": "jioni", "night1": "usiku", "morning2": "asubuhi", "noon": "saa sita za mchana"}}}}