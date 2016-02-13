package es_CU

import "github.com/go-playground/universal-translator"

var currencies = ut.CurrencyFormatValue{"ZAR": ut.Currency{Currency: "ZAR", DisplayName: "rand", Symbol: "ZAR"}, "PLN": ut.Currency{Currency: "PLN", DisplayName: "esloti", Symbol: "PLN"}, "ARP": ut.Currency{Currency: "ARP", DisplayName: "peso argentino (1983–1985)", Symbol: ""}, "XBD": ut.Currency{Currency: "XBD", DisplayName: "unidad de cuenta europea (XBD)", Symbol: ""}, "AOR": ut.Currency{Currency: "AOR", DisplayName: "kwanza reajustado angoleño (1995–1999)", Symbol: ""}, "MTP": ut.Currency{Currency: "MTP", DisplayName: "libra maltesa", Symbol: ""}, "DJF": ut.Currency{Currency: "DJF", DisplayName: "franco yibutiano", Symbol: "DJF"}, "WST": ut.Currency{Currency: "WST", DisplayName: "tala", Symbol: "WST"}, "LSL": ut.Currency{Currency: "LSL", DisplayName: "loti lesothense", Symbol: ""}, "UZS": ut.Currency{Currency: "UZS", DisplayName: "sum", Symbol: "UZS"}, "ESB": ut.Currency{Currency: "ESB", DisplayName: "peseta española (cuenta convertible)", Symbol: ""}, "JPY": ut.Currency{Currency: "JPY", DisplayName: "yen", Symbol: "JPY"}, "KES": ut.Currency{Currency: "KES", DisplayName: "chelín keniano", Symbol: "KES"}, "ANG": ut.Currency{Currency: "ANG", DisplayName: "florín de las Antillas Neerlandesas", Symbol: "ANG"}, "CYP": ut.Currency{Currency: "CYP", DisplayName: "libra chipriota", Symbol: ""}, "AED": ut.Currency{Currency: "AED", DisplayName: "dírham de los Emiratos Árabes Unidos", Symbol: "AED"}, "KYD": ut.Currency{Currency: "KYD", DisplayName: "dólar de las Islas Caimán", Symbol: "KYD"}, "GYD": ut.Currency{Currency: "GYD", DisplayName: "dólar guyanés", Symbol: "GYD"}, "BAD": ut.Currency{Currency: "BAD", DisplayName: "dinar bosnio", Symbol: ""}, "TOP": ut.Currency{Currency: "TOP", DisplayName: "paanga", Symbol: "TOP"}, "BZD": ut.Currency{Currency: "BZD", DisplayName: "dólar beliceño", Symbol: "BZD"}, "ADP": ut.Currency{Currency: "ADP", DisplayName: "peseta andorrana", Symbol: ""}, "BSD": ut.Currency{Currency: "BSD", DisplayName: "dólar bahameño", Symbol: "BSD"}, "PHP": ut.Currency{Currency: "PHP", DisplayName: "peso filipino", Symbol: "PHP"}, "CHW": ut.Currency{Currency: "CHW", DisplayName: "franco WIR", Symbol: ""}, "BWP": ut.Currency{Currency: "BWP", DisplayName: "pula", Symbol: "BWP"}, "NGN": ut.Currency{Currency: "NGN", DisplayName: "naira", Symbol: "NGN"}, "NPR": ut.Currency{Currency: "NPR", DisplayName: "rupia nepalí", Symbol: "NPR"}, "IEP": ut.Currency{Currency: "IEP", DisplayName: "libra irlandesa", Symbol: ""}, "GHC": ut.Currency{Currency: "GHC", DisplayName: "cedi ghanés (1979–2007)", Symbol: ""}, "MMK": ut.Currency{Currency: "MMK", DisplayName: "kiat", Symbol: "MMK"}, "ZAL": ut.Currency{Currency: "ZAL", DisplayName: "rand sudafricano (financiero)", Symbol: ""}, "CHE": ut.Currency{Currency: "CHE", DisplayName: "euro WIR", Symbol: ""}, "SRG": ut.Currency{Currency: "SRG", DisplayName: "florín surinamés", Symbol: ""}, "ZWL": ut.Currency{Currency: "ZWL", DisplayName: "dólar zimbabuense", Symbol: ""}, "EGP": ut.Currency{Currency: "EGP", DisplayName: "libra egipcia", Symbol: "EGP"}, "SIT": ut.Currency{Currency: "SIT", DisplayName: "tólar esloveno", Symbol: ""}, "UGS": ut.Currency{Currency: "UGS", DisplayName: "chelín ugandés (1966–1987)", Symbol: ""}, "CNY": ut.Currency{Currency: "CNY", DisplayName: "yuan", Symbol: "CNY"}, "SDG": ut.Currency{Currency: "SDG", DisplayName: "libra sudanesa", Symbol: "SDG"}, "XDR": ut.Currency{Currency: "XDR", DisplayName: "derechos especiales de giro", Symbol: ""}, "MAF": ut.Currency{Currency: "MAF", DisplayName: "franco marroquí", Symbol: ""}, "XPD": ut.Currency{Currency: "XPD", DisplayName: "paladio", Symbol: ""}, "PAB": ut.Currency{Currency: "PAB", DisplayName: "balboa panameño", Symbol: "PAB"}, "HUF": ut.Currency{Currency: "HUF", DisplayName: "forinto húngaro", Symbol: "HUF"}, "RUR": ut.Currency{Currency: "RUR", DisplayName: "rublo ruso (1991–1998)", Symbol: ""}, "BGN": ut.Currency{Currency: "BGN", DisplayName: "lev búlgaro", Symbol: "BGN"}, "LVL": ut.Currency{Currency: "LVL", DisplayName: "lats letón", Symbol: "LVL"}, "STD": ut.Currency{Currency: "STD", DisplayName: "dobra", Symbol: "STD"}, "KHR": ut.Currency{Currency: "KHR", DisplayName: "riel", Symbol: "KHR"}, "AOK": ut.Currency{Currency: "AOK", DisplayName: "kwanza angoleño (1977–1990)", Symbol: ""}, "FIM": ut.Currency{Currency: "FIM", DisplayName: "marco finlandés", Symbol: ""}, "PEI": ut.Currency{Currency: "PEI", DisplayName: "inti peruano", Symbol: ""}, "AMD": ut.Currency{Currency: "AMD", DisplayName: "dram", Symbol: "AMD"}, "JOD": ut.Currency{Currency: "JOD", DisplayName: "dinar jordano", Symbol: "JOD"}, "BOV": ut.Currency{Currency: "BOV", DisplayName: "MVDOL boliviano", Symbol: ""}, "GHS": ut.Currency{Currency: "GHS", DisplayName: "cedi", Symbol: "GHS"}, "GTQ": ut.Currency{Currency: "GTQ", DisplayName: "quetzal guatemalteco", Symbol: "GTQ"}, "XFU": ut.Currency{Currency: "XFU", DisplayName: "franco UIC francés", Symbol: ""}, "MXV": ut.Currency{Currency: "MXV", DisplayName: "unidad de inversión (UDI) mexicana", Symbol: ""}, "IQD": ut.Currency{Currency: "IQD", DisplayName: "dinar iraquí", Symbol: "IQD"}, "HNL": ut.Currency{Currency: "HNL", DisplayName: "lempira hondureño", Symbol: "HNL"}, "BRC": ut.Currency{Currency: "BRC", DisplayName: "cruzado brasileño", Symbol: ""}, "CRC": ut.Currency{Currency: "CRC", DisplayName: "colón costarricense", Symbol: "CRC"}, "TTD": ut.Currency{Currency: "TTD", DisplayName: "dólar de Trinidad y Tobago", Symbol: "TTD"}, "LTL": ut.Currency{Currency: "LTL", DisplayName: "litas lituano", Symbol: "LTL"}, "DKK": ut.Currency{Currency: "DKK", DisplayName: "corona danesa", Symbol: "DKK"}, "MLF": ut.Currency{Currency: "MLF", DisplayName: "franco malí", Symbol: ""}, "CVE": ut.Currency{Currency: "CVE", DisplayName: "escudo de Cabo Verde", Symbol: "CVE"}, "QAR": ut.Currency{Currency: "QAR", DisplayName: "rial catarí", Symbol: "QAR"}, "LTT": ut.Currency{Currency: "LTT", DisplayName: "talonas lituano", Symbol: ""}, "ESA": ut.Currency{Currency: "ESA", DisplayName: "peseta española (cuenta A)", Symbol: ""}, "LBP": ut.Currency{Currency: "LBP", DisplayName: "libra libanesa", Symbol: "LBP"}, "KWD": ut.Currency{Currency: "KWD", DisplayName: "dinar kuwaití", Symbol: "KWD"}, "XAF": ut.Currency{Currency: "XAF", DisplayName: "franco CFA BEAC", Symbol: "XAF"}, "BGL": ut.Currency{Currency: "BGL", DisplayName: "lev fuerte búlgaro", Symbol: ""}, "MUR": ut.Currency{Currency: "MUR", DisplayName: "rupia mauriciana", Symbol: "MUR"}, "MGF": ut.Currency{Currency: "MGF", DisplayName: "franco malgache", Symbol: ""}, "RSD": ut.Currency{Currency: "RSD", DisplayName: "dinar serbio", Symbol: "RSD"}, "UAH": ut.Currency{Currency: "UAH", DisplayName: "grivna", Symbol: "UAH"}, "TJS": ut.Currency{Currency: "TJS", DisplayName: "somoni tayiko", Symbol: "TJS"}, "RWF": ut.Currency{Currency: "RWF", DisplayName: "franco ruandés", Symbol: "RWF"}, "AFA": ut.Currency{Currency: "AFA", DisplayName: "afgani (1927–2002)", Symbol: ""}, "ECV": ut.Currency{Currency: "ECV", DisplayName: "unidad de valor constante (UVC) ecuatoriana", Symbol: ""}, "AUD": ut.Currency{Currency: "AUD", DisplayName: "dólar australiano", Symbol: "AUD"}, "RUB": ut.Currency{Currency: "RUB", DisplayName: "rublo ruso", Symbol: "RUB"}, "PES": ut.Currency{Currency: "PES", DisplayName: "sol peruano", Symbol: ""}, "USN": ut.Currency{Currency: "USN", DisplayName: "dólar estadounidense (día siguiente)", Symbol: ""}, "CZK": ut.Currency{Currency: "CZK", DisplayName: "corona checa", Symbol: "CZK"}, "SBD": ut.Currency{Currency: "SBD", DisplayName: "dólar salomonense", Symbol: "SBD"}, "PLZ": ut.Currency{Currency: "PLZ", DisplayName: "zloty polaco (1950–1995)", Symbol: ""}, "LAK": ut.Currency{Currency: "LAK", DisplayName: "kip", Symbol: "LAK"}, "PEN": ut.Currency{Currency: "PEN", DisplayName: "nuevo sol peruano", Symbol: "PEN"}, "BAM": ut.Currency{Currency: "BAM", DisplayName: "marco convertible de Bosnia-Herzegovina", Symbol: "BAM"}, "ATS": ut.Currency{Currency: "ATS", DisplayName: "chelín austriaco", Symbol: ""}, "GNS": ut.Currency{Currency: "GNS", DisplayName: "syli guineano", Symbol: ""}, "NZD": ut.Currency{Currency: "NZD", DisplayName: "dólar neozelandés", Symbol: "NZD"}, "JMD": ut.Currency{Currency: "JMD", DisplayName: "dólar jamaicano", Symbol: "JMD"}, "MYR": ut.Currency{Currency: "MYR", DisplayName: "ringit", Symbol: "MYR"}, "ZRZ": ut.Currency{Currency: "ZRZ", DisplayName: "zaire zaireño", Symbol: ""}, "AWG": ut.Currency{Currency: "AWG", DisplayName: "florín arubeño", Symbol: "AWG"}, "SVC": ut.Currency{Currency: "SVC", DisplayName: "colón salvadoreño", Symbol: ""}, "AFN": ut.Currency{Currency: "AFN", DisplayName: "afgani", Symbol: "AFN"}, "XPF": ut.Currency{Currency: "XPF", DisplayName: "franco CFP", Symbol: "CFPF"}, "GRD": ut.Currency{Currency: "GRD", DisplayName: "dracma griego", Symbol: ""}, "ZMK": ut.Currency{Currency: "ZMK", DisplayName: "kwacha zambiano (1968–2012)", Symbol: ""}, "TRL": ut.Currency{Currency: "TRL", DisplayName: "lira turca (1922–2005)", Symbol: ""}, "ILS": ut.Currency{Currency: "ILS", DisplayName: "nuevo séquel israelí", Symbol: "ILS"}, "MZM": ut.Currency{Currency: "MZM", DisplayName: "antiguo metical mozambiqueño", Symbol: ""}, "XXX": ut.Currency{Currency: "XXX", DisplayName: "moneda desconocida", Symbol: ""}, "YDD": ut.Currency{Currency: "YDD", DisplayName: "dinar yemení", Symbol: ""}, "RON": ut.Currency{Currency: "RON", DisplayName: "leu rumano", Symbol: "RON"}, "GBP": ut.Currency{Currency: "GBP", DisplayName: "libra británica", Symbol: "GBP"}, "OMR": ut.Currency{Currency: "OMR", DisplayName: "rial omaní", Symbol: "OMR"}, "INR": ut.Currency{Currency: "INR", DisplayName: "rupia india", Symbol: "INR"}, "MZE": ut.Currency{Currency: "MZE", DisplayName: "escudo mozambiqueño", Symbol: ""}, "MKD": ut.Currency{Currency: "MKD", DisplayName: "dinar macedonio", Symbol: "MKD"}, "ZWD": ut.Currency{Currency: "ZWD", DisplayName: "dólar de Zimbabue", Symbol: ""}, "BRL": ut.Currency{Currency: "BRL", DisplayName: "real brasileño", Symbol: "BRL"}, "VEF": ut.Currency{Currency: "VEF", DisplayName: "bolívar venezolano", Symbol: "VEF"}, "BYB": ut.Currency{Currency: "BYB", DisplayName: "nuevo rublo bielorruso (1994–1999)", Symbol: ""}, "TWD": ut.Currency{Currency: "TWD", DisplayName: "nuevo dólar taiwanés", Symbol: "TWD"}, "SGD": ut.Currency{Currency: "SGD", DisplayName: "dólar singapurense", Symbol: "SGD"}, "CDF": ut.Currency{Currency: "CDF", DisplayName: "franco congoleño", Symbol: "CDF"}, "ARS": ut.Currency{Currency: "ARS", DisplayName: "peso argentino", Symbol: "ARS"}, "LUL": ut.Currency{Currency: "LUL", DisplayName: "franco financiero luxemburgués", Symbol: ""}, "BRR": ut.Currency{Currency: "BRR", DisplayName: "cruceiro brasileño", Symbol: ""}, "XEU": ut.Currency{Currency: "XEU", DisplayName: "unidad de moneda europea", Symbol: ""}, "SZL": ut.Currency{Currency: "SZL", DisplayName: "lilangeni", Symbol: "SZL"}, "XCD": ut.Currency{Currency: "XCD", DisplayName: "dólar del Caribe Oriental", Symbol: "XCD"}, "MTL": ut.Currency{Currency: "MTL", DisplayName: "lira maltesa", Symbol: ""}, "SCR": ut.Currency{Currency: "SCR", DisplayName: "rupia seychellense", Symbol: "SCR"}, "COU": ut.Currency{Currency: "COU", DisplayName: "unidad de valor real colombiana", Symbol: ""}, "FKP": ut.Currency{Currency: "FKP", DisplayName: "libra malvinense", Symbol: "FKP"}, "BOB": ut.Currency{Currency: "BOB", DisplayName: "boliviano", Symbol: "BOB"}, "BBD": ut.Currency{Currency: "BBD", DisplayName: "dólar barbadense", Symbol: "BBD"}, "GQE": ut.Currency{Currency: "GQE", DisplayName: "ekuele de Guinea Ecuatorial", Symbol: ""}, "GWP": ut.Currency{Currency: "GWP", DisplayName: "peso de Guinea-Bissáu", Symbol: ""}, "AZN": ut.Currency{Currency: "AZN", DisplayName: "manat azerí", Symbol: "AZN"}, "IRR": ut.Currency{Currency: "IRR", DisplayName: "rial iraní", Symbol: "IRR"}, "ISK": ut.Currency{Currency: "ISK", DisplayName: "corona islandesa", Symbol: "ISK"}, "TMT": ut.Currency{Currency: "TMT", DisplayName: "manat turcomano", Symbol: "TMT"}, "BRE": ut.Currency{Currency: "BRE", DisplayName: "cruceiro brasileño (1990–1993)", Symbol: ""}, "CHF": ut.Currency{Currency: "CHF", DisplayName: "franco suizo", Symbol: "CHF"}, "YUD": ut.Currency{Currency: "YUD", DisplayName: "dinar fuerte yugoslavo", Symbol: ""}, "KZT": ut.Currency{Currency: "KZT", DisplayName: "tenge kazako", Symbol: "KZT"}, "SEK": ut.Currency{Currency: "SEK", DisplayName: "corona sueca", Symbol: "SEK"}, "MOP": ut.Currency{Currency: "MOP", DisplayName: "pataca de Macao", Symbol: "MOP"}, "BMD": ut.Currency{Currency: "BMD", DisplayName: "dólar de Bermudas", Symbol: "BMD"}, "ILP": ut.Currency{Currency: "ILP", DisplayName: "libra israelí", Symbol: ""}, "SRD": ut.Currency{Currency: "SRD", DisplayName: "dólar surinamés", Symbol: "SRD"}, "UAK": ut.Currency{Currency: "UAK", DisplayName: "karbovanet ucraniano", Symbol: ""}, "BEL": ut.Currency{Currency: "BEL", DisplayName: "franco belga (financiero)", Symbol: ""}, "CAD": ut.Currency{Currency: "CAD", DisplayName: "dólar canadiense", Symbol: "CA$"}, "KMF": ut.Currency{Currency: "KMF", DisplayName: "franco comorense", Symbol: "KMF"}, "LUC": ut.Currency{Currency: "LUC", DisplayName: "franco convertible luxemburgués", Symbol: ""}, "SDD": ut.Currency{Currency: "SDD", DisplayName: "dinar sudanés", Symbol: ""}, "LVR": ut.Currency{Currency: "LVR", DisplayName: "rublo letón", Symbol: ""}, "USD": ut.Currency{Currency: "USD", DisplayName: "dólar estadounidense", Symbol: "US$"}, "GIP": ut.Currency{Currency: "GIP", DisplayName: "libra gibraltareña", Symbol: "GIP"}, "SYP": ut.Currency{Currency: "SYP", DisplayName: "libra siria", Symbol: "SYP"}, "GWE": ut.Currency{Currency: "GWE", DisplayName: "escudo de Guinea Portuguesa", Symbol: ""}, "DDM": ut.Currency{Currency: "DDM", DisplayName: "ostmark de Alemania del Este", Symbol: ""}, "UYU": ut.Currency{Currency: "UYU", DisplayName: "peso uruguayo", Symbol: "UYU"}, "BND": ut.Currency{Currency: "BND", DisplayName: "dólar bruneano", Symbol: "BND"}, "BOP": ut.Currency{Currency: "BOP", DisplayName: "peso boliviano", Symbol: ""}, "PYG": ut.Currency{Currency: "PYG", DisplayName: "guaraní paraguayo", Symbol: "PYG"}, "XRE": ut.Currency{Currency: "XRE", DisplayName: "fondos RINET", Symbol: ""}, "RHD": ut.Currency{Currency: "RHD", DisplayName: "dólar rodesiano", Symbol: ""}, "TMM": ut.Currency{Currency: "TMM", DisplayName: "manat turcomano (1993–2009)", Symbol: ""}, "FRF": ut.Currency{Currency: "FRF", DisplayName: "franco francés", Symbol: ""}, "ERN": ut.Currency{Currency: "ERN", DisplayName: "nakfa", Symbol: "ERN"}, "MRO": ut.Currency{Currency: "MRO", DisplayName: "uguiya", Symbol: "MRO"}, "BHD": ut.Currency{Currency: "BHD", DisplayName: "dinar bahreiní", Symbol: "BHD"}, "COP": ut.Currency{Currency: "COP", DisplayName: "peso colombiano", Symbol: "COP"}, "SOS": ut.Currency{Currency: "SOS", DisplayName: "chelín somalí", Symbol: "SOS"}, "SLL": ut.Currency{Currency: "SLL", DisplayName: "leona", Symbol: "SLL"}, "HTG": ut.Currency{Currency: "HTG", DisplayName: "gourde haitiano", Symbol: "HTG"}, "ROL": ut.Currency{Currency: "ROL", DisplayName: "antiguo leu rumano", Symbol: ""}, "CLF": ut.Currency{Currency: "CLF", DisplayName: "unidad de fomento chilena", Symbol: ""}, "XTS": ut.Currency{Currency: "XTS", DisplayName: "código reservado para pruebas", Symbol: ""}, "CUC": ut.Currency{Currency: "CUC", DisplayName: "peso cubano convertible", Symbol: "CUC"}, "GEK": ut.Currency{Currency: "GEK", DisplayName: "kupon larit georgiano", Symbol: ""}, "CUP": ut.Currency{Currency: "CUP", DisplayName: "peso cubano", Symbol: "$"}, "MXN": ut.Currency{Currency: "MXN", DisplayName: "peso mexicano", Symbol: "MXN"}, "BTN": ut.Currency{Currency: "BTN", DisplayName: "gultrum", Symbol: "BTN"}, "SKK": ut.Currency{Currency: "SKK", DisplayName: "corona eslovaca", Symbol: ""}, "ALL": ut.Currency{Currency: "ALL", DisplayName: "lek", Symbol: "ALL"}, "TND": ut.Currency{Currency: "TND", DisplayName: "dinar tunecino", Symbol: "TND"}, "MXP": ut.Currency{Currency: "MXP", DisplayName: "peso de plata mexicano (1861–1992)", Symbol: ""}, "XAG": ut.Currency{Currency: "XAG", DisplayName: "plata", Symbol: ""}, "DOP": ut.Currency{Currency: "DOP", DisplayName: "peso dominicano", Symbol: "DOP"}, "CSK": ut.Currency{Currency: "CSK", DisplayName: "corona fuerte checoslovaca", Symbol: ""}, "VND": ut.Currency{Currency: "VND", DisplayName: "dong", Symbol: "₫"}, "VUV": ut.Currency{Currency: "VUV", DisplayName: "vatu", Symbol: "VUV"}, "BYR": ut.Currency{Currency: "BYR", DisplayName: "rublo bielorruso", Symbol: "BYR"}, "LKR": ut.Currency{Currency: "LKR", DisplayName: "rupia esrilanquesa", Symbol: "LKR"}, "MZN": ut.Currency{Currency: "MZN", DisplayName: "metical", Symbol: "MZN"}, "ETB": ut.Currency{Currency: "ETB", DisplayName: "bir", Symbol: "ETB"}, "BDT": ut.Currency{Currency: "BDT", DisplayName: "taka", Symbol: "BDT"}, "BRN": ut.Currency{Currency: "BRN", DisplayName: "nuevo cruzado brasileño", Symbol: ""}, "SAR": ut.Currency{Currency: "SAR", DisplayName: "rial saudí", Symbol: "SAR"}, "SHP": ut.Currency{Currency: "SHP", DisplayName: "libra de Santa Elena", Symbol: "SHP"}, "PTE": ut.Currency{Currency: "PTE", DisplayName: "escudo portugués", Symbol: ""}, "FJD": ut.Currency{Currency: "FJD", DisplayName: "dólar fiyiano", Symbol: "FJD"}, "SDP": ut.Currency{Currency: "SDP", DisplayName: "libra sudanesa antigua", Symbol: ""}, "YER": ut.Currency{Currency: "YER", DisplayName: "rial yemení", Symbol: "YER"}, "MDL": ut.Currency{Currency: "MDL", DisplayName: "leu moldavo", Symbol: "MDL"}, "TRY": ut.Currency{Currency: "TRY", DisplayName: "lira turca", Symbol: "TRY"}, "BUK": ut.Currency{Currency: "BUK", DisplayName: "kyat birmano", Symbol: ""}, "TJR": ut.Currency{Currency: "TJR", DisplayName: "rublo tayiko", Symbol: ""}, "GNF": ut.Currency{Currency: "GNF", DisplayName: "franco guineano", Symbol: "GNF"}, "SSP": ut.Currency{Currency: "SSP", DisplayName: "libra sursudanesa", Symbol: "SSP"}, "GEL": ut.Currency{Currency: "GEL", DisplayName: "lari", Symbol: "GEL"}, "TPE": ut.Currency{Currency: "TPE", DisplayName: "escudo timorense", Symbol: ""}, "THB": ut.Currency{Currency: "THB", DisplayName: "bat", Symbol: "฿"}, "USS": ut.Currency{Currency: "USS", DisplayName: "dólar estadounidense (mismo día)", Symbol: ""}, "MWK": ut.Currency{Currency: "MWK", DisplayName: "kwacha malauí", Symbol: "MWK"}, "CSD": ut.Currency{Currency: "CSD", DisplayName: "antiguo dinar serbio", Symbol: ""}, "GMD": ut.Currency{Currency: "GMD", DisplayName: "dalasi", Symbol: "GMD"}, "ECS": ut.Currency{Currency: "ECS", DisplayName: "sucre ecuatoriano", Symbol: ""}, "LUF": ut.Currency{Currency: "LUF", DisplayName: "franco luxemburgués", Symbol: ""}, "MGA": ut.Currency{Currency: "MGA", DisplayName: "ariari", Symbol: "MGA"}, "UYP": ut.Currency{Currency: "UYP", DisplayName: "peso uruguayo (1975–1993)", Symbol: ""}, "PKR": ut.Currency{Currency: "PKR", DisplayName: "rupia pakistaní", Symbol: "PKR"}, "UYI": ut.Currency{Currency: "UYI", DisplayName: "peso uruguayo en unidades indexadas", Symbol: ""}, "MNT": ut.Currency{Currency: "MNT", DisplayName: "tugrik", Symbol: "MNT"}, "XOF": ut.Currency{Currency: "XOF", DisplayName: "franco CFA BCEAO", Symbol: "XOF"}, "XBC": ut.Currency{Currency: "XBC", DisplayName: "unidad de cuenta europea (XBC)", Symbol: ""}, "VEB": ut.Currency{Currency: "VEB", DisplayName: "bolívar venezolano (1871–2008)", Symbol: ""}, "NIO": ut.Currency{Currency: "NIO", DisplayName: "córdoba nicaragüense", Symbol: "NIO"}, "XAU": ut.Currency{Currency: "XAU", DisplayName: "oro", Symbol: ""}, "LRD": ut.Currency{Currency: "LRD", DisplayName: "dólar liberiano", Symbol: "LRD"}, "AOA": ut.Currency{Currency: "AOA", DisplayName: "kuanza", Symbol: "AOA"}, "IDR": ut.Currency{Currency: "IDR", DisplayName: "rupia indonesia", Symbol: "IDR"}, "XBA": ut.Currency{Currency: "XBA", DisplayName: "unidad compuesta europea", Symbol: ""}, "BEC": ut.Currency{Currency: "BEC", DisplayName: "franco belga (convertible)", Symbol: ""}, "MVR": ut.Currency{Currency: "MVR", DisplayName: "rufiya", Symbol: "MVR"}, "EEK": ut.Currency{Currency: "EEK", DisplayName: "corona estonia", Symbol: ""}, "BRB": ut.Currency{Currency: "BRB", DisplayName: "nuevo cruceiro brasileño (1967–1986)", Symbol: ""}, "HRK": ut.Currency{Currency: "HRK", DisplayName: "kuna", Symbol: "HRK"}, "CLP": ut.Currency{Currency: "CLP", DisplayName: "peso chileno", Symbol: "CLP"}, "BEF": ut.Currency{Currency: "BEF", DisplayName: "franco belga", Symbol: ""}, "XPT": ut.Currency{Currency: "XPT", DisplayName: "platino", Symbol: ""}, "EUR": ut.Currency{Currency: "EUR", DisplayName: "euro", Symbol: "€"}, "XFO": ut.Currency{Currency: "XFO", DisplayName: "franco oro francés", Symbol: ""}, "MAD": ut.Currency{Currency: "MAD", DisplayName: "dírham marroquí", Symbol: "MAD"}, "NOK": ut.Currency{Currency: "NOK", DisplayName: "corona noruega", Symbol: "NOK"}, "NAD": ut.Currency{Currency: "NAD", DisplayName: "dólar namibio", Symbol: "NAD"}, "YUN": ut.Currency{Currency: "YUN", DisplayName: "dinar convertible yugoslavo", Symbol: ""}, "ZRN": ut.Currency{Currency: "ZRN", DisplayName: "nuevo zaire zaireño", Symbol: ""}, "UGX": ut.Currency{Currency: "UGX", DisplayName: "chelín ugandés", Symbol: "UGX"}, "LYD": ut.Currency{Currency: "LYD", DisplayName: "dinar libio", Symbol: "LYD"}, "XBB": ut.Currency{Currency: "XBB", DisplayName: "unidad monetaria europea", Symbol: ""}, "NIC": ut.Currency{Currency: "NIC", DisplayName: "córdoba nicaragüense (1988–1991)", Symbol: ""}, "ARA": ut.Currency{Currency: "ARA", DisplayName: "austral argentino", Symbol: ""}, "SUR": ut.Currency{Currency: "SUR", DisplayName: "rublo soviético", Symbol: ""}, "AON": ut.Currency{Currency: "AON", DisplayName: "nuevo kwanza angoleño (1990–2000)", Symbol: ""}, "ZMW": ut.Currency{Currency: "ZMW", DisplayName: "kuacha zambiano", Symbol: "ZMW"}, "DZD": ut.Currency{Currency: "DZD", DisplayName: "dinar argelino", Symbol: "DZD"}, "AZM": ut.Currency{Currency: "AZM", DisplayName: "manat azerí (1993–2006)", Symbol: ""}, "TZS": ut.Currency{Currency: "TZS", DisplayName: "chelín tanzano", Symbol: "TZS"}, "KPW": ut.Currency{Currency: "KPW", DisplayName: "won norcoreano", Symbol: "KPW"}, "BIF": ut.Currency{Currency: "BIF", DisplayName: "franco burundés", Symbol: "BIF"}, "NLG": ut.Currency{Currency: "NLG", DisplayName: "florín neerlandés", Symbol: ""}, "HKD": ut.Currency{Currency: "HKD", DisplayName: "dólar hongkonés", Symbol: "HKD"}, "DEM": ut.Currency{Currency: "DEM", DisplayName: "marco alemán", Symbol: ""}, "KGS": ut.Currency{Currency: "KGS", DisplayName: "som", Symbol: "KGS"}, "YUM": ut.Currency{Currency: "YUM", DisplayName: "super dinar yugoslavo", Symbol: ""}, "PGK": ut.Currency{Currency: "PGK", DisplayName: "kina", Symbol: "PGK"}, "ESP": ut.Currency{Currency: "ESP", DisplayName: "peseta española", Symbol: "₧"}, "KRW": ut.Currency{Currency: "KRW", DisplayName: "won surcoreano", Symbol: "KRW"}, "HRD": ut.Currency{Currency: "HRD", DisplayName: "dinar croata", Symbol: ""}, "ITL": ut.Currency{Currency: "ITL", DisplayName: "lira italiana", Symbol: ""}}