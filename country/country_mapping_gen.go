package country

var CountryByName = map[string]Country{
	"Afghanistan":                            Afghanistan,
	"Albania":                                Albania,
	"Algeria":                                Algeria,
	"American Samoa":                         AmericanSamoa,
	"Andorra":                                Andorra,
	"Angola":                                 Angola,
	"Anguilla":                               Anguilla,
	"Antarctica":                             Antarctica,
	"Antigua and Barbuda":                    AntiguaAndBarbuda,
	"Argentina":                              Argentina,
	"Armenia":                                Armenia,
	"Aruba":                                  Aruba,
	"Australia":                              Australia,
	"Austria":                                Austria,
	"Azerbaijan":                             Azerbaijan,
	"Bahamas (the)":                          Bahamas,
	"Bahrain":                                Bahrain,
	"Bangladesh":                             Bangladesh,
	"Barbados":                               Barbados,
	"Belarus":                                Belarus,
	"Belgium":                                Belgium,
	"Belize":                                 Belize,
	"Benin":                                  Benin,
	"Bermuda":                                Bermuda,
	"Bhutan":                                 Bhutan,
	"Bolivia (Plurinational State of)":       Bolivia,
	"Bonaire, Sint Eustatius and Saba":       BonaireSintEustatiusAndSaba,
	"Bosnia and Herzegovina":                 BosniaAndHerzegovina,
	"Botswana":                               Botswana,
	"Bouvet Island":                          BouvetIsland,
	"Brazil":                                 Brazil,
	"British Indian Ocean Territory (the)":   BritishIndianOceanTerritory,
	"Brunei Darussalam":                      BruneiDarussalam,
	"Bulgaria":                               Bulgaria,
	"Burkina Faso":                           BurkinaFaso,
	"Burundi":                                Burundi,
	"Cabo Verde":                             CaboVerde,
	"Cambodia":                               Cambodia,
	"Cameroon":                               Cameroon,
	"Canada":                                 Canada,
	"Cayman Islands (the)":                   CaymanIslands,
	"Central African Republic (the)":         CentralAfricanRepublic,
	"Chad":                                   Chad,
	"Chile":                                  Chile,
	"China":                                  China,
	"Christmas Island":                       ChristmasIsland,
	"Cocos (Keeling) Islands (the)":          Cocos,
	"Colombia":                               Colombia,
	"Comoros (the)":                          Comoros,
	"Congo (the Democratic Republic of the)": DemocraticRepublicOfTheCongo,
	"Congo (the)":                            Congo,
	"Cook Islands (the)":                     CookIslands,
	"Costa Rica":                             CostaRica,
	"Croatia":                                Croatia,
	"Cuba":                                   Cuba,
	"Curaçao":                                Curacao,
	"Cyprus":                                 Cyprus,
	"Czechia":                                Czechia,
	"Côte d'Ivoire":                          CoteDIvoire,
	"Denmark":                                Denmark,
	"Djibouti":                               Djibouti,
	"Dominica":                               Dominica,
	"Dominican Republic (the)":               DominicanRepublic,
	"Ecuador":                                Ecuador,
	"Egypt":                                  Egypt,
	"El Salvador":                            ElSalvador,
	"Equatorial Guinea":                      EquatorialGuinea,
	"Eritrea":                                Eritrea,
	"Estonia":                                Estonia,
	"Eswatini":                               Eswatini,
	"Ethiopia":                               Ethiopia,
	"Falkland Islands (the) [Malvinas]":      FalklandIslands,
	"Faroe Islands (the)":                    FaroeIslands,
	"Fiji":                                   Fiji,
	"Finland":                                Finland,
	"France":                                 France,
	"French Guiana":                          FrenchGuiana,
	"French Polynesia":                       FrenchPolynesia,
	"French Southern Territories (the)":      FrenchSouthernTerritories,
	"Gabon":                                  Gabon,
	"Gambia (the)":                           Gambia,
	"Georgia":                                Georgia,
	"Germany":                                Germany,
	"Ghana":                                  Ghana,
	"Gibraltar":                              Gibraltar,
	"Greece":                                 Greece,
	"Greenland":                              Greenland,
	"Grenada":                                Grenada,
	"Guadeloupe":                             Guadeloupe,
	"Guam":                                   Guam,
	"Guatemala":                              Guatemala,
	"Guernsey":                               Guernsey,
	"Guinea":                                 Guinea,
	"Guinea-Bissau":                          GuineaBissau,
	"Guyana":                                 Guyana,
	"Haiti":                                  Haiti,
	"Heard Island and McDonald Islands":      HeardIslandAndMcDonaldIslands,
	"Holy See (the)":                         HolySee,
	"Honduras":                               Honduras,
	"Hong Kong":                              HongKong,
	"Hungary":                                Hungary,
	"Iceland":                                Iceland,
	"India":                                  India,
	"Indonesia":                              Indonesia,
	"Iran (Islamic Republic of)":             Iran,
	"Iraq":                                   Iraq,
	"Ireland":                                Ireland,
	"Isle of Man":                            IsleOfMan,
	"Israel":                                 Israel,
	"Italy":                                  Italy,
	"Jamaica":                                Jamaica,
	"Japan":                                  Japan,
	"Jersey":                                 Jersey,
	"Jordan":                                 Jordan,
	"Kazakhstan":                             Kazakhstan,
	"Kenya":                                  Kenya,
	"Kiribati":                               Kiribati,
	"Korea (the Democratic People's Republic of)": NorthKorea,
	"Korea (the Republic of)":                     SouthKorea,
	"Kuwait":                                      Kuwait,
	"Kyrgyzstan":                                  Kyrgyzstan,
	"Lao People's Democratic Republic (the)":      LaoPeoplesDemocraticRepublic,
	"Latvia":                                      Latvia,
	"Lebanon":                                     Lebanon,
	"Lesotho":                                     Lesotho,
	"Liberia":                                     Liberia,
	"Libya":                                       Libya,
	"Liechtenstein":                               Liechtenstein,
	"Lithuania":                                   Lithuania,
	"Luxembourg":                                  Luxembourg,
	"Macao":                                       Macao,
	"Madagascar":                                  Madagascar,
	"Malawi":                                      Malawi,
	"Malaysia":                                    Malaysia,
	"Maldives":                                    Maldives,
	"Mali":                                        Mali,
	"Malta":                                       Malta,
	"Marshall Islands (the)":                      MarshallIslands,
	"Martinique":                                  Martinique,
	"Mauritania":                                  Mauritania,
	"Mauritius":                                   Mauritius,
	"Mayotte":                                     Mayotte,
	"Mexico":                                      Mexico,
	"Micronesia (Federated States of)":            Micronesia,
	"Moldova (the Republic of)":                   Moldova,
	"Monaco":                                      Monaco,
	"Mongolia":                                    Mongolia,
	"Montenegro":                                  Montenegro,
	"Montserrat":                                  Montserrat,
	"Morocco":                                     Morocco,
	"Mozambique":                                  Mozambique,
	"Myanmar":                                     Myanmar,
	"Namibia":                                     Namibia,
	"Nauru":                                       Nauru,
	"Nepal":                                       Nepal,
	"Netherlands (Kingdom of the)":                Netherlands,
	"New Caledonia":                               NewCaledonia,
	"New Zealand":                                 NewZealand,
	"Nicaragua":                                   Nicaragua,
	"Niger (the)":                                 Niger,
	"Nigeria":                                     Nigeria,
	"Niue":                                        Niue,
	"Norfolk Island":                              NorfolkIsland,
	"North Macedonia":                             NorthMacedonia,
	"Northern Mariana Islands (the)":              NorthernMarianaIslands,
	"Norway":                                      Norway,
	"Oman":                                        Oman,
	"Pakistan":                                    Pakistan,
	"Palau":                                       Palau,
	"Palestine, State of":                         Palestine,
	"Panama":                                      Panama,
	"Papua New Guinea":                            PapuaNewGuinea,
	"Paraguay":                                    Paraguay,
	"Peru":                                        Peru,
	"Philippines (the)":                           Philippines,
	"Pitcairn":                                    Pitcairn,
	"Poland":                                      Poland,
	"Portugal":                                    Portugal,
	"Puerto Rico":                                 PuertoRico,
	"Qatar":                                       Qatar,
	"Romania":                                     Romania,
	"Russian Federation (the)":                    RussianFederation,
	"Rwanda":                                      Rwanda,
	"Réunion":                                     Reunion,
	"Saint Barthélemy":                            SaintBarthelemy,
	"Saint Helena, Ascension and Tristan da Cunha": SaintHelenaAscensionAndTristanDaCunha,
	"Saint Kitts and Nevis":                        SaintKittsAndNevis,
	"Saint Lucia":                                  SaintLucia,
	"Saint Martin (French part)":                   SaintMartin,
	"Saint Pierre and Miquelon":                    SaintPierreAndMiquelon,
	"Saint Vincent and the Grenadines":             SaintVincentAndTheGrenadines,
	"Samoa":                                        Samoa,
	"San Marino":                                   SanMarino,
	"Sao Tome and Principe":                        SaoTomeAndPrincipe,
	"Saudi Arabia":                                 SaudiArabia,
	"Senegal":                                      Senegal,
	"Serbia":                                       Serbia,
	"Seychelles":                                   Seychelles,
	"Sierra Leone":                                 SierraLeone,
	"Singapore":                                    Singapore,
	"Sint Maarten (Dutch part)":                    SintMaarten,
	"Slovakia":                                     Slovakia,
	"Slovenia":                                     Slovenia,
	"Solomon Islands":                              SolomonIslands,
	"Somalia":                                      Somalia,
	"South Africa":                                 SouthAfrica,
	"South Georgia and the South Sandwich Islands": SouthGeorgiaAndTheSouthSandwichIslands,
	"South Sudan":                      SouthSudan,
	"Spain":                            Spain,
	"Sri Lanka":                        SriLanka,
	"Sudan (the)":                      Sudan,
	"Suriname":                         Suriname,
	"Svalbard and Jan Mayen":           SvalbardAndJanMayen,
	"Sweden":                           Sweden,
	"Switzerland":                      Switzerland,
	"Syrian Arab Republic (the)":       SyrianArabRepublic,
	"Taiwan (Province of China)":       Taiwan,
	"Tajikistan":                       Tajikistan,
	"Tanzania, the United Republic of": TanzaniaTheUnitedRepublicOf,
	"Thailand":                         Thailand,
	"Timor-Leste":                      TimorLeste,
	"Togo":                             Togo,
	"Tokelau":                          Tokelau,
	"Tonga":                            Tonga,
	"Trinidad and Tobago":              TrinidadAndTobago,
	"Tunisia":                          Tunisia,
	"Turkmenistan":                     Turkmenistan,
	"Turks and Caicos Islands (the)":   TurksAndCaicosIslands,
	"Tuvalu":                           Tuvalu,
	"Türkiye":                          Turkiye,
	"Uganda":                           Uganda,
	"Ukraine":                          Ukraine,
	"United Arab Emirates (the)":       UnitedArabEmirates,
	"United Kingdom of Great Britain and Northern Ireland (the)": UnitedKingdomOfGreatBritainAndNorthernIreland,
	"United States Minor Outlying Islands (the)":                 UnitedStatesMinorOutlyingIslands,
	"United States of America (the)":                             UnitedStatesOfAmerica,
	"Uruguay":                                                    Uruguay,
	"Uzbekistan":                                                 Uzbekistan,
	"Vanuatu":                                                    Vanuatu,
	"Venezuela (Bolivarian Republic of)":                         Venezuela,
	"Viet Nam":                                                   VietNam,
	"Virgin Islands (British)":                                   BritishVirginIslands,
	"Virgin Islands (U.S.)":                                      USVirginIslands,
	"Wallis and Futuna":                                          WallisAndFutuna,
	"Western Sahara*":                                            WesternSahara,
	"Yemen":                                                      Yemen,
	"Zambia":                                                     Zambia,
	"Zimbabwe":                                                   Zimbabwe,
	"Åland Islands":                                              AlandIslands,
}

