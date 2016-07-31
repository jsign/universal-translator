package mr

import "github.com/go-playground/universal-translator"

var calendar = ut.Calendar{Formats: ut.CalendarFormats{DateEra: ut.DateEra{BC: ut.CalendarDateFormat{Full: "EEEE, d MMMM, y", Long: "d MMMM, y", Medium: "d MMM, y", Short: "d/M/yy"}, AD: ut.CalendarDateFormat{Full: "EEEE, d MMMM, y", Long: "d MMMM, y", Medium: "d MMM, y", Short: "d/M/yy"}}, Time: ut.CalendarDateFormat{Full: "h:mm:ss a zzzz", Long: "h:mm:ss a z", Medium: "h:mm:ss a", Short: "h:mm a"}, DateTime: ut.CalendarDateFormat{Full: "{1} रोजी {0}", Long: "{1} रोजी {0}", Medium: "{1}, {0}", Short: "{1}, {0}"}}, FormatNames: ut.CalendarFormatNames{Months: ut.CalendarMonthFormatNames{Abbreviated: ut.CalendarMonthFormatNameValue{6: "जून", 11: "नोव्हें", 1: "जाने", 4: "एप्रि", 5: "मे", 8: "ऑग", 9: "सप्टें", 10: "ऑक्टो", 12: "डिसें", 2: "फेब्रु", 3: "मार्च", 7: "जुलै"}, Narrow: ut.CalendarMonthFormatNameValue{2: "फे", 3: "मा", 4: "ए", 9: "स", 10: "ऑ", 11: "नो", 12: "डि", 1: "जा", 5: "मे", 6: "जू", 7: "जु", 8: "ऑ"}, Short: ut.CalendarMonthFormatNameValue(nil), Wide: ut.CalendarMonthFormatNameValue{9: "सप्टेंबर", 10: "ऑक्टोबर", 1: "जानेवारी", 2: "फेब्रुवारी", 3: "मार्च", 4: "एप्रिल", 5: "मे", 8: "ऑगस्ट", 11: "नोव्हेंबर", 6: "जून", 7: "जुलै", 12: "डिसेंबर"}}, Days: ut.CalendarDayFormatNames{Abbreviated: ut.CalendarDayFormatNameValue{0: "रवि", 1: "सोम", 2: "मंगळ", 3: "बुध", 4: "गुरु", 5: "शुक्र", 6: "शनि"}, Narrow: ut.CalendarDayFormatNameValue{0: "र", 1: "सो", 2: "मं", 3: "बु", 4: "गु", 5: "शु", 6: "श"}, Short: ut.CalendarDayFormatNameValue{1: "सो", 2: "मं", 3: "बु", 4: "गु", 5: "शु", 6: "श", 0: "र"}, Wide: ut.CalendarDayFormatNameValue{5: "शुक्रवार", 6: "शनिवार", 0: "रविवार", 1: "सोमवार", 2: "मंगळवार", 3: "बुधवार", 4: "गुरुवार"}}, Periods: ut.CalendarPeriodFormatNames{Abbreviated: ut.CalendarPeriodFormatNameValue{"pm": "म.उ.", "evening1": "संध्याकाळ", "night1": "रात्र", "am": "म.पू.", "noon": "मध्यान्ह", "morning2": "सकाळ", "afternoon1": "दुपार", "evening2": "सायंकाळ", "midnight": "मध्यरात्र", "morning1": "पहाट"}, Narrow: ut.CalendarPeriodFormatNameValue{"morning1": "प", "afternoon1": "दु", "evening1": "सं", "evening2": "सा", "am": "म.पू.", "noon": "म", "morning2": "स", "night1": "रा", "midnight": "म.रा.", "pm": "म.उ."}, Short: ut.CalendarPeriodFormatNameValue(nil), Wide: ut.CalendarPeriodFormatNameValue{"night1": "रात्र", "midnight": "मध्यरात्र", "am": "म.पू.", "pm": "म.उ.", "evening1": "संध्याकाळ", "evening2": "सायंकाळ", "noon": "मध्यान्ह", "morning1": "पहाट", "morning2": "सकाळ", "afternoon1": "दुपार"}}, Eras: ut.Eras{AD: ut.CalendarEraFormatNames{Full: "ईसवीसन", Abbrev: "इ. स.", Narrow: ""}, BC: ut.CalendarEraFormatNames{Full: "ईसवीसनपूर्व", Abbrev: "इ. स. पू.", Narrow: ""}}}}
