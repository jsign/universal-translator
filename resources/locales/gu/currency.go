package gu

import "github.com/go-playground/universal-translator"

var currencies = ut.CurrencyFormatValue{"UGX": ut.Currency{Currency: "UGX", DisplayName: "યુગાંડન શિલિંગ", Symbol: "UGX"}, "HTG": ut.Currency{Currency: "HTG", DisplayName: "હાઇટિઇન ગોર્ડ", Symbol: "HTG"}, "MYR": ut.Currency{Currency: "MYR", DisplayName: "મલેશિયન રિંગ્ગેટ", Symbol: "MYR"}, "SHP": ut.Currency{Currency: "SHP", DisplayName: "સેંટ હેલેના પાઉન્ડ", Symbol: "SHP"}, "KMF": ut.Currency{Currency: "KMF", DisplayName: "કોમોરિઅન ફ્રેંક", Symbol: "KMF"}, "NAD": ut.Currency{Currency: "NAD", DisplayName: "નામિબિયન ડોલર", Symbol: "NAD"}, "LKR": ut.Currency{Currency: "LKR", DisplayName: "શ્રી લંકન રૂપી", Symbol: "LKR"}, "MVR": ut.Currency{Currency: "MVR", DisplayName: "માલ્દિવિયન રુફિયા", Symbol: "MVR"}, "VUV": ut.Currency{Currency: "VUV", DisplayName: "વનૌતુ વાતુ", Symbol: "VUV"}, "AFN": ut.Currency{Currency: "AFN", DisplayName: "અફ્ગાન અફ્ગાની", Symbol: "AFN"}, "KGS": ut.Currency{Currency: "KGS", DisplayName: "કિર્ગિસ્તાની સોમ", Symbol: "KGS"}, "KYD": ut.Currency{Currency: "KYD", DisplayName: "કેયમેન આઇલેંડ્સ ડોલર", Symbol: "KYD"}, "LTL": ut.Currency{Currency: "LTL", DisplayName: "લિથુએનિયન લિતાસ", Symbol: "LTL"}, "PEN": ut.Currency{Currency: "PEN", DisplayName: "પેરુવિયન ન્યુવો સોલ", Symbol: "PEN"}, "TTD": ut.Currency{Currency: "TTD", DisplayName: "ત્રિનિદાદ અને ટોબેગો ડોલર", Symbol: "TTD"}, "AED": ut.Currency{Currency: "AED", DisplayName: "યુનાઈટેડ આરબ અમિરાત દિરહામ", Symbol: "AED"}, "AWG": ut.Currency{Currency: "AWG", DisplayName: "અરુબન ફ્લોરિન", Symbol: "AWG"}, "JPY": ut.Currency{Currency: "JPY", DisplayName: "જાપાનીઝ યેન", Symbol: "JP¥"}, "BMD": ut.Currency{Currency: "BMD", DisplayName: "બર્મુડન ડોલર", Symbol: "BMD"}, "DKK": ut.Currency{Currency: "DKK", DisplayName: "ડેનિશ ક્રોન", Symbol: "DKK"}, "MXN": ut.Currency{Currency: "MXN", DisplayName: "મેક્સિકન પેસો", Symbol: "MX$"}, "NGN": ut.Currency{Currency: "NGN", DisplayName: "નાઇજીરિયન નૈરા", Symbol: "NGN"}, "RON": ut.Currency{Currency: "RON", DisplayName: "રોમાનિયન લેઉ", Symbol: "RON"}, "THB": ut.Currency{Currency: "THB", DisplayName: "થાઇ બાહ્ત", Symbol: "฿"}, "TJS": ut.Currency{Currency: "TJS", DisplayName: "તાજિકિસ્તાની સોમોની", Symbol: "TJS"}, "ARS": ut.Currency{Currency: "ARS", DisplayName: "અર્જેન્ટીના પેસો", Symbol: "ARS"}, "KHR": ut.Currency{Currency: "KHR", DisplayName: "કેમ્બોડિયન રીઅલ", Symbol: "KHR"}, "EUR": ut.Currency{Currency: "EUR", DisplayName: "યુરો", Symbol: "€"}, "GNF": ut.Currency{Currency: "GNF", DisplayName: "ગિનીયન ફ્રેંક", Symbol: "GNF"}, "NZD": ut.Currency{Currency: "NZD", DisplayName: "ન્યૂઝિલેંડ ડૉલર", Symbol: "NZ$"}, "RWF": ut.Currency{Currency: "RWF", DisplayName: "રવાંડન ફ્રેંક", Symbol: "RWF"}, "SAR": ut.Currency{Currency: "SAR", DisplayName: "સાઉદી રિયાલ", Symbol: "SAR"}, "USD": ut.Currency{Currency: "USD", DisplayName: "યુઍસ ડોલર", Symbol: "US$"}, "BAM": ut.Currency{Currency: "BAM", DisplayName: "બોસ્નિયા અને હર્ઝેગોવિના રૂપાંતર યોગ્ય માર્ક", Symbol: "BAM"}, "BYR": ut.Currency{Currency: "BYR", DisplayName: "બેલારુશિયન રૂબલ", Symbol: "BYR"}, "PHP": ut.Currency{Currency: "PHP", DisplayName: "ફિલીપાઇન પેસો", Symbol: "PHP"}, "GYD": ut.Currency{Currency: "GYD", DisplayName: "ગયાનિઝ ડોલર", Symbol: "GYD"}, "LYD": ut.Currency{Currency: "LYD", DisplayName: "લિબ્યન દિનાર", Symbol: "LYD"}, "FJD": ut.Currency{Currency: "FJD", DisplayName: "ફિજિઅન ડોલર", Symbol: "FJD"}, "XAF": ut.Currency{Currency: "XAF", DisplayName: "મધ્ય આફ્રિકન [CFA] ફ્રેંક", Symbol: "FCFA"}, "MOP": ut.Currency{Currency: "MOP", DisplayName: "માકાનિઝ પતાકા", Symbol: "MOP"}, "NOK": ut.Currency{Currency: "NOK", DisplayName: "નૉર્વેજિયન ક્રોન", Symbol: "NOK"}, "ZAR": ut.Currency{Currency: "ZAR", DisplayName: "દક્ષિણ આફ્રિકી રેંડ", Symbol: "ZAR"}, "ZMK": ut.Currency{Currency: "ZMK", DisplayName: "ઝામ્બિયન ક્વાચા (1968–2012)", Symbol: ""}, "BIF": ut.Currency{Currency: "BIF", DisplayName: "બુરુન્ડિયન ફ્રેંક", Symbol: "BIF"}, "CUP": ut.Currency{Currency: "CUP", DisplayName: "ક્યુબન પેસો", Symbol: "CUP"}, "CAD": ut.Currency{Currency: "CAD", DisplayName: "કેનેડિયન ડૉલર", Symbol: "CA$"}, "HNL": ut.Currency{Currency: "HNL", DisplayName: "હોન્ડ્યુરન લેમ્પિરા", Symbol: "HNL"}, "SRD": ut.Currency{Currency: "SRD", DisplayName: "સૂરીનામિઝ ડોલર", Symbol: "SRD"}, "SSP": ut.Currency{Currency: "SSP", DisplayName: "દક્ષિણ સુદાનિઝ પાઉન્ડ", Symbol: "SSP"}, "CNY": ut.Currency{Currency: "CNY", DisplayName: "ચાઇનિઝ યુઆન", Symbol: "CN¥"}, "RSD": ut.Currency{Currency: "RSD", DisplayName: "સર્બિયન દિનાર", Symbol: "RSD"}, "KPW": ut.Currency{Currency: "KPW", DisplayName: "ઉત્તર કોરિયન વન", Symbol: "KPW"}, "MMK": ut.Currency{Currency: "MMK", DisplayName: "મ્યાંમાર ક્યાત", Symbol: "MMK"}, "AZN": ut.Currency{Currency: "AZN", DisplayName: "અઝરબૈજાની મનાત", Symbol: "AZN"}, "HKD": ut.Currency{Currency: "HKD", DisplayName: "હોંગ કોંગ ડૉલર", Symbol: "HK$"}, "KRW": ut.Currency{Currency: "KRW", DisplayName: "દક્ષિણ કોરિયન વન", Symbol: "₩"}, "MNT": ut.Currency{Currency: "MNT", DisplayName: "મોંગોલિયન ટગરિક", Symbol: "MNT"}, "STD": ut.Currency{Currency: "STD", DisplayName: "સાઓ ટૉમ એન્ડ પ્રિંસાઇપ ડોબ્રા", Symbol: "STD"}, "VND": ut.Currency{Currency: "VND", DisplayName: "વિયેતનામીસ ડોંગ", Symbol: "₫"}, "WST": ut.Currency{Currency: "WST", DisplayName: "સમોઅન તાલા", Symbol: "WST"}, "BZD": ut.Currency{Currency: "BZD", DisplayName: "બેલિઝ ડોલર", Symbol: "BZD"}, "CUC": ut.Currency{Currency: "CUC", DisplayName: "ક્યુબન રૂપાંતર યોગ્ય પેસો", Symbol: "CUC"}, "BRL": ut.Currency{Currency: "BRL", DisplayName: "બ્રાઝિલીયન રિઆલ", Symbol: "R$"}, "BWP": ut.Currency{Currency: "BWP", DisplayName: "બોત્સવાનન પુલા", Symbol: "BWP"}, "CDF": ut.Currency{Currency: "CDF", DisplayName: "કોંગોલિઝ ફ્રેંક", Symbol: "CDF"}, "IRR": ut.Currency{Currency: "IRR", DisplayName: "ઇરાનિયન રિયાલ", Symbol: "IRR"}, "ANG": ut.Currency{Currency: "ANG", DisplayName: "નેધરલેંડ એંટિલિન ગિલ્ડર", Symbol: "ANG"}, "BHD": ut.Currency{Currency: "BHD", DisplayName: "બેહરિની દિનાર", Symbol: "BHD"}, "GEL": ut.Currency{Currency: "GEL", DisplayName: "જ્યોર્જિઅન લારી", Symbol: "GEL"}, "MAD": ut.Currency{Currency: "MAD", DisplayName: "મોરોક્કન દિરહામ", Symbol: "MAD"}, "MKD": ut.Currency{Currency: "MKD", DisplayName: "મેસેડોનિયન દિનાર", Symbol: "MKD"}, "MRO": ut.Currency{Currency: "MRO", DisplayName: "મોરીશેનિયન ઓગુયા", Symbol: "MRO"}, "SDG": ut.Currency{Currency: "SDG", DisplayName: "સુદાનિઝ પાઉન્ડ", Symbol: "SDG"}, "SZL": ut.Currency{Currency: "SZL", DisplayName: "સ્વાઝી લિલાન્ગેની", Symbol: "SZL"}, "COP": ut.Currency{Currency: "COP", DisplayName: "કોલમ્બિયન પેસો", Symbol: "COP"}, "ETB": ut.Currency{Currency: "ETB", DisplayName: "ઇથિયોપીયન બિર", Symbol: "ETB"}, "ERN": ut.Currency{Currency: "ERN", DisplayName: "એરિટ્રેયન નક્ફા", Symbol: "ERN"}, "PYG": ut.Currency{Currency: "PYG", DisplayName: "પરાગ્વેયન ગુઆરાની", Symbol: "PYG"}, "TND": ut.Currency{Currency: "TND", DisplayName: "ટ્યુનિશિયન દિનાર", Symbol: "TND"}, "TZS": ut.Currency{Currency: "TZS", DisplayName: "તાન્ઝાનિયન શિલિંગ", Symbol: "TZS"}, "UAH": ut.Currency{Currency: "UAH", DisplayName: "યુક્રેનિયન હ્રિવિનિયા", Symbol: "UAH"}, "VEF": ut.Currency{Currency: "VEF", DisplayName: "વેનેઝુએલન બોલિવર", Symbol: "VEF"}, "BGN": ut.Currency{Currency: "BGN", DisplayName: "બલ્ગેરીયન લેવ", Symbol: "BGN"}, "EGP": ut.Currency{Currency: "EGP", DisplayName: "ઇજિપ્તિયન પાઉન્ડ", Symbol: "EGP"}, "HRK": ut.Currency{Currency: "HRK", DisplayName: "ક્રોએશિયન ક્યુના", Symbol: "HRK"}, "ISK": ut.Currency{Currency: "ISK", DisplayName: "આઇસલેન્ડિક ક્રોના", Symbol: "ISK"}, "OMR": ut.Currency{Currency: "OMR", DisplayName: "ઓમાની રિયાલ", Symbol: "OMR"}, "PAB": ut.Currency{Currency: "PAB", DisplayName: "પનામેનિયન બાલ્બોઆ", Symbol: "PAB"}, "TMT": ut.Currency{Currency: "TMT", DisplayName: "તુર્કમેનિસ્તાની મનત", Symbol: "TMT"}, "YER": ut.Currency{Currency: "YER", DisplayName: "યેમેની રિઆલ", Symbol: "YER"}, "CRC": ut.Currency{Currency: "CRC", DisplayName: "કોસ્ટા રિકન કોલોન", Symbol: "CRC"}, "CVE": ut.Currency{Currency: "CVE", DisplayName: "કેપ વર્દિયન એસ્કુડો", Symbol: "CVE"}, "SLL": ut.Currency{Currency: "SLL", DisplayName: "સિએરા લિઓનિઅન લિઓન", Symbol: "SLL"}, "SOS": ut.Currency{Currency: "SOS", DisplayName: "સોમાલી શિલિંગ", Symbol: "SOS"}, "LBP": ut.Currency{Currency: "LBP", DisplayName: "લેબેનિઝ પાઉન્ડ", Symbol: "LBP"}, "SEK": ut.Currency{Currency: "SEK", DisplayName: "સ્વીડિશ ક્રોના", Symbol: "SEK"}, "FKP": ut.Currency{Currency: "FKP", DisplayName: "ફૉકલેન્ડ આઇલેંડ્સ પાઉન્ડ", Symbol: "FKP"}, "PKR": ut.Currency{Currency: "PKR", DisplayName: "પાકિસ્તાની રૂપી", Symbol: "PKR"}, "PLN": ut.Currency{Currency: "PLN", DisplayName: "પોલિસ ઝ્લોટી", Symbol: "PLN"}, "TRY": ut.Currency{Currency: "TRY", DisplayName: "તુર્કિશ લિરા", Symbol: "TRY"}, "BSD": ut.Currency{Currency: "BSD", DisplayName: "બહામિયન ડોલર", Symbol: "BSD"}, "MUR": ut.Currency{Currency: "MUR", DisplayName: "મોરેશીઅન રૂપી", Symbol: "MUR"}, "GTQ": ut.Currency{Currency: "GTQ", DisplayName: "ગ્વાટેમાલા કુઇટ્ઝલ", Symbol: "GTQ"}, "IQD": ut.Currency{Currency: "IQD", DisplayName: "ઇરાકી દિનાર", Symbol: "IQD"}, "KES": ut.Currency{Currency: "KES", DisplayName: "કેન્યેન શિલિંગ", Symbol: "KES"}, "KWD": ut.Currency{Currency: "KWD", DisplayName: "કુવૈતી દિનાર", Symbol: "KWD"}, "KZT": ut.Currency{Currency: "KZT", DisplayName: "કઝાકિસ્તાની ટેંગ", Symbol: "KZT"}, "RUB": ut.Currency{Currency: "RUB", DisplayName: "રશિયન રૂબલ", Symbol: "RUB"}, "ALL": ut.Currency{Currency: "ALL", DisplayName: "અલ્બેનિયન લેક", Symbol: "ALL"}, "BDT": ut.Currency{Currency: "BDT", DisplayName: "બાંગલાદેશી ટાકા", Symbol: "BDT"}, "XXX": ut.Currency{Currency: "XXX", DisplayName: "અજ્ઞાત ચલણ", Symbol: ""}, "SGD": ut.Currency{Currency: "SGD", DisplayName: "સિંગાપુર ડૉલર", Symbol: "SGD"}, "UZS": ut.Currency{Currency: "UZS", DisplayName: "ઉઝ્બેકિસ્તાન સોમ", Symbol: "UZS"}, "CZK": ut.Currency{Currency: "CZK", DisplayName: "ચેક રીપબ્લિક કોરુના", Symbol: "CZK"}, "DZD": ut.Currency{Currency: "DZD", DisplayName: "અલ્જિરિયન દિનાર", Symbol: "DZD"}, "LVL": ut.Currency{Currency: "LVL", DisplayName: "લાતવિયન લેત્સ", Symbol: "LVL"}, "ZMW": ut.Currency{Currency: "ZMW", DisplayName: "ઝામ્બિયન ક્વાચા", Symbol: "ZMW"}, "AUD": ut.Currency{Currency: "AUD", DisplayName: "ઑસ્ટ્રેલિયન ડૉલર", Symbol: "A$"}, "BTN": ut.Currency{Currency: "BTN", DisplayName: "ભુતાનિઝ એંગુલ્ત્રમ", Symbol: "BTN"}, "SBD": ut.Currency{Currency: "SBD", DisplayName: "સોલોમન આઇલેંડ્સ ડોલર", Symbol: "SBD"}, "DOP": ut.Currency{Currency: "DOP", DisplayName: "ડોમિનિકન પેસો", Symbol: "DOP"}, "ILS": ut.Currency{Currency: "ILS", DisplayName: "ઇઝરાયેલી ન્યુ શેકલ", Symbol: "₪"}, "GMD": ut.Currency{Currency: "GMD", DisplayName: "ગેમ્બિયન દલાસી", Symbol: "GMD"}, "MDL": ut.Currency{Currency: "MDL", DisplayName: "મોલડોવેન લિયુ", Symbol: "MDL"}, "NPR": ut.Currency{Currency: "NPR", DisplayName: "નેપાલિઝ રૂપી", Symbol: "NPR"}, "SYP": ut.Currency{Currency: "SYP", DisplayName: "સાઇરિયન પાઉન્ડ", Symbol: "SYP"}, "TWD": ut.Currency{Currency: "TWD", DisplayName: "ન્યુ તાઇવાન ડૉલર", Symbol: "NT$"}, "XPF": ut.Currency{Currency: "XPF", DisplayName: "[CFP] ફ્રેંક", Symbol: "CFPF"}, "GHS": ut.Currency{Currency: "GHS", DisplayName: "ઘાનાઇયન સેડી", Symbol: "GHS"}, "GIP": ut.Currency{Currency: "GIP", DisplayName: "જીબ્રાલ્ટર પાઉન્ડ", Symbol: "GIP"}, "HUF": ut.Currency{Currency: "HUF", DisplayName: "હંગેરીયન ફોરિન્ત", Symbol: "HUF"}, "XCD": ut.Currency{Currency: "XCD", DisplayName: "ઇસ્ટ કેરિબિયન ડોલર", Symbol: "EC$"}, "BND": ut.Currency{Currency: "BND", DisplayName: "બ્રુનેઇ ડોલર", Symbol: "BND"}, "CHF": ut.Currency{Currency: "CHF", DisplayName: "સ્વિસ ફ્રેંક", Symbol: "CHF"}, "MWK": ut.Currency{Currency: "MWK", DisplayName: "માલાવિયન ક્વાચા", Symbol: "MWK"}, "LRD": ut.Currency{Currency: "LRD", DisplayName: "લિબેરિયન ડોલર", Symbol: "LRD"}, "MGA": ut.Currency{Currency: "MGA", DisplayName: "માલાગેસી અરીઆરી", Symbol: "MGA"}, "INR": ut.Currency{Currency: "INR", DisplayName: "ભારતીય રૂપિયા", Symbol: "₹"}, "JOD": ut.Currency{Currency: "JOD", DisplayName: "જોર્ડનિયન દિનાર", Symbol: "JOD"}, "SCR": ut.Currency{Currency: "SCR", DisplayName: "સેશેલોઈ રૂપી", Symbol: "SCR"}, "TOP": ut.Currency{Currency: "TOP", DisplayName: "ટોંગન પ’અંગા", Symbol: "TOP"}, "UYU": ut.Currency{Currency: "UYU", DisplayName: "ઉરુગ્વેયન પેસો", Symbol: "UYU"}, "XOF": ut.Currency{Currency: "XOF", DisplayName: "પશ્ચિમી આફ્રિકન [CFA] ફ્રેંક", Symbol: "CFA"}, "AMD": ut.Currency{Currency: "AMD", DisplayName: "અર્મેનિયન ડ્રેમ", Symbol: "AMD"}, "CLP": ut.Currency{Currency: "CLP", DisplayName: "ચિલિઅન પેસો", Symbol: "CLP"}, "LSL": ut.Currency{Currency: "LSL", DisplayName: "લેસોથો લોતી", Symbol: ""}, "PGK": ut.Currency{Currency: "PGK", DisplayName: "પાપુઆ ન્યૂ ગિનીયન કિના", Symbol: "PGK"}, "QAR": ut.Currency{Currency: "QAR", DisplayName: "કતારી રિયાલ", Symbol: "QAR"}, "DJF": ut.Currency{Currency: "DJF", DisplayName: "જિબુટિયન ફ્રેંક", Symbol: "DJF"}, "IDR": ut.Currency{Currency: "IDR", DisplayName: "ઇન્ડોનેશિયન રૂપિયા", Symbol: "IDR"}, "BOB": ut.Currency{Currency: "BOB", DisplayName: "બોલિવિયન બોલિવિયાનો", Symbol: "BOB"}, "GBP": ut.Currency{Currency: "GBP", DisplayName: "બ્રિટિશ પાઉન્ડ", Symbol: "£"}, "JMD": ut.Currency{Currency: "JMD", DisplayName: "જમૈકિયન ડોલર", Symbol: "JMD"}, "LAK": ut.Currency{Currency: "LAK", DisplayName: "લાઓશિયન કિપ", Symbol: "LAK"}, "MZN": ut.Currency{Currency: "MZN", DisplayName: "મોઝામ્બિકન મેટિકલ", Symbol: "MZN"}, "NIO": ut.Currency{Currency: "NIO", DisplayName: "નિકારાગુઅન કોર્ડોબા", Symbol: "NIO"}, "AOA": ut.Currency{Currency: "AOA", DisplayName: "અંગોલિયન ક્વાન્ઝા", Symbol: "AOA"}, "BBD": ut.Currency{Currency: "BBD", DisplayName: "બાર્બાડિયન ડોલર", Symbol: "BBD"}}
