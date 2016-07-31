package ro

import "github.com/go-playground/universal-translator"

var currencies = ut.CurrencyFormatValue{"FIM": ut.Currency{Currency: "FIM", DisplayName: "marcă finlandeză", Symbol: ""}, "LUC": ut.Currency{Currency: "LUC", DisplayName: "franc convertibil luxemburghez", Symbol: ""}, "TMM": ut.Currency{Currency: "TMM", DisplayName: "manat turkmen (1993–2009)", Symbol: ""}, "XTS": ut.Currency{Currency: "XTS", DisplayName: "cod monetar de test", Symbol: ""}, "EEK": ut.Currency{Currency: "EEK", DisplayName: "coroană estoniană", Symbol: ""}, "ROL": ut.Currency{Currency: "ROL", DisplayName: "leu românesc (1952–2006)", Symbol: ""}, "TZS": ut.Currency{Currency: "TZS", DisplayName: "șiling tanzanian", Symbol: "TZS"}, "GTQ": ut.Currency{Currency: "GTQ", DisplayName: "quetzal guatemalez", Symbol: "GTQ"}, "MOP": ut.Currency{Currency: "MOP", DisplayName: "pataca din Macao", Symbol: "MOP"}, "MYR": ut.Currency{Currency: "MYR", DisplayName: "ringgit malaiezian", Symbol: "MYR"}, "NZD": ut.Currency{Currency: "NZD", DisplayName: "dolar neozeelandez", Symbol: "NZD"}, "SOS": ut.Currency{Currency: "SOS", DisplayName: "șiling somalez", Symbol: "SOS"}, "BOP": ut.Currency{Currency: "BOP", DisplayName: "peso bolivian", Symbol: ""}, "CUP": ut.Currency{Currency: "CUP", DisplayName: "peso cubanez", Symbol: "CUP"}, "FJD": ut.Currency{Currency: "FJD", DisplayName: "dolar fijian", Symbol: "FJD"}, "XDR": ut.Currency{Currency: "XDR", DisplayName: "drepturi speciale de tragere", Symbol: ""}, "XPT": ut.Currency{Currency: "XPT", DisplayName: "platină", Symbol: ""}, "TRL": ut.Currency{Currency: "TRL", DisplayName: "liră turcească (1922–2005)", Symbol: ""}, "BRR": ut.Currency{Currency: "BRR", DisplayName: "cruzeiro brazilian (1993–1994)", Symbol: ""}, "INR": ut.Currency{Currency: "INR", DisplayName: "rupie indiană", Symbol: "INR"}, "STD": ut.Currency{Currency: "STD", DisplayName: "dobra Sao Tome și Principe", Symbol: "STD"}, "LSL": ut.Currency{Currency: "LSL", DisplayName: "loti lesothian", Symbol: ""}, "ARS": ut.Currency{Currency: "ARS", DisplayName: "peso argentinian", Symbol: "ARS"}, "BOB": ut.Currency{Currency: "BOB", DisplayName: "boliviano bolivian", Symbol: "BOB"}, "ISK": ut.Currency{Currency: "ISK", DisplayName: "coroană islandeză", Symbol: "ISK"}, "GMD": ut.Currency{Currency: "GMD", DisplayName: "dalasi din Gambia", Symbol: "GMD"}, "JPY": ut.Currency{Currency: "JPY", DisplayName: "yen japonez", Symbol: "JPY"}, "CAD": ut.Currency{Currency: "CAD", DisplayName: "dolar canadian", Symbol: "CAD"}, "CZK": ut.Currency{Currency: "CZK", DisplayName: "coroană cehă", Symbol: "CZK"}, "GHS": ut.Currency{Currency: "GHS", DisplayName: "cedi ghanez", Symbol: "GHS"}, "ESB": ut.Currency{Currency: "ESB", DisplayName: "peseta spaniolă (cont convertibil)", Symbol: ""}, "XPF": ut.Currency{Currency: "XPF", DisplayName: "franc CFP", Symbol: "CFPF"}, "KZT": ut.Currency{Currency: "KZT", DisplayName: "tenge kazahă", Symbol: "KZT"}, "ZMW": ut.Currency{Currency: "ZMW", DisplayName: "kwacha zambian", Symbol: "ZMW"}, "ESP": ut.Currency{Currency: "ESP", DisplayName: "pesetă spaniolă", Symbol: ""}, "NAD": ut.Currency{Currency: "NAD", DisplayName: "dolar namibian", Symbol: "NAD"}, "SCR": ut.Currency{Currency: "SCR", DisplayName: "rupie din Seychelles", Symbol: "SCR"}, "SUR": ut.Currency{Currency: "SUR", DisplayName: "rublă sovietică", Symbol: ""}, "VEF": ut.Currency{Currency: "VEF", DisplayName: "bolivar venezuelean", Symbol: "VEF"}, "ARP": ut.Currency{Currency: "ARP", DisplayName: "peso argentinian (1983–1985)", Symbol: ""}, "BBD": ut.Currency{Currency: "BBD", DisplayName: "dolar Barbados", Symbol: "BBD"}, "BND": ut.Currency{Currency: "BND", DisplayName: "dolar din Brunei", Symbol: "BND"}, "LAK": ut.Currency{Currency: "LAK", DisplayName: "kip laoțian", Symbol: "LAK"}, "LRD": ut.Currency{Currency: "LRD", DisplayName: "dolar liberian", Symbol: "LRD"}, "MNT": ut.Currency{Currency: "MNT", DisplayName: "tugrik mongol", Symbol: "MNT"}, "PAB": ut.Currency{Currency: "PAB", DisplayName: "balboa panameză", Symbol: "PAB"}, "UGX": ut.Currency{Currency: "UGX", DisplayName: "șiling ugandez", Symbol: "UGX"}, "BIF": ut.Currency{Currency: "BIF", DisplayName: "franc burundez", Symbol: "BIF"}, "GIP": ut.Currency{Currency: "GIP", DisplayName: "liră din Gibraltar", Symbol: "GIP"}, "KRW": ut.Currency{Currency: "KRW", DisplayName: "won sud-coreean", Symbol: "KRW"}, "KYD": ut.Currency{Currency: "KYD", DisplayName: "dolar din Insulele Cayman", Symbol: "KYD"}, "UAK": ut.Currency{Currency: "UAK", DisplayName: "carboavă ucraineană", Symbol: ""}, "HTG": ut.Currency{Currency: "HTG", DisplayName: "gourde din Haiti", Symbol: "HTG"}, "NIO": ut.Currency{Currency: "NIO", DisplayName: "cordoba nicaraguană", Symbol: "NIO"}, "MAF": ut.Currency{Currency: "MAF", DisplayName: "franc marocan", Symbol: ""}, "RSD": ut.Currency{Currency: "RSD", DisplayName: "dinar sârbesc", Symbol: "RSD"}, "SLL": ut.Currency{Currency: "SLL", DisplayName: "leone din Sierra Leone", Symbol: "SLL"}, "AWG": ut.Currency{Currency: "AWG", DisplayName: "florin aruban", Symbol: "AWG"}, "BDT": ut.Currency{Currency: "BDT", DisplayName: "taka din Bangladesh", Symbol: "BDT"}, "KHR": ut.Currency{Currency: "KHR", DisplayName: "riel cambodgian", Symbol: "KHR"}, "GWP": ut.Currency{Currency: "GWP", DisplayName: "peso Guineea-Bissau", Symbol: ""}, "KWD": ut.Currency{Currency: "KWD", DisplayName: "dinar kuweitian", Symbol: "KWD"}, "YER": ut.Currency{Currency: "YER", DisplayName: "rial yemenit", Symbol: "YER"}, "TOP": ut.Currency{Currency: "TOP", DisplayName: "pa’anga tongană", Symbol: "TOP"}, "USD": ut.Currency{Currency: "USD", DisplayName: "dolar american", Symbol: "USD"}, "BEL": ut.Currency{Currency: "BEL", DisplayName: "franc belgian (financiar)", Symbol: ""}, "ITL": ut.Currency{Currency: "ITL", DisplayName: "liră italiană", Symbol: ""}, "JMD": ut.Currency{Currency: "JMD", DisplayName: "dolar jamaican", Symbol: "JMD"}, "MXP": ut.Currency{Currency: "MXP", DisplayName: "peso mexican de argint (1861–1992)", Symbol: ""}, "MZN": ut.Currency{Currency: "MZN", DisplayName: "metical mozambican", Symbol: "MZN"}, "SAR": ut.Currency{Currency: "SAR", DisplayName: "rial saudit", Symbol: "SAR"}, "UYP": ut.Currency{Currency: "UYP", DisplayName: "peso Uruguay (1975–1993)", Symbol: ""}, "CHF": ut.Currency{Currency: "CHF", DisplayName: "franc elvețian", Symbol: "CHF"}, "CYP": ut.Currency{Currency: "CYP", DisplayName: "liră cipriotă", Symbol: ""}, "IEP": ut.Currency{Currency: "IEP", DisplayName: "liră irlandeză", Symbol: ""}, "RHD": ut.Currency{Currency: "RHD", DisplayName: "dolar rhodesian", Symbol: ""}, "AMD": ut.Currency{Currency: "AMD", DisplayName: "dram armenesc", Symbol: "AMD"}, "BEC": ut.Currency{Currency: "BEC", DisplayName: "franc belgian (convertibil)", Symbol: ""}, "NLG": ut.Currency{Currency: "NLG", DisplayName: "gulden olandez", Symbol: ""}, "GHC": ut.Currency{Currency: "GHC", DisplayName: "cedi Ghana (1979–2007)", Symbol: ""}, "PKR": ut.Currency{Currency: "PKR", DisplayName: "rupie pakistaneză", Symbol: "PKR"}, "QAR": ut.Currency{Currency: "QAR", DisplayName: "rial qatarian", Symbol: "QAR"}, "XBB": ut.Currency{Currency: "XBB", DisplayName: "unitate monetară europeană", Symbol: ""}, "BGN": ut.Currency{Currency: "BGN", DisplayName: "leva bulgărească", Symbol: "BGN"}, "COP": ut.Currency{Currency: "COP", DisplayName: "peso columbian", Symbol: "COP"}, "DDM": ut.Currency{Currency: "DDM", DisplayName: "marcă est-germană", Symbol: ""}, "AFN": ut.Currency{Currency: "AFN", DisplayName: "afgani afgan", Symbol: "AFN"}, "HUF": ut.Currency{Currency: "HUF", DisplayName: "forint maghiar", Symbol: "HUF"}, "SGD": ut.Currency{Currency: "SGD", DisplayName: "dolar Singapore", Symbol: "SGD"}, "ZAL": ut.Currency{Currency: "ZAL", DisplayName: "rand sud-african (financiar)", Symbol: ""}, "PLN": ut.Currency{Currency: "PLN", DisplayName: "zlot polonez", Symbol: "PLN"}, "XCD": ut.Currency{Currency: "XCD", DisplayName: "dolar din Caraibele de Est", Symbol: "XCD"}, "ZWL": ut.Currency{Currency: "ZWL", DisplayName: "dolar Zimbabwe (2009)", Symbol: ""}, "AOA": ut.Currency{Currency: "AOA", DisplayName: "kwanza angoleză", Symbol: "AOA"}, "BAM": ut.Currency{Currency: "BAM", DisplayName: "marcă convertibilă din Bosnia și Herțegovina", Symbol: "BAM"}, "MLF": ut.Currency{Currency: "MLF", DisplayName: "franc Mali", Symbol: ""}, "BMD": ut.Currency{Currency: "BMD", DisplayName: "dolar din Bermuda", Symbol: "BMD"}, "LBP": ut.Currency{Currency: "LBP", DisplayName: "liră libaneză", Symbol: "LBP"}, "XOF": ut.Currency{Currency: "XOF", DisplayName: "franc CFA BCEAO", Symbol: "CFA"}, "ADP": ut.Currency{Currency: "ADP", DisplayName: "pesetă andorrană", Symbol: ""}, "RWF": ut.Currency{Currency: "RWF", DisplayName: "franc rwandez", Symbol: "RWF"}, "XBA": ut.Currency{Currency: "XBA", DisplayName: "unitate compusă europeană", Symbol: ""}, "ECS": ut.Currency{Currency: "ECS", DisplayName: "sucre Ecuador", Symbol: ""}, "ESA": ut.Currency{Currency: "ESA", DisplayName: "peseta spaniolă (cont A)", Symbol: ""}, "TMT": ut.Currency{Currency: "TMT", DisplayName: "manat turkmen", Symbol: "TMT"}, "VUV": ut.Currency{Currency: "VUV", DisplayName: "vatu din Vanuatu", Symbol: "VUV"}, "YUM": ut.Currency{Currency: "YUM", DisplayName: "dinar iugoslav nou", Symbol: ""}, "CDF": ut.Currency{Currency: "CDF", DisplayName: "franc congolez", Symbol: "CDF"}, "ERN": ut.Currency{Currency: "ERN", DisplayName: "nakfa eritreeană", Symbol: "ERN"}, "FRF": ut.Currency{Currency: "FRF", DisplayName: "franc francez", Symbol: ""}, "YUN": ut.Currency{Currency: "YUN", DisplayName: "dinar iugoslav convertibil", Symbol: ""}, "DZD": ut.Currency{Currency: "DZD", DisplayName: "dinar algerian", Symbol: "DZD"}, "MGF": ut.Currency{Currency: "MGF", DisplayName: "franc Madagascar", Symbol: ""}, "NGN": ut.Currency{Currency: "NGN", DisplayName: "naira nigeriană", Symbol: "NGN"}, "LUL": ut.Currency{Currency: "LUL", DisplayName: "franc financiar luxemburghez", Symbol: ""}, "RUB": ut.Currency{Currency: "RUB", DisplayName: "rublă rusească", Symbol: "RUB"}, "ATS": ut.Currency{Currency: "ATS", DisplayName: "șiling austriac", Symbol: ""}, "PEN": ut.Currency{Currency: "PEN", DisplayName: "sol nou peruvian", Symbol: "PEN"}, "SZL": ut.Currency{Currency: "SZL", DisplayName: "lilangeni din Swaziland", Symbol: "SZL"}, "TTD": ut.Currency{Currency: "TTD", DisplayName: "dolar din Trinidad-Tobago", Symbol: "TTD"}, "BOV": ut.Currency{Currency: "BOV", DisplayName: "mvdol bolivian", Symbol: ""}, "PES": ut.Currency{Currency: "PES", DisplayName: "sol peruvian (1863–1965)", Symbol: ""}, "SEK": ut.Currency{Currency: "SEK", DisplayName: "coroană suedeză", Symbol: "SEK"}, "LUF": ut.Currency{Currency: "LUF", DisplayName: "franc luxemburghez", Symbol: ""}, "MVR": ut.Currency{Currency: "MVR", DisplayName: "rufiyaa maldiviană", Symbol: "MVR"}, "MXN": ut.Currency{Currency: "MXN", DisplayName: "peso mexican", Symbol: "MXN"}, "SDG": ut.Currency{Currency: "SDG", DisplayName: "liră sudaneză", Symbol: "SDG"}, "TJR": ut.Currency{Currency: "TJR", DisplayName: "rublă Tadjikistan", Symbol: ""}, "USN": ut.Currency{Currency: "USN", DisplayName: "dolar american (ziua următoare)", Symbol: ""}, "ZMK": ut.Currency{Currency: "ZMK", DisplayName: "kwacha zambian (1968–2012)", Symbol: ""}, "GBP": ut.Currency{Currency: "GBP", DisplayName: "liră sterlină", Symbol: "GBP"}, "GRD": ut.Currency{Currency: "GRD", DisplayName: "drahmă grecească", Symbol: ""}, "LTL": ut.Currency{Currency: "LTL", DisplayName: "litu lituanian", Symbol: "LTL"}, "IDR": ut.Currency{Currency: "IDR", DisplayName: "rupie indoneziană", Symbol: "IDR"}, "SHP": ut.Currency{Currency: "SHP", DisplayName: "liră Insula Sf. Elena", Symbol: "SHP"}, "SIT": ut.Currency{Currency: "SIT", DisplayName: "tolar sloven", Symbol: ""}, "XPD": ut.Currency{Currency: "XPD", DisplayName: "paladiu", Symbol: ""}, "NIC": ut.Currency{Currency: "NIC", DisplayName: "cordoba nicaraguană (1988–1991)", Symbol: ""}, "UGS": ut.Currency{Currency: "UGS", DisplayName: "șiling ugandez (1966–1987)", Symbol: ""}, "VND": ut.Currency{Currency: "VND", DisplayName: "dong vietnamez", Symbol: "VND"}, "XFU": ut.Currency{Currency: "XFU", DisplayName: "franc UIC francez", Symbol: ""}, "EGP": ut.Currency{Currency: "EGP", DisplayName: "liră egipteană", Symbol: "EGP"}, "NOK": ut.Currency{Currency: "NOK", DisplayName: "coroană norvegiană", Symbol: "NOK"}, "SYP": ut.Currency{Currency: "SYP", DisplayName: "liră siriană", Symbol: "SYP"}, "NPR": ut.Currency{Currency: "NPR", DisplayName: "rupie nepaleză", Symbol: "NPR"}, "BSD": ut.Currency{Currency: "BSD", DisplayName: "dolar din Bahamas", Symbol: "BSD"}, "FKP": ut.Currency{Currency: "FKP", DisplayName: "liră din Insulele Falkland", Symbol: "FKP"}, "MDL": ut.Currency{Currency: "MDL", DisplayName: "leu moldovenesc", Symbol: "MDL"}, "SSP": ut.Currency{Currency: "SSP", DisplayName: "liră sud-sudaneză", Symbol: "SSP"}, "VEB": ut.Currency{Currency: "VEB", DisplayName: "bolivar Venezuela (1871–2008)", Symbol: ""}, "PGK": ut.Currency{Currency: "PGK", DisplayName: "kina din Papua-Noua Guinee", Symbol: "PGK"}, "RON": ut.Currency{Currency: "RON", DisplayName: "leu românesc", Symbol: "RON"}, "SKK": ut.Currency{Currency: "SKK", DisplayName: "coroană slovacă", Symbol: ""}, "USS": ut.Currency{Currency: "USS", DisplayName: "dolar american (aceeași zi)", Symbol: ""}, "KMF": ut.Currency{Currency: "KMF", DisplayName: "franc comorian", Symbol: "KMF"}, "MRO": ut.Currency{Currency: "MRO", DisplayName: "ouguiya mauritană", Symbol: "MRO"}, "PLZ": ut.Currency{Currency: "PLZ", DisplayName: "zlot polonez (1950–1995)", Symbol: ""}, "HNL": ut.Currency{Currency: "HNL", DisplayName: "lempira honduriană", Symbol: "HNL"}, "ILS": ut.Currency{Currency: "ILS", DisplayName: "șechel israelian nou", Symbol: "ILS"}, "SDP": ut.Currency{Currency: "SDP", DisplayName: "liră sudaneză (1957–1998)", Symbol: ""}, "PHP": ut.Currency{Currency: "PHP", DisplayName: "peso filipinez", Symbol: "PHP"}, "PYG": ut.Currency{Currency: "PYG", DisplayName: "guarani paraguayan", Symbol: "PYG"}, "SBD": ut.Currency{Currency: "SBD", DisplayName: "dolar din Insulele Solomon", Symbol: "SBD"}, "XAG": ut.Currency{Currency: "XAG", DisplayName: "argint", Symbol: ""}, "XFO": ut.Currency{Currency: "XFO", DisplayName: "franc francez de aur", Symbol: ""}, "BTN": ut.Currency{Currency: "BTN", DisplayName: "ngultrum din Bhutan", Symbol: "BTN"}, "CUC": ut.Currency{Currency: "CUC", DisplayName: "peso cubanez convertibil", Symbol: "CUC"}, "DEM": ut.Currency{Currency: "DEM", DisplayName: "marcă germană", Symbol: ""}, "SRG": ut.Currency{Currency: "SRG", DisplayName: "gulden Surinam", Symbol: ""}, "ZWD": ut.Currency{Currency: "ZWD", DisplayName: "dolar Zimbabwe (1980–2008)", Symbol: ""}, "GNF": ut.Currency{Currency: "GNF", DisplayName: "franc guineean", Symbol: "GNF"}, "UZS": ut.Currency{Currency: "UZS", DisplayName: "sum Uzbekistan", Symbol: "UZS"}, "XEU": ut.Currency{Currency: "XEU", DisplayName: "unitate de monedă europeană", Symbol: ""}, "CVE": ut.Currency{Currency: "CVE", DisplayName: "escudo din Capul Verde", Symbol: "CVE"}, "MZM": ut.Currency{Currency: "MZM", DisplayName: "metical Mozambic vechi", Symbol: ""}, "TWD": ut.Currency{Currency: "TWD", DisplayName: "dolar nou din Taiwan", Symbol: "TWD"}, "PEI": ut.Currency{Currency: "PEI", DisplayName: "inti peruvian", Symbol: ""}, "BAD": ut.Currency{Currency: "BAD", DisplayName: "dinar Bosnia-Herțegovina (1992–1994)", Symbol: ""}, "CLP": ut.Currency{Currency: "CLP", DisplayName: "peso chilian", Symbol: "CLP"}, "CNY": ut.Currency{Currency: "CNY", DisplayName: "yuan chinezesc", Symbol: "CNY"}, "BEF": ut.Currency{Currency: "BEF", DisplayName: "franc belgian", Symbol: ""}, "CSD": ut.Currency{Currency: "CSD", DisplayName: "dinar Serbia și Muntenegru (2002–2006)", Symbol: ""}, "MMK": ut.Currency{Currency: "MMK", DisplayName: "kyat din Myanmar", Symbol: "MMK"}, "HRK": ut.Currency{Currency: "HRK", DisplayName: "kuna croată", Symbol: "HRK"}, "XBD": ut.Currency{Currency: "XBD", DisplayName: "unitate de cont europeană (XBD)", Symbol: ""}, "ANG": ut.Currency{Currency: "ANG", DisplayName: "gulden din Antilele Olandeze", Symbol: "ANG"}, "BRE": ut.Currency{Currency: "BRE", DisplayName: "cruzeiro brazilian (1990–1993)", Symbol: ""}, "HRD": ut.Currency{Currency: "HRD", DisplayName: "dinar croat", Symbol: ""}, "OMR": ut.Currency{Currency: "OMR", DisplayName: "rial omanez", Symbol: "OMR"}, "DKK": ut.Currency{Currency: "DKK", DisplayName: "coroană daneză", Symbol: "DKK"}, "EUR": ut.Currency{Currency: "EUR", DisplayName: "euro", Symbol: "EUR"}, "JOD": ut.Currency{Currency: "JOD", DisplayName: "dinar iordanian", Symbol: "JOD"}, "AZM": ut.Currency{Currency: "AZM", DisplayName: "manat azer (1993–2006)", Symbol: ""}, "KPW": ut.Currency{Currency: "KPW", DisplayName: "won nord-coreean", Symbol: "KPW"}, "LVL": ut.Currency{Currency: "LVL", DisplayName: "lats letonian", Symbol: "LVL"}, "AED": ut.Currency{Currency: "AED", DisplayName: "dirham din Emiratele Arabe Unite", Symbol: "AED"}, "ALL": ut.Currency{Currency: "ALL", DisplayName: "leka albaneză", Symbol: "ALL"}, "WST": ut.Currency{Currency: "WST", DisplayName: "tala samoană", Symbol: "WST"}, "XXX": ut.Currency{Currency: "XXX", DisplayName: "monedă necunoscută", Symbol: ""}, "XAU": ut.Currency{Currency: "XAU", DisplayName: "aur", Symbol: ""}, "ZWR": ut.Currency{Currency: "ZWR", DisplayName: "dolar Zimbabwe (2008)", Symbol: ""}, "BUK": ut.Currency{Currency: "BUK", DisplayName: "kyat birman", Symbol: ""}, "ETB": ut.Currency{Currency: "ETB", DisplayName: "birr etiopian", Symbol: "ETB"}, "GEL": ut.Currency{Currency: "GEL", DisplayName: "lari georgian", Symbol: "GEL"}, "THB": ut.Currency{Currency: "THB", DisplayName: "baht thailandez", Symbol: "THB"}, "TND": ut.Currency{Currency: "TND", DisplayName: "dinar tunisian", Symbol: "TND"}, "BYR": ut.Currency{Currency: "BYR", DisplayName: "rublă belarusă", Symbol: "BYR"}, "LYD": ut.Currency{Currency: "LYD", DisplayName: "dinar libian", Symbol: "LYD"}, "SRD": ut.Currency{Currency: "SRD", DisplayName: "dolar surinamez", Symbol: "SRD"}, "HKD": ut.Currency{Currency: "HKD", DisplayName: "dolar din Hong Kong", Symbol: "HKD"}, "LVR": ut.Currency{Currency: "LVR", DisplayName: "rublă Letonia", Symbol: ""}, "YDD": ut.Currency{Currency: "YDD", DisplayName: "dinar Yemen", Symbol: ""}, "AZN": ut.Currency{Currency: "AZN", DisplayName: "manat azer", Symbol: "AZN"}, "BWP": ut.Currency{Currency: "BWP", DisplayName: "pula Botswana", Symbol: "BWP"}, "GYD": ut.Currency{Currency: "GYD", DisplayName: "dolar guyanez", Symbol: "GYD"}, "IQD": ut.Currency{Currency: "IQD", DisplayName: "dinar irakian", Symbol: "IQD"}, "MKD": ut.Currency{Currency: "MKD", DisplayName: "dinar macedonean", Symbol: "MKD"}, "UYU": ut.Currency{Currency: "UYU", DisplayName: "peso uruguayan", Symbol: "UYU"}, "DJF": ut.Currency{Currency: "DJF", DisplayName: "franc djiboutian", Symbol: "DJF"}, "MTL": ut.Currency{Currency: "MTL", DisplayName: "liră malteză", Symbol: ""}, "MWK": ut.Currency{Currency: "MWK", DisplayName: "kwacha malawiană", Symbol: "MWK"}, "KES": ut.Currency{Currency: "KES", DisplayName: "șiling kenyan", Symbol: "KES"}, "YUD": ut.Currency{Currency: "YUD", DisplayName: "dinar iugoslav greu", Symbol: ""}, "SDD": ut.Currency{Currency: "SDD", DisplayName: "dinar sudanez", Symbol: ""}, "ZAR": ut.Currency{Currency: "ZAR", DisplayName: "rand sud-african", Symbol: "ZAR"}, "LKR": ut.Currency{Currency: "LKR", DisplayName: "rupie din Sri Lanka", Symbol: "LKR"}, "MAD": ut.Currency{Currency: "MAD", DisplayName: "dirham marocan", Symbol: "MAD"}, "MZE": ut.Currency{Currency: "MZE", DisplayName: "escudo Mozambic", Symbol: ""}, "TJS": ut.Currency{Currency: "TJS", DisplayName: "somoni tadjic", Symbol: "TJS"}, "BRL": ut.Currency{Currency: "BRL", DisplayName: "real brazilian", Symbol: "BRL"}, "BZD": ut.Currency{Currency: "BZD", DisplayName: "dolar din Belize", Symbol: "BZD"}, "CRC": ut.Currency{Currency: "CRC", DisplayName: "colon costarican", Symbol: "CRC"}, "DOP": ut.Currency{Currency: "DOP", DisplayName: "peso dominican", Symbol: "DOP"}, "UAH": ut.Currency{Currency: "UAH", DisplayName: "hryvna ucraineană", Symbol: "UAH"}, "XAF": ut.Currency{Currency: "XAF", DisplayName: "franc CFA BEAC", Symbol: "FCFA"}, "XBC": ut.Currency{Currency: "XBC", DisplayName: "unitate de cont europeană (XBC)", Symbol: ""}, "ZRN": ut.Currency{Currency: "ZRN", DisplayName: "zair nou", Symbol: ""}, "AUD": ut.Currency{Currency: "AUD", DisplayName: "dolar australian", Symbol: "AUD"}, "BHD": ut.Currency{Currency: "BHD", DisplayName: "dinar din Bahrain", Symbol: "BHD"}, "MGA": ut.Currency{Currency: "MGA", DisplayName: "ariary malgaș", Symbol: "MGA"}, "MUR": ut.Currency{Currency: "MUR", DisplayName: "rupie mauritiană", Symbol: "MUR"}, "SVC": ut.Currency{Currency: "SVC", DisplayName: "colon El Salvador", Symbol: ""}, "TRY": ut.Currency{Currency: "TRY", DisplayName: "liră turcească", Symbol: "TRY"}, "ILP": ut.Currency{Currency: "ILP", DisplayName: "liră israeliană", Symbol: ""}, "IRR": ut.Currency{Currency: "IRR", DisplayName: "rial iranian", Symbol: "IRR"}, "KGS": ut.Currency{Currency: "KGS", DisplayName: "som kârgâz", Symbol: "KGS"}}
