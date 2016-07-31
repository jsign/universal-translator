package it

import "github.com/go-playground/universal-translator"

var currencies = ut.CurrencyFormatValue{"GWP": ut.Currency{Currency: "GWP", DisplayName: "peso della Guinea-Bissau", Symbol: ""}, "RUB": ut.Currency{Currency: "RUB", DisplayName: "rublo russo", Symbol: "RUB"}, "SZL": ut.Currency{Currency: "SZL", DisplayName: "lilangeni dello Swaziland", Symbol: "SZL"}, "MYR": ut.Currency{Currency: "MYR", DisplayName: "ringgit malese", Symbol: "MYR"}, "SAR": ut.Currency{Currency: "SAR", DisplayName: "riyal saudita", Symbol: "SAR"}, "SYP": ut.Currency{Currency: "SYP", DisplayName: "lira siriana", Symbol: "SYP"}, "XTS": ut.Currency{Currency: "XTS", DisplayName: "codice di verifica della valuta", Symbol: ""}, "ZAR": ut.Currency{Currency: "ZAR", DisplayName: "rand sudafricano", Symbol: "ZAR"}, "DJF": ut.Currency{Currency: "DJF", DisplayName: "franco di Gibuti", Symbol: "DJF"}, "GMD": ut.Currency{Currency: "GMD", DisplayName: "dalasi gambiano", Symbol: "GMD"}, "KWD": ut.Currency{Currency: "KWD", DisplayName: "dinaro kuwaitiano", Symbol: "KWD"}, "LVL": ut.Currency{Currency: "LVL", DisplayName: "lats lettone", Symbol: "LVL"}, "STD": ut.Currency{Currency: "STD", DisplayName: "dobra di Sao Tomé e Principe", Symbol: "STD"}, "HRD": ut.Currency{Currency: "HRD", DisplayName: "dinaro croato", Symbol: ""}, "NGN": ut.Currency{Currency: "NGN", DisplayName: "naira nigeriana", Symbol: "NGN"}, "PLZ": ut.Currency{Currency: "PLZ", DisplayName: "złoty Polacco (1950–1995)", Symbol: ""}, "UGS": ut.Currency{Currency: "UGS", DisplayName: "scellino ugandese (1966–1987)", Symbol: ""}, "CRC": ut.Currency{Currency: "CRC", DisplayName: "colón costaricano", Symbol: "CRC"}, "ZMK": ut.Currency{Currency: "ZMK", DisplayName: "kwacha dello Zambia (1968–2012)", Symbol: ""}, "BOV": ut.Currency{Currency: "BOV", DisplayName: "mvdol boliviano", Symbol: ""}, "ESA": ut.Currency{Currency: "ESA", DisplayName: "peseta spagnola account", Symbol: ""}, "GEK": ut.Currency{Currency: "GEK", DisplayName: "kupon larit georgiano", Symbol: ""}, "MDL": ut.Currency{Currency: "MDL", DisplayName: "leu moldavo", Symbol: "MDL"}, "MVR": ut.Currency{Currency: "MVR", DisplayName: "rufiyaa delle Maldive", Symbol: "MVR"}, "PAB": ut.Currency{Currency: "PAB", DisplayName: "balboa panamense", Symbol: "PAB"}, "SRD": ut.Currency{Currency: "SRD", DisplayName: "dollaro del Suriname", Symbol: "SRD"}, "UYI": ut.Currency{Currency: "UYI", DisplayName: "peso uruguaiano in unità indicizzate", Symbol: ""}, "BRE": ut.Currency{Currency: "BRE", DisplayName: "cruzeiro brasiliano (1990–1993)", Symbol: ""}, "INR": ut.Currency{Currency: "INR", DisplayName: "rupia indiana", Symbol: "₹"}, "MGA": ut.Currency{Currency: "MGA", DisplayName: "ariary malgascio", Symbol: "MGA"}, "MGF": ut.Currency{Currency: "MGF", DisplayName: "franco malgascio", Symbol: ""}, "NIO": ut.Currency{Currency: "NIO", DisplayName: "córdoba nicaraguense", Symbol: "NIO"}, "ZAL": ut.Currency{Currency: "ZAL", DisplayName: "rand sudafricano (finanziario)", Symbol: ""}, "DZD": ut.Currency{Currency: "DZD", DisplayName: "dinaro algerino", Symbol: "DZD"}, "TMM": ut.Currency{Currency: "TMM", DisplayName: "manat turkmeno (1993–2009)", Symbol: ""}, "TMT": ut.Currency{Currency: "TMT", DisplayName: "manat turkmeno", Symbol: "TMT"}, "XOF": ut.Currency{Currency: "XOF", DisplayName: "franco CFA BCEAO", Symbol: "CFA"}, "ZWL": ut.Currency{Currency: "ZWL", DisplayName: "dollaro zimbabwiano (2009)", Symbol: ""}, "PLN": ut.Currency{Currency: "PLN", DisplayName: "złoty polacco", Symbol: "PLN"}, "SHP": ut.Currency{Currency: "SHP", DisplayName: "sterlina di Sant’Elena", Symbol: "SHP"}, "DKK": ut.Currency{Currency: "DKK", DisplayName: "corona danese", Symbol: "DKK"}, "FKP": ut.Currency{Currency: "FKP", DisplayName: "sterlina delle Falkland", Symbol: "FKP"}, "MRO": ut.Currency{Currency: "MRO", DisplayName: "ouguiya della Mauritania", Symbol: "MRO"}, "MXV": ut.Currency{Currency: "MXV", DisplayName: "unidad de inversion (UDI) messicana", Symbol: ""}, "MZE": ut.Currency{Currency: "MZE", DisplayName: "escudo del Mozambico", Symbol: ""}, "BMD": ut.Currency{Currency: "BMD", DisplayName: "dollaro delle Bermuda", Symbol: "BMD"}, "EGP": ut.Currency{Currency: "EGP", DisplayName: "sterlina egiziana", Symbol: "EGP"}, "MXP": ut.Currency{Currency: "MXP", DisplayName: "peso messicano d’argento (1861–1992)", Symbol: ""}, "UGX": ut.Currency{Currency: "UGX", DisplayName: "scellino ugandese", Symbol: "UGX"}, "XBA": ut.Currency{Currency: "XBA", DisplayName: "unità composita europea", Symbol: ""}, "AFA": ut.Currency{Currency: "AFA", DisplayName: "afgani (1927–2002)", Symbol: ""}, "CUP": ut.Currency{Currency: "CUP", DisplayName: "peso cubano", Symbol: "CUP"}, "ESP": ut.Currency{Currency: "ESP", DisplayName: "peseta spagnola", Symbol: ""}, "ITL": ut.Currency{Currency: "ITL", DisplayName: "lira italiana", Symbol: ""}, "VEF": ut.Currency{Currency: "VEF", DisplayName: "bolívar venezuelano", Symbol: "VEF"}, "XBC": ut.Currency{Currency: "XBC", DisplayName: "unità di acconto europea (XBC)", Symbol: ""}, "ZMW": ut.Currency{Currency: "ZMW", DisplayName: "kwacha dello Zambia", Symbol: "ZMW"}, "BEL": ut.Currency{Currency: "BEL", DisplayName: "franco belga (finanziario)", Symbol: ""}, "CLP": ut.Currency{Currency: "CLP", DisplayName: "peso cileno", Symbol: "CLP"}, "IDR": ut.Currency{Currency: "IDR", DisplayName: "rupia indonesiana", Symbol: "IDR"}, "PTE": ut.Currency{Currency: "PTE", DisplayName: "escudo portoghese", Symbol: ""}, "UAH": ut.Currency{Currency: "UAH", DisplayName: "grivnia ucraina", Symbol: "UAH"}, "BRC": ut.Currency{Currency: "BRC", DisplayName: "cruzado brasiliano", Symbol: ""}, "ECS": ut.Currency{Currency: "ECS", DisplayName: "sucre dell’Ecuador", Symbol: ""}, "FJD": ut.Currency{Currency: "FJD", DisplayName: "dollaro delle Figi", Symbol: "FJD"}, "MUR": ut.Currency{Currency: "MUR", DisplayName: "rupia mauriziana", Symbol: "MUR"}, "ZRZ": ut.Currency{Currency: "ZRZ", DisplayName: "zaire dello Zaire", Symbol: ""}, "BAD": ut.Currency{Currency: "BAD", DisplayName: "dinar Bosnia-Herzegovina", Symbol: ""}, "PKR": ut.Currency{Currency: "PKR", DisplayName: "rupia pakistana", Symbol: "PKR"}, "SBD": ut.Currency{Currency: "SBD", DisplayName: "dollaro delle Isole Salomone", Symbol: "SBD"}, "SCR": ut.Currency{Currency: "SCR", DisplayName: "rupia delle Seychelles", Symbol: "SCR"}, "AOK": ut.Currency{Currency: "AOK", DisplayName: "kwanza angolano (1977–1990)", Symbol: ""}, "CSD": ut.Currency{Currency: "CSD", DisplayName: "antico dinaro serbo", Symbol: ""}, "GRD": ut.Currency{Currency: "GRD", DisplayName: "dracma greca", Symbol: ""}, "HKD": ut.Currency{Currency: "HKD", DisplayName: "dollaro di Hong Kong", Symbol: "HKD"}, "YDD": ut.Currency{Currency: "YDD", DisplayName: "dinaro dello Yemen", Symbol: ""}, "IEP": ut.Currency{Currency: "IEP", DisplayName: "sterlina irlandese", Symbol: ""}, "YER": ut.Currency{Currency: "YER", DisplayName: "riyal yemenita", Symbol: "YER"}, "CDF": ut.Currency{Currency: "CDF", DisplayName: "franco congolese", Symbol: "CDF"}, "GIP": ut.Currency{Currency: "GIP", DisplayName: "sterlina di Gibilterra", Symbol: "GIP"}, "MNT": ut.Currency{Currency: "MNT", DisplayName: "tugrik mongolo", Symbol: "MNT"}, "RUR": ut.Currency{Currency: "RUR", DisplayName: "rublo della CSI", Symbol: ""}, "DDM": ut.Currency{Currency: "DDM", DisplayName: "ostmark della Germania Orientale", Symbol: ""}, "MLF": ut.Currency{Currency: "MLF", DisplayName: "franco di Mali", Symbol: ""}, "MTL": ut.Currency{Currency: "MTL", DisplayName: "lira maltese", Symbol: ""}, "XAF": ut.Currency{Currency: "XAF", DisplayName: "franco CFA BEAC", Symbol: "FCFA"}, "HRK": ut.Currency{Currency: "HRK", DisplayName: "kuna croata", Symbol: "HRK"}, "KGS": ut.Currency{Currency: "KGS", DisplayName: "som kirghiso", Symbol: "KGS"}, "XPT": ut.Currency{Currency: "XPT", DisplayName: "platino", Symbol: ""}, "XXX": ut.Currency{Currency: "XXX", DisplayName: "valuta sconosciuta", Symbol: ""}, "BSD": ut.Currency{Currency: "BSD", DisplayName: "dollaro delle Bahamas", Symbol: "BSD"}, "BTN": ut.Currency{Currency: "BTN", DisplayName: "ngultrum bhutanese", Symbol: "BTN"}, "ILS": ut.Currency{Currency: "ILS", DisplayName: "nuovo siclo israeliano", Symbol: "₪"}, "ROL": ut.Currency{Currency: "ROL", DisplayName: "leu della Romania", Symbol: ""}, "TRL": ut.Currency{Currency: "TRL", DisplayName: "lira turca (1922–2005)", Symbol: ""}, "SDG": ut.Currency{Currency: "SDG", DisplayName: "sterlina sudanese", Symbol: "SDG"}, "YUM": ut.Currency{Currency: "YUM", DisplayName: "dinaro noviy yugoslavo", Symbol: ""}, "BRB": ut.Currency{Currency: "BRB", DisplayName: "cruzeiro novo brasiliano (1967–1986)", Symbol: ""}, "ISK": ut.Currency{Currency: "ISK", DisplayName: "corona islandese", Symbol: "ISK"}, "LBP": ut.Currency{Currency: "LBP", DisplayName: "lira libanese", Symbol: "LBP"}, "MWK": ut.Currency{Currency: "MWK", DisplayName: "kwacha malawiano", Symbol: "MWK"}, "PHP": ut.Currency{Currency: "PHP", DisplayName: "peso filippino", Symbol: "PHP"}, "MMK": ut.Currency{Currency: "MMK", DisplayName: "kyat di Myanmar", Symbol: "MMK"}, "NIC": ut.Currency{Currency: "NIC", DisplayName: "cordoba nicaraguense", Symbol: ""}, "TJS": ut.Currency{Currency: "TJS", DisplayName: "somoni tagiko", Symbol: "TJS"}, "XDR": ut.Currency{Currency: "XDR", DisplayName: "diritti speciali di incasso", Symbol: ""}, "CUC": ut.Currency{Currency: "CUC", DisplayName: "peso cubano convertibile", Symbol: "CUC"}, "LYD": ut.Currency{Currency: "LYD", DisplayName: "dinaro libico", Symbol: "LYD"}, "TRY": ut.Currency{Currency: "TRY", DisplayName: "lira turca", Symbol: "TRY"}, "SLL": ut.Currency{Currency: "SLL", DisplayName: "leone della Sierra Leone", Symbol: "SLL"}, "XRE": ut.Currency{Currency: "XRE", DisplayName: "fondi RINET", Symbol: ""}, "ALL": ut.Currency{Currency: "ALL", DisplayName: "lek albanese", Symbol: "ALL"}, "BRN": ut.Currency{Currency: "BRN", DisplayName: "cruzado novo brasiliano", Symbol: ""}, "GHS": ut.Currency{Currency: "GHS", DisplayName: "cedi ghanese", Symbol: "GHS"}, "KES": ut.Currency{Currency: "KES", DisplayName: "scellino keniota", Symbol: "KES"}, "LKR": ut.Currency{Currency: "LKR", DisplayName: "rupia di Sri Lanka", Symbol: "LKR"}, "CNY": ut.Currency{Currency: "CNY", DisplayName: "renminbi cinese", Symbol: "CN¥"}, "CYP": ut.Currency{Currency: "CYP", DisplayName: "sterlina cipriota", Symbol: ""}, "KYD": ut.Currency{Currency: "KYD", DisplayName: "dollaro delle Isole Cayman", Symbol: "KYD"}, "PYG": ut.Currency{Currency: "PYG", DisplayName: "guaraní paraguayano", Symbol: "PYG"}, "XCD": ut.Currency{Currency: "XCD", DisplayName: "dollaro dei Caraibi orientali", Symbol: "EC$"}, "XFO": ut.Currency{Currency: "XFO", DisplayName: "franco oro francese", Symbol: ""}, "BEC": ut.Currency{Currency: "BEC", DisplayName: "franco belga (convertibile)", Symbol: ""}, "BGL": ut.Currency{Currency: "BGL", DisplayName: "lev bulgaro (1962–1999)", Symbol: ""}, "EEK": ut.Currency{Currency: "EEK", DisplayName: "corona dell’Estonia", Symbol: ""}, "JMD": ut.Currency{Currency: "JMD", DisplayName: "dollaro giamaicano", Symbol: "JMD"}, "MKD": ut.Currency{Currency: "MKD", DisplayName: "denar macedone", Symbol: "MKD"}, "CZK": ut.Currency{Currency: "CZK", DisplayName: "corona ceca", Symbol: "CZK"}, "MZN": ut.Currency{Currency: "MZN", DisplayName: "metical mozambicano", Symbol: "MZN"}, "ETB": ut.Currency{Currency: "ETB", DisplayName: "birr etiope", Symbol: "ETB"}, "PGK": ut.Currency{Currency: "PGK", DisplayName: "kina papuana", Symbol: "PGK"}, "USS": ut.Currency{Currency: "USS", DisplayName: "dollaro statunitense (same day)", Symbol: ""}, "AWG": ut.Currency{Currency: "AWG", DisplayName: "fiorino di Aruba", Symbol: "AWG"}, "BAM": ut.Currency{Currency: "BAM", DisplayName: "marco convertibile della Bosnia-Herzegovina", Symbol: "BAM"}, "BDT": ut.Currency{Currency: "BDT", DisplayName: "taka bangladese", Symbol: "BDT"}, "BWP": ut.Currency{Currency: "BWP", DisplayName: "pula del Botswana", Symbol: "BWP"}, "BYB": ut.Currency{Currency: "BYB", DisplayName: "nuovo rublo bielorusso (1994–1999)", Symbol: ""}, "BOP": ut.Currency{Currency: "BOP", DisplayName: "peso boliviano", Symbol: ""}, "EUR": ut.Currency{Currency: "EUR", DisplayName: "euro", Symbol: "€"}, "NAD": ut.Currency{Currency: "NAD", DisplayName: "dollaro namibiano", Symbol: "NAD"}, "NZD": ut.Currency{Currency: "NZD", DisplayName: "dollaro neozelandese", Symbol: "NZ$"}, "USN": ut.Currency{Currency: "USN", DisplayName: "dollaro statunitense (next day)", Symbol: ""}, "IRR": ut.Currency{Currency: "IRR", DisplayName: "rial iraniano", Symbol: "IRR"}, "SUR": ut.Currency{Currency: "SUR", DisplayName: "rublo sovietico", Symbol: ""}, "USD": ut.Currency{Currency: "USD", DisplayName: "dollaro statunitense", Symbol: "US$"}, "AON": ut.Currency{Currency: "AON", DisplayName: "nuovo kwanza angolano (1990–2000)", Symbol: ""}, "BRL": ut.Currency{Currency: "BRL", DisplayName: "real brasiliano", Symbol: "BRL"}, "BYR": ut.Currency{Currency: "BYR", DisplayName: "rublo bielorusso", Symbol: "BYR"}, "CVE": ut.Currency{Currency: "CVE", DisplayName: "escudo capoverdiano", Symbol: "CVE"}, "GQE": ut.Currency{Currency: "GQE", DisplayName: "ekwele della Guinea Equatoriale", Symbol: ""}, "XAU": ut.Currency{Currency: "XAU", DisplayName: "oro", Symbol: ""}, "XFU": ut.Currency{Currency: "XFU", DisplayName: "franco UIC francese", Symbol: ""}, "DEM": ut.Currency{Currency: "DEM", DisplayName: "marco tedesco", Symbol: ""}, "GWE": ut.Currency{Currency: "GWE", DisplayName: "escudo della Guinea portoghese", Symbol: ""}, "KHR": ut.Currency{Currency: "KHR", DisplayName: "riel cambogiano", Symbol: "KHR"}, "XPD": ut.Currency{Currency: "XPD", DisplayName: "palladio", Symbol: ""}, "AZN": ut.Currency{Currency: "AZN", DisplayName: "manat azero", Symbol: "AZN"}, "FRF": ut.Currency{Currency: "FRF", DisplayName: "franco francese", Symbol: ""}, "LUL": ut.Currency{Currency: "LUL", DisplayName: "franco finanziario del Lussemburgo", Symbol: ""}, "WST": ut.Currency{Currency: "WST", DisplayName: "tala samoano", Symbol: "WST"}, "SRG": ut.Currency{Currency: "SRG", DisplayName: "fiorino del Suriname", Symbol: ""}, "YUD": ut.Currency{Currency: "YUD", DisplayName: "dinaro forte yugoslavo", Symbol: ""}, "ADP": ut.Currency{Currency: "ADP", DisplayName: "peseta andorrana", Symbol: ""}, "AED": ut.Currency{Currency: "AED", DisplayName: "dirham degli Emirati Arabi Uniti", Symbol: "AED"}, "ARS": ut.Currency{Currency: "ARS", DisplayName: "peso argentino", Symbol: "ARS"}, "KRW": ut.Currency{Currency: "KRW", DisplayName: "won sudcoreano", Symbol: "KRW"}, "LUC": ut.Currency{Currency: "LUC", DisplayName: "franco convertibile del Lussemburgo", Symbol: ""}, "LAK": ut.Currency{Currency: "LAK", DisplayName: "kip laotiano", Symbol: "LAK"}, "CAD": ut.Currency{Currency: "CAD", DisplayName: "dollaro canadese", Symbol: "CA$"}, "GYD": ut.Currency{Currency: "GYD", DisplayName: "dollaro della Guyana", Symbol: "GYD"}, "KPW": ut.Currency{Currency: "KPW", DisplayName: "won nordcoreano", Symbol: "KPW"}, "TOP": ut.Currency{Currency: "TOP", DisplayName: "paʻanga tongano", Symbol: "TOP"}, "XAG": ut.Currency{Currency: "XAG", DisplayName: "argento", Symbol: ""}, "ERN": ut.Currency{Currency: "ERN", DisplayName: "nakfa eritreo", Symbol: "ERN"}, "LRD": ut.Currency{Currency: "LRD", DisplayName: "dollaro liberiano", Symbol: "LRD"}, "RON": ut.Currency{Currency: "RON", DisplayName: "leu rumeno", Symbol: "RON"}, "SSP": ut.Currency{Currency: "SSP", DisplayName: "sterlina sud-sudanese", Symbol: "SSP"}, "PEI": ut.Currency{Currency: "PEI", DisplayName: "inti peruviano", Symbol: ""}, "CSK": ut.Currency{Currency: "CSK", DisplayName: "corona forte cecoslovacca", Symbol: ""}, "BGN": ut.Currency{Currency: "BGN", DisplayName: "lev bulgaro", Symbol: "BGN"}, "BIF": ut.Currency{Currency: "BIF", DisplayName: "franco del Burundi", Symbol: "BIF"}, "KMF": ut.Currency{Currency: "KMF", DisplayName: "franco comoriano", Symbol: "KMF"}, "RHD": ut.Currency{Currency: "RHD", DisplayName: "dollaro della Rhodesia", Symbol: ""}, "UZS": ut.Currency{Currency: "UZS", DisplayName: "sum uzbeco", Symbol: "UZS"}, "UYP": ut.Currency{Currency: "UYP", DisplayName: "peso uruguaiano (1975–1993)", Symbol: ""}, "ILP": ut.Currency{Currency: "ILP", DisplayName: "sterlina israeliana", Symbol: ""}, "HNL": ut.Currency{Currency: "HNL", DisplayName: "lempira honduregna", Symbol: "HNL"}, "MAF": ut.Currency{Currency: "MAF", DisplayName: "franco marocchino", Symbol: ""}, "XBB": ut.Currency{Currency: "XBB", DisplayName: "unità monetaria europea", Symbol: ""}, "GNF": ut.Currency{Currency: "GNF", DisplayName: "franco della Guinea", Symbol: "GNF"}, "IQD": ut.Currency{Currency: "IQD", DisplayName: "dinaro iracheno", Symbol: "IQD"}, "SDD": ut.Currency{Currency: "SDD", DisplayName: "dinaro sudanese", Symbol: ""}, "UAK": ut.Currency{Currency: "UAK", DisplayName: "karbovanetz ucraino", Symbol: ""}, "GEL": ut.Currency{Currency: "GEL", DisplayName: "lari georgiano", Symbol: "GEL"}, "SOS": ut.Currency{Currency: "SOS", DisplayName: "scellino somalo", Symbol: "SOS"}, "MXN": ut.Currency{Currency: "MXN", DisplayName: "peso messicano", Symbol: "MXN"}, "ANG": ut.Currency{Currency: "ANG", DisplayName: "fiorino delle Antille olandesi", Symbol: "ANG"}, "BRR": ut.Currency{Currency: "BRR", DisplayName: "cruzeiro brasiliano", Symbol: ""}, "XPF": ut.Currency{Currency: "XPF", DisplayName: "franco CFP", Symbol: "CFPF"}, "ARP": ut.Currency{Currency: "ARP", DisplayName: "peso argentino (vecchio Cod.)", Symbol: ""}, "BND": ut.Currency{Currency: "BND", DisplayName: "dollaro del Brunei", Symbol: "BND"}, "FIM": ut.Currency{Currency: "FIM", DisplayName: "markka finlandese", Symbol: ""}, "GNS": ut.Currency{Currency: "GNS", DisplayName: "syli della Guinea", Symbol: ""}, "SKK": ut.Currency{Currency: "SKK", DisplayName: "corona slovacca", Symbol: ""}, "HTG": ut.Currency{Currency: "HTG", DisplayName: "gourde haitiano", Symbol: "HTG"}, "LUF": ut.Currency{Currency: "LUF", DisplayName: "franco del Lussemburgo", Symbol: ""}, "SGD": ut.Currency{Currency: "SGD", DisplayName: "dollaro di Singapore", Symbol: "SGD"}, "SIT": ut.Currency{Currency: "SIT", DisplayName: "tallero sloveno", Symbol: ""}, "THB": ut.Currency{Currency: "THB", DisplayName: "baht thailandese", Symbol: "฿"}, "AOA": ut.Currency{Currency: "AOA", DisplayName: "kwanza angolano", Symbol: "AOA"}, "AZM": ut.Currency{Currency: "AZM", DisplayName: "manat azero (1993–2006)", Symbol: ""}, "CHF": ut.Currency{Currency: "CHF", DisplayName: "franco svizzero", Symbol: "CHF"}, "MTP": ut.Currency{Currency: "MTP", DisplayName: "sterlina maltese", Symbol: ""}, "NPR": ut.Currency{Currency: "NPR", DisplayName: "rupia nepalese", Symbol: "NPR"}, "TZS": ut.Currency{Currency: "TZS", DisplayName: "scellino della Tanzania", Symbol: "TZS"}, "PES": ut.Currency{Currency: "PES", DisplayName: "sol peruviano", Symbol: ""}, "TTD": ut.Currency{Currency: "TTD", DisplayName: "dollaro di Trinidad e Tobago", Symbol: "TTD"}, "ATS": ut.Currency{Currency: "ATS", DisplayName: "scellino austriaco", Symbol: ""}, "BUK": ut.Currency{Currency: "BUK", DisplayName: "kyat birmano", Symbol: ""}, "SVC": ut.Currency{Currency: "SVC", DisplayName: "colón salvadoregno", Symbol: ""}, "ZWD": ut.Currency{Currency: "ZWD", DisplayName: "dollaro dello Zimbabwe", Symbol: ""}, "XBD": ut.Currency{Currency: "XBD", DisplayName: "unità di acconto europea (XBD)", Symbol: ""}, "BOB": ut.Currency{Currency: "BOB", DisplayName: "boliviano", Symbol: "BOB"}, "COP": ut.Currency{Currency: "COP", DisplayName: "peso colombiano", Symbol: "COP"}, "LVR": ut.Currency{Currency: "LVR", DisplayName: "rublo lettone", Symbol: ""}, "MOP": ut.Currency{Currency: "MOP", DisplayName: "pataca di Macao", Symbol: "MOP"}, "TJR": ut.Currency{Currency: "TJR", DisplayName: "rublo del Tajikistan", Symbol: ""}, "AOR": ut.Currency{Currency: "AOR", DisplayName: "kwanza reajustado angolano (1995–1999)", Symbol: ""}, "AUD": ut.Currency{Currency: "AUD", DisplayName: "dollaro australiano", Symbol: "A$"}, "JPY": ut.Currency{Currency: "JPY", DisplayName: "yen giapponese", Symbol: "JPY"}, "KZT": ut.Currency{Currency: "KZT", DisplayName: "tenge kazako", Symbol: "KZT"}, "BZD": ut.Currency{Currency: "BZD", DisplayName: "dollaro del Belize", Symbol: "BZD"}, "LSL": ut.Currency{Currency: "LSL", DisplayName: "loti del Lesotho", Symbol: ""}, "TPE": ut.Currency{Currency: "TPE", DisplayName: "escudo di Timor", Symbol: ""}, "ZRN": ut.Currency{Currency: "ZRN", DisplayName: "nuovo zaire dello Zaire", Symbol: ""}, "ECV": ut.Currency{Currency: "ECV", DisplayName: "unidad de valor constante (UVC) dell’Ecuador", Symbol: ""}, "OMR": ut.Currency{Currency: "OMR", DisplayName: "rial omanita", Symbol: "OMR"}, "ARA": ut.Currency{Currency: "ARA", DisplayName: "austral argentino", Symbol: ""}, "GBP": ut.Currency{Currency: "GBP", DisplayName: "sterlina britannica", Symbol: "£"}, "TND": ut.Currency{Currency: "TND", DisplayName: "dinaro tunisino", Symbol: "TND"}, "VUV": ut.Currency{Currency: "VUV", DisplayName: "vatu di Vanuatu", Symbol: "VUV"}, "CLF": ut.Currency{Currency: "CLF", DisplayName: "unidades de fomento chilene", Symbol: ""}, "JOD": ut.Currency{Currency: "JOD", DisplayName: "dinaro giordano", Symbol: "JOD"}, "YUN": ut.Currency{Currency: "YUN", DisplayName: "dinaro convertibile yugoslavo", Symbol: ""}, "AMD": ut.Currency{Currency: "AMD", DisplayName: "dram armeno", Symbol: "AMD"}, "BEF": ut.Currency{Currency: "BEF", DisplayName: "franco belga", Symbol: ""}, "DOP": ut.Currency{Currency: "DOP", DisplayName: "peso dominicano", Symbol: "DOP"}, "GHC": ut.Currency{Currency: "GHC", DisplayName: "cedi del Ghana", Symbol: ""}, "MAD": ut.Currency{Currency: "MAD", DisplayName: "dirham marocchino", Symbol: "MAD"}, "LTL": ut.Currency{Currency: "LTL", DisplayName: "litas lituano", Symbol: "LTL"}, "PEN": ut.Currency{Currency: "PEN", DisplayName: "nuovo sol peruviano", Symbol: "PEN"}, "RWF": ut.Currency{Currency: "RWF", DisplayName: "franco ruandese", Symbol: "RWF"}, "UYU": ut.Currency{Currency: "UYU", DisplayName: "peso uruguayano", Symbol: "UYU"}, "VND": ut.Currency{Currency: "VND", DisplayName: "dong vietnamita", Symbol: "₫"}, "BBD": ut.Currency{Currency: "BBD", DisplayName: "dollaro di Barbados", Symbol: "BBD"}, "BHD": ut.Currency{Currency: "BHD", DisplayName: "dinaro del Bahrein", Symbol: "BHD"}, "ESB": ut.Currency{Currency: "ESB", DisplayName: "peseta spagnola account convertibile", Symbol: ""}, "HUF": ut.Currency{Currency: "HUF", DisplayName: "fiorino ungherese", Symbol: "HUF"}, "NOK": ut.Currency{Currency: "NOK", DisplayName: "corona norvegese", Symbol: "NOK"}, "AFN": ut.Currency{Currency: "AFN", DisplayName: "afghani", Symbol: "AFN"}, "GTQ": ut.Currency{Currency: "GTQ", DisplayName: "quetzal guatemalteco", Symbol: "GTQ"}, "LTT": ut.Currency{Currency: "LTT", DisplayName: "talonas lituani", Symbol: ""}, "NLG": ut.Currency{Currency: "NLG", DisplayName: "fiorino olandese", Symbol: ""}, "QAR": ut.Currency{Currency: "QAR", DisplayName: "rial qatariano", Symbol: "QAR"}, "RSD": ut.Currency{Currency: "RSD", DisplayName: "dinaro serbo", Symbol: "RSD"}, "SEK": ut.Currency{Currency: "SEK", DisplayName: "corona svedese", Symbol: "SEK"}, "TWD": ut.Currency{Currency: "TWD", DisplayName: "nuovo dollaro taiwanese", Symbol: "TWD"}, "VEB": ut.Currency{Currency: "VEB", DisplayName: "bolivar venezuelano (1871–2008)", Symbol: ""}}