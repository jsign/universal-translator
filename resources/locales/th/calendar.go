package th

import "github.com/go-playground/universal-translator"

var calendar = ut.Calendar{Formats: ut.CalendarFormats{DateEra: ut.DateEra{BC: ut.CalendarDateFormat{Full: "EEEEที่ d MMMM y GGGG", Long: "d MMMM y GG", Medium: "d MMM y", Short: "d/M/yy"}, AD: ut.CalendarDateFormat{Full: "EEEEที่ d MMMM GGGG y", Long: "d MMMM GG y", Medium: "d MMM y", Short: "d/M/yy"}}, Time: ut.CalendarDateFormat{Full: "H นาฬิกา mm นาที ss วินาที zzzz", Long: "H นาฬิกา mm นาที ss วินาที z", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: ut.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"}}, FormatNames: ut.CalendarFormatNames{Months: ut.CalendarMonthFormatNames{Abbreviated: ut.CalendarMonthFormatNameValue{2: "ก.พ.", 7: "ก.ค.", 8: "ส.ค.", 9: "ก.ย.", 1: "ม.ค.", 3: "มี.ค.", 4: "เม.ย.", 5: "พ.ค.", 6: "มิ.ย.", 10: "ต.ค.", 11: "พ.ย.", 12: "ธ.ค."}, Narrow: ut.CalendarMonthFormatNameValue{8: "ส.ค.", 11: "พ.ย.", 2: "ก.พ.", 4: "เม.ย.", 5: "พ.ค.", 6: "มิ.ย.", 7: "ก.ค.", 1: "ม.ค.", 3: "มี.ค.", 9: "ก.ย.", 10: "ต.ค.", 12: "ธ.ค."}, Short: ut.CalendarMonthFormatNameValue(nil), Wide: ut.CalendarMonthFormatNameValue{8: "สิงหาคม", 11: "พฤศจิกายน", 12: "ธันวาคม", 3: "มีนาคม", 4: "เมษายน", 7: "กรกฎาคม", 6: "มิถุนายน", 9: "กันยายน", 10: "ตุลาคม", 1: "มกราคม", 2: "กุมภาพันธ์", 5: "พฤษภาคม"}}, Days: ut.CalendarDayFormatNames{Abbreviated: ut.CalendarDayFormatNameValue{3: "พ.", 4: "พฤ.", 5: "ศ.", 6: "ส.", 0: "อา.", 1: "จ.", 2: "อ."}, Narrow: ut.CalendarDayFormatNameValue{6: "ส", 0: "อา", 1: "จ", 2: "อ", 3: "พ", 4: "พฤ", 5: "ศ"}, Short: ut.CalendarDayFormatNameValue{3: "พ.", 4: "พฤ.", 5: "ศ.", 6: "ส.", 0: "อา.", 1: "จ.", 2: "อ."}, Wide: ut.CalendarDayFormatNameValue{6: "วันเสาร์", 0: "วันอาทิตย์", 1: "วันจันทร์", 2: "วันอังคาร", 3: "วันพุธ", 4: "วันพฤหัสบดี", 5: "วันศุกร์"}}, Periods: ut.CalendarPeriodFormatNames{Abbreviated: ut.CalendarPeriodFormatNameValue{"am": "ก่อนเที่ยง", "pm": "หลังเที่ยง"}, Narrow: ut.CalendarPeriodFormatNameValue{"evening1": "เย็น", "midnight": "เที่ยงคืน", "noon": "เที่ยง", "pm": "หลังเที่ยง", "morning1": "เช้า", "night1": "กลางคืน", "am": "ก่อนเที่ยง", "afternoon1": "ช่วงเที่ยง", "afternoon2": "บ่าย", "evening2": "ค่ำ"}, Short: ut.CalendarPeriodFormatNameValue(nil), Wide: ut.CalendarPeriodFormatNameValue{"midnight": "เที่ยงคืน", "pm": "หลังเที่ยง", "afternoon1": "ช่วงเที่ยง", "afternoon2": "บ่าย", "evening1": "เย็น", "am": "ก่อนเที่ยง", "noon": "เที่ยง", "morning1": "เช้า", "evening2": "ค่ำ", "night1": "กลางคืน"}}, Eras: ut.Eras{AD: ut.CalendarEraFormatNames{Full: "คริสต์ศักราช", Abbrev: "ค.ศ.", Narrow: "ค.ศ."}, BC: ut.CalendarEraFormatNames{Full: "ปีก่อนคริสต์ศักราช", Abbrev: "ปีก่อน ค.ศ.", Narrow: "ก่อน ค.ศ."}}}}