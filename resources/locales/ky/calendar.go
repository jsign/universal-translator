package ky

import "github.com/go-playground/universal-translator"

var calendar = ut.Calendar{Formats: ut.CalendarFormats{DateEra: ut.DateEra{BC: ut.CalendarDateFormat{Full: "EEEE, d-MMMM, y-'ж'.", Long: "y MMMM d", Medium: "y MMM d", Short: "dd.MM.yy"}, AD: ut.CalendarDateFormat{Full: "EEEE, d-MMMM, y-'ж'.", Long: "y MMMM d", Medium: "y MMM d", Short: "dd.MM.yy"}}, Time: ut.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: ut.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"}}, FormatNames: ut.CalendarFormatNames{Months: ut.CalendarMonthFormatNames{Abbreviated: ut.CalendarMonthFormatNameValue{11: "Ноя", 1: "Янв", 4: "Апр", 6: "Июн", 7: "Июл", 9: "Сен", 10: "Окт", 2: "Фев", 3: "Мар", 5: "Май", 8: "Авг", 12: "Дек"}, Narrow: ut.CalendarMonthFormatNameValue{11: "Н", 3: "М", 5: "М", 7: "И", 9: "С", 10: "О", 12: "Д", 1: "Я", 2: "Ф", 4: "А", 6: "И", 8: "А"}, Short: ut.CalendarMonthFormatNameValue(nil), Wide: ut.CalendarMonthFormatNameValue{5: "Май", 6: "Июнь", 7: "Июль", 10: "Октябрь", 12: "Декабрь", 1: "Январь", 2: "Февраль", 8: "Август", 9: "Сентябрь", 11: "Ноябрь", 3: "Март", 4: "Апрель"}}, Days: ut.CalendarDayFormatNames{Abbreviated: ut.CalendarDayFormatNameValue{0: "жек.", 1: "дүй.", 2: "шейш.", 3: "шарш.", 4: "бейш.", 5: "жума", 6: "ишм."}, Narrow: ut.CalendarDayFormatNameValue{1: "Д", 2: "Ш", 3: "Ш", 4: "Б", 5: "Ж", 6: "И", 0: "Ж"}, Short: ut.CalendarDayFormatNameValue{2: "шш.", 3: "шр.", 4: "бш.", 5: "жм.", 6: "иш.", 0: "жк", 1: "дш."}, Wide: ut.CalendarDayFormatNameValue{3: "шаршемби", 4: "бейшемби", 5: "жума", 6: "ишемби", 0: "жекшемби", 1: "дүйшөмбү", 2: "шейшемби"}}, Periods: ut.CalendarPeriodFormatNames{Abbreviated: ut.CalendarPeriodFormatNameValue{"evening1": "кечкурун", "night1": "түн", "midnight": "түн ортосу", "am": "тң", "noon": "чак түш", "pm": "тк", "morning1": "эртең менен", "afternoon1": "түштөн кийин"}, Narrow: ut.CalendarPeriodFormatNameValue{"morning1": "эртең менен", "afternoon1": "түштөн кийин", "evening1": "кечкурун", "night1": "түн", "midnight": "түн ортосу", "am": "тң", "noon": "чак түш", "pm": "тк"}, Short: ut.CalendarPeriodFormatNameValue(nil), Wide: ut.CalendarPeriodFormatNameValue{"night1": "түн", "midnight": "түн ортосу", "am": "таңкы", "noon": "чак түш", "pm": "түштөн кийинки", "morning1": "эртең менен", "afternoon1": "түштөн кийин", "evening1": "кечкурун"}}, Eras: ut.Eras{AD: ut.CalendarEraFormatNames{Full: "", Abbrev: "", Narrow: ""}, BC: ut.CalendarEraFormatNames{Full: "биздин заманга чейин", Abbrev: "б.з.ч.", Narrow: "б.з.ч."}}}}
