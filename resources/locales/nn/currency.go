package nn

import "github.com/go-playground/universal-translator"

var currencies = ut.CurrencyFormatValue{"GYD": ut.Currency{Currency: "GYD", DisplayName: "guyansk dollar", Symbol: ""}, "XFU": ut.Currency{Currency: "XFU", DisplayName: "fransk UIC-franc", Symbol: ""}, "AFA": ut.Currency{Currency: "AFA", DisplayName: "afghani (1927–2002)", Symbol: ""}, "MVR": ut.Currency{Currency: "MVR", DisplayName: "maldivisk rufiyaa", Symbol: ""}, "XAU": ut.Currency{Currency: "XAU", DisplayName: "gull", Symbol: ""}, "XOF": ut.Currency{Currency: "XOF", DisplayName: "CFA franc BCEAO", Symbol: ""}, "ZMK": ut.Currency{Currency: "ZMK", DisplayName: "zambisk kwacha (1968–2012)", Symbol: ""}, "AZN": ut.Currency{Currency: "AZN", DisplayName: "aserbajdsjansk manat", Symbol: ""}, "DDM": ut.Currency{Currency: "DDM", DisplayName: "austtysk mark", Symbol: ""}, "PYG": ut.Currency{Currency: "PYG", DisplayName: "paraguayansk guarani", Symbol: ""}, "TMM": ut.Currency{Currency: "TMM", DisplayName: "turkmensk manat", Symbol: ""}, "UAH": ut.Currency{Currency: "UAH", DisplayName: "ukrainsk hryvnia", Symbol: ""}, "UZS": ut.Currency{Currency: "UZS", DisplayName: "usbekisk sum", Symbol: ""}, "DZD": ut.Currency{Currency: "DZD", DisplayName: "algerisk dinar", Symbol: ""}, "HUF": ut.Currency{Currency: "HUF", DisplayName: "ungarsk forint", Symbol: ""}, "BRE": ut.Currency{Currency: "BRE", DisplayName: "brasiliansk cruzeiro (1990–1993)", Symbol: ""}, "CSD": ut.Currency{Currency: "CSD", DisplayName: "gammal serbisk dinar", Symbol: ""}, "LUL": ut.Currency{Currency: "LUL", DisplayName: "luxemburgsk finansiell franc", Symbol: ""}, "QAR": ut.Currency{Currency: "QAR", DisplayName: "qatarsk rial", Symbol: ""}, "TTD": ut.Currency{Currency: "TTD", DisplayName: "trinidadisk dollar", Symbol: ""}, "AFN": ut.Currency{Currency: "AFN", DisplayName: "afghani", Symbol: ""}, "ARP": ut.Currency{Currency: "ARP", DisplayName: "argentinsk peso (1983–1985)", Symbol: ""}, "GHC": ut.Currency{Currency: "GHC", DisplayName: "ghanesisk cedi (1979–2007)", Symbol: ""}, "LRD": ut.Currency{Currency: "LRD", DisplayName: "liberisk dollar", Symbol: ""}, "VEB": ut.Currency{Currency: "VEB", DisplayName: "venezuelansk bolivar (1871–2008)", Symbol: ""}, "LAK": ut.Currency{Currency: "LAK", DisplayName: "laotisk kip", Symbol: ""}, "MXV": ut.Currency{Currency: "MXV", DisplayName: "meksikansk unidad de inversion (UDI)", Symbol: ""}, "PLZ": ut.Currency{Currency: "PLZ", DisplayName: "polsk zloty (1950–1995)", Symbol: ""}, "SOS": ut.Currency{Currency: "SOS", DisplayName: "somalisk shilling", Symbol: ""}, "SVC": ut.Currency{Currency: "SVC", DisplayName: "salvadoransk colon", Symbol: ""}, "UYU": ut.Currency{Currency: "UYU", DisplayName: "uruguayansk peso", Symbol: ""}, "GEL": ut.Currency{Currency: "GEL", DisplayName: "georgisk lari", Symbol: ""}, "HRK": ut.Currency{Currency: "HRK", DisplayName: "kroatisk kuna", Symbol: ""}, "LKR": ut.Currency{Currency: "LKR", DisplayName: "srilankisk rupi", Symbol: ""}, "USS": ut.Currency{Currency: "USS", DisplayName: "amerikansk dollar (same dag)", Symbol: ""}, "VND": ut.Currency{Currency: "VND", DisplayName: "vietnamesisk dong", Symbol: ""}, "XDR": ut.Currency{Currency: "XDR", DisplayName: "spesielle trekkrettar", Symbol: ""}, "BIF": ut.Currency{Currency: "BIF", DisplayName: "burundisk franc", Symbol: ""}, "GWP": ut.Currency{Currency: "GWP", DisplayName: "Guinea-Bissau-peso", Symbol: ""}, "LTL": ut.Currency{Currency: "LTL", DisplayName: "litauisk lita", Symbol: ""}, "VEF": ut.Currency{Currency: "VEF", DisplayName: "venezuelansk bolivar", Symbol: ""}, "RSD": ut.Currency{Currency: "RSD", DisplayName: "serbisk dinar", Symbol: ""}, "SBD": ut.Currency{Currency: "SBD", DisplayName: "salomonsk dollar", Symbol: ""}, "AOA": ut.Currency{Currency: "AOA", DisplayName: "angolsk kwanza", Symbol: ""}, "ARA": ut.Currency{Currency: "ARA", DisplayName: "argentisk austral", Symbol: ""}, "CDF": ut.Currency{Currency: "CDF", DisplayName: "kongolesisk franc", Symbol: ""}, "GMD": ut.Currency{Currency: "GMD", DisplayName: "gambisk dalasi", Symbol: ""}, "MNT": ut.Currency{Currency: "MNT", DisplayName: "mongolsk tugrik", Symbol: ""}, "NGN": ut.Currency{Currency: "NGN", DisplayName: "nigeriansk naira", Symbol: ""}, "TMT": ut.Currency{Currency: "TMT", DisplayName: "turkmenistansk manat", Symbol: ""}, "XXX": ut.Currency{Currency: "XXX", DisplayName: "ukjend eller ugyldig valuta", Symbol: ""}, "BBD": ut.Currency{Currency: "BBD", DisplayName: "barbadisk dollar", Symbol: ""}, "BSD": ut.Currency{Currency: "BSD", DisplayName: "bahamisk dollar", Symbol: ""}, "GTQ": ut.Currency{Currency: "GTQ", DisplayName: "guatemalansk quetzal", Symbol: ""}, "LUF": ut.Currency{Currency: "LUF", DisplayName: "luxemburgsk franc", Symbol: ""}, "STD": ut.Currency{Currency: "STD", DisplayName: "Sao Tome og Principe-dobra", Symbol: ""}, "CUP": ut.Currency{Currency: "CUP", DisplayName: "kubansk peso", Symbol: ""}, "PTE": ut.Currency{Currency: "PTE", DisplayName: "portugisisk escudo", Symbol: ""}, "XAG": ut.Currency{Currency: "XAG", DisplayName: "sølv", Symbol: ""}, "DKK": ut.Currency{Currency: "DKK", DisplayName: "dansk krone", Symbol: ""}, "JPY": ut.Currency{Currency: "JPY", DisplayName: "japansk yen", Symbol: ""}, "MAD": ut.Currency{Currency: "MAD", DisplayName: "marokkansk dirham", Symbol: ""}, "MZN": ut.Currency{Currency: "MZN", DisplayName: "mosambikisk metical", Symbol: ""}, "BAM": ut.Currency{Currency: "BAM", DisplayName: "bosnisk-hercegovinsk mark (konvertibel)", Symbol: ""}, "LBP": ut.Currency{Currency: "LBP", DisplayName: "libanesisk pund", Symbol: ""}, "MTP": ut.Currency{Currency: "MTP", DisplayName: "maltesisk pund", Symbol: ""}, "XRE": ut.Currency{Currency: "XRE", DisplayName: "RINET-fond", Symbol: ""}, "ZWL": ut.Currency{Currency: "ZWL", DisplayName: "Zimbabwe-dollar (2009)", Symbol: ""}, "NPR": ut.Currency{Currency: "NPR", DisplayName: "nepalsk rupi", Symbol: ""}, "XCD": ut.Currency{Currency: "XCD", DisplayName: "austkaribisk dollar", Symbol: ""}, "XEU": ut.Currency{Currency: "XEU", DisplayName: "europeisk valutaeining", Symbol: ""}, "XPD": ut.Currency{Currency: "XPD", DisplayName: "palladium", Symbol: ""}, "ZMW": ut.Currency{Currency: "ZMW", DisplayName: "zambisk kwacha", Symbol: ""}, "ZRZ": ut.Currency{Currency: "ZRZ", DisplayName: "zairisk zaire", Symbol: ""}, "BRR": ut.Currency{Currency: "BRR", DisplayName: "brasiliansk cruzeiro", Symbol: ""}, "BZD": ut.Currency{Currency: "BZD", DisplayName: "belizisk dollar", Symbol: ""}, "ETB": ut.Currency{Currency: "ETB", DisplayName: "etiopisk birr", Symbol: ""}, "KGS": ut.Currency{Currency: "KGS", DisplayName: "kirgisisk som", Symbol: ""}, "SDG": ut.Currency{Currency: "SDG", DisplayName: "sudansk pund", Symbol: ""}, "ADP": ut.Currency{Currency: "ADP", DisplayName: "andorransk peseta", Symbol: ""}, "LVR": ut.Currency{Currency: "LVR", DisplayName: "latvisk rubel", Symbol: ""}, "ALL": ut.Currency{Currency: "ALL", DisplayName: "albansk lek", Symbol: ""}, "EUR": ut.Currency{Currency: "EUR", DisplayName: "euro", Symbol: ""}, "MXN": ut.Currency{Currency: "MXN", DisplayName: "meksikansk peso", Symbol: ""}, "SAR": ut.Currency{Currency: "SAR", DisplayName: "saudiarabisk rial", Symbol: ""}, "AOR": ut.Currency{Currency: "AOR", DisplayName: "angolsk kwanza reajustado (1995–1999)", Symbol: ""}, "CLF": ut.Currency{Currency: "CLF", DisplayName: "chilensk unidades de fomento", Symbol: ""}, "NOK": ut.Currency{Currency: "NOK", DisplayName: "norsk krone", Symbol: "kr"}, "ZAL": ut.Currency{Currency: "ZAL", DisplayName: "sørafrikansk rand (finansiell)", Symbol: ""}, "FRF": ut.Currency{Currency: "FRF", DisplayName: "fransk franc", Symbol: ""}, "VUV": ut.Currency{Currency: "VUV", DisplayName: "vanuatuisk vatu", Symbol: ""}, "CUC": ut.Currency{Currency: "CUC", DisplayName: "kubansk peso (konvertibel)", Symbol: ""}, "DEM": ut.Currency{Currency: "DEM", DisplayName: "tysk mark", Symbol: ""}, "FKP": ut.Currency{Currency: "FKP", DisplayName: "Falkland-pund", Symbol: ""}, "SRD": ut.Currency{Currency: "SRD", DisplayName: "surinamsk dollar", Symbol: ""}, "AWG": ut.Currency{Currency: "AWG", DisplayName: "arubisk gylden", Symbol: ""}, "CSK": ut.Currency{Currency: "CSK", DisplayName: "tsjekkoslovakisk koruna (hard)", Symbol: ""}, "MGA": ut.Currency{Currency: "MGA", DisplayName: "madagassisk ariary", Symbol: ""}, "PES": ut.Currency{Currency: "PES", DisplayName: "peruansk sol", Symbol: ""}, "SDP": ut.Currency{Currency: "SDP", DisplayName: "gammalt sudanesisk pund", Symbol: ""}, "BEC": ut.Currency{Currency: "BEC", DisplayName: "belgisk franc (konvertibel)", Symbol: ""}, "CLP": ut.Currency{Currency: "CLP", DisplayName: "chilensk peso", Symbol: ""}, "TRY": ut.Currency{Currency: "TRY", DisplayName: "tyrkisk lire", Symbol: ""}, "AMD": ut.Currency{Currency: "AMD", DisplayName: "armensk dram", Symbol: ""}, "ILS": ut.Currency{Currency: "ILS", DisplayName: "israelsk ny shekel", Symbol: ""}, "LTT": ut.Currency{Currency: "LTT", DisplayName: "litauisk talona", Symbol: ""}, "MXP": ut.Currency{Currency: "MXP", DisplayName: "meksikansk sølvpeso (1861–1992)", Symbol: ""}, "UYI": ut.Currency{Currency: "UYI", DisplayName: "uruguayansk peso en unidades indexadas", Symbol: ""}, "JMD": ut.Currency{Currency: "JMD", DisplayName: "jamaikansk dollar", Symbol: ""}, "YUN": ut.Currency{Currency: "YUN", DisplayName: "jugoslavisk konvertibel dinar", Symbol: ""}, "BAD": ut.Currency{Currency: "BAD", DisplayName: "bosnisk-hercegovinsk dinar", Symbol: ""}, "CVE": ut.Currency{Currency: "CVE", DisplayName: "kappverdisk escudo", Symbol: ""}, "GIP": ut.Currency{Currency: "GIP", DisplayName: "gibraltarsk pund", Symbol: ""}, "KPW": ut.Currency{Currency: "KPW", DisplayName: "nordkoreansk won", Symbol: ""}, "MAF": ut.Currency{Currency: "MAF", DisplayName: "marokkansk franc", Symbol: ""}, "IEP": ut.Currency{Currency: "IEP", DisplayName: "irsk pund", Symbol: ""}, "PAB": ut.Currency{Currency: "PAB", DisplayName: "panamansk balboa", Symbol: ""}, "BHD": ut.Currency{Currency: "BHD", DisplayName: "bahrainsk dinar", Symbol: ""}, "ESP": ut.Currency{Currency: "ESP", DisplayName: "spansk peseta", Symbol: ""}, "MDL": ut.Currency{Currency: "MDL", DisplayName: "moldovsk leu", Symbol: ""}, "MGF": ut.Currency{Currency: "MGF", DisplayName: "madagassisk franc", Symbol: ""}, "SUR": ut.Currency{Currency: "SUR", DisplayName: "sovjetisk rubel", Symbol: ""}, "PEN": ut.Currency{Currency: "PEN", DisplayName: "peruansk nuevo sol", Symbol: ""}, "RON": ut.Currency{Currency: "RON", DisplayName: "rumensk leu", Symbol: ""}, "YUD": ut.Currency{Currency: "YUD", DisplayName: "jugoslavisk dinar (hard)", Symbol: ""}, "RUR": ut.Currency{Currency: "RUR", DisplayName: "russisk rubel (1991–1998)", Symbol: ""}, "TWD": ut.Currency{Currency: "TWD", DisplayName: "taiwansk ny dollar", Symbol: ""}, "SCR": ut.Currency{Currency: "SCR", DisplayName: "seychellisk rupi", Symbol: ""}, "SGD": ut.Currency{Currency: "SGD", DisplayName: "singaporsk dollar", Symbol: ""}, "SLL": ut.Currency{Currency: "SLL", DisplayName: "sierraleonsk leone", Symbol: ""}, "GEK": ut.Currency{Currency: "GEK", DisplayName: "georgisk kupon larit", Symbol: ""}, "ROL": ut.Currency{Currency: "ROL", DisplayName: "gammal rumensk leu", Symbol: ""}, "TJR": ut.Currency{Currency: "TJR", DisplayName: "tadsjikisk rubel", Symbol: ""}, "XPT": ut.Currency{Currency: "XPT", DisplayName: "platina", Symbol: ""}, "BRN": ut.Currency{Currency: "BRN", DisplayName: "brasiliansk cruzado novo", Symbol: ""}, "CHF": ut.Currency{Currency: "CHF", DisplayName: "sveitsisk franc", Symbol: ""}, "IQD": ut.Currency{Currency: "IQD", DisplayName: "iraksk dinar", Symbol: ""}, "RUB": ut.Currency{Currency: "RUB", DisplayName: "russisk rubel", Symbol: ""}, "UGX": ut.Currency{Currency: "UGX", DisplayName: "ugandisk shilling", Symbol: ""}, "XPF": ut.Currency{Currency: "XPF", DisplayName: "CFP franc", Symbol: ""}, "ZWD": ut.Currency{Currency: "ZWD", DisplayName: "zimbabwisk dollar", Symbol: ""}, "BGL": ut.Currency{Currency: "BGL", DisplayName: "bulgarsk hard lev", Symbol: ""}, "GHS": ut.Currency{Currency: "GHS", DisplayName: "ghanesisk cedi", Symbol: ""}, "KES": ut.Currency{Currency: "KES", DisplayName: "kenyansk shilling", Symbol: ""}, "NLG": ut.Currency{Currency: "NLG", DisplayName: "nederlandsk gylden", Symbol: ""}, "TOP": ut.Currency{Currency: "TOP", DisplayName: "tongansk paʻanga", Symbol: ""}, "XAF": ut.Currency{Currency: "XAF", DisplayName: "CFA franc BEAC", Symbol: ""}, "MMK": ut.Currency{Currency: "MMK", DisplayName: "myanmarsk kyat", Symbol: ""}, "TZS": ut.Currency{Currency: "TZS", DisplayName: "tanzaniansk shilling", Symbol: ""}, "AOK": ut.Currency{Currency: "AOK", DisplayName: "angolsk kwanza (1977–1990)", Symbol: ""}, "ATS": ut.Currency{Currency: "ATS", DisplayName: "austerriksk schilling", Symbol: ""}, "BND": ut.Currency{Currency: "BND", DisplayName: "bruneisk dollar", Symbol: ""}, "BRL": ut.Currency{Currency: "BRL", DisplayName: "brasiliansk real", Symbol: ""}, "HKD": ut.Currency{Currency: "HKD", DisplayName: "Hongkong-dollar", Symbol: ""}, "KMF": ut.Currency{Currency: "KMF", DisplayName: "komorisk franc", Symbol: ""}, "BYB": ut.Currency{Currency: "BYB", DisplayName: "kviterussisk ny rubel (1994–1999)", Symbol: ""}, "CHW": ut.Currency{Currency: "CHW", DisplayName: "WIR franc", Symbol: ""}, "SHP": ut.Currency{Currency: "SHP", DisplayName: "sankthelensk pund", Symbol: ""}, "SKK": ut.Currency{Currency: "SKK", DisplayName: "slovakisk koruna", Symbol: ""}, "XBC": ut.Currency{Currency: "XBC", DisplayName: "europeisk kontoeining (XBC)", Symbol: ""}, "ARS": ut.Currency{Currency: "ARS", DisplayName: "argentinsk peso", Symbol: ""}, "CHE": ut.Currency{Currency: "CHE", DisplayName: "WIR euro", Symbol: ""}, "CYP": ut.Currency{Currency: "CYP", DisplayName: "kypriotisk pund", Symbol: ""}, "HRD": ut.Currency{Currency: "HRD", DisplayName: "kroatisk dinar", Symbol: ""}, "KHR": ut.Currency{Currency: "KHR", DisplayName: "kambodsjansk riel", Symbol: ""}, "SEK": ut.Currency{Currency: "SEK", DisplayName: "svensk krone", Symbol: ""}, "KYD": ut.Currency{Currency: "KYD", DisplayName: "caymansk dollar", Symbol: ""}, "MZM": ut.Currency{Currency: "MZM", DisplayName: "gammal mosambikisk metical", Symbol: ""}, "SZL": ut.Currency{Currency: "SZL", DisplayName: "swazilandsk lilangeni", Symbol: ""}, "BDT": ut.Currency{Currency: "BDT", DisplayName: "bangladeshisk taka", Symbol: ""}, "BRC": ut.Currency{Currency: "BRC", DisplayName: "brasiliansk cruzado", Symbol: ""}, "GNF": ut.Currency{Currency: "GNF", DisplayName: "guineansk franc", Symbol: ""}, "HTG": ut.Currency{Currency: "HTG", DisplayName: "haitisk gourde", Symbol: ""}, "ITL": ut.Currency{Currency: "ITL", DisplayName: "italiensk lire", Symbol: ""}, "BGN": ut.Currency{Currency: "BGN", DisplayName: "bulgarsk ny lev", Symbol: ""}, "LSL": ut.Currency{Currency: "LSL", DisplayName: "lesothisk loti", Symbol: ""}, "MLF": ut.Currency{Currency: "MLF", DisplayName: "malisk franc", Symbol: ""}, "ZRN": ut.Currency{Currency: "ZRN", DisplayName: "zairisk ny zaire", Symbol: ""}, "MYR": ut.Currency{Currency: "MYR", DisplayName: "malaysisk ringgit", Symbol: ""}, "NAD": ut.Currency{Currency: "NAD", DisplayName: "namibisk dollar", Symbol: ""}, "PHP": ut.Currency{Currency: "PHP", DisplayName: "filippinsk peso", Symbol: ""}, "ZAR": ut.Currency{Currency: "ZAR", DisplayName: "sørafrikansk rand", Symbol: ""}, "BEF": ut.Currency{Currency: "BEF", DisplayName: "belgisk franc", Symbol: ""}, "CZK": ut.Currency{Currency: "CZK", DisplayName: "tsjekkisk koruna", Symbol: ""}, "GNS": ut.Currency{Currency: "GNS", DisplayName: "guineansk syli", Symbol: ""}, "SRG": ut.Currency{Currency: "SRG", DisplayName: "surinamsk gylden", Symbol: ""}, "TRL": ut.Currency{Currency: "TRL", DisplayName: "gammal tyrkiske lire", Symbol: ""}, "XBB": ut.Currency{Currency: "XBB", DisplayName: "europeisk monetær eining", Symbol: ""}, "NIC": ut.Currency{Currency: "NIC", DisplayName: "nicaraguansk cordoba", Symbol: ""}, "XBD": ut.Currency{Currency: "XBD", DisplayName: "europeisk kontoeining (XBD)", Symbol: ""}, "ISK": ut.Currency{Currency: "ISK", DisplayName: "islandsk krone", Symbol: ""}, "KRW": ut.Currency{Currency: "KRW", DisplayName: "sørkoreansk won", Symbol: ""}, "MOP": ut.Currency{Currency: "MOP", DisplayName: "makaosk pataca", Symbol: ""}, "MTL": ut.Currency{Currency: "MTL", DisplayName: "maltesisk lira", Symbol: ""}, "MWK": ut.Currency{Currency: "MWK", DisplayName: "malawisk kwacha", Symbol: ""}, "MZE": ut.Currency{Currency: "MZE", DisplayName: "mosambikisk escudo", Symbol: ""}, "BMD": ut.Currency{Currency: "BMD", DisplayName: "bermudisk dollar", Symbol: ""}, "CRC": ut.Currency{Currency: "CRC", DisplayName: "costaricansk colon", Symbol: ""}, "ECS": ut.Currency{Currency: "ECS", DisplayName: "ecuadoriansk sucre", Symbol: ""}, "EGP": ut.Currency{Currency: "EGP", DisplayName: "egyptisk pund", Symbol: ""}, "HNL": ut.Currency{Currency: "HNL", DisplayName: "honduransk lempira", Symbol: ""}, "SDD": ut.Currency{Currency: "SDD", DisplayName: "gammal sudanesisk dinar", Symbol: ""}, "EEK": ut.Currency{Currency: "EEK", DisplayName: "estisk kroon", Symbol: ""}, "KZT": ut.Currency{Currency: "KZT", DisplayName: "kasakhstansk tenge", Symbol: ""}, "XBA": ut.Currency{Currency: "XBA", DisplayName: "europeisk samansett eining", Symbol: ""}, "CNY": ut.Currency{Currency: "CNY", DisplayName: "kinesisk yuan renminbi", Symbol: ""}, "GBP": ut.Currency{Currency: "GBP", DisplayName: "britisk pund", Symbol: ""}, "JOD": ut.Currency{Currency: "JOD", DisplayName: "jordansk dinar", Symbol: ""}, "MUR": ut.Currency{Currency: "MUR", DisplayName: "mauritansk rupi", Symbol: ""}, "AON": ut.Currency{Currency: "AON", DisplayName: "angolsk ny kwanza (1990–2000)", Symbol: ""}, "BUK": ut.Currency{Currency: "BUK", DisplayName: "burmesisk kyat", Symbol: ""}, "BYR": ut.Currency{Currency: "BYR", DisplayName: "kviterussisk rubel", Symbol: ""}, "GRD": ut.Currency{Currency: "GRD", DisplayName: "gresk drakme", Symbol: ""}, "IRR": ut.Currency{Currency: "IRR", DisplayName: "iransk rial", Symbol: ""}, "PGK": ut.Currency{Currency: "PGK", DisplayName: "papuansk kina", Symbol: ""}, "ILP": ut.Currency{Currency: "ILP", DisplayName: "israelsk pund", Symbol: ""}, "TPE": ut.Currency{Currency: "TPE", DisplayName: "timoresisk escudo", Symbol: ""}, "USN": ut.Currency{Currency: "USN", DisplayName: "amerikansk dollar (neste dag)", Symbol: ""}, "WST": ut.Currency{Currency: "WST", DisplayName: "vestsamoisk tala", Symbol: ""}, "XFO": ut.Currency{Currency: "XFO", DisplayName: "fransk gullfranc", Symbol: ""}, "XTS": ut.Currency{Currency: "XTS", DisplayName: "testvalutakode", Symbol: ""}, "SIT": ut.Currency{Currency: "SIT", DisplayName: "slovensk tolar", Symbol: ""}, "UYP": ut.Currency{Currency: "UYP", DisplayName: "uruguayansk peso (1975–1993)", Symbol: ""}, "AED": ut.Currency{Currency: "AED", DisplayName: "UAE dirham", Symbol: ""}, "SYP": ut.Currency{Currency: "SYP", DisplayName: "syrisk pund", Symbol: ""}, "TJS": ut.Currency{Currency: "TJS", DisplayName: "tadsjikisk somoni", Symbol: ""}, "YUM": ut.Currency{Currency: "YUM", DisplayName: "jugoslavisk noviy-dinar", Symbol: ""}, "PEI": ut.Currency{Currency: "PEI", DisplayName: "peruansk inti", Symbol: ""}, "BOP": ut.Currency{Currency: "BOP", DisplayName: "bolivisk peso", Symbol: ""}, "BRB": ut.Currency{Currency: "BRB", DisplayName: "brasiliansk cruzeiro novo (1967–1986)", Symbol: ""}, "ESB": ut.Currency{Currency: "ESB", DisplayName: "spansk peseta (konvertibel konto)", Symbol: ""}, "LVL": ut.Currency{Currency: "LVL", DisplayName: "latvisk lat", Symbol: ""}, "OMR": ut.Currency{Currency: "OMR", DisplayName: "omansk rial", Symbol: ""}, "PLN": ut.Currency{Currency: "PLN", DisplayName: "polsk zloty", Symbol: ""}, "BTN": ut.Currency{Currency: "BTN", DisplayName: "bhutansk ngultrum", Symbol: ""}, "GQE": ut.Currency{Currency: "GQE", DisplayName: "ekvatorialguineansk ekwele guineana", Symbol: ""}, "KWD": ut.Currency{Currency: "KWD", DisplayName: "kuwaitisk dinar", Symbol: ""}, "UAK": ut.Currency{Currency: "UAK", DisplayName: "ukrainsk karbovanetz", Symbol: ""}, "BWP": ut.Currency{Currency: "BWP", DisplayName: "botswansk pula", Symbol: ""}, "DJF": ut.Currency{Currency: "DJF", DisplayName: "djiboutisk franc", Symbol: ""}, "FJD": ut.Currency{Currency: "FJD", DisplayName: "fijiansk dollar", Symbol: ""}, "UGS": ut.Currency{Currency: "UGS", DisplayName: "ugandisk shilling (1966–1987)", Symbol: ""}, "BOV": ut.Currency{Currency: "BOV", DisplayName: "bolivisk mvdol", Symbol: ""}, "GWE": ut.Currency{Currency: "GWE", DisplayName: "portugisisk guinea escudo", Symbol: ""}, "AZM": ut.Currency{Currency: "AZM", DisplayName: "aserbaijansk manat", Symbol: ""}, "COP": ut.Currency{Currency: "COP", DisplayName: "colombiansk peso", Symbol: ""}, "PKR": ut.Currency{Currency: "PKR", DisplayName: "pakistansk rupi", Symbol: ""}, "USD": ut.Currency{Currency: "USD", DisplayName: "amerikansk dollar", Symbol: ""}, "ANG": ut.Currency{Currency: "ANG", DisplayName: "nederlansk antillegylden", Symbol: ""}, "AUD": ut.Currency{Currency: "AUD", DisplayName: "australsk dollar", Symbol: ""}, "COU": ut.Currency{Currency: "COU", DisplayName: "unidad de valor real", Symbol: ""}, "DOP": ut.Currency{Currency: "DOP", DisplayName: "dominikansk peso", Symbol: ""}, "ECV": ut.Currency{Currency: "ECV", DisplayName: "ecuadoriansk unidad de valor constante (UVC)", Symbol: ""}, "ERN": ut.Currency{Currency: "ERN", DisplayName: "eritreisk nakfa", Symbol: ""}, "YER": ut.Currency{Currency: "YER", DisplayName: "jemenittisk rial", Symbol: ""}, "BEL": ut.Currency{Currency: "BEL", DisplayName: "belgisk franc (finansiell)", Symbol: ""}, "BOB": ut.Currency{Currency: "BOB", DisplayName: "boliviano", Symbol: ""}, "FIM": ut.Currency{Currency: "FIM", DisplayName: "finsk mark", Symbol: ""}, "RHD": ut.Currency{Currency: "RHD", DisplayName: "rhodesisk dollar", Symbol: ""}, "THB": ut.Currency{Currency: "THB", DisplayName: "thailandsk baht", Symbol: ""}, "TND": ut.Currency{Currency: "TND", DisplayName: "tunisisk dinar", Symbol: ""}, "ESA": ut.Currency{Currency: "ESA", DisplayName: "spansk peseta (A–konto)", Symbol: ""}, "INR": ut.Currency{Currency: "INR", DisplayName: "indisk rupi", Symbol: ""}, "LUC": ut.Currency{Currency: "LUC", DisplayName: "luxemburgsk konvertibel franc", Symbol: ""}, "CAD": ut.Currency{Currency: "CAD", DisplayName: "kanadisk dollar", Symbol: ""}, "IDR": ut.Currency{Currency: "IDR", DisplayName: "indonesisk rupi", Symbol: ""}, "MRO": ut.Currency{Currency: "MRO", DisplayName: "mauritansk ouguiya", Symbol: ""}, "NZD": ut.Currency{Currency: "NZD", DisplayName: "new zealandsk dollar", Symbol: ""}, "YDD": ut.Currency{Currency: "YDD", DisplayName: "jemenittisk dinar", Symbol: ""}, "LYD": ut.Currency{Currency: "LYD", DisplayName: "libysk dinar", Symbol: ""}, "MKD": ut.Currency{Currency: "MKD", DisplayName: "makedonsk denar", Symbol: ""}, "NIO": ut.Currency{Currency: "NIO", DisplayName: "nicaraguansk cordoba oro", Symbol: ""}, "RWF": ut.Currency{Currency: "RWF", DisplayName: "rwandisk franc", Symbol: ""}}