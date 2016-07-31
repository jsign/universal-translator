package ff_GN

import "github.com/go-playground/universal-translator"

var calendar = ut.Calendar{Formats: ut.CalendarFormats{DateEra: ut.DateEra{BC: ut.CalendarDateFormat{Full: "EEEE d MMMM y", Long: "d MMMM y", Medium: "d MMM, y", Short: "d/M/y"}, AD: ut.CalendarDateFormat{Full: "EEEE d MMMM y", Long: "d MMMM y", Medium: "d MMM, y", Short: "d/M/y"}}, Time: ut.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: ut.CalendarDateFormat{Full: "", Long: "", Medium: "", Short: ""}}, FormatNames: ut.CalendarFormatNames{Months: ut.CalendarMonthFormatNames{Abbreviated: ut.CalendarMonthFormatNameValue{1: "sii", 7: "mor", 3: "mbo", 2: "col", 5: "duu", 9: "slt", 12: "bow", 6: "kor", 8: "juk", 4: "see", 10: "yar", 11: "jol"}, Narrow: ut.CalendarMonthFormatNameValue{3: "m", 12: "b", 9: "s", 10: "y", 6: "k", 7: "m", 2: "c", 4: "s", 11: "j", 8: "j", 1: "s", 5: "d"}, Short: ut.CalendarMonthFormatNameValue{}, Wide: ut.CalendarMonthFormatNameValue{11: "jolal", 2: "colte", 8: "juko", 12: "bowte", 1: "siilo", 4: "seeɗto", 3: "mbooy", 5: "duujal", 9: "siilto", 6: "korse", 7: "morso", 10: "yarkomaa"}}, Days: ut.CalendarDayFormatNames{Abbreviated: ut.CalendarDayFormatNameValue{2: "maw", 3: "nje", 4: "naa", 5: "mwd", 6: "hbi", 0: "dew", 1: "aaɓ"}, Narrow: ut.CalendarDayFormatNameValue{0: "d", 1: "a", 2: "m", 3: "n", 4: "n", 5: "m", 6: "h"}, Short: ut.CalendarDayFormatNameValue{}, Wide: ut.CalendarDayFormatNameValue{2: "mawbaare", 3: "njeslaare", 4: "naasaande", 5: "mawnde", 6: "hoore-biir", 0: "dewo", 1: "aaɓnde"}}, Periods: ut.CalendarPeriodFormatNames{Abbreviated: ut.CalendarPeriodFormatNameValue{"am": "subaka", "pm": "kikiiɗe"}, Narrow: ut.CalendarPeriodFormatNameValue{}, Short: ut.CalendarPeriodFormatNameValue{}, Wide: ut.CalendarPeriodFormatNameValue{"am": "subaka", "pm": "kikiiɗe"}}, Eras: ut.Eras{AD: ut.CalendarEraFormatNames{Full: "", Abbrev: "", Narrow: ""}, BC: ut.CalendarEraFormatNames{Full: "", Abbrev: "", Narrow: ""}}}}