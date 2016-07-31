package uk

import "github.com/go-playground/universal-translator"

var currencies = ut.CurrencyFormatValue{"XBC": ut.Currency{Currency: "XBC", DisplayName: "європейська розрахункова одиниця XBC", Symbol: ""}, "DJF": ut.Currency{Currency: "DJF", DisplayName: "джибутійський франк", Symbol: "DJF"}, "LBP": ut.Currency{Currency: "LBP", DisplayName: "ліванський фунт", Symbol: "LBP"}, "NIC": ut.Currency{Currency: "NIC", DisplayName: "нікарагуанська кордоба (1988–1991)", Symbol: ""}, "CSD": ut.Currency{Currency: "CSD", DisplayName: "старий сербський динар", Symbol: ""}, "TTD": ut.Currency{Currency: "TTD", DisplayName: "долар Тринідаду і Тобаго", Symbol: "TTD"}, "YUD": ut.Currency{Currency: "YUD", DisplayName: "югославський твердий динар", Symbol: ""}, "AFN": ut.Currency{Currency: "AFN", DisplayName: "афганський афгані", Symbol: "AFN"}, "BEC": ut.Currency{Currency: "BEC", DisplayName: "бельгійський франк (конвертований)", Symbol: ""}, "BTN": ut.Currency{Currency: "BTN", DisplayName: "бутанський нгултрум", Symbol: "BTN"}, "GMD": ut.Currency{Currency: "GMD", DisplayName: "гамбійський даласі", Symbol: "GMD"}, "KGS": ut.Currency{Currency: "KGS", DisplayName: "киргизький сом", Symbol: "KGS"}, "PGK": ut.Currency{Currency: "PGK", DisplayName: "кіна Папуа Нової Гвінеї", Symbol: "PGK"}, "RWF": ut.Currency{Currency: "RWF", DisplayName: "руандійський франк", Symbol: "RWF"}, "THB": ut.Currency{Currency: "THB", DisplayName: "таїландський бат", Symbol: "THB"}, "BDT": ut.Currency{Currency: "BDT", DisplayName: "бангладеська така", Symbol: "BDT"}, "ERN": ut.Currency{Currency: "ERN", DisplayName: "еритрейська накфа", Symbol: "ERN"}, "GNS": ut.Currency{Currency: "GNS", DisplayName: "гвінейське сілі", Symbol: ""}, "GIP": ut.Currency{Currency: "GIP", DisplayName: "гібралтарський фунт", Symbol: "GIP"}, "PLZ": ut.Currency{Currency: "PLZ", DisplayName: "польський злотий (1950–1995)", Symbol: ""}, "AON": ut.Currency{Currency: "AON", DisplayName: "ангольська нова кванза (1990–2000)", Symbol: ""}, "BYR": ut.Currency{Currency: "BYR", DisplayName: "білоруський рубль", Symbol: "BYR"}, "GHC": ut.Currency{Currency: "GHC", DisplayName: "ганський седі (1979–2007)", Symbol: ""}, "SGD": ut.Currency{Currency: "SGD", DisplayName: "сінгапурський долар", Symbol: "SGD"}, "XBD": ut.Currency{Currency: "XBD", DisplayName: "європейська розрахункова одиниця XBD", Symbol: ""}, "ZMW": ut.Currency{Currency: "ZMW", DisplayName: "замбійська квача", Symbol: "ZMW"}, "COU": ut.Currency{Currency: "COU", DisplayName: "одиниця реальної вартості", Symbol: ""}, "KMF": ut.Currency{Currency: "KMF", DisplayName: "коморський франк", Symbol: "KMF"}, "MKD": ut.Currency{Currency: "MKD", DisplayName: "македонський денар", Symbol: "MKD"}, "SCR": ut.Currency{Currency: "SCR", DisplayName: "сейшельська рупія", Symbol: "SCR"}, "ILS": ut.Currency{Currency: "ILS", DisplayName: "ізраїльський новий шекель", Symbol: "ILS"}, "NOK": ut.Currency{Currency: "NOK", DisplayName: "норвезька крона", Symbol: "NOK"}, "DZD": ut.Currency{Currency: "DZD", DisplayName: "алжирський динар", Symbol: "DZD"}, "ETB": ut.Currency{Currency: "ETB", DisplayName: "ефіопський бир", Symbol: "ETB"}, "PKR": ut.Currency{Currency: "PKR", DisplayName: "пакистанська рупія", Symbol: "PKR"}, "TOP": ut.Currency{Currency: "TOP", DisplayName: "тонганська паанга", Symbol: "TOP"}, "BBD": ut.Currency{Currency: "BBD", DisplayName: "барбадоський долар", Symbol: "BBD"}, "HTG": ut.Currency{Currency: "HTG", DisplayName: "гаїтянський гурд", Symbol: "HTG"}, "NAD": ut.Currency{Currency: "NAD", DisplayName: "намібійський долар", Symbol: "NAD"}, "KRW": ut.Currency{Currency: "KRW", DisplayName: "південнокорейський вон", Symbol: "KRW"}, "MZM": ut.Currency{Currency: "MZM", DisplayName: "старий мозамбіцький метикал", Symbol: ""}, "IRR": ut.Currency{Currency: "IRR", DisplayName: "іранський ріал", Symbol: "IRR"}, "SEK": ut.Currency{Currency: "SEK", DisplayName: "шведська крона", Symbol: "SEK"}, "SHP": ut.Currency{Currency: "SHP", DisplayName: "фунт острова Святої Єлени", Symbol: "SHP"}, "STD": ut.Currency{Currency: "STD", DisplayName: "добра Сан-Томе і Принсіпі", Symbol: "STD"}, "XFO": ut.Currency{Currency: "XFO", DisplayName: "французький золотий франк", Symbol: ""}, "AZN": ut.Currency{Currency: "AZN", DisplayName: "азербайджанський манат", Symbol: "AZN"}, "GNF": ut.Currency{Currency: "GNF", DisplayName: "гвінейський франк", Symbol: "GNF"}, "INR": ut.Currency{Currency: "INR", DisplayName: "індійська рупія", Symbol: "INR"}, "MTL": ut.Currency{Currency: "MTL", DisplayName: "мальтійська ліра", Symbol: ""}, "SDP": ut.Currency{Currency: "SDP", DisplayName: "старий суданський фунт", Symbol: ""}, "TJS": ut.Currency{Currency: "TJS", DisplayName: "таджицький сомоні", Symbol: "TJS"}, "XRE": ut.Currency{Currency: "XRE", DisplayName: "фонди RINET", Symbol: ""}, "ARP": ut.Currency{Currency: "ARP", DisplayName: "аргентинський песо (1983–1985)", Symbol: ""}, "BRL": ut.Currency{Currency: "BRL", DisplayName: "бразильський реал", Symbol: "BRL"}, "CLF": ut.Currency{Currency: "CLF", DisplayName: "чилійський юнідадес де фоменто", Symbol: ""}, "MMK": ut.Currency{Currency: "MMK", DisplayName: "кʼят Мʼянми", Symbol: "MMK"}, "MOP": ut.Currency{Currency: "MOP", DisplayName: "патака Макао", Symbol: "MOP"}, "TRY": ut.Currency{Currency: "TRY", DisplayName: "турецька ліра", Symbol: "TRY"}, "VUV": ut.Currency{Currency: "VUV", DisplayName: "вануатський вату", Symbol: "VUV"}, "XBA": ut.Currency{Currency: "XBA", DisplayName: "європейська складена валютна одиниця", Symbol: ""}, "ARS": ut.Currency{Currency: "ARS", DisplayName: "аргентинський песо", Symbol: "ARS"}, "ATS": ut.Currency{Currency: "ATS", DisplayName: "австрійський шилінг", Symbol: ""}, "BOB": ut.Currency{Currency: "BOB", DisplayName: "болівійський болівіано", Symbol: "BOB"}, "GYD": ut.Currency{Currency: "GYD", DisplayName: "гаянський долар", Symbol: "GYD"}, "PEI": ut.Currency{Currency: "PEI", DisplayName: "перуанський інті", Symbol: ""}, "SSP": ut.Currency{Currency: "SSP", DisplayName: "південносуданський фунт", Symbol: "SSP"}, "ISK": ut.Currency{Currency: "ISK", DisplayName: "ісландська крона", Symbol: "ISK"}, "SKK": ut.Currency{Currency: "SKK", DisplayName: "словацька крона", Symbol: ""}, "SVC": ut.Currency{Currency: "SVC", DisplayName: "сальвадорський колон", Symbol: ""}, "ZWD": ut.Currency{Currency: "ZWD", DisplayName: "зімбабвійський долар", Symbol: ""}, "BUK": ut.Currency{Currency: "BUK", DisplayName: "бірманський кіат", Symbol: ""}, "GBP": ut.Currency{Currency: "GBP", DisplayName: "англійський фунт", Symbol: "GBP"}, "GWE": ut.Currency{Currency: "GWE", DisplayName: "ескудо португальської гвінеї", Symbol: ""}, "LTT": ut.Currency{Currency: "LTT", DisplayName: "литовський талон", Symbol: ""}, "BRE": ut.Currency{Currency: "BRE", DisplayName: "бразильське крузейро (1990–1993)", Symbol: ""}, "ECV": ut.Currency{Currency: "ECV", DisplayName: "еквадорський юнідад де валор константе", Symbol: ""}, "GEK": ut.Currency{Currency: "GEK", DisplayName: "грузинський купон", Symbol: ""}, "YER": ut.Currency{Currency: "YER", DisplayName: "єменський ріал", Symbol: "YER"}, "BAD": ut.Currency{Currency: "BAD", DisplayName: "динар (Боснія і Герцеговина)", Symbol: ""}, "IEP": ut.Currency{Currency: "IEP", DisplayName: "ірландський фунт", Symbol: ""}, "NZD": ut.Currency{Currency: "NZD", DisplayName: "новозеландський долар", Symbol: "NZD"}, "SBD": ut.Currency{Currency: "SBD", DisplayName: "долар Соломонових Островів", Symbol: "SBD"}, "SYP": ut.Currency{Currency: "SYP", DisplayName: "сирійський фунт", Symbol: "SYP"}, "XDR": ut.Currency{Currency: "XDR", DisplayName: "спеціальні права запозичення", Symbol: ""}, "BOP": ut.Currency{Currency: "BOP", DisplayName: "болівійське песо", Symbol: ""}, "GTQ": ut.Currency{Currency: "GTQ", DisplayName: "гватемальський кетсаль", Symbol: "GTQ"}, "SAR": ut.Currency{Currency: "SAR", DisplayName: "саудівський ріал", Symbol: "SAR"}, "PES": ut.Currency{Currency: "PES", DisplayName: "перуанський сол", Symbol: ""}, "SRD": ut.Currency{Currency: "SRD", DisplayName: "суринамський долар", Symbol: "SRD"}, "XXX": ut.Currency{Currency: "XXX", DisplayName: "невідома грошова одиниця", Symbol: ""}, "CUP": ut.Currency{Currency: "CUP", DisplayName: "кубинський песо", Symbol: "CUP"}, "GQE": ut.Currency{Currency: "GQE", DisplayName: "еквеле (Екваторіальна Ґвінея)", Symbol: ""}, "LVR": ut.Currency{Currency: "LVR", DisplayName: "латвійський рубль", Symbol: ""}, "LYD": ut.Currency{Currency: "LYD", DisplayName: "лівійський динар", Symbol: "LYD"}, "ADP": ut.Currency{Currency: "ADP", DisplayName: "андоррська песета", Symbol: ""}, "IQD": ut.Currency{Currency: "IQD", DisplayName: "іракський динар", Symbol: "IQD"}, "JPY": ut.Currency{Currency: "JPY", DisplayName: "японська єна", Symbol: "¥"}, "TZS": ut.Currency{Currency: "TZS", DisplayName: "танзанійський шилінг", Symbol: "TZS"}, "ZRN": ut.Currency{Currency: "ZRN", DisplayName: "заїрський новий заїр", Symbol: ""}, "IDR": ut.Currency{Currency: "IDR", DisplayName: "індонезійська рупія", Symbol: "IDR"}, "MWK": ut.Currency{Currency: "MWK", DisplayName: "малавійська квача", Symbol: "MWK"}, "NIO": ut.Currency{Currency: "NIO", DisplayName: "нікарагуанська кордоба", Symbol: "NIO"}, "MDL": ut.Currency{Currency: "MDL", DisplayName: "молдовський лей", Symbol: "MDL"}, "SIT": ut.Currency{Currency: "SIT", DisplayName: "словенський толар", Symbol: ""}, "TWD": ut.Currency{Currency: "TWD", DisplayName: "новий тайванський долар", Symbol: "TWD"}, "XOF": ut.Currency{Currency: "XOF", DisplayName: "західноафриканський франк", Symbol: "CFA"}, "BRN": ut.Currency{Currency: "BRN", DisplayName: "бразильське нове крузадо", Symbol: ""}, "CLP": ut.Currency{Currency: "CLP", DisplayName: "чилійський песо", Symbol: "CLP"}, "LUC": ut.Currency{Currency: "LUC", DisplayName: "люксембурґський франк (конвертований)", Symbol: ""}, "ALL": ut.Currency{Currency: "ALL", DisplayName: "албанський лек", Symbol: "ALL"}, "GEL": ut.Currency{Currency: "GEL", DisplayName: "грузинський ларі", Symbol: "GEL"}, "XFU": ut.Currency{Currency: "XFU", DisplayName: "французький франк UIC", Symbol: ""}, "MRO": ut.Currency{Currency: "MRO", DisplayName: "мавританська угія", Symbol: "MRO"}, "MZE": ut.Currency{Currency: "MZE", DisplayName: "мозамбіцький ескудо", Symbol: ""}, "BHD": ut.Currency{Currency: "BHD", DisplayName: "бахрейнський динар", Symbol: "BHD"}, "HKD": ut.Currency{Currency: "HKD", DisplayName: "гонконгський долар", Symbol: "HKD"}, "XAF": ut.Currency{Currency: "XAF", DisplayName: "центральноафриканський франк", Symbol: "FCFA"}, "BYB": ut.Currency{Currency: "BYB", DisplayName: "білоруський новий рубль (1994–1999)", Symbol: ""}, "DOP": ut.Currency{Currency: "DOP", DisplayName: "домініканський песо", Symbol: "DOP"}, "PAB": ut.Currency{Currency: "PAB", DisplayName: "панамська бальбоа", Symbol: "PAB"}, "LSL": ut.Currency{Currency: "LSL", DisplayName: "лесотський лоті", Symbol: ""}, "MUR": ut.Currency{Currency: "MUR", DisplayName: "маврикійська рупія", Symbol: "MUR"}, "ZWL": ut.Currency{Currency: "ZWL", DisplayName: "зімбабвійський долар (2009)", Symbol: ""}, "ZAL": ut.Currency{Currency: "ZAL", DisplayName: "південноафриканський фінансовий ранд", Symbol: ""}, "COP": ut.Currency{Currency: "COP", DisplayName: "колумбійський песо", Symbol: "COP"}, "CYP": ut.Currency{Currency: "CYP", DisplayName: "кіпрський фунт", Symbol: ""}, "FKP": ut.Currency{Currency: "FKP", DisplayName: "фунт Фолклендських островів", Symbol: "FKP"}, "AOR": ut.Currency{Currency: "AOR", DisplayName: "ангольська кванза реаджастадо (1995–1999)", Symbol: ""}, "TMT": ut.Currency{Currency: "TMT", DisplayName: "туркменський манат", Symbol: "TMT"}, "CHW": ut.Currency{Currency: "CHW", DisplayName: "франк WIR", Symbol: ""}, "LVL": ut.Currency{Currency: "LVL", DisplayName: "латвійський лат", Symbol: "LVL"}, "MTP": ut.Currency{Currency: "MTP", DisplayName: "мальтійський фунт", Symbol: ""}, "RHD": ut.Currency{Currency: "RHD", DisplayName: "родезійський долар", Symbol: ""}, "TMM": ut.Currency{Currency: "TMM", DisplayName: "туркменський манат (1993–2009)", Symbol: ""}, "AFA": ut.Currency{Currency: "AFA", DisplayName: "афгані (1927–2002)", Symbol: ""}, "BRR": ut.Currency{Currency: "BRR", DisplayName: "бразильське крузейро", Symbol: ""}, "CHE": ut.Currency{Currency: "CHE", DisplayName: "євро WIR", Symbol: ""}, "YUM": ut.Currency{Currency: "YUM", DisplayName: "югославський новий динар", Symbol: ""}, "VND": ut.Currency{Currency: "VND", DisplayName: "вʼєтнамський донг", Symbol: "VND"}, "MXN": ut.Currency{Currency: "MXN", DisplayName: "мексиканський песо", Symbol: "MXN"}, "TPE": ut.Currency{Currency: "TPE", DisplayName: "тіморський ескудо", Symbol: ""}, "UYP": ut.Currency{Currency: "UYP", DisplayName: "уругвайське песо (1975–1993)", Symbol: ""}, "ITL": ut.Currency{Currency: "ITL", DisplayName: "італійська ліра", Symbol: ""}, "OMR": ut.Currency{Currency: "OMR", DisplayName: "оманський ріал", Symbol: "OMR"}, "AOK": ut.Currency{Currency: "AOK", DisplayName: "ангольська кванза (1977–1990)", Symbol: ""}, "CHF": ut.Currency{Currency: "CHF", DisplayName: "швейцарський франк", Symbol: "CHF"}, "HRK": ut.Currency{Currency: "HRK", DisplayName: "хорватська куна", Symbol: "HRK"}, "ANG": ut.Currency{Currency: "ANG", DisplayName: "нідерландський антильський гульден", Symbol: "ANG"}, "BWP": ut.Currency{Currency: "BWP", DisplayName: "ботсванська пула", Symbol: "BWP"}, "ESP": ut.Currency{Currency: "ESP", DisplayName: "іспанська песета", Symbol: ""}, "PTE": ut.Currency{Currency: "PTE", DisplayName: "португальський ескудо", Symbol: ""}, "BRB": ut.Currency{Currency: "BRB", DisplayName: "бразильське нове крузейро (1967–1986)", Symbol: ""}, "KZT": ut.Currency{Currency: "KZT", DisplayName: "казахстанський тенге", Symbol: "KZT"}, "LTL": ut.Currency{Currency: "LTL", DisplayName: "литовський літ", Symbol: "LTL"}, "PEN": ut.Currency{Currency: "PEN", DisplayName: "перуанський новий сол", Symbol: "PEN"}, "AOA": ut.Currency{Currency: "AOA", DisplayName: "ангольська кванза", Symbol: "AOA"}, "BGN": ut.Currency{Currency: "BGN", DisplayName: "болгарський лев", Symbol: "BGN"}, "FRF": ut.Currency{Currency: "FRF", DisplayName: "французький франк", Symbol: ""}, "BND": ut.Currency{Currency: "BND", DisplayName: "брунейський долар", Symbol: "BND"}, "MGF": ut.Currency{Currency: "MGF", DisplayName: "мадагаскарський франк", Symbol: ""}, "XTS": ut.Currency{Currency: "XTS", DisplayName: "код тестування валюти", Symbol: ""}, "ILP": ut.Currency{Currency: "ILP", DisplayName: "ізраїльський фунт", Symbol: ""}, "SOS": ut.Currency{Currency: "SOS", DisplayName: "сомалійський шилінг", Symbol: "SOS"}, "XAG": ut.Currency{Currency: "XAG", DisplayName: "срібло", Symbol: ""}, "PLN": ut.Currency{Currency: "PLN", DisplayName: "польський злотий", Symbol: "PLN"}, "SDD": ut.Currency{Currency: "SDD", DisplayName: "суданський динар", Symbol: ""}, "SDG": ut.Currency{Currency: "SDG", DisplayName: "суданський фунт", Symbol: "SDG"}, "TND": ut.Currency{Currency: "TND", DisplayName: "туніський динар", Symbol: "TND"}, "TRL": ut.Currency{Currency: "TRL", DisplayName: "турецька ліра (1922–2005)", Symbol: ""}, "BEL": ut.Currency{Currency: "BEL", DisplayName: "бельгійський франк (фінансовий)", Symbol: ""}, "ESB": ut.Currency{Currency: "ESB", DisplayName: "іспанська песета (конвертовані рахунки)", Symbol: ""}, "KPW": ut.Currency{Currency: "KPW", DisplayName: "північнокорейський вон", Symbol: "KPW"}, "SZL": ut.Currency{Currency: "SZL", DisplayName: "свазілендський лілангені", Symbol: "SZL"}, "XCD": ut.Currency{Currency: "XCD", DisplayName: "східнокарибський долар", Symbol: "XCD"}, "BZD": ut.Currency{Currency: "BZD", DisplayName: "белізький долар", Symbol: "BZD"}, "CNY": ut.Currency{Currency: "CNY", DisplayName: "китайський юань", Symbol: "CNY"}, "NGN": ut.Currency{Currency: "NGN", DisplayName: "нігерійська найра", Symbol: "NGN"}, "AUD": ut.Currency{Currency: "AUD", DisplayName: "австралійський долар", Symbol: "AUD"}, "FIM": ut.Currency{Currency: "FIM", DisplayName: "фінляндська марка", Symbol: ""}, "MNT": ut.Currency{Currency: "MNT", DisplayName: "монгольський тугрик", Symbol: "MNT"}, "CRC": ut.Currency{Currency: "CRC", DisplayName: "костариканський колон", Symbol: "CRC"}, "MVR": ut.Currency{Currency: "MVR", DisplayName: "мальдівська руфія", Symbol: "MVR"}, "MXP": ut.Currency{Currency: "MXP", DisplayName: "мексиканське срібне песо (1861–1992)", Symbol: ""}, "ARA": ut.Currency{Currency: "ARA", DisplayName: "аргентинський австрал", Symbol: ""}, "AZM": ut.Currency{Currency: "AZM", DisplayName: "азербайджанський манат (1993–2006)", Symbol: ""}, "BAM": ut.Currency{Currency: "BAM", DisplayName: "конвертована марка Боснії і Герцеговини", Symbol: "BAM"}, "TJR": ut.Currency{Currency: "TJR", DisplayName: "таджицький рубль", Symbol: ""}, "USD": ut.Currency{Currency: "USD", DisplayName: "долар США", Symbol: "USD"}, "BIF": ut.Currency{Currency: "BIF", DisplayName: "бурундійський франк", Symbol: "BIF"}, "BRC": ut.Currency{Currency: "BRC", DisplayName: "бразильське крузадо", Symbol: ""}, "MXV": ut.Currency{Currency: "MXV", DisplayName: "мексиканський юнідад де інверсіон", Symbol: ""}, "HUF": ut.Currency{Currency: "HUF", DisplayName: "угорський форинт", Symbol: "HUF"}, "KHR": ut.Currency{Currency: "KHR", DisplayName: "камбоджійський рієль", Symbol: "KHR"}, "UZS": ut.Currency{Currency: "UZS", DisplayName: "узбецький сум", Symbol: "UZS"}, "ZMK": ut.Currency{Currency: "ZMK", DisplayName: "замбійська квача (1968–2012)", Symbol: ""}, "AMD": ut.Currency{Currency: "AMD", DisplayName: "вірменський драм", Symbol: "AMD"}, "BEF": ut.Currency{Currency: "BEF", DisplayName: "бельгійський франк", Symbol: ""}, "EEK": ut.Currency{Currency: "EEK", DisplayName: "естонська крона", Symbol: ""}, "CDF": ut.Currency{Currency: "CDF", DisplayName: "конголезький франк", Symbol: "CDF"}, "RON": ut.Currency{Currency: "RON", DisplayName: "румунський лей", Symbol: "RON"}, "XAU": ut.Currency{Currency: "XAU", DisplayName: "золото", Symbol: ""}, "KES": ut.Currency{Currency: "KES", DisplayName: "кенійський шилінг", Symbol: "KES"}, "PHP": ut.Currency{Currency: "PHP", DisplayName: "філіппінський песо", Symbol: "PHP"}, "UYI": ut.Currency{Currency: "UYI", DisplayName: "уругвайський песо в індексованих одиницях", Symbol: ""}, "XPD": ut.Currency{Currency: "XPD", DisplayName: "паладій", Symbol: ""}, "CUC": ut.Currency{Currency: "CUC", DisplayName: "кубинський конвертований песо", Symbol: "CUC"}, "LKR": ut.Currency{Currency: "LKR", DisplayName: "шрі-ланкійська рупія", Symbol: "LKR"}, "UGX": ut.Currency{Currency: "UGX", DisplayName: "угандійський шилінг", Symbol: "UGX"}, "UYU": ut.Currency{Currency: "UYU", DisplayName: "уругвайський песо", Symbol: "UYU"}, "ZAR": ut.Currency{Currency: "ZAR", DisplayName: "південноафриканський ранд", Symbol: "ZAR"}, "AED": ut.Currency{Currency: "AED", DisplayName: "дирхам ОАЕ", Symbol: "AED"}, "CVE": ut.Currency{Currency: "CVE", DisplayName: "ескудо Кабо-Верде", Symbol: "CVE"}, "DDM": ut.Currency{Currency: "DDM", DisplayName: "марка НДР", Symbol: ""}, "BMD": ut.Currency{Currency: "BMD", DisplayName: "бермудський долар", Symbol: "BMD"}, "PYG": ut.Currency{Currency: "PYG", DisplayName: "парагвайський гуарані", Symbol: "PYG"}, "UAH": ut.Currency{Currency: "UAH", DisplayName: "українська гривня", Symbol: "₴"}, "UAK": ut.Currency{Currency: "UAK", DisplayName: "український карбованець", Symbol: "крб."}, "YDD": ut.Currency{Currency: "YDD", DisplayName: "єменський динар", Symbol: ""}, "DKK": ut.Currency{Currency: "DKK", DisplayName: "данська крона", Symbol: "DKK"}, "MAF": ut.Currency{Currency: "MAF", DisplayName: "марокканський франк", Symbol: ""}, "RUB": ut.Currency{Currency: "RUB", DisplayName: "російський рубль", Symbol: "RUB"}, "JOD": ut.Currency{Currency: "JOD", DisplayName: "йорданський динар", Symbol: "JOD"}, "LUF": ut.Currency{Currency: "LUF", DisplayName: "люксембурзький франк", Symbol: ""}, "NPR": ut.Currency{Currency: "NPR", DisplayName: "непальська рупія", Symbol: "NPR"}, "HRD": ut.Currency{Currency: "HRD", DisplayName: "хорватський динар", Symbol: ""}, "LRD": ut.Currency{Currency: "LRD", DisplayName: "ліберійський долар", Symbol: "LRD"}, "BOV": ut.Currency{Currency: "BOV", DisplayName: "болівійський мвдол", Symbol: ""}, "BSD": ut.Currency{Currency: "BSD", DisplayName: "багамський долар", Symbol: "BSD"}, "GWP": ut.Currency{Currency: "GWP", DisplayName: "песо Гвінеї-Бісау", Symbol: ""}, "WST": ut.Currency{Currency: "WST", DisplayName: "самоанська тала", Symbol: "WST"}, "HNL": ut.Currency{Currency: "HNL", DisplayName: "гондураська лемпіра", Symbol: "HNL"}, "MZN": ut.Currency{Currency: "MZN", DisplayName: "мозамбіцький метикал", Symbol: "MZN"}, "UGS": ut.Currency{Currency: "UGS", DisplayName: "угандійський шилінг (1966–1987)", Symbol: ""}, "YUN": ut.Currency{Currency: "YUN", DisplayName: "югославський конвертований динар", Symbol: ""}, "KYD": ut.Currency{Currency: "KYD", DisplayName: "долар Кайманових островів", Symbol: "KYD"}, "MAD": ut.Currency{Currency: "MAD", DisplayName: "марокканський дирхам", Symbol: "MAD"}, "SLL": ut.Currency{Currency: "SLL", DisplayName: "леоне Сьєрра-Леоне", Symbol: "SLL"}, "NLG": ut.Currency{Currency: "NLG", DisplayName: "нідерландський гульден", Symbol: ""}, "SRG": ut.Currency{Currency: "SRG", DisplayName: "суринамський гульден", Symbol: ""}, "ZRZ": ut.Currency{Currency: "ZRZ", DisplayName: "заїрський заїр", Symbol: ""}, "EGP": ut.Currency{Currency: "EGP", DisplayName: "єгипетський фунт", Symbol: "EGP"}, "GHS": ut.Currency{Currency: "GHS", DisplayName: "ганський седі", Symbol: "GHS"}, "RUR": ut.Currency{Currency: "RUR", DisplayName: "російський рубль (1991–1998)", Symbol: ""}, "BGL": ut.Currency{Currency: "BGL", DisplayName: "болгарський твердий лев", Symbol: ""}, "CSK": ut.Currency{Currency: "CSK", DisplayName: "чехословацька тверда крона", Symbol: ""}, "JMD": ut.Currency{Currency: "JMD", DisplayName: "ямайський долар", Symbol: "JMD"}, "LUL": ut.Currency{Currency: "LUL", DisplayName: "люксембурґський франк (фінансовий)", Symbol: ""}, "VEB": ut.Currency{Currency: "VEB", DisplayName: "венесуельський болівар (1871–2008)", Symbol: ""}, "XPT": ut.Currency{Currency: "XPT", DisplayName: "платина", Symbol: ""}, "ZWR": ut.Currency{Currency: "ZWR", DisplayName: "зімбабвійський долар (2008)", Symbol: ""}, "EUR": ut.Currency{Currency: "EUR", DisplayName: "євро", Symbol: "EUR"}, "QAR": ut.Currency{Currency: "QAR", DisplayName: "катарський ріал", Symbol: "QAR"}, "SUR": ut.Currency{Currency: "SUR", DisplayName: "радянський рубль", Symbol: ""}, "MLF": ut.Currency{Currency: "MLF", DisplayName: "малійський франк", Symbol: ""}, "MYR": ut.Currency{Currency: "MYR", DisplayName: "малайзійський рингіт", Symbol: "MYR"}, "XEU": ut.Currency{Currency: "XEU", DisplayName: "європейська валютна одиниця", Symbol: ""}, "XPF": ut.Currency{Currency: "XPF", DisplayName: "французький тихоокеанський франк", Symbol: "CFPF"}, "CAD": ut.Currency{Currency: "CAD", DisplayName: "канадський долар", Symbol: "CAD"}, "CZK": ut.Currency{Currency: "CZK", DisplayName: "чеська крона", Symbol: "CZK"}, "DEM": ut.Currency{Currency: "DEM", DisplayName: "німецька марка", Symbol: ""}, "MGA": ut.Currency{Currency: "MGA", DisplayName: "малагасійський аріарі", Symbol: "MGA"}, "RSD": ut.Currency{Currency: "RSD", DisplayName: "сербський динар", Symbol: "RSD"}, "USS": ut.Currency{Currency: "USS", DisplayName: "долар США (цього дня)", Symbol: ""}, "FJD": ut.Currency{Currency: "FJD", DisplayName: "фіджійський долар", Symbol: "FJD"}, "GRD": ut.Currency{Currency: "GRD", DisplayName: "грецька драхма", Symbol: ""}, "LAK": ut.Currency{Currency: "LAK", DisplayName: "лаоський кіп", Symbol: "LAK"}, "AWG": ut.Currency{Currency: "AWG", DisplayName: "арубський флорин", Symbol: "AWG"}, "KWD": ut.Currency{Currency: "KWD", DisplayName: "кувейтський динар", Symbol: "KWD"}, "XBB": ut.Currency{Currency: "XBB", DisplayName: "одиниця європейського валютного фонду", Symbol: ""}, "USN": ut.Currency{Currency: "USN", DisplayName: "долар США (наступного дня)", Symbol: ""}, "VEF": ut.Currency{Currency: "VEF", DisplayName: "венесуельський болівар", Symbol: "VEF"}, "ECS": ut.Currency{Currency: "ECS", DisplayName: "еквадорський сукре", Symbol: ""}, "ESA": ut.Currency{Currency: "ESA", DisplayName: "іспанська песета (\"А\" рахунок)", Symbol: ""}, "ROL": ut.Currency{Currency: "ROL", DisplayName: "старий румунський лей", Symbol: ""}}