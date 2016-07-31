package nb

import "github.com/go-playground/universal-translator"

var calendar = ut.Calendar{Formats: ut.CalendarFormats{DateEra: ut.DateEra{BC: ut.CalendarDateFormat{Full: "EEEE d. MMMM y", Long: "d. MMMM y", Medium: "d. MMM y", Short: "dd.MM.y"}, AD: ut.CalendarDateFormat{Full: "EEEE d. MMMM y", Long: "d. MMMM y", Medium: "d. MMM y", Short: "dd.MM.y"}}, Time: ut.CalendarDateFormat{Full: "HH.mm.ss zzzz", Long: "HH.mm.ss z", Medium: "HH.mm.ss", Short: "HH.mm"}, DateTime: ut.CalendarDateFormat{Full: "{1} {0}", Long: "{1} 'kl'. {0}", Medium: "{1}, {0}", Short: "{1}, {0}"}}, FormatNames: ut.CalendarFormatNames{Months: ut.CalendarMonthFormatNames{Abbreviated: ut.CalendarMonthFormatNameValue{7: "jul", 8: "aug", 10: "okt", 12: "des", 2: "feb", 3: "mar", 4: "apr", 6: "jun", 1: "jan", 5: "mai", 9: "sep", 11: "nov"}, Narrow: ut.CalendarMonthFormatNameValue{1: "J", 9: "S", 10: "O", 11: "N", 12: "D", 2: "F", 3: "M", 4: "A", 5: "M", 6: "J", 7: "J", 8: "A"}, Short: ut.CalendarMonthFormatNameValue(nil), Wide: ut.CalendarMonthFormatNameValue{2: "februar", 5: "mai", 6: "juni", 7: "juli", 12: "desember", 1: "januar", 3: "mars", 4: "april", 8: "august", 9: "september", 10: "oktober", 11: "november"}}, Days: ut.CalendarDayFormatNames{Abbreviated: ut.CalendarDayFormatNameValue{6: "lør.", 0: "søn.", 1: "man.", 2: "tir.", 3: "ons.", 4: "tor.", 5: "fre."}, Narrow: ut.CalendarDayFormatNameValue{3: "O", 4: "T", 5: "F", 6: "L", 0: "S", 1: "M", 2: "T"}, Short: ut.CalendarDayFormatNameValue{1: "ma.", 2: "ti.", 3: "on.", 4: "to.", 5: "fr.", 6: "lø.", 0: "sø."}, Wide: ut.CalendarDayFormatNameValue{5: "fredag", 6: "lørdag", 0: "søndag", 1: "mandag", 2: "tirsdag", 3: "onsdag", 4: "torsdag"}}, Periods: ut.CalendarPeriodFormatNames{Abbreviated: ut.CalendarPeriodFormatNameValue{"pm": "p.m.", "morning1": "morg.", "morning2": "form.", "afternoon1": "etterm.", "evening1": "kveld", "night1": "natt", "midnight": "midn.", "am": "a.m."}, Narrow: ut.CalendarPeriodFormatNameValue{"morning1": "mg.", "morning2": "fm.", "afternoon1": "em.", "evening1": "kv.", "night1": "nt.", "midnight": "mn.", "am": "a.m.", "pm": "p.m."}, Short: ut.CalendarPeriodFormatNameValue(nil), Wide: ut.CalendarPeriodFormatNameValue{"evening1": "kveld", "night1": "natt", "midnight": "midnatt", "am": "a.m.", "pm": "p.m.", "morning1": "morgen", "morning2": "formiddag", "afternoon1": "ettermiddag"}}, Eras: ut.Eras{AD: ut.CalendarEraFormatNames{Full: "etter Kristus", Abbrev: "e.Kr.", Narrow: "e.Kr."}, BC: ut.CalendarEraFormatNames{Full: "før Kristus", Abbrev: "f.Kr.", Narrow: "f.Kr."}}}}
