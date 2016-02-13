package yo_BJ

import "github.com/go-playground/universal-translator"

var calendar = ut.Calendar{Formats: ut.CalendarFormats{Date: ut.CalendarDateFormat{Full: "EEEE, d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "dd/MM/y"}, Time: ut.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: ut.CalendarDateFormat{Full: "", Long: "", Medium: "", Short: ""}}, FormatNames: ut.CalendarFormatNames{Months: ut.CalendarMonthFormatNames{Abbreviated: ut.CalendarMonthFormatNameValue{6: "Òkúdu", 7: "Agɛmɔ", 8: "Ògún", 10: "Ɔ̀wàrà", 11: "Bélú", 1: "Shɛ́rɛ́", 3: "Ɛrɛ̀nà", 4: "Ìgbé", 5: "Ɛ̀bibi", 9: "Owewe", 12: "Ɔ̀pɛ̀", 2: "Èrèlè"}, Narrow: ut.CalendarMonthFormatNameValue{}, Short: ut.CalendarMonthFormatNameValue{}, Wide: ut.CalendarMonthFormatNameValue{10: "Oshù Ɔ̀wàrà", 1: "Oshù Shɛ́rɛ́", 4: "Oshù Ìgbé", 7: "Oshù Agɛmɔ", 8: "Oshù Ògún", 9: "Oshù Owewe", 12: "Oshù Ɔ̀pɛ̀", 2: "Oshù Èrèlè", 3: "Oshù Ɛrɛ̀nà", 5: "Oshù Ɛ̀bibi", 6: "Oshù Òkúdu", 11: "Oshù Bélú"}}, Days: ut.CalendarDayFormatNames{Abbreviated: ut.CalendarDayFormatNameValue{1: "Ajé", 2: "Ìsɛ́gun", 3: "Ɔjɔ́rú", 4: "Ɔjɔ́bɔ", 5: "Ɛtì", 6: "Àbámɛ́ta", 0: "Àìkú"}, Narrow: ut.CalendarDayFormatNameValue{}, Short: ut.CalendarDayFormatNameValue{}, Wide: ut.CalendarDayFormatNameValue{2: "Ɔjɔ́ Ìsɛ́gun", 3: "Ɔjɔ́rú", 4: "Ɔjɔ́bɔ", 5: "Ɔjɔ́ Ɛtì", 6: "Ɔjɔ́ Àbámɛ́ta", 0: "Ɔjɔ́ Àìkú", 1: "Ɔjɔ́ Ajé"}}, Periods: ut.CalendarPeriodFormatNames{Abbreviated: ut.CalendarPeriodFormatNameValue{}, Narrow: ut.CalendarPeriodFormatNameValue{}, Short: ut.CalendarPeriodFormatNameValue{}, Wide: ut.CalendarPeriodFormatNameValue{"am": "Àárɔ̀", "pm": "Ɔ̀sán"}}}}