package bg

import "github.com/go-playground/universal-translator"

var currencies = ut.CurrencyFormatValue{"LSL": ut.Currency{Currency: "LSL", DisplayName: "Лесотско лоти", Symbol: ""}, "MWK": ut.Currency{Currency: "MWK", DisplayName: "Малавийска квача", Symbol: "MWK"}, "NIC": ut.Currency{Currency: "NIC", DisplayName: "Никарагуанска кордоба (1988–1991)", Symbol: "NIC"}, "TRY": ut.Currency{Currency: "TRY", DisplayName: "Турска лира", Symbol: "TRY"}, "VND": ut.Currency{Currency: "VND", DisplayName: "Виетнамски донг", Symbol: "VND"}, "YER": ut.Currency{Currency: "YER", DisplayName: "Йеменски риал", Symbol: "YER"}, "DEM": ut.Currency{Currency: "DEM", DisplayName: "Германска марка", Symbol: ""}, "BOV": ut.Currency{Currency: "BOV", DisplayName: "Боливийски мвдол", Symbol: ""}, "MUR": ut.Currency{Currency: "MUR", DisplayName: "Маврицийска рупия", Symbol: "MUR"}, "VEB": ut.Currency{Currency: "VEB", DisplayName: "Венецуелски боливар (1871–2008)", Symbol: ""}, "XDR": ut.Currency{Currency: "XDR", DisplayName: "Специални права на тираж", Symbol: ""}, "YUM": ut.Currency{Currency: "YUM", DisplayName: "Югославски динар", Symbol: ""}, "AOA": ut.Currency{Currency: "AOA", DisplayName: "Анголска кванза", Symbol: "AOA"}, "HTG": ut.Currency{Currency: "HTG", DisplayName: "Хаитски гурд", Symbol: "HTG"}, "INR": ut.Currency{Currency: "INR", DisplayName: "Индийска рупия", Symbol: "INR"}, "NZD": ut.Currency{Currency: "NZD", DisplayName: "Новозеландски долар", Symbol: "NZD"}, "SOS": ut.Currency{Currency: "SOS", DisplayName: "Сомалийски шилинг", Symbol: "SOS"}, "USS": ut.Currency{Currency: "USS", DisplayName: "", Symbol: "USS"}, "VUV": ut.Currency{Currency: "VUV", DisplayName: "Вануатско вату", Symbol: "VUV"}, "XBB": ut.Currency{Currency: "XBB", DisplayName: "Европейска валутна единица", Symbol: ""}, "BRN": ut.Currency{Currency: "BRN", DisplayName: "Бразилско ново крозадо", Symbol: ""}, "BTN": ut.Currency{Currency: "BTN", DisplayName: "Бутански нгултрум", Symbol: "BTN"}, "RSD": ut.Currency{Currency: "RSD", DisplayName: "Сръбски динар", Symbol: "RSD"}, "XAG": ut.Currency{Currency: "XAG", DisplayName: "Сребро", Symbol: ""}, "XPF": ut.Currency{Currency: "XPF", DisplayName: "Френскополинезийски франк", Symbol: "CFPF"}, "BHD": ut.Currency{Currency: "BHD", DisplayName: "Бахрейнски динар", Symbol: "BHD"}, "MRO": ut.Currency{Currency: "MRO", DisplayName: "Мавританска угия", Symbol: "MRO"}, "XAU": ut.Currency{Currency: "XAU", DisplayName: "Злато", Symbol: ""}, "AFA": ut.Currency{Currency: "AFA", DisplayName: "Афганистански афган (1927–2002)", Symbol: ""}, "TMM": ut.Currency{Currency: "TMM", DisplayName: "Туркменистански манат", Symbol: ""}, "UAH": ut.Currency{Currency: "UAH", DisplayName: "Украинска хривня", Symbol: "UAH"}, "ZRZ": ut.Currency{Currency: "ZRZ", DisplayName: "Заирско зайре", Symbol: ""}, "SBD": ut.Currency{Currency: "SBD", DisplayName: "Долар на Соломоновите острови", Symbol: "SBD"}, "DDM": ut.Currency{Currency: "DDM", DisplayName: "Източногерманска марка", Symbol: ""}, "ECS": ut.Currency{Currency: "ECS", DisplayName: "Еквадорско сукре", Symbol: ""}, "MXV": ut.Currency{Currency: "MXV", DisplayName: "Мексиканска конвертируема единица (UDI)", Symbol: "MXV"}, "QAR": ut.Currency{Currency: "QAR", DisplayName: "Катарски риал", Symbol: "QAR"}, "TWD": ut.Currency{Currency: "TWD", DisplayName: "Тайвански долар", Symbol: "TWD"}, "CRC": ut.Currency{Currency: "CRC", DisplayName: "Костарикански колон", Symbol: "CRC"}, "UGS": ut.Currency{Currency: "UGS", DisplayName: "Угандийски шилинг (1966–1987)", Symbol: ""}, "OMR": ut.Currency{Currency: "OMR", DisplayName: "Омански риал", Symbol: "OMR"}, "PGK": ut.Currency{Currency: "PGK", DisplayName: "Папуа-новогвинейска кина", Symbol: "PGK"}, "XXX": ut.Currency{Currency: "XXX", DisplayName: "Непозната валута", Symbol: ""}, "GWE": ut.Currency{Currency: "GWE", DisplayName: "Ескудо от Португалска Гвинея", Symbol: ""}, "KHR": ut.Currency{Currency: "KHR", DisplayName: "Камбоджански риел", Symbol: "KHR"}, "TRL": ut.Currency{Currency: "TRL", DisplayName: "Турска лира (1922–2005)", Symbol: ""}, "TZS": ut.Currency{Currency: "TZS", DisplayName: "Танзанийски шилинг", Symbol: "TZS"}, "UGX": ut.Currency{Currency: "UGX", DisplayName: "Угандски шилинг", Symbol: "UGX"}, "YUD": ut.Currency{Currency: "YUD", DisplayName: "Югославски твърд динар", Symbol: ""}, "CZK": ut.Currency{Currency: "CZK", DisplayName: "Чешка крона", Symbol: "CZK"}, "COP": ut.Currency{Currency: "COP", DisplayName: "Колумбийско песо", Symbol: "COP"}, "FRF": ut.Currency{Currency: "FRF", DisplayName: "Френски франк", Symbol: ""}, "LBP": ut.Currency{Currency: "LBP", DisplayName: "Ливанска лира", Symbol: "LBP"}, "LVL": ut.Currency{Currency: "LVL", DisplayName: "Латвийски лат", Symbol: "LVL"}, "NGN": ut.Currency{Currency: "NGN", DisplayName: "Нигерийска найра", Symbol: "NGN"}, "UAK": ut.Currency{Currency: "UAK", DisplayName: "Украински карбованец", Symbol: ""}, "BRE": ut.Currency{Currency: "BRE", DisplayName: "Бразилско крузейро (1990–1993)", Symbol: ""}, "RUR": ut.Currency{Currency: "RUR", DisplayName: "Руска рубла (1991–1998)", Symbol: ""}, "XBC": ut.Currency{Currency: "XBC", DisplayName: "Европейска единица по сметка (XBC)", Symbol: ""}, "NIO": ut.Currency{Currency: "NIO", DisplayName: "Никарагуанска кордоба", Symbol: "NIO"}, "LVR": ut.Currency{Currency: "LVR", DisplayName: "Латвийска рубла", Symbol: ""}, "PKR": ut.Currency{Currency: "PKR", DisplayName: "Пакистанска рупия", Symbol: "PKR"}, "SRG": ut.Currency{Currency: "SRG", DisplayName: "Суринамски гилдер", Symbol: ""}, "IQD": ut.Currency{Currency: "IQD", DisplayName: "Иракски динар", Symbol: "IQD"}, "ECV": ut.Currency{Currency: "ECV", DisplayName: "Еквадорска банкова единица", Symbol: ""}, "MLF": ut.Currency{Currency: "MLF", DisplayName: "Малийски франк", Symbol: ""}, "MZN": ut.Currency{Currency: "MZN", DisplayName: "Мозамбикски метикал", Symbol: "MZN"}, "XEU": ut.Currency{Currency: "XEU", DisplayName: "Еку на ЕИО", Symbol: ""}, "CYP": ut.Currency{Currency: "CYP", DisplayName: "Кипърска лира", Symbol: ""}, "EGP": ut.Currency{Currency: "EGP", DisplayName: "Египетска лира", Symbol: "EGP"}, "FJD": ut.Currency{Currency: "FJD", DisplayName: "Фиджийски долар", Symbol: "FJD"}, "GMD": ut.Currency{Currency: "GMD", DisplayName: "Гамбийско даласи", Symbol: "GMD"}, "WST": ut.Currency{Currency: "WST", DisplayName: "Самоанска тала", Symbol: "WST"}, "BDT": ut.Currency{Currency: "BDT", DisplayName: "Бангладешка така", Symbol: "BDT"}, "ERN": ut.Currency{Currency: "ERN", DisplayName: "Еритрейска накфа", Symbol: "ERN"}, "LTT": ut.Currency{Currency: "LTT", DisplayName: "Литовски талон", Symbol: ""}, "MZM": ut.Currency{Currency: "MZM", DisplayName: "Мозамбикски метикал (1980–2006)", Symbol: ""}, "NPR": ut.Currency{Currency: "NPR", DisplayName: "Непалска рупия", Symbol: "NPR"}, "BRC": ut.Currency{Currency: "BRC", DisplayName: "Бразилско крозадо", Symbol: ""}, "IEP": ut.Currency{Currency: "IEP", DisplayName: "Ирландска лира", Symbol: ""}, "KWD": ut.Currency{Currency: "KWD", DisplayName: "Кувейтски динар", Symbol: "KWD"}, "MOP": ut.Currency{Currency: "MOP", DisplayName: "Патака на Макао", Symbol: "MOP"}, "MVR": ut.Currency{Currency: "MVR", DisplayName: "Малдивска руфия", Symbol: "MVR"}, "ZWD": ut.Currency{Currency: "ZWD", DisplayName: "Зимбабвийски долар", Symbol: ""}, "GNS": ut.Currency{Currency: "GNS", DisplayName: "Гвинейска сили", Symbol: ""}, "MZE": ut.Currency{Currency: "MZE", DisplayName: "Мозамбикско ескудо", Symbol: ""}, "UYP": ut.Currency{Currency: "UYP", DisplayName: "Уругвайско песо (1975–1993)", Symbol: ""}, "GHS": ut.Currency{Currency: "GHS", DisplayName: "Ганайско седи", Symbol: "GHS"}, "DOP": ut.Currency{Currency: "DOP", DisplayName: "Доминиканско песо", Symbol: "DOP"}, "GIP": ut.Currency{Currency: "GIP", DisplayName: "Гибралтарска лира", Symbol: "GIP"}, "XPT": ut.Currency{Currency: "XPT", DisplayName: "Платина", Symbol: ""}, "BEC": ut.Currency{Currency: "BEC", DisplayName: "Белгийски франк (конвертируем)", Symbol: ""}, "KES": ut.Currency{Currency: "KES", DisplayName: "Кенийски шилинг", Symbol: "KES"}, "SHP": ut.Currency{Currency: "SHP", DisplayName: "Лира на Света Елена", Symbol: "SHP"}, "ITL": ut.Currency{Currency: "ITL", DisplayName: "Италианска лира", Symbol: ""}, "GEK": ut.Currency{Currency: "GEK", DisplayName: "Грузински купон", Symbol: ""}, "RON": ut.Currency{Currency: "RON", DisplayName: "Румънска лея", Symbol: "RON"}, "ZRN": ut.Currency{Currency: "ZRN", DisplayName: "Заирско ново зайре", Symbol: ""}, "CHE": ut.Currency{Currency: "CHE", DisplayName: "WIR евро", Symbol: ""}, "BWP": ut.Currency{Currency: "BWP", DisplayName: "Ботсванска пула", Symbol: "BWP"}, "KPW": ut.Currency{Currency: "KPW", DisplayName: "Севернокорейски вон", Symbol: "KPW"}, "SUR": ut.Currency{Currency: "SUR", DisplayName: "Съветска рубла", Symbol: ""}, "XAF": ut.Currency{Currency: "XAF", DisplayName: "Централноафрикански франк", Symbol: "FCFA"}, "BEL": ut.Currency{Currency: "BEL", DisplayName: "Белгийски франк (финансов)", Symbol: ""}, "GWP": ut.Currency{Currency: "GWP", DisplayName: "Гвинея-Бисау песо", Symbol: ""}, "HKD": ut.Currency{Currency: "HKD", DisplayName: "Хонконгски долар", Symbol: "HKD"}, "SRD": ut.Currency{Currency: "SRD", DisplayName: "Суринамски долар", Symbol: "SRD"}, "XCD": ut.Currency{Currency: "XCD", DisplayName: "Източнокарибски долар", Symbol: "XCD"}, "ATS": ut.Currency{Currency: "ATS", DisplayName: "Австрийски шилинг", Symbol: ""}, "DJF": ut.Currency{Currency: "DJF", DisplayName: "Джибутски франк", Symbol: "DJF"}, "ETB": ut.Currency{Currency: "ETB", DisplayName: "Етиопски бир", Symbol: "ETB"}, "MXN": ut.Currency{Currency: "MXN", DisplayName: "Мексиканско песо", Symbol: "MXN"}, "AWG": ut.Currency{Currency: "AWG", DisplayName: "Арубски флорин", Symbol: "AWG"}, "FIM": ut.Currency{Currency: "FIM", DisplayName: "Финландска марка", Symbol: ""}, "PTE": ut.Currency{Currency: "PTE", DisplayName: "Португалско ескудо", Symbol: ""}, "BRR": ut.Currency{Currency: "BRR", DisplayName: "Бразилско крузейро", Symbol: ""}, "ESP": ut.Currency{Currency: "ESP", DisplayName: "Испанска песета", Symbol: ""}, "RWF": ut.Currency{Currency: "RWF", DisplayName: "Руандски франк", Symbol: "RWF"}, "CUC": ut.Currency{Currency: "CUC", DisplayName: "Кубинско конвертируемо песо", Symbol: "CUC"}, "GRD": ut.Currency{Currency: "GRD", DisplayName: "Гръцка драхма", Symbol: ""}, "UZS": ut.Currency{Currency: "UZS", DisplayName: "Узбекски сум", Symbol: "UZS"}, "ALL": ut.Currency{Currency: "ALL", DisplayName: "Албански лек", Symbol: "ALL"}, "KYD": ut.Currency{Currency: "KYD", DisplayName: "Кайманов долар", Symbol: "KYD"}, "MDL": ut.Currency{Currency: "MDL", DisplayName: "Молдовско леу", Symbol: "MDL"}, "NAD": ut.Currency{Currency: "NAD", DisplayName: "Намибийски долар", Symbol: "NAD"}, "SCR": ut.Currency{Currency: "SCR", DisplayName: "Сейшелска рупия", Symbol: "SCR"}, "TOP": ut.Currency{Currency: "TOP", DisplayName: "Тонганска паанга", Symbol: "TOP"}, "BYR": ut.Currency{Currency: "BYR", DisplayName: "Беларуска рубла", Symbol: "BYR"}, "ILS": ut.Currency{Currency: "ILS", DisplayName: "Израелски нов шекел", Symbol: "ILS"}, "ROL": ut.Currency{Currency: "ROL", DisplayName: "Стара румънска лея", Symbol: ""}, "SGD": ut.Currency{Currency: "SGD", DisplayName: "Сингапурски долар", Symbol: "SGD"}, "STD": ut.Currency{Currency: "STD", DisplayName: "Добра на Сао Томе и Принсипи", Symbol: "STD"}, "ZAL": ut.Currency{Currency: "ZAL", DisplayName: "Южноафрикански ранд (финансов)", Symbol: ""}, "AOK": ut.Currency{Currency: "AOK", DisplayName: "Анголска кванца (1977–1990)", Symbol: ""}, "SVC": ut.Currency{Currency: "SVC", DisplayName: "Салвадорски колон", Symbol: "SVC"}, "BBD": ut.Currency{Currency: "BBD", DisplayName: "Барбадоски долар", Symbol: "BBD"}, "RHD": ut.Currency{Currency: "RHD", DisplayName: "Родезийски долар", Symbol: ""}, "RUB": ut.Currency{Currency: "RUB", DisplayName: "Руска рубла", Symbol: "RUB"}, "BUK": ut.Currency{Currency: "BUK", DisplayName: "Бирмански киат", Symbol: ""}, "CNY": ut.Currency{Currency: "CNY", DisplayName: "Китайски юан", Symbol: "CNY"}, "HNL": ut.Currency{Currency: "HNL", DisplayName: "Хондураска лемпира", Symbol: "HNL"}, "JMD": ut.Currency{Currency: "JMD", DisplayName: "Ямайски долар", Symbol: "JMD"}, "PES": ut.Currency{Currency: "PES", DisplayName: "Перуански сол", Symbol: ""}, "BYB": ut.Currency{Currency: "BYB", DisplayName: "Беларуска нова рубла (1994–1999)", Symbol: ""}, "CLF": ut.Currency{Currency: "CLF", DisplayName: "Условна разчетна единица на Чили", Symbol: ""}, "COU": ut.Currency{Currency: "COU", DisplayName: "Колумбийска единица на реалната стойност", Symbol: ""}, "SZL": ut.Currency{Currency: "SZL", DisplayName: "Свазилендски лилангени", Symbol: "SZL"}, "TMT": ut.Currency{Currency: "TMT", DisplayName: "Туркменски манат", Symbol: "TMT"}, "TPE": ut.Currency{Currency: "TPE", DisplayName: "Тиморско ескудо", Symbol: ""}, "CDF": ut.Currency{Currency: "CDF", DisplayName: "Конгоански франк", Symbol: "CDF"}, "BZD": ut.Currency{Currency: "BZD", DisplayName: "Белизийски долар", Symbol: "BZD"}, "DKK": ut.Currency{Currency: "DKK", DisplayName: "Датска крона", Symbol: "DKK"}, "BAM": ut.Currency{Currency: "BAM", DisplayName: "Босненска конвертируема марка", Symbol: "BAM"}, "BMD": ut.Currency{Currency: "BMD", DisplayName: "Бермудски долар", Symbol: "BMD"}, "IDR": ut.Currency{Currency: "IDR", DisplayName: "Индонезийска рупия", Symbol: "IDR"}, "KRW": ut.Currency{Currency: "KRW", DisplayName: "Южнокорейски вон", Symbol: "KRW"}, "SKK": ut.Currency{Currency: "SKK", DisplayName: "Словашка крона", Symbol: ""}, "BGL": ut.Currency{Currency: "BGL", DisplayName: "Български конвертируем лев (1962–1999)", Symbol: ""}, "BEF": ut.Currency{Currency: "BEF", DisplayName: "Белгийски франк", Symbol: ""}, "XPD": ut.Currency{Currency: "XPD", DisplayName: "Паладий", Symbol: ""}, "XTS": ut.Currency{Currency: "XTS", DisplayName: "Код резервиран за целите на тестване", Symbol: ""}, "ARP": ut.Currency{Currency: "ARP", DisplayName: "Аржентинско песо (1983–1985)", Symbol: ""}, "LUF": ut.Currency{Currency: "LUF", DisplayName: "Люксембургски франк", Symbol: ""}, "SIT": ut.Currency{Currency: "SIT", DisplayName: "Словенски толар", Symbol: ""}, "VEF": ut.Currency{Currency: "VEF", DisplayName: "Венецуелски боливар", Symbol: "VEF"}, "AUD": ut.Currency{Currency: "AUD", DisplayName: "Австралийски долар", Symbol: "AUD"}, "LAK": ut.Currency{Currency: "LAK", DisplayName: "Лаоски кип", Symbol: "LAK"}, "MGF": ut.Currency{Currency: "MGF", DisplayName: "Малгашки франк - Мадагаскар", Symbol: ""}, "SAR": ut.Currency{Currency: "SAR", DisplayName: "Саудитскоарабски риал", Symbol: "SAR"}, "BOP": ut.Currency{Currency: "BOP", DisplayName: "Боливийско песо", Symbol: ""}, "ISK": ut.Currency{Currency: "ISK", DisplayName: "Исландска крона", Symbol: "ISK"}, "KZT": ut.Currency{Currency: "KZT", DisplayName: "Казахстанско тенге", Symbol: "KZT"}, "PEI": ut.Currency{Currency: "PEI", DisplayName: "Перуанско инти", Symbol: ""}, "PHP": ut.Currency{Currency: "PHP", DisplayName: "Филипинско песо", Symbol: "PHP"}, "THB": ut.Currency{Currency: "THB", DisplayName: "Тайландски бат", Symbol: "THB"}, "TJR": ut.Currency{Currency: "TJR", DisplayName: "Таджикистанска рубла", Symbol: ""}, "AZM": ut.Currency{Currency: "AZM", DisplayName: "Азербайджански манат (1993–2006)", Symbol: ""}, "CLP": ut.Currency{Currency: "CLP", DisplayName: "Чилийско песо", Symbol: "CLP"}, "FKP": ut.Currency{Currency: "FKP", DisplayName: "Фолклендска лира", Symbol: "FKP"}, "LYD": ut.Currency{Currency: "LYD", DisplayName: "Либийски динар", Symbol: "LYD"}, "NOK": ut.Currency{Currency: "NOK", DisplayName: "Норвежка крона", Symbol: "NOK"}, "BGN": ut.Currency{Currency: "BGN", DisplayName: "Български лев", Symbol: "лв."}, "MAD": ut.Currency{Currency: "MAD", DisplayName: "Марокански дирхам", Symbol: "MAD"}, "EUR": ut.Currency{Currency: "EUR", DisplayName: "Евро", Symbol: "€"}, "CVE": ut.Currency{Currency: "CVE", DisplayName: "Ескудо на Кабо Верде", Symbol: "CVE"}, "XFO": ut.Currency{Currency: "XFO", DisplayName: "Френски златен франк", Symbol: ""}, "AON": ut.Currency{Currency: "AON", DisplayName: "Анголска нова кванца (1990–2000)", Symbol: ""}, "UYI": ut.Currency{Currency: "UYI", DisplayName: "Уругвайско песо (индекс на инфлацията)", Symbol: ""}, "LRD": ut.Currency{Currency: "LRD", DisplayName: "Либерийски долар", Symbol: "LRD"}, "MAF": ut.Currency{Currency: "MAF", DisplayName: "Марокански франк", Symbol: ""}, "NLG": ut.Currency{Currency: "NLG", DisplayName: "Холандски гулден", Symbol: ""}, "ZAR": ut.Currency{Currency: "ZAR", DisplayName: "Южноафрикански ранд", Symbol: "ZAR"}, "DZD": ut.Currency{Currency: "DZD", DisplayName: "Алжирски динар", Symbol: "DZD"}, "CSD": ut.Currency{Currency: "CSD", DisplayName: "Стар сръбски динар", Symbol: ""}, "GQE": ut.Currency{Currency: "GQE", DisplayName: "Екваториално гвинейско еквеле", Symbol: ""}, "YUN": ut.Currency{Currency: "YUN", DisplayName: "Югославски конвертируем динар", Symbol: ""}, "CHF": ut.Currency{Currency: "CHF", DisplayName: "Швейцарски франк", Symbol: "CHF"}, "AZN": ut.Currency{Currency: "AZN", DisplayName: "Азербайджански манат", Symbol: "AZN"}, "BRL": ut.Currency{Currency: "BRL", DisplayName: "Бразилски реал", Symbol: "BRL"}, "BSD": ut.Currency{Currency: "BSD", DisplayName: "Бахамски долар", Symbol: "BSD"}, "ILP": ut.Currency{Currency: "ILP", DisplayName: "Израелска лира", Symbol: ""}, "PLN": ut.Currency{Currency: "PLN", DisplayName: "Полска злота", Symbol: "PLN"}, "USD": ut.Currency{Currency: "USD", DisplayName: "Щатски долар", Symbol: "щ.д."}, "AOR": ut.Currency{Currency: "AOR", DisplayName: "Анголска нова кванца (1995–1999)", Symbol: ""}, "XOF": ut.Currency{Currency: "XOF", DisplayName: "Западноафрикански франк", Symbol: "CFA"}, "SEK": ut.Currency{Currency: "SEK", DisplayName: "Шведска крона", Symbol: "SEK"}, "MMK": ut.Currency{Currency: "MMK", DisplayName: "Мианмарски кият", Symbol: "MMK"}, "SYP": ut.Currency{Currency: "SYP", DisplayName: "Сирийска лира", Symbol: "SYP"}, "IRR": ut.Currency{Currency: "IRR", DisplayName: "Ирански риал", Symbol: "IRR"}, "ANG": ut.Currency{Currency: "ANG", DisplayName: "Антилски гулден", Symbol: "ANG"}, "LKR": ut.Currency{Currency: "LKR", DisplayName: "Шриланкска рупия", Symbol: "LKR"}, "SDG": ut.Currency{Currency: "SDG", DisplayName: "Суданска лира", Symbol: "SDG"}, "ADP": ut.Currency{Currency: "ADP", DisplayName: "Андорска песета", Symbol: ""}, "GBP": ut.Currency{Currency: "GBP", DisplayName: "Британска лира", Symbol: "GBP"}, "GHC": ut.Currency{Currency: "GHC", DisplayName: "Ганайско седи (1979–2007)", Symbol: ""}, "GYD": ut.Currency{Currency: "GYD", DisplayName: "Гаянски долар", Symbol: "GYD"}, "PLZ": ut.Currency{Currency: "PLZ", DisplayName: "Полска злота (1950–1995)", Symbol: ""}, "CAD": ut.Currency{Currency: "CAD", DisplayName: "Канадски долар", Symbol: "CAD"}, "HRK": ut.Currency{Currency: "HRK", DisplayName: "Хърватска куна", Symbol: "HRK"}, "LTL": ut.Currency{Currency: "LTL", DisplayName: "Литовски литас", Symbol: "LTL"}, "SLL": ut.Currency{Currency: "SLL", DisplayName: "Сиералеонско леоне", Symbol: "SLL"}, "CSK": ut.Currency{Currency: "CSK", DisplayName: "Чехословашка конвертируема крона", Symbol: ""}, "TTD": ut.Currency{Currency: "TTD", DisplayName: "Долар на Тринидад и Тобаго", Symbol: "TTD"}, "XBD": ut.Currency{Currency: "XBD", DisplayName: "Европейска единица по сметка (XBD)", Symbol: ""}, "MKD": ut.Currency{Currency: "MKD", DisplayName: "Македонски денар", Symbol: "MKD"}, "EEK": ut.Currency{Currency: "EEK", DisplayName: "Естонска крона", Symbol: ""}, "KMF": ut.Currency{Currency: "KMF", DisplayName: "Коморски франк", Symbol: "KMF"}, "BIF": ut.Currency{Currency: "BIF", DisplayName: "Бурундийски франк", Symbol: "BIF"}, "TND": ut.Currency{Currency: "TND", DisplayName: "Тунизийски динар", Symbol: "TND"}, "XBA": ut.Currency{Currency: "XBA", DisplayName: "Европейска съставна единица", Symbol: ""}, "MGA": ut.Currency{Currency: "MGA", DisplayName: "Малгашко ариари", Symbol: "MGA"}, "GNF": ut.Currency{Currency: "GNF", DisplayName: "Гвинейски франк", Symbol: "GNF"}, "ARS": ut.Currency{Currency: "ARS", DisplayName: "Аржентинско песо", Symbol: "ARS"}, "SSP": ut.Currency{Currency: "SSP", DisplayName: "Южносуданска лира", Symbol: "SSP"}, "USN": ut.Currency{Currency: "USN", DisplayName: "", Symbol: "USN"}, "ZMK": ut.Currency{Currency: "ZMK", DisplayName: "Замбийска квача (1968–2012)", Symbol: ""}, "ZWL": ut.Currency{Currency: "ZWL", DisplayName: "Зимбабвийски долар (2009)", Symbol: ""}, "CHW": ut.Currency{Currency: "CHW", DisplayName: "WIR франк", Symbol: ""}, "KGS": ut.Currency{Currency: "KGS", DisplayName: "Киргизстански сом", Symbol: "KGS"}, "BOB": ut.Currency{Currency: "BOB", DisplayName: "Боливийско боливиано", Symbol: "BOB"}, "JOD": ut.Currency{Currency: "JOD", DisplayName: "Йордански динар", Symbol: "JOD"}, "MNT": ut.Currency{Currency: "MNT", DisplayName: "Монголски тугрик", Symbol: "MNT"}, "MTL": ut.Currency{Currency: "MTL", DisplayName: "Малтийска лира", Symbol: ""}, "PYG": ut.Currency{Currency: "PYG", DisplayName: "Парагвайско гуарани", Symbol: "PYG"}, "BAD": ut.Currency{Currency: "BAD", DisplayName: "Босна и Херцеговина-динар", Symbol: ""}, "PEN": ut.Currency{Currency: "PEN", DisplayName: "Перуански нов сол", Symbol: "PEN"}, "ZMW": ut.Currency{Currency: "ZMW", DisplayName: "Замбийска куача", Symbol: "ZMW"}, "JPY": ut.Currency{Currency: "JPY", DisplayName: "Японска йена", Symbol: "JPY"}, "BND": ut.Currency{Currency: "BND", DisplayName: "Брунейски долар", Symbol: "BND"}, "BRB": ut.Currency{Currency: "BRB", DisplayName: "Бразилско ново крузейро (1967–1986)", Symbol: ""}, "MXP": ut.Currency{Currency: "MXP", DisplayName: "Мексиканско сребърно песо (1861–1992)", Symbol: "MXP"}, "AFN": ut.Currency{Currency: "AFN", DisplayName: "Афганистански афган", Symbol: "AFN"}, "HRD": ut.Currency{Currency: "HRD", DisplayName: "Хърватски динар", Symbol: ""}, "HUF": ut.Currency{Currency: "HUF", DisplayName: "Унгарски форинт", Symbol: "HUF"}, "GTQ": ut.Currency{Currency: "GTQ", DisplayName: "Гватемалски кетцал", Symbol: "GTQ"}, "CUP": ut.Currency{Currency: "CUP", DisplayName: "Кубинско песо", Symbol: "CUP"}, "PAB": ut.Currency{Currency: "PAB", DisplayName: "Панамска балбоа", Symbol: "PAB"}, "SDD": ut.Currency{Currency: "SDD", DisplayName: "Судански динар", Symbol: ""}, "TJS": ut.Currency{Currency: "TJS", DisplayName: "Таджикистански сомони", Symbol: "TJS"}, "UYU": ut.Currency{Currency: "UYU", DisplayName: "Уругвайско песо", Symbol: "UYU"}, "AMD": ut.Currency{Currency: "AMD", DisplayName: "Арменски драм", Symbol: "AMD"}, "ARA": ut.Currency{Currency: "ARA", DisplayName: "Аржентински австрал", Symbol: ""}, "GEL": ut.Currency{Currency: "GEL", DisplayName: "Грузински лари", Symbol: "GEL"}, "MYR": ut.Currency{Currency: "MYR", DisplayName: "Малайзийски рингит", Symbol: "MYR"}, "YDD": ut.Currency{Currency: "YDD", DisplayName: "Йеменски динар", Symbol: ""}, "AED": ut.Currency{Currency: "AED", DisplayName: "Дирхам на Обединените арабски емирства", Symbol: "AED"}}