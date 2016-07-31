package to

import "github.com/go-playground/universal-translator"

var calendar = ut.Calendar{Formats: ut.CalendarFormats{DateEra: ut.DateEra{BC: ut.CalendarDateFormat{Full: "EEEE d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "d/M/yy"}, AD: ut.CalendarDateFormat{Full: "EEEE d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "d/M/yy"}}, Time: ut.CalendarDateFormat{Full: "h:mm:ss a zzzz", Long: "h:mm:ss a z", Medium: "h:mm:ss a", Short: "h:mm a"}, DateTime: ut.CalendarDateFormat{Full: "{1}, {0}", Long: "{1}, {0}", Medium: "{1}, {0}", Short: "{1} {0}"}}, FormatNames: ut.CalendarFormatNames{Months: ut.CalendarMonthFormatNames{Abbreviated: ut.CalendarMonthFormatNameValue{2: "Fēp", 5: "Mē", 9: "Sep", 8: "ʻAok", 10: "ʻOka", 11: "Nōv", 1: "Sān", 3: "Maʻa", 4: "ʻEpe", 6: "Sun", 7: "Siu", 12: "Tīs"}, Narrow: ut.CalendarMonthFormatNameValue{3: "M", 5: "M", 7: "S", 9: "S", 10: "O", 11: "N", 1: "S", 2: "F", 12: "T", 8: "A", 4: "E", 6: "S"}, Short: ut.CalendarMonthFormatNameValue(nil), Wide: ut.CalendarMonthFormatNameValue{7: "Siulai", 8: "ʻAokosi", 9: "Sepitema", 10: "ʻOkatopa", 1: "Sānuali", 2: "Fēpueli", 3: "Maʻasi", 11: "Nōvema", 12: "Tīsema", 4: "ʻEpeleli", 5: "Mē", 6: "Sune"}}, Days: ut.CalendarDayFormatNames{Abbreviated: ut.CalendarDayFormatNameValue{1: "Mōn", 2: "Tūs", 3: "Pul", 4: "Tuʻa", 5: "Fal", 6: "Tok", 0: "Sāp"}, Narrow: ut.CalendarDayFormatNameValue{5: "F", 6: "T", 0: "S", 1: "M", 2: "T", 3: "P", 4: "T"}, Short: ut.CalendarDayFormatNameValue{2: "Tūs", 3: "Pul", 4: "Tuʻa", 5: "Fal", 6: "Tok", 0: "Sāp", 1: "Mōn"}, Wide: ut.CalendarDayFormatNameValue{4: "Tuʻapulelulu", 5: "Falaite", 6: "Tokonaki", 0: "Sāpate", 1: "Mōnite", 2: "Tūsite", 3: "Pulelulu"}}, Periods: ut.CalendarPeriodFormatNames{Abbreviated: ut.CalendarPeriodFormatNameValue{"am": "AM", "pm": "PM"}, Narrow: ut.CalendarPeriodFormatNameValue{"am": "AM", "pm": "PM"}, Short: ut.CalendarPeriodFormatNameValue(nil), Wide: ut.CalendarPeriodFormatNameValue{"am": "AM", "pm": "PM"}}, Eras: ut.Eras{AD: ut.CalendarEraFormatNames{Full: "taʻu ʻo Sīsū", Abbrev: "", Narrow: ""}, BC: ut.CalendarEraFormatNames{Full: "ki muʻa", Abbrev: "KM", Narrow: ""}}}}