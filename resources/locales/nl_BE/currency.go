package nl_BE

import "github.com/go-playground/universal-translator"

var currencies = ut.CurrencyFormatValue{"BHD": ut.Currency{Currency: "BHD", DisplayName: "Bahreinse dinar", Symbol: "BHD"}, "MGA": ut.Currency{Currency: "MGA", DisplayName: "Malagassische ariary", Symbol: "MGA"}, "XXX": ut.Currency{Currency: "XXX", DisplayName: "onbekende munteenheid", Symbol: "XXX"}, "ATS": ut.Currency{Currency: "ATS", DisplayName: "Oostenrijkse schilling", Symbol: "ATS"}, "ILR": ut.Currency{Currency: "ILR", DisplayName: "Israëlische sjekel (1980–1985)", Symbol: "ILR"}, "CHE": ut.Currency{Currency: "CHE", DisplayName: "WIR euro", Symbol: "CHE"}, "XSU": ut.Currency{Currency: "XSU", DisplayName: "Sucre", Symbol: "XSU"}, "YER": ut.Currency{Currency: "YER", DisplayName: "Jemenitische rial", Symbol: "YER"}, "NIO": ut.Currency{Currency: "NIO", DisplayName: "Nicaraguaanse córdoba", Symbol: "NIO"}, "TMT": ut.Currency{Currency: "TMT", DisplayName: "Turkmeense manat", Symbol: "TMT"}, "UYI": ut.Currency{Currency: "UYI", DisplayName: "Uruguayaanse peso en geïndexeerde eenheden", Symbol: "UYI"}, "AWG": ut.Currency{Currency: "AWG", DisplayName: "Arubaanse gulden", Symbol: "AWG"}, "USS": ut.Currency{Currency: "USS", DisplayName: "Amerikaanse dollar (zelfde dag)", Symbol: "USS"}, "TWD": ut.Currency{Currency: "TWD", DisplayName: "Nieuwe Taiwanese dollar", Symbol: "NT$"}, "MXV": ut.Currency{Currency: "MXV", DisplayName: "Mexicaanse unidad de inversion (UDI)", Symbol: "MXV"}, "UAH": ut.Currency{Currency: "UAH", DisplayName: "Oekraïense hryvnia", Symbol: "UAH"}, "AON": ut.Currency{Currency: "AON", DisplayName: "Angolese nieuwe kwanza (1990–2000)", Symbol: "AON"}, "BRR": ut.Currency{Currency: "BRR", DisplayName: "Braziliaanse cruzeiro", Symbol: "BRR"}, "LUF": ut.Currency{Currency: "LUF", DisplayName: "Luxemburgse frank", Symbol: "LUF"}, "BOL": ut.Currency{Currency: "BOL", DisplayName: "Boliviaanse boliviano (1863–1963)", Symbol: "BOL"}, "MZE": ut.Currency{Currency: "MZE", DisplayName: "Mozambikaanse escudo", Symbol: "MZE"}, "CDF": ut.Currency{Currency: "CDF", DisplayName: "Congolese frank", Symbol: "CDF"}, "LUL": ut.Currency{Currency: "LUL", DisplayName: "Luxemburgse financiële franc", Symbol: "LUL"}, "HUF": ut.Currency{Currency: "HUF", DisplayName: "Hongaarse forint", Symbol: "HUF"}, "TPE": ut.Currency{Currency: "TPE", DisplayName: "Timorese escudo", Symbol: "TPE"}, "FIM": ut.Currency{Currency: "FIM", DisplayName: "Finse markka", Symbol: "FIM"}, "XTS": ut.Currency{Currency: "XTS", DisplayName: "Valutacode voor testdoeleinden", Symbol: "XTS"}, "RHD": ut.Currency{Currency: "RHD", DisplayName: "Rhodesische dollar", Symbol: "RHD"}, "CLP": ut.Currency{Currency: "CLP", DisplayName: "Chileense peso", Symbol: "CLP"}, "LAK": ut.Currency{Currency: "LAK", DisplayName: "Laotiaanse kip", Symbol: "LAK"}, "DOP": ut.Currency{Currency: "DOP", DisplayName: "Dominicaanse peso", Symbol: "DOP"}, "GBP": ut.Currency{Currency: "GBP", DisplayName: "Brits pond", Symbol: "£"}, "PEI": ut.Currency{Currency: "PEI", DisplayName: "Peruaanse inti", Symbol: "PEI"}, "SRG": ut.Currency{Currency: "SRG", DisplayName: "Surinaamse gulden", Symbol: "SRG"}, "BYR": ut.Currency{Currency: "BYR", DisplayName: "Wit-Russische roebel", Symbol: "BYR"}, "MZM": ut.Currency{Currency: "MZM", DisplayName: "Oude Mozambikaanse metical", Symbol: "MZM"}, "SEK": ut.Currency{Currency: "SEK", DisplayName: "Zweedse kroon", Symbol: "SEK"}, "BAM": ut.Currency{Currency: "BAM", DisplayName: "Bosnische convertibele mark", Symbol: "BAM"}, "GWP": ut.Currency{Currency: "GWP", DisplayName: "Guinee-Bissause peso", Symbol: "GWP"}, "NOK": ut.Currency{Currency: "NOK", DisplayName: "Noorse kroon", Symbol: "NOK"}, "LUC": ut.Currency{Currency: "LUC", DisplayName: "Luxemburgse convertibele franc", Symbol: "LUC"}, "IEP": ut.Currency{Currency: "IEP", DisplayName: "Iers pond", Symbol: "IEP"}, "KPW": ut.Currency{Currency: "KPW", DisplayName: "Noord-Koreaanse won", Symbol: "KPW"}, "UZS": ut.Currency{Currency: "UZS", DisplayName: "Oezbeekse sum", Symbol: "UZS"}, "MUR": ut.Currency{Currency: "MUR", DisplayName: "Mauritiaanse roepie", Symbol: "MUR"}, "ARP": ut.Currency{Currency: "ARP", DisplayName: "Argentijnse peso (1983–1985)", Symbol: "ARP"}, "GRD": ut.Currency{Currency: "GRD", DisplayName: "Griekse drachme", Symbol: "GRD"}, "XBB": ut.Currency{Currency: "XBB", DisplayName: "Europese monetaire eenheid", Symbol: "XBB"}, "KRH": ut.Currency{Currency: "KRH", DisplayName: "Zuid-Koreaanse hwan (1953–1962)", Symbol: "KRH"}, "XAU": ut.Currency{Currency: "XAU", DisplayName: "Goud", Symbol: "XAU"}, "ETB": ut.Currency{Currency: "ETB", DisplayName: "Ethiopische birr", Symbol: "ETB"}, "BEF": ut.Currency{Currency: "BEF", DisplayName: "Belgische frank", Symbol: "BEF"}, "BUK": ut.Currency{Currency: "BUK", DisplayName: "Birmese kyat", Symbol: "BUK"}, "CUC": ut.Currency{Currency: "CUC", DisplayName: "Cubaanse convertibele peso", Symbol: "CUC"}, "LTL": ut.Currency{Currency: "LTL", DisplayName: "Litouwse litas", Symbol: "LTL"}, "SRD": ut.Currency{Currency: "SRD", DisplayName: "Surinaamse dollar", Symbol: "SRD"}, "EUR": ut.Currency{Currency: "EUR", DisplayName: "Euro", Symbol: "€"}, "GHS": ut.Currency{Currency: "GHS", DisplayName: "Ghanese cedi", Symbol: "GHS"}, "AED": ut.Currency{Currency: "AED", DisplayName: "Verenigde Arabische Emiraten-dirham", Symbol: "AED"}, "BRL": ut.Currency{Currency: "BRL", DisplayName: "Braziliaanse real", Symbol: "R$"}, "BYB": ut.Currency{Currency: "BYB", DisplayName: "Wit-Russische nieuwe roebel (1994–1999)", Symbol: "BYB"}, "XAF": ut.Currency{Currency: "XAF", DisplayName: "CFA-frank", Symbol: "FCFA"}, "MKD": ut.Currency{Currency: "MKD", DisplayName: "Macedonische denar", Symbol: "MKD"}, "UYU": ut.Currency{Currency: "UYU", DisplayName: "Uruguayaanse peso", Symbol: "UYU"}, "CLE": ut.Currency{Currency: "CLE", DisplayName: "Chileense escudo", Symbol: "CLE"}, "BRE": ut.Currency{Currency: "BRE", DisplayName: "Braziliaanse cruzeiro (1990–1993)", Symbol: "BRE"}, "MXP": ut.Currency{Currency: "MXP", DisplayName: "Mexicaanse zilveren peso (1861–1992)", Symbol: "MXP"}, "KWD": ut.Currency{Currency: "KWD", DisplayName: "Koeweitse dinar", Symbol: "KWD"}, "YUN": ut.Currency{Currency: "YUN", DisplayName: "Joegoslavische convertibele dinar", Symbol: "YUN"}, "TZS": ut.Currency{Currency: "TZS", DisplayName: "Tanzaniaanse shilling", Symbol: "TZS"}, "KRW": ut.Currency{Currency: "KRW", DisplayName: "Zuid-Koreaanse won", Symbol: "₩"}, "LVR": ut.Currency{Currency: "LVR", DisplayName: "Letse roebel", Symbol: "LVR"}, "BOV": ut.Currency{Currency: "BOV", DisplayName: "Boliviaanse mvdol", Symbol: "BOV"}, "NLG": ut.Currency{Currency: "NLG", DisplayName: "Nederlandse gulden", Symbol: "NLG"}, "XFO": ut.Currency{Currency: "XFO", DisplayName: "Franse gouden franc", Symbol: "XFO"}, "BGO": ut.Currency{Currency: "BGO", DisplayName: "Bulgaarse lev (1879–1952)", Symbol: "BGO"}, "PGK": ut.Currency{Currency: "PGK", DisplayName: "Papoea-Nieuw-Guinese kina", Symbol: "PGK"}, "BGN": ut.Currency{Currency: "BGN", DisplayName: "Bulgaarse lev", Symbol: "BGN"}, "MGF": ut.Currency{Currency: "MGF", DisplayName: "Malagassische franc", Symbol: "MGF"}, "FKP": ut.Currency{Currency: "FKP", DisplayName: "Falklandeilands pond", Symbol: "FKP"}, "JOD": ut.Currency{Currency: "JOD", DisplayName: "Jordaanse dinar", Symbol: "JOD"}, "XAG": ut.Currency{Currency: "XAG", DisplayName: "Zilver", Symbol: "XAG"}, "GNS": ut.Currency{Currency: "GNS", DisplayName: "Guinese syli", Symbol: "GNS"}, "MAD": ut.Currency{Currency: "MAD", DisplayName: "Marokkaanse dirham", Symbol: "MAD"}, "GNF": ut.Currency{Currency: "GNF", DisplayName: "Guinese frank", Symbol: "GNF"}, "ZAR": ut.Currency{Currency: "ZAR", DisplayName: "Zuid-Afrikaanse rand", Symbol: "ZAR"}, "MCF": ut.Currency{Currency: "MCF", DisplayName: "Monegaskische frank", Symbol: "MCF"}, "ALK": ut.Currency{Currency: "ALK", DisplayName: "Albanese lek (1946–1965)", Symbol: "ALK"}, "BOP": ut.Currency{Currency: "BOP", DisplayName: "Boliviaanse peso", Symbol: "BOP"}, "AOR": ut.Currency{Currency: "AOR", DisplayName: "Angolese kwanza reajustado (1995–1999)", Symbol: "AOR"}, "CZK": ut.Currency{Currency: "CZK", DisplayName: "Tsjechische kroon", Symbol: "CZK"}, "XPF": ut.Currency{Currency: "XPF", DisplayName: "CFP-frank", Symbol: "XPF"}, "VEB": ut.Currency{Currency: "VEB", DisplayName: "Venezolaanse bolivar (1871–2008)", Symbol: "VEB"}, "SIT": ut.Currency{Currency: "SIT", DisplayName: "Sloveense tolar", Symbol: "SIT"}, "JPY": ut.Currency{Currency: "JPY", DisplayName: "Japanse yen", Symbol: "JP¥"}, "MZN": ut.Currency{Currency: "MZN", DisplayName: "Mozambikaanse metical", Symbol: "MZN"}, "DDM": ut.Currency{Currency: "DDM", DisplayName: "Oost-Duitse ostmark", Symbol: "DDM"}, "SLL": ut.Currency{Currency: "SLL", DisplayName: "Sierraleoonse leone", Symbol: "SLL"}, "UAK": ut.Currency{Currency: "UAK", DisplayName: "Oekraïense karbovanetz", Symbol: "UAK"}, "MRO": ut.Currency{Currency: "MRO", DisplayName: "Mauritaanse ouguiya", Symbol: "MRO"}, "GHC": ut.Currency{Currency: "GHC", DisplayName: "Ghanese cedi (1979–2007)", Symbol: "GHC"}, "XUA": ut.Currency{Currency: "XUA", DisplayName: "ADB-rekeneenheid", Symbol: "XUA"}, "SUR": ut.Currency{Currency: "SUR", DisplayName: "Sovjet-roebel", Symbol: "SUR"}, "HNL": ut.Currency{Currency: "HNL", DisplayName: "Hondurese lempira", Symbol: "HNL"}, "IDR": ut.Currency{Currency: "IDR", DisplayName: "Indonesische roepia", Symbol: "IDR"}, "CNX": ut.Currency{Currency: "CNX", DisplayName: "dollar van de Chinese Volksbank", Symbol: "CNX"}, "GWE": ut.Currency{Currency: "GWE", DisplayName: "Portugees-Guinese escudo", Symbol: "GWE"}, "GQE": ut.Currency{Currency: "GQE", DisplayName: "Equatoriaal-Guinese ekwele guineana", Symbol: "GQE"}, "PLN": ut.Currency{Currency: "PLN", DisplayName: "Poolse zloty", Symbol: "PLN"}, "MDC": ut.Currency{Currency: "MDC", DisplayName: "Moldavische cupon", Symbol: "MDC"}, "BRZ": ut.Currency{Currency: "BRZ", DisplayName: "Braziliaanse cruzeiro (1942–1967)", Symbol: "BRZ"}, "PAB": ut.Currency{Currency: "PAB", DisplayName: "Panamese balboa", Symbol: "PAB"}, "MDL": ut.Currency{Currency: "MDL", DisplayName: "Moldavische leu", Symbol: "MDL"}, "RUB": ut.Currency{Currency: "RUB", DisplayName: "Russische roebel", Symbol: "RUB"}, "GIP": ut.Currency{Currency: "GIP", DisplayName: "Gibraltarees pond", Symbol: "GIP"}, "BTN": ut.Currency{Currency: "BTN", DisplayName: "Bhutaanse ngultrum", Symbol: "BTN"}, "GEK": ut.Currency{Currency: "GEK", DisplayName: "Georgische kupon larit", Symbol: "GEK"}, "MVP": ut.Currency{Currency: "MVP", DisplayName: "Maldivische roepie", Symbol: "MVP"}, "MKN": ut.Currency{Currency: "MKN", DisplayName: "Macedonische denar (1992–1993)", Symbol: "MKN"}, "PEN": ut.Currency{Currency: "PEN", DisplayName: "Peruaanse nieuwe sol", Symbol: "PEN"}, "ESB": ut.Currency{Currency: "ESB", DisplayName: "Spaanse peseta (convertibele account)", Symbol: "ESB"}, "FJD": ut.Currency{Currency: "FJD", DisplayName: "Fiji-dollar", Symbol: "FJ$"}, "NAD": ut.Currency{Currency: "NAD", DisplayName: "Namibische dollar", Symbol: "NAD"}, "DKK": ut.Currency{Currency: "DKK", DisplayName: "Deense kroon", Symbol: "DKK"}, "BOB": ut.Currency{Currency: "BOB", DisplayName: "Boliviaanse boliviano", Symbol: "BOB"}, "DEM": ut.Currency{Currency: "DEM", DisplayName: "Duitse mark", Symbol: "DEM"}, "SBD": ut.Currency{Currency: "SBD", DisplayName: "Salomon-dollar", Symbol: "SI$"}, "RON": ut.Currency{Currency: "RON", DisplayName: "Roemeense leu", Symbol: "RON"}, "BBD": ut.Currency{Currency: "BBD", DisplayName: "Barbadaanse dollar", Symbol: "BBD"}, "YUD": ut.Currency{Currency: "YUD", DisplayName: "Joegoslavische harde dinar", Symbol: "YUD"}, "GYD": ut.Currency{Currency: "GYD", DisplayName: "Guyaanse dollar", Symbol: "GYD"}, "PES": ut.Currency{Currency: "PES", DisplayName: "Peruaanse sol", Symbol: "PES"}, "COP": ut.Currency{Currency: "COP", DisplayName: "Colombiaanse peso", Symbol: "COP"}, "SDD": ut.Currency{Currency: "SDD", DisplayName: "Soedanese dinar", Symbol: "SDD"}, "IQD": ut.Currency{Currency: "IQD", DisplayName: "Iraakse dinar", Symbol: "IQD"}, "MYR": ut.Currency{Currency: "MYR", DisplayName: "Maleisische ringgit", Symbol: "MYR"}, "KMF": ut.Currency{Currency: "KMF", DisplayName: "Comorese frank", Symbol: "KMF"}, "ARS": ut.Currency{Currency: "ARS", DisplayName: "Argentijnse peso", Symbol: "ARS"}, "MNT": ut.Currency{Currency: "MNT", DisplayName: "Mongoolse tugrik", Symbol: "MNT"}, "FRF": ut.Currency{Currency: "FRF", DisplayName: "Franse franc", Symbol: "FRF"}, "SZL": ut.Currency{Currency: "SZL", DisplayName: "Swazische lilangeni", Symbol: "SZL"}, "AFN": ut.Currency{Currency: "AFN", DisplayName: "Afghaanse afghani", Symbol: "AFN"}, "LKR": ut.Currency{Currency: "LKR", DisplayName: "Sri Lankaanse roepie", Symbol: "LKR"}, "BAN": ut.Currency{Currency: "BAN", DisplayName: "Nieuwe Bosnische dinar (1994–1997)", Symbol: "BAN"}, "ESA": ut.Currency{Currency: "ESA", DisplayName: "Spaanse peseta (account A)", Symbol: "ESA"}, "USN": ut.Currency{Currency: "USN", DisplayName: "Amerikaanse dollar (volgende dag)", Symbol: "USN"}, "RSD": ut.Currency{Currency: "RSD", DisplayName: "Servische dinar", Symbol: "RSD"}, "BSD": ut.Currency{Currency: "BSD", DisplayName: "Bahamaanse dollar", Symbol: "BSD"}, "SDP": ut.Currency{Currency: "SDP", DisplayName: "Soedanees pond (1957–1998)", Symbol: "SDP"}, "MAF": ut.Currency{Currency: "MAF", DisplayName: "Marokkaanse franc", Symbol: "MAF"}, "MOP": ut.Currency{Currency: "MOP", DisplayName: "Macause pataca", Symbol: "MOP"}, "KHR": ut.Currency{Currency: "KHR", DisplayName: "Cambodjaanse riel", Symbol: "KHR"}, "AZM": ut.Currency{Currency: "AZM", DisplayName: "Azerbeidzjaanse manat (1993–2006)", Symbol: "AZM"}, "ARM": ut.Currency{Currency: "ARM", DisplayName: "Argentijnse peso (1881–1970)", Symbol: "ARM"}, "SKK": ut.Currency{Currency: "SKK", DisplayName: "Slowaakse koruna", Symbol: "SKK"}, "MTP": ut.Currency{Currency: "MTP", DisplayName: "Maltees pond", Symbol: "MTP"}, "ZWD": ut.Currency{Currency: "ZWD", DisplayName: "Zimbabwaanse dollar", Symbol: "ZWD"}, "JMD": ut.Currency{Currency: "JMD", DisplayName: "Jamaicaanse dollar", Symbol: "JMD"}, "ECS": ut.Currency{Currency: "ECS", DisplayName: "Ecuadoraanse sucre", Symbol: "ECS"}, "GEL": ut.Currency{Currency: "GEL", DisplayName: "Georgische lari", Symbol: "GEL"}, "BMD": ut.Currency{Currency: "BMD", DisplayName: "Bermuda-dollar", Symbol: "BMD"}, "NIC": ut.Currency{Currency: "NIC", DisplayName: "Nicaraguaanse córdoba (1988–1991)", Symbol: "NIC"}, "LVL": ut.Currency{Currency: "LVL", DisplayName: "Letse lats", Symbol: "LVL"}, "ERN": ut.Currency{Currency: "ERN", DisplayName: "Eritrese nakfa", Symbol: "ERN"}, "ZMW": ut.Currency{Currency: "ZMW", DisplayName: "Zambiaanse kwacha", Symbol: "ZMW"}, "XBC": ut.Currency{Currency: "XBC", DisplayName: "Europese rekeneenheid (XBC)", Symbol: "XBC"}, "PTE": ut.Currency{Currency: "PTE", DisplayName: "Portugese escudo", Symbol: "PTE"}, "BGL": ut.Currency{Currency: "BGL", DisplayName: "Bulgaarse harde lev", Symbol: "BGL"}, "CYP": ut.Currency{Currency: "CYP", DisplayName: "Cyprisch pond", Symbol: "CYP"}, "ZAL": ut.Currency{Currency: "ZAL", DisplayName: "Zuid-Afrikaanse rand (financieel)", Symbol: "ZAL"}, "CSD": ut.Currency{Currency: "CSD", DisplayName: "Oude Servische dinar", Symbol: "CSD"}, "KES": ut.Currency{Currency: "KES", DisplayName: "Keniaanse shilling", Symbol: "KES"}, "BZD": ut.Currency{Currency: "BZD", DisplayName: "Belizaanse dollar", Symbol: "BZD"}, "TMM": ut.Currency{Currency: "TMM", DisplayName: "Turkmeense manat (1993–2009)", Symbol: "TMM"}, "CUP": ut.Currency{Currency: "CUP", DisplayName: "Cubaanse peso", Symbol: "CUP"}, "IRR": ut.Currency{Currency: "IRR", DisplayName: "Iraanse rial", Symbol: "IRR"}, "SSP": ut.Currency{Currency: "SSP", DisplayName: "Zuid-Soedanees pond", Symbol: "SSP"}, "NZD": ut.Currency{Currency: "NZD", DisplayName: "Nieuw-Zeelandse dollar", Symbol: "NZ$"}, "XOF": ut.Currency{Currency: "XOF", DisplayName: "CFA-franc BCEAO", Symbol: "CFA"}, "YDD": ut.Currency{Currency: "YDD", DisplayName: "Jemenitische dinar", Symbol: "YDD"}, "TJS": ut.Currency{Currency: "TJS", DisplayName: "Tadzjiekse somoni", Symbol: "TJS"}, "PKR": ut.Currency{Currency: "PKR", DisplayName: "Pakistaanse roepie", Symbol: "PKR"}, "CNY": ut.Currency{Currency: "CNY", DisplayName: "Chinese renminbi", Symbol: "CN¥"}, "CHF": ut.Currency{Currency: "CHF", DisplayName: "Zwitserse frank", Symbol: "CHF"}, "ROL": ut.Currency{Currency: "ROL", DisplayName: "Oude Roemeense leu", Symbol: "ROL"}, "YUR": ut.Currency{Currency: "YUR", DisplayName: "Joegoslavische hervormde dinar (1992–1993)", Symbol: "YUR"}, "DJF": ut.Currency{Currency: "DJF", DisplayName: "Djiboutiaanse frank", Symbol: "DJF"}, "BRC": ut.Currency{Currency: "BRC", DisplayName: "Braziliaanse cruzado", Symbol: "BRC"}, "SCR": ut.Currency{Currency: "SCR", DisplayName: "Seychelse roepie", Symbol: "SCR"}, "LSL": ut.Currency{Currency: "LSL", DisplayName: "Lesothaanse loti", Symbol: "LSL"}, "PHP": ut.Currency{Currency: "PHP", DisplayName: "Filipijnse peso", Symbol: "PHP"}, "AFA": ut.Currency{Currency: "AFA", DisplayName: "Afghani (1927–2002)", Symbol: "AFA"}, "SOS": ut.Currency{Currency: "SOS", DisplayName: "Somalische shilling", Symbol: "SOS"}, "KRO": ut.Currency{Currency: "KRO", DisplayName: "Oude Zuid-Koreaanse won (1945–1953)", Symbol: "KRO"}, "BAD": ut.Currency{Currency: "BAD", DisplayName: "Bosnische dinar", Symbol: "BAD"}, "AOK": ut.Currency{Currency: "AOK", DisplayName: "Angolese kwanza (1977–1990)", Symbol: "AOK"}, "MMK": ut.Currency{Currency: "MMK", DisplayName: "Myanmarese kyat", Symbol: "MMK"}, "TRL": ut.Currency{Currency: "TRL", DisplayName: "Turkse lire", Symbol: "TRL"}, "ECV": ut.Currency{Currency: "ECV", DisplayName: "Ecuadoraanse unidad de valor constante (UVC)", Symbol: "ECV"}, "ALL": ut.Currency{Currency: "ALL", DisplayName: "Albanese lek", Symbol: "ALL"}, "TTD": ut.Currency{Currency: "TTD", DisplayName: "Trinidad en Tobago-dollar", Symbol: "TTD"}, "XCD": ut.Currency{Currency: "XCD", DisplayName: "Oost-Caribische dollar", Symbol: "EC$"}, "KGS": ut.Currency{Currency: "KGS", DisplayName: "Kirgizische som", Symbol: "KGS"}, "USD": ut.Currency{Currency: "USD", DisplayName: "Amerikaanse dollar", Symbol: "US$"}, "AZN": ut.Currency{Currency: "AZN", DisplayName: "Azerbeidzjaanse manat", Symbol: "AZN"}, "SYP": ut.Currency{Currency: "SYP", DisplayName: "Syrisch pond", Symbol: "SYP"}, "AOA": ut.Currency{Currency: "AOA", DisplayName: "Angolese kwanza", Symbol: "AOA"}, "XBD": ut.Currency{Currency: "XBD", DisplayName: "Europese rekeneenheid (XBD)", Symbol: "XBD"}, "ISJ": ut.Currency{Currency: "ISJ", DisplayName: "IJslandse kroon (1918–1981)", Symbol: "ISJ"}, "XPD": ut.Currency{Currency: "XPD", DisplayName: "Palladium", Symbol: "XPD"}, "BND": ut.Currency{Currency: "BND", DisplayName: "Bruneise dollar", Symbol: "BND"}, "ZWL": ut.Currency{Currency: "ZWL", DisplayName: "Zimbabwaanse dollar (2009)", Symbol: "ZWL"}, "XEU": ut.Currency{Currency: "XEU", DisplayName: "European Currency Unit", Symbol: "XEU"}, "MLF": ut.Currency{Currency: "MLF", DisplayName: "Malinese franc", Symbol: "MLF"}, "ARA": ut.Currency{Currency: "ARA", DisplayName: "Argentijnse austral", Symbol: "ARA"}, "XPT": ut.Currency{Currency: "XPT", DisplayName: "Platina", Symbol: "XPT"}, "COU": ut.Currency{Currency: "COU", DisplayName: "Unidad de Valor Real", Symbol: "COU"}, "RUR": ut.Currency{Currency: "RUR", DisplayName: "Russische roebel (1991–1998)", Symbol: "RUR"}, "CVE": ut.Currency{Currency: "CVE", DisplayName: "Kaapverdische escudo", Symbol: "CVE"}, "RWF": ut.Currency{Currency: "RWF", DisplayName: "Rwandese frank", Symbol: "RWF"}, "MWK": ut.Currency{Currency: "MWK", DisplayName: "Malawische kwacha", Symbol: "MWK"}, "BEC": ut.Currency{Currency: "BEC", DisplayName: "Belgische frank (convertibel)", Symbol: "BEC"}, "KZT": ut.Currency{Currency: "KZT", DisplayName: "Kazachse tenge", Symbol: "KZT"}, "LYD": ut.Currency{Currency: "LYD", DisplayName: "Libische dinar", Symbol: "LYD"}, "SGD": ut.Currency{Currency: "SGD", DisplayName: "Singaporese dollar", Symbol: "SGD"}, "DZD": ut.Currency{Currency: "DZD", DisplayName: "Algerijnse dinar", Symbol: "DZD"}, "BEL": ut.Currency{Currency: "BEL", DisplayName: "Belgische frank (financieel)", Symbol: "BEL"}, "ILP": ut.Currency{Currency: "ILP", DisplayName: "Israëlisch pond", Symbol: "ILP"}, "UGX": ut.Currency{Currency: "UGX", DisplayName: "Oegandese shilling", Symbol: "UGX"}, "CAD": ut.Currency{Currency: "CAD", DisplayName: "Canadese dollar", Symbol: "C$"}, "MVR": ut.Currency{Currency: "MVR", DisplayName: "Maldivische rufiyaa", Symbol: "MVR"}, "PLZ": ut.Currency{Currency: "PLZ", DisplayName: "Poolse zloty (1950–1995)", Symbol: "PLZ"}, "ANG": ut.Currency{Currency: "ANG", DisplayName: "Nederlands-Antilliaanse gulden", Symbol: "ANG"}, "MTL": ut.Currency{Currency: "MTL", DisplayName: "Maltese lire", Symbol: "MTL"}, "TOP": ut.Currency{Currency: "TOP", DisplayName: "Tongaanse paʻanga", Symbol: "TOP"}, "THB": ut.Currency{Currency: "THB", DisplayName: "Thaise baht", Symbol: "฿"}, "CLF": ut.Currency{Currency: "CLF", DisplayName: "Chileense unidades de fomento", Symbol: "CLF"}, "BDT": ut.Currency{Currency: "BDT", DisplayName: "Bengalese taka", Symbol: "BDT"}, "ARL": ut.Currency{Currency: "ARL", DisplayName: "Argentijnse peso ley (1970–1983)", Symbol: "ARL"}, "EEK": ut.Currency{Currency: "EEK", DisplayName: "Estlandse kroon", Symbol: "EEK"}, "ILS": ut.Currency{Currency: "ILS", DisplayName: "Israëlische nieuwe shekel", Symbol: "₪"}, "UGS": ut.Currency{Currency: "UGS", DisplayName: "Oegandese shilling (1966–1987)", Symbol: "UGS"}, "ZRZ": ut.Currency{Currency: "ZRZ", DisplayName: "Zaïrese zaïre", Symbol: "ZRZ"}, "XDR": ut.Currency{Currency: "XDR", DisplayName: "Special Drawing Rights", Symbol: "XDR"}, "TJR": ut.Currency{Currency: "TJR", DisplayName: "Tadzjikistaanse roebel", Symbol: "TJR"}, "LTT": ut.Currency{Currency: "LTT", DisplayName: "Litouwse talonas", Symbol: "LTT"}, "XFU": ut.Currency{Currency: "XFU", DisplayName: "Franse UIC-franc", Symbol: "XFU"}, "SHP": ut.Currency{Currency: "SHP", DisplayName: "Sint-Heleens pond", Symbol: "SHP"}, "BIF": ut.Currency{Currency: "BIF", DisplayName: "Burundese frank", Symbol: "BIF"}, "EGP": ut.Currency{Currency: "EGP", DisplayName: "Egyptisch pond", Symbol: "EGP"}, "LBP": ut.Currency{Currency: "LBP", DisplayName: "Libanees pond", Symbol: "LBP"}, "XBA": ut.Currency{Currency: "XBA", DisplayName: "Europese samengestelde eenheid", Symbol: "XBA"}, "TND": ut.Currency{Currency: "TND", DisplayName: "Tunesische dinar", Symbol: "TND"}, "NPR": ut.Currency{Currency: "NPR", DisplayName: "Nepalese roepie", Symbol: "NPR"}, "AUD": ut.Currency{Currency: "AUD", DisplayName: "Australische dollar", Symbol: "AU$"}, "HRK": ut.Currency{Currency: "HRK", DisplayName: "Kroatische kuna", Symbol: "HRK"}, "HRD": ut.Currency{Currency: "HRD", DisplayName: "Kroatische dinar", Symbol: "HRD"}, "BWP": ut.Currency{Currency: "BWP", DisplayName: "Botswaanse pula", Symbol: "BWP"}, "TRY": ut.Currency{Currency: "TRY", DisplayName: "Turkse lira", Symbol: "TRY"}, "AMD": ut.Currency{Currency: "AMD", DisplayName: "Armeense dram", Symbol: "AMD"}, "VUV": ut.Currency{Currency: "VUV", DisplayName: "Vanuatuaanse vatu", Symbol: "VUV"}, "PYG": ut.Currency{Currency: "PYG", DisplayName: "Paraguayaanse guarani", Symbol: "PYG"}, "VND": ut.Currency{Currency: "VND", DisplayName: "Vietnamese dong", Symbol: "₫"}, "ESP": ut.Currency{Currency: "ESP", DisplayName: "Spaanse peseta", Symbol: "ESP"}, "BRB": ut.Currency{Currency: "BRB", DisplayName: "Braziliaanse cruzeiro novo (1967–1986)", Symbol: "BRB"}, "KYD": ut.Currency{Currency: "KYD", DisplayName: "Caymaneilandse dollar", Symbol: "KYD"}, "ZMK": ut.Currency{Currency: "ZMK", DisplayName: "Zambiaanse kwacha (1968–2012)", Symbol: "ZMK"}, "HTG": ut.Currency{Currency: "HTG", DisplayName: "Haïtiaanse gourde", Symbol: "HTG"}, "SVC": ut.Currency{Currency: "SVC", DisplayName: "Salvadoraanse colón", Symbol: "SVC"}, "HKD": ut.Currency{Currency: "HKD", DisplayName: "Hongkongse dollar", Symbol: "HK$"}, "MXN": ut.Currency{Currency: "MXN", DisplayName: "Mexicaanse peso", Symbol: "MX$"}, "UYP": ut.Currency{Currency: "UYP", DisplayName: "Uruguayaanse peso (1975–1993)", Symbol: "UYP"}, "CSK": ut.Currency{Currency: "CSK", DisplayName: "Tsjechoslowaakse harde koruna", Symbol: "CSK"}, "VNN": ut.Currency{Currency: "VNN", DisplayName: "Vietnamese dong (1978–1985)", Symbol: "VNN"}, "ISK": ut.Currency{Currency: "ISK", DisplayName: "IJslandse kroon", Symbol: "ISK"}, "STD": ut.Currency{Currency: "STD", DisplayName: "Santomese dobra", Symbol: "STD"}, "XRE": ut.Currency{Currency: "XRE", DisplayName: "RINET-fondsen", Symbol: "XRE"}, "VEF": ut.Currency{Currency: "VEF", DisplayName: "Venezolaanse bolivar", Symbol: "VEF"}, "CRC": ut.Currency{Currency: "CRC", DisplayName: "Costa Ricaanse colon", Symbol: "CRC"}, "SDG": ut.Currency{Currency: "SDG", DisplayName: "Soedanees pond", Symbol: "SDG"}, "YUM": ut.Currency{Currency: "YUM", DisplayName: "Joegoslavische noviy-dinar", Symbol: "YUM"}, "ZRN": ut.Currency{Currency: "ZRN", DisplayName: "Zaïrese nieuwe zaïre", Symbol: "ZRN"}, "SAR": ut.Currency{Currency: "SAR", DisplayName: "Saoedi-Arabische riyal", Symbol: "SAR"}, "ZWR": ut.Currency{Currency: "ZWR", DisplayName: "Zimbabwaanse dollar (2008)", Symbol: "ZWR"}, "INR": ut.Currency{Currency: "INR", DisplayName: "Indiase roepie", Symbol: "₹"}, "LRD": ut.Currency{Currency: "LRD", DisplayName: "Liberiaanse dollar", Symbol: "LRD"}, "ITL": ut.Currency{Currency: "ITL", DisplayName: "Italiaanse lire", Symbol: "ITL"}, "QAR": ut.Currency{Currency: "QAR", DisplayName: "Qatarese rial", Symbol: "QAR"}, "WST": ut.Currency{Currency: "WST", DisplayName: "Samoaanse tala", Symbol: "WST"}, "NGN": ut.Currency{Currency: "NGN", DisplayName: "Nigeriaanse naira", Symbol: "NGN"}, "CHW": ut.Currency{Currency: "CHW", DisplayName: "WIR franc", Symbol: "CHW"}, "BRN": ut.Currency{Currency: "BRN", DisplayName: "Braziliaanse cruzado novo", Symbol: "BRN"}, "BGM": ut.Currency{Currency: "BGM", DisplayName: "Bulgaarse socialistische lev", Symbol: "BGM"}, "GMD": ut.Currency{Currency: "GMD", DisplayName: "Gambiaanse dalasi", Symbol: "GMD"}, "GTQ": ut.Currency{Currency: "GTQ", DisplayName: "Guatemalteekse quetzal", Symbol: "GTQ"}, "ADP": ut.Currency{Currency: "ADP", DisplayName: "Andorrese peseta", Symbol: "ADP"}, "OMR": ut.Currency{Currency: "OMR", DisplayName: "Omaanse rial", Symbol: "OMR"}}