var CountryByAlpha2 = map[string]Country{
	"AF": Afghanistan,
	"AL": Albania,
	"DZ": Algeria,
	"AS": AmericanSamoa,
	"AD": Andorra,
	"AO": Angola,
	"AI": Anguilla,
	"AQ": Antarctica,
	"AG": AntiguaAndBarbuda,
	"AR": Argentina,
	"AM": Armenia,
	"AW": Aruba,
	"AU": Australia,
	"AT": Austria,
	"AZ": Azerbaijan,
	"BS": Bahamas,
	"BH": Bahrain,
	"BD": Bangladesh,
	"BB": Barbados,
	"BY": Belarus,
	"BE": Belgium,
	"BZ": Belize,
	"BJ": Benin,
	"BM": Bermuda,
	"BT": Bhutan,
	"BO": Bolivia,
	"BQ": BonaireSintEustatiusAndSaba,
	"BA": BosniaAndHerzegovina,
	"BW": Botswana,
	"BV": BouvetIsland,
	"BR": Brazil,
	"IO": BritishIndianOceanTerritory,
	"BN": BruneiDarussalam,
	"BG": Bulgaria,
	"BF": BurkinaFaso,
	"BI": Burundi,
	"CV": CaboVerde,
	"KH": Cambodia,
	"CM": Cameroon,
	"CA": Canada,
	"KY": CaymanIslands,
	"CF": CentralAfricanRepublic,
	"TD": Chad,
	"CL": Chile,
	"CN": China,
	"CX": ChristmasIsland,
	"CC": Cocos,
	"CO": Colombia,
	"KM": Comoros,
	"CD": DemocraticRepublicOfTheCongo,
	"CG": Congo,
	"CK": CookIslands,
	"CR": CostaRica,
	"HR": Croatia,
	"CU": Cuba,
	"CW": Curacao,
	"CY": Cyprus,
	"CZ": Czechia,
	"CI": CoteDIvoire,
	"DK": Denmark,
	"DJ": Djibouti,
	"DM": Dominica,
	"DO": DominicanRepublic,
	"EC": Ecuador,
	"EG": Egypt,
	"SV": ElSalvador,
	"GQ": EquatorialGuinea,
	"ER": Eritrea,
	"EE": Estonia,
	"SZ": Eswatini,
	"ET": Ethiopia,
	"FK": FalklandIslands,
	"FO": FaroeIslands,
	"FJ": Fiji,
	"FI": Finland,
	"FR": France,
	"GF": FrenchGuiana,
	"PF": FrenchPolynesia,
	"TF": FrenchSouthernTerritories,
	"GA": Gabon,
	"GM": Gambia,
	"GE": Georgia,
	"DE": Germany,
	"GH": Ghana,
	"GI": Gibraltar,
	"GR": Greece,
	"GL": Greenland,
	"GD": Grenada,
	"GP": Guadeloupe,
	"GU": Guam,
	"GT": Guatemala,
	"GG": Guernsey,
	"GN": Guinea,
	"GW": GuineaBissau,
	"GY": Guyana,
	"HT": Haiti,
	"HM": HeardIslandAndMcDonaldIslands,
	"VA": HolySee,
	"HN": Honduras,
	"HK": HongKong,
	"HU": Hungary,
	"IS": Iceland,
	"IN": India,
	"ID": Indonesia,
	"IR": Iran,
	"IQ": Iraq,
	"IE": Ireland,
	"IM": IsleOfMan,
	"IL": Israel,
	"IT": Italy,
	"JM": Jamaica,
	"JP": Japan,
	"JE": Jersey,
	"JO": Jordan,
	"KZ": Kazakhstan,
	"KE": Kenya,
	"KI": Kiribati,
	"KP": NorthKorea,
	"KR": SouthKorea,
	"KW": Kuwait,
	"KG": Kyrgyzstan,
	"LA": LaoPeoplesDemocraticRepublic,
	"LV": Latvia,
	"LB": Lebanon,
	"LS": Lesotho,
	"LR": Liberia,
	"LY": Libya,
	"LI": Liechtenstein,
	"LT": Lithuania,
	"LU": Luxembourg,
	"MO": Macao,
	"MG": Madagascar,
	"MW": Malawi,
	"MY": Malaysia,
	"MV": Maldives,
	"ML": Mali,
	"MT": Malta,
	"MH": MarshallIslands,
	"MQ": Martinique,
	"MR": Mauritania,
	"MU": Mauritius,
	"YT": Mayotte,
	"MX": Mexico,
	"FM": Micronesia,
	"MD": Moldova,
	"MC": Monaco,
	"MN": Mongolia,
	"ME": Montenegro,
	"MS": Montserrat,
	"MA": Morocco,
	"MZ": Mozambique,
	"MM": Myanmar,
	"NA": Namibia,
	"NR": Nauru,
	"NP": Nepal,
	"NL": Netherlands,
	"NC": NewCaledonia,
	"NZ": NewZealand,
	"NI": Nicaragua,
	"NE": Niger,
	"NG": Nigeria,
	"NU": Niue,
	"NF": NorfolkIsland,
	"MK": NorthMacedonia,
	"MP": NorthernMarianaIslands,
	"NO": Norway,
	"OM": Oman,
	"PK": Pakistan,
	"PW": Palau,
	"PS": Palestine,
	"PA": Panama,
	"PG": PapuaNewGuinea,
	"PY": Paraguay,
	"PE": Peru,
	"PH": Philippines,
	"PN": Pitcairn,
	"PL": Poland,
	"PT": Portugal,
	"PR": PuertoRico,
	"QA": Qatar,
	"RO": Romania,
	"RU": RussianFederation,
	"RW": Rwanda,
	"RE": Reunion,
	"BL": SaintBarthelemy,
	"SH": SaintHelenaAscensionAndTristanDaCunha,
	"KN": SaintKittsAndNevis,
	"LC": SaintLucia,
	"MF": SaintMartin,
	"PM": SaintPierreAndMiquelon,
	"VC": SaintVincentAndTheGrenadines,
	"WS": Samoa,
	"SM": SanMarino,
	"ST": SaoTomeAndPrincipe,
	"SA": SaudiArabia,
	"SN": Senegal,
	"RS": Serbia,
	"SC": Seychelles,
	"SL": SierraLeone,
	"SG": Singapore,
	"SX": SintMaarten,
	"SK": Slovakia,
	"SI": Slovenia,
	"SB": SolomonIslands,
	"SO": Somalia,
	"ZA": SouthAfrica,
	"GS": SouthGeorgiaAndTheSouthSandwichIslands,
	"SS": SouthSudan,
	"ES": Spain,
	"LK": SriLanka,
	"SD": Sudan,
	"SR": Suriname,
	"SJ": SvalbardAndJanMayen,
	"SE": Sweden,
	"CH": Switzerland,
	"SY": SyrianArabRepublic,
	"TW": Taiwan,
	"TJ": Tajikistan,
	"TZ": TanzaniaTheUnitedRepublicOf,
	"TH": Thailand,
	"TL": TimorLeste,
	"TG": Togo,
	"TK": Tokelau,
	"TO": Tonga,
	"TT": TrinidadAndTobago,
	"TN": Tunisia,
	"TM": Turkmenistan,
	"TC": TurksAndCaicosIslands,
	"TV": Tuvalu,
	"TR": Turkiye,
	"UG": Uganda,
	"UA": Ukraine,
	"AE": UnitedArabEmirates,
	"GB": UnitedKingdomOfGreatBritainAndNorthernIreland,
	"UM": UnitedStatesMinorOutlyingIslands,
	"US": UnitedStatesOfAmerica,
	"UY": Uruguay,
	"UZ": Uzbekistan,
	"VU": Vanuatu,
	"VE": Venezuela,
	"VN": VietNam,
	"VG": BritishVirginIslands,
	"VI": USVirginIslands,
	"WF": WallisAndFutuna,
	"EH": WesternSahara,
	"YE": Yemen,
	"ZM": Zambia,
	"ZW": Zimbabwe,
	"AX": AlandIslands,
}

