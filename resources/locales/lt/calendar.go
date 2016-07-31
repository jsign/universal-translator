package lt

import "github.com/go-playground/universal-translator"

var calendar = ut.Calendar{Formats: ut.CalendarFormats{DateEra: ut.DateEra{BC: ut.CalendarDateFormat{Full: "y 'm'. MMMM d 'd'., EEEE", Long: "y 'm'. MMMM d 'd'.", Medium: "y-MM-dd", Short: "y-MM-dd"}, AD: ut.CalendarDateFormat{Full: "y 'm'. MMMM d 'd'., EEEE", Long: "y 'm'. MMMM d 'd'.", Medium: "y-MM-dd", Short: "y-MM-dd"}}, Time: ut.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: ut.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"}}, FormatNames: ut.CalendarFormatNames{Months: ut.CalendarMonthFormatNames{Abbreviated: ut.CalendarMonthFormatNameValue{3: "kov.", 4: "bal.", 6: "birž.", 9: "rugs.", 10: "spal.", 11: "lapkr.", 1: "saus.", 2: "vas.", 12: "gruod.", 8: "rugp.", 5: "geg.", 7: "liep."}, Narrow: ut.CalendarMonthFormatNameValue{6: "B", 7: "L", 8: "R", 9: "R", 11: "L", 12: "G", 3: "K", 5: "G", 4: "B", 10: "S", 1: "S", 2: "V"}, Short: ut.CalendarMonthFormatNameValue(nil), Wide: ut.CalendarMonthFormatNameValue{11: "lapkritis", 3: "kovas", 5: "gegužė", 6: "birželis", 7: "liepa", 9: "rugsėjis", 12: "gruodis", 1: "sausis", 2: "vasaris", 4: "balandis", 8: "rugpjūtis", 10: "spalis"}}, Days: ut.CalendarDayFormatNames{Abbreviated: ut.CalendarDayFormatNameValue{4: "kt", 5: "pn", 6: "št", 0: "sk", 1: "pr", 2: "an", 3: "tr"}, Narrow: ut.CalendarDayFormatNameValue{0: "S", 1: "P", 2: "A", 3: "T", 4: "K", 5: "P", 6: "Š"}, Short: ut.CalendarDayFormatNameValue{1: "Pr", 2: "An", 3: "Tr", 4: "Kt", 5: "Pn", 6: "Št", 0: "Sk"}, Wide: ut.CalendarDayFormatNameValue{5: "penktadienis", 6: "šeštadienis", 0: "sekmadienis", 1: "pirmadienis", 2: "antradienis", 3: "trečiadienis", 4: "ketvirtadienis"}}, Periods: ut.CalendarPeriodFormatNames{Abbreviated: ut.CalendarPeriodFormatNameValue{"noon": "vidurdienis", "pm": "popiet", "morning1": "rytas", "afternoon1": "diena", "evening1": "vakaras", "night1": "naktis", "midnight": "vidurnaktis", "am": "priešpiet"}, Narrow: ut.CalendarPeriodFormatNameValue{"morning1": "rytas", "afternoon1": "diena", "evening1": "vakaras", "night1": "naktis", "midnight": "vidurnaktis", "am": "pr. p.", "noon": "vidurdienis", "pm": "pop."}, Short: ut.CalendarPeriodFormatNameValue(nil), Wide: ut.CalendarPeriodFormatNameValue{"midnight": "vidurnaktis", "am": "priešpiet", "noon": "vidurdienis", "pm": "popiet", "morning1": "rytas", "afternoon1": "diena", "evening1": "vakaras", "night1": "naktis"}}, Eras: ut.Eras{AD: ut.CalendarEraFormatNames{Full: "po Kristaus", Abbrev: "po Kr.", Narrow: "po Kr."}, BC: ut.CalendarEraFormatNames{Full: "prieš Kristų", Abbrev: "pr. Kr.", Narrow: "pr. Kr."}}}}