package bs_Cyrl

import "github.com/go-playground/universal-translator"

var currencies = ut.CurrencyFormatValue{"KWD": ut.Currency{Currency: "KWD", DisplayName: "Кувајтски динар", Symbol: "KWD"}, "TND": ut.Currency{Currency: "TND", DisplayName: "Тунизијски долар", Symbol: "TND"}, "XAU": ut.Currency{Currency: "XAU", DisplayName: "Злато", Symbol: ""}, "ARA": ut.Currency{Currency: "ARA", DisplayName: "Аргентински аустрал", Symbol: ""}, "ARL": ut.Currency{Currency: "ARL", DisplayName: "аргентински пезос леј", Symbol: ""}, "CVE": ut.Currency{Currency: "CVE", DisplayName: "Зеленортски ескудо", Symbol: "CVE"}, "DZD": ut.Currency{Currency: "DZD", DisplayName: "Алжирски динар", Symbol: "DZD"}, "ETB": ut.Currency{Currency: "ETB", DisplayName: "етиопијски бир", Symbol: "ETB"}, "MKD": ut.Currency{Currency: "MKD", DisplayName: "Македонски денар", Symbol: "MKD"}, "MTP": ut.Currency{Currency: "MTP", DisplayName: "Малтешка фунта", Symbol: ""}, "NIC": ut.Currency{Currency: "NIC", DisplayName: "Никарагванска кордоба", Symbol: ""}, "AZM": ut.Currency{Currency: "AZM", DisplayName: "Азербејџански манат (1993–2006)", Symbol: ""}, "BND": ut.Currency{Currency: "BND", DisplayName: "Брунејски долар", Symbol: "BND"}, "BYB": ut.Currency{Currency: "BYB", DisplayName: "Белоруска нова рубља (1994–1999)", Symbol: ""}, "CSD": ut.Currency{Currency: "CSD", DisplayName: "Стари српски динар", Symbol: ""}, "XEU": ut.Currency{Currency: "XEU", DisplayName: "Европска валутна јединица", Symbol: ""}, "CRC": ut.Currency{Currency: "CRC", DisplayName: "Костарикански колон", Symbol: "CRC"}, "GNF": ut.Currency{Currency: "GNF", DisplayName: "Гвинејски франак", Symbol: "GNF"}, "TMT": ut.Currency{Currency: "TMT", DisplayName: "Туркменистански манат", Symbol: "TMT"}, "AMD": ut.Currency{Currency: "AMD", DisplayName: "Јерменски драм", Symbol: "AMD"}, "FJD": ut.Currency{Currency: "FJD", DisplayName: "Фиџи долар", Symbol: "FJD"}, "QAR": ut.Currency{Currency: "QAR", DisplayName: "Катаршки ријал", Symbol: "QAR"}, "THB": ut.Currency{Currency: "THB", DisplayName: "Таи бахт", Symbol: "฿"}, "ARS": ut.Currency{Currency: "ARS", DisplayName: "аргентински пезо", Symbol: "ARS"}, "BSD": ut.Currency{Currency: "BSD", DisplayName: "Бахамски долар", Symbol: "BSD"}, "DEM": ut.Currency{Currency: "DEM", DisplayName: "Немачка марка", Symbol: ""}, "SZL": ut.Currency{Currency: "SZL", DisplayName: "Свази лилангени", Symbol: "SZL"}, "HTG": ut.Currency{Currency: "HTG", DisplayName: "Хаићански гурд", Symbol: "HTG"}, "OMR": ut.Currency{Currency: "OMR", DisplayName: "Омански ријал", Symbol: "OMR"}, "PLZ": ut.Currency{Currency: "PLZ", DisplayName: "Пољски злоти (1950–1995)", Symbol: ""}, "GMD": ut.Currency{Currency: "GMD", DisplayName: "Гамбијски даласи", Symbol: "GMD"}, "KGS": ut.Currency{Currency: "KGS", DisplayName: "Киргистански сом", Symbol: "KGS"}, "NPR": ut.Currency{Currency: "NPR", DisplayName: "Непалски рупи", Symbol: "NPR"}, "UGS": ut.Currency{Currency: "UGS", DisplayName: "Угандски шилинг (1966–1987)", Symbol: ""}, "SDG": ut.Currency{Currency: "SDG", DisplayName: "Суданска фунта", Symbol: "SDG"}, "XPF": ut.Currency{Currency: "XPF", DisplayName: "CFP франак", Symbol: "XPF"}, "BEL": ut.Currency{Currency: "BEL", DisplayName: "Белгијски франак (финансијски)", Symbol: ""}, "BRN": ut.Currency{Currency: "BRN", DisplayName: "Бразилијски нови крузадо", Symbol: ""}, "CAD": ut.Currency{Currency: "CAD", DisplayName: "Канадски долар", Symbol: "CAD"}, "LVR": ut.Currency{Currency: "LVR", DisplayName: "атвијска рубља", Symbol: ""}, "GIP": ut.Currency{Currency: "GIP", DisplayName: "Гибралташка фунта", Symbol: "GIP"}, "INR": ut.Currency{Currency: "INR", DisplayName: "Индијски Рупи", Symbol: "₹"}, "KMF": ut.Currency{Currency: "KMF", DisplayName: "Коморски франак", Symbol: "KMF"}, "XAF": ut.Currency{Currency: "XAF", DisplayName: "CFA франак BEAC", Symbol: "FCFA"}, "MCF": ut.Currency{Currency: "MCF", DisplayName: "монегаскански франак", Symbol: ""}, "XDR": ut.Currency{Currency: "XDR", DisplayName: "Посебна цртаћа права", Symbol: ""}, "XFO": ut.Currency{Currency: "XFO", DisplayName: "Француски златни франак", Symbol: ""}, "YUD": ut.Currency{Currency: "YUD", DisplayName: "Југословенски тврди динар", Symbol: ""}, "MDL": ut.Currency{Currency: "MDL", DisplayName: "Молдовски љу", Symbol: "MDL"}, "PTE": ut.Currency{Currency: "PTE", DisplayName: "Португалски ескудо", Symbol: ""}, "VUV": ut.Currency{Currency: "VUV", DisplayName: "Вануатски вату", Symbol: "VUV"}, "XPT": ut.Currency{Currency: "XPT", DisplayName: "Платина", Symbol: ""}, "TZS": ut.Currency{Currency: "TZS", DisplayName: "Танзанијски шилинг", Symbol: "TZS"}, "ADP": ut.Currency{Currency: "ADP", DisplayName: "Андорска пезета", Symbol: ""}, "BBD": ut.Currency{Currency: "BBD", DisplayName: "Барбадошки долар", Symbol: "BBD"}, "EUR": ut.Currency{Currency: "EUR", DisplayName: "Евро", Symbol: "€"}, "SEK": ut.Currency{Currency: "SEK", DisplayName: "Шведска круна", Symbol: "SEK"}, "GTQ": ut.Currency{Currency: "GTQ", DisplayName: "Гватемалски квецал", Symbol: "GTQ"}, "MAD": ut.Currency{Currency: "MAD", DisplayName: "Марокански дирхам", Symbol: "MAD"}, "SBD": ut.Currency{Currency: "SBD", DisplayName: "Соломонско-острвски долар", Symbol: "SBD"}, "TMM": ut.Currency{Currency: "TMM", DisplayName: "Туркменистански манат (1993–2009)", Symbol: ""}, "TRL": ut.Currency{Currency: "TRL", DisplayName: "Стара турска лира", Symbol: ""}, "USN": ut.Currency{Currency: "USN", DisplayName: "САД долар (следећи дан)", Symbol: ""}, "ZMW": ut.Currency{Currency: "ZMW", DisplayName: "Замбијска квача", Symbol: "ZMW"}, "BAN": ut.Currency{Currency: "BAN", DisplayName: "босанско-херцеговачки нови динар", Symbol: ""}, "BRZ": ut.Currency{Currency: "BRZ", DisplayName: "стари бразилски крузеиро", Symbol: ""}, "LBP": ut.Currency{Currency: "LBP", DisplayName: "Лебанска фунта", Symbol: "LBP"}, "MGA": ut.Currency{Currency: "MGA", DisplayName: "Малагасијски ариари", Symbol: "MGA"}, "NLG": ut.Currency{Currency: "NLG", DisplayName: "Холандски гулден", Symbol: ""}, "CNY": ut.Currency{Currency: "CNY", DisplayName: "Кинески јуан ренминби", Symbol: "CNY"}, "HUF": ut.Currency{Currency: "HUF", DisplayName: "Мађарска форинта", Symbol: "HUF"}, "JPY": ut.Currency{Currency: "JPY", DisplayName: "Јапански јен", Symbol: "¥"}, "MNT": ut.Currency{Currency: "MNT", DisplayName: "Монголски тугрик", Symbol: "MNT"}, "CUC": ut.Currency{Currency: "CUC", DisplayName: "кубански конвертибилни песо", Symbol: "CUC"}, "LUC": ut.Currency{Currency: "LUC", DisplayName: "Луксембуршки конвертибилни франак", Symbol: ""}, "SAR": ut.Currency{Currency: "SAR", DisplayName: "Саудијски ријал", Symbol: "SAR"}, "AOR": ut.Currency{Currency: "AOR", DisplayName: "Анголска кванза реађустадо (1995–1999)", Symbol: ""}, "MZN": ut.Currency{Currency: "MZN", DisplayName: "Мозамбијски метикал", Symbol: "MZN"}, "KRO": ut.Currency{Currency: "KRO", DisplayName: "стари јужнокорејски вон", Symbol: ""}, "LRD": ut.Currency{Currency: "LRD", DisplayName: "Либеријски долар", Symbol: "LRD"}, "RUR": ut.Currency{Currency: "RUR", DisplayName: "Руска рубља (1991–1998)", Symbol: ""}, "BAM": ut.Currency{Currency: "BAM", DisplayName: "Конвертибилна марка", Symbol: "КМ"}, "BOV": ut.Currency{Currency: "BOV", DisplayName: "Боливијски мвдол", Symbol: ""}, "GHS": ut.Currency{Currency: "GHS", DisplayName: "Гански цеди", Symbol: "GHS"}, "ILR": ut.Currency{Currency: "ILR", DisplayName: "стари израелски шекели", Symbol: ""}, "SUR": ut.Currency{Currency: "SUR", DisplayName: "Совјетска рубља", Symbol: ""}, "WST": ut.Currency{Currency: "WST", DisplayName: "Самоанска тала", Symbol: "WST"}, "AOK": ut.Currency{Currency: "AOK", DisplayName: "анголијска кванза (1977–1990)", Symbol: ""}, "BGN": ut.Currency{Currency: "BGN", DisplayName: "Бугарски лев", Symbol: "BGN"}, "BTN": ut.Currency{Currency: "BTN", DisplayName: "Бутански нгултрум", Symbol: "BTN"}, "NAD": ut.Currency{Currency: "NAD", DisplayName: "намбијски долар", Symbol: "NAD"}, "SLL": ut.Currency{Currency: "SLL", DisplayName: "Сијера-леоншки леоне", Symbol: "SLL"}, "MTL": ut.Currency{Currency: "MTL", DisplayName: "Малтешка лира", Symbol: ""}, "PEN": ut.Currency{Currency: "PEN", DisplayName: "Перуански нуево сол", Symbol: "PEN"}, "ROL": ut.Currency{Currency: "ROL", DisplayName: "Стари румунски љу", Symbol: ""}, "SCR": ut.Currency{Currency: "SCR", DisplayName: "Сејшелска рупија", Symbol: "SCR"}, "USD": ut.Currency{Currency: "USD", DisplayName: "Амерички долар", Symbol: "USD"}, "CLE": ut.Currency{Currency: "CLE", DisplayName: "чилеански ескудо", Symbol: ""}, "CSK": ut.Currency{Currency: "CSK", DisplayName: "Чехословачка тврда круна", Symbol: ""}, "ESB": ut.Currency{Currency: "ESB", DisplayName: "Шпанска пезета (конвертибилнирачун)", Symbol: ""}, "GEK": ut.Currency{Currency: "GEK", DisplayName: "Грузијски купон ларит", Symbol: ""}, "RON": ut.Currency{Currency: "RON", DisplayName: "Румунски леу", Symbol: "RON"}, "ZAR": ut.Currency{Currency: "ZAR", DisplayName: "Јужно-афрички ранд", Symbol: "ZAR"}, "AED": ut.Currency{Currency: "AED", DisplayName: "Уједињени арапски емирати дирхам", Symbol: "AED"}, "AUD": ut.Currency{Currency: "AUD", DisplayName: "Аустралијски долар", Symbol: "AUD"}, "BIF": ut.Currency{Currency: "BIF", DisplayName: "Бурундски франак", Symbol: "BIF"}, "GWE": ut.Currency{Currency: "GWE", DisplayName: "португалска гвинеја ескудо", Symbol: ""}, "LKR": ut.Currency{Currency: "LKR", DisplayName: "Шриланкански рупи", Symbol: "LKR"}, "MAF": ut.Currency{Currency: "MAF", DisplayName: "Марокански франак", Symbol: ""}, "XPD": ut.Currency{Currency: "XPD", DisplayName: "Паладијум", Symbol: ""}, "XTS": ut.Currency{Currency: "XTS", DisplayName: "Код тестиране валуте", Symbol: ""}, "YER": ut.Currency{Currency: "YER", DisplayName: "Јеменски риал", Symbol: "YER"}, "FKP": ut.Currency{Currency: "FKP", DisplayName: "Фокландска острва фунта", Symbol: "FKP"}, "GEL": ut.Currency{Currency: "GEL", DisplayName: "Грузијски лари", Symbol: "GEL"}, "JOD": ut.Currency{Currency: "JOD", DisplayName: "Јордански динар", Symbol: "JOD"}, "LTL": ut.Currency{Currency: "LTL", DisplayName: "Литвански литас", Symbol: "LTL"}, "ESP": ut.Currency{Currency: "ESP", DisplayName: "Шпанска пезета", Symbol: ""}, "MVR": ut.Currency{Currency: "MVR", DisplayName: "Малдивијска руфија", Symbol: "MVR"}, "MZE": ut.Currency{Currency: "MZE", DisplayName: "Мозамбијски ескудо", Symbol: ""}, "ARM": ut.Currency{Currency: "ARM", DisplayName: "аргентински пезос монедо национал", Symbol: ""}, "CYP": ut.Currency{Currency: "CYP", DisplayName: "Кипарска фунта", Symbol: ""}, "KZT": ut.Currency{Currency: "KZT", DisplayName: "Казахстански тенџ", Symbol: "KZT"}, "GWP": ut.Currency{Currency: "GWP", DisplayName: "Гвинеја Бисао Пезо", Symbol: ""}, "UYI": ut.Currency{Currency: "UYI", DisplayName: "Уругвајски пезо ен унидадес индексадас", Symbol: ""}, "XBB": ut.Currency{Currency: "XBB", DisplayName: "Европска новчана јединица", Symbol: ""}, "MYR": ut.Currency{Currency: "MYR", DisplayName: "Малезијски ринггит", Symbol: "MYR"}, "PEI": ut.Currency{Currency: "PEI", DisplayName: "Перуански инти", Symbol: ""}, "AFA": ut.Currency{Currency: "AFA", DisplayName: "Авганистански авган (1927–2002)", Symbol: ""}, "KRW": ut.Currency{Currency: "KRW", DisplayName: "Јужнокорејски Вон", Symbol: "₩"}, "LUL": ut.Currency{Currency: "LUL", DisplayName: "Луксембуршки финансијски франак", Symbol: ""}, "MXP": ut.Currency{Currency: "MXP", DisplayName: "Мексички сребрни пезо (1861–1992)", Symbol: ""}, "YUR": ut.Currency{Currency: "YUR", DisplayName: "југословенски реформирани динар", Symbol: ""}, "BAD": ut.Currency{Currency: "BAD", DisplayName: "Босанско-Херцеговачки динар", Symbol: ""}, "BGO": ut.Currency{Currency: "BGO", DisplayName: "стари бугарски лев", Symbol: ""}, "DJF": ut.Currency{Currency: "DJF", DisplayName: "Џибутански франак", Symbol: "DJF"}, "HNL": ut.Currency{Currency: "HNL", DisplayName: "Хондурашка лемпира", Symbol: "HNL"}, "SDD": ut.Currency{Currency: "SDD", DisplayName: "Стари судански динар", Symbol: ""}, "SRG": ut.Currency{Currency: "SRG", DisplayName: "Суринамски гилдер", Symbol: ""}, "XXX": ut.Currency{Currency: "XXX", DisplayName: "Непозната или неважећа валута", Symbol: ""}, "AON": ut.Currency{Currency: "AON", DisplayName: "Анголијска нова кванза (1990–2000)", Symbol: ""}, "GYD": ut.Currency{Currency: "GYD", DisplayName: "Гујански долар", Symbol: "GYD"}, "STD": ut.Currency{Currency: "STD", DisplayName: "Сао Томе и Принципе добра", Symbol: "STD"}, "VNN": ut.Currency{Currency: "VNN", DisplayName: "стари вијетнамски донг", Symbol: ""}, "ISK": ut.Currency{Currency: "ISK", DisplayName: "Исландска круна", Symbol: "ISK"}, "KPW": ut.Currency{Currency: "KPW", DisplayName: "Севернокорејски вон", Symbol: "KPW"}, "TJR": ut.Currency{Currency: "TJR", DisplayName: "Таџихистанска рубља", Symbol: ""}, "YDD": ut.Currency{Currency: "YDD", DisplayName: "Јеменски динар", Symbol: ""}, "BEC": ut.Currency{Currency: "BEC", DisplayName: "Белгијски франак (конвертибилни)", Symbol: ""}, "BGL": ut.Currency{Currency: "BGL", DisplayName: "Бугарски тврди лев", Symbol: ""}, "DOP": ut.Currency{Currency: "DOP", DisplayName: "Доминикански пезо", Symbol: "DOP"}, "FRF": ut.Currency{Currency: "FRF", DisplayName: "Француски франак", Symbol: ""}, "MMK": ut.Currency{Currency: "MMK", DisplayName: "Мјанмашки кјат", Symbol: "MMK"}, "VEB": ut.Currency{Currency: "VEB", DisplayName: "Венецуелански боливар (1871–2008)", Symbol: ""}, "ALL": ut.Currency{Currency: "ALL", DisplayName: "Албански лек", Symbol: "ALL"}, "CHW": ut.Currency{Currency: "CHW", DisplayName: "WIR франак", Symbol: ""}, "CNX": ut.Currency{Currency: "CNX", DisplayName: "долар кинеске народне банке", Symbol: ""}, "KRH": ut.Currency{Currency: "KRH", DisplayName: "јужнокорејски хван", Symbol: ""}, "EEK": ut.Currency{Currency: "EEK", DisplayName: "Естонска кроон", Symbol: ""}, "GHC": ut.Currency{Currency: "GHC", DisplayName: "Гански цеди (1979–2007)", Symbol: ""}, "HRK": ut.Currency{Currency: "HRK", DisplayName: "Хрватска куна", Symbol: "kn"}, "MZM": ut.Currency{Currency: "MZM", DisplayName: "Стари мозамбијски метикал", Symbol: ""}, "SDP": ut.Currency{Currency: "SDP", DisplayName: "Стара суданска фунта", Symbol: ""}, "XFU": ut.Currency{Currency: "XFU", DisplayName: "Француски UIC-франак", Symbol: ""}, "BZD": ut.Currency{Currency: "BZD", DisplayName: "Белизе долар", Symbol: "BZD"}, "MLF": ut.Currency{Currency: "MLF", DisplayName: "Малијански франак", Symbol: ""}, "MXN": ut.Currency{Currency: "MXN", DisplayName: "Мексички пезо", Symbol: "MXN"}, "NIO": ut.Currency{Currency: "NIO", DisplayName: "Никарагванска златна кордоба", Symbol: "NIO"}, "AZN": ut.Currency{Currency: "AZN", DisplayName: "Азербејџански манат", Symbol: "AZN"}, "MUR": ut.Currency{Currency: "MUR", DisplayName: "Маурицијски рупи", Symbol: "MUR"}, "RSD": ut.Currency{Currency: "RSD", DisplayName: "Српски динар", Symbol: "дин."}, "MWK": ut.Currency{Currency: "MWK", DisplayName: "Малавијска квача", Symbol: "MWK"}, "ZRZ": ut.Currency{Currency: "ZRZ", DisplayName: "Заирски заир", Symbol: ""}, "CHF": ut.Currency{Currency: "CHF", DisplayName: "Швајцарски франак", Symbol: "CHF"}, "CZK": ut.Currency{Currency: "CZK", DisplayName: "Чешка круна", Symbol: "Кч"}, "DKK": ut.Currency{Currency: "DKK", DisplayName: "Данска круна", Symbol: "DKK"}, "LSL": ut.Currency{Currency: "LSL", DisplayName: "Лесото лоти", Symbol: ""}, "IQD": ut.Currency{Currency: "IQD", DisplayName: "Ирачки динар", Symbol: "IQD"}, "XBA": ut.Currency{Currency: "XBA", DisplayName: "Европска композитна јединица", Symbol: ""}, "TJS": ut.Currency{Currency: "TJS", DisplayName: "Таљихистански сомони", Symbol: "TJS"}, "TRY": ut.Currency{Currency: "TRY", DisplayName: "Турска лира", Symbol: "Тл"}, "BOB": ut.Currency{Currency: "BOB", DisplayName: "Боливијски Боливиано", Symbol: "BOB"}, "IEP": ut.Currency{Currency: "IEP", DisplayName: "Ирска фунта", Symbol: ""}, "LVL": ut.Currency{Currency: "LVL", DisplayName: "Латвијски лати", Symbol: "LVL"}, "NZD": ut.Currency{Currency: "NZD", DisplayName: "Новозеландски долар", Symbol: "NZD"}, "PES": ut.Currency{Currency: "PES", DisplayName: "Перуански сол", Symbol: ""}, "VND": ut.Currency{Currency: "VND", DisplayName: "Вијетнамски донг", Symbol: "₫"}, "XOF": ut.Currency{Currency: "XOF", DisplayName: "CFA франак BCEAO", Symbol: "CFA"}, "ZRN": ut.Currency{Currency: "ZRN", DisplayName: "Заирски нови заир", Symbol: ""}, "BHD": ut.Currency{Currency: "BHD", DisplayName: "Бахреински динар", Symbol: "BHD"}, "GRD": ut.Currency{Currency: "GRD", DisplayName: "Грчка драхма", Symbol: ""}, "PAB": ut.Currency{Currency: "PAB", DisplayName: "Панамска балбоа", Symbol: "PAB"}, "SYP": ut.Currency{Currency: "SYP", DisplayName: "Сиријска фунта", Symbol: "SYP"}, "ZMK": ut.Currency{Currency: "ZMK", DisplayName: "Замбијска квача (1968–2012)", Symbol: ""}, "BYR": ut.Currency{Currency: "BYR", DisplayName: "Белоруска рубља", Symbol: "BYR"}, "EGP": ut.Currency{Currency: "EGP", DisplayName: "Египатска фунта", Symbol: "EGP"}, "FIM": ut.Currency{Currency: "FIM", DisplayName: "Финска марка", Symbol: ""}, "MGF": ut.Currency{Currency: "MGF", DisplayName: "Малагасијски франак", Symbol: ""}, "BDT": ut.Currency{Currency: "BDT", DisplayName: "Бангладешка така", Symbol: "BDT"}, "DDM": ut.Currency{Currency: "DDM", DisplayName: "Источно-немачка марка", Symbol: ""}, "XCD": ut.Currency{Currency: "XCD", DisplayName: "Источно-карибски долар", Symbol: "XCD"}, "MKN": ut.Currency{Currency: "MKN", DisplayName: "стари македонски денар", Symbol: ""}, "PGK": ut.Currency{Currency: "PGK", DisplayName: "Папуа ново-гвинејшка кина", Symbol: "PGK"}, "RHD": ut.Currency{Currency: "RHD", DisplayName: "Родејскидолар", Symbol: ""}, "UAK": ut.Currency{Currency: "UAK", DisplayName: "Украјински карбованети", Symbol: ""}, "IRR": ut.Currency{Currency: "IRR", DisplayName: "Ирански риал", Symbol: "IRR"}, "KES": ut.Currency{Currency: "KES", DisplayName: "Кенијски шилинг", Symbol: "KES"}, "KHR": ut.Currency{Currency: "KHR", DisplayName: "Камбоџијски риел", Symbol: "KHR"}, "MXV": ut.Currency{Currency: "MXV", DisplayName: "Мексички унидад де инверсион (UDI)", Symbol: ""}, "BGM": ut.Currency{Currency: "BGM", DisplayName: "бугарски социјалистички лев", Symbol: ""}, "BRL": ut.Currency{Currency: "BRL", DisplayName: "Бразилски Реал", Symbol: "BRL"}, "CLP": ut.Currency{Currency: "CLP", DisplayName: "Чилеански пезо", Symbol: "CLP"}, "HKD": ut.Currency{Currency: "HKD", DisplayName: "Хонгконшки долар", Symbol: "HKD"}, "USS": ut.Currency{Currency: "USS", DisplayName: "САД долар (исти дан)", Symbol: ""}, "BRC": ut.Currency{Currency: "BRC", DisplayName: "Бразилијски крузадо", Symbol: ""}, "BRR": ut.Currency{Currency: "BRR", DisplayName: "Бразилски крузеиро", Symbol: ""}, "ECV": ut.Currency{Currency: "ECV", DisplayName: "Еквадорски унидад де валор константе", Symbol: ""}, "LYD": ut.Currency{Currency: "LYD", DisplayName: "Либијски динар", Symbol: "LYD"}, "TWD": ut.Currency{Currency: "TWD", DisplayName: "Нови тајвански долар", Symbol: "NT$"}, "VEF": ut.Currency{Currency: "VEF", DisplayName: "Венецуелански боливар", Symbol: "VEF"}, "XRE": ut.Currency{Currency: "XRE", DisplayName: "RINET фонд", Symbol: ""}, "LAK": ut.Currency{Currency: "LAK", DisplayName: "Лаошки кип", Symbol: "LAK"}, "LTT": ut.Currency{Currency: "LTT", DisplayName: "Литвански талонас", Symbol: ""}, "MRO": ut.Currency{Currency: "MRO", DisplayName: "Мауританијска угвија", Symbol: "MRO"}, "SRD": ut.Currency{Currency: "SRD", DisplayName: "СУринамски долар", Symbol: "SRD"}, "BEF": ut.Currency{Currency: "BEF", DisplayName: "Белгијски франак", Symbol: ""}, "SOS": ut.Currency{Currency: "SOS", DisplayName: "СОмалијски шилинг", Symbol: "SOS"}, "XBD": ut.Currency{Currency: "XBD", DisplayName: "Европска јединица рачуна (XBD)", Symbol: ""}, "AOA": ut.Currency{Currency: "AOA", DisplayName: "анголска кванза", Symbol: "AOA"}, "BOL": ut.Currency{Currency: "BOL", DisplayName: "стари боливијски боливијано", Symbol: ""}, "MDC": ut.Currency{Currency: "MDC", DisplayName: "молдовански купон", Symbol: ""}, "YUM": ut.Currency{Currency: "YUM", DisplayName: "Југословенски нови динар", Symbol: ""}, "TTD": ut.Currency{Currency: "TTD", DisplayName: "Тринидад тобагошки долар", Symbol: "TTD"}, "XBC": ut.Currency{Currency: "XBC", DisplayName: "Европска јединица рачуна (XBC)", Symbol: ""}, "ISJ": ut.Currency{Currency: "ISJ", DisplayName: "стара исландска круна", Symbol: ""}, "JMD": ut.Currency{Currency: "JMD", DisplayName: "Јамајски долар", Symbol: "JMD"}, "SGD": ut.Currency{Currency: "SGD", DisplayName: "Сингапурски долар", Symbol: "SGD"}, "SIT": ut.Currency{Currency: "SIT", DisplayName: "Словеначки толар", Symbol: ""}, "ALK": ut.Currency{Currency: "ALK", DisplayName: "стари албански лек", Symbol: ""}, "COP": ut.Currency{Currency: "COP", DisplayName: "Колумбијски пезо", Symbol: "COP"}, "ILP": ut.Currency{Currency: "ILP", DisplayName: "Израелска фунта", Symbol: ""}, "SHP": ut.Currency{Currency: "SHP", DisplayName: "Св. јеленска фунта", Symbol: "SHP"}, "BUK": ut.Currency{Currency: "BUK", DisplayName: "Бурмански кјат", Symbol: ""}, "CUP": ut.Currency{Currency: "CUP", DisplayName: "Кубански пезо", Symbol: "CUP"}, "ZWD": ut.Currency{Currency: "ZWD", DisplayName: "Зимбабвејски долар", Symbol: ""}, "ZWR": ut.Currency{Currency: "ZWR", DisplayName: "Зимбабвеански долар (2008)", Symbol: ""}, "CLF": ut.Currency{Currency: "CLF", DisplayName: "Чилеовски унидадес се фоменто", Symbol: ""}, "GQE": ut.Currency{Currency: "GQE", DisplayName: "Екваторијално-гвинејски еквеле", Symbol: ""}, "NGN": ut.Currency{Currency: "NGN", DisplayName: "Нигеријска наира", Symbol: "NGN"}, "SVC": ut.Currency{Currency: "SVC", DisplayName: "Салвадорски колон", Symbol: ""}, "ARP": ut.Currency{Currency: "ARP", DisplayName: "аргентински пезо (1983–1985)", Symbol: ""}, "BMD": ut.Currency{Currency: "BMD", DisplayName: "Бермудски долар", Symbol: "BMD"}, "BRE": ut.Currency{Currency: "BRE", DisplayName: "Бразилски крузеиро (1990–1993)", Symbol: ""}, "CHE": ut.Currency{Currency: "CHE", DisplayName: "WIR евро", Symbol: ""}, "UYU": ut.Currency{Currency: "UYU", DisplayName: "Уругвајски пезо", Symbol: "UYU"}, "GBP": ut.Currency{Currency: "GBP", DisplayName: "Британска фунта", Symbol: "GBP"}, "UGX": ut.Currency{Currency: "UGX", DisplayName: "Угандски шилинг", Symbol: "UGX"}, "PKR": ut.Currency{Currency: "PKR", DisplayName: "Пакистански рупи", Symbol: "PKR"}, "ANG": ut.Currency{Currency: "ANG", DisplayName: "Холандски антили гилдер", Symbol: "ANG"}, "COU": ut.Currency{Currency: "COU", DisplayName: "Унидад де валоршки реал", Symbol: ""}, "LUF": ut.Currency{Currency: "LUF", DisplayName: "Луксембуршки франак", Symbol: ""}, "NOK": ut.Currency{Currency: "NOK", DisplayName: "Норвешка круна", Symbol: "NOK"}, "RUB": ut.Currency{Currency: "RUB", DisplayName: "Руска рубља", Symbol: "RUB"}, "XAG": ut.Currency{Currency: "XAG", DisplayName: "Сребро", Symbol: ""}, "BWP": ut.Currency{Currency: "BWP", DisplayName: "Боцванска пула", Symbol: "BWP"}, "GNS": ut.Currency{Currency: "GNS", DisplayName: "Гвинејски сили", Symbol: ""}, "IDR": ut.Currency{Currency: "IDR", DisplayName: "Индонезијска рупиа", Symbol: "IDR"}, "PLN": ut.Currency{Currency: "PLN", DisplayName: "Пољски злот", Symbol: "зл"}, "ATS": ut.Currency{Currency: "ATS", DisplayName: "Аустријски шилинг", Symbol: ""}, "ITL": ut.Currency{Currency: "ITL", DisplayName: "Италијанска лира", Symbol: ""}, "KYD": ut.Currency{Currency: "KYD", DisplayName: "Кајманска острва долар", Symbol: "KYD"}, "TOP": ut.Currency{Currency: "TOP", DisplayName: "Тонгоншка Панга", Symbol: "TOP"}, "BRB": ut.Currency{Currency: "BRB", DisplayName: "Бразилски нови крузеиро (1967–1986)", Symbol: ""}, "CDF": ut.Currency{Currency: "CDF", DisplayName: "Конголски франак", Symbol: "CDF"}, "ERN": ut.Currency{Currency: "ERN", DisplayName: "Еритреанска накфа", Symbol: "ERN"}, "SKK": ut.Currency{Currency: "SKK", DisplayName: "Словачка круна", Symbol: ""}, "ZAL": ut.Currency{Currency: "ZAL", DisplayName: "Јужно-афрички ранд (финансијски)", Symbol: ""}, "UYP": ut.Currency{Currency: "UYP", DisplayName: "Уругвајски пезо (1975–1993)", Symbol: ""}, "ECS": ut.Currency{Currency: "ECS", DisplayName: "Еквадорски сакр", Symbol: ""}, "MOP": ut.Currency{Currency: "MOP", DisplayName: "Маканишка патака", Symbol: "MOP"}, "PYG": ut.Currency{Currency: "PYG", DisplayName: "Парагвајски гуарни", Symbol: "PYG"}, "TPE": ut.Currency{Currency: "TPE", DisplayName: "Тиморшки ескудо", Symbol: ""}, "AFN": ut.Currency{Currency: "AFN", DisplayName: "Авганистански авган", Symbol: "AFN"}, "ZWL": ut.Currency{Currency: "ZWL", DisplayName: "Зимбабвеански долар (2009)", Symbol: ""}, "BOP": ut.Currency{Currency: "BOP", DisplayName: "Боливијски пезо", Symbol: ""}, "PHP": ut.Currency{Currency: "PHP", DisplayName: "Филипински пезо", Symbol: "PHP"}, "UZS": ut.Currency{Currency: "UZS", DisplayName: "Узбекистански сом", Symbol: "UZS"}, "UAH": ut.Currency{Currency: "UAH", DisplayName: "Украјинска хривња", Symbol: "UAH"}, "YUN": ut.Currency{Currency: "YUN", DisplayName: "Југословенски конвертибилни динар", Symbol: ""}, "ESA": ut.Currency{Currency: "ESA", DisplayName: "Шпанска пезета (рачун)", Symbol: ""}, "HRD": ut.Currency{Currency: "HRD", DisplayName: "Хрватски динар", Symbol: ""}, "ILS": ut.Currency{Currency: "ILS", DisplayName: "Израелски нови шекел", Symbol: "ILS"}, "RWF": ut.Currency{Currency: "RWF", DisplayName: "Руандански франак", Symbol: "RWF"}, "AWG": ut.Currency{Currency: "AWG", DisplayName: "Арубански флорин", Symbol: "AWG"}}