package khq

import "github.com/go-playground/universal-translator"

var currencies = ut.CurrencyFormatValue{"GMD": ut.Currency{Currency: "GMD", DisplayName: "Gambi Dalasi", Symbol: ""}, "ZMW": ut.Currency{Currency: "ZMW", DisplayName: "Zambi Kwaca", Symbol: ""}, "CNY": ut.Currency{Currency: "CNY", DisplayName: "Sinwa Yuan Renminbi", Symbol: ""}, "CVE": ut.Currency{Currency: "CVE", DisplayName: "Kapuver Escudo", Symbol: ""}, "USD": ut.Currency{Currency: "USD", DisplayName: "Ameriki Dollar", Symbol: ""}, "ZAR": ut.Currency{Currency: "ZAR", DisplayName: "Hawasa Afriki Rand", Symbol: ""}, "AOA": ut.Currency{Currency: "AOA", DisplayName: "Angoola Kwanza", Symbol: ""}, "BWP": ut.Currency{Currency: "BWP", DisplayName: "Botswaana Pund", Symbol: ""}, "ERN": ut.Currency{Currency: "ERN", DisplayName: "Eritree Nafka", Symbol: ""}, "MGA": ut.Currency{Currency: "MGA", DisplayName: "Malgaaši Fraŋ", Symbol: ""}, "BIF": ut.Currency{Currency: "BIF", DisplayName: "Burundi Fraŋ", Symbol: ""}, "MAD": ut.Currency{Currency: "MAD", DisplayName: "Maarok Dirham", Symbol: ""}, "NAD": ut.Currency{Currency: "NAD", DisplayName: "Naamibi Dollar", Symbol: ""}, "SZL": ut.Currency{Currency: "SZL", DisplayName: "Lilangeni", Symbol: ""}, "TND": ut.Currency{Currency: "TND", DisplayName: "Tunizi Dinar", Symbol: ""}, "CDF": ut.Currency{Currency: "CDF", DisplayName: "Kongo Fraŋ", Symbol: ""}, "EUR": ut.Currency{Currency: "EUR", DisplayName: "Eero", Symbol: ""}, "TZS": ut.Currency{Currency: "TZS", DisplayName: "Tanzaani Šiiliŋ", Symbol: ""}, "ETB": ut.Currency{Currency: "ETB", DisplayName: "Ecioopi Birr", Symbol: ""}, "GHC": ut.Currency{Currency: "GHC", DisplayName: "Gaana Šiidi", Symbol: ""}, "JPY": ut.Currency{Currency: "JPY", DisplayName: "Jaapoŋ Yen", Symbol: ""}, "ZWD": ut.Currency{Currency: "ZWD", DisplayName: "Zimbabwe Dollar", Symbol: ""}, "SDP": ut.Currency{Currency: "SDP", DisplayName: "Suudaŋ Pund", Symbol: ""}, "SAR": ut.Currency{Currency: "SAR", DisplayName: "Saudiya Riyal", Symbol: ""}, "BHD": ut.Currency{Currency: "BHD", DisplayName: "Bahareen Dinar", Symbol: ""}, "DJF": ut.Currency{Currency: "DJF", DisplayName: "Jibuuti Fraŋ", Symbol: ""}, "DZD": ut.Currency{Currency: "DZD", DisplayName: "Alžeeri Dinar", Symbol: ""}, "STD": ut.Currency{Currency: "STD", DisplayName: "Sao Tome nda Prinsipe Dobra", Symbol: ""}, "XOF": ut.Currency{Currency: "XOF", DisplayName: "CFA Fraŋ (BCEAO)", Symbol: ""}, "RWF": ut.Currency{Currency: "RWF", DisplayName: "Rwanda Fraŋ", Symbol: ""}, "AED": ut.Currency{Currency: "AED", DisplayName: "Laaraw Immaara Margantey Dirham", Symbol: ""}, "LRD": ut.Currency{Currency: "LRD", DisplayName: "Liberia Dollar", Symbol: ""}, "ZMK": ut.Currency{Currency: "ZMK", DisplayName: "Zambi Kwaca (1968–2012)", Symbol: ""}, "EGP": ut.Currency{Currency: "EGP", DisplayName: "Misra Pund", Symbol: ""}, "AUD": ut.Currency{Currency: "AUD", DisplayName: "Ostraali Dollar", Symbol: ""}, "CAD": ut.Currency{Currency: "CAD", DisplayName: "Kanaada Dollar", Symbol: ""}, "GBP": ut.Currency{Currency: "GBP", DisplayName: "Britin Pund", Symbol: ""}, "GNS": ut.Currency{Currency: "GNS", DisplayName: "Ginee Fraŋ", Symbol: ""}, "KMF": ut.Currency{Currency: "KMF", DisplayName: "Komoor Fraŋ", Symbol: ""}, "MRO": ut.Currency{Currency: "MRO", DisplayName: "Mooritaani Ugiya", Symbol: ""}, "SOS": ut.Currency{Currency: "SOS", DisplayName: "Somaali Šiiliŋ", Symbol: ""}, "UGX": ut.Currency{Currency: "UGX", DisplayName: "Uganda Šiiliŋ", Symbol: ""}, "INR": ut.Currency{Currency: "INR", DisplayName: "Indu Rupii", Symbol: ""}, "MWK": ut.Currency{Currency: "MWK", DisplayName: "Malaawi Kwaca", Symbol: ""}, "XAF": ut.Currency{Currency: "XAF", DisplayName: "CFA Fraŋ (BEAC)", Symbol: ""}, "KES": ut.Currency{Currency: "KES", DisplayName: "Keeniya Šiiliŋ", Symbol: ""}, "LYD": ut.Currency{Currency: "LYD", DisplayName: "Liibi Dinar", Symbol: ""}, "MUR": ut.Currency{Currency: "MUR", DisplayName: "Mooris Rupii", Symbol: ""}, "SCR": ut.Currency{Currency: "SCR", DisplayName: "Seešel Rupii", Symbol: ""}, "SDG": ut.Currency{Currency: "SDG", DisplayName: "Suudaŋ Dinar", Symbol: ""}, "SHP": ut.Currency{Currency: "SHP", DisplayName: "Seŋ Helena Fraŋ", Symbol: ""}, "NGN": ut.Currency{Currency: "NGN", DisplayName: "Naajiriya Neera", Symbol: ""}, "CHF": ut.Currency{Currency: "CHF", DisplayName: "Swisu Fraŋ", Symbol: ""}, "LSL": ut.Currency{Currency: "LSL", DisplayName: "Leezoto Loti", Symbol: ""}, "MZM": ut.Currency{Currency: "MZM", DisplayName: "Mozambik Metikal", Symbol: ""}, "SLL": ut.Currency{Currency: "SLL", DisplayName: "Leeon", Symbol: ""}}