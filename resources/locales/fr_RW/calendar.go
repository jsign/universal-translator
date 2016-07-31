package fr_RW

import "github.com/go-playground/universal-translator"

var calendar = ut.Calendar{Formats: ut.CalendarFormats{DateEra: ut.DateEra{BC: ut.CalendarDateFormat{Full: "EEEE d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "dd/MM/y"}, AD: ut.CalendarDateFormat{Full: "EEEE d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "dd/MM/y"}}, Time: ut.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: ut.CalendarDateFormat{Full: "{1} 'à' {0}", Long: "{1} 'à' {0}", Medium: "{1} 'à' {0}", Short: "{1} {0}"}}, FormatNames: ut.CalendarFormatNames{Months: ut.CalendarMonthFormatNames{Abbreviated: ut.CalendarMonthFormatNameValue{12: "déc.", 9: "sept.", 5: "mai", 6: "juin", 1: "janv.", 2: "févr.", 8: "août", 10: "oct.", 11: "nov.", 4: "avr.", 7: "juil.", 3: "mars"}, Narrow: ut.CalendarMonthFormatNameValue{6: "J", 9: "S", 3: "M", 4: "A", 8: "A", 5: "M", 7: "J", 1: "J", 2: "F", 10: "O", 11: "N", 12: "D"}, Short: ut.CalendarMonthFormatNameValue{}, Wide: ut.CalendarMonthFormatNameValue{5: "mai", 10: "octobre", 11: "novembre", 2: "février", 4: "avril", 6: "juin", 8: "août", 9: "septembre", 3: "mars", 12: "décembre", 1: "janvier", 7: "juillet"}}, Days: ut.CalendarDayFormatNames{Abbreviated: ut.CalendarDayFormatNameValue{4: "jeu.", 5: "ven.", 6: "sam.", 0: "dim.", 1: "lun.", 2: "mar.", 3: "mer."}, Narrow: ut.CalendarDayFormatNameValue{6: "S", 0: "D", 1: "L", 2: "M", 3: "M", 4: "J", 5: "V"}, Short: ut.CalendarDayFormatNameValue{4: "je", 5: "ve", 6: "sa", 0: "di", 1: "lu", 2: "ma", 3: "me"}, Wide: ut.CalendarDayFormatNameValue{1: "lundi", 2: "mardi", 3: "mercredi", 4: "jeudi", 5: "vendredi", 6: "samedi", 0: "dimanche"}}, Periods: ut.CalendarPeriodFormatNames{Abbreviated: ut.CalendarPeriodFormatNameValue{"afternoon1": "ap.m.", "evening1": "soir", "night1": "nuit", "midnight": "min.", "am": "AM", "noon": "midi", "pm": "PM", "morning1": "mat."}, Narrow: ut.CalendarPeriodFormatNameValue{"afternoon1": "ap.m.", "evening1": "soir", "night1": "nuit", "midnight": "min.", "am": "AM", "noon": "midi", "pm": "PM", "morning1": "mat."}, Short: ut.CalendarPeriodFormatNameValue{}, Wide: ut.CalendarPeriodFormatNameValue{"morning1": "matin", "afternoon1": "après-midi", "evening1": "soir", "night1": "nuit", "midnight": "minuit", "am": "AM", "noon": "midi", "pm": "PM"}}, Eras: ut.Eras{AD: ut.CalendarEraFormatNames{Full: "", Abbrev: "", Narrow: ""}, BC: ut.CalendarEraFormatNames{Full: "", Abbrev: "", Narrow: ""}}}}