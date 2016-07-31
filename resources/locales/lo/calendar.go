package lo

import "github.com/go-playground/universal-translator"

var calendar = ut.Calendar{Formats: ut.CalendarFormats{DateEra: ut.DateEra{BC: ut.CalendarDateFormat{Full: "EEEE ທີ d MMMM G y", Long: "d MMMM y", Medium: "d MMM y", Short: "d/M/y"}, AD: ut.CalendarDateFormat{Full: "EEEE ທີ d MMMM G y", Long: "d MMMM y", Medium: "d MMM y", Short: "d/M/y"}}, Time: ut.CalendarDateFormat{Full: "H ໂມງ m ນາທີ ss ວິນາທີ zzzz", Long: "H ໂມງ m ນາທີ ss ວິນາທີ z", Medium: "H:mm:ss", Short: "H:mm"}, DateTime: ut.CalendarDateFormat{Full: "{1}, {0}", Long: "{1}, {0}", Medium: "{1}, {0}", Short: "{1}, {0}"}}, FormatNames: ut.CalendarFormatNames{Months: ut.CalendarMonthFormatNames{Abbreviated: ut.CalendarMonthFormatNameValue{3: "ມ.ນ.", 5: "ພ.ພ.", 8: "ສ.ຫ.", 9: "ກ.ຍ.", 10: "ຕ.ລ.", 12: "ທ.ວ.", 1: "ມ.ກ.", 4: "ມ.ສ.", 6: "ມິ.ຖ.", 7: "ກ.ລ.", 11: "ພ.ຈ.", 2: "ກ.ພ."}, Narrow: ut.CalendarMonthFormatNameValue{6: "6", 7: "7", 9: "9", 10: "10", 12: "12", 1: "1", 3: "3", 5: "5", 8: "8", 11: "11", 2: "2", 4: "4"}, Short: ut.CalendarMonthFormatNameValue(nil), Wide: ut.CalendarMonthFormatNameValue{9: "ກັນຍາ", 10: "ຕຸລາ", 1: "ມັງກອນ", 3: "ມີນາ", 5: "ພຶດສະພາ", 6: "ມິຖຸນາ", 11: "ພະຈິກ", 12: "ທັນວາ", 2: "ກຸມພາ", 4: "ເມສາ", 7: "ກໍລະກົດ", 8: "ສິງຫາ"}}, Days: ut.CalendarDayFormatNames{Abbreviated: ut.CalendarDayFormatNameValue{0: "ອາທິດ", 1: "ຈັນ", 2: "ອັງຄານ", 3: "ພຸດ", 4: "ພະຫັດ", 5: "ສຸກ", 6: "ເສົາ"}, Narrow: ut.CalendarDayFormatNameValue{0: "ອ", 1: "ຈ", 2: "ອ", 3: "ພ", 4: "ພຫ", 5: "ສຸ", 6: "ສ"}, Short: ut.CalendarDayFormatNameValue{0: "ອາ.", 1: "ຈ.", 2: "ອ.", 3: "ພ.", 4: "ພຫ.", 5: "ສຸ.", 6: "ສ."}, Wide: ut.CalendarDayFormatNameValue{0: "ວັນອາທິດ", 1: "ວັນຈັນ", 2: "ວັນອັງຄານ", 3: "ວັນພຸດ", 4: "ວັນພະຫັດ", 5: "ວັນສຸກ", 6: "ວັນເສົາ"}}, Periods: ut.CalendarPeriodFormatNames{Abbreviated: ut.CalendarPeriodFormatNameValue{"midnight": "ທ່ຽງ\u200bຄືນ", "am": "ກ່ອນທ່ຽງ", "noon": "ຕອນທ່ຽງ", "pm": "ຫຼັງທ່ຽງ", "morning1": "\u200bເຊົ້າ", "afternoon1": "ສວຍ", "evening1": "ແລງ", "night1": "\u200bກາງ\u200bຄືນ"}, Narrow: ut.CalendarPeriodFormatNameValue{"night1": "\u200bກາງ\u200bຄືນ", "midnight": "ທ່ຽງ\u200bຄືນ", "am": "ກທ", "noon": "ຕອນທ່ຽງ", "pm": "ຫຼທ", "morning1": "\u200bເຊົ້າ", "afternoon1": "ສວຍ", "evening1": "ແລງ"}, Short: ut.CalendarPeriodFormatNameValue(nil), Wide: ut.CalendarPeriodFormatNameValue{"am": "ກ່ອນທ່ຽງ", "noon": "ຕອນທ່ຽງ", "pm": "ຫຼັງທ່ຽງ", "morning1": "\u200bເຊົ້າ", "afternoon1": "ສວຍ", "evening1": "ແລງ", "night1": "\u200bກາງ\u200bຄືນ", "midnight": "ທ່ຽງ\u200bຄືນ"}}, Eras: ut.Eras{AD: ut.CalendarEraFormatNames{Full: "ຄຣິດສັກກະລາດ", Abbrev: "ຄ.ສ.", Narrow: ""}, BC: ut.CalendarEraFormatNames{Full: "ກ່ອນຄຣິດສັກກະລາດ", Abbrev: "ກ່ອນ ຄ.ສ.", Narrow: ""}}}}
