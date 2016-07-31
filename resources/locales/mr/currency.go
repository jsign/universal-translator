package mr

import "github.com/go-playground/universal-translator"

var currencies = ut.CurrencyFormatValue{"MDL": ut.Currency{Currency: "MDL", DisplayName: "मोल्डोवन लेउ", Symbol: "MDL"}, "PAB": ut.Currency{Currency: "PAB", DisplayName: "पनामा बाल्बोआ", Symbol: "PAB"}, "SBD": ut.Currency{Currency: "SBD", DisplayName: "सोलोमन आयलँड्स डॉलर", Symbol: "SBD"}, "ETB": ut.Currency{Currency: "ETB", DisplayName: "इथिओपियन बिर", Symbol: "ETB"}, "FKP": ut.Currency{Currency: "FKP", DisplayName: "फॉकलंड आयलंड पाउंड", Symbol: "FKP"}, "NZD": ut.Currency{Currency: "NZD", DisplayName: "न्यूझीलँड डॉलर", Symbol: "NZ$"}, "PKR": ut.Currency{Currency: "PKR", DisplayName: "पाकिस्तानी रुपया", Symbol: "PKR"}, "AMD": ut.Currency{Currency: "AMD", DisplayName: "आर्मेनियन द्रॅम", Symbol: "AMD"}, "MGA": ut.Currency{Currency: "MGA", DisplayName: "मालागासी एरियारी", Symbol: "MGA"}, "OMR": ut.Currency{Currency: "OMR", DisplayName: "ओमानी रियाल", Symbol: "OMR"}, "BHD": ut.Currency{Currency: "BHD", DisplayName: "बाहरिनी दिनार", Symbol: "BHD"}, "CRC": ut.Currency{Currency: "CRC", DisplayName: "कोस्टा रिका कोलोन", Symbol: "CRC"}, "MMK": ut.Currency{Currency: "MMK", DisplayName: "म्यानमार क्याट", Symbol: "MMK"}, "PYG": ut.Currency{Currency: "PYG", DisplayName: "पराग्वे ग्वारानी", Symbol: "PYG"}, "TND": ut.Currency{Currency: "TND", DisplayName: "ट्यूनिशियन दिनार", Symbol: "TND"}, "JOD": ut.Currency{Currency: "JOD", DisplayName: "जॉर्डनियन दिनार", Symbol: "JOD"}, "XCD": ut.Currency{Currency: "XCD", DisplayName: "पूर्व कॅरीबियन डॉलर", Symbol: "EC$"}, "XOF": ut.Currency{Currency: "XOF", DisplayName: "पश्चिम आफ्रिकन [CFA] फ्रँक", Symbol: "CFA"}, "ILS": ut.Currency{Currency: "ILS", DisplayName: "ईस्त्रायली न्यू शेकेल", Symbol: "₪"}, "KHR": ut.Currency{Currency: "KHR", DisplayName: "कंबोडियन रियेल", Symbol: "KHR"}, "KWD": ut.Currency{Currency: "KWD", DisplayName: "कुवैती दिनार", Symbol: "KWD"}, "AED": ut.Currency{Currency: "AED", DisplayName: "संयुक्त अरब अमीरात दिरहॅम", Symbol: "AED"}, "AUD": ut.Currency{Currency: "AUD", DisplayName: "ऑस्ट्रेलियन डॉलर", Symbol: "A$"}, "BTN": ut.Currency{Currency: "BTN", DisplayName: "भूतानी एंगल्ट्रम", Symbol: "BTN"}, "BZD": ut.Currency{Currency: "BZD", DisplayName: "बेलीझ डॉलर", Symbol: "BZD"}, "GHS": ut.Currency{Currency: "GHS", DisplayName: "घानीयन सेडी", Symbol: "GHS"}, "YER": ut.Currency{Currency: "YER", DisplayName: "येमेनी रियाल", Symbol: "YER"}, "MUR": ut.Currency{Currency: "MUR", DisplayName: "मॉरिशियन रुपी", Symbol: "MUR"}, "SEK": ut.Currency{Currency: "SEK", DisplayName: "स्वीडिश क्रोना", Symbol: "SEK"}, "TOP": ut.Currency{Currency: "TOP", DisplayName: "टोंगन पाआंगा", Symbol: "TOP"}, "UGX": ut.Currency{Currency: "UGX", DisplayName: "युगांडा शिलिंग", Symbol: "UGX"}, "UYU": ut.Currency{Currency: "UYU", DisplayName: "उरुग्वेचा पेसो", Symbol: "UYU"}, "BYR": ut.Currency{Currency: "BYR", DisplayName: "बेलारुशियन रुबल", Symbol: "BYR"}, "DJF": ut.Currency{Currency: "DJF", DisplayName: "जिबौटियन फ्रँक", Symbol: "DJF"}, "FJD": ut.Currency{Currency: "FJD", DisplayName: "फिजियन डॉलर", Symbol: "FJD"}, "KGS": ut.Currency{Currency: "KGS", DisplayName: "किरगिस्तानी सॉम", Symbol: "KGS"}, "LBP": ut.Currency{Currency: "LBP", DisplayName: "लेबनीज पाउंड", Symbol: "LBP"}, "SDG": ut.Currency{Currency: "SDG", DisplayName: "सुदानी पाउंड", Symbol: "SDG"}, "UZS": ut.Currency{Currency: "UZS", DisplayName: "उझबेकिस्तानी सोम", Symbol: "UZS"}, "AWG": ut.Currency{Currency: "AWG", DisplayName: "अरुबा फ्लोरिन", Symbol: "AWG"}, "BOB": ut.Currency{Currency: "BOB", DisplayName: "बोलिव्हियन बोलिव्हियानो", Symbol: "BOB"}, "DZD": ut.Currency{Currency: "DZD", DisplayName: "अल्जेरियन दिनार", Symbol: "DZD"}, "HKD": ut.Currency{Currency: "HKD", DisplayName: "हाँगकाँग डॉलर", Symbol: "HK$"}, "INR": ut.Currency{Currency: "INR", DisplayName: "भारतीय रुपया", Symbol: "₹"}, "TRY": ut.Currency{Currency: "TRY", DisplayName: "तुर्की लिरा", Symbol: "TRY"}, "VUV": ut.Currency{Currency: "VUV", DisplayName: "वानुआतु वाटु", Symbol: "VUV"}, "BDT": ut.Currency{Currency: "BDT", DisplayName: "बांगलादेशी टका", Symbol: "BDT"}, "HNL": ut.Currency{Currency: "HNL", DisplayName: "होन्डुरन लेंपिरा", Symbol: "HNL"}, "KPW": ut.Currency{Currency: "KPW", DisplayName: "उत्तर कोरियन वॉन", Symbol: "KPW"}, "MNT": ut.Currency{Currency: "MNT", DisplayName: "मंगोलियन टुग्रिक", Symbol: "MNT"}, "NGN": ut.Currency{Currency: "NGN", DisplayName: "नायजेरियन नायरा", Symbol: "NGN"}, "DKK": ut.Currency{Currency: "DKK", DisplayName: "डॅनिश क्रोन", Symbol: "DKK"}, "KRW": ut.Currency{Currency: "KRW", DisplayName: "दक्षिण कोरियन वॉन", Symbol: "₩"}, "QAR": ut.Currency{Currency: "QAR", DisplayName: "कतारी रियाल", Symbol: "QAR"}, "SGD": ut.Currency{Currency: "SGD", DisplayName: "सिंगापूर डॉलर", Symbol: "SGD"}, "TMT": ut.Currency{Currency: "TMT", DisplayName: "तुर्कमेनिस्तानी मानाट", Symbol: "TMT"}, "RWF": ut.Currency{Currency: "RWF", DisplayName: "रवांडा फ्रँक", Symbol: "RWF"}, "BRL": ut.Currency{Currency: "BRL", DisplayName: "ब्राझिलियन रियाल", Symbol: "R$"}, "CAD": ut.Currency{Currency: "CAD", DisplayName: "कॅनडियन डॉलर", Symbol: "CA$"}, "COP": ut.Currency{Currency: "COP", DisplayName: "कोलंबियन पेसो", Symbol: "COP"}, "IRR": ut.Currency{Currency: "IRR", DisplayName: "इराणी रियाल", Symbol: "IRR"}, "KYD": ut.Currency{Currency: "KYD", DisplayName: "केमेन आयलॅंड डॉलर", Symbol: "KYD"}, "BAM": ut.Currency{Currency: "BAM", DisplayName: "बोस्निया-हर्जेगोविना विनिमय मार्क", Symbol: "BAM"}, "MVR": ut.Currency{Currency: "MVR", DisplayName: "मालदीवियन रुफिया", Symbol: "MVR"}, "RSD": ut.Currency{Currency: "RSD", DisplayName: "सर्बियन दिनार", Symbol: "RSD"}, "SYP": ut.Currency{Currency: "SYP", DisplayName: "सीरियन पाउंड", Symbol: "SYP"}, "CDF": ut.Currency{Currency: "CDF", DisplayName: "काँगोलीज फ्रँक", Symbol: "CDF"}, "KZT": ut.Currency{Currency: "KZT", DisplayName: "कझाकिस्तानी तेंगे", Symbol: "KZT"}, "LTL": ut.Currency{Currency: "LTL", DisplayName: "लिथुआनियन लिटास", Symbol: "LTL"}, "MOP": ut.Currency{Currency: "MOP", DisplayName: "मॅकॅनीज् पटाका", Symbol: "MOP"}, "STD": ut.Currency{Currency: "STD", DisplayName: "साओ टोम आणि प्रिन्सिपे डोबरा", Symbol: "STD"}, "ZMW": ut.Currency{Currency: "ZMW", DisplayName: "झांबियन क्वाचा", Symbol: "ZMW"}, "ALL": ut.Currency{Currency: "ALL", DisplayName: "अल्बानियन लेक", Symbol: "ALL"}, "AOA": ut.Currency{Currency: "AOA", DisplayName: "अँगोलन क्वॅन्झा", Symbol: "AOA"}, "NOK": ut.Currency{Currency: "NOK", DisplayName: "नॉर्वेजियन क्रोन", Symbol: "NOK"}, "THB": ut.Currency{Currency: "THB", DisplayName: "थाई बाहत", Symbol: "฿"}, "TWD": ut.Currency{Currency: "TWD", DisplayName: "न्यू तैवान डॉलर", Symbol: "NT$"}, "BWP": ut.Currency{Currency: "BWP", DisplayName: "बोट्सवानन पुला", Symbol: "BWP"}, "CVE": ut.Currency{Currency: "CVE", DisplayName: "केप व्हर्डेयन एस्कुडो", Symbol: "CVE"}, "SAR": ut.Currency{Currency: "SAR", DisplayName: "सौदी रियाल", Symbol: "SAR"}, "SOS": ut.Currency{Currency: "SOS", DisplayName: "सोमाली शिलिंग", Symbol: "SOS"}, "GMD": ut.Currency{Currency: "GMD", DisplayName: "गाम्बियन डालासी", Symbol: "GMD"}, "LVL": ut.Currency{Currency: "LVL", DisplayName: "लाट्व्हियन लाट्झ", Symbol: "LVL"}, "MRO": ut.Currency{Currency: "MRO", DisplayName: "मॉरिटानियन ओगिया", Symbol: "MRO"}, "NAD": ut.Currency{Currency: "NAD", DisplayName: "नमिबियन डॉलर", Symbol: "NAD"}, "PEN": ut.Currency{Currency: "PEN", DisplayName: "पेरुवियन नुइव्हो सोल", Symbol: "PEN"}, "CLP": ut.Currency{Currency: "CLP", DisplayName: "चिली पेसो", Symbol: "CLP"}, "IDR": ut.Currency{Currency: "IDR", DisplayName: "इंडोनेशियन रुपिया", Symbol: "IDR"}, "KES": ut.Currency{Currency: "KES", DisplayName: "केनियन शिलिंग", Symbol: "KES"}, "SSP": ut.Currency{Currency: "SSP", DisplayName: "दक्षिण सुदानी पाउंड", Symbol: "SSP"}, "BIF": ut.Currency{Currency: "BIF", DisplayName: "बुरुंडियन फ्रँक", Symbol: "BIF"}, "BMD": ut.Currency{Currency: "BMD", DisplayName: "बर्मुडा डॉलर", Symbol: "BMD"}, "GBP": ut.Currency{Currency: "GBP", DisplayName: "ब्रिटिश पाऊंड", Symbol: "£"}, "JPY": ut.Currency{Currency: "JPY", DisplayName: "जपानी येन", Symbol: "JP¥"}, "LAK": ut.Currency{Currency: "LAK", DisplayName: "लाओशियन किप", Symbol: "LAK"}, "MKD": ut.Currency{Currency: "MKD", DisplayName: "मॅसेडोनियन देनार", Symbol: "MKD"}, "VND": ut.Currency{Currency: "VND", DisplayName: "व्हिएतनामी डोंग", Symbol: "₫"}, "CUC": ut.Currency{Currency: "CUC", DisplayName: "क्यूबन विनिमय पेसो", Symbol: "CUC"}, "CUP": ut.Currency{Currency: "CUP", DisplayName: "क्यूबन पेसो", Symbol: "CUP"}, "ERN": ut.Currency{Currency: "ERN", DisplayName: "इरिट्रियन नाक्फा", Symbol: "ERN"}, "IQD": ut.Currency{Currency: "IQD", DisplayName: "इराकी दिनार", Symbol: "IQD"}, "LKR": ut.Currency{Currency: "LKR", DisplayName: "श्रीलंकन रुपया", Symbol: "LKR"}, "AZN": ut.Currency{Currency: "AZN", DisplayName: "अझरबैझानी मानाट", Symbol: "AZN"}, "EUR": ut.Currency{Currency: "EUR", DisplayName: "युरो", Symbol: "€"}, "XAF": ut.Currency{Currency: "XAF", DisplayName: "मध्य आफ्रिकन [CFA] फ्रँक", Symbol: "FCFA"}, "ZAR": ut.Currency{Currency: "ZAR", DisplayName: "दक्षिण आफ्रिकी रँड", Symbol: "ZAR"}, "ZMK": ut.Currency{Currency: "ZMK", DisplayName: "झांबियन क्वाचा (1968–2012)", Symbol: ""}, "PLN": ut.Currency{Currency: "PLN", DisplayName: "पोलिश झ्लॉटी", Symbol: "PLN"}, "TZS": ut.Currency{Currency: "TZS", DisplayName: "टांझानियन शिलिंग", Symbol: "TZS"}, "XPF": ut.Currency{Currency: "XPF", DisplayName: "[CFP] फ्रँक", Symbol: "CFPF"}, "EGP": ut.Currency{Currency: "EGP", DisplayName: "इजिप्शियन पाउंड", Symbol: "EGP"}, "GTQ": ut.Currency{Currency: "GTQ", DisplayName: "ग्वाटेमालाचे क्वेत्झाल", Symbol: "GTQ"}, "LRD": ut.Currency{Currency: "LRD", DisplayName: "लाइबेरियन डॉलर", Symbol: "LRD"}, "LSL": ut.Currency{Currency: "LSL", DisplayName: "लेसोटो लोटी", Symbol: ""}, "NIO": ut.Currency{Currency: "NIO", DisplayName: "निकाराग्वेचा कोर्डोबा", Symbol: "NIO"}, "AFN": ut.Currency{Currency: "AFN", DisplayName: "अफगाण अफगाणी", Symbol: "AFN"}, "BBD": ut.Currency{Currency: "BBD", DisplayName: "बार्बाडियन डॉलर", Symbol: "BBD"}, "CHF": ut.Currency{Currency: "CHF", DisplayName: "स्विस फ्रँक", Symbol: "CHF"}, "HTG": ut.Currency{Currency: "HTG", DisplayName: "हैती गोअर्ड", Symbol: "HTG"}, "MZN": ut.Currency{Currency: "MZN", DisplayName: "मोझांबिकन मेटिकल", Symbol: "MZN"}, "HRK": ut.Currency{Currency: "HRK", DisplayName: "क्रोएशियन कूना", Symbol: "HRK"}, "BND": ut.Currency{Currency: "BND", DisplayName: "ब्रुनेई डॉलर", Symbol: "BND"}, "BSD": ut.Currency{Currency: "BSD", DisplayName: "बहामी डॉलर", Symbol: "BSD"}, "CZK": ut.Currency{Currency: "CZK", DisplayName: "झेक प्रजासत्ताक कोरुना", Symbol: "CZK"}, "GEL": ut.Currency{Currency: "GEL", DisplayName: "जॉर्जियन लारी", Symbol: "GEL"}, "GIP": ut.Currency{Currency: "GIP", DisplayName: "जिब्राल्टर पाउंड", Symbol: "GIP"}, "RON": ut.Currency{Currency: "RON", DisplayName: "रोमानियन लेऊ", Symbol: "RON"}, "RUB": ut.Currency{Currency: "RUB", DisplayName: "रशियन रुबल", Symbol: "RUB"}, "ANG": ut.Currency{Currency: "ANG", DisplayName: "नेदरलँडचा अँटिलीन गिल्डर", Symbol: "ANG"}, "ISK": ut.Currency{Currency: "ISK", DisplayName: "आइसलँडिक क्रोना", Symbol: "ISK"}, "JMD": ut.Currency{Currency: "JMD", DisplayName: "जमैकन डॉलर", Symbol: "JMD"}, "MXN": ut.Currency{Currency: "MXN", DisplayName: "मेक्सिको पेसो", Symbol: "MX$"}, "MYR": ut.Currency{Currency: "MYR", DisplayName: "मलेशियन रिंगिट", Symbol: "MYR"}, "DOP": ut.Currency{Currency: "DOP", DisplayName: "डोमिनिकन पेसो", Symbol: "DOP"}, "MWK": ut.Currency{Currency: "MWK", DisplayName: "मालावियन क्वाचा", Symbol: "MWK"}, "NPR": ut.Currency{Currency: "NPR", DisplayName: "नेपाळी रुपया", Symbol: "NPR"}, "PHP": ut.Currency{Currency: "PHP", DisplayName: "फिलिपिनी पेसो", Symbol: "PHP"}, "KMF": ut.Currency{Currency: "KMF", DisplayName: "कोमोरियन फ्रँक", Symbol: "KMF"}, "LYD": ut.Currency{Currency: "LYD", DisplayName: "लिबियाचा दिनार", Symbol: "LYD"}, "TJS": ut.Currency{Currency: "TJS", DisplayName: "ताजकीस्तानी सोमोनी", Symbol: "TJS"}, "UAH": ut.Currency{Currency: "UAH", DisplayName: "युक्रेनियन रिवनिया", Symbol: "UAH"}, "XXX": ut.Currency{Currency: "XXX", DisplayName: "अज्ञात चलन", Symbol: ""}, "GNF": ut.Currency{Currency: "GNF", DisplayName: "गिनी फ्रँक", Symbol: "GNF"}, "GYD": ut.Currency{Currency: "GYD", DisplayName: "गयाना डॉलर", Symbol: "GYD"}, "MAD": ut.Currency{Currency: "MAD", DisplayName: "मोरोक्को दिरहॅम", Symbol: "MAD"}, "USD": ut.Currency{Currency: "USD", DisplayName: "यूएस डॉलर", Symbol: "$"}, "WST": ut.Currency{Currency: "WST", DisplayName: "सामोअन टाला", Symbol: "WST"}, "SCR": ut.Currency{Currency: "SCR", DisplayName: "सेशेलोईस रुपी", Symbol: "SCR"}, "SHP": ut.Currency{Currency: "SHP", DisplayName: "सेंट हेलेना पाउंड", Symbol: "SHP"}, "SLL": ut.Currency{Currency: "SLL", DisplayName: "सिएरा लिऑनचा लिऑन", Symbol: "SLL"}, "BGN": ut.Currency{Currency: "BGN", DisplayName: "बल्गेरियन लेव", Symbol: "BGN"}, "CNY": ut.Currency{Currency: "CNY", DisplayName: "चीनी युआन", Symbol: "CN¥"}, "PGK": ut.Currency{Currency: "PGK", DisplayName: "पापुआ न्यू गिनीयन किना", Symbol: "PGK"}, "SZL": ut.Currency{Currency: "SZL", DisplayName: "स्वाझी लीलांगेनी", Symbol: "SZL"}, "TTD": ut.Currency{Currency: "TTD", DisplayName: "त्रिनिदाद आणि टोबॅगो डॉलर", Symbol: "TTD"}, "ARS": ut.Currency{Currency: "ARS", DisplayName: "अर्जेंटाइन पेसो", Symbol: "ARS"}, "HUF": ut.Currency{Currency: "HUF", DisplayName: "हंगेरियन फॉरिन्ट", Symbol: "HUF"}, "SRD": ut.Currency{Currency: "SRD", DisplayName: "सुरिनामी डॉलर", Symbol: "SRD"}, "VEF": ut.Currency{Currency: "VEF", DisplayName: "व्हेनेझुएला बोलिव्हार", Symbol: "VEF"}}