var CountryByAlpha3 = map[string]Country{
	"AFG": Afghanistan,
	"ALB": Albania,
	"DZA": Algeria,
	"ASM": AmericanSamoa,
	"AND": Andorra,
	"AGO": Angola,
	"AIA": Anguilla,
	"ATA": Antarctica,
	"ATG": AntiguaAndBarbuda,
	"ARG": Argentina,
	"ARM": Armenia,
	"ABW": Aruba,
	"AUS": Australia,
	"AUT": Austria,
	"AZE": Azerbaijan,
	"BHS": Bahamas,
	"BHR": Bahrain,
	"BGD": Bangladesh,
	"BRB": Barbados,
	"BLR": Belarus,
	"BEL": Belgium,
	"BLZ": Belize,
	"BEN": Benin,
	"BMU": Bermuda,
	"BTN": Bhutan,
	"BOL": Bolivia,
	"BES": BonaireSintEustatiusAndSaba,
	"BIH": BosniaAndHerzegovina,
	"BWA": Botswana,
	"BVT": BouvetIsland,
	"BRA": Brazil,
	"IOT": BritishIndianOceanTerritory,
	"BRN": BruneiDarussalam,
	"BGR": Bulgaria,
	"BFA": BurkinaFaso,
	"BDI": Burundi,
	"CPV": CaboVerde,
	"KHM": Cambodia,
	"CMR": Cameroon,
	"CAN": Canada,
	"CYM": CaymanIslands,
	"CAF": CentralAfricanRepublic,
	"TCD": Chad,
	"CHL": Chile,
	"CHN": China,
	"CXR": ChristmasIsland,
	"CCK": Cocos,
	"COL": Colombia,
	"COM": Comoros,
	"COD": DemocraticRepublicOfTheCongo,
	"COG": Congo,
	"COK": CookIslands,
	"CRI": CostaRica,
	"HRV": Croatia,
	"CUB": Cuba,
	"CUW": Curacao,
	"CYP": Cyprus,
	"CZE": Czechia,
	"CIV": CoteDIvoire,
	"DNK": Denmark,
	"DJI": Djibouti,
	"DMA": Dominica,
	"DOM": DominicanRepublic,
	"ECU": Ecuador,
	"EGY": Egypt,
	"SLV": ElSalvador,
	"GNQ": EquatorialGuinea,
	"ERI": Eritrea,
	"EST": Estonia,
	"SWZ": Eswatini,
	"ETH": Ethiopia,
	"FLK": FalklandIslands,
	"FRO": FaroeIslands,
	"FJI": Fiji,
	"FIN": Finland,
	"FRA": France,
	"GUF": FrenchGuiana,
	"PYF": FrenchPolynesia,
	"ATF": FrenchSouthernTerritories,
	"GAB": Gabon,
	"GMB": Gambia,
	"GEO": Georgia,
	"DEU": Germany,
	"GHA": Ghana,
	"GIB": Gibraltar,
	"GRC": Greece,
	"GRL": Greenland,
	"GRD": Grenada,
	"GLP": Guadeloupe,
	"GUM": Guam,
	"GTM": Guatemala,
	"GGY": Guernsey,
	"GIN": Guinea,
	"GNB": GuineaBissau,
	"GUY": Guyana,
	"HTI": Haiti,
	"HMD": HeardIslandAndMcDonaldIslands,
	"VAT": HolySee,
	"HND": Honduras,
	"HKG": HongKong,
	"HUN": Hungary,
	"ISL": Iceland,
	"IND": India,
	"IDN": Indonesia,
	"IRN": Iran,
	"IRQ": Iraq,
	"IRL": Ireland,
	"IMN": IsleOfMan,
	"ISR": Israel,
	"ITA": Italy,
	"JAM": Jamaica,
	"JPN": Japan,
	"JEY": Jersey,
	"JOR": Jordan,
	"KAZ": Kazakhstan,
	"KEN": Kenya,
	"KIR": Kiribati,
	"PRK": NorthKorea,
	"KOR": SouthKorea,
	"KWT": Kuwait,
	"KGZ": Kyrgyzstan,
	"LAO": LaoPeoplesDemocraticRepublic,
	"LVA": Latvia,
	"LBN": Lebanon,
	"LSO": Lesotho,
	"LBR": Liberia,
	"LBY": Libya,
	"LIE": Liechtenstein,
	"LTU": Lithuania,
	"LUX": Luxembourg,
	"MAC": Macao,
	"MDG": Madagascar,
	"MWI": Malawi,
	"MYS": Malaysia,
	"MDV": Maldives,
	"MLI": Mali,
	"MLT": Malta,
	"MHL": MarshallIslands,
	"MTQ": Martinique,
	"MRT": Mauritania,
	"MUS": Mauritius,
	"MYT": Mayotte,
	"MEX": Mexico,
	"FSM": Micronesia,
	"MDA": Moldova,
	"MCO": Monaco,
	"MNG": Mongolia,
	"MNE": Montenegro,
	"MSR": Montserrat,
	"MAR": Morocco,
	"MOZ": Mozambique,
	"MMR": Myanmar,
	"NAM": Namibia,
	"NRU": Nauru,
	"NPL": Nepal,
	"NLD": Netherlands,
	"NCL": NewCaledonia,
	"NZL": NewZealand,
	"NIC": Nicaragua,
	"NER": Niger,
	"NGA": Nigeria,
	"NIU": Niue,
	"NFK": NorfolkIsland,
	"MKD": NorthMacedonia,
	"MNP": NorthernMarianaIslands,
	"NOR": Norway,
	"OMN": Oman,
	"PAK": Pakistan,
	"PLW": Palau,
	"PSE": Palestine,
	"PAN": Panama,
	"PNG": PapuaNewGuinea,
	"PRY": Paraguay,
	"PER": Peru,
	"PHL": Philippines,
	"PCN": Pitcairn,
	"POL": Poland,
	"PRT": Portugal,
	"PRI": PuertoRico,
	"QAT": Qatar,
	"ROU": Romania,
	"RUS": RussianFederation,
	"RWA": Rwanda,
	"REU": Reunion,
	"BLM": SaintBarthelemy,
	"SHN": SaintHelenaAscensionAndTristanDaCunha,
	"KNA": SaintKittsAndNevis,
	"LCA": SaintLucia,
	"MAF": SaintMartin,
	"SPM": SaintPierreAndMiquelon,
	"VCT": SaintVincentAndTheGrenadines,
	"WSM": Samoa,
	"SMR": SanMarino,
	"STP": SaoTomeAndPrincipe,
	"SAU": SaudiArabia,
	"SEN": Senegal,
	"SRB": Serbia,
	"SYC": Seychelles,
	"SLE": SierraLeone,
	"SGP": Singapore,
	"SXM": SintMaarten,
	"SVK": Slovakia,
	"SVN": Slovenia,
	"SLB": SolomonIslands,
	"SOM": Somalia,
	"ZAF": SouthAfrica,
	"SGS": SouthGeorgiaAndTheSouthSandwichIslands,
	"SSD": SouthSudan,
	"ESP": Spain,
	"LKA": SriLanka,
	"SDN": Sudan,
	"SUR": Suriname,
	"SJM": SvalbardAndJanMayen,
	"SWE": Sweden,
	"CHE": Switzerland,
	"SYR": SyrianArabRepublic,
	"TWN": Taiwan,
	"TJK": Tajikistan,
	"TZA": TanzaniaTheUnitedRepublicOf,
	"THA": Thailand,
	"TLS": TimorLeste,
	"TGO": Togo,
	"TKL": Tokelau,
	"TON": Tonga,
	"TTO": TrinidadAndTobago,
	"TUN": Tunisia,
	"TKM": Turkmenistan,
	"TCA": TurksAndCaicosIslands,
	"TUV": Tuvalu,
	"TUR": Turkiye,
	"UGA": Uganda,
	"UKR": Ukraine,
	"ARE": UnitedArabEmirates,
	"GBR": UnitedKingdomOfGreatBritainAndNorthernIreland,
	"UMI": UnitedStatesMinorOutlyingIslands,
	"USA": UnitedStatesOfAmerica,
	"URY": Uruguay,
	"UZB": Uzbekistan,
	"VUT": Vanuatu,
	"VEN": Venezuela,
	"VNM": VietNam,
	"VGB": BritishVirginIslands,
	"VIR": USVirginIslands,
	"WLF": WallisAndFutuna,
	"ESH": WesternSahara,
	"YEM": Yemen,
	"ZMB": Zambia,
	"ZWE": Zimbabwe,
	"ALA": AlandIslands,
}
