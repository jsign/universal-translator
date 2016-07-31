package fr_DJ

import "github.com/go-playground/universal-translator"

var calendar = ut.Calendar{Formats: ut.CalendarFormats{DateEra: ut.DateEra{BC: ut.CalendarDateFormat{Full: "EEEE d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "dd/MM/y"}, AD: ut.CalendarDateFormat{Full: "EEEE d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "dd/MM/y"}}, Time: ut.CalendarDateFormat{Full: "h:mm:ss a zzzz", Long: "h:mm:ss a z", Medium: "h:mm:ss a", Short: "h:mm a"}, DateTime: ut.CalendarDateFormat{Full: "{1} 'à' {0}", Long: "{1} 'à' {0}", Medium: "{1} 'à' {0}", Short: "{1} {0}"}}, FormatNames: ut.CalendarFormatNames{Months: ut.CalendarMonthFormatNames{Abbreviated: ut.CalendarMonthFormatNameValue{2: "févr.", 10: "oct.", 12: "déc.", 3: "mars", 7: "juil.", 8: "août", 5: "mai", 9: "sept.", 1: "janv.", 4: "avr.", 11: "nov.", 6: "juin"}, Narrow: ut.CalendarMonthFormatNameValue{5: "M", 6: "J", 7: "J", 8: "A", 12: "D", 2: "F", 4: "A", 3: "M", 9: "S", 10: "O", 11: "N", 1: "J"}, Short: ut.CalendarMonthFormatNameValue{}, Wide: ut.CalendarMonthFormatNameValue{5: "mai", 10: "octobre", 11: "novembre", 4: "avril", 12: "décembre", 1: "janvier", 2: "février", 3: "mars", 6: "juin", 8: "août", 9: "septembre", 7: "juillet"}}, Days: ut.CalendarDayFormatNames{Abbreviated: ut.CalendarDayFormatNameValue{2: "mar.", 3: "mer.", 4: "jeu.", 5: "ven.", 6: "sam.", 0: "dim.", 1: "lun."}, Narrow: ut.CalendarDayFormatNameValue{5: "V", 6: "S", 0: "D", 1: "L", 2: "M", 3: "M", 4: "J"}, Short: ut.CalendarDayFormatNameValue{2: "ma", 3: "me", 4: "je", 5: "ve", 6: "sa", 0: "di", 1: "lu"}, Wide: ut.CalendarDayFormatNameValue{4: "jeudi", 5: "vendredi", 6: "samedi", 0: "dimanche", 1: "lundi", 2: "mardi", 3: "mercredi"}}, Periods: ut.CalendarPeriodFormatNames{Abbreviated: ut.CalendarPeriodFormatNameValue{"noon": "midi", "pm": "PM", "morning1": "mat.", "afternoon1": "ap.m.", "evening1": "soir", "night1": "nuit", "midnight": "min.", "am": "AM"}, Narrow: ut.CalendarPeriodFormatNameValue{"noon": "midi", "pm": "PM", "morning1": "mat.", "afternoon1": "ap.m.", "evening1": "soir", "night1": "nuit", "midnight": "min.", "am": "AM"}, Short: ut.CalendarPeriodFormatNameValue{}, Wide: ut.CalendarPeriodFormatNameValue{"morning1": "matin", "afternoon1": "après-midi", "evening1": "soir", "night1": "nuit", "midnight": "minuit", "am": "AM", "noon": "midi", "pm": "PM"}}, Eras: ut.Eras{AD: ut.CalendarEraFormatNames{Full: "", Abbrev: "", Narrow: ""}, BC: ut.CalendarEraFormatNames{Full: "", Abbrev: "", Narrow: ""}}}}
