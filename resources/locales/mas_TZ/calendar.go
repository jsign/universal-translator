package mas_TZ

import "github.com/go-playground/universal-translator"

var calendar = ut.Calendar{Formats: ut.CalendarFormats{Date: ut.CalendarDateFormat{Full: "EEEE, d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "dd/MM/y"}, Time: ut.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: ut.CalendarDateFormat{Full: "", Long: "", Medium: "", Short: ""}}, FormatNames: ut.CalendarFormatNames{Months: ut.CalendarMonthFormatNames{Abbreviated: ut.CalendarMonthFormatNameValue{12: "Ntʉ́", 1: "Dal", 3: "Ɔɛn", 7: "Sás", 8: "Bɔ́r", 5: "Lép", 9: "Kús", 10: "Gís", 11: "Shʉ́", 6: "Rok", 2: "Ará", 4: "Doy"}, Narrow: ut.CalendarMonthFormatNameValue{}, Short: ut.CalendarMonthFormatNameValue{}, Wide: ut.CalendarMonthFormatNameValue{6: "Kújúɔrɔk", 3: "Ɔɛnɨ́ɔɨŋɔk", 12: "Ntʉ́ŋʉ́s", 10: "Olgísan", 11: "Pʉshʉ́ka", 1: "Oladalʉ́", 7: "Mórusásin", 8: "Ɔlɔ́ɨ́bɔ́rárɛ", 9: "Kúshîn", 2: "Arát", 4: "Olodoyíóríê inkókúâ", 5: "Oloilépūnyīē inkókúâ"}}, Days: ut.CalendarDayFormatNames{Abbreviated: ut.CalendarDayFormatNameValue{5: "Iju", 6: "Jmo", 0: "Jpi", 1: "Jtt", 2: "Jnn", 3: "Jtn", 4: "Alh"}, Narrow: ut.CalendarDayFormatNameValue{3: "5", 4: "6", 5: "7", 6: "1", 0: "2", 1: "3", 2: "4"}, Short: ut.CalendarDayFormatNameValue{}, Wide: ut.CalendarDayFormatNameValue{4: "Alaámisi", 5: "Jumáa", 6: "Jumamósi", 0: "Jumapílí", 1: "Jumatátu", 2: "Jumane", 3: "Jumatánɔ"}}, Periods: ut.CalendarPeriodFormatNames{Abbreviated: ut.CalendarPeriodFormatNameValue{}, Narrow: ut.CalendarPeriodFormatNameValue{}, Short: ut.CalendarPeriodFormatNameValue{}, Wide: ut.CalendarPeriodFormatNameValue{"am": "Ɛnkakɛnyá", "pm": "Ɛndámâ"}}}}