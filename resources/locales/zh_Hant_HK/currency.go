package zh_Hant_HK

import "github.com/go-playground/universal-translator"

var currencies = ut.CurrencyFormatValue{"OMR": ut.Currency{Currency: "OMR", DisplayName: "阿曼里奧", Symbol: "OMR"}, "SAR": ut.Currency{Currency: "SAR", DisplayName: "沙特阿拉伯里亞爾", Symbol: "SAR"}, "YER": ut.Currency{Currency: "YER", DisplayName: "也門里雅", Symbol: "YER"}, "KES": ut.Currency{Currency: "KES", DisplayName: "肯雅先令", Symbol: "KES"}, "MVR": ut.Currency{Currency: "MVR", DisplayName: "馬爾代夫盧非亞", Symbol: "MVR"}, "HRK": ut.Currency{Currency: "HRK", DisplayName: "克羅地亞庫納", Symbol: "HRK"}, "MUR": ut.Currency{Currency: "MUR", DisplayName: "毛里裘斯盧布", Symbol: "MUR"}, "RSD": ut.Currency{Currency: "RSD", DisplayName: "塞爾維亞第納爾", Symbol: "RSD"}, "XAF": ut.Currency{Currency: "XAF", DisplayName: "中非法郎", Symbol: "FCFA"}, "XPF": ut.Currency{Currency: "XPF", DisplayName: "太平洋法郎", Symbol: "CFPF"}, "CVE": ut.Currency{Currency: "CVE", DisplayName: "佛得角埃斯庫多", Symbol: "CVE"}, "GTQ": ut.Currency{Currency: "GTQ", DisplayName: "危地馬拉格查爾", Symbol: "GTQ"}, "BZD": ut.Currency{Currency: "BZD", DisplayName: "伯利茲元", Symbol: "BZD"}, "SEK": ut.Currency{Currency: "SEK", DisplayName: "瑞典克羅納", Symbol: "SEK"}, "VUV": ut.Currency{Currency: "VUV", DisplayName: "瓦努阿圖瓦圖", Symbol: "VUV"}, "CRC": ut.Currency{Currency: "CRC", DisplayName: "哥斯達黎加科郎", Symbol: "CRC"}, "DOP": ut.Currency{Currency: "DOP", DisplayName: "多米尼加披索", Symbol: "DOP"}, "BWP": ut.Currency{Currency: "BWP", DisplayName: "博茨瓦納普拉", Symbol: "BWP"}, "CAD": ut.Currency{Currency: "CAD", DisplayName: "加拿大元", Symbol: "CA$"}, "LTL": ut.Currency{Currency: "LTL", DisplayName: "立陶宛里塔", Symbol: "LTL"}, "XCD": ut.Currency{Currency: "XCD", DisplayName: "東加勒比元", Symbol: "EC$"}, "ZMW": ut.Currency{Currency: "ZMW", DisplayName: "贊比亞克瓦查", Symbol: "ZMW"}, "AZN": ut.Currency{Currency: "AZN", DisplayName: "亞塞拜疆馬納特", Symbol: "AZN"}, "BIF": ut.Currency{Currency: "BIF", DisplayName: "布隆迪法郎", Symbol: "BIF"}, "KPW": ut.Currency{Currency: "KPW", DisplayName: "北韓圜", Symbol: "KPW"}, "TTD": ut.Currency{Currency: "TTD", DisplayName: "千里達和多巴哥元", Symbol: "TTD"}, "DKK": ut.Currency{Currency: "DKK", DisplayName: "丹麥克羅納", Symbol: "DKK"}, "ITL": ut.Currency{Currency: "ITL", DisplayName: "意大利里拉", Symbol: "ITL"}, "LRD": ut.Currency{Currency: "LRD", DisplayName: "利比利亞元", Symbol: "LRD"}, "MDL": ut.Currency{Currency: "MDL", DisplayName: "摩爾多瓦列伊", Symbol: "MDL"}, "SCR": ut.Currency{Currency: "SCR", DisplayName: "塞舌爾盧比", Symbol: "SCR"}, "NGN": ut.Currency{Currency: "NGN", DisplayName: "尼日利亞奈拉", Symbol: "NGN"}, "PGK": ut.Currency{Currency: "PGK", DisplayName: "巴布亞新幾內亞基那", Symbol: "PGK"}, "NZD": ut.Currency{Currency: "NZD", DisplayName: "紐西蘭元", Symbol: "NZ$"}, "SZL": ut.Currency{Currency: "SZL", DisplayName: "斯威士蘭里朗吉尼", Symbol: "SZL"}, "TOP": ut.Currency{Currency: "TOP", DisplayName: "湯加潘加", Symbol: "TOP"}, "BBD": ut.Currency{Currency: "BBD", DisplayName: "巴巴多斯元", Symbol: "BBD"}, "LAK": ut.Currency{Currency: "LAK", DisplayName: "老撾基普", Symbol: "LAK"}, "SOS": ut.Currency{Currency: "SOS", DisplayName: "索馬里先令", Symbol: "SOS"}, "DJF": ut.Currency{Currency: "DJF", DisplayName: "吉布提法郎", Symbol: "DJF"}, "GHS": ut.Currency{Currency: "GHS", DisplayName: "加納塞地", Symbol: "GHS"}, "KZT": ut.Currency{Currency: "KZT", DisplayName: "哈薩克坦吉", Symbol: "KZT"}, "MZN": ut.Currency{Currency: "MZN", DisplayName: "莫桑比克梅蒂卡爾", Symbol: "MZN"}, "NOK": ut.Currency{Currency: "NOK", DisplayName: "挪威克羅納", Symbol: "NOK"}, "SRD": ut.Currency{Currency: "SRD", DisplayName: "蘇里南元", Symbol: "SRD"}, "ETB": ut.Currency{Currency: "ETB", DisplayName: "埃塞俄比亞比爾", Symbol: "ETB"}, "GMD": ut.Currency{Currency: "GMD", DisplayName: "岡比亞達拉西", Symbol: "GMD"}, "QAR": ut.Currency{Currency: "QAR", DisplayName: "卡塔爾里亞爾", Symbol: "QAR"}, "AED": ut.Currency{Currency: "AED", DisplayName: "阿拉伯聯合酋長國迪爾汗", Symbol: "AED"}, "BAD": ut.Currency{Currency: "BAD", DisplayName: "波斯尼亞-赫塞哥維納第納爾", Symbol: "BAD"}, "MRO": ut.Currency{Currency: "MRO", DisplayName: "毛里塔尼亞烏吉亞", Symbol: "MRO"}, "SBD": ut.Currency{Currency: "SBD", DisplayName: "所羅門群島元", Symbol: "SBD"}, "STD": ut.Currency{Currency: "STD", DisplayName: "聖多美和普林西比多布拉", Symbol: "STD"}, "TZS": ut.Currency{Currency: "TZS", DisplayName: "坦桑尼亞先令", Symbol: "TZS"}, "XOF": ut.Currency{Currency: "XOF", DisplayName: "多哥非洲共同體法郎", Symbol: "CFA"}, "BAM": ut.Currency{Currency: "BAM", DisplayName: "波斯尼亞-赫塞哥維納可轉換馬克", Symbol: "BAM"}, "HKD": ut.Currency{Currency: "HKD", DisplayName: "港元", Symbol: "HK$"}, "TWD": ut.Currency{Currency: "TWD", DisplayName: "新台币", Symbol: "NT$"}, "GEL": ut.Currency{Currency: "GEL", DisplayName: "格魯吉亞拉里", Symbol: "GEL"}, "KRW": ut.Currency{Currency: "KRW", DisplayName: "韩元", Symbol: "₩"}, "SGD": ut.Currency{Currency: "SGD", DisplayName: "新加坡元", Symbol: "SGD"}, "SLL": ut.Currency{Currency: "SLL", DisplayName: "塞拉利昂利昂", Symbol: "SLL"}, "KHR": ut.Currency{Currency: "KHR", DisplayName: "柬埔寨里爾", Symbol: "KHR"}, "KYD": ut.Currency{Currency: "KYD", DisplayName: "開曼群島美元", Symbol: "KYD"}, "RWF": ut.Currency{Currency: "RWF", DisplayName: "盧旺達法郎", Symbol: "RWF"}, "AUD": ut.Currency{Currency: "AUD", DisplayName: "澳元", Symbol: "AU$"}, "AWG": ut.Currency{Currency: "AWG", DisplayName: "阿魯巴盾", Symbol: "AWG"}}
