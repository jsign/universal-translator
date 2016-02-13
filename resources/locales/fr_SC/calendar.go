package fr_SC

import "github.com/go-playground/universal-translator"

var calendar = ut.Calendar{Formats: ut.CalendarFormats{Date: ut.CalendarDateFormat{Full: "EEEE d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "dd/MM/y"}, Time: ut.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: ut.CalendarDateFormat{Full: "{1} 'à' {0}", Long: "{1} 'à' {0}", Medium: "{1} 'à' {0}", Short: "{1} {0}"}}, FormatNames: ut.CalendarFormatNames{Months: ut.CalendarMonthFormatNames{Abbreviated: ut.CalendarMonthFormatNameValue{10: "oct.", 1: "janv.", 6: "juin", 12: "déc.", 2: "févr.", 3: "mars", 5: "mai", 8: "août", 4: "avr.", 7: "juil.", 9: "sept.", 11: "nov."}, Narrow: ut.CalendarMonthFormatNameValue{1: "J", 2: "F", 5: "M", 4: "A", 8: "A", 10: "O", 11: "N", 7: "J", 9: "S", 3: "M", 6: "J", 12: "D"}, Short: ut.CalendarMonthFormatNameValue{}, Wide: ut.CalendarMonthFormatNameValue{5: "mai", 7: "juillet", 8: "août", 1: "janvier", 6: "juin", 11: "novembre", 2: "février", 3: "mars", 9: "septembre", 4: "avril", 10: "octobre", 12: "décembre"}}, Days: ut.CalendarDayFormatNames{Abbreviated: ut.CalendarDayFormatNameValue{4: "jeu.", 5: "ven.", 6: "sam.", 0: "dim.", 1: "lun.", 2: "mar.", 3: "mer."}, Narrow: ut.CalendarDayFormatNameValue{6: "S", 0: "D", 1: "L", 2: "M", 3: "M", 4: "J", 5: "V"}, Short: ut.CalendarDayFormatNameValue{0: "di", 1: "lu", 2: "ma", 3: "me", 4: "je", 5: "ve", 6: "sa"}, Wide: ut.CalendarDayFormatNameValue{1: "lundi", 2: "mardi", 3: "mercredi", 4: "jeudi", 5: "vendredi", 6: "samedi", 0: "dimanche"}}, Periods: ut.CalendarPeriodFormatNames{Abbreviated: ut.CalendarPeriodFormatNameValue{"morning1": "mat.", "afternoon1": "ap.m.", "evening1": "soir", "night1": "nuit", "midnight": "min.", "am": "AM", "noon": "midi", "pm": "PM"}, Narrow: ut.CalendarPeriodFormatNameValue{"evening1": "soir", "night1": "nuit", "midnight": "min.", "am": "AM", "noon": "midi", "pm": "PM", "morning1": "mat.", "afternoon1": "ap.m."}, Short: ut.CalendarPeriodFormatNameValue{}, Wide: ut.CalendarPeriodFormatNameValue{"night1": "nuit", "midnight": "minuit", "am": "AM", "noon": "midi", "pm": "PM", "morning1": "matin", "afternoon1": "après-midi", "evening1": "soir"}}}}