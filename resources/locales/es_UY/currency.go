package es_UY

import "github.com/go-playground/universal-translator"

var currencies = ut.CurrencyFormatValue{"GYD": ut.Currency{Currency: "GYD", DisplayName: "dólar guyanés", Symbol: "GYD"}, "MTL": ut.Currency{Currency: "MTL", DisplayName: "lira maltesa", Symbol: ""}, "ETB": ut.Currency{Currency: "ETB", DisplayName: "bir", Symbol: "ETB"}, "PKR": ut.Currency{Currency: "PKR", DisplayName: "rupia pakistaní", Symbol: "PKR"}, "KRW": ut.Currency{Currency: "KRW", DisplayName: "won surcoreano", Symbol: "KRW"}, "LRD": ut.Currency{Currency: "LRD", DisplayName: "dólar liberiano", Symbol: "LRD"}, "GEL": ut.Currency{Currency: "GEL", DisplayName: "lari", Symbol: "GEL"}, "PEI": ut.Currency{Currency: "PEI", DisplayName: "inti peruano", Symbol: ""}, "AUD": ut.Currency{Currency: "AUD", DisplayName: "dólar australiano", Symbol: "AUD"}, "ESB": ut.Currency{Currency: "ESB", DisplayName: "peseta española (cuenta convertible)", Symbol: ""}, "UYP": ut.Currency{Currency: "UYP", DisplayName: "peso uruguayo (1975–1993)", Symbol: ""}, "SOS": ut.Currency{Currency: "SOS", DisplayName: "chelín somalí", Symbol: "SOS"}, "AZM": ut.Currency{Currency: "AZM", DisplayName: "manat azerí (1993–2006)", Symbol: ""}, "LKR": ut.Currency{Currency: "LKR", DisplayName: "rupia esrilanquesa", Symbol: "LKR"}, "ZAL": ut.Currency{Currency: "ZAL", DisplayName: "rand sudafricano (financiero)", Symbol: ""}, "TMT": ut.Currency{Currency: "TMT", DisplayName: "manat turcomano", Symbol: "TMT"}, "XPD": ut.Currency{Currency: "XPD", DisplayName: "paladio", Symbol: ""}, "INR": ut.Currency{Currency: "INR", DisplayName: "rupia india", Symbol: "INR"}, "NLG": ut.Currency{Currency: "NLG", DisplayName: "florín neerlandés", Symbol: ""}, "ARS": ut.Currency{Currency: "ARS", DisplayName: "peso argentino", Symbol: "ARS"}, "TND": ut.Currency{Currency: "TND", DisplayName: "dinar tunecino", Symbol: "TND"}, "PHP": ut.Currency{Currency: "PHP", DisplayName: "peso filipino", Symbol: "PHP"}, "BRR": ut.Currency{Currency: "BRR", DisplayName: "cruceiro brasileño", Symbol: ""}, "XAF": ut.Currency{Currency: "XAF", DisplayName: "franco CFA BEAC", Symbol: "XAF"}, "SYP": ut.Currency{Currency: "SYP", DisplayName: "libra siria", Symbol: "SYP"}, "RHD": ut.Currency{Currency: "RHD", DisplayName: "dólar rodesiano", Symbol: ""}, "SIT": ut.Currency{Currency: "SIT", DisplayName: "tólar esloveno", Symbol: ""}, "SLL": ut.Currency{Currency: "SLL", DisplayName: "leona", Symbol: "SLL"}, "LTT": ut.Currency{Currency: "LTT", DisplayName: "talonas lituano", Symbol: ""}, "MTP": ut.Currency{Currency: "MTP", DisplayName: "libra maltesa", Symbol: ""}, "FJD": ut.Currency{Currency: "FJD", DisplayName: "dólar fiyiano", Symbol: "FJD"}, "CVE": ut.Currency{Currency: "CVE", DisplayName: "escudo de Cabo Verde", Symbol: "CVE"}, "SUR": ut.Currency{Currency: "SUR", DisplayName: "rublo soviético", Symbol: ""}, "XAU": ut.Currency{Currency: "XAU", DisplayName: "oro", Symbol: ""}, "AFA": ut.Currency{Currency: "AFA", DisplayName: "afgani (1927–2002)", Symbol: ""}, "LUC": ut.Currency{Currency: "LUC", DisplayName: "franco convertible luxemburgués", Symbol: ""}, "PAB": ut.Currency{Currency: "PAB", DisplayName: "balboa panameño", Symbol: "PAB"}, "BRL": ut.Currency{Currency: "BRL", DisplayName: "real brasileño", Symbol: "BRL"}, "IEP": ut.Currency{Currency: "IEP", DisplayName: "libra irlandesa", Symbol: ""}, "SGD": ut.Currency{Currency: "SGD", DisplayName: "dólar singapurense", Symbol: "SGD"}, "GBP": ut.Currency{Currency: "GBP", DisplayName: "libra británica", Symbol: "GBP"}, "MZN": ut.Currency{Currency: "MZN", DisplayName: "metical", Symbol: "MZN"}, "UGX": ut.Currency{Currency: "UGX", DisplayName: "chelín ugandés", Symbol: "UGX"}, "XOF": ut.Currency{Currency: "XOF", DisplayName: "franco CFA BCEAO", Symbol: "XOF"}, "PES": ut.Currency{Currency: "PES", DisplayName: "sol peruano", Symbol: ""}, "TJR": ut.Currency{Currency: "TJR", DisplayName: "rublo tayiko", Symbol: ""}, "CYP": ut.Currency{Currency: "CYP", DisplayName: "libra chipriota", Symbol: ""}, "RWF": ut.Currency{Currency: "RWF", DisplayName: "franco ruandés", Symbol: "RWF"}, "JOD": ut.Currency{Currency: "JOD", DisplayName: "dinar jordano", Symbol: "JOD"}, "CUC": ut.Currency{Currency: "CUC", DisplayName: "peso cubano convertible", Symbol: "CUC"}, "CRC": ut.Currency{Currency: "CRC", DisplayName: "colón costarricense", Symbol: "CRC"}, "MYR": ut.Currency{Currency: "MYR", DisplayName: "ringit", Symbol: "MYR"}, "DJF": ut.Currency{Currency: "DJF", DisplayName: "franco yibutiano", Symbol: "DJF"}, "AZN": ut.Currency{Currency: "AZN", DisplayName: "manat azerí", Symbol: "AZN"}, "GTQ": ut.Currency{Currency: "GTQ", DisplayName: "quetzal guatemalteco", Symbol: "GTQ"}, "BAD": ut.Currency{Currency: "BAD", DisplayName: "dinar bosnio", Symbol: ""}, "ZRZ": ut.Currency{Currency: "ZRZ", DisplayName: "zaire zaireño", Symbol: ""}, "COP": ut.Currency{Currency: "COP", DisplayName: "peso colombiano", Symbol: "COP"}, "LVR": ut.Currency{Currency: "LVR", DisplayName: "rublo letón", Symbol: ""}, "FKP": ut.Currency{Currency: "FKP", DisplayName: "libra malvinense", Symbol: "FKP"}, "TWD": ut.Currency{Currency: "TWD", DisplayName: "nuevo dólar taiwanés", Symbol: "TWD"}, "COU": ut.Currency{Currency: "COU", DisplayName: "unidad de valor real colombiana", Symbol: ""}, "NZD": ut.Currency{Currency: "NZD", DisplayName: "dólar neozelandés", Symbol: "NZD"}, "DKK": ut.Currency{Currency: "DKK", DisplayName: "corona danesa", Symbol: "DKK"}, "BGL": ut.Currency{Currency: "BGL", DisplayName: "lev fuerte búlgaro", Symbol: ""}, "BRN": ut.Currency{Currency: "BRN", DisplayName: "nuevo cruzado brasileño", Symbol: ""}, "HUF": ut.Currency{Currency: "HUF", DisplayName: "forinto húngaro", Symbol: "HUF"}, "YUM": ut.Currency{Currency: "YUM", DisplayName: "super dinar yugoslavo", Symbol: ""}, "HNL": ut.Currency{Currency: "HNL", DisplayName: "lempira hondureño", Symbol: "HNL"}, "MRO": ut.Currency{Currency: "MRO", DisplayName: "uguiya", Symbol: "MRO"}, "XFO": ut.Currency{Currency: "XFO", DisplayName: "franco oro francés", Symbol: ""}, "AOR": ut.Currency{Currency: "AOR", DisplayName: "kwanza reajustado angoleño (1995–1999)", Symbol: ""}, "DEM": ut.Currency{Currency: "DEM", DisplayName: "marco alemán", Symbol: ""}, "SKK": ut.Currency{Currency: "SKK", DisplayName: "corona eslovaca", Symbol: ""}, "BOP": ut.Currency{Currency: "BOP", DisplayName: "peso boliviano", Symbol: ""}, "ITL": ut.Currency{Currency: "ITL", DisplayName: "lira italiana", Symbol: ""}, "CNY": ut.Currency{Currency: "CNY", DisplayName: "yuan", Symbol: "CNY"}, "MKD": ut.Currency{Currency: "MKD", DisplayName: "dinar macedonio", Symbol: "MKD"}, "BBD": ut.Currency{Currency: "BBD", DisplayName: "dólar barbadense", Symbol: "BBD"}, "KHR": ut.Currency{Currency: "KHR", DisplayName: "riel", Symbol: "KHR"}, "SRG": ut.Currency{Currency: "SRG", DisplayName: "florín surinamés", Symbol: ""}, "UZS": ut.Currency{Currency: "UZS", DisplayName: "sum", Symbol: "UZS"}, "GQE": ut.Currency{Currency: "GQE", DisplayName: "ekuele de Guinea Ecuatorial", Symbol: ""}, "BEF": ut.Currency{Currency: "BEF", DisplayName: "franco belga", Symbol: ""}, "ZMW": ut.Currency{Currency: "ZMW", DisplayName: "kuacha zambiano", Symbol: "ZMW"}, "BSD": ut.Currency{Currency: "BSD", DisplayName: "dólar bahameño", Symbol: "BSD"}, "KMF": ut.Currency{Currency: "KMF", DisplayName: "franco comorense", Symbol: "KMF"}, "AED": ut.Currency{Currency: "AED", DisplayName: "dírham de los Emiratos Árabes Unidos", Symbol: "AED"}, "GHC": ut.Currency{Currency: "GHC", DisplayName: "cedi ghanés (1979–2007)", Symbol: ""}, "MDL": ut.Currency{Currency: "MDL", DisplayName: "leu moldavo", Symbol: "MDL"}, "UYI": ut.Currency{Currency: "UYI", DisplayName: "peso uruguayo en unidades indexadas", Symbol: ""}, "EUR": ut.Currency{Currency: "EUR", DisplayName: "euro", Symbol: "€"}, "BOV": ut.Currency{Currency: "BOV", DisplayName: "MVDOL boliviano", Symbol: ""}, "WST": ut.Currency{Currency: "WST", DisplayName: "tala", Symbol: "WST"}, "ERN": ut.Currency{Currency: "ERN", DisplayName: "nakfa", Symbol: "ERN"}, "SCR": ut.Currency{Currency: "SCR", DisplayName: "rupia seychellense", Symbol: "SCR"}, "YUD": ut.Currency{Currency: "YUD", DisplayName: "dinar fuerte yugoslavo", Symbol: ""}, "BDT": ut.Currency{Currency: "BDT", DisplayName: "taka", Symbol: "BDT"}, "XTS": ut.Currency{Currency: "XTS", DisplayName: "código reservado para pruebas", Symbol: ""}, "USN": ut.Currency{Currency: "USN", DisplayName: "dólar estadounidense (día siguiente)", Symbol: ""}, "ILS": ut.Currency{Currency: "ILS", DisplayName: "nuevo séquel israelí", Symbol: "ILS"}, "JMD": ut.Currency{Currency: "JMD", DisplayName: "dólar jamaicano", Symbol: "JMD"}, "GWE": ut.Currency{Currency: "GWE", DisplayName: "escudo de Guinea Portuguesa", Symbol: ""}, "CAD": ut.Currency{Currency: "CAD", DisplayName: "dólar canadiense", Symbol: "CA$"}, "BUK": ut.Currency{Currency: "BUK", DisplayName: "kyat birmano", Symbol: ""}, "DOP": ut.Currency{Currency: "DOP", DisplayName: "peso dominicano", Symbol: "DOP"}, "ECS": ut.Currency{Currency: "ECS", DisplayName: "sucre ecuatoriano", Symbol: ""}, "HTG": ut.Currency{Currency: "HTG", DisplayName: "gourde haitiano", Symbol: "HTG"}, "ZWD": ut.Currency{Currency: "ZWD", DisplayName: "dólar de Zimbabue", Symbol: ""}, "BTN": ut.Currency{Currency: "BTN", DisplayName: "gultrum", Symbol: "BTN"}, "IDR": ut.Currency{Currency: "IDR", DisplayName: "rupia indonesia", Symbol: "IDR"}, "MLF": ut.Currency{Currency: "MLF", DisplayName: "franco malí", Symbol: ""}, "HRK": ut.Currency{Currency: "HRK", DisplayName: "kuna", Symbol: "HRK"}, "MOP": ut.Currency{Currency: "MOP", DisplayName: "pataca de Macao", Symbol: "MOP"}, "SDD": ut.Currency{Currency: "SDD", DisplayName: "dinar sudanés", Symbol: ""}, "GHS": ut.Currency{Currency: "GHS", DisplayName: "cedi", Symbol: "GHS"}, "XCD": ut.Currency{Currency: "XCD", DisplayName: "dólar del Caribe Oriental", Symbol: "XCD"}, "LAK": ut.Currency{Currency: "LAK", DisplayName: "kip", Symbol: "LAK"}, "CHW": ut.Currency{Currency: "CHW", DisplayName: "franco WIR", Symbol: ""}, "PYG": ut.Currency{Currency: "PYG", DisplayName: "guaraní paraguayo", Symbol: "PYG"}, "BRB": ut.Currency{Currency: "BRB", DisplayName: "nuevo cruceiro brasileño (1967–1986)", Symbol: ""}, "XBB": ut.Currency{Currency: "XBB", DisplayName: "unidad monetaria europea", Symbol: ""}, "MXV": ut.Currency{Currency: "MXV", DisplayName: "unidad de inversión (UDI) mexicana", Symbol: ""}, "KPW": ut.Currency{Currency: "KPW", DisplayName: "won norcoreano", Symbol: "KPW"}, "MXN": ut.Currency{Currency: "MXN", DisplayName: "peso mexicano", Symbol: "MXN"}, "CLF": ut.Currency{Currency: "CLF", DisplayName: "unidad de fomento chilena", Symbol: ""}, "MGF": ut.Currency{Currency: "MGF", DisplayName: "franco malgache", Symbol: ""}, "RSD": ut.Currency{Currency: "RSD", DisplayName: "dinar serbio", Symbol: "RSD"}, "ZMK": ut.Currency{Currency: "ZMK", DisplayName: "kwacha zambiano (1968–2012)", Symbol: ""}, "XXX": ut.Currency{Currency: "XXX", DisplayName: "moneda desconocida", Symbol: ""}, "SEK": ut.Currency{Currency: "SEK", DisplayName: "corona sueca", Symbol: "SEK"}, "DDM": ut.Currency{Currency: "DDM", DisplayName: "ostmark de Alemania del Este", Symbol: ""}, "AOK": ut.Currency{Currency: "AOK", DisplayName: "kwanza angoleño (1977–1990)", Symbol: ""}, "AFN": ut.Currency{Currency: "AFN", DisplayName: "afgani", Symbol: "AFN"}, "BOB": ut.Currency{Currency: "BOB", DisplayName: "boliviano", Symbol: "BOB"}, "GRD": ut.Currency{Currency: "GRD", DisplayName: "dracma griego", Symbol: ""}, "SDG": ut.Currency{Currency: "SDG", DisplayName: "libra sudanesa", Symbol: "SDG"}, "ZWL": ut.Currency{Currency: "ZWL", DisplayName: "dólar zimbabuense", Symbol: ""}, "MNT": ut.Currency{Currency: "MNT", DisplayName: "tugrik", Symbol: "MNT"}, "BEL": ut.Currency{Currency: "BEL", DisplayName: "franco belga (financiero)", Symbol: ""}, "SBD": ut.Currency{Currency: "SBD", DisplayName: "dólar salomonense", Symbol: "SBD"}, "GEK": ut.Currency{Currency: "GEK", DisplayName: "kupon larit georgiano", Symbol: ""}, "ESA": ut.Currency{Currency: "ESA", DisplayName: "peseta española (cuenta A)", Symbol: ""}, "KYD": ut.Currency{Currency: "KYD", DisplayName: "dólar de las Islas Caimán", Symbol: "KYD"}, "ADP": ut.Currency{Currency: "ADP", DisplayName: "peseta andorrana", Symbol: ""}, "MMK": ut.Currency{Currency: "MMK", DisplayName: "kiat", Symbol: "MMK"}, "XRE": ut.Currency{Currency: "XRE", DisplayName: "fondos RINET", Symbol: ""}, "DZD": ut.Currency{Currency: "DZD", DisplayName: "dinar argelino", Symbol: "DZD"}, "BEC": ut.Currency{Currency: "BEC", DisplayName: "franco belga (convertible)", Symbol: ""}, "UGS": ut.Currency{Currency: "UGS", DisplayName: "chelín ugandés (1966–1987)", Symbol: ""}, "LSL": ut.Currency{Currency: "LSL", DisplayName: "loti lesothense", Symbol: ""}, "QAR": ut.Currency{Currency: "QAR", DisplayName: "rial catarí", Symbol: "QAR"}, "UYU": ut.Currency{Currency: "UYU", DisplayName: "peso uruguayo", Symbol: "$"}, "SRD": ut.Currency{Currency: "SRD", DisplayName: "dólar surinamés", Symbol: "SRD"}, "ARA": ut.Currency{Currency: "ARA", DisplayName: "austral argentino", Symbol: ""}, "SDP": ut.Currency{Currency: "SDP", DisplayName: "libra sudanesa antigua", Symbol: ""}, "LUF": ut.Currency{Currency: "LUF", DisplayName: "franco luxemburgués", Symbol: ""}, "STD": ut.Currency{Currency: "STD", DisplayName: "dobra", Symbol: "STD"}, "ILP": ut.Currency{Currency: "ILP", DisplayName: "libra israelí", Symbol: ""}, "MXP": ut.Currency{Currency: "MXP", DisplayName: "peso de plata mexicano (1861–1992)", Symbol: ""}, "LTL": ut.Currency{Currency: "LTL", DisplayName: "litas lituano", Symbol: "LTL"}, "TJS": ut.Currency{Currency: "TJS", DisplayName: "somoni tayiko", Symbol: "TJS"}, "GIP": ut.Currency{Currency: "GIP", DisplayName: "libra gibraltareña", Symbol: "GIP"}, "UAK": ut.Currency{Currency: "UAK", DisplayName: "karbovanet ucraniano", Symbol: ""}, "CZK": ut.Currency{Currency: "CZK", DisplayName: "corona checa", Symbol: "CZK"}, "ZRN": ut.Currency{Currency: "ZRN", DisplayName: "nuevo zaire zaireño", Symbol: ""}, "THB": ut.Currency{Currency: "THB", DisplayName: "bat", Symbol: "฿"}, "BND": ut.Currency{Currency: "BND", DisplayName: "dólar bruneano", Symbol: "BND"}, "GNS": ut.Currency{Currency: "GNS", DisplayName: "syli guineano", Symbol: ""}, "BRC": ut.Currency{Currency: "BRC", DisplayName: "cruzado brasileño", Symbol: ""}, "TRY": ut.Currency{Currency: "TRY", DisplayName: "lira turca", Symbol: "TRY"}, "LVL": ut.Currency{Currency: "LVL", DisplayName: "lats letón", Symbol: "LVL"}, "BMD": ut.Currency{Currency: "BMD", DisplayName: "dólar de Bermudas", Symbol: "BMD"}, "ECV": ut.Currency{Currency: "ECV", DisplayName: "unidad de valor constante (UVC) ecuatoriana", Symbol: ""}, "ROL": ut.Currency{Currency: "ROL", DisplayName: "antiguo leu rumano", Symbol: ""}, "ARP": ut.Currency{Currency: "ARP", DisplayName: "peso argentino (1983–1985)", Symbol: ""}, "GMD": ut.Currency{Currency: "GMD", DisplayName: "dalasi", Symbol: "GMD"}, "LBP": ut.Currency{Currency: "LBP", DisplayName: "libra libanesa", Symbol: "LBP"}, "BWP": ut.Currency{Currency: "BWP", DisplayName: "pula", Symbol: "BWP"}, "GWP": ut.Currency{Currency: "GWP", DisplayName: "peso de Guinea-Bissáu", Symbol: ""}, "BZD": ut.Currency{Currency: "BZD", DisplayName: "dólar beliceño", Symbol: "BZD"}, "KGS": ut.Currency{Currency: "KGS", DisplayName: "som", Symbol: "KGS"}, "MAD": ut.Currency{Currency: "MAD", DisplayName: "dírham marroquí", Symbol: "MAD"}, "FIM": ut.Currency{Currency: "FIM", DisplayName: "marco finlandés", Symbol: ""}, "VEF": ut.Currency{Currency: "VEF", DisplayName: "bolívar venezolano", Symbol: "VEF"}, "MZM": ut.Currency{Currency: "MZM", DisplayName: "antiguo metical mozambiqueño", Symbol: ""}, "RUB": ut.Currency{Currency: "RUB", DisplayName: "rublo ruso", Symbol: "RUB"}, "VEB": ut.Currency{Currency: "VEB", DisplayName: "bolívar venezolano (1871–2008)", Symbol: ""}, "USS": ut.Currency{Currency: "USS", DisplayName: "dólar estadounidense (mismo día)", Symbol: ""}, "CDF": ut.Currency{Currency: "CDF", DisplayName: "franco congoleño", Symbol: "CDF"}, "XBA": ut.Currency{Currency: "XBA", DisplayName: "unidad compuesta europea", Symbol: ""}, "XPT": ut.Currency{Currency: "XPT", DisplayName: "platino", Symbol: ""}, "CSD": ut.Currency{Currency: "CSD", DisplayName: "antiguo dinar serbio", Symbol: ""}, "XBD": ut.Currency{Currency: "XBD", DisplayName: "unidad de cuenta europea (XBD)", Symbol: ""}, "CLP": ut.Currency{Currency: "CLP", DisplayName: "peso chileno", Symbol: "CLP"}, "SZL": ut.Currency{Currency: "SZL", DisplayName: "lilangeni", Symbol: "SZL"}, "XDR": ut.Currency{Currency: "XDR", DisplayName: "derechos especiales de giro", Symbol: ""}, "PGK": ut.Currency{Currency: "PGK", DisplayName: "kina", Symbol: "PGK"}, "HRD": ut.Currency{Currency: "HRD", DisplayName: "dinar croata", Symbol: ""}, "PLN": ut.Currency{Currency: "PLN", DisplayName: "esloti", Symbol: "PLN"}, "YER": ut.Currency{Currency: "YER", DisplayName: "rial yemení", Symbol: "YER"}, "ZAR": ut.Currency{Currency: "ZAR", DisplayName: "rand", Symbol: "ZAR"}, "FRF": ut.Currency{Currency: "FRF", DisplayName: "franco francés", Symbol: ""}, "AMD": ut.Currency{Currency: "AMD", DisplayName: "dram", Symbol: "AMD"}, "BAM": ut.Currency{Currency: "BAM", DisplayName: "marco convertible de Bosnia-Herzegovina", Symbol: "BAM"}, "PLZ": ut.Currency{Currency: "PLZ", DisplayName: "zloty polaco (1950–1995)", Symbol: ""}, "BHD": ut.Currency{Currency: "BHD", DisplayName: "dinar bahreiní", Symbol: "BHD"}, "PTE": ut.Currency{Currency: "PTE", DisplayName: "escudo portugués", Symbol: ""}, "GNF": ut.Currency{Currency: "GNF", DisplayName: "franco guineano", Symbol: "GNF"}, "JPY": ut.Currency{Currency: "JPY", DisplayName: "yen", Symbol: "JPY"}, "CUP": ut.Currency{Currency: "CUP", DisplayName: "peso cubano", Symbol: "CUP"}, "SSP": ut.Currency{Currency: "SSP", DisplayName: "libra sursudanesa", Symbol: "SSP"}, "NIO": ut.Currency{Currency: "NIO", DisplayName: "córdoba nicaragüense", Symbol: "NIO"}, "CHE": ut.Currency{Currency: "CHE", DisplayName: "euro WIR", Symbol: ""}, "MAF": ut.Currency{Currency: "MAF", DisplayName: "franco marroquí", Symbol: ""}, "AWG": ut.Currency{Currency: "AWG", DisplayName: "florín arubeño", Symbol: "AWG"}, "BRE": ut.Currency{Currency: "BRE", DisplayName: "cruceiro brasileño (1990–1993)", Symbol: ""}, "RUR": ut.Currency{Currency: "RUR", DisplayName: "rublo ruso (1991–1998)", Symbol: ""}, "NPR": ut.Currency{Currency: "NPR", DisplayName: "rupia nepalí", Symbol: "NPR"}, "MUR": ut.Currency{Currency: "MUR", DisplayName: "rupia mauriciana", Symbol: "MUR"}, "ESP": ut.Currency{Currency: "ESP", DisplayName: "peseta española", Symbol: "₧"}, "ALL": ut.Currency{Currency: "ALL", DisplayName: "lek", Symbol: "ALL"}, "XEU": ut.Currency{Currency: "XEU", DisplayName: "unidad de moneda europea", Symbol: ""}, "CHF": ut.Currency{Currency: "CHF", DisplayName: "franco suizo", Symbol: "CHF"}, "SAR": ut.Currency{Currency: "SAR", DisplayName: "rial saudí", Symbol: "SAR"}, "EEK": ut.Currency{Currency: "EEK", DisplayName: "corona estonia", Symbol: ""}, "TRL": ut.Currency{Currency: "TRL", DisplayName: "lira turca (1922–2005)", Symbol: ""}, "YDD": ut.Currency{Currency: "YDD", DisplayName: "dinar yemení", Symbol: ""}, "XFU": ut.Currency{Currency: "XFU", DisplayName: "franco UIC francés", Symbol: ""}, "KWD": ut.Currency{Currency: "KWD", DisplayName: "dinar kuwaití", Symbol: "KWD"}, "BYB": ut.Currency{Currency: "BYB", DisplayName: "nuevo rublo bielorruso (1994–1999)", Symbol: ""}, "TZS": ut.Currency{Currency: "TZS", DisplayName: "chelín tanzano", Symbol: "TZS"}, "SHP": ut.Currency{Currency: "SHP", DisplayName: "libra de Santa Elena", Symbol: "SHP"}, "TOP": ut.Currency{Currency: "TOP", DisplayName: "paanga", Symbol: "TOP"}, "ISK": ut.Currency{Currency: "ISK", DisplayName: "corona islandesa", Symbol: "ISK"}, "ATS": ut.Currency{Currency: "ATS", DisplayName: "chelín austriaco", Symbol: ""}, "NOK": ut.Currency{Currency: "NOK", DisplayName: "corona noruega", Symbol: "NOK"}, "SVC": ut.Currency{Currency: "SVC", DisplayName: "colón salvadoreño", Symbol: ""}, "NAD": ut.Currency{Currency: "NAD", DisplayName: "dólar namibio", Symbol: "NAD"}, "IQD": ut.Currency{Currency: "IQD", DisplayName: "dinar iraquí", Symbol: "IQD"}, "LUL": ut.Currency{Currency: "LUL", DisplayName: "franco financiero luxemburgués", Symbol: ""}, "XBC": ut.Currency{Currency: "XBC", DisplayName: "unidad de cuenta europea (XBC)", Symbol: ""}, "KES": ut.Currency{Currency: "KES", DisplayName: "chelín keniano", Symbol: "KES"}, "RON": ut.Currency{Currency: "RON", DisplayName: "leu rumano", Symbol: "RON"}, "VUV": ut.Currency{Currency: "VUV", DisplayName: "vatu", Symbol: "VUV"}, "NGN": ut.Currency{Currency: "NGN", DisplayName: "naira", Symbol: "NGN"}, "OMR": ut.Currency{Currency: "OMR", DisplayName: "rial omaní", Symbol: "OMR"}, "HKD": ut.Currency{Currency: "HKD", DisplayName: "dólar hongkonés", Symbol: "HKD"}, "MGA": ut.Currency{Currency: "MGA", DisplayName: "ariari", Symbol: "MGA"}, "TTD": ut.Currency{Currency: "TTD", DisplayName: "dólar de Trinidad y Tobago", Symbol: "TTD"}, "USD": ut.Currency{Currency: "USD", DisplayName: "dólar estadounidense", Symbol: "US$"}, "BGN": ut.Currency{Currency: "BGN", DisplayName: "lev búlgaro", Symbol: "BGN"}, "NIC": ut.Currency{Currency: "NIC", DisplayName: "córdoba nicaragüense (1988–1991)", Symbol: ""}, "AOA": ut.Currency{Currency: "AOA", DisplayName: "kuanza", Symbol: "AOA"}, "XAG": ut.Currency{Currency: "XAG", DisplayName: "plata", Symbol: ""}, "LYD": ut.Currency{Currency: "LYD", DisplayName: "dinar libio", Symbol: "LYD"}, "MVR": ut.Currency{Currency: "MVR", DisplayName: "rufiya", Symbol: "MVR"}, "YUN": ut.Currency{Currency: "YUN", DisplayName: "dinar convertible yugoslavo", Symbol: ""}, "BIF": ut.Currency{Currency: "BIF", DisplayName: "franco burundés", Symbol: "BIF"}, "XPF": ut.Currency{Currency: "XPF", DisplayName: "franco CFP", Symbol: "CFPF"}, "ANG": ut.Currency{Currency: "ANG", DisplayName: "florín de las Antillas Neerlandesas", Symbol: "ANG"}, "PEN": ut.Currency{Currency: "PEN", DisplayName: "nuevo sol peruano", Symbol: "PEN"}, "UAH": ut.Currency{Currency: "UAH", DisplayName: "grivna", Symbol: "UAH"}, "KZT": ut.Currency{Currency: "KZT", DisplayName: "tenge kazako", Symbol: "KZT"}, "CSK": ut.Currency{Currency: "CSK", DisplayName: "corona fuerte checoslovaca", Symbol: ""}, "TPE": ut.Currency{Currency: "TPE", DisplayName: "escudo timorense", Symbol: ""}, "MWK": ut.Currency{Currency: "MWK", DisplayName: "kwacha malauí", Symbol: "MWK"}, "AON": ut.Currency{Currency: "AON", DisplayName: "nuevo kwanza angoleño (1990–2000)", Symbol: ""}, "MZE": ut.Currency{Currency: "MZE", DisplayName: "escudo mozambiqueño", Symbol: ""}, "IRR": ut.Currency{Currency: "IRR", DisplayName: "rial iraní", Symbol: "IRR"}, "TMM": ut.Currency{Currency: "TMM", DisplayName: "manat turcomano (1993–2009)", Symbol: ""}, "VND": ut.Currency{Currency: "VND", DisplayName: "dong", Symbol: "₫"}, "EGP": ut.Currency{Currency: "EGP", DisplayName: "libra egipcia", Symbol: "EGP"}, "BYR": ut.Currency{Currency: "BYR", DisplayName: "rublo bielorruso", Symbol: "BYR"}}