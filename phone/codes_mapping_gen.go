package phone

import "github.com/mikekonan/go-types/v2/country"

var dialCodeByCountryAlpha2Str = map[string]DialCode{
	"TW": TW,
	"AF": AF,
	"AL": AL,
	"DZ": DZ,
	"AS": AS,
	"AD": AD,
	"AO": AO,
	"AI": AI,
	"AQ": AQ,
	"AG": AG,
	"AR": AR,
	"AM": AM,
	"AW": AW,
	"AU": AU,
	"AT": AT,
	"AZ": AZ,
	"BS": BS,
	"BH": BH,
	"BD": BD,
	"BB": BB,
	"BY": BY,
	"BE": BE,
	"BZ": BZ,
	"BJ": BJ,
	"BM": BM,
	"BT": BT,
	"BO": BO,
	"BQ": BQ,
	"BA": BA,
	"BW": BW,
	"BV": BV,
	"BR": BR,
	"IO": IO,
	"VG": VG,
	"BN": BN,
	"BG": BG,
	"BF": BF,
	"BI": BI,
	"CV": CV,
	"KH": KH,
	"CM": CM,
	"CA": CA,
	"KY": KY,
	"CF": CF,
	"TD": TD,
	"CL": CL,
	"CN": CN,
	"HK": HK,
	"MO": MO,
	"CX": CX,
	"CC": CC,
	"CO": CO,
	"KM": KM,
	"CG": CG,
	"CK": CK,
	"CR": CR,
	"HR": HR,
	"CU": CU,
	"CW": CW,
	"CY": CY,
	"CZ": CZ,
	"CI": CI,
	"KP": KP,
	"CD": CD,
	"DK": DK,
	"DJ": DJ,
	"DM": DM,
	"DO": DO,
	"EC": EC,
	"EG": EG,
	"SV": SV,
	"GQ": GQ,
	"ER": ER,
	"EE": EE,
	"SZ": SZ,
	"ET": ET,
	"FK": FK,
	"FO": FO,
	"FJ": FJ,
	"FI": FI,
	"FR": FR,
	"GF": GF,
	"PF": PF,
	"TF": TF,
	"GA": GA,
	"GM": GM,
	"GE": GE,
	"DE": DE,
	"GH": GH,
	"GI": GI,
	"GR": GR,
	"GL": GL,
	"GD": GD,
	"GP": GP,
	"GU": GU,
	"GT": GT,
	"GG": GG,
	"GN": GN,
	"GW": GW,
	"GY": GY,
	"HT": HT,
	"HM": HM,
	"VA": VA,
	"HN": HN,
	"HU": HU,
	"IS": IS,
	"IN": IN,
	"ID": ID,
	"IR": IR,
	"IQ": IQ,
	"IE": IE,
	"IM": IM,
	"IL": IL,
	"IT": IT,
	"JM": JM,
	"JP": JP,
	"JE": JE,
	"JO": JO,
	"KZ": KZ,
	"KE": KE,
	"KI": KI,
	"KW": KW,
	"KG": KG,
	"LA": LA,
	"LV": LV,
	"LB": LB,
	"LS": LS,
	"LR": LR,
	"LY": LY,
	"LI": LI,
	"LT": LT,
	"LU": LU,
	"MG": MG,
	"MW": MW,
	"MY": MY,
	"MV": MV,
	"ML": ML,
	"MT": MT,
	"MH": MH,
	"MQ": MQ,
	"MR": MR,
	"MU": MU,
	"YT": YT,
	"MX": MX,
	"FM": FM,
	"MC": MC,
	"MN": MN,
	"ME": ME,
	"MS": MS,
	"MA": MA,
	"MZ": MZ,
	"MM": MM,
	"NA": NA,
	"NR": NR,
	"NP": NP,
	"NL": NL,
	"NC": NC,
	"NZ": NZ,
	"NI": NI,
	"NE": NE,
	"NG": NG,
	"NU": NU,
	"NF": NF,
	"MP": MP,
	"NO": NO,
	"OM": OM,
	"PK": PK,
	"PW": PW,
	"PA": PA,
	"PG": PG,
	"PY": PY,
	"PE": PE,
	"PH": PH,
	"PN": PN,
	"PL": PL,
	"PT": PT,
	"PR": PR,
	"QA": QA,
	"KR": KR,
	"MD": MD,
	"RO": RO,
	"RU": RU,
	"RW": RW,
	"RE": RE,
	"BL": BL,
	"SH": SH,
	"KN": KN,
	"LC": LC,
	"MF": MF,
	"PM": PM,
	"VC": VC,
	"WS": WS,
	"SM": SM,
	"ST": ST,
	"SA": SA,
	"SN": SN,
	"RS": RS,
	"SC": SC,
	"SL": SL,
	"SG": SG,
	"SX": SX,
	"SK": SK,
	"SI": SI,
	"SB": SB,
	"SO": SO,
	"ZA": ZA,
	"GS": GS,
	"SS": SS,
	"ES": ES,
	"LK": LK,
	"PS": PS,
	"SD": SD,
	"SR": SR,
	"SJ": SJ,
	"SE": SE,
	"CH": CH,
	"SY": SY,
	"TJ": TJ,
	"TH": TH,
	"MK": MK,
	"TL": TL,
	"TG": TG,
	"TK": TK,
	"TO": TO,
	"TT": TT,
	"TN": TN,
	"TR": TR,
	"TM": TM,
	"TC": TC,
	"TV": TV,
	"UG": UG,
	"UA": UA,
	"AE": AE,
	"GB": GB,
	"TZ": TZ,
	"VI": VI,
	"US": US,
	"UY": UY,
	"UZ": UZ,
	"VU": VU,
	"VE": VE,
	"VN": VN,
	"WF": WF,
	"EH": EH,
	"YE": YE,
	"ZM": ZM,
	"ZW": ZW,
	"AX": AX,
}

