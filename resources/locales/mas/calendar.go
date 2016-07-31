package mas

import "github.com/go-playground/universal-translator"

var calendar = ut.Calendar{Formats: ut.CalendarFormats{DateEra: ut.DateEra{BC: ut.CalendarDateFormat{Full: "EEEE, d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "dd/MM/y"}, AD: ut.CalendarDateFormat{Full: "EEEE, d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "dd/MM/y"}}, Time: ut.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: ut.CalendarDateFormat{Full: "", Long: "", Medium: "", Short: ""}}, FormatNames: ut.CalendarFormatNames{Months: ut.CalendarMonthFormatNames{Abbreviated: ut.CalendarMonthFormatNameValue{4: "Doy", 5: "Lép", 10: "Gís", 12: "Ntʉ́", 1: "Dal", 2: "Ará", 7: "Sás", 8: "Bɔ́r", 9: "Kús", 11: "Shʉ́", 3: "Ɔɛn", 6: "Rok"}, Narrow: ut.CalendarMonthFormatNameValue(nil), Short: ut.CalendarMonthFormatNameValue(nil), Wide: ut.CalendarMonthFormatNameValue{9: "Kúshîn", 10: "Olgísan", 11: "Pʉshʉ́ka", 5: "Oloilépūnyīē inkókúâ", 6: "Kújúɔrɔk", 3: "Ɔɛnɨ́ɔɨŋɔk", 4: "Olodoyíóríê inkókúâ", 7: "Mórusásin", 8: "Ɔlɔ́ɨ́bɔ́rárɛ", 12: "Ntʉ́ŋʉ́s", 1: "Oladalʉ́", 2: "Arát"}}, Days: ut.CalendarDayFormatNames{Abbreviated: ut.CalendarDayFormatNameValue{5: "Iju", 6: "Jmo", 0: "Jpi", 1: "Jtt", 2: "Jnn", 3: "Jtn", 4: "Alh"}, Narrow: ut.CalendarDayFormatNameValue{2: "4", 3: "5", 4: "6", 5: "7", 6: "1", 0: "2", 1: "3"}, Short: ut.CalendarDayFormatNameValue(nil), Wide: ut.CalendarDayFormatNameValue{3: "Jumatánɔ", 4: "Alaámisi", 5: "Jumáa", 6: "Jumamósi", 0: "Jumapílí", 1: "Jumatátu", 2: "Jumane"}}, Periods: ut.CalendarPeriodFormatNames{Abbreviated: ut.CalendarPeriodFormatNameValue{"pm": "Ɛndámâ", "am": "Ɛnkakɛnyá"}, Narrow: ut.CalendarPeriodFormatNameValue(nil), Short: ut.CalendarPeriodFormatNameValue(nil), Wide: ut.CalendarPeriodFormatNameValue{"am": "Ɛnkakɛnyá", "pm": "Ɛndámâ"}}, Eras: ut.Eras{AD: ut.CalendarEraFormatNames{Full: "", Abbrev: "", Narrow: ""}, BC: ut.CalendarEraFormatNames{Full: "Meínō Yɛ́sʉ", Abbrev: "MY", Narrow: ""}}}}