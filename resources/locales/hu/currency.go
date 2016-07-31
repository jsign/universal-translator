package hu

import "github.com/go-playground/universal-translator"

var currencies = ut.CurrencyFormatValue{"ILS": ut.Currency{Currency: "ILS", DisplayName: "izraeli új sékel", Symbol: "ILS"}, "MAF": ut.Currency{Currency: "MAF", DisplayName: "Marokkói frank", Symbol: ""}, "YER": ut.Currency{Currency: "YER", DisplayName: "jemeni rial", Symbol: "YER"}, "ZRZ": ut.Currency{Currency: "ZRZ", DisplayName: "Zairei zaire", Symbol: ""}, "AFA": ut.Currency{Currency: "AFA", DisplayName: "afgán afghani (1927–2002)", Symbol: ""}, "BBD": ut.Currency{Currency: "BBD", DisplayName: "barbadosi dollár", Symbol: "BBD"}, "GEK": ut.Currency{Currency: "GEK", DisplayName: "Grúz kupon larit", Symbol: ""}, "TJR": ut.Currency{Currency: "TJR", DisplayName: "Tádzsikisztáni rubel", Symbol: ""}, "ZMK": ut.Currency{Currency: "ZMK", DisplayName: "Zambiai kwacha (1968–2012)", Symbol: "ZMK"}, "MDL": ut.Currency{Currency: "MDL", DisplayName: "moldován lei", Symbol: "MDL"}, "FRF": ut.Currency{Currency: "FRF", DisplayName: "Francia frank", Symbol: ""}, "HKD": ut.Currency{Currency: "HKD", DisplayName: "hongkongi dollár", Symbol: "HKD"}, "LUF": ut.Currency{Currency: "LUF", DisplayName: "Luxemburgi frank", Symbol: ""}, "ECV": ut.Currency{Currency: "ECV", DisplayName: "Ecuadori Unidad de Valor Constante (UVC)", Symbol: ""}, "AZN": ut.Currency{Currency: "AZN", DisplayName: "azerbajdzsáni manat", Symbol: "AZN"}, "CHE": ut.Currency{Currency: "CHE", DisplayName: "WIR euro", Symbol: ""}, "CLF": ut.Currency{Currency: "CLF", DisplayName: "Chilei unidades de fomento", Symbol: ""}, "SVC": ut.Currency{Currency: "SVC", DisplayName: "Salvadori colón", Symbol: ""}, "AWG": ut.Currency{Currency: "AWG", DisplayName: "arubai florin", Symbol: "AWG"}, "BDT": ut.Currency{Currency: "BDT", DisplayName: "bangladesi taka", Symbol: "BDT"}, "KMF": ut.Currency{Currency: "KMF", DisplayName: "comorei frank", Symbol: "KMF"}, "AOA": ut.Currency{Currency: "AOA", DisplayName: "angolai kwanza", Symbol: "AOA"}, "ZMW": ut.Currency{Currency: "ZMW", DisplayName: "zambiai kwacha", Symbol: "ZMW"}, "JMD": ut.Currency{Currency: "JMD", DisplayName: "jamaicai dollár", Symbol: "JMD"}, "LYD": ut.Currency{Currency: "LYD", DisplayName: "líbiai dínár", Symbol: "LYD"}, "PAB": ut.Currency{Currency: "PAB", DisplayName: "panamai balboa", Symbol: "PAB"}, "RUR": ut.Currency{Currency: "RUR", DisplayName: "orosz rubel (1991–1998)", Symbol: ""}, "AED": ut.Currency{Currency: "AED", DisplayName: "EAE-dirham", Symbol: "AED"}, "BND": ut.Currency{Currency: "BND", DisplayName: "brunei dollár", Symbol: "BND"}, "MAD": ut.Currency{Currency: "MAD", DisplayName: "marokkói dirham", Symbol: "MAD"}, "ARA": ut.Currency{Currency: "ARA", DisplayName: "Argentín austral", Symbol: ""}, "LTT": ut.Currency{Currency: "LTT", DisplayName: "Litvániai talonas", Symbol: ""}, "SUR": ut.Currency{Currency: "SUR", DisplayName: "Szovjet rubel", Symbol: ""}, "USS": ut.Currency{Currency: "USS", DisplayName: "USA dollár (aznapi)", Symbol: ""}, "GWP": ut.Currency{Currency: "GWP", DisplayName: "Guinea-Bissaui peso", Symbol: ""}, "OMR": ut.Currency{Currency: "OMR", DisplayName: "ománi rial", Symbol: "OMR"}, "RWF": ut.Currency{Currency: "RWF", DisplayName: "ruandai frank", Symbol: "RWF"}, "VEB": ut.Currency{Currency: "VEB", DisplayName: "Venezuelai bolivar (1871–2008)", Symbol: ""}, "SAR": ut.Currency{Currency: "SAR", DisplayName: "szaúdi riyal", Symbol: "SAR"}, "SKK": ut.Currency{Currency: "SKK", DisplayName: "Szlovák korona", Symbol: ""}, "TJS": ut.Currency{Currency: "TJS", DisplayName: "tádzsikisztáni somoni", Symbol: "TJS"}, "XAG": ut.Currency{Currency: "XAG", DisplayName: "Ezüst", Symbol: ""}, "XFU": ut.Currency{Currency: "XFU", DisplayName: "Francia UIC-frank", Symbol: ""}, "YUD": ut.Currency{Currency: "YUD", DisplayName: "Jugoszláv kemény dínár", Symbol: ""}, "KPW": ut.Currency{Currency: "KPW", DisplayName: "észak-koreai won", Symbol: "KPW"}, "AOK": ut.Currency{Currency: "AOK", DisplayName: "Angolai kwanza (1977–1990)", Symbol: ""}, "CAD": ut.Currency{Currency: "CAD", DisplayName: "kanadai dollár", Symbol: "CAD"}, "COU": ut.Currency{Currency: "COU", DisplayName: "Unidad de Valor Real", Symbol: ""}, "CLP": ut.Currency{Currency: "CLP", DisplayName: "chilei peso", Symbol: "CLP"}, "MXN": ut.Currency{Currency: "MXN", DisplayName: "mexikói peso", Symbol: "MXN"}, "NIO": ut.Currency{Currency: "NIO", DisplayName: "nicaraguai córdoba", Symbol: "NIO"}, "TND": ut.Currency{Currency: "TND", DisplayName: "tunéziai dínár", Symbol: "TND"}, "CRC": ut.Currency{Currency: "CRC", DisplayName: "Costa Rica-i colon", Symbol: "CRC"}, "MGA": ut.Currency{Currency: "MGA", DisplayName: "madagaszkári ariary", Symbol: "MGA"}, "MOP": ut.Currency{Currency: "MOP", DisplayName: "makaói pataca", Symbol: "MOP"}, "NZD": ut.Currency{Currency: "NZD", DisplayName: "új-zélandi dollár", Symbol: "NZD"}, "BOV": ut.Currency{Currency: "BOV", DisplayName: "Bolíviai mvdol", Symbol: ""}, "BZD": ut.Currency{Currency: "BZD", DisplayName: "belize-i dollár", Symbol: "BZD"}, "NAD": ut.Currency{Currency: "NAD", DisplayName: "namíbiai dollár", Symbol: "NAD"}, "CZK": ut.Currency{Currency: "CZK", DisplayName: "cseh korona", Symbol: "CZK"}, "EEK": ut.Currency{Currency: "EEK", DisplayName: "Észt korona", Symbol: ""}, "LBP": ut.Currency{Currency: "LBP", DisplayName: "libanoni font", Symbol: "LBP"}, "DKK": ut.Currency{Currency: "DKK", DisplayName: "dán korona", Symbol: "DKK"}, "PTE": ut.Currency{Currency: "PTE", DisplayName: "Portugál escudo", Symbol: ""}, "SCR": ut.Currency{Currency: "SCR", DisplayName: "seychelle-szigeteki rúpia", Symbol: "SCR"}, "SLL": ut.Currency{Currency: "SLL", DisplayName: "Sierra Leone-i leone", Symbol: "SLL"}, "VEF": ut.Currency{Currency: "VEF", DisplayName: "venezuelai bolivar", Symbol: "VEF"}, "BAD": ut.Currency{Currency: "BAD", DisplayName: "Bosznia-hercegovinai dínár (1992–1994)", Symbol: ""}, "BYR": ut.Currency{Currency: "BYR", DisplayName: "fehérorosz rubel", Symbol: "BYR"}, "CHF": ut.Currency{Currency: "CHF", DisplayName: "svájci frank", Symbol: "CHF"}, "XOF": ut.Currency{Currency: "XOF", DisplayName: "CFA frank BCEAO", Symbol: "CFA"}, "GNF": ut.Currency{Currency: "GNF", DisplayName: "guineai frank", Symbol: "GNF"}, "MUR": ut.Currency{Currency: "MUR", DisplayName: "mauritiusi rúpia", Symbol: "MUR"}, "PKR": ut.Currency{Currency: "PKR", DisplayName: "pakisztáni rúpia", Symbol: "PKR"}, "PYG": ut.Currency{Currency: "PYG", DisplayName: "paraguayi guarani", Symbol: "PYG"}, "YUN": ut.Currency{Currency: "YUN", DisplayName: "Jugoszláv konvertibilis dínár", Symbol: ""}, "HRD": ut.Currency{Currency: "HRD", DisplayName: "Horvát dínár", Symbol: ""}, "KRW": ut.Currency{Currency: "KRW", DisplayName: "dél-koreai won", Symbol: "KRW"}, "MZN": ut.Currency{Currency: "MZN", DisplayName: "mozambiki metikális", Symbol: "MZN"}, "RSD": ut.Currency{Currency: "RSD", DisplayName: "szerb dínár", Symbol: "RSD"}, "XBD": ut.Currency{Currency: "XBD", DisplayName: "Európai kontó egység (XBD)", Symbol: ""}, "AZM": ut.Currency{Currency: "AZM", DisplayName: "azerbajdzsáni manat (1993–2006)", Symbol: ""}, "BRR": ut.Currency{Currency: "BRR", DisplayName: "Brazil cruzeiro (1993–1994)", Symbol: ""}, "BTN": ut.Currency{Currency: "BTN", DisplayName: "bhutáni ngultrum", Symbol: "BTN"}, "ZAL": ut.Currency{Currency: "ZAL", DisplayName: "Dél-afrikai rand (pénzügyi)", Symbol: ""}, "CNY": ut.Currency{Currency: "CNY", DisplayName: "kínai jüan", Symbol: "CNY"}, "IDR": ut.Currency{Currency: "IDR", DisplayName: "indonéz rúpia", Symbol: "IDR"}, "XAF": ut.Currency{Currency: "XAF", DisplayName: "CFA frank BEAC", Symbol: "FCFA"}, "TTD": ut.Currency{Currency: "TTD", DisplayName: "Trinidad és Tobago-i dollár", Symbol: "TTD"}, "XBA": ut.Currency{Currency: "XBA", DisplayName: "Európai kompozit egység", Symbol: ""}, "ARS": ut.Currency{Currency: "ARS", DisplayName: "argentin peso", Symbol: "ARS"}, "MYR": ut.Currency{Currency: "MYR", DisplayName: "malajziai ringgit", Symbol: "MYR"}, "MZM": ut.Currency{Currency: "MZM", DisplayName: "Mozambik metical", Symbol: ""}, "CDF": ut.Currency{Currency: "CDF", DisplayName: "kongói frank", Symbol: "CDF"}, "HNL": ut.Currency{Currency: "HNL", DisplayName: "hodurasi lempira", Symbol: "HNL"}, "ILP": ut.Currency{Currency: "ILP", DisplayName: "Izraeli font", Symbol: ""}, "ESP": ut.Currency{Currency: "ESP", DisplayName: "Spanyol peseta", Symbol: ""}, "LAK": ut.Currency{Currency: "LAK", DisplayName: "laoszi kip", Symbol: "LAK"}, "MLF": ut.Currency{Currency: "MLF", DisplayName: "Mali frank", Symbol: ""}, "PES": ut.Currency{Currency: "PES", DisplayName: "Perui sol", Symbol: ""}, "XPF": ut.Currency{Currency: "XPF", DisplayName: "csendes-óceáni valutaközösségi frank", Symbol: "CFPF"}, "BEL": ut.Currency{Currency: "BEL", DisplayName: "Belga frank (pénzügyi)", Symbol: ""}, "BIF": ut.Currency{Currency: "BIF", DisplayName: "burundi frank", Symbol: "BIF"}, "EGP": ut.Currency{Currency: "EGP", DisplayName: "egyiptomi font", Symbol: "EGP"}, "MRO": ut.Currency{Currency: "MRO", DisplayName: "mauritániai ouguiya", Symbol: "MRO"}, "ZRN": ut.Currency{Currency: "ZRN", DisplayName: "Zairei új zaire", Symbol: ""}, "GBP": ut.Currency{Currency: "GBP", DisplayName: "brit font", Symbol: "GBP"}, "MWK": ut.Currency{Currency: "MWK", DisplayName: "malawi kwacha", Symbol: "MWK"}, "NOK": ut.Currency{Currency: "NOK", DisplayName: "norvég korona", Symbol: "NOK"}, "MTP": ut.Currency{Currency: "MTP", DisplayName: "Máltai font", Symbol: ""}, "SRD": ut.Currency{Currency: "SRD", DisplayName: "suriname-i dollár", Symbol: "SRD"}, "SSP": ut.Currency{Currency: "SSP", DisplayName: "dél-szudáni font", Symbol: "SSP"}, "AFN": ut.Currency{Currency: "AFN", DisplayName: "afgán afghani", Symbol: "AFN"}, "CSD": ut.Currency{Currency: "CSD", DisplayName: "szerb dinár", Symbol: ""}, "DJF": ut.Currency{Currency: "DJF", DisplayName: "dzsibuti frank", Symbol: "DJF"}, "HRK": ut.Currency{Currency: "HRK", DisplayName: "horvát kuna", Symbol: "HRK"}, "BHD": ut.Currency{Currency: "BHD", DisplayName: "bahreini dinár", Symbol: "BHD"}, "COP": ut.Currency{Currency: "COP", DisplayName: "kolumbiai peso", Symbol: "COP"}, "CYP": ut.Currency{Currency: "CYP", DisplayName: "Ciprusi font", Symbol: ""}, "USN": ut.Currency{Currency: "USN", DisplayName: "USA dollár (következő napi)", Symbol: ""}, "BRC": ut.Currency{Currency: "BRC", DisplayName: "Brazi cruzado (1986–1989)", Symbol: ""}, "GHC": ut.Currency{Currency: "GHC", DisplayName: "Ghánai cedi (1979–2007)", Symbol: ""}, "MMK": ut.Currency{Currency: "MMK", DisplayName: "mianmari kyat", Symbol: "MMK"}, "BGN": ut.Currency{Currency: "BGN", DisplayName: "bolgár új leva", Symbol: "BGN"}, "YUM": ut.Currency{Currency: "YUM", DisplayName: "Jugoszláv új dínár", Symbol: ""}, "ZWL": ut.Currency{Currency: "ZWL", DisplayName: "Zimbabwei dollár (2009)", Symbol: ""}, "PEI": ut.Currency{Currency: "PEI", DisplayName: "Perui inti", Symbol: ""}, "BEC": ut.Currency{Currency: "BEC", DisplayName: "Belga frank (konvertibilis)", Symbol: ""}, "BEF": ut.Currency{Currency: "BEF", DisplayName: "Belga frank", Symbol: ""}, "ITL": ut.Currency{Currency: "ITL", DisplayName: "Olasz líra", Symbol: ""}, "KYD": ut.Currency{Currency: "KYD", DisplayName: "kajmán-szigeteki dollár", Symbol: "KYD"}, "MTL": ut.Currency{Currency: "MTL", DisplayName: "Máltai líra", Symbol: ""}, "ANG": ut.Currency{Currency: "ANG", DisplayName: "holland antilláki forint", Symbol: "ANG"}, "GQE": ut.Currency{Currency: "GQE", DisplayName: "Egyenlítői-guineai ekwele guineana", Symbol: ""}, "GTQ": ut.Currency{Currency: "GTQ", DisplayName: "guatemalai quetzal", Symbol: "GTQ"}, "TOP": ut.Currency{Currency: "TOP", DisplayName: "tongai paanga", Symbol: "TOP"}, "BRE": ut.Currency{Currency: "BRE", DisplayName: "Brazil cruzeiro (1990–1993)", Symbol: ""}, "FIM": ut.Currency{Currency: "FIM", DisplayName: "Finn markka", Symbol: ""}, "TMM": ut.Currency{Currency: "TMM", DisplayName: "türkmenisztáni manat (1993–2009)", Symbol: ""}, "ESA": ut.Currency{Currency: "ESA", DisplayName: "spanyol peseta (A–kontó)", Symbol: ""}, "NPR": ut.Currency{Currency: "NPR", DisplayName: "nepáli rúpia", Symbol: "NPR"}, "SDP": ut.Currency{Currency: "SDP", DisplayName: "Szudáni font (1957–1998)", Symbol: ""}, "JOD": ut.Currency{Currency: "JOD", DisplayName: "jordániai dínár", Symbol: "JOD"}, "LSL": ut.Currency{Currency: "LSL", DisplayName: "Lesothoi loti", Symbol: ""}, "LUC": ut.Currency{Currency: "LUC", DisplayName: "luxemburgi konvertibilis frank", Symbol: ""}, "SZL": ut.Currency{Currency: "SZL", DisplayName: "szváziföldi lilangeni", Symbol: "SZL"}, "UAK": ut.Currency{Currency: "UAK", DisplayName: "Ukrán karbovanec", Symbol: ""}, "UYI": ut.Currency{Currency: "UYI", DisplayName: "Uruguayi peso en unidades indexadas", Symbol: ""}, "BAM": ut.Currency{Currency: "BAM", DisplayName: "bosznia-hercegovinai konvertibilis márka", Symbol: "BAM"}, "ERN": ut.Currency{Currency: "ERN", DisplayName: "eritreai nakfa", Symbol: "ERN"}, "NGN": ut.Currency{Currency: "NGN", DisplayName: "nigériai naira", Symbol: "NGN"}, "KES": ut.Currency{Currency: "KES", DisplayName: "kenyai shilling", Symbol: "KES"}, "PLZ": ut.Currency{Currency: "PLZ", DisplayName: "Lengyel zloty (1950–1995)", Symbol: ""}, "RHD": ut.Currency{Currency: "RHD", DisplayName: "rhodéziai dollár", Symbol: ""}, "SRG": ut.Currency{Currency: "SRG", DisplayName: "Suriname-i gulden", Symbol: ""}, "TRY": ut.Currency{Currency: "TRY", DisplayName: "török líra", Symbol: "TRY"}, "BYB": ut.Currency{Currency: "BYB", DisplayName: "Fehérorosz új rubel (1994–1999)", Symbol: ""}, "GEL": ut.Currency{Currency: "GEL", DisplayName: "grúz lari", Symbol: "GEL"}, "ISK": ut.Currency{Currency: "ISK", DisplayName: "izlandi korona", Symbol: "ISK"}, "XTS": ut.Currency{Currency: "XTS", DisplayName: "Tesztelési pénznemkód", Symbol: ""}, "ZWD": ut.Currency{Currency: "ZWD", DisplayName: "Zimbabwei dollár (1980–2008)", Symbol: ""}, "DEM": ut.Currency{Currency: "DEM", DisplayName: "Német márka", Symbol: ""}, "ETB": ut.Currency{Currency: "ETB", DisplayName: "etiópiai birr", Symbol: "ETB"}, "LVL": ut.Currency{Currency: "LVL", DisplayName: "lett lats", Symbol: "LVL"}, "TMT": ut.Currency{Currency: "TMT", DisplayName: "türkmenisztáni manat", Symbol: "TMT"}, "MXP": ut.Currency{Currency: "MXP", DisplayName: "Mexikói ezüst peso (1861–1992)", Symbol: ""}, "SEK": ut.Currency{Currency: "SEK", DisplayName: "svéd korona", Symbol: "SEK"}, "SOS": ut.Currency{Currency: "SOS", DisplayName: "szomáli shilling", Symbol: "SOS"}, "XDR": ut.Currency{Currency: "XDR", DisplayName: "Special Drawing Rights", Symbol: ""}, "YDD": ut.Currency{Currency: "YDD", DisplayName: "Jemeni dínár", Symbol: ""}, "BOP": ut.Currency{Currency: "BOP", DisplayName: "Bolíviai peso", Symbol: ""}, "DOP": ut.Currency{Currency: "DOP", DisplayName: "dominikai peso", Symbol: "DOP"}, "XCD": ut.Currency{Currency: "XCD", DisplayName: "kelet-karibi dollár", Symbol: "XCD"}, "ATS": ut.Currency{Currency: "ATS", DisplayName: "Osztrák schilling", Symbol: ""}, "KGS": ut.Currency{Currency: "KGS", DisplayName: "kirgizisztáni szom", Symbol: "KGS"}, "SIT": ut.Currency{Currency: "SIT", DisplayName: "Szlovén tolar", Symbol: ""}, "AON": ut.Currency{Currency: "AON", DisplayName: "Angolai új kwanza (1990–2000)", Symbol: ""}, "ESB": ut.Currency{Currency: "ESB", DisplayName: "spanyol peseta (konvertibilis kontó)", Symbol: ""}, "GIP": ut.Currency{Currency: "GIP", DisplayName: "gibraltári font", Symbol: "GIP"}, "THB": ut.Currency{Currency: "THB", DisplayName: "thai baht", Symbol: "THB"}, "UAH": ut.Currency{Currency: "UAH", DisplayName: "ukrán hrivnya", Symbol: "UAH"}, "WST": ut.Currency{Currency: "WST", DisplayName: "nyugat-szamoai tala", Symbol: "WST"}, "ADP": ut.Currency{Currency: "ADP", DisplayName: "Andorrai peseta", Symbol: ""}, "CHW": ut.Currency{Currency: "CHW", DisplayName: "WIR frank", Symbol: ""}, "MGF": ut.Currency{Currency: "MGF", DisplayName: "Madagaszkári frank", Symbol: ""}, "SDD": ut.Currency{Currency: "SDD", DisplayName: "Szudáni dínár (1992–2007)", Symbol: ""}, "SHP": ut.Currency{Currency: "SHP", DisplayName: "Szent Ilona-i font", Symbol: "SHP"}, "VUV": ut.Currency{Currency: "VUV", DisplayName: "vanuatui vatu", Symbol: "VUV"}, "AMD": ut.Currency{Currency: "AMD", DisplayName: "örmény dram", Symbol: "AMD"}, "DDM": ut.Currency{Currency: "DDM", DisplayName: "Kelet-Német márka", Symbol: ""}, "ROL": ut.Currency{Currency: "ROL", DisplayName: "román lej (1952–2006)", Symbol: ""}, "VND": ut.Currency{Currency: "VND", DisplayName: "vietnami dong", Symbol: "VND"}, "XRE": ut.Currency{Currency: "XRE", DisplayName: "RINET tőke", Symbol: ""}, "CUC": ut.Currency{Currency: "CUC", DisplayName: "kubai konvertibilis peso", Symbol: "CUC"}, "NIC": ut.Currency{Currency: "NIC", DisplayName: "Nikaraguai cordoba", Symbol: ""}, "UGX": ut.Currency{Currency: "UGX", DisplayName: "ugandai shilling", Symbol: "UGX"}, "PHP": ut.Currency{Currency: "PHP", DisplayName: "fülöp-szigeteki peso", Symbol: "PHP"}, "XEU": ut.Currency{Currency: "XEU", DisplayName: "európai pénznemegység", Symbol: ""}, "CUP": ut.Currency{Currency: "CUP", DisplayName: "kubai peso", Symbol: "CUP"}, "FKP": ut.Currency{Currency: "FKP", DisplayName: "falkland-szigeteki font", Symbol: "FKP"}, "GHS": ut.Currency{Currency: "GHS", DisplayName: "ghánai cedi", Symbol: "GHS"}, "AUD": ut.Currency{Currency: "AUD", DisplayName: "ausztrál dollár", Symbol: "AUD"}, "PEN": ut.Currency{Currency: "PEN", DisplayName: "perui sol nuevo", Symbol: "PEN"}, "XBB": ut.Currency{Currency: "XBB", DisplayName: "Európai monetáris egység", Symbol: ""}, "LTL": ut.Currency{Currency: "LTL", DisplayName: "litvániai litas", Symbol: "LTL"}, "STD": ut.Currency{Currency: "STD", DisplayName: "São Tomé és Príncipe-i dobra", Symbol: "STD"}, "BWP": ut.Currency{Currency: "BWP", DisplayName: "botswanai pula", Symbol: "BWP"}, "ECS": ut.Currency{Currency: "ECS", DisplayName: "Ecuadori sucre", Symbol: ""}, "LKR": ut.Currency{Currency: "LKR", DisplayName: "Srí Lanka-i rúpia", Symbol: "LKR"}, "JPY": ut.Currency{Currency: "JPY", DisplayName: "japán jen", Symbol: "¥"}, "KWD": ut.Currency{Currency: "KWD", DisplayName: "kuvaiti dínár", Symbol: "KWD"}, "XPT": ut.Currency{Currency: "XPT", DisplayName: "platina", Symbol: ""}, "BUK": ut.Currency{Currency: "BUK", DisplayName: "Burmai kyat", Symbol: ""}, "FJD": ut.Currency{Currency: "FJD", DisplayName: "fidzsi dollár", Symbol: "FJD"}, "HTG": ut.Currency{Currency: "HTG", DisplayName: "haiti gourde", Symbol: "HTG"}, "INR": ut.Currency{Currency: "INR", DisplayName: "indiai rúpia", Symbol: "INR"}, "LRD": ut.Currency{Currency: "LRD", DisplayName: "libériai dollár", Symbol: "LRD"}, "SGD": ut.Currency{Currency: "SGD", DisplayName: "szingapúri dollár", Symbol: "SGD"}, "TPE": ut.Currency{Currency: "TPE", DisplayName: "Timori escudo", Symbol: ""}, "ZWR": ut.Currency{Currency: "ZWR", DisplayName: "Zimbabwei dollár (2008)", Symbol: ""}, "BMD": ut.Currency{Currency: "BMD", DisplayName: "bermudai dollár", Symbol: "BMD"}, "GMD": ut.Currency{Currency: "GMD", DisplayName: "gambiai dalasi", Symbol: "GMD"}, "GYD": ut.Currency{Currency: "GYD", DisplayName: "guyanai dollár", Symbol: "GYD"}, "USD": ut.Currency{Currency: "USD", DisplayName: "USA-dollár", Symbol: "USD"}, "ARP": ut.Currency{Currency: "ARP", DisplayName: "Argentín peso (1983–1985)", Symbol: ""}, "BRL": ut.Currency{Currency: "BRL", DisplayName: "brazil real", Symbol: "BRL"}, "MKD": ut.Currency{Currency: "MKD", DisplayName: "macedon dínár", Symbol: "MKD"}, "TZS": ut.Currency{Currency: "TZS", DisplayName: "tanzániai shilling", Symbol: "TZS"}, "UYU": ut.Currency{Currency: "UYU", DisplayName: "uruguay-i peso", Symbol: "UYU"}, "GNS": ut.Currency{Currency: "GNS", DisplayName: "Guineai syli", Symbol: ""}, "MNT": ut.Currency{Currency: "MNT", DisplayName: "mongóliai tugrik", Symbol: "MNT"}, "MZE": ut.Currency{Currency: "MZE", DisplayName: "Mozambik escudo", Symbol: ""}, "ALL": ut.Currency{Currency: "ALL", DisplayName: "albán lek", Symbol: "ALL"}, "IQD": ut.Currency{Currency: "IQD", DisplayName: "iraki dínár", Symbol: "IQD"}, "UYP": ut.Currency{Currency: "UYP", DisplayName: "Uruguay-i peso (1975–1993)", Symbol: ""}, "XFO": ut.Currency{Currency: "XFO", DisplayName: "Francia arany frank", Symbol: ""}, "DZD": ut.Currency{Currency: "DZD", DisplayName: "algériai dínár", Symbol: "DZD"}, "NLG": ut.Currency{Currency: "NLG", DisplayName: "Holland forint", Symbol: ""}, "TWD": ut.Currency{Currency: "TWD", DisplayName: "tajvani új dollár", Symbol: "TWD"}, "ZAR": ut.Currency{Currency: "ZAR", DisplayName: "dél-afrikai rand", Symbol: "ZAR"}, "BSD": ut.Currency{Currency: "BSD", DisplayName: "bahamai dollár", Symbol: "BSD"}, "EUR": ut.Currency{Currency: "EUR", DisplayName: "euró", Symbol: "EUR"}, "SYP": ut.Currency{Currency: "SYP", DisplayName: "szíriai font", Symbol: "SYP"}, "SDG": ut.Currency{Currency: "SDG", DisplayName: "szudáni font", Symbol: "SDG"}, "UGS": ut.Currency{Currency: "UGS", DisplayName: "Ugandai shilling (1966–1987)", Symbol: ""}, "XAU": ut.Currency{Currency: "XAU", DisplayName: "Arany", Symbol: ""}, "XBC": ut.Currency{Currency: "XBC", DisplayName: "Európai kontó egység (XBC)", Symbol: ""}, "XXX": ut.Currency{Currency: "XXX", DisplayName: "ismeretlen pénznem", Symbol: ""}, "AOR": ut.Currency{Currency: "AOR", DisplayName: "Angolai kwanza reajustado (1995–1999)", Symbol: ""}, "CSK": ut.Currency{Currency: "CSK", DisplayName: "Csehszlovák kemény korona", Symbol: ""}, "RUB": ut.Currency{Currency: "RUB", DisplayName: "orosz rubel", Symbol: "RUB"}, "HUF": ut.Currency{Currency: "HUF", DisplayName: "magyar forint", Symbol: "Ft"}, "LUL": ut.Currency{Currency: "LUL", DisplayName: "luxemburgi pénzügyi frank", Symbol: ""}, "BGL": ut.Currency{Currency: "BGL", DisplayName: "Bolgár kemény leva", Symbol: ""}, "BRB": ut.Currency{Currency: "BRB", DisplayName: "Brazi cruzeiro novo (1967–1986)", Symbol: ""}, "BRN": ut.Currency{Currency: "BRN", DisplayName: "Brazil cruzado novo (1989–1990)", Symbol: ""}, "RON": ut.Currency{Currency: "RON", DisplayName: "román lej", Symbol: "RON"}, "UZS": ut.Currency{Currency: "UZS", DisplayName: "üzbegisztáni szum", Symbol: "UZS"}, "IRR": ut.Currency{Currency: "IRR", DisplayName: "iráni rial", Symbol: "IRR"}, "KHR": ut.Currency{Currency: "KHR", DisplayName: "kambodzsai riel", Symbol: "KHR"}, "PLN": ut.Currency{Currency: "PLN", DisplayName: "lengyel zloty", Symbol: "PLN"}, "TRL": ut.Currency{Currency: "TRL", DisplayName: "török líra (1922–2005)", Symbol: ""}, "KZT": ut.Currency{Currency: "KZT", DisplayName: "kazahsztáni tenge", Symbol: "KZT"}, "LVR": ut.Currency{Currency: "LVR", DisplayName: "Lett rubel", Symbol: ""}, "QAR": ut.Currency{Currency: "QAR", DisplayName: "katari rial", Symbol: "QAR"}, "CVE": ut.Currency{Currency: "CVE", DisplayName: "Cape Verde-i escudo", Symbol: "CVE"}, "GRD": ut.Currency{Currency: "GRD", DisplayName: "Görög drachma", Symbol: ""}, "SBD": ut.Currency{Currency: "SBD", DisplayName: "salamon-szigeteki dollár", Symbol: "SBD"}, "XPD": ut.Currency{Currency: "XPD", DisplayName: "palládium", Symbol: ""}, "BOB": ut.Currency{Currency: "BOB", DisplayName: "bolíviai boliviano", Symbol: "BOB"}, "IEP": ut.Currency{Currency: "IEP", DisplayName: "Ír font", Symbol: ""}, "MVR": ut.Currency{Currency: "MVR", DisplayName: "maldív-szigeteki rufiyaa", Symbol: "MVR"}, "GWE": ut.Currency{Currency: "GWE", DisplayName: "Portugál guinea escudo", Symbol: ""}, "MXV": ut.Currency{Currency: "MXV", DisplayName: "Mexikói Unidad de Inversion (UDI)", Symbol: ""}, "PGK": ut.Currency{Currency: "PGK", DisplayName: "pápua új-guineai kina", Symbol: "PGK"}}