var countryCodeByDialString = map[string][]country.Alpha2Code{
	"1": {
		country.Alpha2Code("CA"),
		country.Alpha2Code("PR"),
		country.Alpha2Code("US"),
	},
	"7": {
		country.Alpha2Code("KZ"),
		country.Alpha2Code("RU"),
	},
	"20": {
		country.Alpha2Code("EG"),
	},
	"27": {
		country.Alpha2Code("ZA"),
	},
	"30": {
		country.Alpha2Code("GR"),
	},
	"31": {
		country.Alpha2Code("NL"),
	},
	"32": {
		country.Alpha2Code("BE"),
	},
	"33": {
		country.Alpha2Code("FR"),
	},
	"34": {
		country.Alpha2Code("ES"),
	},
	"36": {
		country.Alpha2Code("HU"),
	},
	"39": {
		country.Alpha2Code("IT"),
	},
	"40": {
		country.Alpha2Code("RO"),
	},
	"41": {
		country.Alpha2Code("CH"),
	},
	"43": {
		country.Alpha2Code("AT"),
	},
	"44": {
		country.Alpha2Code("GG"),
		country.Alpha2Code("IM"),
		country.Alpha2Code("JE"),
		country.Alpha2Code("GB"),
	},
	"45": {
		country.Alpha2Code("DK"),
	},
	"46": {
		country.Alpha2Code("SE"),
	},
	"47": {
		country.Alpha2Code("BV"),
		country.Alpha2Code("NO"),
		country.Alpha2Code("SJ"),
	},
	"48": {
		country.Alpha2Code("PL"),
	},
	"49": {
		country.Alpha2Code("DE"),
	},
	"51": {
		country.Alpha2Code("PE"),
	},
	"52": {
		country.Alpha2Code("MX"),
	},
	"53": {
		country.Alpha2Code("CU"),
	},
	"54": {
		country.Alpha2Code("AR"),
	},
	"55": {
		country.Alpha2Code("BR"),
	},
	"56": {
		country.Alpha2Code("CL"),
	},
	"57": {
		country.Alpha2Code("CO"),
	},
	"58": {
		country.Alpha2Code("VE"),
	},
	"60": {
		country.Alpha2Code("MY"),
	},
	"61": {
		country.Alpha2Code("AU"),
		country.Alpha2Code("CX"),
		country.Alpha2Code("CC"),
	},
	"62": {
		country.Alpha2Code("ID"),
	},
	"63": {
		country.Alpha2Code("PH"),
	},
	"64": {
		country.Alpha2Code("NZ"),
	},
	"65": {
		country.Alpha2Code("SG"),
	},
	"66": {
		country.Alpha2Code("TH"),
	},
	"81": {
		country.Alpha2Code("JP"),
	},
	"82": {
		country.Alpha2Code("KR"),
	},
	"84": {
		country.Alpha2Code("VN"),
	},
	"86": {
		country.Alpha2Code("CN"),
	},
	"90": {
		country.Alpha2Code("TR"),
	},
	"91": {
		country.Alpha2Code("IN"),
	},
	"92": {
		country.Alpha2Code("PK"),
	},
	"93": {
		country.Alpha2Code("AF"),
	},
	"94": {
		country.Alpha2Code("LK"),
	},
	"95": {
		country.Alpha2Code("MM"),
	},
	"98": {
		country.Alpha2Code("IR"),
	},
	"211": {
		country.Alpha2Code("SS"),
	},
	"212": {
		country.Alpha2Code("MA"),
		country.Alpha2Code("EH"),
	},
	"213": {
		country.Alpha2Code("DZ"),
	},
	"216": {
		country.Alpha2Code("TN"),
	},
	"218": {
		country.Alpha2Code("LY"),
	},
	"220": {
		country.Alpha2Code("GM"),
	},
	"221": {
		country.Alpha2Code("SN"),
	},
	"222": {
		country.Alpha2Code("MR"),
	},
	"223": {
		country.Alpha2Code("ML"),
	},
	"224": {
		country.Alpha2Code("GN"),
	},
	"225": {
		country.Alpha2Code("CI"),
	},
	"226": {
		country.Alpha2Code("BF"),
	},
	"227": {
		country.Alpha2Code("NE"),
	},
	"228": {
		country.Alpha2Code("TG"),
	},
	"229": {
		country.Alpha2Code("BJ"),
	},
	"230": {
		country.Alpha2Code("MU"),
	},
	"231": {
		country.Alpha2Code("LR"),
	},
	"232": {
		country.Alpha2Code("SL"),
	},
	"233": {
		country.Alpha2Code("GH"),
	},
	"234": {
		country.Alpha2Code("NG"),
	},
	"235": {
		country.Alpha2Code("TD"),
	},
	"236": {
		country.Alpha2Code("CF"),
	},
	"237": {
		country.Alpha2Code("CM"),
	},
	"238": {
		country.Alpha2Code("CV"),
	},
	"239": {
		country.Alpha2Code("ST"),
	},
	"240": {
		country.Alpha2Code("GQ"),
	},
	"241": {
		country.Alpha2Code("GA"),
	},
	"242": {
		country.Alpha2Code("CG"),
	},
	"243": {
		country.Alpha2Code("CD"),
	},
	"244": {
		country.Alpha2Code("AO"),
	},
	"245": {
		country.Alpha2Code("GW"),
	},
	"246": {
		country.Alpha2Code("IO"),
	},
	"248": {
		country.Alpha2Code("SC"),
	},
	"249": {
		country.Alpha2Code("SD"),
	},
	"250": {
		country.Alpha2Code("RW"),
	},
	"251": {
		country.Alpha2Code("ET"),
	},
	"252": {
		country.Alpha2Code("SO"),
	},
	"253": {
		country.Alpha2Code("DJ"),
	},
	"254": {
		country.Alpha2Code("KE"),
	},
	"255": {
		country.Alpha2Code("TZ"),
	},
	"256": {
		country.Alpha2Code("UG"),
	},
	"257": {
		country.Alpha2Code("BI"),
	},
	"258": {
		country.Alpha2Code("MZ"),
	},
	"260": {
		country.Alpha2Code("ZM"),
	},
	"261": {
		country.Alpha2Code("MG"),
	},
	"262": {
		country.Alpha2Code("TF"),
		country.Alpha2Code("YT"),
		country.Alpha2Code("RE"),
	},
	"263": {
		country.Alpha2Code("ZW"),
	},
	"264": {
		country.Alpha2Code("NA"),
	},
	"265": {
		country.Alpha2Code("MW"),
	},
	"266": {
		country.Alpha2Code("LS"),
	},
	"267": {
		country.Alpha2Code("BW"),
	},
	"268": {
		country.Alpha2Code("SZ"),
	},
	"269": {
		country.Alpha2Code("KM"),
	},
	"290": {
		country.Alpha2Code("SH"),
	},
	"291": {
		country.Alpha2Code("ER"),
	},
	"297": {
		country.Alpha2Code("AW"),
	},
	"298": {
		country.Alpha2Code("FO"),
	},
	"299": {
		country.Alpha2Code("GL"),
	},
	"350": {
		country.Alpha2Code("GI"),
	},
	"351": {
		country.Alpha2Code("PT"),
	},
	"352": {
		country.Alpha2Code("LU"),
	},
	"353": {
		country.Alpha2Code("IE"),
	},
	"354": {
		country.Alpha2Code("IS"),
	},
	"355": {
		country.Alpha2Code("AL"),
	},
	"356": {
		country.Alpha2Code("MT"),
	},
	"357": {
		country.Alpha2Code("CY"),
	},
	"358": {
		country.Alpha2Code("FI"),
		country.Alpha2Code("AX"),
	},
	"359": {
		country.Alpha2Code("BG"),
	},
	"370": {
		country.Alpha2Code("LT"),
	},
	"371": {
		country.Alpha2Code("LV"),
	},
	"372": {
		country.Alpha2Code("EE"),
	},
	"373": {
		country.Alpha2Code("MD"),
	},
	"374": {
		country.Alpha2Code("AM"),
	},
	"375": {
		country.Alpha2Code("BY"),
	},
	"376": {
		country.Alpha2Code("AD"),
	},
	"377": {
		country.Alpha2Code("MC"),
	},
	"378": {
		country.Alpha2Code("SM"),
	},
	"380": {
		country.Alpha2Code("UA"),
	},
	"381": {
		country.Alpha2Code("RS"),
	},
	"382": {
		country.Alpha2Code("ME"),
	},
	"385": {
		country.Alpha2Code("HR"),
	},
	"386": {
		country.Alpha2Code("SI"),
	},
	"387": {
		country.Alpha2Code("BA"),
	},
	"389": {
		country.Alpha2Code("MK"),
	},
	"420": {
		country.Alpha2Code("CZ"),
	},
	"421": {
		country.Alpha2Code("SK"),
	},
	"423": {
		country.Alpha2Code("LI"),
	},
	"500": {
		country.Alpha2Code("FK"),
		country.Alpha2Code("GS"),
	},
	"501": {
		country.Alpha2Code("BZ"),
	},
	"502": {
		country.Alpha2Code("GT"),
	},
	"503": {
		country.Alpha2Code("SV"),
	},
	"504": {
		country.Alpha2Code("HN"),
	},
	"505": {
		country.Alpha2Code("NI"),
	},
	"506": {
		country.Alpha2Code("CR"),
	},
	"507": {
		country.Alpha2Code("PA"),
	},
	"508": {
		country.Alpha2Code("PM"),
	},
	"509": {
		country.Alpha2Code("HT"),
	},
	"590": {
		country.Alpha2Code("GP"),
		country.Alpha2Code("BL"),
		country.Alpha2Code("MF"),
	},
	"591": {
		country.Alpha2Code("BO"),
	},
	"592": {
		country.Alpha2Code("GY"),
	},
	"593": {
		country.Alpha2Code("EC"),
	},
	"594": {
		country.Alpha2Code("GF"),
	},
	"595": {
		country.Alpha2Code("PY"),
	},
	"596": {
		country.Alpha2Code("MQ"),
	},
	"597": {
		country.Alpha2Code("SR"),
	},
	"598": {
		country.Alpha2Code("UY"),
	},
	"599": {
		country.Alpha2Code("BQ"),
		country.Alpha2Code("CW"),
	},
	"670": {
		country.Alpha2Code("TL"),
	},
	"672": {
		country.Alpha2Code("AQ"),
		country.Alpha2Code("HM"),
		country.Alpha2Code("NF"),
	},
	"673": {
		country.Alpha2Code("BN"),
	},
	"674": {
		country.Alpha2Code("NR"),
	},
	"675": {
		country.Alpha2Code("PG"),
	},
	"676": {
		country.Alpha2Code("TO"),
	},
	"677": {
		country.Alpha2Code("SB"),
	},
	"678": {
		country.Alpha2Code("VU"),
	},
	"679": {
		country.Alpha2Code("FJ"),
	},
	"680": {
		country.Alpha2Code("PW"),
	},
	"681": {
		country.Alpha2Code("WF"),
	},
	"682": {
		country.Alpha2Code("CK"),
	},
	"683": {
		country.Alpha2Code("NU"),
	},
	"685": {
		country.Alpha2Code("WS"),
	},
	"686": {
		country.Alpha2Code("KI"),
	},
	"687": {
		country.Alpha2Code("NC"),
	},
	"688": {
		country.Alpha2Code("TV"),
	},
	"689": {
		country.Alpha2Code("PF"),
	},
	"690": {
		country.Alpha2Code("TK"),
	},
	"691": {
		country.Alpha2Code("FM"),
	},
	"692": {
		country.Alpha2Code("MH"),
	},
	"850": {
		country.Alpha2Code("KP"),
	},
	"852": {
		country.Alpha2Code("HK"),
	},
	"853": {
		country.Alpha2Code("MO"),
	},
	"855": {
		country.Alpha2Code("KH"),
	},
	"856": {
		country.Alpha2Code("LA"),
	},
	"870": {
		country.Alpha2Code("PN"),
	},
	"880": {
		country.Alpha2Code("BD"),
	},
	"886": {
		country.Alpha2Code("TW"),
	},
	"960": {
		country.Alpha2Code("MV"),
	},
	"961": {
		country.Alpha2Code("LB"),
	},
	"962": {
		country.Alpha2Code("JO"),
	},
	"963": {
		country.Alpha2Code("SY"),
	},
	"964": {
		country.Alpha2Code("IQ"),
	},
	"965": {
		country.Alpha2Code("KW"),
	},
	"966": {
		country.Alpha2Code("SA"),
	},
	"967": {
		country.Alpha2Code("YE"),
	},
	"968": {
		country.Alpha2Code("OM"),
	},
	"970": {
		country.Alpha2Code("PS"),
	},
	"971": {
		country.Alpha2Code("AE"),
	},
	"972": {
		country.Alpha2Code("IL"),
	},
	"973": {
		country.Alpha2Code("BH"),
	},
	"974": {
		country.Alpha2Code("QA"),
	},
	"975": {
		country.Alpha2Code("BT"),
	},
	"976": {
		country.Alpha2Code("MN"),
	},
	"977": {
		country.Alpha2Code("NP"),
	},
	"992": {
		country.Alpha2Code("TJ"),
	},
	"993": {
		country.Alpha2Code("TM"),
	},
	"994": {
		country.Alpha2Code("AZ"),
	},
	"995": {
		country.Alpha2Code("GE"),
	},
	"996": {
		country.Alpha2Code("KG"),
	},
	"998": {
		country.Alpha2Code("UZ"),
	},
	"1242": {
		country.Alpha2Code("BS"),
	},
	"1246": {
		country.Alpha2Code("BB"),
	},
	"1264": {
		country.Alpha2Code("AI"),
	},
	"1268": {
		country.Alpha2Code("AG"),
	},
	"1284": {
		country.Alpha2Code("VG"),
	},
	"1340": {
		country.Alpha2Code("VI"),
	},
	"1345": {
		country.Alpha2Code("KY"),
	},
	"1441": {
		country.Alpha2Code("BM"),
	},
	"1473": {
		country.Alpha2Code("GD"),
	},
	"1649": {
		country.Alpha2Code("TC"),
	},
	"1664": {
		country.Alpha2Code("MS"),
	},
	"1670": {
		country.Alpha2Code("MP"),
	},
	"1671": {
		country.Alpha2Code("GU"),
	},
	"1684": {
		country.Alpha2Code("AS"),
	},
	"1721": {
		country.Alpha2Code("SX"),
	},
	"1758": {
		country.Alpha2Code("LC"),
	},
	"1767": {
		country.Alpha2Code("DM"),
	},
	"1784": {
		country.Alpha2Code("VC"),
	},
	"1809": {
		country.Alpha2Code("DO"),
	},
	"1829": {
		country.Alpha2Code("DO"),
	},
	"1849": {
		country.Alpha2Code("DO"),
	},
	"1868": {
		country.Alpha2Code("TT"),
	},
	"1869": {
		country.Alpha2Code("KN"),
	},
	"1876": {
		country.Alpha2Code("JM"),
	},
	"3906": {
		country.Alpha2Code("VA"),
	},
}
