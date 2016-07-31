package ksb

import "github.com/go-playground/universal-translator"

var currencies = ut.CurrencyFormatValue{"EGP": ut.Currency{Currency: "EGP", DisplayName: "pauni ya Misli", Symbol: ""}, "LRD": ut.Currency{Currency: "LRD", DisplayName: "dola ya Libelia", Symbol: ""}, "BHD": ut.Currency{Currency: "BHD", DisplayName: "dinali ya Bahaleni", Symbol: ""}, "GMD": ut.Currency{Currency: "GMD", DisplayName: "dalasi ya Gambia", Symbol: ""}, "KMF": ut.Currency{Currency: "KMF", DisplayName: "falanga ya Komolo", Symbol: ""}, "MUR": ut.Currency{Currency: "MUR", DisplayName: "lupia ya Molisi", Symbol: ""}, "TZS": ut.Currency{Currency: "TZS", DisplayName: "shilingi ya Tanzania", Symbol: "TSh"}, "USD": ut.Currency{Currency: "USD", DisplayName: "dola ya Malekani", Symbol: ""}, "ZMW": ut.Currency{Currency: "ZMW", DisplayName: "kwacha ya Zambia", Symbol: ""}, "CVE": ut.Currency{Currency: "CVE", DisplayName: "eskudo ya Kepuvede", Symbol: ""}, "ZMK": ut.Currency{Currency: "ZMK", DisplayName: "kwacha ya Zambia (1968–2012)", Symbol: ""}, "STD": ut.Currency{Currency: "STD", DisplayName: "dobla ya Sao Tome na Plincipe", Symbol: ""}, "BWP": ut.Currency{Currency: "BWP", DisplayName: "pula ya Botswana", Symbol: ""}, "MGA": ut.Currency{Currency: "MGA", DisplayName: "falanga ya Bukini", Symbol: ""}, "RWF": ut.Currency{Currency: "RWF", DisplayName: "falanga ya Lwanda", Symbol: ""}, "XOF": ut.Currency{Currency: "XOF", DisplayName: "falanga CFA BCEAO", Symbol: ""}, "GNS": ut.Currency{Currency: "GNS", DisplayName: "falanga ya Gine", Symbol: ""}, "SOS": ut.Currency{Currency: "SOS", DisplayName: "shilingi ya Somalia", Symbol: ""}, "ERN": ut.Currency{Currency: "ERN", DisplayName: "nakfa ya Elitlea", Symbol: ""}, "AOA": ut.Currency{Currency: "AOA", DisplayName: "kwanza ya Angola", Symbol: ""}, "AUD": ut.Currency{Currency: "AUD", DisplayName: "dola ya Austlalia", Symbol: ""}, "CHF": ut.Currency{Currency: "CHF", DisplayName: "falanga ya Uswisi", Symbol: ""}, "JPY": ut.Currency{Currency: "JPY", DisplayName: "salafu ya Kijapani", Symbol: ""}, "MAD": ut.Currency{Currency: "MAD", DisplayName: "dilham ya Moloko", Symbol: ""}, "AED": ut.Currency{Currency: "AED", DisplayName: "dilham ya Falme za Kialabu", Symbol: ""}, "DZD": ut.Currency{Currency: "DZD", DisplayName: "dinali ya Aljelia", Symbol: ""}, "XAF": ut.Currency{Currency: "XAF", DisplayName: "falanga CFA BEAC", Symbol: ""}, "DJF": ut.Currency{Currency: "DJF", DisplayName: "falanga ya Jibuti", Symbol: ""}, "CDF": ut.Currency{Currency: "CDF", DisplayName: "falanga ya Kongo", Symbol: ""}, "EUR": ut.Currency{Currency: "EUR", DisplayName: "yulo", Symbol: ""}, "LYD": ut.Currency{Currency: "LYD", DisplayName: "dinali ya Libya", Symbol: ""}, "SCR": ut.Currency{Currency: "SCR", DisplayName: "lupia ya Shelisheli", Symbol: ""}, "NAD": ut.Currency{Currency: "NAD", DisplayName: "dola ya Namibia", Symbol: ""}, "UGX": ut.Currency{Currency: "UGX", DisplayName: "shilingi ya Uganda", Symbol: ""}, "ZAR": ut.Currency{Currency: "ZAR", DisplayName: "landi ya Aflika Kusini", Symbol: ""}, "KES": ut.Currency{Currency: "KES", DisplayName: "shilingi ya Kenya", Symbol: ""}, "TND": ut.Currency{Currency: "TND", DisplayName: "dinali ya Tunisia", Symbol: ""}, "BIF": ut.Currency{Currency: "BIF", DisplayName: "falanga ya Bulundi", Symbol: ""}, "SDP": ut.Currency{Currency: "SDP", DisplayName: "pauni ya Sudani", Symbol: ""}, "NGN": ut.Currency{Currency: "NGN", DisplayName: "naila ya Naijelia", Symbol: ""}, "ETB": ut.Currency{Currency: "ETB", DisplayName: "bil ya Uhabeshi", Symbol: ""}, "LSL": ut.Currency{Currency: "LSL", DisplayName: "loti ya Lesoto", Symbol: ""}, "SAR": ut.Currency{Currency: "SAR", DisplayName: "liyal ya Saudia", Symbol: ""}, "SHP": ut.Currency{Currency: "SHP", DisplayName: "pauni ya Santahelena", Symbol: ""}, "SLL": ut.Currency{Currency: "SLL", DisplayName: "leoni", Symbol: ""}, "SZL": ut.Currency{Currency: "SZL", DisplayName: "lilangeni", Symbol: ""}, "CNY": ut.Currency{Currency: "CNY", DisplayName: "yaun lenminbi ya China", Symbol: ""}, "GHC": ut.Currency{Currency: "GHC", DisplayName: "sedi ya Ghana", Symbol: ""}, "MZM": ut.Currency{Currency: "MZM", DisplayName: "metikali ya Msumbiji", Symbol: ""}, "ZWD": ut.Currency{Currency: "ZWD", DisplayName: "dola ya Zimbabwe", Symbol: ""}, "CAD": ut.Currency{Currency: "CAD", DisplayName: "dola ya Kanada", Symbol: ""}, "INR": ut.Currency{Currency: "INR", DisplayName: "lupia ya India", Symbol: ""}, "MRO": ut.Currency{Currency: "MRO", DisplayName: "ugwiya ya Molitania", Symbol: ""}, "MWK": ut.Currency{Currency: "MWK", DisplayName: "kwacha ya Malawi", Symbol: ""}, "GBP": ut.Currency{Currency: "GBP", DisplayName: "pauni ya Uingeeza", Symbol: ""}, "SDG": ut.Currency{Currency: "SDG", DisplayName: "dinali ya Sudani", Symbol: ""}}