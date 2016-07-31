package fo

import "github.com/go-playground/universal-translator"

var currencies = ut.CurrencyFormatValue{"AMD": ut.Currency{Currency: "AMD", DisplayName: "Armenia dram", Symbol: "AMD"}, "DKK": ut.Currency{Currency: "DKK", DisplayName: "donsk króna", Symbol: "kr"}, "LKR": ut.Currency{Currency: "LKR", DisplayName: "Sri Lanka rupi", Symbol: "LKR"}, "BWP": ut.Currency{Currency: "BWP", DisplayName: "Botsvana pula", Symbol: "BWP"}, "CZK": ut.Currency{Currency: "CZK", DisplayName: "Kekkia koruna", Symbol: "CZK"}, "ETB": ut.Currency{Currency: "ETB", DisplayName: "Etiopia birr", Symbol: "ETB"}, "GNF": ut.Currency{Currency: "GNF", DisplayName: "Guinea frankur", Symbol: "GNF"}, "SBD": ut.Currency{Currency: "SBD", DisplayName: "Salomonoyggjar dollari", Symbol: "SBD"}, "CLP": ut.Currency{Currency: "CLP", DisplayName: "Kili peso", Symbol: "CLP"}, "HRK": ut.Currency{Currency: "HRK", DisplayName: "Kroatia kuna", Symbol: "HRK"}, "PAB": ut.Currency{Currency: "PAB", DisplayName: "Panama balboa", Symbol: "PAB"}, "TTD": ut.Currency{Currency: "TTD", DisplayName: "Trinidad & Tobago dollari", Symbol: "TTD"}, "BHD": ut.Currency{Currency: "BHD", DisplayName: "Barein dinar", Symbol: "BHD"}, "UYU": ut.Currency{Currency: "UYU", DisplayName: "Uruguai peso", Symbol: "UYU"}, "XPD": ut.Currency{Currency: "XPD", DisplayName: "unse palladium", Symbol: ""}, "GEL": ut.Currency{Currency: "GEL", DisplayName: "Georgia lari", Symbol: "GEL"}, "HNL": ut.Currency{Currency: "HNL", DisplayName: "Honduras lempira", Symbol: "HNL"}, "KHR": ut.Currency{Currency: "KHR", DisplayName: "Kambodja riel", Symbol: "KHR"}, "SLL": ut.Currency{Currency: "SLL", DisplayName: "Sierra Leona leone", Symbol: "SLL"}, "TOP": ut.Currency{Currency: "TOP", DisplayName: "Tonga paʻanga", Symbol: "TOP"}, "BRL": ut.Currency{Currency: "BRL", DisplayName: "Brasilianskur real", Symbol: "R$"}, "MAD": ut.Currency{Currency: "MAD", DisplayName: "Marokko dirham", Symbol: "MAD"}, "MKD": ut.Currency{Currency: "MKD", DisplayName: "Makedónia denar", Symbol: "MKD"}, "AOA": ut.Currency{Currency: "AOA", DisplayName: "Angola kwanza", Symbol: "AOA"}, "COP": ut.Currency{Currency: "COP", DisplayName: "Kolombia peso", Symbol: "COP"}, "NZD": ut.Currency{Currency: "NZD", DisplayName: "Nýsæland dollari", Symbol: "NZ$"}, "UGX": ut.Currency{Currency: "UGX", DisplayName: "Uganda skillingur", Symbol: "UGX"}, "TMT": ut.Currency{Currency: "TMT", DisplayName: "Turkmenistan manat", Symbol: "TMT"}, "UAH": ut.Currency{Currency: "UAH", DisplayName: "Ukraina hryvnia", Symbol: "UAH"}, "XAF": ut.Currency{Currency: "XAF", DisplayName: "Miðafrika CFA frankur", Symbol: "FCFA"}, "YER": ut.Currency{Currency: "YER", DisplayName: "Jemen rial", Symbol: "YER"}, "CUC": ut.Currency{Currency: "CUC", DisplayName: "Kuba peso (sum kann vekslast)", Symbol: "CUC"}, "JPY": ut.Currency{Currency: "JPY", DisplayName: "japanskur yen", Symbol: "JP¥"}, "KPW": ut.Currency{Currency: "KPW", DisplayName: "Norðurkorea won", Symbol: "KPW"}, "PEN": ut.Currency{Currency: "PEN", DisplayName: "Peru nuevo sol", Symbol: "PEN"}, "NAD": ut.Currency{Currency: "NAD", DisplayName: "Namibia dollari", Symbol: "NAD"}, "RWF": ut.Currency{Currency: "RWF", DisplayName: "Ruanda frankur", Symbol: "RWF"}, "SRD": ut.Currency{Currency: "SRD", DisplayName: "Surinam dollari", Symbol: "SRD"}, "AFN": ut.Currency{Currency: "AFN", DisplayName: "Afganistan afghani", Symbol: "AFN"}, "AWG": ut.Currency{Currency: "AWG", DisplayName: "Aruba florin", Symbol: "AWG"}, "BSD": ut.Currency{Currency: "BSD", DisplayName: "Bahamaoyggjar dollari", Symbol: "BSD"}, "CAD": ut.Currency{Currency: "CAD", DisplayName: "Kanada dollari", Symbol: "CA$"}, "KGS": ut.Currency{Currency: "KGS", DisplayName: "Kirgisia som", Symbol: "KGS"}, "XPF": ut.Currency{Currency: "XPF", DisplayName: "CFP frankur", Symbol: "CFPF"}, "SZL": ut.Currency{Currency: "SZL", DisplayName: "Svasiland lilangeni", Symbol: "SZL"}, "TRY": ut.Currency{Currency: "TRY", DisplayName: "Turkaland liri", Symbol: "TRY"}, "USD": ut.Currency{Currency: "USD", DisplayName: "US dollari", Symbol: "US$"}, "AED": ut.Currency{Currency: "AED", DisplayName: "Sameindu Emirríkini dirham", Symbol: "AED"}, "BND": ut.Currency{Currency: "BND", DisplayName: "Brunei dollari", Symbol: "BND"}, "CNY": ut.Currency{Currency: "CNY", DisplayName: "kinesiskur yuan", Symbol: "CN¥"}, "ERN": ut.Currency{Currency: "ERN", DisplayName: "Eritrea nakfa", Symbol: "ERN"}, "SOS": ut.Currency{Currency: "SOS", DisplayName: "Somalia skillingur", Symbol: "SOS"}, "BIF": ut.Currency{Currency: "BIF", DisplayName: "Burundi frankur", Symbol: "BIF"}, "CDF": ut.Currency{Currency: "CDF", DisplayName: "Kongo frankur", Symbol: "CDF"}, "LBP": ut.Currency{Currency: "LBP", DisplayName: "Libanon pund", Symbol: "LBP"}, "BBD": ut.Currency{Currency: "BBD", DisplayName: "Barbados dollari", Symbol: "BBD"}, "BDT": ut.Currency{Currency: "BDT", DisplayName: "Bangladesj taka", Symbol: "BDT"}, "EUR": ut.Currency{Currency: "EUR", DisplayName: "Evra", Symbol: "€"}, "MMK": ut.Currency{Currency: "MMK", DisplayName: "Myanmar (Burma) kyat", Symbol: "MMK"}, "SHP": ut.Currency{Currency: "SHP", DisplayName: "St. Helena pund", Symbol: "SHP"}, "KWD": ut.Currency{Currency: "KWD", DisplayName: "Kuvait dinar", Symbol: "KWD"}, "PKR": ut.Currency{Currency: "PKR", DisplayName: "Pakistan rupi", Symbol: "PKR"}, "XPT": ut.Currency{Currency: "XPT", DisplayName: "unse platin", Symbol: ""}, "BAM": ut.Currency{Currency: "BAM", DisplayName: "Bosnia-Hersegovina mark (kann vekslast)", Symbol: "BAM"}, "CRC": ut.Currency{Currency: "CRC", DisplayName: "Kosta Rika colón", Symbol: "CRC"}, "DOP": ut.Currency{Currency: "DOP", DisplayName: "Dominika peso", Symbol: "DOP"}, "IDR": ut.Currency{Currency: "IDR", DisplayName: "Indonesia rupiah", Symbol: "IDR"}, "JMD": ut.Currency{Currency: "JMD", DisplayName: "Jamaika dollari", Symbol: "JMD"}, "LAK": ut.Currency{Currency: "LAK", DisplayName: "Laos kip", Symbol: "LAK"}, "NPR": ut.Currency{Currency: "NPR", DisplayName: "Nepal rupi", Symbol: "NPR"}, "THB": ut.Currency{Currency: "THB", DisplayName: "Tailand baht", Symbol: "THB"}, "XCD": ut.Currency{Currency: "XCD", DisplayName: "Eystur Karibia dollari", Symbol: "EC$"}, "XXX": ut.Currency{Currency: "XXX", DisplayName: "ókent gjaldoyra", Symbol: ""}, "BTN": ut.Currency{Currency: "BTN", DisplayName: "Butan ngultrum", Symbol: "BTN"}, "BYR": ut.Currency{Currency: "BYR", DisplayName: "Hvítarussland ruble", Symbol: "BYR"}, "EGP": ut.Currency{Currency: "EGP", DisplayName: "Egyptaland pund", Symbol: "EGP"}, "FKP": ut.Currency{Currency: "FKP", DisplayName: "Falklandsoyggjar pund", Symbol: "FKP"}, "SGD": ut.Currency{Currency: "SGD", DisplayName: "Singapor dollari", Symbol: "SGD"}, "GMD": ut.Currency{Currency: "GMD", DisplayName: "Gambia dalasi", Symbol: "GMD"}, "IRR": ut.Currency{Currency: "IRR", DisplayName: "iranskir rials", Symbol: "IRR"}, "SYP": ut.Currency{Currency: "SYP", DisplayName: "Sýria pund", Symbol: "SYP"}, "ZMW": ut.Currency{Currency: "ZMW", DisplayName: "Sambia kwacha", Symbol: "ZMW"}, "BZD": ut.Currency{Currency: "BZD", DisplayName: "Belis dollari", Symbol: "BZD"}, "ISK": ut.Currency{Currency: "ISK", DisplayName: "íslendsk króna", Symbol: "ISK"}, "RON": ut.Currency{Currency: "RON", DisplayName: "Rumenia leu", Symbol: "RON"}, "SDG": ut.Currency{Currency: "SDG", DisplayName: "Sudan pund", Symbol: "SDG"}, "SSP": ut.Currency{Currency: "SSP", DisplayName: "Suðursudan pund", Symbol: "SSP"}, "VND": ut.Currency{Currency: "VND", DisplayName: "Vjetnam dong", Symbol: "₫"}, "ARS": ut.Currency{Currency: "ARS", DisplayName: "Argentina peso", Symbol: "ARS"}, "AZN": ut.Currency{Currency: "AZN", DisplayName: "Aserbadjan manat", Symbol: "AZN"}, "BGN": ut.Currency{Currency: "BGN", DisplayName: "Bulgaria lev", Symbol: "BGN"}, "INR": ut.Currency{Currency: "INR", DisplayName: "indiskir rupis", Symbol: "₹"}, "JOD": ut.Currency{Currency: "JOD", DisplayName: "Jordan dinar", Symbol: "JOD"}, "DJF": ut.Currency{Currency: "DJF", DisplayName: "Djibuti frankur", Symbol: "DJF"}, "VEF": ut.Currency{Currency: "VEF", DisplayName: "Venesuela bolívar", Symbol: "VEF"}, "XAG": ut.Currency{Currency: "XAG", DisplayName: "unse sølv", Symbol: ""}, "XOF": ut.Currency{Currency: "XOF", DisplayName: "Vesturafrika CFA frankur", Symbol: "CFA"}, "WST": ut.Currency{Currency: "WST", DisplayName: "Samoa tala", Symbol: "WST"}, "LYD": ut.Currency{Currency: "LYD", DisplayName: "Libya dinar", Symbol: "LYD"}, "NIO": ut.Currency{Currency: "NIO", DisplayName: "Nikaragua córdoba", Symbol: "NIO"}, "PYG": ut.Currency{Currency: "PYG", DisplayName: "Paraguai guarani", Symbol: "PYG"}, "RSD": ut.Currency{Currency: "RSD", DisplayName: "Serbia dinar", Symbol: "RSD"}, "VUV": ut.Currency{Currency: "VUV", DisplayName: "Vanuatu vatu", Symbol: "VUV"}, "GHS": ut.Currency{Currency: "GHS", DisplayName: "Gana cedi", Symbol: "GHS"}, "HUF": ut.Currency{Currency: "HUF", DisplayName: "Ungarn forint", Symbol: "HUF"}, "MOP": ut.Currency{Currency: "MOP", DisplayName: "Makao pataca", Symbol: "MOP"}, "MRO": ut.Currency{Currency: "MRO", DisplayName: "Móritania ouguiya", Symbol: "MRO"}, "OMR": ut.Currency{Currency: "OMR", DisplayName: "Oman rial", Symbol: "OMR"}, "MGA": ut.Currency{Currency: "MGA", DisplayName: "Madagaskar ariary", Symbol: "MGA"}, "FJD": ut.Currency{Currency: "FJD", DisplayName: "Fiji dollari", Symbol: "FJD"}, "GYD": ut.Currency{Currency: "GYD", DisplayName: "Gujana dollari", Symbol: "GYD"}, "KRW": ut.Currency{Currency: "KRW", DisplayName: "Suðurkorea won", Symbol: "₩"}, "KYD": ut.Currency{Currency: "KYD", DisplayName: "Caymanoyggjar dollari", Symbol: "KYD"}, "UZS": ut.Currency{Currency: "UZS", DisplayName: "Usbekistan som", Symbol: "UZS"}, "MVR": ut.Currency{Currency: "MVR", DisplayName: "Maldivoyggjar rufiyaa", Symbol: "MVR"}, "MWK": ut.Currency{Currency: "MWK", DisplayName: "Malavi kwacha", Symbol: "MWK"}, "NGN": ut.Currency{Currency: "NGN", DisplayName: "Nigeria naira", Symbol: "NGN"}, "RUB": ut.Currency{Currency: "RUB", DisplayName: "Russland ruble", Symbol: "RUB"}, "TWD": ut.Currency{Currency: "TWD", DisplayName: "Taivan new dollari", Symbol: "NT$"}, "ILS": ut.Currency{Currency: "ILS", DisplayName: "Ísrael new sheqel", Symbol: "₪"}, "MNT": ut.Currency{Currency: "MNT", DisplayName: "Mongolia tugrik", Symbol: "MNT"}, "QAR": ut.Currency{Currency: "QAR", DisplayName: "Katar rial", Symbol: "QAR"}, "ALL": ut.Currency{Currency: "ALL", DisplayName: "Albania lek", Symbol: "ALL"}, "AUD": ut.Currency{Currency: "AUD", DisplayName: "Avstralskur dollari", Symbol: "A$"}, "CHF": ut.Currency{Currency: "CHF", DisplayName: "sveisiskur frankur", Symbol: "CHF"}, "CVE": ut.Currency{Currency: "CVE", DisplayName: "Grønhøvdaoyggjar escudo", Symbol: "CVE"}, "GIP": ut.Currency{Currency: "GIP", DisplayName: "Gibraltar pund", Symbol: "GIP"}, "SAR": ut.Currency{Currency: "SAR", DisplayName: "Saudiarabia riyal", Symbol: "SAR"}, "SEK": ut.Currency{Currency: "SEK", DisplayName: "svensk króna", Symbol: "SEK"}, "MZN": ut.Currency{Currency: "MZN", DisplayName: "Mosambik metical", Symbol: "MZN"}, "PLN": ut.Currency{Currency: "PLN", DisplayName: "Pólland zloty", Symbol: "PLN"}, "STD": ut.Currency{Currency: "STD", DisplayName: "Sao Tome & Prinsipi dobra", Symbol: "STD"}, "ANG": ut.Currency{Currency: "ANG", DisplayName: "Niðurlonds Karibia gyllin", Symbol: "ANG"}, "BOB": ut.Currency{Currency: "BOB", DisplayName: "Bolivia boliviano", Symbol: "BOB"}, "GTQ": ut.Currency{Currency: "GTQ", DisplayName: "Guatemala quetzal", Symbol: "GTQ"}, "LRD": ut.Currency{Currency: "LRD", DisplayName: "Liberia dollari", Symbol: "LRD"}, "MUR": ut.Currency{Currency: "MUR", DisplayName: "Móritius rupi", Symbol: "MUR"}, "BMD": ut.Currency{Currency: "BMD", DisplayName: "Bermuda dollari", Symbol: "BMD"}, "DZD": ut.Currency{Currency: "DZD", DisplayName: "Algeria dinar", Symbol: "DZD"}, "MYR": ut.Currency{Currency: "MYR", DisplayName: "Malaisia ringgit", Symbol: "MYR"}, "HKD": ut.Currency{Currency: "HKD", DisplayName: "Hong Kong dollari", Symbol: "HK$"}, "IQD": ut.Currency{Currency: "IQD", DisplayName: "Irak dinar", Symbol: "IQD"}, "MXN": ut.Currency{Currency: "MXN", DisplayName: "Meksiko peso", Symbol: "MX$"}, "SCR": ut.Currency{Currency: "SCR", DisplayName: "Seyskelloyggjar rupi", Symbol: "SCR"}, "TND": ut.Currency{Currency: "TND", DisplayName: "Tunesia dinar", Symbol: "TND"}, "CUP": ut.Currency{Currency: "CUP", DisplayName: "Kuba peso", Symbol: "CUP"}, "KMF": ut.Currency{Currency: "KMF", DisplayName: "Komoroyggjar frankur", Symbol: "KMF"}, "KZT": ut.Currency{Currency: "KZT", DisplayName: "Kasakstan tenge", Symbol: "KZT"}, "TZS": ut.Currency{Currency: "TZS", DisplayName: "Tansania skillingur", Symbol: "TZS"}, "XAU": ut.Currency{Currency: "XAU", DisplayName: "unse guld", Symbol: ""}, "HTG": ut.Currency{Currency: "HTG", DisplayName: "Haiti gourde", Symbol: "HTG"}, "MDL": ut.Currency{Currency: "MDL", DisplayName: "Moldova leu", Symbol: "MDL"}, "TJS": ut.Currency{Currency: "TJS", DisplayName: "Tadsjikistan somoni", Symbol: "TJS"}, "ZAR": ut.Currency{Currency: "ZAR", DisplayName: "Suðurafrika rand", Symbol: "ZAR"}, "GBP": ut.Currency{Currency: "GBP", DisplayName: "bretsk pund", Symbol: "£"}, "KES": ut.Currency{Currency: "KES", DisplayName: "kenjanskur skillingur", Symbol: "KES"}, "NOK": ut.Currency{Currency: "NOK", DisplayName: "norsk króna", Symbol: "NOK"}, "PGK": ut.Currency{Currency: "PGK", DisplayName: "Papua Nýguinea kina", Symbol: "PGK"}, "PHP": ut.Currency{Currency: "PHP", DisplayName: "Filipsoyggjar peso", Symbol: "PHP"}}