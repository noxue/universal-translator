package af_NA

import "github.com/go-playground/universal-translator"

var currencies = ut.CurrencyFormatValue{"TND": ut.Currency{Currency: "TND", DisplayName: "Tunisiese dinar", Symbol: "TND"}, "USD": ut.Currency{Currency: "USD", DisplayName: "Amerikaanse dollar", Symbol: "US$"}, "BTN": ut.Currency{Currency: "BTN", DisplayName: "Bhoetanese ngoeltroem", Symbol: "BTN"}, "TWD": ut.Currency{Currency: "TWD", DisplayName: "Nuwe Taiwanese dollar", Symbol: "NT$"}, "KPW": ut.Currency{Currency: "KPW", DisplayName: "Noord-Koreaanse won", Symbol: "KPW"}, "AZN": ut.Currency{Currency: "AZN", DisplayName: "Azerbeidjaanse manat", Symbol: "AZN"}, "LVL": ut.Currency{Currency: "LVL", DisplayName: "Lettiese lats", Symbol: ""}, "NIO": ut.Currency{Currency: "NIO", DisplayName: "Nicaraguaanse córdoba", Symbol: "NIO"}, "CVE": ut.Currency{Currency: "CVE", DisplayName: "Kaap Verdiese escudo", Symbol: "CVE"}, "ILS": ut.Currency{Currency: "ILS", DisplayName: "Israeliese nuwe sikkel", Symbol: "₪"}, "PAB": ut.Currency{Currency: "PAB", DisplayName: "Panamese balboa", Symbol: "PAB"}, "AMD": ut.Currency{Currency: "AMD", DisplayName: "Armeense dram", Symbol: "AMD"}, "EUR": ut.Currency{Currency: "EUR", DisplayName: "Euro", Symbol: "€"}, "BBD": ut.Currency{Currency: "BBD", DisplayName: "Barbados-dollar", Symbol: "BBD"}, "UYU": ut.Currency{Currency: "UYU", DisplayName: "Uruguaanse peso", Symbol: "UYU"}, "JPY": ut.Currency{Currency: "JPY", DisplayName: "Japannese jen", Symbol: "JP¥"}, "ZAR": ut.Currency{Currency: "ZAR", DisplayName: "Suid-Afrikaanse rand", Symbol: "R"}, "SGD": ut.Currency{Currency: "SGD", DisplayName: "Singapoer dollar", Symbol: "SGD"}, "BMD": ut.Currency{Currency: "BMD", DisplayName: "Bermuda-dollar", Symbol: "BMD"}, "SLL": ut.Currency{Currency: "SLL", DisplayName: "Sierra Leoniese leone", Symbol: "SLL"}, "NOK": ut.Currency{Currency: "NOK", DisplayName: "Noorse kroon", Symbol: "NOK"}, "CZK": ut.Currency{Currency: "CZK", DisplayName: "Tsjeggiese kroon", Symbol: "CZK"}, "IQD": ut.Currency{Currency: "IQD", DisplayName: "Irakse dinar", Symbol: "IQD"}, "RUB": ut.Currency{Currency: "RUB", DisplayName: "Russiese roebel", Symbol: "RUB"}, "SDP": ut.Currency{Currency: "SDP", DisplayName: "Soedannese pond (1957–1998)", Symbol: ""}, "GHS": ut.Currency{Currency: "GHS", DisplayName: "Ghanese cedi", Symbol: "GHS"}, "MYR": ut.Currency{Currency: "MYR", DisplayName: "Maleisiese ringgit", Symbol: "MYR"}, "TTD": ut.Currency{Currency: "TTD", DisplayName: "Trinidad en Tobago dollar", Symbol: "TTD"}, "MZM": ut.Currency{Currency: "MZM", DisplayName: "Mosambiekse metical (1980–2006)", Symbol: ""}, "AFN": ut.Currency{Currency: "AFN", DisplayName: "Afgaanse afgani", Symbol: "AFN"}, "MRO": ut.Currency{Currency: "MRO", DisplayName: "Mauritaniese ouguiya", Symbol: "MRO"}, "MXN": ut.Currency{Currency: "MXN", DisplayName: "Meksikaanse peso", Symbol: "MXN"}, "ITL": ut.Currency{Currency: "ITL", DisplayName: "Italiaanse lier", Symbol: ""}, "XPF": ut.Currency{Currency: "XPF", DisplayName: "CFP-frank", Symbol: "CFPF"}, "IDR": ut.Currency{Currency: "IDR", DisplayName: "Indonesiese roepia", Symbol: "IDR"}, "HNL": ut.Currency{Currency: "HNL", DisplayName: "Hondurese lempira", Symbol: "HNL"}, "BAM": ut.Currency{Currency: "BAM", DisplayName: "Bosnies-Herzegowiniese omskakelbare marka", Symbol: "BAM"}, "DOP": ut.Currency{Currency: "DOP", DisplayName: "Dominikaanse peso", Symbol: "DOP"}, "PEN": ut.Currency{Currency: "PEN", DisplayName: "Peruaanse nuwe sol", Symbol: "PEN"}, "KWD": ut.Currency{Currency: "KWD", DisplayName: "Koeweitse dinar", Symbol: "KWD"}, "EGP": ut.Currency{Currency: "EGP", DisplayName: "Egiptiese pond", Symbol: "EGP"}, "MAD": ut.Currency{Currency: "MAD", DisplayName: "Marokkaanse dirham", Symbol: "MAD"}, "MMK": ut.Currency{Currency: "MMK", DisplayName: "Mianmese kyat", Symbol: "MMK"}, "KZT": ut.Currency{Currency: "KZT", DisplayName: "Kazakse tenge", Symbol: "KZT"}, "ALL": ut.Currency{Currency: "ALL", DisplayName: "Albanese lek", Symbol: "ALL"}, "CDF": ut.Currency{Currency: "CDF", DisplayName: "Kongolese frank", Symbol: "CDF"}, "LAK": ut.Currency{Currency: "LAK", DisplayName: "Laosiaanse kip", Symbol: "LAK"}, "HKD": ut.Currency{Currency: "HKD", DisplayName: "Hong Kong dollar", Symbol: "HK$"}, "GBP": ut.Currency{Currency: "GBP", DisplayName: "Britse pond", Symbol: "£"}, "TMT": ut.Currency{Currency: "TMT", DisplayName: "Turkmeense manat", Symbol: "TMT"}, "GYD": ut.Currency{Currency: "GYD", DisplayName: "Guyanese dollar", Symbol: "GYD"}, "SOS": ut.Currency{Currency: "SOS", DisplayName: "Somaliese sjieling", Symbol: "SOS"}, "GIP": ut.Currency{Currency: "GIP", DisplayName: "Gibraltarese pond", Symbol: "GIP"}, "NAD": ut.Currency{Currency: "NAD", DisplayName: "Namibiese dollar", Symbol: "$"}, "KRW": ut.Currency{Currency: "KRW", DisplayName: "Suid-Koreaanse won", Symbol: "₩"}, "ETB": ut.Currency{Currency: "ETB", DisplayName: "Etiopiese birr", Symbol: "ETB"}, "JMD": ut.Currency{Currency: "JMD", DisplayName: "Jamaikaanse dollar", Symbol: "JMD"}, "LSL": ut.Currency{Currency: "LSL", DisplayName: "Lesotho loti", Symbol: ""}, "DKK": ut.Currency{Currency: "DKK", DisplayName: "Deense kroon", Symbol: "DKK"}, "XXX": ut.Currency{Currency: "XXX", DisplayName: "Onbekende geldeenheid", Symbol: ""}, "MDL": ut.Currency{Currency: "MDL", DisplayName: "Moldowiese leu", Symbol: "MDL"}, "DJF": ut.Currency{Currency: "DJF", DisplayName: "Djiboeti frank", Symbol: "DJF"}, "PLN": ut.Currency{Currency: "PLN", DisplayName: "Poolse zloty", Symbol: "PLN"}, "CAD": ut.Currency{Currency: "CAD", DisplayName: "Kanadese dollar", Symbol: "CA$"}, "MWK": ut.Currency{Currency: "MWK", DisplayName: "Malawiese kwacha", Symbol: "MWK"}, "MUR": ut.Currency{Currency: "MUR", DisplayName: "Mauritiaanse rupee", Symbol: "MUR"}, "PKR": ut.Currency{Currency: "PKR", DisplayName: "Pakistanse roepee", Symbol: "PKR"}, "BGN": ut.Currency{Currency: "BGN", DisplayName: "Bulgaarse lev", Symbol: "BGN"}, "IRR": ut.Currency{Currency: "IRR", DisplayName: "Iranse rial", Symbol: "IRR"}, "SHP": ut.Currency{Currency: "SHP", DisplayName: "Sint Helena pond", Symbol: "SHP"}, "AED": ut.Currency{Currency: "AED", DisplayName: "Verenigde Arabiese Emirate dirham", Symbol: "AED"}, "BRL": ut.Currency{Currency: "BRL", DisplayName: "Brasilliaanse reaal", Symbol: "R$"}, "VEF": ut.Currency{Currency: "VEF", DisplayName: "Venezolaanse bolivar", Symbol: "VEF"}, "XAF": ut.Currency{Currency: "XAF", DisplayName: "CFA frank BEAC", Symbol: "FCFA"}, "VND": ut.Currency{Currency: "VND", DisplayName: "Viëtnamese dong", Symbol: "₫"}, "YER": ut.Currency{Currency: "YER", DisplayName: "Jemenitiese rial", Symbol: "YER"}, "CUC": ut.Currency{Currency: "CUC", DisplayName: "Kubaanse omskakelbare peso", Symbol: "CUC"}, "BZD": ut.Currency{Currency: "BZD", DisplayName: "Beliziese dollar", Symbol: "BZD"}, "MOP": ut.Currency{Currency: "MOP", DisplayName: "Macaose pataca", Symbol: "MOP"}, "INR": ut.Currency{Currency: "INR", DisplayName: "Indiese roepee", Symbol: "₹"}, "HUF": ut.Currency{Currency: "HUF", DisplayName: "Hongaarse florint", Symbol: "HUF"}, "RON": ut.Currency{Currency: "RON", DisplayName: "Roemeense leu", Symbol: "RON"}, "HTG": ut.Currency{Currency: "HTG", DisplayName: "Haïtiaanse gourde", Symbol: "HTG"}, "TRL": ut.Currency{Currency: "TRL", DisplayName: "Turkse lier (1922–2005)", Symbol: ""}, "XOF": ut.Currency{Currency: "XOF", DisplayName: "CFA frank BCEAO", Symbol: "CFA"}, "BHD": ut.Currency{Currency: "BHD", DisplayName: "Bahreinse dinar", Symbol: "BHD"}, "RWF": ut.Currency{Currency: "RWF", DisplayName: "Rwandiese frank", Symbol: "RWF"}, "AOA": ut.Currency{Currency: "AOA", DisplayName: "Angolese kwanza", Symbol: "AOA"}, "ERN": ut.Currency{Currency: "ERN", DisplayName: "Eritrese nakfa", Symbol: "ERN"}, "ZWD": ut.Currency{Currency: "ZWD", DisplayName: "Zimbabwiese dollar", Symbol: ""}, "OMR": ut.Currency{Currency: "OMR", DisplayName: "Omaanse rial", Symbol: "OMR"}, "JOD": ut.Currency{Currency: "JOD", DisplayName: "Jordaniese dinar", Symbol: "JOD"}, "PHP": ut.Currency{Currency: "PHP", DisplayName: "Filippynse peso", Symbol: "PHP"}, "KGS": ut.Currency{Currency: "KGS", DisplayName: "Kirgisiese som", Symbol: "KGS"}, "UZS": ut.Currency{Currency: "UZS", DisplayName: "Oezbekiese som", Symbol: "UZS"}, "TZS": ut.Currency{Currency: "TZS", DisplayName: "Tanzaniese sjieling", Symbol: "TZS"}, "GHC": ut.Currency{Currency: "GHC", DisplayName: "Ghanese cedi (1979–2007)", Symbol: ""}, "MNT": ut.Currency{Currency: "MNT", DisplayName: "Mongoolse toegrik", Symbol: "MNT"}, "THB": ut.Currency{Currency: "THB", DisplayName: "Thaise baht", Symbol: "฿"}, "CLP": ut.Currency{Currency: "CLP", DisplayName: "Chileense peso", Symbol: "CLP"}, "XCD": ut.Currency{Currency: "XCD", DisplayName: "Oos-Karibbiese dollar", Symbol: "EC$"}, "MZN": ut.Currency{Currency: "MZN", DisplayName: "Mosambiekse metical", Symbol: "MZN"}, "GMD": ut.Currency{Currency: "GMD", DisplayName: "Gambiese dalasi", Symbol: "GMD"}, "CUP": ut.Currency{Currency: "CUP", DisplayName: "Kubaanse peso", Symbol: "CUP"}, "SEK": ut.Currency{Currency: "SEK", DisplayName: "Sweedse kroon", Symbol: "SEK"}, "SBD": ut.Currency{Currency: "SBD", DisplayName: "Salomonseilande dollar", Symbol: "SBD"}, "SCR": ut.Currency{Currency: "SCR", DisplayName: "Seychellese rupee", Symbol: "SCR"}, "WST": ut.Currency{Currency: "WST", DisplayName: "Samoaanse tala", Symbol: "WST"}, "AWG": ut.Currency{Currency: "AWG", DisplayName: "Arubaanse floryn", Symbol: "AWG"}, "TRY": ut.Currency{Currency: "TRY", DisplayName: "Turkse lier", Symbol: "TRY"}, "BOB": ut.Currency{Currency: "BOB", DisplayName: "Boliviaanse boliviano", Symbol: "BOB"}, "DZD": ut.Currency{Currency: "DZD", DisplayName: "Algeriese dinar", Symbol: "DZD"}, "ZMK": ut.Currency{Currency: "ZMK", DisplayName: "Zambiese kwacha (1968–2012)", Symbol: ""}, "BYR": ut.Currency{Currency: "BYR", DisplayName: "Belo-Russiese roebel", Symbol: "BYR"}, "STD": ut.Currency{Currency: "STD", DisplayName: "São Tomé en Príncipe dobra", Symbol: "STD"}, "PYG": ut.Currency{Currency: "PYG", DisplayName: "Paraguaanse guarani", Symbol: "PYG"}, "COP": ut.Currency{Currency: "COP", DisplayName: "Colombiaanse peso", Symbol: "COP"}, "LTL": ut.Currency{Currency: "LTL", DisplayName: "Litause litas", Symbol: "LTL"}, "LRD": ut.Currency{Currency: "LRD", DisplayName: "Liberiese dollar", Symbol: "LRD"}, "GNF": ut.Currency{Currency: "GNF", DisplayName: "Guinese frank", Symbol: "GNF"}, "BSD": ut.Currency{Currency: "BSD", DisplayName: "Bahamiaanse dollar", Symbol: "BSD"}, "KYD": ut.Currency{Currency: "KYD", DisplayName: "Cayman-eilande dollar", Symbol: "KYD"}, "NGN": ut.Currency{Currency: "NGN", DisplayName: "Nigeriese naira", Symbol: "NGN"}, "ARS": ut.Currency{Currency: "ARS", DisplayName: "Argentynse peso", Symbol: "ARS"}, "LBP": ut.Currency{Currency: "LBP", DisplayName: "Lebanese pond", Symbol: "LBP"}, "ANG": ut.Currency{Currency: "ANG", DisplayName: "Nederlands-Antilliaanse gulde", Symbol: "ANG"}, "GNS": ut.Currency{Currency: "GNS", DisplayName: "Guinese syli", Symbol: ""}, "GEL": ut.Currency{Currency: "GEL", DisplayName: "Georgiese lari", Symbol: "GEL"}, "SDG": ut.Currency{Currency: "SDG", DisplayName: "Soedannese pond", Symbol: "SDG"}, "UGX": ut.Currency{Currency: "UGX", DisplayName: "Ugandese sjieling", Symbol: "UGX"}, "NPR": ut.Currency{Currency: "NPR", DisplayName: "Nepalese roepee", Symbol: "NPR"}, "TJS": ut.Currency{Currency: "TJS", DisplayName: "Tadjikse roebel", Symbol: "TJS"}, "NZD": ut.Currency{Currency: "NZD", DisplayName: "Nieu-Seeland dollar", Symbol: "NZ$"}, "LKR": ut.Currency{Currency: "LKR", DisplayName: "Sri Lankaanse roepee", Symbol: "LKR"}, "BND": ut.Currency{Currency: "BND", DisplayName: "Broeneise dollar", Symbol: "BND"}, "CNY": ut.Currency{Currency: "CNY", DisplayName: "Sjinese joean renminbi", Symbol: "CN¥"}, "QAR": ut.Currency{Currency: "QAR", DisplayName: "Katarrese rial", Symbol: "QAR"}, "RSD": ut.Currency{Currency: "RSD", DisplayName: "Serbiese dinar", Symbol: "RSD"}, "MVR": ut.Currency{Currency: "MVR", DisplayName: "Malediviese rufia", Symbol: "MVR"}, "SYP": ut.Currency{Currency: "SYP", DisplayName: "Siriese pond", Symbol: "SYP"}, "CRC": ut.Currency{Currency: "CRC", DisplayName: "Costa Ricaanse colón", Symbol: "CRC"}, "BIF": ut.Currency{Currency: "BIF", DisplayName: "Burundiese frank", Symbol: "BIF"}, "UAH": ut.Currency{Currency: "UAH", DisplayName: "Oekraïnse hriwna", Symbol: "UAH"}, "TOP": ut.Currency{Currency: "TOP", DisplayName: "Tongaanse pa’anga", Symbol: "TOP"}, "VUV": ut.Currency{Currency: "VUV", DisplayName: "Vanuatuse vatu", Symbol: "VUV"}, "PGK": ut.Currency{Currency: "PGK", DisplayName: "Papoea-Nieu-Guinese kina", Symbol: "PGK"}, "KHR": ut.Currency{Currency: "KHR", DisplayName: "Kambodjaanse riel", Symbol: "KHR"}, "CHF": ut.Currency{Currency: "CHF", DisplayName: "Switserse frank", Symbol: "CHF"}, "KMF": ut.Currency{Currency: "KMF", DisplayName: "Comoraanse frank", Symbol: "KMF"}, "LYD": ut.Currency{Currency: "LYD", DisplayName: "Libiese dinar", Symbol: "LYD"}, "ZMW": ut.Currency{Currency: "ZMW", DisplayName: "Zambiese kwacha", Symbol: "ZMW"}, "ISK": ut.Currency{Currency: "ISK", DisplayName: "Yslandse kroon", Symbol: "ISK"}, "SAR": ut.Currency{Currency: "SAR", DisplayName: "Saoedi-Arabiese riyal", Symbol: "SAR"}, "GTQ": ut.Currency{Currency: "GTQ", DisplayName: "Guatemalaanse quetzal", Symbol: "GTQ"}, "SZL": ut.Currency{Currency: "SZL", DisplayName: "Swazilandse lilangeni", Symbol: "SZL"}, "SSP": ut.Currency{Currency: "SSP", DisplayName: "Suid-Soedanese pond", Symbol: "SSP"}, "AUD": ut.Currency{Currency: "AUD", DisplayName: "Australiese dollar", Symbol: "A$"}, "HRK": ut.Currency{Currency: "HRK", DisplayName: "Kroatiese kuna", Symbol: "HRK"}, "BDT": ut.Currency{Currency: "BDT", DisplayName: "Bangladesjiese taka", Symbol: "BDT"}, "FKP": ut.Currency{Currency: "FKP", DisplayName: "Falkland-eilande pond", Symbol: "FKP"}, "BWP": ut.Currency{Currency: "BWP", DisplayName: "Botswana pula", Symbol: "BWP"}, "KES": ut.Currency{Currency: "KES", DisplayName: "Keniaanse sjieling", Symbol: "KES"}, "MKD": ut.Currency{Currency: "MKD", DisplayName: "Macedoniese denar", Symbol: "MKD"}, "SRD": ut.Currency{Currency: "SRD", DisplayName: "Surinaamse dollar", Symbol: "SRD"}, "MGA": ut.Currency{Currency: "MGA", DisplayName: "Malgassiese ariary", Symbol: "MGA"}, "FJD": ut.Currency{Currency: "FJD", DisplayName: "Fidjiaanse dollar", Symbol: "FJD"}}