package be

import "github.com/go-playground/universal-translator"

var currencies = ut.CurrencyFormatValue{"TRY": ut.Currency{Currency: "TRY", DisplayName: "турэцкая ліра", Symbol: "TRY"}, "DOP": ut.Currency{Currency: "DOP", DisplayName: "дамініканскае песа", Symbol: "DOP"}, "OMR": ut.Currency{Currency: "OMR", DisplayName: "аманскі рыал", Symbol: "OMR"}, "PHP": ut.Currency{Currency: "PHP", DisplayName: "філіпінскае песа", Symbol: "PHP"}, "SCR": ut.Currency{Currency: "SCR", DisplayName: "сейшэльская рупія", Symbol: "SCR"}, "TMT": ut.Currency{Currency: "TMT", DisplayName: "туркменскі манат", Symbol: "TMT"}, "MYR": ut.Currency{Currency: "MYR", DisplayName: "малайзійскі рынгіт", Symbol: "MYR"}, "NAD": ut.Currency{Currency: "NAD", DisplayName: "намібійскі долар", Symbol: "NAD"}, "SSP": ut.Currency{Currency: "SSP", DisplayName: "паўднёвасуданскі фунт", Symbol: "SSP"}, "ARS": ut.Currency{Currency: "ARS", DisplayName: "аргенцінскае песа", Symbol: "ARS"}, "CLP": ut.Currency{Currency: "CLP", DisplayName: "чылійскае песа", Symbol: "CLP"}, "DZD": ut.Currency{Currency: "DZD", DisplayName: "алжырскі дынар", Symbol: "DZD"}, "TND": ut.Currency{Currency: "TND", DisplayName: "туніскі дынар", Symbol: "TND"}, "UAH": ut.Currency{Currency: "UAH", DisplayName: "украінская грыўна", Symbol: "UAH"}, "WST": ut.Currency{Currency: "WST", DisplayName: "самаанская тала", Symbol: "WST"}, "AOA": ut.Currency{Currency: "AOA", DisplayName: "ангольская кванза", Symbol: "AOA"}, "BND": ut.Currency{Currency: "BND", DisplayName: "брунейскі долар", Symbol: "BND"}, "EUR": ut.Currency{Currency: "EUR", DisplayName: "еўра", Symbol: "€"}, "ILS": ut.Currency{Currency: "ILS", DisplayName: "новы ізраільскі шэкель", Symbol: "₪"}, "SZL": ut.Currency{Currency: "SZL", DisplayName: "свазілендскі лілангені", Symbol: "SZL"}, "KMF": ut.Currency{Currency: "KMF", DisplayName: "каморскі франк", Symbol: "KMF"}, "LRD": ut.Currency{Currency: "LRD", DisplayName: "ліберыйскі долар", Symbol: "LRD"}, "MGA": ut.Currency{Currency: "MGA", DisplayName: "малагасійскі арыяры", Symbol: "MGA"}, "XXX": ut.Currency{Currency: "XXX", DisplayName: "невядомая валюта", Symbol: ""}, "NZD": ut.Currency{Currency: "NZD", DisplayName: "новазеландскі долар", Symbol: "NZD"}, "PGK": ut.Currency{Currency: "PGK", DisplayName: "кіна", Symbol: "PGK"}, "RWF": ut.Currency{Currency: "RWF", DisplayName: "руандыйскі франк", Symbol: "RWF"}, "BWP": ut.Currency{Currency: "BWP", DisplayName: "батсванская пула", Symbol: "BWP"}, "DKK": ut.Currency{Currency: "DKK", DisplayName: "дацкая крона", Symbol: "DKK"}, "ETB": ut.Currency{Currency: "ETB", DisplayName: "эфіопскі быр", Symbol: "ETB"}, "HTG": ut.Currency{Currency: "HTG", DisplayName: "гаіцянскі гурд", Symbol: "HTG"}, "NGN": ut.Currency{Currency: "NGN", DisplayName: "нігерыйская найра", Symbol: "NGN"}, "VUV": ut.Currency{Currency: "VUV", DisplayName: "вату", Symbol: "VUV"}, "XCD": ut.Currency{Currency: "XCD", DisplayName: "усходнекарыбскі долар", Symbol: "EC$"}, "BGN": ut.Currency{Currency: "BGN", DisplayName: "балгарскі леў", Symbol: "BGN"}, "MOP": ut.Currency{Currency: "MOP", DisplayName: "патака Макаа", Symbol: "MOP"}, "MVR": ut.Currency{Currency: "MVR", DisplayName: "мальдыўская руфія", Symbol: "MVR"}, "PEN": ut.Currency{Currency: "PEN", DisplayName: "перуанскі новы соль", Symbol: "PEN"}, "SLL": ut.Currency{Currency: "SLL", DisplayName: "леонэ", Symbol: "SLL"}, "TTD": ut.Currency{Currency: "TTD", DisplayName: "долар Трынідада і Табага", Symbol: "TTD"}, "BIF": ut.Currency{Currency: "BIF", DisplayName: "бурундзійскі франк", Symbol: "BIF"}, "DJF": ut.Currency{Currency: "DJF", DisplayName: "джыбуційскі франк", Symbol: "DJF"}, "LYD": ut.Currency{Currency: "LYD", DisplayName: "лівійскі дынар", Symbol: "LYD"}, "SGD": ut.Currency{Currency: "SGD", DisplayName: "сінгапурскі долар", Symbol: "SGD"}, "TOP": ut.Currency{Currency: "TOP", DisplayName: "танганская паанга", Symbol: "TOP"}, "HRK": ut.Currency{Currency: "HRK", DisplayName: "харвацкая куна", Symbol: "HRK"}, "JPY": ut.Currency{Currency: "JPY", DisplayName: "японская іена", Symbol: "¥"}, "XOF": ut.Currency{Currency: "XOF", DisplayName: "заходнеафрыканскі франк КФА", Symbol: "CFA"}, "CUC": ut.Currency{Currency: "CUC", DisplayName: "кубінскае канверсоўнае песа", Symbol: "CUC"}, "MAD": ut.Currency{Currency: "MAD", DisplayName: "мараканскі дырхам", Symbol: "MAD"}, "MXN": ut.Currency{Currency: "MXN", DisplayName: "мексіканскае песа", Symbol: "MX$"}, "FJD": ut.Currency{Currency: "FJD", DisplayName: "фіджыйскі долар", Symbol: "FJD"}, "NIO": ut.Currency{Currency: "NIO", DisplayName: "нікарагуанская кордаба", Symbol: "NIO"}, "ZAR": ut.Currency{Currency: "ZAR", DisplayName: "паўднёваафрыканскі ранд", Symbol: "ZAR"}, "GNF": ut.Currency{Currency: "GNF", DisplayName: "гвінейскі франк", Symbol: "GNF"}, "BMD": ut.Currency{Currency: "BMD", DisplayName: "бермудскі долар", Symbol: "BMD"}, "HKD": ut.Currency{Currency: "HKD", DisplayName: "ганконгскі долар", Symbol: "HK$"}, "SDG": ut.Currency{Currency: "SDG", DisplayName: "суданскі фунт", Symbol: "SDG"}, "TWD": ut.Currency{Currency: "TWD", DisplayName: "новы тайваньскі долар", Symbol: "NT$"}, "ZMW": ut.Currency{Currency: "ZMW", DisplayName: "замбійская квача", Symbol: "ZMW"}, "YER": ut.Currency{Currency: "YER", DisplayName: "еменскі рыал", Symbol: "YER"}, "BDT": ut.Currency{Currency: "BDT", DisplayName: "бангладэшская така", Symbol: "BDT"}, "IRR": ut.Currency{Currency: "IRR", DisplayName: "іранскі рыал", Symbol: "IRR"}, "LAK": ut.Currency{Currency: "LAK", DisplayName: "лаоскі кіп", Symbol: "LAK"}, "LKR": ut.Currency{Currency: "LKR", DisplayName: "шры-ланкійская рупія", Symbol: "LKR"}, "RUB": ut.Currency{Currency: "RUB", DisplayName: "расійскі рубель", Symbol: "₽"}, "PYG": ut.Currency{Currency: "PYG", DisplayName: "парагвайскі гуарані", Symbol: "PYG"}, "XPF": ut.Currency{Currency: "XPF", DisplayName: "французскі ціхаакіянскі франк", Symbol: "CFPF"}, "ALL": ut.Currency{Currency: "ALL", DisplayName: "албанскі лек", Symbol: "ALL"}, "CAD": ut.Currency{Currency: "CAD", DisplayName: "канадскі долар", Symbol: "CAD"}, "IQD": ut.Currency{Currency: "IQD", DisplayName: "іракскі дынар", Symbol: "IQD"}, "JOD": ut.Currency{Currency: "JOD", DisplayName: "іарданскі дынар", Symbol: "JOD"}, "MWK": ut.Currency{Currency: "MWK", DisplayName: "малавійская квача", Symbol: "MWK"}, "MDL": ut.Currency{Currency: "MDL", DisplayName: "малдаўскі лей", Symbol: "MDL"}, "PLN": ut.Currency{Currency: "PLN", DisplayName: "польскі злоты", Symbol: "PLN"}, "BOB": ut.Currency{Currency: "BOB", DisplayName: "балівіяна", Symbol: "BOB"}, "BTN": ut.Currency{Currency: "BTN", DisplayName: "бутанскі нгултрум", Symbol: "BTN"}, "CHF": ut.Currency{Currency: "CHF", DisplayName: "швейцарскі франк", Symbol: "CHF"}, "GTQ": ut.Currency{Currency: "GTQ", DisplayName: "гватэмальскі кетсаль", Symbol: "GTQ"}, "THB": ut.Currency{Currency: "THB", DisplayName: "тайскі бат", Symbol: "THB"}, "AED": ut.Currency{Currency: "AED", DisplayName: "дырхем ААЭ", Symbol: "AED"}, "KZT": ut.Currency{Currency: "KZT", DisplayName: "казахстанскі тэнгэ", Symbol: "KZT"}, "LBP": ut.Currency{Currency: "LBP", DisplayName: "ліванскі фунт", Symbol: "LBP"}, "MKD": ut.Currency{Currency: "MKD", DisplayName: "македонскі дэнар", Symbol: "MKD"}, "NOK": ut.Currency{Currency: "NOK", DisplayName: "нарвежская крона", Symbol: "NOK"}, "SYP": ut.Currency{Currency: "SYP", DisplayName: "сірыйскі фунт", Symbol: "SYP"}, "CRC": ut.Currency{Currency: "CRC", DisplayName: "коста-рыканскі калон", Symbol: "CRC"}, "CVE": ut.Currency{Currency: "CVE", DisplayName: "эскуда Каба-Вердэ", Symbol: "CVE"}, "CZK": ut.Currency{Currency: "CZK", DisplayName: "чэшская крона", Symbol: "CZK"}, "FKP": ut.Currency{Currency: "FKP", DisplayName: "фунт Фалклендскіх астравоў", Symbol: "FKP"}, "GBP": ut.Currency{Currency: "GBP", DisplayName: "брытанскі фунт стэрлінгаў", Symbol: "£"}, "BYR": ut.Currency{Currency: "BYR", DisplayName: "беларускі рубель", Symbol: "р."}, "KGS": ut.Currency{Currency: "KGS", DisplayName: "кіргізскі сом", Symbol: "KGS"}, "KYD": ut.Currency{Currency: "KYD", DisplayName: "долар Кайманавых астравоў", Symbol: "KYD"}, "MRO": ut.Currency{Currency: "MRO", DisplayName: "маўрытанская ўгія", Symbol: "MRO"}, "NPR": ut.Currency{Currency: "NPR", DisplayName: "непальская рупія", Symbol: "NPR"}, "MNT": ut.Currency{Currency: "MNT", DisplayName: "мангольскі тугрык", Symbol: "MNT"}, "SAR": ut.Currency{Currency: "SAR", DisplayName: "саудаўскі рыял", Symbol: "SAR"}, "UGX": ut.Currency{Currency: "UGX", DisplayName: "угандыйскі шылінг", Symbol: "UGX"}, "AMD": ut.Currency{Currency: "AMD", DisplayName: "армянскі драм", Symbol: "AMD"}, "AZN": ut.Currency{Currency: "AZN", DisplayName: "азербайджанскі манат", Symbol: "AZN"}, "CNY": ut.Currency{Currency: "CNY", DisplayName: "кітайскі юань", Symbol: "CN¥"}, "COP": ut.Currency{Currency: "COP", DisplayName: "калумбійскае песа", Symbol: "COP"}, "HUF": ut.Currency{Currency: "HUF", DisplayName: "венгерскі форынт", Symbol: "HUF"}, "BRL": ut.Currency{Currency: "BRL", DisplayName: "бразільскі рэал", Symbol: "BRL"}, "GIP": ut.Currency{Currency: "GIP", DisplayName: "гібралтарскі фунт", Symbol: "GIP"}, "SBD": ut.Currency{Currency: "SBD", DisplayName: "долар Саламонавых Астравоў", Symbol: "SBD"}, "VND": ut.Currency{Currency: "VND", DisplayName: "в’етнамскі донг", Symbol: "₫"}, "TZS": ut.Currency{Currency: "TZS", DisplayName: "танзанійскі шылінг", Symbol: "TZS"}, "UZS": ut.Currency{Currency: "UZS", DisplayName: "узбекскі сум", Symbol: "UZS"}, "XAF": ut.Currency{Currency: "XAF", DisplayName: "цэнтральнаафрыканскі франк КФА", Symbol: "FCFA"}, "ANG": ut.Currency{Currency: "ANG", DisplayName: "нідэрландскі антыльскі гульдэн", Symbol: "ANG"}, "GEL": ut.Currency{Currency: "GEL", DisplayName: "грузінскі лары", Symbol: "GEL"}, "GYD": ut.Currency{Currency: "GYD", DisplayName: "гаянскі долар", Symbol: "GYD"}, "KWD": ut.Currency{Currency: "KWD", DisplayName: "кувейцкі дынар", Symbol: "KWD"}, "SHP": ut.Currency{Currency: "SHP", DisplayName: "фунт Святой Алены", Symbol: "SHP"}, "JMD": ut.Currency{Currency: "JMD", DisplayName: "ямайскі долар", Symbol: "JMD"}, "KRW": ut.Currency{Currency: "KRW", DisplayName: "паўднёвакарэйская вона", Symbol: "₩"}, "VEF": ut.Currency{Currency: "VEF", DisplayName: "венесуальскі балівар", Symbol: "VEF"}, "CUP": ut.Currency{Currency: "CUP", DisplayName: "кубінскае песа", Symbol: "CUP"}, "PAB": ut.Currency{Currency: "PAB", DisplayName: "панамскае бальбоа", Symbol: "PAB"}, "SOS": ut.Currency{Currency: "SOS", DisplayName: "самалійскі шылінг", Symbol: "SOS"}, "GMD": ut.Currency{Currency: "GMD", DisplayName: "гамбійскі даласі", Symbol: "GMD"}, "MZN": ut.Currency{Currency: "MZN", DisplayName: "мазамбікскі метыкал", Symbol: "MZN"}, "PKR": ut.Currency{Currency: "PKR", DisplayName: "пакістанская рупія", Symbol: "PKR"}, "RON": ut.Currency{Currency: "RON", DisplayName: "румынскі лей", Symbol: "RON"}, "TJS": ut.Currency{Currency: "TJS", DisplayName: "таджыкскі самані", Symbol: "TJS"}, "KPW": ut.Currency{Currency: "KPW", DisplayName: "паўночнакарэйская вона", Symbol: "KPW"}, "STD": ut.Currency{Currency: "STD", DisplayName: "добра Сан-Тамэ і Прынсіпі", Symbol: "STD"}, "USD": ut.Currency{Currency: "USD", DisplayName: "долар ЗША", Symbol: "$"}, "AUD": ut.Currency{Currency: "AUD", DisplayName: "аўстралійскі долар", Symbol: "A$"}, "BHD": ut.Currency{Currency: "BHD", DisplayName: "бахрэйнскі дынар", Symbol: "BHD"}, "BSD": ut.Currency{Currency: "BSD", DisplayName: "багамскі долар", Symbol: "BSD"}, "BZD": ut.Currency{Currency: "BZD", DisplayName: "белізскі долар", Symbol: "BZD"}, "CDF": ut.Currency{Currency: "CDF", DisplayName: "кангалезскі франк", Symbol: "CDF"}, "KES": ut.Currency{Currency: "KES", DisplayName: "кенійскі шылінг", Symbol: "KES"}, "RSD": ut.Currency{Currency: "RSD", DisplayName: "сербскі дынар", Symbol: "RSD"}, "SEK": ut.Currency{Currency: "SEK", DisplayName: "шведская крона", Symbol: "SEK"}, "AFN": ut.Currency{Currency: "AFN", DisplayName: "афганскі афгані", Symbol: "AFN"}, "BAM": ut.Currency{Currency: "BAM", DisplayName: "канверсоўная марка", Symbol: "BAM"}, "GHS": ut.Currency{Currency: "GHS", DisplayName: "ганскі седзі", Symbol: "GHS"}, "HNL": ut.Currency{Currency: "HNL", DisplayName: "гандураская лемпіра", Symbol: "HNL"}, "ISK": ut.Currency{Currency: "ISK", DisplayName: "ісландская крона", Symbol: "ISK"}, "AWG": ut.Currency{Currency: "AWG", DisplayName: "арубанскі фларын", Symbol: "AWG"}, "IDR": ut.Currency{Currency: "IDR", DisplayName: "інданезійская рупія", Symbol: "IDR"}, "MMK": ut.Currency{Currency: "MMK", DisplayName: "м’янманскі к’ят", Symbol: "MMK"}, "SRD": ut.Currency{Currency: "SRD", DisplayName: "сурынамскі долар", Symbol: "SRD"}, "BBD": ut.Currency{Currency: "BBD", DisplayName: "барбадоскі долар", Symbol: "BBD"}, "INR": ut.Currency{Currency: "INR", DisplayName: "індыйская рупія", Symbol: "₹"}, "MUR": ut.Currency{Currency: "MUR", DisplayName: "маўрыкійская рупія", Symbol: "MUR"}, "EGP": ut.Currency{Currency: "EGP", DisplayName: "егіпецкі фунт", Symbol: "EGP"}, "ERN": ut.Currency{Currency: "ERN", DisplayName: "эрытрэйская накфа", Symbol: "ERN"}, "KHR": ut.Currency{Currency: "KHR", DisplayName: "камбаджыйскі рыэль", Symbol: "KHR"}, "QAR": ut.Currency{Currency: "QAR", DisplayName: "катарскі рыал", Symbol: "QAR"}, "UYU": ut.Currency{Currency: "UYU", DisplayName: "уругвайскае песа", Symbol: "UYU"}}