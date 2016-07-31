package teo

import "github.com/go-playground/universal-translator"

var calendar = ut.Calendar{Formats: ut.CalendarFormats{DateEra: ut.DateEra{BC: ut.CalendarDateFormat{Full: "EEEE, d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "dd/MM/y"}, AD: ut.CalendarDateFormat{Full: "EEEE, d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "dd/MM/y"}}, Time: ut.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: ut.CalendarDateFormat{Full: "", Long: "", Medium: "", Short: ""}}, FormatNames: ut.CalendarFormatNames{Months: ut.CalendarMonthFormatNames{Abbreviated: ut.CalendarMonthFormatNameValue{4: "Dun", 5: "Mar", 6: "Mod", 7: "Jol", 8: "Ped", 11: "Lab", 2: "Muk", 3: "Kwa", 9: "Sok", 10: "Tib", 12: "Poo", 1: "Rar"}, Narrow: ut.CalendarMonthFormatNameValue{7: "J", 9: "S", 11: "L", 12: "P", 1: "R", 3: "K", 6: "M", 8: "P", 10: "T", 2: "M", 4: "D", 5: "M"}, Short: ut.CalendarMonthFormatNameValue(nil), Wide: ut.CalendarMonthFormatNameValue{1: "Orara", 4: "Odung’el", 5: "Omaruk", 6: "Omodok’king’ol", 7: "Ojola", 12: "Opoo", 2: "Omuk", 3: "Okwamg’", 8: "Opedel", 9: "Osokosokoma", 10: "Otibar", 11: "Olabor"}}, Days: ut.CalendarDayFormatNames{Abbreviated: ut.CalendarDayFormatNameValue{0: "Jum", 1: "Bar", 2: "Aar", 3: "Uni", 4: "Ung", 5: "Kan", 6: "Sab"}, Narrow: ut.CalendarDayFormatNameValue{0: "J", 1: "B", 2: "A", 3: "U", 4: "U", 5: "K", 6: "S"}, Short: ut.CalendarDayFormatNameValue(nil), Wide: ut.CalendarDayFormatNameValue{0: "Nakaejuma", 1: "Nakaebarasa", 2: "Nakaare", 3: "Nakauni", 4: "Nakaung’on", 5: "Nakakany", 6: "Nakasabiti"}}, Periods: ut.CalendarPeriodFormatNames{Abbreviated: ut.CalendarPeriodFormatNameValue{"am": "Taparachu", "pm": "Ebongi"}, Narrow: ut.CalendarPeriodFormatNameValue(nil), Short: ut.CalendarPeriodFormatNameValue(nil), Wide: ut.CalendarPeriodFormatNameValue{"am": "Taparachu", "pm": "Ebongi"}}, Eras: ut.Eras{AD: ut.CalendarEraFormatNames{Full: "", Abbrev: "", Narrow: ""}, BC: ut.CalendarEraFormatNames{Full: "Kabla ya Christo", Abbrev: "KK", Narrow: ""}}}}
