package fil

import "github.com/go-playground/universal-translator"

var currencies = ut.CurrencyFormatValue{"MWK": ut.Currency{Currency: "MWK", DisplayName: "Malawian Kwacha", Symbol: "MWK"}, "PHP": ut.Currency{Currency: "PHP", DisplayName: "Piso ng Pilipinas", Symbol: "₱"}, "ZAR": ut.Currency{Currency: "ZAR", DisplayName: "Rand ng Timog Africa", Symbol: "ZAR"}, "AMD": ut.Currency{Currency: "AMD", DisplayName: "Armenian Dram", Symbol: "AMD"}, "VND": ut.Currency{Currency: "VND", DisplayName: "Vietnamese Dong", Symbol: "₫"}, "GNF": ut.Currency{Currency: "GNF", DisplayName: "Guinean Franc", Symbol: "GNF"}, "MNT": ut.Currency{Currency: "MNT", DisplayName: "Mongolian Tugrik", Symbol: "MNT"}, "NPR": ut.Currency{Currency: "NPR", DisplayName: "Nepalese Rupee", Symbol: "NPR"}, "THB": ut.Currency{Currency: "THB", DisplayName: "Thai Baht", Symbol: "฿"}, "DKK": ut.Currency{Currency: "DKK", DisplayName: "Danish Krone", Symbol: "DKK"}, "DZD": ut.Currency{Currency: "DZD", DisplayName: "Algerian Dinar", Symbol: "DZD"}, "FJD": ut.Currency{Currency: "FJD", DisplayName: "Dolyar ng Fiji", Symbol: "FJD"}, "ISK": ut.Currency{Currency: "ISK", DisplayName: "Icelandic Króna", Symbol: "ISK"}, "KWD": ut.Currency{Currency: "KWD", DisplayName: "Kuwaiti Dinar", Symbol: "KWD"}, "PKR": ut.Currency{Currency: "PKR", DisplayName: "Pakistani Rupee", Symbol: "PKR"}, "SRD": ut.Currency{Currency: "SRD", DisplayName: "Dolyar ng Suriname", Symbol: "SRD"}, "TZS": ut.Currency{Currency: "TZS", DisplayName: "Tanzanian Shilling", Symbol: "TZS"}, "BRL": ut.Currency{Currency: "BRL", DisplayName: "Real ng Barzil", Symbol: "R$"}, "CDF": ut.Currency{Currency: "CDF", DisplayName: "Congolese Franc", Symbol: "CDF"}, "GHS": ut.Currency{Currency: "GHS", DisplayName: "Ghanaian Cedi", Symbol: "GHS"}, "HTG": ut.Currency{Currency: "HTG", DisplayName: "Gourde ng Haiti", Symbol: "HTG"}, "MMK": ut.Currency{Currency: "MMK", DisplayName: "Myanmar Kyat", Symbol: "MMK"}, "MZN": ut.Currency{Currency: "MZN", DisplayName: "Mozambican Metical", Symbol: "MZN"}, "PYG": ut.Currency{Currency: "PYG", DisplayName: "Paraguayan Guarani", Symbol: "PYG"}, "SGD": ut.Currency{Currency: "SGD", DisplayName: "Dolyar ng Singapore", Symbol: "SGD"}, "BBD": ut.Currency{Currency: "BBD", DisplayName: "Dolyar ng Barbados", Symbol: "BBD"}, "CHF": ut.Currency{Currency: "CHF", DisplayName: "Swiss Franc", Symbol: "CHF"}, "COP": ut.Currency{Currency: "COP", DisplayName: "Piso ng Colombia", Symbol: "COP"}, "ZMW": ut.Currency{Currency: "ZMW", DisplayName: "Zambian Kwacha", Symbol: "ZMW"}, "BAM": ut.Currency{Currency: "BAM", DisplayName: "Bosnia-Herzegovina Convertible Mark", Symbol: "BAM"}, "JPY": ut.Currency{Currency: "JPY", DisplayName: "Japanese Yen", Symbol: "¥"}, "VEF": ut.Currency{Currency: "VEF", DisplayName: "Bolívar ng Venezuela", Symbol: "VEF"}, "BHD": ut.Currency{Currency: "BHD", DisplayName: "Bahraini Dinar", Symbol: "BHD"}, "ILS": ut.Currency{Currency: "ILS", DisplayName: "Israeli New Sheqel", Symbol: "₪"}, "KPW": ut.Currency{Currency: "KPW", DisplayName: "Won ng Hilagang Korea", Symbol: "KPW"}, "UAH": ut.Currency{Currency: "UAH", DisplayName: "Ukrainian Hryvnia", Symbol: "UAH"}, "XXX": ut.Currency{Currency: "XXX", DisplayName: "Hindi Kilalang Pera", Symbol: ""}, "BSD": ut.Currency{Currency: "BSD", DisplayName: "Dolyar ng Bahamas", Symbol: "BSD"}, "NIO": ut.Currency{Currency: "NIO", DisplayName: "Córdoba ng Nicaragua", Symbol: "NIO"}, "NOK": ut.Currency{Currency: "NOK", DisplayName: "Norwegian Krone", Symbol: "NOK"}, "ANG": ut.Currency{Currency: "ANG", DisplayName: "Antillean Guilder ng Netherlands", Symbol: "ANG"}, "CNY": ut.Currency{Currency: "CNY", DisplayName: "Chinese Yuan", Symbol: "CN¥"}, "DOP": ut.Currency{Currency: "DOP", DisplayName: "Piso ng Dominican", Symbol: "DOP"}, "EEK": ut.Currency{Currency: "EEK", DisplayName: "Estonian Kroon", Symbol: ""}, "KGS": ut.Currency{Currency: "KGS", DisplayName: "Kyrgystani Som", Symbol: "KGS"}, "MYR": ut.Currency{Currency: "MYR", DisplayName: "Malaysian Ringgit", Symbol: "MYR"}, "AED": ut.Currency{Currency: "AED", DisplayName: "United Arab Emirates Dirham", Symbol: "AED"}, "AUD": ut.Currency{Currency: "AUD", DisplayName: "Dolyar ng Australya", Symbol: "A$"}, "BDT": ut.Currency{Currency: "BDT", DisplayName: "Bangladeshi Taka", Symbol: "BDT"}, "SSP": ut.Currency{Currency: "SSP", DisplayName: "Pound ng Timog Sudan", Symbol: "SSP"}, "EUR": ut.Currency{Currency: "EUR", DisplayName: "Euro", Symbol: "€"}, "INR": ut.Currency{Currency: "INR", DisplayName: "Indian Rupee", Symbol: "₹"}, "UYU": ut.Currency{Currency: "UYU", DisplayName: "Piso ng Uruguay", Symbol: "UYU"}, "ARS": ut.Currency{Currency: "ARS", DisplayName: "Piso ng Argentina", Symbol: "ARS"}, "CZK": ut.Currency{Currency: "CZK", DisplayName: "Czech Republic Koruna", Symbol: "CZK"}, "DEM": ut.Currency{Currency: "DEM", DisplayName: "Deutsche Marks", Symbol: ""}, "TTD": ut.Currency{Currency: "TTD", DisplayName: "Dolyar ng Trinidad and Tobago", Symbol: "TTD"}, "GBP": ut.Currency{Currency: "GBP", DisplayName: "British Pound", Symbol: "£"}, "MXN": ut.Currency{Currency: "MXN", DisplayName: "Piso ng Mexico", Symbol: "MX$"}, "QAR": ut.Currency{Currency: "QAR", DisplayName: "Qatari Rial", Symbol: "QAR"}, "TWD": ut.Currency{Currency: "TWD", DisplayName: "Dolyar ng New Taiwan", Symbol: "NT$"}, "USD": ut.Currency{Currency: "USD", DisplayName: "Dolyar ng US", Symbol: "$"}, "MUR": ut.Currency{Currency: "MUR", DisplayName: "Mauritian Rupee", Symbol: "MUR"}, "OMR": ut.Currency{Currency: "OMR", DisplayName: "Omani Rial", Symbol: "OMR"}, "TJS": ut.Currency{Currency: "TJS", DisplayName: "Tajikistani Somoni", Symbol: "TJS"}, "SHP": ut.Currency{Currency: "SHP", DisplayName: "Saint Helena Pound", Symbol: "SHP"}, "EGP": ut.Currency{Currency: "EGP", DisplayName: "Egyptian Pound", Symbol: "EGP"}, "MAD": ut.Currency{Currency: "MAD", DisplayName: "Moroccan Dirham", Symbol: "MAD"}, "SBD": ut.Currency{Currency: "SBD", DisplayName: "Dolyar ng Solomon Islands", Symbol: "SBD"}, "RUB": ut.Currency{Currency: "RUB", DisplayName: "Russian Ruble", Symbol: "RUB"}, "VEB": ut.Currency{Currency: "VEB", DisplayName: "Venezuelan Bolívar (1871–2008)", Symbol: ""}, "CLP": ut.Currency{Currency: "CLP", DisplayName: "Piso ng Chile", Symbol: "CLP"}, "HUF": ut.Currency{Currency: "HUF", DisplayName: "Hungarian Forint", Symbol: "HUF"}, "KES": ut.Currency{Currency: "KES", DisplayName: "Kenyan Shilling", Symbol: "KES"}, "KZT": ut.Currency{Currency: "KZT", DisplayName: "Kazakhstani Tenge", Symbol: "KZT"}, "WST": ut.Currency{Currency: "WST", DisplayName: "Samoan Tala", Symbol: "WST"}, "AFN": ut.Currency{Currency: "AFN", DisplayName: "Afghan Afghani", Symbol: "AFN"}, "BGN": ut.Currency{Currency: "BGN", DisplayName: "Bulgarian Lev", Symbol: "BGN"}, "KRW": ut.Currency{Currency: "KRW", DisplayName: "Won ng Timog Korea", Symbol: "₩"}, "JMD": ut.Currency{Currency: "JMD", DisplayName: "Dolyar ng Jamaica", Symbol: "JMD"}, "SCR": ut.Currency{Currency: "SCR", DisplayName: "Seychellois Rupee", Symbol: "SCR"}, "YER": ut.Currency{Currency: "YER", DisplayName: "Yemeni Rial", Symbol: "YER"}, "SZL": ut.Currency{Currency: "SZL", DisplayName: "Swazi Lilangeni", Symbol: "SZL"}, "BIF": ut.Currency{Currency: "BIF", DisplayName: "Burundian Franc", Symbol: "BIF"}, "NZD": ut.Currency{Currency: "NZD", DisplayName: "Dolyar ng New Zealand", Symbol: "NZ$"}, "SIT": ut.Currency{Currency: "SIT", DisplayName: "Slovenian Tolar", Symbol: ""}, "RWF": ut.Currency{Currency: "RWF", DisplayName: "Rwandan Franc", Symbol: "RWF"}, "BWP": ut.Currency{Currency: "BWP", DisplayName: "Botswanan Pula", Symbol: "BWP"}, "FRF": ut.Currency{Currency: "FRF", DisplayName: "French Franc", Symbol: ""}, "PAB": ut.Currency{Currency: "PAB", DisplayName: "Balboa ng Panama", Symbol: "PAB"}, "LSL": ut.Currency{Currency: "LSL", DisplayName: "Lesotho Loti", Symbol: ""}, "SDG": ut.Currency{Currency: "SDG", DisplayName: "Pound ng Sudan", Symbol: "SDG"}, "ZMK": ut.Currency{Currency: "ZMK", DisplayName: "Zambian Kwacha (1968–2012)", Symbol: ""}, "AWG": ut.Currency{Currency: "AWG", DisplayName: "Florin ng Aruba", Symbol: "AWG"}, "HKD": ut.Currency{Currency: "HKD", DisplayName: "Dolyar ng Hong Kong", Symbol: "HK$"}, "LRD": ut.Currency{Currency: "LRD", DisplayName: "Dolyar ng Liberia", Symbol: "LRD"}, "GTQ": ut.Currency{Currency: "GTQ", DisplayName: "Quetzal ng Guatemala", Symbol: "GTQ"}, "UZS": ut.Currency{Currency: "UZS", DisplayName: "Uzbekistan Som", Symbol: "UZS"}, "AOA": ut.Currency{Currency: "AOA", DisplayName: "Angolan Kwanza", Symbol: "AOA"}, "BZD": ut.Currency{Currency: "BZD", DisplayName: "Dolyar ng Belize", Symbol: "BZD"}, "CVE": ut.Currency{Currency: "CVE", DisplayName: "Cape Verdean Escudo", Symbol: "CVE"}, "NAD": ut.Currency{Currency: "NAD", DisplayName: "Dolyar ng Namibia", Symbol: "NAD"}, "TND": ut.Currency{Currency: "TND", DisplayName: "Tunisian Dinar", Symbol: "TND"}, "VUV": ut.Currency{Currency: "VUV", DisplayName: "Vanuatu Vatu", Symbol: "VUV"}, "BND": ut.Currency{Currency: "BND", DisplayName: "Dolyar ng Brunei", Symbol: "BND"}, "HNL": ut.Currency{Currency: "HNL", DisplayName: "Lempira ng Honduras", Symbol: "HNL"}, "KYD": ut.Currency{Currency: "KYD", DisplayName: "Dolyar ng Cayman Islands", Symbol: "KYD"}, "JOD": ut.Currency{Currency: "JOD", DisplayName: "Jordanian Dinar", Symbol: "JOD"}, "LBP": ut.Currency{Currency: "LBP", DisplayName: "Pound ng Lebanon", Symbol: "LBP"}, "MKD": ut.Currency{Currency: "MKD", DisplayName: "Macedonian Denar", Symbol: "MKD"}, "SEK": ut.Currency{Currency: "SEK", DisplayName: "Swedish Krona", Symbol: "SEK"}, "TRY": ut.Currency{Currency: "TRY", DisplayName: "Turkish Lira", Symbol: "TRY"}, "CRC": ut.Currency{Currency: "CRC", DisplayName: "Colón ng Costa Rica", Symbol: "CRC"}, "ERN": ut.Currency{Currency: "ERN", DisplayName: "Eritrean Nakfa", Symbol: "ERN"}, "HRK": ut.Currency{Currency: "HRK", DisplayName: "Croatian Kuna", Symbol: "HRK"}, "GEL": ut.Currency{Currency: "GEL", DisplayName: "Georgian Lari", Symbol: "GEL"}, "MDL": ut.Currency{Currency: "MDL", DisplayName: "Moldovan Leu", Symbol: "MDL"}, "UGX": ut.Currency{Currency: "UGX", DisplayName: "Ugandan Shilling", Symbol: "UGX"}, "XAF": ut.Currency{Currency: "XAF", DisplayName: "CFA Franc BEAC", Symbol: "FCFA"}, "BMD": ut.Currency{Currency: "BMD", DisplayName: "Dolyar ng Bermuda", Symbol: "BMD"}, "BYR": ut.Currency{Currency: "BYR", DisplayName: "Belarusian Ruble", Symbol: "BYR"}, "CUP": ut.Currency{Currency: "CUP", DisplayName: "Piso ng Cuba", Symbol: "CUP"}, "GMD": ut.Currency{Currency: "GMD", DisplayName: "Gambian Dalasi", Symbol: "GMD"}, "PGK": ut.Currency{Currency: "PGK", DisplayName: "Papua New Guinean Kina", Symbol: "PGK"}, "PLN": ut.Currency{Currency: "PLN", DisplayName: "Polish Zloty", Symbol: "PLN"}, "ALL": ut.Currency{Currency: "ALL", DisplayName: "Albanian Lek", Symbol: "ALL"}, "AZN": ut.Currency{Currency: "AZN", DisplayName: "Azerbaijani Manat", Symbol: "AZN"}, "MRO": ut.Currency{Currency: "MRO", DisplayName: "Mauritanian Ouguiya", Symbol: "MRO"}, "LVL": ut.Currency{Currency: "LVL", DisplayName: "Latvian Lats", Symbol: "LVL"}, "TMT": ut.Currency{Currency: "TMT", DisplayName: "Turkmenistani Manat", Symbol: "TMT"}, "CUC": ut.Currency{Currency: "CUC", DisplayName: "Convertible na Piso ng Cuba", Symbol: "CUC"}, "FKP": ut.Currency{Currency: "FKP", DisplayName: "Pound ng Falkland Islands", Symbol: "FKP"}, "KMF": ut.Currency{Currency: "KMF", DisplayName: "Comorian Franc", Symbol: "KMF"}, "IDR": ut.Currency{Currency: "IDR", DisplayName: "Indonesian Rupiah", Symbol: "IDR"}, "LKR": ut.Currency{Currency: "LKR", DisplayName: "Sri Lankan Rupee", Symbol: "LKR"}, "LTL": ut.Currency{Currency: "LTL", DisplayName: "Lithuanian Litas", Symbol: "LTL"}, "MOP": ut.Currency{Currency: "MOP", DisplayName: "Macanese Pataca", Symbol: "MOP"}, "RSD": ut.Currency{Currency: "RSD", DisplayName: "Serbian Dinar", Symbol: "RSD"}, "BTN": ut.Currency{Currency: "BTN", DisplayName: "Bhutanese Ngultrum", Symbol: "BTN"}, "CAD": ut.Currency{Currency: "CAD", DisplayName: "Dolyar ng Canada", Symbol: "CA$"}, "GYD": ut.Currency{Currency: "GYD", DisplayName: "Dolyar ng Guyanese", Symbol: "GYD"}, "SOS": ut.Currency{Currency: "SOS", DisplayName: "Somali Shilling", Symbol: "SOS"}, "XCD": ut.Currency{Currency: "XCD", DisplayName: "Dolyar ng Silangang Caribbean", Symbol: "EC$"}, "XPF": ut.Currency{Currency: "XPF", DisplayName: "CFP Franc", Symbol: "CFPF"}, "DJF": ut.Currency{Currency: "DJF", DisplayName: "Djiboutian Franc", Symbol: "DJF"}, "IRR": ut.Currency{Currency: "IRR", DisplayName: "Iranian Rial", Symbol: "IRR"}, "SYP": ut.Currency{Currency: "SYP", DisplayName: "Syrian Pound", Symbol: "SYP"}, "SLL": ut.Currency{Currency: "SLL", DisplayName: "Sierra Leonean Leone", Symbol: "SLL"}, "ETB": ut.Currency{Currency: "ETB", DisplayName: "Ethiopian Birr", Symbol: "ETB"}, "LAK": ut.Currency{Currency: "LAK", DisplayName: "Laotian Kip", Symbol: "LAK"}, "LYD": ut.Currency{Currency: "LYD", DisplayName: "Libyan Dinar", Symbol: "LYD"}, "NGN": ut.Currency{Currency: "NGN", DisplayName: "Nigerian Naira", Symbol: "NGN"}, "SAR": ut.Currency{Currency: "SAR", DisplayName: "Saudi Riyal", Symbol: "SAR"}, "SKK": ut.Currency{Currency: "SKK", DisplayName: "Slovak Koruna", Symbol: ""}, "STD": ut.Currency{Currency: "STD", DisplayName: "São Tomé & Príncipe Dobra", Symbol: "STD"}, "TOP": ut.Currency{Currency: "TOP", DisplayName: "Tongan Paʻanga", Symbol: "TOP"}, "BOB": ut.Currency{Currency: "BOB", DisplayName: "Boliviano ng Bolivia", Symbol: "BOB"}, "GIP": ut.Currency{Currency: "GIP", DisplayName: "Gibraltar Pound", Symbol: "GIP"}, "MVR": ut.Currency{Currency: "MVR", DisplayName: "Maldivian Rufiyaa", Symbol: "MVR"}, "XOF": ut.Currency{Currency: "XOF", DisplayName: "CFA Franc ng Kanlurang Africa", Symbol: "CFA"}, "PEN": ut.Currency{Currency: "PEN", DisplayName: "Peruvian Nuevo Sol", Symbol: "PEN"}, "RON": ut.Currency{Currency: "RON", DisplayName: "Romanian Leu", Symbol: "RON"}, "IQD": ut.Currency{Currency: "IQD", DisplayName: "Iraqi Dinar", Symbol: "IQD"}, "KHR": ut.Currency{Currency: "KHR", DisplayName: "Cambodian Riel", Symbol: "KHR"}, "MGA": ut.Currency{Currency: "MGA", DisplayName: "Malagasy Ariary", Symbol: "MGA"}}