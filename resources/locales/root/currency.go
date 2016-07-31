package root

import "github.com/go-playground/universal-translator"

var currencies = ut.CurrencyFormatValue{"PYG": ut.Currency{Currency: "PYG", DisplayName: "", Symbol: "₲"}, "SBD": ut.Currency{Currency: "SBD", DisplayName: "", Symbol: "$"}, "BDT": ut.Currency{Currency: "BDT", DisplayName: "", Symbol: "৳"}, "BSD": ut.Currency{Currency: "BSD", DisplayName: "", Symbol: "$"}, "COP": ut.Currency{Currency: "COP", DisplayName: "", Symbol: "$"}, "GBP": ut.Currency{Currency: "GBP", DisplayName: "", Symbol: "£"}, "MYR": ut.Currency{Currency: "MYR", DisplayName: "", Symbol: "RM"}, "NAD": ut.Currency{Currency: "NAD", DisplayName: "", Symbol: "$"}, "SEK": ut.Currency{Currency: "SEK", DisplayName: "", Symbol: "kr"}, "ILS": ut.Currency{Currency: "ILS", DisplayName: "", Symbol: "₪"}, "MGA": ut.Currency{Currency: "MGA", DisplayName: "", Symbol: "Ar"}, "MNT": ut.Currency{Currency: "MNT", DisplayName: "", Symbol: "₮"}, "NIO": ut.Currency{Currency: "NIO", DisplayName: "", Symbol: "C$"}, "XAF": ut.Currency{Currency: "XAF", DisplayName: "", Symbol: "FCFA"}, "TWD": ut.Currency{Currency: "TWD", DisplayName: "", Symbol: "NT$"}, "XCD": ut.Currency{Currency: "XCD", DisplayName: "", Symbol: "EC$"}, "CUC": ut.Currency{Currency: "CUC", DisplayName: "", Symbol: "$"}, "DKK": ut.Currency{Currency: "DKK", DisplayName: "", Symbol: "kr"}, "GIP": ut.Currency{Currency: "GIP", DisplayName: "", Symbol: "£"}, "HRK": ut.Currency{Currency: "HRK", DisplayName: "", Symbol: "kn"}, "KZT": ut.Currency{Currency: "KZT", DisplayName: "", Symbol: "₸"}, "LVL": ut.Currency{Currency: "LVL", DisplayName: "", Symbol: "Ls"}, "XPF": ut.Currency{Currency: "XPF", DisplayName: "", Symbol: "CFPF"}, "GTQ": ut.Currency{Currency: "GTQ", DisplayName: "", Symbol: "Q"}, "JMD": ut.Currency{Currency: "JMD", DisplayName: "", Symbol: "$"}, "JPY": ut.Currency{Currency: "JPY", DisplayName: "", Symbol: "JP¥"}, "KRW": ut.Currency{Currency: "KRW", DisplayName: "", Symbol: "₩"}, "NZD": ut.Currency{Currency: "NZD", DisplayName: "", Symbol: "NZ$"}, "CRC": ut.Currency{Currency: "CRC", DisplayName: "", Symbol: "₡"}, "KHR": ut.Currency{Currency: "KHR", DisplayName: "", Symbol: "៛"}, "STD": ut.Currency{Currency: "STD", DisplayName: "", Symbol: "Db"}, "TRY": ut.Currency{Currency: "TRY", DisplayName: "", Symbol: "₺"}, "BAM": ut.Currency{Currency: "BAM", DisplayName: "", Symbol: "KM"}, "BRL": ut.Currency{Currency: "BRL", DisplayName: "", Symbol: "R$"}, "GYD": ut.Currency{Currency: "GYD", DisplayName: "", Symbol: "$"}, "NPR": ut.Currency{Currency: "NPR", DisplayName: "", Symbol: "Rs"}, "NGN": ut.Currency{Currency: "NGN", DisplayName: "", Symbol: "₦"}, "AOA": ut.Currency{Currency: "AOA", DisplayName: "", Symbol: "Kz"}, "BZD": ut.Currency{Currency: "BZD", DisplayName: "", Symbol: "$"}, "EGP": ut.Currency{Currency: "EGP", DisplayName: "", Symbol: "E£"}, "ESP": ut.Currency{Currency: "ESP", DisplayName: "", Symbol: "₧"}, "FJD": ut.Currency{Currency: "FJD", DisplayName: "", Symbol: "$"}, "FKP": ut.Currency{Currency: "FKP", DisplayName: "", Symbol: "£"}, "BWP": ut.Currency{Currency: "BWP", DisplayName: "", Symbol: "P"}, "RUR": ut.Currency{Currency: "RUR", DisplayName: "", Symbol: "р."}, "ZAR": ut.Currency{Currency: "ZAR", DisplayName: "", Symbol: "R"}, "ZMW": ut.Currency{Currency: "ZMW", DisplayName: "", Symbol: "ZK"}, "SRD": ut.Currency{Currency: "SRD", DisplayName: "", Symbol: "$"}, "THB": ut.Currency{Currency: "THB", DisplayName: "", Symbol: "฿"}, "ARS": ut.Currency{Currency: "ARS", DisplayName: "", Symbol: "$"}, "BND": ut.Currency{Currency: "BND", DisplayName: "", Symbol: "$"}, "GNF": ut.Currency{Currency: "GNF", DisplayName: "", Symbol: "FG"}, "LAK": ut.Currency{Currency: "LAK", DisplayName: "", Symbol: "₭"}, "LRD": ut.Currency{Currency: "LRD", DisplayName: "", Symbol: "$"}, "PLN": ut.Currency{Currency: "PLN", DisplayName: "", Symbol: "zł"}, "CAD": ut.Currency{Currency: "CAD", DisplayName: "", Symbol: "CA$"}, "GEL": ut.Currency{Currency: "GEL", DisplayName: "", Symbol: "₾"}, "CNY": ut.Currency{Currency: "CNY", DisplayName: "", Symbol: "CN¥"}, "CZK": ut.Currency{Currency: "CZK", DisplayName: "", Symbol: "Kč"}, "LBP": ut.Currency{Currency: "LBP", DisplayName: "", Symbol: "L£"}, "PKR": ut.Currency{Currency: "PKR", DisplayName: "", Symbol: "Rs"}, "SSP": ut.Currency{Currency: "SSP", DisplayName: "", Symbol: "£"}, "DOP": ut.Currency{Currency: "DOP", DisplayName: "", Symbol: "$"}, "HUF": ut.Currency{Currency: "HUF", DisplayName: "", Symbol: "Ft"}, "IDR": ut.Currency{Currency: "IDR", DisplayName: "", Symbol: "Rp"}, "RWF": ut.Currency{Currency: "RWF", DisplayName: "", Symbol: "RF"}, "SHP": ut.Currency{Currency: "SHP", DisplayName: "", Symbol: "£"}, "TTD": ut.Currency{Currency: "TTD", DisplayName: "", Symbol: "$"}, "XOF": ut.Currency{Currency: "XOF", DisplayName: "", Symbol: "CFA"}, "BBD": ut.Currency{Currency: "BBD", DisplayName: "", Symbol: "$"}, "BMD": ut.Currency{Currency: "BMD", DisplayName: "", Symbol: "$"}, "INR": ut.Currency{Currency: "INR", DisplayName: "", Symbol: "₹"}, "LKR": ut.Currency{Currency: "LKR", DisplayName: "", Symbol: "Rs"}, "MXN": ut.Currency{Currency: "MXN", DisplayName: "", Symbol: "MX$"}, "UAH": ut.Currency{Currency: "UAH", DisplayName: "", Symbol: "₴"}, "KYD": ut.Currency{Currency: "KYD", DisplayName: "", Symbol: "$"}, "LTL": ut.Currency{Currency: "LTL", DisplayName: "", Symbol: "Lt"}, "BOB": ut.Currency{Currency: "BOB", DisplayName: "", Symbol: "Bs"}, "CUP": ut.Currency{Currency: "CUP", DisplayName: "", Symbol: "$"}, "HKD": ut.Currency{Currency: "HKD", DisplayName: "", Symbol: "HK$"}, "HNL": ut.Currency{Currency: "HNL", DisplayName: "", Symbol: "L"}, "ISK": ut.Currency{Currency: "ISK", DisplayName: "", Symbol: "kr"}, "KPW": ut.Currency{Currency: "KPW", DisplayName: "", Symbol: "₩"}, "PHP": ut.Currency{Currency: "PHP", DisplayName: "", Symbol: "₱"}, "RUB": ut.Currency{Currency: "RUB", DisplayName: "", Symbol: "₽"}, "SYP": ut.Currency{Currency: "SYP", DisplayName: "", Symbol: "£"}, "USD": ut.Currency{Currency: "USD", DisplayName: "", Symbol: "US$"}, "VND": ut.Currency{Currency: "VND", DisplayName: "", Symbol: "₫"}, "BYR": ut.Currency{Currency: "BYR", DisplayName: "", Symbol: "р."}, "KMF": ut.Currency{Currency: "KMF", DisplayName: "", Symbol: "CF"}, "MMK": ut.Currency{Currency: "MMK", DisplayName: "", Symbol: "K"}, "NOK": ut.Currency{Currency: "NOK", DisplayName: "", Symbol: "kr"}, "VEF": ut.Currency{Currency: "VEF", DisplayName: "", Symbol: "Bs"}, "UYU": ut.Currency{Currency: "UYU", DisplayName: "", Symbol: "$"}, "AUD": ut.Currency{Currency: "AUD", DisplayName: "", Symbol: "A$"}, "CLP": ut.Currency{Currency: "CLP", DisplayName: "", Symbol: "$"}, "EUR": ut.Currency{Currency: "EUR", DisplayName: "", Symbol: "€"}, "MUR": ut.Currency{Currency: "MUR", DisplayName: "", Symbol: "Rs"}, "SGD": ut.Currency{Currency: "SGD", DisplayName: "", Symbol: "$"}, "TOP": ut.Currency{Currency: "TOP", DisplayName: "", Symbol: "T$"}}