package currency

var currenciesByCode = map[string]currency{
	`AFN`: {
		countries: Countries{`AFGHANISTAN`},
		currency:  `Afghani`,
		code:      `AFN`,
		number:    `971`,
	},
	`EUR`: {
		countries: Countries{`ÅLAND ISLANDS`, `ANDORRA`, `AUSTRIA`, `BELGIUM`, `CYPRUS`, `ESTONIA`, `EUROPEAN UNION`, `FINLAND`, `FRANCE`, `FRENCH GUIANA`, `FRENCH SOUTHERN TERRITORIES (THE)`, `GERMANY`, `GREECE`, `GUADELOUPE`, `HOLY SEE (THE)`, `IRELAND`, `ITALY`, `LATVIA`, `LITHUANIA`, `LUXEMBOURG`, `MALTA`, `MARTINIQUE`, `MAYOTTE`, `MONACO`, `MONTENEGRO`, `NETHERLANDS (THE)`, `PORTUGAL`, `RÉUNION`, `SAINT BARTHÉLEMY`, `SAINT MARTIN (FRENCH PART)`, `SAINT PIERRE AND MIQUELON`, `SAN MARINO`, `SLOVAKIA`, `SLOVENIA`, `SPAIN`, `FRENCH SOUTHERN TERRITORIES`, `HOLY SEE`, `NETHERLANDS`, `SAINT MARTIN`},
		currency:  `Euro`,
		code:      `EUR`,
		number:    `978`,
	},
	`ALL`: {
		countries: Countries{`ALBANIA`},
		currency:  `Lek`,
		code:      `ALL`,
		number:    `008`,
	},
	`DZD`: {
		countries: Countries{`ALGERIA`},
		currency:  `Algerian Dinar`,
		code:      `DZD`,
		number:    `012`,
	},
	`USD`: {
		countries: Countries{`AMERICAN SAMOA`, `BONAIRE, SINT EUSTATIUS AND SABA`, `BRITISH INDIAN OCEAN TERRITORY (THE)`, `ECUADOR`, `EL SALVADOR`, `GUAM`, `HAITI`, `MARSHALL ISLANDS (THE)`, `MICRONESIA (FEDERATED STATES OF)`, `NORTHERN MARIANA ISLANDS (THE)`, `PALAU`, `PANAMA`, `PUERTO RICO`, `TIMOR-LESTE`, `TURKS AND CAICOS ISLANDS (THE)`, `UNITED STATES MINOR OUTLYING ISLANDS (THE)`, `UNITED STATES OF AMERICA (THE)`, `VIRGIN ISLANDS (BRITISH)`, `VIRGIN ISLANDS (U.S.)`, `BRITISH INDIAN OCEAN TERRITORY`, `MARSHALL ISLANDS`, `MICRONESIA`, `NORTHERN MARIANA ISLANDS`, `TURKS AND CAICOS ISLANDS`, `UNITED STATES MINOR OUTLYING ISLANDS`, `UNITED STATES OF AMERICA`, `VIRGIN ISLANDS`, `VIRGIN ISLANDS`},
		currency:  `US Dollar`,
		code:      `USD`,
		number:    `840`,
	},
	`AOA`: {
		countries: Countries{`ANGOLA`},
		currency:  `Kwanza`,
		code:      `AOA`,
		number:    `973`,
	},
	`XCD`: {
		countries: Countries{`ANGUILLA`, `ANTIGUA AND BARBUDA`, `DOMINICA`, `GRENADA`, `MONTSERRAT`, `SAINT KITTS AND NEVIS`, `SAINT LUCIA`, `SAINT VINCENT AND THE GRENADINES`},
		currency:  `East Caribbean Dollar`,
		code:      `XCD`,
		number:    `951`,
	},
	`ARS`: {
		countries: Countries{`ARGENTINA`},
		currency:  `Argentine Peso`,
		code:      `ARS`,
		number:    `032`,
	},
	`AMD`: {
		countries: Countries{`ARMENIA`},
		currency:  `Armenian Dram`,
		code:      `AMD`,
		number:    `051`,
	},
	`AWG`: {
		countries: Countries{`ARUBA`},
		currency:  `Aruban Florin`,
		code:      `AWG`,
		number:    `533`,
	},
	`AUD`: {
		countries: Countries{`AUSTRALIA`, `CHRISTMAS ISLAND`, `COCOS (KEELING) ISLANDS (THE)`, `HEARD ISLAND AND McDONALD ISLANDS`, `KIRIBATI`, `NAURU`, `NORFOLK ISLAND`, `TUVALU`, `COCOS ISLANDS`},
		currency:  `Australian Dollar`,
		code:      `AUD`,
		number:    `036`,
	},
	`AZN`: {
		countries: Countries{`AZERBAIJAN`},
		currency:  `Azerbaijan Manat`,
		code:      `AZN`,
		number:    `944`,
	},
	`BSD`: {
		countries: Countries{`BAHAMAS (THE)`, `BAHAMAS`},
		currency:  `Bahamian Dollar`,
		code:      `BSD`,
		number:    `044`,
	},
	`BHD`: {
		countries: Countries{`BAHRAIN`},
		currency:  `Bahraini Dinar`,
		code:      `BHD`,
		number:    `048`,
	},
	`BDT`: {
		countries: Countries{`BANGLADESH`},
		currency:  `Taka`,
		code:      `BDT`,
		number:    `050`,
	},
	`BBD`: {
		countries: Countries{`BARBADOS`},
		currency:  `Barbados Dollar`,
		code:      `BBD`,
		number:    `052`,
	},
	`BYN`: {
		countries: Countries{`BELARUS`},
		currency:  `Belarusian Ruble`,
		code:      `BYN`,
		number:    `933`,
	},
	`BZD`: {
		countries: Countries{`BELIZE`},
		currency:  `Belize Dollar`,
		code:      `BZD`,
		number:    `084`,
	},
	`XOF`: {
		countries: Countries{`BENIN`, `BURKINA FASO`, `CÔTE D'IVOIRE`, `GUINEA-BISSAU`, `MALI`, `NIGER (THE)`, `SENEGAL`, `TOGO`, `NIGER`},
		currency:  `CFA Franc BCEAO`,
		code:      `XOF`,
		number:    `952`,
	},
	`BMD`: {
		countries: Countries{`BERMUDA`},
		currency:  `Bermudian Dollar`,
		code:      `BMD`,
		number:    `060`,
	},
	`INR`: {
		countries: Countries{`BHUTAN`, `INDIA`},
		currency:  `Indian Rupee`,
		code:      `INR`,
		number:    `356`,
	},
	`BTN`: {
		countries: Countries{`BHUTAN`},
		currency:  `Ngultrum`,
		code:      `BTN`,
		number:    `064`,
	},
	`BOB`: {
		countries: Countries{`BOLIVIA (PLURINATIONAL STATE OF)`, `BOLIVIA`},
		currency:  `Boliviano`,
		code:      `BOB`,
		number:    `068`,
	},
	`BOV`: {
		countries: Countries{`BOLIVIA (PLURINATIONAL STATE OF)`, `BOLIVIA`},
		currency:  `Mvdol`,
		code:      `BOV`,
		number:    `984`,
	},
	`BAM`: {
		countries: Countries{`BOSNIA AND HERZEGOVINA`},
		currency:  `Convertible Mark`,
		code:      `BAM`,
		number:    `977`,
	},
	`BWP`: {
		countries: Countries{`BOTSWANA`},
		currency:  `Pula`,
		code:      `BWP`,
		number:    `072`,
	},
	`NOK`: {
		countries: Countries{`BOUVET ISLAND`, `NORWAY`, `SVALBARD AND JAN MAYEN`},
		currency:  `Norwegian Krone`,
		code:      `NOK`,
		number:    `578`,
	},
	`BRL`: {
		countries: Countries{`BRAZIL`},
		currency:  `Brazilian Real`,
		code:      `BRL`,
		number:    `986`,
	},
	`BND`: {
		countries: Countries{`BRUNEI DARUSSALAM`},
		currency:  `Brunei Dollar`,
		code:      `BND`,
		number:    `096`,
	},
	`BGN`: {
		countries: Countries{`BULGARIA`},
		currency:  `Bulgarian Lev`,
		code:      `BGN`,
		number:    `975`,
	},
	`BIF`: {
		countries: Countries{`BURUNDI`},
		currency:  `Burundi Franc`,
		code:      `BIF`,
		number:    `108`,
	},
	`CVE`: {
		countries: Countries{`CABO VERDE`},
		currency:  `Cabo Verde Escudo`,
		code:      `CVE`,
		number:    `132`,
	},
	`KHR`: {
		countries: Countries{`CAMBODIA`},
		currency:  `Riel`,
		code:      `KHR`,
		number:    `116`,
	},
	`XAF`: {
		countries: Countries{`CAMEROON`, `CENTRAL AFRICAN REPUBLIC (THE)`, `CHAD`, `CONGO (THE)`, `EQUATORIAL GUINEA`, `GABON`, `CENTRAL AFRICAN REPUBLIC`, `CONGO`},
		currency:  `CFA Franc BEAC`,
		code:      `XAF`,
		number:    `950`,
	},
	`CAD`: {
		countries: Countries{`CANADA`},
		currency:  `Canadian Dollar`,
		code:      `CAD`,
		number:    `124`,
	},
	`KYD`: {
		countries: Countries{`CAYMAN ISLANDS (THE)`, `CAYMAN ISLANDS`},
		currency:  `Cayman Islands Dollar`,
		code:      `KYD`,
		number:    `136`,
	},
	`CLP`: {
		countries: Countries{`CHILE`},
		currency:  `Chilean Peso`,
		code:      `CLP`,
		number:    `152`,
	},
	`CLF`: {
		countries: Countries{`CHILE`},
		currency:  `Unidad de Fomento`,
		code:      `CLF`,
		number:    `990`,
	},
	`CNY`: {
		countries: Countries{`CHINA`},
		currency:  `Yuan Renminbi`,
		code:      `CNY`,
		number:    `156`,
	},
	`COP`: {
		countries: Countries{`COLOMBIA`},
		currency:  `Colombian Peso`,
		code:      `COP`,
		number:    `170`,
	},
	`COU`: {
		countries: Countries{`COLOMBIA`},
		currency:  `Unidad de Valor Real`,
		code:      `COU`,
		number:    `970`,
	},
	`KMF`: {
		countries: Countries{`COMOROS (THE)`, `COMOROS`},
		currency:  `Comorian Franc `,
		code:      `KMF`,
		number:    `174`,
	},
	`CDF`: {
		countries: Countries{`CONGO (THE DEMOCRATIC REPUBLIC OF THE)`, `CONGO`},
		currency:  `Congolese Franc`,
		code:      `CDF`,
		number:    `976`,
	},
	`NZD`: {
		countries: Countries{`COOK ISLANDS (THE)`, `NEW ZEALAND`, `NIUE`, `PITCAIRN`, `TOKELAU`, `COOK ISLANDS`},
		currency:  `New Zealand Dollar`,
		code:      `NZD`,
		number:    `554`,
	},
	`CRC`: {
		countries: Countries{`COSTA RICA`},
		currency:  `Costa Rican Colon`,
		code:      `CRC`,
		number:    `188`,
	},
	`HRK`: {
		countries: Countries{`CROATIA`},
		currency:  `Kuna`,
		code:      `HRK`,
		number:    `191`,
	},
	`CUP`: {
		countries: Countries{`CUBA`},
		currency:  `Cuban Peso`,
		code:      `CUP`,
		number:    `192`,
	},
	`CUC`: {
		countries: Countries{`CUBA`},
		currency:  `Peso Convertible`,
		code:      `CUC`,
		number:    `931`,
	},
	`ANG`: {
		countries: Countries{`CURAÇAO`, `SINT MAARTEN (DUTCH PART)`, `SINT MAARTEN`},
		currency:  `Netherlands Antillean Guilder`,
		code:      `ANG`,
		number:    `532`,
	},
	`CZK`: {
		countries: Countries{`CZECHIA`},
		currency:  `Czech Koruna`,
		code:      `CZK`,
		number:    `203`,
	},
	`DKK`: {
		countries: Countries{`DENMARK`, `FAROE ISLANDS (THE)`, `GREENLAND`, `FAROE ISLANDS`},
		currency:  `Danish Krone`,
		code:      `DKK`,
		number:    `208`,
	},
	`DJF`: {
		countries: Countries{`DJIBOUTI`},
		currency:  `Djibouti Franc`,
		code:      `DJF`,
		number:    `262`,
	},
	`DOP`: {
		countries: Countries{`DOMINICAN REPUBLIC (THE)`, `DOMINICAN REPUBLIC`},
		currency:  `Dominican Peso`,
		code:      `DOP`,
		number:    `214`,
	},
	`EGP`: {
		countries: Countries{`EGYPT`},
		currency:  `Egyptian Pound`,
		code:      `EGP`,
		number:    `818`,
	},
	`SVC`: {
		countries: Countries{`EL SALVADOR`},
		currency:  `El Salvador Colon`,
		code:      `SVC`,
		number:    `222`,
	},
	`ERN`: {
		countries: Countries{`ERITREA`},
		currency:  `Nakfa`,
		code:      `ERN`,
		number:    `232`,
	},
	`SZL`: {
		countries: Countries{`ESWATINI`},
		currency:  `Lilangeni`,
		code:      `SZL`,
		number:    `748`,
	},
	`ETB`: {
		countries: Countries{`ETHIOPIA`},
		currency:  `Ethiopian Birr`,
		code:      `ETB`,
		number:    `230`,
	},
	`FKP`: {
		countries: Countries{`FALKLAND ISLANDS (THE) [MALVINAS]`, `FALKLAND ISLANDS [MALVINAS]`},
		currency:  `Falkland Islands Pound`,
		code:      `FKP`,
		number:    `238`,
	},
	`FJD`: {
		countries: Countries{`FIJI`},
		currency:  `Fiji Dollar`,
		code:      `FJD`,
		number:    `242`,
	},
	`XPF`: {
		countries: Countries{`FRENCH POLYNESIA`, `NEW CALEDONIA`, `WALLIS AND FUTUNA`},
		currency:  `CFP Franc`,
		code:      `XPF`,
		number:    `953`,
	},
	`GMD`: {
		countries: Countries{`GAMBIA (THE)`, `GAMBIA`},
		currency:  `Dalasi`,
		code:      `GMD`,
		number:    `270`,
	},
	`GEL`: {
		countries: Countries{`GEORGIA`},
		currency:  `Lari`,
		code:      `GEL`,
		number:    `981`,
	},
	`GHS`: {
		countries: Countries{`GHANA`},
		currency:  `Ghana Cedi`,
		code:      `GHS`,
		number:    `936`,
	},
	`GIP`: {
		countries: Countries{`GIBRALTAR`},
		currency:  `Gibraltar Pound`,
		code:      `GIP`,
		number:    `292`,
	},
	`GTQ`: {
		countries: Countries{`GUATEMALA`},
		currency:  `Quetzal`,
		code:      `GTQ`,
		number:    `320`,
	},
	`GBP`: {
		countries: Countries{`GUERNSEY`, `ISLE OF MAN`, `JERSEY`, `UNITED KINGDOM OF GREAT BRITAIN AND NORTHERN IRELAND (THE)`, `UNITED KINGDOM OF GREAT BRITAIN AND NORTHERN IRELAND`},
		currency:  `Pound Sterling`,
		code:      `GBP`,
		number:    `826`,
	},
	`GNF`: {
		countries: Countries{`GUINEA`},
		currency:  `Guinean Franc`,
		code:      `GNF`,
		number:    `324`,
	},
	`GYD`: {
		countries: Countries{`GUYANA`},
		currency:  `Guyana Dollar`,
		code:      `GYD`,
		number:    `328`,
	},
	`HTG`: {
		countries: Countries{`HAITI`},
		currency:  `Gourde`,
		code:      `HTG`,
		number:    `332`,
	},
	`HNL`: {
		countries: Countries{`HONDURAS`},
		currency:  `Lempira`,
		code:      `HNL`,
		number:    `340`,
	},
	`HKD`: {
		countries: Countries{`HONG KONG`},
		currency:  `Hong Kong Dollar`,
		code:      `HKD`,
		number:    `344`,
	},
	`HUF`: {
		countries: Countries{`HUNGARY`},
		currency:  `Forint`,
		code:      `HUF`,
		number:    `348`,
	},
	`ISK`: {
		countries: Countries{`ICELAND`},
		currency:  `Iceland Krona`,
		code:      `ISK`,
		number:    `352`,
	},
	`IDR`: {
		countries: Countries{`INDONESIA`},
		currency:  `Rupiah`,
		code:      `IDR`,
		number:    `360`,
	},
	`XDR`: {
		countries: Countries{`INTERNATIONAL MONETARY FUND (IMF) `, `INTERNATIONAL MONETARY FUND `},
		currency:  `SDR (Special Drawing Right)`,
		code:      `XDR`,
		number:    `960`,
	},
	`IRR`: {
		countries: Countries{`IRAN (ISLAMIC REPUBLIC OF)`, `IRAN`},
		currency:  `Iranian Rial`,
		code:      `IRR`,
		number:    `364`,
	},
	`IQD`: {
		countries: Countries{`IRAQ`},
		currency:  `Iraqi Dinar`,
		code:      `IQD`,
		number:    `368`,
	},
	`ILS`: {
		countries: Countries{`ISRAEL`},
		currency:  `New Israeli Sheqel`,
		code:      `ILS`,
		number:    `376`,
	},
	`JMD`: {
		countries: Countries{`JAMAICA`},
		currency:  `Jamaican Dollar`,
		code:      `JMD`,
		number:    `388`,
	},
	`JPY`: {
		countries: Countries{`JAPAN`},
		currency:  `Yen`,
		code:      `JPY`,
		number:    `392`,
	},
	`JOD`: {
		countries: Countries{`JORDAN`},
		currency:  `Jordanian Dinar`,
		code:      `JOD`,
		number:    `400`,
	},
	`KZT`: {
		countries: Countries{`KAZAKHSTAN`},
		currency:  `Tenge`,
		code:      `KZT`,
		number:    `398`,
	},
	`KES`: {
		countries: Countries{`KENYA`},
		currency:  `Kenyan Shilling`,
		code:      `KES`,
		number:    `404`,
	},
	`KPW`: {
		countries: Countries{`KOREA (THE DEMOCRATIC PEOPLE’S REPUBLIC OF)`, `KOREA`},
		currency:  `North Korean Won`,
		code:      `KPW`,
		number:    `408`,
	},
	`KRW`: {
		countries: Countries{`KOREA (THE REPUBLIC OF)`, `KOREA`},
		currency:  `Won`,
		code:      `KRW`,
		number:    `410`,
	},
	`KWD`: {
		countries: Countries{`KUWAIT`},
		currency:  `Kuwaiti Dinar`,
		code:      `KWD`,
		number:    `414`,
	},
	`KGS`: {
		countries: Countries{`KYRGYZSTAN`},
		currency:  `Som`,
		code:      `KGS`,
		number:    `417`,
	},
	`LAK`: {
		countries: Countries{`LAO PEOPLE’S DEMOCRATIC REPUBLIC (THE)`, `LAO PEOPLE’S DEMOCRATIC REPUBLIC`},
		currency:  `Lao Kip`,
		code:      `LAK`,
		number:    `418`,
	},
	`LBP`: {
		countries: Countries{`LEBANON`},
		currency:  `Lebanese Pound`,
		code:      `LBP`,
		number:    `422`,
	},
	`LSL`: {
		countries: Countries{`LESOTHO`},
		currency:  `Loti`,
		code:      `LSL`,
		number:    `426`,
	},
	`ZAR`: {
		countries: Countries{`LESOTHO`, `NAMIBIA`, `SOUTH AFRICA`},
		currency:  `Rand`,
		code:      `ZAR`,
		number:    `710`,
	},
	`LRD`: {
		countries: Countries{`LIBERIA`},
		currency:  `Liberian Dollar`,
		code:      `LRD`,
		number:    `430`,
	},
	`LYD`: {
		countries: Countries{`LIBYA`},
		currency:  `Libyan Dinar`,
		code:      `LYD`,
		number:    `434`,
	},
	`CHF`: {
		countries: Countries{`LIECHTENSTEIN`, `SWITZERLAND`},
		currency:  `Swiss Franc`,
		code:      `CHF`,
		number:    `756`,
	},
	`MOP`: {
		countries: Countries{`MACAO`},
		currency:  `Pataca`,
		code:      `MOP`,
		number:    `446`,
	},
	`MKD`: {
		countries: Countries{`NORTH MACEDONIA`},
		currency:  `Denar`,
		code:      `MKD`,
		number:    `807`,
	},
	`MGA`: {
		countries: Countries{`MADAGASCAR`},
		currency:  `Malagasy Ariary`,
		code:      `MGA`,
		number:    `969`,
	},
	`MWK`: {
		countries: Countries{`MALAWI`},
		currency:  `Malawi Kwacha`,
		code:      `MWK`,
		number:    `454`,
	},
	`MYR`: {
		countries: Countries{`MALAYSIA`},
		currency:  `Malaysian Ringgit`,
		code:      `MYR`,
		number:    `458`,
	},
	`MVR`: {
		countries: Countries{`MALDIVES`},
		currency:  `Rufiyaa`,
		code:      `MVR`,
		number:    `462`,
	},
	`MRU`: {
		countries: Countries{`MAURITANIA`},
		currency:  `Ouguiya`,
		code:      `MRU`,
		number:    `929`,
	},
	`MUR`: {
		countries: Countries{`MAURITIUS`},
		currency:  `Mauritius Rupee`,
		code:      `MUR`,
		number:    `480`,
	},
	`XUA`: {
		countries: Countries{`MEMBER COUNTRIES OF THE AFRICAN DEVELOPMENT BANK GROUP`},
		currency:  `ADB Unit of Account`,
		code:      `XUA`,
		number:    `965`,
	},
	`MXN`: {
		countries: Countries{`MEXICO`},
		currency:  `Mexican Peso`,
		code:      `MXN`,
		number:    `484`,
	},
	`MXV`: {
		countries: Countries{`MEXICO`},
		currency:  `Mexican Unidad de Inversion (UDI)`,
		code:      `MXV`,
		number:    `979`,
	},
	`MDL`: {
		countries: Countries{`MOLDOVA (THE REPUBLIC OF)`, `MOLDOVA`},
		currency:  `Moldovan Leu`,
		code:      `MDL`,
		number:    `498`,
	},
	`MNT`: {
		countries: Countries{`MONGOLIA`},
		currency:  `Tugrik`,
		code:      `MNT`,
		number:    `496`,
	},
	`MAD`: {
		countries: Countries{`MOROCCO`, `WESTERN SAHARA`},
		currency:  `Moroccan Dirham`,
		code:      `MAD`,
		number:    `504`,
	},
	`MZN`: {
		countries: Countries{`MOZAMBIQUE`},
		currency:  `Mozambique Metical`,
		code:      `MZN`,
		number:    `943`,
	},
	`MMK`: {
		countries: Countries{`MYANMAR`},
		currency:  `Kyat`,
		code:      `MMK`,
		number:    `104`,
	},
	`NAD`: {
		countries: Countries{`NAMIBIA`},
		currency:  `Namibia Dollar`,
		code:      `NAD`,
		number:    `516`,
	},
	`NPR`: {
		countries: Countries{`NEPAL`},
		currency:  `Nepalese Rupee`,
		code:      `NPR`,
		number:    `524`,
	},
	`NIO`: {
		countries: Countries{`NICARAGUA`},
		currency:  `Cordoba Oro`,
		code:      `NIO`,
		number:    `558`,
	},
	`NGN`: {
		countries: Countries{`NIGERIA`},
		currency:  `Naira`,
		code:      `NGN`,
		number:    `566`,
	},
	`OMR`: {
		countries: Countries{`OMAN`},
		currency:  `Rial Omani`,
		code:      `OMR`,
		number:    `512`,
	},
	`PKR`: {
		countries: Countries{`PAKISTAN`},
		currency:  `Pakistan Rupee`,
		code:      `PKR`,
		number:    `586`,
	},
	`PAB`: {
		countries: Countries{`PANAMA`},
		currency:  `Balboa`,
		code:      `PAB`,
		number:    `590`,
	},
	`PGK`: {
		countries: Countries{`PAPUA NEW GUINEA`},
		currency:  `Kina`,
		code:      `PGK`,
		number:    `598`,
	},
	`PYG`: {
		countries: Countries{`PARAGUAY`},
		currency:  `Guarani`,
		code:      `PYG`,
		number:    `600`,
	},
	`PEN`: {
		countries: Countries{`PERU`},
		currency:  `Sol`,
		code:      `PEN`,
		number:    `604`,
	},
	`PHP`: {
		countries: Countries{`PHILIPPINES (THE)`, `PHILIPPINES`},
		currency:  `Philippine Peso`,
		code:      `PHP`,
		number:    `608`,
	},
	`PLN`: {
		countries: Countries{`POLAND`},
		currency:  `Zloty`,
		code:      `PLN`,
		number:    `985`,
	},
	`QAR`: {
		countries: Countries{`QATAR`},
		currency:  `Qatari Rial`,
		code:      `QAR`,
		number:    `634`,
	},
	`RON`: {
		countries: Countries{`ROMANIA`},
		currency:  `Romanian Leu`,
		code:      `RON`,
		number:    `946`,
	},
	`RUB`: {
		countries: Countries{`RUSSIAN FEDERATION (THE)`, `RUSSIAN FEDERATION`},
		currency:  `Russian Ruble`,
		code:      `RUB`,
		number:    `643`,
	},
	`RWF`: {
		countries: Countries{`RWANDA`},
		currency:  `Rwanda Franc`,
		code:      `RWF`,
		number:    `646`,
	},
	`SHP`: {
		countries: Countries{`SAINT HELENA, ASCENSION AND TRISTAN DA CUNHA`},
		currency:  `Saint Helena Pound`,
		code:      `SHP`,
		number:    `654`,
	},
	`WST`: {
		countries: Countries{`SAMOA`},
		currency:  `Tala`,
		code:      `WST`,
		number:    `882`,
	},
	`STN`: {
		countries: Countries{`SAO TOME AND PRINCIPE`},
		currency:  `Dobra`,
		code:      `STN`,
		number:    `930`,
	},
	`SAR`: {
		countries: Countries{`SAUDI ARABIA`},
		currency:  `Saudi Riyal`,
		code:      `SAR`,
		number:    `682`,
	},
	`RSD`: {
		countries: Countries{`SERBIA`},
		currency:  `Serbian Dinar`,
		code:      `RSD`,
		number:    `941`,
	},
	`SCR`: {
		countries: Countries{`SEYCHELLES`},
		currency:  `Seychelles Rupee`,
		code:      `SCR`,
		number:    `690`,
	},
	`SLL`: {
		countries: Countries{`SIERRA LEONE`},
		currency:  `Leone`,
		code:      `SLL`,
		number:    `694`,
	},
	`SGD`: {
		countries: Countries{`SINGAPORE`},
		currency:  `Singapore Dollar`,
		code:      `SGD`,
		number:    `702`,
	},
	`XSU`: {
		countries: Countries{`SISTEMA UNITARIO DE COMPENSACION REGIONAL DE PAGOS "SUCRE"`},
		currency:  `Sucre`,
		code:      `XSU`,
		number:    `994`,
	},
	`SBD`: {
		countries: Countries{`SOLOMON ISLANDS`},
		currency:  `Solomon Islands Dollar`,
		code:      `SBD`,
		number:    `090`,
	},
	`SOS`: {
		countries: Countries{`SOMALIA`},
		currency:  `Somali Shilling`,
		code:      `SOS`,
		number:    `706`,
	},
	`SSP`: {
		countries: Countries{`SOUTH SUDAN`},
		currency:  `South Sudanese Pound`,
		code:      `SSP`,
		number:    `728`,
	},
	`LKR`: {
		countries: Countries{`SRI LANKA`},
		currency:  `Sri Lanka Rupee`,
		code:      `LKR`,
		number:    `144`,
	},
	`SDG`: {
		countries: Countries{`SUDAN (THE)`, `SUDAN`},
		currency:  `Sudanese Pound`,
		code:      `SDG`,
		number:    `938`,
	},
	`SRD`: {
		countries: Countries{`SURINAME`},
		currency:  `Surinam Dollar`,
		code:      `SRD`,
		number:    `968`,
	},
	`SEK`: {
		countries: Countries{`SWEDEN`},
		currency:  `Swedish Krona`,
		code:      `SEK`,
		number:    `752`,
	},
	`CHE`: {
		countries: Countries{`SWITZERLAND`},
		currency:  `WIR Euro`,
		code:      `CHE`,
		number:    `947`,
	},
	`CHW`: {
		countries: Countries{`SWITZERLAND`},
		currency:  `WIR Franc`,
		code:      `CHW`,
		number:    `948`,
	},
	`SYP`: {
		countries: Countries{`SYRIAN ARAB REPUBLIC`},
		currency:  `Syrian Pound`,
		code:      `SYP`,
		number:    `760`,
	},
	`TWD`: {
		countries: Countries{`TAIWAN (PROVINCE OF CHINA)`, `TAIWAN`},
		currency:  `New Taiwan Dollar`,
		code:      `TWD`,
		number:    `901`,
	},
	`TJS`: {
		countries: Countries{`TAJIKISTAN`},
		currency:  `Somoni`,
		code:      `TJS`,
		number:    `972`,
	},
	`TZS`: {
		countries: Countries{`TANZANIA, UNITED REPUBLIC OF`},
		currency:  `Tanzanian Shilling`,
		code:      `TZS`,
		number:    `834`,
	},
	`THB`: {
		countries: Countries{`THAILAND`},
		currency:  `Baht`,
		code:      `THB`,
		number:    `764`,
	},
	`TOP`: {
		countries: Countries{`TONGA`},
		currency:  `Pa’anga`,
		code:      `TOP`,
		number:    `776`,
	},
	`TTD`: {
		countries: Countries{`TRINIDAD AND TOBAGO`},
		currency:  `Trinidad and Tobago Dollar`,
		code:      `TTD`,
		number:    `780`,
	},
	`TND`: {
		countries: Countries{`TUNISIA`},
		currency:  `Tunisian Dinar`,
		code:      `TND`,
		number:    `788`,
	},
	`TRY`: {
		countries: Countries{`TURKEY`},
		currency:  `Turkish Lira`,
		code:      `TRY`,
		number:    `949`,
	},
	`TMT`: {
		countries: Countries{`TURKMENISTAN`},
		currency:  `Turkmenistan New Manat`,
		code:      `TMT`,
		number:    `934`,
	},
	`UGX`: {
		countries: Countries{`UGANDA`},
		currency:  `Uganda Shilling`,
		code:      `UGX`,
		number:    `800`,
	},
	`UAH`: {
		countries: Countries{`UKRAINE`},
		currency:  `Hryvnia`,
		code:      `UAH`,
		number:    `980`,
	},
	`AED`: {
		countries: Countries{`UNITED ARAB EMIRATES (THE)`, `UNITED ARAB EMIRATES`},
		currency:  `UAE Dirham`,
		code:      `AED`,
		number:    `784`,
	},
	`USN`: {
		countries: Countries{`UNITED STATES OF AMERICA (THE)`, `UNITED STATES OF AMERICA`},
		currency:  `US Dollar (Next day)`,
		code:      `USN`,
		number:    `997`,
	},
	`UYU`: {
		countries: Countries{`URUGUAY`},
		currency:  `Peso Uruguayo`,
		code:      `UYU`,
		number:    `858`,
	},
	`UYI`: {
		countries: Countries{`URUGUAY`},
		currency:  `Uruguay Peso en Unidades Indexadas (UI)`,
		code:      `UYI`,
		number:    `940`,
	},
	`UYW`: {
		countries: Countries{`URUGUAY`},
		currency:  `Unidad Previsional`,
		code:      `UYW`,
		number:    `927`,
	},
	`UZS`: {
		countries: Countries{`UZBEKISTAN`},
		currency:  `Uzbekistan Sum`,
		code:      `UZS`,
		number:    `860`,
	},
	`VUV`: {
		countries: Countries{`VANUATU`},
		currency:  `Vatu`,
		code:      `VUV`,
		number:    `548`,
	},
	`VES`: {
		countries: Countries{`VENEZUELA (BOLIVARIAN REPUBLIC OF)`, `VENEZUELA`},
		currency:  `Bolívar Soberano`,
		code:      `VES`,
		number:    `928`,
	},
	`VND`: {
		countries: Countries{`VIET NAM`},
		currency:  `Dong`,
		code:      `VND`,
		number:    `704`,
	},
	`YER`: {
		countries: Countries{`YEMEN`},
		currency:  `Yemeni Rial`,
		code:      `YER`,
		number:    `886`,
	},
	`ZMW`: {
		countries: Countries{`ZAMBIA`},
		currency:  `Zambian Kwacha`,
		code:      `ZMW`,
		number:    `967`,
	},
	`ZWL`: {
		countries: Countries{`ZIMBABWE`},
		currency:  `Zimbabwe Dollar`,
		code:      `ZWL`,
		number:    `932`,
	},
}

var currenciesByNumber = map[string]currency{
	`104`: {
		countries: Countries{`MYANMAR`},
		currency:  `Kyat`,
		code:      `MMK`,
		number:    `104`,
	},
	`108`: {
		countries: Countries{`BURUNDI`},
		currency:  `Burundi Franc`,
		code:      `BIF`,
		number:    `108`,
	},
	`116`: {
		countries: Countries{`CAMBODIA`},
		currency:  `Riel`,
		code:      `KHR`,
		number:    `116`,
	},
	`124`: {
		countries: Countries{`CANADA`},
		currency:  `Canadian Dollar`,
		code:      `CAD`,
		number:    `124`,
	},
	`132`: {
		countries: Countries{`CABO VERDE`},
		currency:  `Cabo Verde Escudo`,
		code:      `CVE`,
		number:    `132`,
	},
	`136`: {
		countries: Countries{`CAYMAN ISLANDS (THE)`, `CAYMAN ISLANDS`},
		currency:  `Cayman Islands Dollar`,
		code:      `KYD`,
		number:    `136`,
	},
	`144`: {
		countries: Countries{`SRI LANKA`},
		currency:  `Sri Lanka Rupee`,
		code:      `LKR`,
		number:    `144`,
	},
	`152`: {
		countries: Countries{`CHILE`},
		currency:  `Chilean Peso`,
		code:      `CLP`,
		number:    `152`,
	},
	`156`: {
		countries: Countries{`CHINA`},
		currency:  `Yuan Renminbi`,
		code:      `CNY`,
		number:    `156`,
	},
	`170`: {
		countries: Countries{`COLOMBIA`},
		currency:  `Colombian Peso`,
		code:      `COP`,
		number:    `170`,
	},
	`174`: {
		countries: Countries{`COMOROS (THE)`, `COMOROS`},
		currency:  `Comorian Franc `,
		code:      `KMF`,
		number:    `174`,
	},
	`188`: {
		countries: Countries{`COSTA RICA`},
		currency:  `Costa Rican Colon`,
		code:      `CRC`,
		number:    `188`,
	},
	`191`: {
		countries: Countries{`CROATIA`},
		currency:  `Kuna`,
		code:      `HRK`,
		number:    `191`,
	},
	`192`: {
		countries: Countries{`CUBA`},
		currency:  `Cuban Peso`,
		code:      `CUP`,
		number:    `192`,
	},
	`203`: {
		countries: Countries{`CZECHIA`},
		currency:  `Czech Koruna`,
		code:      `CZK`,
		number:    `203`,
	},
	`208`: {
		countries: Countries{`DENMARK`, `FAROE ISLANDS (THE)`, `GREENLAND`, `FAROE ISLANDS`},
		currency:  `Danish Krone`,
		code:      `DKK`,
		number:    `208`,
	},
	`214`: {
		countries: Countries{`DOMINICAN REPUBLIC (THE)`, `DOMINICAN REPUBLIC`},
		currency:  `Dominican Peso`,
		code:      `DOP`,
		number:    `214`,
	},
	`222`: {
		countries: Countries{`EL SALVADOR`},
		currency:  `El Salvador Colon`,
		code:      `SVC`,
		number:    `222`,
	},
	`230`: {
		countries: Countries{`ETHIOPIA`},
		currency:  `Ethiopian Birr`,
		code:      `ETB`,
		number:    `230`,
	},
	`232`: {
		countries: Countries{`ERITREA`},
		currency:  `Nakfa`,
		code:      `ERN`,
		number:    `232`,
	},
	`238`: {
		countries: Countries{`FALKLAND ISLANDS (THE) [MALVINAS]`, `FALKLAND ISLANDS [MALVINAS]`},
		currency:  `Falkland Islands Pound`,
		code:      `FKP`,
		number:    `238`,
	},
	`242`: {
		countries: Countries{`FIJI`},
		currency:  `Fiji Dollar`,
		code:      `FJD`,
		number:    `242`,
	},
	`262`: {
		countries: Countries{`DJIBOUTI`},
		currency:  `Djibouti Franc`,
		code:      `DJF`,
		number:    `262`,
	},
	`270`: {
		countries: Countries{`GAMBIA (THE)`, `GAMBIA`},
		currency:  `Dalasi`,
		code:      `GMD`,
		number:    `270`,
	},
	`292`: {
		countries: Countries{`GIBRALTAR`},
		currency:  `Gibraltar Pound`,
		code:      `GIP`,
		number:    `292`,
	},
	`320`: {
		countries: Countries{`GUATEMALA`},
		currency:  `Quetzal`,
		code:      `GTQ`,
		number:    `320`,
	},
	`324`: {
		countries: Countries{`GUINEA`},
		currency:  `Guinean Franc`,
		code:      `GNF`,
		number:    `324`,
	},
	`328`: {
		countries: Countries{`GUYANA`},
		currency:  `Guyana Dollar`,
		code:      `GYD`,
		number:    `328`,
	},
	`332`: {
		countries: Countries{`HAITI`},
		currency:  `Gourde`,
		code:      `HTG`,
		number:    `332`,
	},
	`340`: {
		countries: Countries{`HONDURAS`},
		currency:  `Lempira`,
		code:      `HNL`,
		number:    `340`,
	},
	`344`: {
		countries: Countries{`HONG KONG`},
		currency:  `Hong Kong Dollar`,
		code:      `HKD`,
		number:    `344`,
	},
	`348`: {
		countries: Countries{`HUNGARY`},
		currency:  `Forint`,
		code:      `HUF`,
		number:    `348`,
	},
	`352`: {
		countries: Countries{`ICELAND`},
		currency:  `Iceland Krona`,
		code:      `ISK`,
		number:    `352`,
	},
	`356`: {
		countries: Countries{`BHUTAN`, `INDIA`},
		currency:  `Indian Rupee`,
		code:      `INR`,
		number:    `356`,
	},
	`360`: {
		countries: Countries{`INDONESIA`},
		currency:  `Rupiah`,
		code:      `IDR`,
		number:    `360`,
	},
	`364`: {
		countries: Countries{`IRAN (ISLAMIC REPUBLIC OF)`, `IRAN`},
		currency:  `Iranian Rial`,
		code:      `IRR`,
		number:    `364`,
	},
	`368`: {
		countries: Countries{`IRAQ`},
		currency:  `Iraqi Dinar`,
		code:      `IQD`,
		number:    `368`,
	},
	`376`: {
		countries: Countries{`ISRAEL`},
		currency:  `New Israeli Sheqel`,
		code:      `ILS`,
		number:    `376`,
	},
	`388`: {
		countries: Countries{`JAMAICA`},
		currency:  `Jamaican Dollar`,
		code:      `JMD`,
		number:    `388`,
	},
	`392`: {
		countries: Countries{`JAPAN`},
		currency:  `Yen`,
		code:      `JPY`,
		number:    `392`,
	},
	`398`: {
		countries: Countries{`KAZAKHSTAN`},
		currency:  `Tenge`,
		code:      `KZT`,
		number:    `398`,
	},
	`400`: {
		countries: Countries{`JORDAN`},
		currency:  `Jordanian Dinar`,
		code:      `JOD`,
		number:    `400`,
	},
	`404`: {
		countries: Countries{`KENYA`},
		currency:  `Kenyan Shilling`,
		code:      `KES`,
		number:    `404`,
	},
	`408`: {
		countries: Countries{`KOREA (THE DEMOCRATIC PEOPLE’S REPUBLIC OF)`, `KOREA`},
		currency:  `North Korean Won`,
		code:      `KPW`,
		number:    `408`,
	},
	`410`: {
		countries: Countries{`KOREA (THE REPUBLIC OF)`, `KOREA`},
		currency:  `Won`,
		code:      `KRW`,
		number:    `410`,
	},
	`414`: {
		countries: Countries{`KUWAIT`},
		currency:  `Kuwaiti Dinar`,
		code:      `KWD`,
		number:    `414`,
	},
	`417`: {
		countries: Countries{`KYRGYZSTAN`},
		currency:  `Som`,
		code:      `KGS`,
		number:    `417`,
	},
	`418`: {
		countries: Countries{`LAO PEOPLE’S DEMOCRATIC REPUBLIC (THE)`, `LAO PEOPLE’S DEMOCRATIC REPUBLIC`},
		currency:  `Lao Kip`,
		code:      `LAK`,
		number:    `418`,
	},
	`422`: {
		countries: Countries{`LEBANON`},
		currency:  `Lebanese Pound`,
		code:      `LBP`,
		number:    `422`,
	},
	`426`: {
		countries: Countries{`LESOTHO`},
		currency:  `Loti`,
		code:      `LSL`,
		number:    `426`,
	},
	`430`: {
		countries: Countries{`LIBERIA`},
		currency:  `Liberian Dollar`,
		code:      `LRD`,
		number:    `430`,
	},
	`434`: {
		countries: Countries{`LIBYA`},
		currency:  `Libyan Dinar`,
		code:      `LYD`,
		number:    `434`,
	},
	`446`: {
		countries: Countries{`MACAO`},
		currency:  `Pataca`,
		code:      `MOP`,
		number:    `446`,
	},
	`454`: {
		countries: Countries{`MALAWI`},
		currency:  `Malawi Kwacha`,
		code:      `MWK`,
		number:    `454`,
	},
	`458`: {
		countries: Countries{`MALAYSIA`},
		currency:  `Malaysian Ringgit`,
		code:      `MYR`,
		number:    `458`,
	},
	`462`: {
		countries: Countries{`MALDIVES`},
		currency:  `Rufiyaa`,
		code:      `MVR`,
		number:    `462`,
	},
	`480`: {
		countries: Countries{`MAURITIUS`},
		currency:  `Mauritius Rupee`,
		code:      `MUR`,
		number:    `480`,
	},
	`484`: {
		countries: Countries{`MEXICO`},
		currency:  `Mexican Peso`,
		code:      `MXN`,
		number:    `484`,
	},
	`496`: {
		countries: Countries{`MONGOLIA`},
		currency:  `Tugrik`,
		code:      `MNT`,
		number:    `496`,
	},
	`498`: {
		countries: Countries{`MOLDOVA (THE REPUBLIC OF)`, `MOLDOVA`},
		currency:  `Moldovan Leu`,
		code:      `MDL`,
		number:    `498`,
	},
	`504`: {
		countries: Countries{`MOROCCO`, `WESTERN SAHARA`},
		currency:  `Moroccan Dirham`,
		code:      `MAD`,
		number:    `504`,
	},
	`512`: {
		countries: Countries{`OMAN`},
		currency:  `Rial Omani`,
		code:      `OMR`,
		number:    `512`,
	},
	`516`: {
		countries: Countries{`NAMIBIA`},
		currency:  `Namibia Dollar`,
		code:      `NAD`,
		number:    `516`,
	},
	`524`: {
		countries: Countries{`NEPAL`},
		currency:  `Nepalese Rupee`,
		code:      `NPR`,
		number:    `524`,
	},
	`532`: {
		countries: Countries{`CURAÇAO`, `SINT MAARTEN (DUTCH PART)`, `SINT MAARTEN`},
		currency:  `Netherlands Antillean Guilder`,
		code:      `ANG`,
		number:    `532`,
	},
	`533`: {
		countries: Countries{`ARUBA`},
		currency:  `Aruban Florin`,
		code:      `AWG`,
		number:    `533`,
	},
	`548`: {
		countries: Countries{`VANUATU`},
		currency:  `Vatu`,
		code:      `VUV`,
		number:    `548`,
	},
	`554`: {
		countries: Countries{`COOK ISLANDS (THE)`, `NEW ZEALAND`, `NIUE`, `PITCAIRN`, `TOKELAU`, `COOK ISLANDS`},
		currency:  `New Zealand Dollar`,
		code:      `NZD`,
		number:    `554`,
	},
	`558`: {
		countries: Countries{`NICARAGUA`},
		currency:  `Cordoba Oro`,
		code:      `NIO`,
		number:    `558`,
	},
	`566`: {
		countries: Countries{`NIGERIA`},
		currency:  `Naira`,
		code:      `NGN`,
		number:    `566`,
	},
	`578`: {
		countries: Countries{`BOUVET ISLAND`, `NORWAY`, `SVALBARD AND JAN MAYEN`},
		currency:  `Norwegian Krone`,
		code:      `NOK`,
		number:    `578`,
	},
	`586`: {
		countries: Countries{`PAKISTAN`},
		currency:  `Pakistan Rupee`,
		code:      `PKR`,
		number:    `586`,
	},
	`590`: {
		countries: Countries{`PANAMA`},
		currency:  `Balboa`,
		code:      `PAB`,
		number:    `590`,
	},
	`598`: {
		countries: Countries{`PAPUA NEW GUINEA`},
		currency:  `Kina`,
		code:      `PGK`,
		number:    `598`,
	},
	`600`: {
		countries: Countries{`PARAGUAY`},
		currency:  `Guarani`,
		code:      `PYG`,
		number:    `600`,
	},
	`604`: {
		countries: Countries{`PERU`},
		currency:  `Sol`,
		code:      `PEN`,
		number:    `604`,
	},
	`608`: {
		countries: Countries{`PHILIPPINES (THE)`, `PHILIPPINES`},
		currency:  `Philippine Peso`,
		code:      `PHP`,
		number:    `608`,
	},
	`634`: {
		countries: Countries{`QATAR`},
		currency:  `Qatari Rial`,
		code:      `QAR`,
		number:    `634`,
	},
	`643`: {
		countries: Countries{`RUSSIAN FEDERATION (THE)`, `RUSSIAN FEDERATION`},
		currency:  `Russian Ruble`,
		code:      `RUB`,
		number:    `643`,
	},
	`646`: {
		countries: Countries{`RWANDA`},
		currency:  `Rwanda Franc`,
		code:      `RWF`,
		number:    `646`,
	},
	`654`: {
		countries: Countries{`SAINT HELENA, ASCENSION AND TRISTAN DA CUNHA`},
		currency:  `Saint Helena Pound`,
		code:      `SHP`,
		number:    `654`,
	},
	`682`: {
		countries: Countries{`SAUDI ARABIA`},
		currency:  `Saudi Riyal`,
		code:      `SAR`,
		number:    `682`,
	},
	`690`: {
		countries: Countries{`SEYCHELLES`},
		currency:  `Seychelles Rupee`,
		code:      `SCR`,
		number:    `690`,
	},
	`694`: {
		countries: Countries{`SIERRA LEONE`},
		currency:  `Leone`,
		code:      `SLL`,
		number:    `694`,
	},
	`702`: {
		countries: Countries{`SINGAPORE`},
		currency:  `Singapore Dollar`,
		code:      `SGD`,
		number:    `702`,
	},
	`704`: {
		countries: Countries{`VIET NAM`},
		currency:  `Dong`,
		code:      `VND`,
		number:    `704`,
	},
	`706`: {
		countries: Countries{`SOMALIA`},
		currency:  `Somali Shilling`,
		code:      `SOS`,
		number:    `706`,
	},
	`710`: {
		countries: Countries{`LESOTHO`, `NAMIBIA`, `SOUTH AFRICA`},
		currency:  `Rand`,
		code:      `ZAR`,
		number:    `710`,
	},
	`728`: {
		countries: Countries{`SOUTH SUDAN`},
		currency:  `South Sudanese Pound`,
		code:      `SSP`,
		number:    `728`,
	},
	`748`: {
		countries: Countries{`ESWATINI`},
		currency:  `Lilangeni`,
		code:      `SZL`,
		number:    `748`,
	},
	`752`: {
		countries: Countries{`SWEDEN`},
		currency:  `Swedish Krona`,
		code:      `SEK`,
		number:    `752`,
	},
	`756`: {
		countries: Countries{`LIECHTENSTEIN`, `SWITZERLAND`},
		currency:  `Swiss Franc`,
		code:      `CHF`,
		number:    `756`,
	},
	`760`: {
		countries: Countries{`SYRIAN ARAB REPUBLIC`},
		currency:  `Syrian Pound`,
		code:      `SYP`,
		number:    `760`,
	},
	`764`: {
		countries: Countries{`THAILAND`},
		currency:  `Baht`,
		code:      `THB`,
		number:    `764`,
	},
	`776`: {
		countries: Countries{`TONGA`},
		currency:  `Pa’anga`,
		code:      `TOP`,
		number:    `776`,
	},
	`780`: {
		countries: Countries{`TRINIDAD AND TOBAGO`},
		currency:  `Trinidad and Tobago Dollar`,
		code:      `TTD`,
		number:    `780`,
	},
	`784`: {
		countries: Countries{`UNITED ARAB EMIRATES (THE)`, `UNITED ARAB EMIRATES`},
		currency:  `UAE Dirham`,
		code:      `AED`,
		number:    `784`,
	},
	`788`: {
		countries: Countries{`TUNISIA`},
		currency:  `Tunisian Dinar`,
		code:      `TND`,
		number:    `788`,
	},
	`800`: {
		countries: Countries{`UGANDA`},
		currency:  `Uganda Shilling`,
		code:      `UGX`,
		number:    `800`,
	},
	`807`: {
		countries: Countries{`NORTH MACEDONIA`},
		currency:  `Denar`,
		code:      `MKD`,
		number:    `807`,
	},
	`818`: {
		countries: Countries{`EGYPT`},
		currency:  `Egyptian Pound`,
		code:      `EGP`,
		number:    `818`,
	},
	`826`: {
		countries: Countries{`GUERNSEY`, `ISLE OF MAN`, `JERSEY`, `UNITED KINGDOM OF GREAT BRITAIN AND NORTHERN IRELAND (THE)`, `UNITED KINGDOM OF GREAT BRITAIN AND NORTHERN IRELAND`},
		currency:  `Pound Sterling`,
		code:      `GBP`,
		number:    `826`,
	},
	`834`: {
		countries: Countries{`TANZANIA, UNITED REPUBLIC OF`},
		currency:  `Tanzanian Shilling`,
		code:      `TZS`,
		number:    `834`,
	},
	`840`: {
		countries: Countries{`AMERICAN SAMOA`, `BONAIRE, SINT EUSTATIUS AND SABA`, `BRITISH INDIAN OCEAN TERRITORY (THE)`, `ECUADOR`, `EL SALVADOR`, `GUAM`, `HAITI`, `MARSHALL ISLANDS (THE)`, `MICRONESIA (FEDERATED STATES OF)`, `NORTHERN MARIANA ISLANDS (THE)`, `PALAU`, `PANAMA`, `PUERTO RICO`, `TIMOR-LESTE`, `TURKS AND CAICOS ISLANDS (THE)`, `UNITED STATES MINOR OUTLYING ISLANDS (THE)`, `UNITED STATES OF AMERICA (THE)`, `VIRGIN ISLANDS (BRITISH)`, `VIRGIN ISLANDS (U.S.)`, `BRITISH INDIAN OCEAN TERRITORY`, `MARSHALL ISLANDS`, `MICRONESIA`, `NORTHERN MARIANA ISLANDS`, `TURKS AND CAICOS ISLANDS`, `UNITED STATES MINOR OUTLYING ISLANDS`, `UNITED STATES OF AMERICA`, `VIRGIN ISLANDS`, `VIRGIN ISLANDS`},
		currency:  `US Dollar`,
		code:      `USD`,
		number:    `840`,
	},
	`858`: {
		countries: Countries{`URUGUAY`},
		currency:  `Peso Uruguayo`,
		code:      `UYU`,
		number:    `858`,
	},
	`860`: {
		countries: Countries{`UZBEKISTAN`},
		currency:  `Uzbekistan Sum`,
		code:      `UZS`,
		number:    `860`,
	},
	`882`: {
		countries: Countries{`SAMOA`},
		currency:  `Tala`,
		code:      `WST`,
		number:    `882`,
	},
	`886`: {
		countries: Countries{`YEMEN`},
		currency:  `Yemeni Rial`,
		code:      `YER`,
		number:    `886`,
	},
	`901`: {
		countries: Countries{`TAIWAN (PROVINCE OF CHINA)`, `TAIWAN`},
		currency:  `New Taiwan Dollar`,
		code:      `TWD`,
		number:    `901`,
	},
	`927`: {
		countries: Countries{`URUGUAY`},
		currency:  `Unidad Previsional`,
		code:      `UYW`,
		number:    `927`,
	},
	`928`: {
		countries: Countries{`VENEZUELA (BOLIVARIAN REPUBLIC OF)`, `VENEZUELA`},
		currency:  `Bolívar Soberano`,
		code:      `VES`,
		number:    `928`,
	},
	`929`: {
		countries: Countries{`MAURITANIA`},
		currency:  `Ouguiya`,
		code:      `MRU`,
		number:    `929`,
	},
	`930`: {
		countries: Countries{`SAO TOME AND PRINCIPE`},
		currency:  `Dobra`,
		code:      `STN`,
		number:    `930`,
	},
	`931`: {
		countries: Countries{`CUBA`},
		currency:  `Peso Convertible`,
		code:      `CUC`,
		number:    `931`,
	},
	`932`: {
		countries: Countries{`ZIMBABWE`},
		currency:  `Zimbabwe Dollar`,
		code:      `ZWL`,
		number:    `932`,
	},
	`933`: {
		countries: Countries{`BELARUS`},
		currency:  `Belarusian Ruble`,
		code:      `BYN`,
		number:    `933`,
	},
	`934`: {
		countries: Countries{`TURKMENISTAN`},
		currency:  `Turkmenistan New Manat`,
		code:      `TMT`,
		number:    `934`,
	},
	`936`: {
		countries: Countries{`GHANA`},
		currency:  `Ghana Cedi`,
		code:      `GHS`,
		number:    `936`,
	},
	`938`: {
		countries: Countries{`SUDAN (THE)`, `SUDAN`},
		currency:  `Sudanese Pound`,
		code:      `SDG`,
		number:    `938`,
	},
	`940`: {
		countries: Countries{`URUGUAY`},
		currency:  `Uruguay Peso en Unidades Indexadas (UI)`,
		code:      `UYI`,
		number:    `940`,
	},
	`941`: {
		countries: Countries{`SERBIA`},
		currency:  `Serbian Dinar`,
		code:      `RSD`,
		number:    `941`,
	},
	`943`: {
		countries: Countries{`MOZAMBIQUE`},
		currency:  `Mozambique Metical`,
		code:      `MZN`,
		number:    `943`,
	},
	`944`: {
		countries: Countries{`AZERBAIJAN`},
		currency:  `Azerbaijan Manat`,
		code:      `AZN`,
		number:    `944`,
	},
	`946`: {
		countries: Countries{`ROMANIA`},
		currency:  `Romanian Leu`,
		code:      `RON`,
		number:    `946`,
	},
	`947`: {
		countries: Countries{`SWITZERLAND`},
		currency:  `WIR Euro`,
		code:      `CHE`,
		number:    `947`,
	},
	`948`: {
		countries: Countries{`SWITZERLAND`},
		currency:  `WIR Franc`,
		code:      `CHW`,
		number:    `948`,
	},
	`949`: {
		countries: Countries{`TURKEY`},
		currency:  `Turkish Lira`,
		code:      `TRY`,
		number:    `949`,
	},
	`950`: {
		countries: Countries{`CAMEROON`, `CENTRAL AFRICAN REPUBLIC (THE)`, `CHAD`, `CONGO (THE)`, `EQUATORIAL GUINEA`, `GABON`, `CENTRAL AFRICAN REPUBLIC`, `CONGO`},
		currency:  `CFA Franc BEAC`,
		code:      `XAF`,
		number:    `950`,
	},
	`951`: {
		countries: Countries{`ANGUILLA`, `ANTIGUA AND BARBUDA`, `DOMINICA`, `GRENADA`, `MONTSERRAT`, `SAINT KITTS AND NEVIS`, `SAINT LUCIA`, `SAINT VINCENT AND THE GRENADINES`},
		currency:  `East Caribbean Dollar`,
		code:      `XCD`,
		number:    `951`,
	},
	`952`: {
		countries: Countries{`BENIN`, `BURKINA FASO`, `CÔTE D'IVOIRE`, `GUINEA-BISSAU`, `MALI`, `NIGER (THE)`, `SENEGAL`, `TOGO`, `NIGER`},
		currency:  `CFA Franc BCEAO`,
		code:      `XOF`,
		number:    `952`,
	},
	`953`: {
		countries: Countries{`FRENCH POLYNESIA`, `NEW CALEDONIA`, `WALLIS AND FUTUNA`},
		currency:  `CFP Franc`,
		code:      `XPF`,
		number:    `953`,
	},
	`960`: {
		countries: Countries{`INTERNATIONAL MONETARY FUND (IMF) `, `INTERNATIONAL MONETARY FUND `},
		currency:  `SDR (Special Drawing Right)`,
		code:      `XDR`,
		number:    `960`,
	},
	`965`: {
		countries: Countries{`MEMBER COUNTRIES OF THE AFRICAN DEVELOPMENT BANK GROUP`},
		currency:  `ADB Unit of Account`,
		code:      `XUA`,
		number:    `965`,
	},
	`967`: {
		countries: Countries{`ZAMBIA`},
		currency:  `Zambian Kwacha`,
		code:      `ZMW`,
		number:    `967`,
	},
	`968`: {
		countries: Countries{`SURINAME`},
		currency:  `Surinam Dollar`,
		code:      `SRD`,
		number:    `968`,
	},
	`969`: {
		countries: Countries{`MADAGASCAR`},
		currency:  `Malagasy Ariary`,
		code:      `MGA`,
		number:    `969`,
	},
	`970`: {
		countries: Countries{`COLOMBIA`},
		currency:  `Unidad de Valor Real`,
		code:      `COU`,
		number:    `970`,
	},
	`971`: {
		countries: Countries{`AFGHANISTAN`},
		currency:  `Afghani`,
		code:      `AFN`,
		number:    `971`,
	},
	`972`: {
		countries: Countries{`TAJIKISTAN`},
		currency:  `Somoni`,
		code:      `TJS`,
		number:    `972`,
	},
	`973`: {
		countries: Countries{`ANGOLA`},
		currency:  `Kwanza`,
		code:      `AOA`,
		number:    `973`,
	},
	`975`: {
		countries: Countries{`BULGARIA`},
		currency:  `Bulgarian Lev`,
		code:      `BGN`,
		number:    `975`,
	},
	`976`: {
		countries: Countries{`CONGO (THE DEMOCRATIC REPUBLIC OF THE)`, `CONGO`},
		currency:  `Congolese Franc`,
		code:      `CDF`,
		number:    `976`,
	},
	`977`: {
		countries: Countries{`BOSNIA AND HERZEGOVINA`},
		currency:  `Convertible Mark`,
		code:      `BAM`,
		number:    `977`,
	},
	`978`: {
		countries: Countries{`ÅLAND ISLANDS`, `ANDORRA`, `AUSTRIA`, `BELGIUM`, `CYPRUS`, `ESTONIA`, `EUROPEAN UNION`, `FINLAND`, `FRANCE`, `FRENCH GUIANA`, `FRENCH SOUTHERN TERRITORIES (THE)`, `GERMANY`, `GREECE`, `GUADELOUPE`, `HOLY SEE (THE)`, `IRELAND`, `ITALY`, `LATVIA`, `LITHUANIA`, `LUXEMBOURG`, `MALTA`, `MARTINIQUE`, `MAYOTTE`, `MONACO`, `MONTENEGRO`, `NETHERLANDS (THE)`, `PORTUGAL`, `RÉUNION`, `SAINT BARTHÉLEMY`, `SAINT MARTIN (FRENCH PART)`, `SAINT PIERRE AND MIQUELON`, `SAN MARINO`, `SLOVAKIA`, `SLOVENIA`, `SPAIN`, `FRENCH SOUTHERN TERRITORIES`, `HOLY SEE`, `NETHERLANDS`, `SAINT MARTIN`},
		currency:  `Euro`,
		code:      `EUR`,
		number:    `978`,
	},
	`979`: {
		countries: Countries{`MEXICO`},
		currency:  `Mexican Unidad de Inversion (UDI)`,
		code:      `MXV`,
		number:    `979`,
	},
	`980`: {
		countries: Countries{`UKRAINE`},
		currency:  `Hryvnia`,
		code:      `UAH`,
		number:    `980`,
	},
	`981`: {
		countries: Countries{`GEORGIA`},
		currency:  `Lari`,
		code:      `GEL`,
		number:    `981`,
	},
	`984`: {
		countries: Countries{`BOLIVIA (PLURINATIONAL STATE OF)`, `BOLIVIA`},
		currency:  `Mvdol`,
		code:      `BOV`,
		number:    `984`,
	},
	`985`: {
		countries: Countries{`POLAND`},
		currency:  `Zloty`,
		code:      `PLN`,
		number:    `985`,
	},
	`986`: {
		countries: Countries{`BRAZIL`},
		currency:  `Brazilian Real`,
		code:      `BRL`,
		number:    `986`,
	},
	`990`: {
		countries: Countries{`CHILE`},
		currency:  `Unidad de Fomento`,
		code:      `CLF`,
		number:    `990`,
	},
	`994`: {
		countries: Countries{`SISTEMA UNITARIO DE COMPENSACION REGIONAL DE PAGOS "SUCRE"`},
		currency:  `Sucre`,
		code:      `XSU`,
		number:    `994`,
	},
	`997`: {
		countries: Countries{`UNITED STATES OF AMERICA (THE)`, `UNITED STATES OF AMERICA`},
		currency:  `US Dollar (Next day)`,
		code:      `USN`,
		number:    `997`,
	},
	`008`: {
		countries: Countries{`ALBANIA`},
		currency:  `Lek`,
		code:      `ALL`,
		number:    `008`,
	},
	`012`: {
		countries: Countries{`ALGERIA`},
		currency:  `Algerian Dinar`,
		code:      `DZD`,
		number:    `012`,
	},
	`032`: {
		countries: Countries{`ARGENTINA`},
		currency:  `Argentine Peso`,
		code:      `ARS`,
		number:    `032`,
	},
	`051`: {
		countries: Countries{`ARMENIA`},
		currency:  `Armenian Dram`,
		code:      `AMD`,
		number:    `051`,
	},
	`036`: {
		countries: Countries{`AUSTRALIA`, `CHRISTMAS ISLAND`, `COCOS (KEELING) ISLANDS (THE)`, `HEARD ISLAND AND McDONALD ISLANDS`, `KIRIBATI`, `NAURU`, `NORFOLK ISLAND`, `TUVALU`, `COCOS ISLANDS`},
		currency:  `Australian Dollar`,
		code:      `AUD`,
		number:    `036`,
	},
	`044`: {
		countries: Countries{`BAHAMAS (THE)`, `BAHAMAS`},
		currency:  `Bahamian Dollar`,
		code:      `BSD`,
		number:    `044`,
	},
	`048`: {
		countries: Countries{`BAHRAIN`},
		currency:  `Bahraini Dinar`,
		code:      `BHD`,
		number:    `048`,
	},
	`050`: {
		countries: Countries{`BANGLADESH`},
		currency:  `Taka`,
		code:      `BDT`,
		number:    `050`,
	},
	`052`: {
		countries: Countries{`BARBADOS`},
		currency:  `Barbados Dollar`,
		code:      `BBD`,
		number:    `052`,
	},
	`084`: {
		countries: Countries{`BELIZE`},
		currency:  `Belize Dollar`,
		code:      `BZD`,
		number:    `084`,
	},
	`060`: {
		countries: Countries{`BERMUDA`},
		currency:  `Bermudian Dollar`,
		code:      `BMD`,
		number:    `060`,
	},
	`064`: {
		countries: Countries{`BHUTAN`},
		currency:  `Ngultrum`,
		code:      `BTN`,
		number:    `064`,
	},
	`068`: {
		countries: Countries{`BOLIVIA (PLURINATIONAL STATE OF)`, `BOLIVIA`},
		currency:  `Boliviano`,
		code:      `BOB`,
		number:    `068`,
	},
	`072`: {
		countries: Countries{`BOTSWANA`},
		currency:  `Pula`,
		code:      `BWP`,
		number:    `072`,
	},
	`096`: {
		countries: Countries{`BRUNEI DARUSSALAM`},
		currency:  `Brunei Dollar`,
		code:      `BND`,
		number:    `096`,
	},
	`090`: {
		countries: Countries{`SOLOMON ISLANDS`},
		currency:  `Solomon Islands Dollar`,
		code:      `SBD`,
		number:    `090`,
	},
}

var currenciesByCountry = map[string]currencies{
	`AFGHANISTAN`: {
		{
			countries: Countries{`AFGHANISTAN`},
			currency:  `Afghani`,
			code:      `AFN`,
			number:    `971`,
		},
	}, `ÅLAND ISLANDS`: {
		{
			countries: Countries{`ÅLAND ISLANDS`, `ANDORRA`, `AUSTRIA`, `BELGIUM`, `CYPRUS`, `ESTONIA`, `EUROPEAN UNION`, `FINLAND`, `FRANCE`, `FRENCH GUIANA`, `FRENCH SOUTHERN TERRITORIES (THE)`, `GERMANY`, `GREECE`, `GUADELOUPE`, `HOLY SEE (THE)`, `IRELAND`, `ITALY`, `LATVIA`, `LITHUANIA`, `LUXEMBOURG`, `MALTA`, `MARTINIQUE`, `MAYOTTE`, `MONACO`, `MONTENEGRO`, `NETHERLANDS (THE)`, `PORTUGAL`, `RÉUNION`, `SAINT BARTHÉLEMY`, `SAINT MARTIN (FRENCH PART)`, `SAINT PIERRE AND MIQUELON`, `SAN MARINO`, `SLOVAKIA`, `SLOVENIA`, `SPAIN`, `FRENCH SOUTHERN TERRITORIES`, `HOLY SEE`, `NETHERLANDS`, `SAINT MARTIN`},
			currency:  `Euro`,
			code:      `EUR`,
			number:    `978`,
		},
	}, `ALBANIA`: {
		{
			countries: Countries{`ALBANIA`},
			currency:  `Lek`,
			code:      `ALL`,
			number:    `008`,
		},
	}, `ALGERIA`: {
		{
			countries: Countries{`ALGERIA`},
			currency:  `Algerian Dinar`,
			code:      `DZD`,
			number:    `012`,
		},
	}, `AMERICAN SAMOA`: {
		{
			countries: Countries{`AMERICAN SAMOA`, `BONAIRE, SINT EUSTATIUS AND SABA`, `BRITISH INDIAN OCEAN TERRITORY (THE)`, `ECUADOR`, `EL SALVADOR`, `GUAM`, `HAITI`, `MARSHALL ISLANDS (THE)`, `MICRONESIA (FEDERATED STATES OF)`, `NORTHERN MARIANA ISLANDS (THE)`, `PALAU`, `PANAMA`, `PUERTO RICO`, `TIMOR-LESTE`, `TURKS AND CAICOS ISLANDS (THE)`, `UNITED STATES MINOR OUTLYING ISLANDS (THE)`, `UNITED STATES OF AMERICA (THE)`, `VIRGIN ISLANDS (BRITISH)`, `VIRGIN ISLANDS (U.S.)`, `BRITISH INDIAN OCEAN TERRITORY`, `MARSHALL ISLANDS`, `MICRONESIA`, `NORTHERN MARIANA ISLANDS`, `TURKS AND CAICOS ISLANDS`, `UNITED STATES MINOR OUTLYING ISLANDS`, `UNITED STATES OF AMERICA`, `VIRGIN ISLANDS`, `VIRGIN ISLANDS`},
			currency:  `US Dollar`,
			code:      `USD`,
			number:    `840`,
		},
	}, `ANDORRA`: {
		{
			countries: Countries{`ÅLAND ISLANDS`, `ANDORRA`, `AUSTRIA`, `BELGIUM`, `CYPRUS`, `ESTONIA`, `EUROPEAN UNION`, `FINLAND`, `FRANCE`, `FRENCH GUIANA`, `FRENCH SOUTHERN TERRITORIES (THE)`, `GERMANY`, `GREECE`, `GUADELOUPE`, `HOLY SEE (THE)`, `IRELAND`, `ITALY`, `LATVIA`, `LITHUANIA`, `LUXEMBOURG`, `MALTA`, `MARTINIQUE`, `MAYOTTE`, `MONACO`, `MONTENEGRO`, `NETHERLANDS (THE)`, `PORTUGAL`, `RÉUNION`, `SAINT BARTHÉLEMY`, `SAINT MARTIN (FRENCH PART)`, `SAINT PIERRE AND MIQUELON`, `SAN MARINO`, `SLOVAKIA`, `SLOVENIA`, `SPAIN`, `FRENCH SOUTHERN TERRITORIES`, `HOLY SEE`, `NETHERLANDS`, `SAINT MARTIN`},
			currency:  `Euro`,
			code:      `EUR`,
			number:    `978`,
		},
	}, `ANGOLA`: {
		{
			countries: Countries{`ANGOLA`},
			currency:  `Kwanza`,
			code:      `AOA`,
			number:    `973`,
		},
	}, `ANGUILLA`: {
		{
			countries: Countries{`ANGUILLA`, `ANTIGUA AND BARBUDA`, `DOMINICA`, `GRENADA`, `MONTSERRAT`, `SAINT KITTS AND NEVIS`, `SAINT LUCIA`, `SAINT VINCENT AND THE GRENADINES`},
			currency:  `East Caribbean Dollar`,
			code:      `XCD`,
			number:    `951`,
		},
	}, `ANTIGUA AND BARBUDA`: {
		{
			countries: Countries{`ANGUILLA`, `ANTIGUA AND BARBUDA`, `DOMINICA`, `GRENADA`, `MONTSERRAT`, `SAINT KITTS AND NEVIS`, `SAINT LUCIA`, `SAINT VINCENT AND THE GRENADINES`},
			currency:  `East Caribbean Dollar`,
			code:      `XCD`,
			number:    `951`,
		},
	}, `ARGENTINA`: {
		{
			countries: Countries{`ARGENTINA`},
			currency:  `Argentine Peso`,
			code:      `ARS`,
			number:    `032`,
		},
	}, `ARMENIA`: {
		{
			countries: Countries{`ARMENIA`},
			currency:  `Armenian Dram`,
			code:      `AMD`,
			number:    `051`,
		},
	}, `ARUBA`: {
		{
			countries: Countries{`ARUBA`},
			currency:  `Aruban Florin`,
			code:      `AWG`,
			number:    `533`,
		},
	}, `AUSTRALIA`: {
		{
			countries: Countries{`AUSTRALIA`, `CHRISTMAS ISLAND`, `COCOS (KEELING) ISLANDS (THE)`, `HEARD ISLAND AND McDONALD ISLANDS`, `KIRIBATI`, `NAURU`, `NORFOLK ISLAND`, `TUVALU`, `COCOS ISLANDS`},
			currency:  `Australian Dollar`,
			code:      `AUD`,
			number:    `036`,
		},
	}, `AUSTRIA`: {
		{
			countries: Countries{`ÅLAND ISLANDS`, `ANDORRA`, `AUSTRIA`, `BELGIUM`, `CYPRUS`, `ESTONIA`, `EUROPEAN UNION`, `FINLAND`, `FRANCE`, `FRENCH GUIANA`, `FRENCH SOUTHERN TERRITORIES (THE)`, `GERMANY`, `GREECE`, `GUADELOUPE`, `HOLY SEE (THE)`, `IRELAND`, `ITALY`, `LATVIA`, `LITHUANIA`, `LUXEMBOURG`, `MALTA`, `MARTINIQUE`, `MAYOTTE`, `MONACO`, `MONTENEGRO`, `NETHERLANDS (THE)`, `PORTUGAL`, `RÉUNION`, `SAINT BARTHÉLEMY`, `SAINT MARTIN (FRENCH PART)`, `SAINT PIERRE AND MIQUELON`, `SAN MARINO`, `SLOVAKIA`, `SLOVENIA`, `SPAIN`, `FRENCH SOUTHERN TERRITORIES`, `HOLY SEE`, `NETHERLANDS`, `SAINT MARTIN`},
			currency:  `Euro`,
			code:      `EUR`,
			number:    `978`,
		},
	}, `AZERBAIJAN`: {
		{
			countries: Countries{`AZERBAIJAN`},
			currency:  `Azerbaijan Manat`,
			code:      `AZN`,
			number:    `944`,
		},
	}, `BAHAMAS (THE)`: {
		{
			countries: Countries{`BAHAMAS (THE)`, `BAHAMAS`},
			currency:  `Bahamian Dollar`,
			code:      `BSD`,
			number:    `044`,
		},
	}, `BAHRAIN`: {
		{
			countries: Countries{`BAHRAIN`},
			currency:  `Bahraini Dinar`,
			code:      `BHD`,
			number:    `048`,
		},
	}, `BANGLADESH`: {
		{
			countries: Countries{`BANGLADESH`},
			currency:  `Taka`,
			code:      `BDT`,
			number:    `050`,
		},
	}, `BARBADOS`: {
		{
			countries: Countries{`BARBADOS`},
			currency:  `Barbados Dollar`,
			code:      `BBD`,
			number:    `052`,
		},
	}, `BELARUS`: {
		{
			countries: Countries{`BELARUS`},
			currency:  `Belarusian Ruble`,
			code:      `BYN`,
			number:    `933`,
		},
	}, `BELGIUM`: {
		{
			countries: Countries{`ÅLAND ISLANDS`, `ANDORRA`, `AUSTRIA`, `BELGIUM`, `CYPRUS`, `ESTONIA`, `EUROPEAN UNION`, `FINLAND`, `FRANCE`, `FRENCH GUIANA`, `FRENCH SOUTHERN TERRITORIES (THE)`, `GERMANY`, `GREECE`, `GUADELOUPE`, `HOLY SEE (THE)`, `IRELAND`, `ITALY`, `LATVIA`, `LITHUANIA`, `LUXEMBOURG`, `MALTA`, `MARTINIQUE`, `MAYOTTE`, `MONACO`, `MONTENEGRO`, `NETHERLANDS (THE)`, `PORTUGAL`, `RÉUNION`, `SAINT BARTHÉLEMY`, `SAINT MARTIN (FRENCH PART)`, `SAINT PIERRE AND MIQUELON`, `SAN MARINO`, `SLOVAKIA`, `SLOVENIA`, `SPAIN`, `FRENCH SOUTHERN TERRITORIES`, `HOLY SEE`, `NETHERLANDS`, `SAINT MARTIN`},
			currency:  `Euro`,
			code:      `EUR`,
			number:    `978`,
		},
	}, `BELIZE`: {
		{
			countries: Countries{`BELIZE`},
			currency:  `Belize Dollar`,
			code:      `BZD`,
			number:    `084`,
		},
	}, `BENIN`: {
		{
			countries: Countries{`BENIN`, `BURKINA FASO`, `CÔTE D'IVOIRE`, `GUINEA-BISSAU`, `MALI`, `NIGER (THE)`, `SENEGAL`, `TOGO`, `NIGER`},
			currency:  `CFA Franc BCEAO`,
			code:      `XOF`,
			number:    `952`,
		},
	}, `BERMUDA`: {
		{
			countries: Countries{`BERMUDA`},
			currency:  `Bermudian Dollar`,
			code:      `BMD`,
			number:    `060`,
		},
	}, `BHUTAN`: {
		{
			countries: Countries{`BHUTAN`, `INDIA`},
			currency:  `Indian Rupee`,
			code:      `INR`,
			number:    `356`,
		}, {
			countries: Countries{`BHUTAN`},
			currency:  `Ngultrum`,
			code:      `BTN`,
			number:    `064`,
		},
	}, `BOLIVIA (PLURINATIONAL STATE OF)`: {
		{
			countries: Countries{`BOLIVIA (PLURINATIONAL STATE OF)`, `BOLIVIA`},
			currency:  `Boliviano`,
			code:      `BOB`,
			number:    `068`,
		}, {
			countries: Countries{`BOLIVIA (PLURINATIONAL STATE OF)`, `BOLIVIA`},
			currency:  `Mvdol`,
			code:      `BOV`,
			number:    `984`,
		},
	}, `BONAIRE, SINT EUSTATIUS AND SABA`: {
		{
			countries: Countries{`AMERICAN SAMOA`, `BONAIRE, SINT EUSTATIUS AND SABA`, `BRITISH INDIAN OCEAN TERRITORY (THE)`, `ECUADOR`, `EL SALVADOR`, `GUAM`, `HAITI`, `MARSHALL ISLANDS (THE)`, `MICRONESIA (FEDERATED STATES OF)`, `NORTHERN MARIANA ISLANDS (THE)`, `PALAU`, `PANAMA`, `PUERTO RICO`, `TIMOR-LESTE`, `TURKS AND CAICOS ISLANDS (THE)`, `UNITED STATES MINOR OUTLYING ISLANDS (THE)`, `UNITED STATES OF AMERICA (THE)`, `VIRGIN ISLANDS (BRITISH)`, `VIRGIN ISLANDS (U.S.)`, `BRITISH INDIAN OCEAN TERRITORY`, `MARSHALL ISLANDS`, `MICRONESIA`, `NORTHERN MARIANA ISLANDS`, `TURKS AND CAICOS ISLANDS`, `UNITED STATES MINOR OUTLYING ISLANDS`, `UNITED STATES OF AMERICA`, `VIRGIN ISLANDS`, `VIRGIN ISLANDS`},
			currency:  `US Dollar`,
			code:      `USD`,
			number:    `840`,
		},
	}, `BOSNIA AND HERZEGOVINA`: {
		{
			countries: Countries{`BOSNIA AND HERZEGOVINA`},
			currency:  `Convertible Mark`,
			code:      `BAM`,
			number:    `977`,
		},
	}, `BOTSWANA`: {
		{
			countries: Countries{`BOTSWANA`},
			currency:  `Pula`,
			code:      `BWP`,
			number:    `072`,
		},
	}, `BOUVET ISLAND`: {
		{
			countries: Countries{`BOUVET ISLAND`, `NORWAY`, `SVALBARD AND JAN MAYEN`},
			currency:  `Norwegian Krone`,
			code:      `NOK`,
			number:    `578`,
		},
	}, `BRAZIL`: {
		{
			countries: Countries{`BRAZIL`},
			currency:  `Brazilian Real`,
			code:      `BRL`,
			number:    `986`,
		},
	}, `BRITISH INDIAN OCEAN TERRITORY (THE)`: {
		{
			countries: Countries{`AMERICAN SAMOA`, `BONAIRE, SINT EUSTATIUS AND SABA`, `BRITISH INDIAN OCEAN TERRITORY (THE)`, `ECUADOR`, `EL SALVADOR`, `GUAM`, `HAITI`, `MARSHALL ISLANDS (THE)`, `MICRONESIA (FEDERATED STATES OF)`, `NORTHERN MARIANA ISLANDS (THE)`, `PALAU`, `PANAMA`, `PUERTO RICO`, `TIMOR-LESTE`, `TURKS AND CAICOS ISLANDS (THE)`, `UNITED STATES MINOR OUTLYING ISLANDS (THE)`, `UNITED STATES OF AMERICA (THE)`, `VIRGIN ISLANDS (BRITISH)`, `VIRGIN ISLANDS (U.S.)`, `BRITISH INDIAN OCEAN TERRITORY`, `MARSHALL ISLANDS`, `MICRONESIA`, `NORTHERN MARIANA ISLANDS`, `TURKS AND CAICOS ISLANDS`, `UNITED STATES MINOR OUTLYING ISLANDS`, `UNITED STATES OF AMERICA`, `VIRGIN ISLANDS`, `VIRGIN ISLANDS`},
			currency:  `US Dollar`,
			code:      `USD`,
			number:    `840`,
		},
	}, `BRUNEI DARUSSALAM`: {
		{
			countries: Countries{`BRUNEI DARUSSALAM`},
			currency:  `Brunei Dollar`,
			code:      `BND`,
			number:    `096`,
		},
	}, `BULGARIA`: {
		{
			countries: Countries{`BULGARIA`},
			currency:  `Bulgarian Lev`,
			code:      `BGN`,
			number:    `975`,
		},
	}, `BURKINA FASO`: {
		{
			countries: Countries{`BENIN`, `BURKINA FASO`, `CÔTE D'IVOIRE`, `GUINEA-BISSAU`, `MALI`, `NIGER (THE)`, `SENEGAL`, `TOGO`, `NIGER`},
			currency:  `CFA Franc BCEAO`,
			code:      `XOF`,
			number:    `952`,
		},
	}, `BURUNDI`: {
		{
			countries: Countries{`BURUNDI`},
			currency:  `Burundi Franc`,
			code:      `BIF`,
			number:    `108`,
		},
	}, `CABO VERDE`: {
		{
			countries: Countries{`CABO VERDE`},
			currency:  `Cabo Verde Escudo`,
			code:      `CVE`,
			number:    `132`,
		},
	}, `CAMBODIA`: {
		{
			countries: Countries{`CAMBODIA`},
			currency:  `Riel`,
			code:      `KHR`,
			number:    `116`,
		},
	}, `CAMEROON`: {
		{
			countries: Countries{`CAMEROON`, `CENTRAL AFRICAN REPUBLIC (THE)`, `CHAD`, `CONGO (THE)`, `EQUATORIAL GUINEA`, `GABON`, `CENTRAL AFRICAN REPUBLIC`, `CONGO`},
			currency:  `CFA Franc BEAC`,
			code:      `XAF`,
			number:    `950`,
		},
	}, `CANADA`: {
		{
			countries: Countries{`CANADA`},
			currency:  `Canadian Dollar`,
			code:      `CAD`,
			number:    `124`,
		},
	}, `CAYMAN ISLANDS (THE)`: {
		{
			countries: Countries{`CAYMAN ISLANDS (THE)`, `CAYMAN ISLANDS`},
			currency:  `Cayman Islands Dollar`,
			code:      `KYD`,
			number:    `136`,
		},
	}, `CENTRAL AFRICAN REPUBLIC (THE)`: {
		{
			countries: Countries{`CAMEROON`, `CENTRAL AFRICAN REPUBLIC (THE)`, `CHAD`, `CONGO (THE)`, `EQUATORIAL GUINEA`, `GABON`, `CENTRAL AFRICAN REPUBLIC`, `CONGO`},
			currency:  `CFA Franc BEAC`,
			code:      `XAF`,
			number:    `950`,
		},
	}, `CHAD`: {
		{
			countries: Countries{`CAMEROON`, `CENTRAL AFRICAN REPUBLIC (THE)`, `CHAD`, `CONGO (THE)`, `EQUATORIAL GUINEA`, `GABON`, `CENTRAL AFRICAN REPUBLIC`, `CONGO`},
			currency:  `CFA Franc BEAC`,
			code:      `XAF`,
			number:    `950`,
		},
	}, `CHILE`: {
		{
			countries: Countries{`CHILE`},
			currency:  `Chilean Peso`,
			code:      `CLP`,
			number:    `152`,
		}, {
			countries: Countries{`CHILE`},
			currency:  `Unidad de Fomento`,
			code:      `CLF`,
			number:    `990`,
		},
	}, `CHINA`: {
		{
			countries: Countries{`CHINA`},
			currency:  `Yuan Renminbi`,
			code:      `CNY`,
			number:    `156`,
		},
	}, `CHRISTMAS ISLAND`: {
		{
			countries: Countries{`AUSTRALIA`, `CHRISTMAS ISLAND`, `COCOS (KEELING) ISLANDS (THE)`, `HEARD ISLAND AND McDONALD ISLANDS`, `KIRIBATI`, `NAURU`, `NORFOLK ISLAND`, `TUVALU`, `COCOS ISLANDS`},
			currency:  `Australian Dollar`,
			code:      `AUD`,
			number:    `036`,
		},
	}, `COCOS (KEELING) ISLANDS (THE)`: {
		{
			countries: Countries{`AUSTRALIA`, `CHRISTMAS ISLAND`, `COCOS (KEELING) ISLANDS (THE)`, `HEARD ISLAND AND McDONALD ISLANDS`, `KIRIBATI`, `NAURU`, `NORFOLK ISLAND`, `TUVALU`, `COCOS ISLANDS`},
			currency:  `Australian Dollar`,
			code:      `AUD`,
			number:    `036`,
		},
	}, `COLOMBIA`: {
		{
			countries: Countries{`COLOMBIA`},
			currency:  `Colombian Peso`,
			code:      `COP`,
			number:    `170`,
		}, {
			countries: Countries{`COLOMBIA`},
			currency:  `Unidad de Valor Real`,
			code:      `COU`,
			number:    `970`,
		},
	}, `COMOROS (THE)`: {
		{
			countries: Countries{`COMOROS (THE)`, `COMOROS`},
			currency:  `Comorian Franc `,
			code:      `KMF`,
			number:    `174`,
		},
	}, `CONGO (THE DEMOCRATIC REPUBLIC OF THE)`: {
		{
			countries: Countries{`CONGO (THE DEMOCRATIC REPUBLIC OF THE)`, `CONGO`},
			currency:  `Congolese Franc`,
			code:      `CDF`,
			number:    `976`,
		},
	}, `CONGO (THE)`: {
		{
			countries: Countries{`CAMEROON`, `CENTRAL AFRICAN REPUBLIC (THE)`, `CHAD`, `CONGO (THE)`, `EQUATORIAL GUINEA`, `GABON`, `CENTRAL AFRICAN REPUBLIC`, `CONGO`},
			currency:  `CFA Franc BEAC`,
			code:      `XAF`,
			number:    `950`,
		},
	}, `COOK ISLANDS (THE)`: {
		{
			countries: Countries{`COOK ISLANDS (THE)`, `NEW ZEALAND`, `NIUE`, `PITCAIRN`, `TOKELAU`, `COOK ISLANDS`},
			currency:  `New Zealand Dollar`,
			code:      `NZD`,
			number:    `554`,
		},
	}, `COSTA RICA`: {
		{
			countries: Countries{`COSTA RICA`},
			currency:  `Costa Rican Colon`,
			code:      `CRC`,
			number:    `188`,
		},
	}, `CÔTE D'IVOIRE`: {
		{
			countries: Countries{`BENIN`, `BURKINA FASO`, `CÔTE D'IVOIRE`, `GUINEA-BISSAU`, `MALI`, `NIGER (THE)`, `SENEGAL`, `TOGO`, `NIGER`},
			currency:  `CFA Franc BCEAO`,
			code:      `XOF`,
			number:    `952`,
		},
	}, `CROATIA`: {
		{
			countries: Countries{`CROATIA`},
			currency:  `Kuna`,
			code:      `HRK`,
			number:    `191`,
		},
	}, `CUBA`: {
		{
			countries: Countries{`CUBA`},
			currency:  `Cuban Peso`,
			code:      `CUP`,
			number:    `192`,
		}, {
			countries: Countries{`CUBA`},
			currency:  `Peso Convertible`,
			code:      `CUC`,
			number:    `931`,
		},
	}, `CURAÇAO`: {
		{
			countries: Countries{`CURAÇAO`, `SINT MAARTEN (DUTCH PART)`, `SINT MAARTEN`},
			currency:  `Netherlands Antillean Guilder`,
			code:      `ANG`,
			number:    `532`,
		},
	}, `CYPRUS`: {
		{
			countries: Countries{`ÅLAND ISLANDS`, `ANDORRA`, `AUSTRIA`, `BELGIUM`, `CYPRUS`, `ESTONIA`, `EUROPEAN UNION`, `FINLAND`, `FRANCE`, `FRENCH GUIANA`, `FRENCH SOUTHERN TERRITORIES (THE)`, `GERMANY`, `GREECE`, `GUADELOUPE`, `HOLY SEE (THE)`, `IRELAND`, `ITALY`, `LATVIA`, `LITHUANIA`, `LUXEMBOURG`, `MALTA`, `MARTINIQUE`, `MAYOTTE`, `MONACO`, `MONTENEGRO`, `NETHERLANDS (THE)`, `PORTUGAL`, `RÉUNION`, `SAINT BARTHÉLEMY`, `SAINT MARTIN (FRENCH PART)`, `SAINT PIERRE AND MIQUELON`, `SAN MARINO`, `SLOVAKIA`, `SLOVENIA`, `SPAIN`, `FRENCH SOUTHERN TERRITORIES`, `HOLY SEE`, `NETHERLANDS`, `SAINT MARTIN`},
			currency:  `Euro`,
			code:      `EUR`,
			number:    `978`,
		},
	}, `CZECHIA`: {
		{
			countries: Countries{`CZECHIA`},
			currency:  `Czech Koruna`,
			code:      `CZK`,
			number:    `203`,
		},
	}, `DENMARK`: {
		{
			countries: Countries{`DENMARK`, `FAROE ISLANDS (THE)`, `GREENLAND`, `FAROE ISLANDS`},
			currency:  `Danish Krone`,
			code:      `DKK`,
			number:    `208`,
		},
	}, `DJIBOUTI`: {
		{
			countries: Countries{`DJIBOUTI`},
			currency:  `Djibouti Franc`,
			code:      `DJF`,
			number:    `262`,
		},
	}, `DOMINICA`: {
		{
			countries: Countries{`ANGUILLA`, `ANTIGUA AND BARBUDA`, `DOMINICA`, `GRENADA`, `MONTSERRAT`, `SAINT KITTS AND NEVIS`, `SAINT LUCIA`, `SAINT VINCENT AND THE GRENADINES`},
			currency:  `East Caribbean Dollar`,
			code:      `XCD`,
			number:    `951`,
		},
	}, `DOMINICAN REPUBLIC (THE)`: {
		{
			countries: Countries{`DOMINICAN REPUBLIC (THE)`, `DOMINICAN REPUBLIC`},
			currency:  `Dominican Peso`,
			code:      `DOP`,
			number:    `214`,
		},
	}, `ECUADOR`: {
		{
			countries: Countries{`AMERICAN SAMOA`, `BONAIRE, SINT EUSTATIUS AND SABA`, `BRITISH INDIAN OCEAN TERRITORY (THE)`, `ECUADOR`, `EL SALVADOR`, `GUAM`, `HAITI`, `MARSHALL ISLANDS (THE)`, `MICRONESIA (FEDERATED STATES OF)`, `NORTHERN MARIANA ISLANDS (THE)`, `PALAU`, `PANAMA`, `PUERTO RICO`, `TIMOR-LESTE`, `TURKS AND CAICOS ISLANDS (THE)`, `UNITED STATES MINOR OUTLYING ISLANDS (THE)`, `UNITED STATES OF AMERICA (THE)`, `VIRGIN ISLANDS (BRITISH)`, `VIRGIN ISLANDS (U.S.)`, `BRITISH INDIAN OCEAN TERRITORY`, `MARSHALL ISLANDS`, `MICRONESIA`, `NORTHERN MARIANA ISLANDS`, `TURKS AND CAICOS ISLANDS`, `UNITED STATES MINOR OUTLYING ISLANDS`, `UNITED STATES OF AMERICA`, `VIRGIN ISLANDS`, `VIRGIN ISLANDS`},
			currency:  `US Dollar`,
			code:      `USD`,
			number:    `840`,
		},
	}, `EGYPT`: {
		{
			countries: Countries{`EGYPT`},
			currency:  `Egyptian Pound`,
			code:      `EGP`,
			number:    `818`,
		},
	}, `EL SALVADOR`: {
		{
			countries: Countries{`EL SALVADOR`},
			currency:  `El Salvador Colon`,
			code:      `SVC`,
			number:    `222`,
		}, {
			countries: Countries{`AMERICAN SAMOA`, `BONAIRE, SINT EUSTATIUS AND SABA`, `BRITISH INDIAN OCEAN TERRITORY (THE)`, `ECUADOR`, `EL SALVADOR`, `GUAM`, `HAITI`, `MARSHALL ISLANDS (THE)`, `MICRONESIA (FEDERATED STATES OF)`, `NORTHERN MARIANA ISLANDS (THE)`, `PALAU`, `PANAMA`, `PUERTO RICO`, `TIMOR-LESTE`, `TURKS AND CAICOS ISLANDS (THE)`, `UNITED STATES MINOR OUTLYING ISLANDS (THE)`, `UNITED STATES OF AMERICA (THE)`, `VIRGIN ISLANDS (BRITISH)`, `VIRGIN ISLANDS (U.S.)`, `BRITISH INDIAN OCEAN TERRITORY`, `MARSHALL ISLANDS`, `MICRONESIA`, `NORTHERN MARIANA ISLANDS`, `TURKS AND CAICOS ISLANDS`, `UNITED STATES MINOR OUTLYING ISLANDS`, `UNITED STATES OF AMERICA`, `VIRGIN ISLANDS`, `VIRGIN ISLANDS`},
			currency:  `US Dollar`,
			code:      `USD`,
			number:    `840`,
		},
	}, `EQUATORIAL GUINEA`: {
		{
			countries: Countries{`CAMEROON`, `CENTRAL AFRICAN REPUBLIC (THE)`, `CHAD`, `CONGO (THE)`, `EQUATORIAL GUINEA`, `GABON`, `CENTRAL AFRICAN REPUBLIC`, `CONGO`},
			currency:  `CFA Franc BEAC`,
			code:      `XAF`,
			number:    `950`,
		},
	}, `ERITREA`: {
		{
			countries: Countries{`ERITREA`},
			currency:  `Nakfa`,
			code:      `ERN`,
			number:    `232`,
		},
	}, `ESTONIA`: {
		{
			countries: Countries{`ÅLAND ISLANDS`, `ANDORRA`, `AUSTRIA`, `BELGIUM`, `CYPRUS`, `ESTONIA`, `EUROPEAN UNION`, `FINLAND`, `FRANCE`, `FRENCH GUIANA`, `FRENCH SOUTHERN TERRITORIES (THE)`, `GERMANY`, `GREECE`, `GUADELOUPE`, `HOLY SEE (THE)`, `IRELAND`, `ITALY`, `LATVIA`, `LITHUANIA`, `LUXEMBOURG`, `MALTA`, `MARTINIQUE`, `MAYOTTE`, `MONACO`, `MONTENEGRO`, `NETHERLANDS (THE)`, `PORTUGAL`, `RÉUNION`, `SAINT BARTHÉLEMY`, `SAINT MARTIN (FRENCH PART)`, `SAINT PIERRE AND MIQUELON`, `SAN MARINO`, `SLOVAKIA`, `SLOVENIA`, `SPAIN`, `FRENCH SOUTHERN TERRITORIES`, `HOLY SEE`, `NETHERLANDS`, `SAINT MARTIN`},
			currency:  `Euro`,
			code:      `EUR`,
			number:    `978`,
		},
	}, `ESWATINI`: {
		{
			countries: Countries{`ESWATINI`},
			currency:  `Lilangeni`,
			code:      `SZL`,
			number:    `748`,
		},
	}, `ETHIOPIA`: {
		{
			countries: Countries{`ETHIOPIA`},
			currency:  `Ethiopian Birr`,
			code:      `ETB`,
			number:    `230`,
		},
	}, `EUROPEAN UNION`: {
		{
			countries: Countries{`ÅLAND ISLANDS`, `ANDORRA`, `AUSTRIA`, `BELGIUM`, `CYPRUS`, `ESTONIA`, `EUROPEAN UNION`, `FINLAND`, `FRANCE`, `FRENCH GUIANA`, `FRENCH SOUTHERN TERRITORIES (THE)`, `GERMANY`, `GREECE`, `GUADELOUPE`, `HOLY SEE (THE)`, `IRELAND`, `ITALY`, `LATVIA`, `LITHUANIA`, `LUXEMBOURG`, `MALTA`, `MARTINIQUE`, `MAYOTTE`, `MONACO`, `MONTENEGRO`, `NETHERLANDS (THE)`, `PORTUGAL`, `RÉUNION`, `SAINT BARTHÉLEMY`, `SAINT MARTIN (FRENCH PART)`, `SAINT PIERRE AND MIQUELON`, `SAN MARINO`, `SLOVAKIA`, `SLOVENIA`, `SPAIN`, `FRENCH SOUTHERN TERRITORIES`, `HOLY SEE`, `NETHERLANDS`, `SAINT MARTIN`},
			currency:  `Euro`,
			code:      `EUR`,
			number:    `978`,
		},
	}, `FALKLAND ISLANDS (THE) [MALVINAS]`: {
		{
			countries: Countries{`FALKLAND ISLANDS (THE) [MALVINAS]`, `FALKLAND ISLANDS [MALVINAS]`},
			currency:  `Falkland Islands Pound`,
			code:      `FKP`,
			number:    `238`,
		},
	}, `FAROE ISLANDS (THE)`: {
		{
			countries: Countries{`DENMARK`, `FAROE ISLANDS (THE)`, `GREENLAND`, `FAROE ISLANDS`},
			currency:  `Danish Krone`,
			code:      `DKK`,
			number:    `208`,
		},
	}, `FIJI`: {
		{
			countries: Countries{`FIJI`},
			currency:  `Fiji Dollar`,
			code:      `FJD`,
			number:    `242`,
		},
	}, `FINLAND`: {
		{
			countries: Countries{`ÅLAND ISLANDS`, `ANDORRA`, `AUSTRIA`, `BELGIUM`, `CYPRUS`, `ESTONIA`, `EUROPEAN UNION`, `FINLAND`, `FRANCE`, `FRENCH GUIANA`, `FRENCH SOUTHERN TERRITORIES (THE)`, `GERMANY`, `GREECE`, `GUADELOUPE`, `HOLY SEE (THE)`, `IRELAND`, `ITALY`, `LATVIA`, `LITHUANIA`, `LUXEMBOURG`, `MALTA`, `MARTINIQUE`, `MAYOTTE`, `MONACO`, `MONTENEGRO`, `NETHERLANDS (THE)`, `PORTUGAL`, `RÉUNION`, `SAINT BARTHÉLEMY`, `SAINT MARTIN (FRENCH PART)`, `SAINT PIERRE AND MIQUELON`, `SAN MARINO`, `SLOVAKIA`, `SLOVENIA`, `SPAIN`, `FRENCH SOUTHERN TERRITORIES`, `HOLY SEE`, `NETHERLANDS`, `SAINT MARTIN`},
			currency:  `Euro`,
			code:      `EUR`,
			number:    `978`,
		},
	}, `FRANCE`: {
		{
			countries: Countries{`ÅLAND ISLANDS`, `ANDORRA`, `AUSTRIA`, `BELGIUM`, `CYPRUS`, `ESTONIA`, `EUROPEAN UNION`, `FINLAND`, `FRANCE`, `FRENCH GUIANA`, `FRENCH SOUTHERN TERRITORIES (THE)`, `GERMANY`, `GREECE`, `GUADELOUPE`, `HOLY SEE (THE)`, `IRELAND`, `ITALY`, `LATVIA`, `LITHUANIA`, `LUXEMBOURG`, `MALTA`, `MARTINIQUE`, `MAYOTTE`, `MONACO`, `MONTENEGRO`, `NETHERLANDS (THE)`, `PORTUGAL`, `RÉUNION`, `SAINT BARTHÉLEMY`, `SAINT MARTIN (FRENCH PART)`, `SAINT PIERRE AND MIQUELON`, `SAN MARINO`, `SLOVAKIA`, `SLOVENIA`, `SPAIN`, `FRENCH SOUTHERN TERRITORIES`, `HOLY SEE`, `NETHERLANDS`, `SAINT MARTIN`},
			currency:  `Euro`,
			code:      `EUR`,
			number:    `978`,
		},
	}, `FRENCH GUIANA`: {
		{
			countries: Countries{`ÅLAND ISLANDS`, `ANDORRA`, `AUSTRIA`, `BELGIUM`, `CYPRUS`, `ESTONIA`, `EUROPEAN UNION`, `FINLAND`, `FRANCE`, `FRENCH GUIANA`, `FRENCH SOUTHERN TERRITORIES (THE)`, `GERMANY`, `GREECE`, `GUADELOUPE`, `HOLY SEE (THE)`, `IRELAND`, `ITALY`, `LATVIA`, `LITHUANIA`, `LUXEMBOURG`, `MALTA`, `MARTINIQUE`, `MAYOTTE`, `MONACO`, `MONTENEGRO`, `NETHERLANDS (THE)`, `PORTUGAL`, `RÉUNION`, `SAINT BARTHÉLEMY`, `SAINT MARTIN (FRENCH PART)`, `SAINT PIERRE AND MIQUELON`, `SAN MARINO`, `SLOVAKIA`, `SLOVENIA`, `SPAIN`, `FRENCH SOUTHERN TERRITORIES`, `HOLY SEE`, `NETHERLANDS`, `SAINT MARTIN`},
			currency:  `Euro`,
			code:      `EUR`,
			number:    `978`,
		},
	}, `FRENCH POLYNESIA`: {
		{
			countries: Countries{`FRENCH POLYNESIA`, `NEW CALEDONIA`, `WALLIS AND FUTUNA`},
			currency:  `CFP Franc`,
			code:      `XPF`,
			number:    `953`,
		},
	}, `FRENCH SOUTHERN TERRITORIES (THE)`: {
		{
			countries: Countries{`ÅLAND ISLANDS`, `ANDORRA`, `AUSTRIA`, `BELGIUM`, `CYPRUS`, `ESTONIA`, `EUROPEAN UNION`, `FINLAND`, `FRANCE`, `FRENCH GUIANA`, `FRENCH SOUTHERN TERRITORIES (THE)`, `GERMANY`, `GREECE`, `GUADELOUPE`, `HOLY SEE (THE)`, `IRELAND`, `ITALY`, `LATVIA`, `LITHUANIA`, `LUXEMBOURG`, `MALTA`, `MARTINIQUE`, `MAYOTTE`, `MONACO`, `MONTENEGRO`, `NETHERLANDS (THE)`, `PORTUGAL`, `RÉUNION`, `SAINT BARTHÉLEMY`, `SAINT MARTIN (FRENCH PART)`, `SAINT PIERRE AND MIQUELON`, `SAN MARINO`, `SLOVAKIA`, `SLOVENIA`, `SPAIN`, `FRENCH SOUTHERN TERRITORIES`, `HOLY SEE`, `NETHERLANDS`, `SAINT MARTIN`},
			currency:  `Euro`,
			code:      `EUR`,
			number:    `978`,
		},
	}, `GABON`: {
		{
			countries: Countries{`CAMEROON`, `CENTRAL AFRICAN REPUBLIC (THE)`, `CHAD`, `CONGO (THE)`, `EQUATORIAL GUINEA`, `GABON`, `CENTRAL AFRICAN REPUBLIC`, `CONGO`},
			currency:  `CFA Franc BEAC`,
			code:      `XAF`,
			number:    `950`,
		},
	}, `GAMBIA (THE)`: {
		{
			countries: Countries{`GAMBIA (THE)`, `GAMBIA`},
			currency:  `Dalasi`,
			code:      `GMD`,
			number:    `270`,
		},
	}, `GEORGIA`: {
		{
			countries: Countries{`GEORGIA`},
			currency:  `Lari`,
			code:      `GEL`,
			number:    `981`,
		},
	}, `GERMANY`: {
		{
			countries: Countries{`ÅLAND ISLANDS`, `ANDORRA`, `AUSTRIA`, `BELGIUM`, `CYPRUS`, `ESTONIA`, `EUROPEAN UNION`, `FINLAND`, `FRANCE`, `FRENCH GUIANA`, `FRENCH SOUTHERN TERRITORIES (THE)`, `GERMANY`, `GREECE`, `GUADELOUPE`, `HOLY SEE (THE)`, `IRELAND`, `ITALY`, `LATVIA`, `LITHUANIA`, `LUXEMBOURG`, `MALTA`, `MARTINIQUE`, `MAYOTTE`, `MONACO`, `MONTENEGRO`, `NETHERLANDS (THE)`, `PORTUGAL`, `RÉUNION`, `SAINT BARTHÉLEMY`, `SAINT MARTIN (FRENCH PART)`, `SAINT PIERRE AND MIQUELON`, `SAN MARINO`, `SLOVAKIA`, `SLOVENIA`, `SPAIN`, `FRENCH SOUTHERN TERRITORIES`, `HOLY SEE`, `NETHERLANDS`, `SAINT MARTIN`},
			currency:  `Euro`,
			code:      `EUR`,
			number:    `978`,
		},
	}, `GHANA`: {
		{
			countries: Countries{`GHANA`},
			currency:  `Ghana Cedi`,
			code:      `GHS`,
			number:    `936`,
		},
	}, `GIBRALTAR`: {
		{
			countries: Countries{`GIBRALTAR`},
			currency:  `Gibraltar Pound`,
			code:      `GIP`,
			number:    `292`,
		},
	}, `GREECE`: {
		{
			countries: Countries{`ÅLAND ISLANDS`, `ANDORRA`, `AUSTRIA`, `BELGIUM`, `CYPRUS`, `ESTONIA`, `EUROPEAN UNION`, `FINLAND`, `FRANCE`, `FRENCH GUIANA`, `FRENCH SOUTHERN TERRITORIES (THE)`, `GERMANY`, `GREECE`, `GUADELOUPE`, `HOLY SEE (THE)`, `IRELAND`, `ITALY`, `LATVIA`, `LITHUANIA`, `LUXEMBOURG`, `MALTA`, `MARTINIQUE`, `MAYOTTE`, `MONACO`, `MONTENEGRO`, `NETHERLANDS (THE)`, `PORTUGAL`, `RÉUNION`, `SAINT BARTHÉLEMY`, `SAINT MARTIN (FRENCH PART)`, `SAINT PIERRE AND MIQUELON`, `SAN MARINO`, `SLOVAKIA`, `SLOVENIA`, `SPAIN`, `FRENCH SOUTHERN TERRITORIES`, `HOLY SEE`, `NETHERLANDS`, `SAINT MARTIN`},
			currency:  `Euro`,
			code:      `EUR`,
			number:    `978`,
		},
	}, `GREENLAND`: {
		{
			countries: Countries{`DENMARK`, `FAROE ISLANDS (THE)`, `GREENLAND`, `FAROE ISLANDS`},
			currency:  `Danish Krone`,
			code:      `DKK`,
			number:    `208`,
		},
	}, `GRENADA`: {
		{
			countries: Countries{`ANGUILLA`, `ANTIGUA AND BARBUDA`, `DOMINICA`, `GRENADA`, `MONTSERRAT`, `SAINT KITTS AND NEVIS`, `SAINT LUCIA`, `SAINT VINCENT AND THE GRENADINES`},
			currency:  `East Caribbean Dollar`,
			code:      `XCD`,
			number:    `951`,
		},
	}, `GUADELOUPE`: {
		{
			countries: Countries{`ÅLAND ISLANDS`, `ANDORRA`, `AUSTRIA`, `BELGIUM`, `CYPRUS`, `ESTONIA`, `EUROPEAN UNION`, `FINLAND`, `FRANCE`, `FRENCH GUIANA`, `FRENCH SOUTHERN TERRITORIES (THE)`, `GERMANY`, `GREECE`, `GUADELOUPE`, `HOLY SEE (THE)`, `IRELAND`, `ITALY`, `LATVIA`, `LITHUANIA`, `LUXEMBOURG`, `MALTA`, `MARTINIQUE`, `MAYOTTE`, `MONACO`, `MONTENEGRO`, `NETHERLANDS (THE)`, `PORTUGAL`, `RÉUNION`, `SAINT BARTHÉLEMY`, `SAINT MARTIN (FRENCH PART)`, `SAINT PIERRE AND MIQUELON`, `SAN MARINO`, `SLOVAKIA`, `SLOVENIA`, `SPAIN`, `FRENCH SOUTHERN TERRITORIES`, `HOLY SEE`, `NETHERLANDS`, `SAINT MARTIN`},
			currency:  `Euro`,
			code:      `EUR`,
			number:    `978`,
		},
	}, `GUAM`: {
		{
			countries: Countries{`AMERICAN SAMOA`, `BONAIRE, SINT EUSTATIUS AND SABA`, `BRITISH INDIAN OCEAN TERRITORY (THE)`, `ECUADOR`, `EL SALVADOR`, `GUAM`, `HAITI`, `MARSHALL ISLANDS (THE)`, `MICRONESIA (FEDERATED STATES OF)`, `NORTHERN MARIANA ISLANDS (THE)`, `PALAU`, `PANAMA`, `PUERTO RICO`, `TIMOR-LESTE`, `TURKS AND CAICOS ISLANDS (THE)`, `UNITED STATES MINOR OUTLYING ISLANDS (THE)`, `UNITED STATES OF AMERICA (THE)`, `VIRGIN ISLANDS (BRITISH)`, `VIRGIN ISLANDS (U.S.)`, `BRITISH INDIAN OCEAN TERRITORY`, `MARSHALL ISLANDS`, `MICRONESIA`, `NORTHERN MARIANA ISLANDS`, `TURKS AND CAICOS ISLANDS`, `UNITED STATES MINOR OUTLYING ISLANDS`, `UNITED STATES OF AMERICA`, `VIRGIN ISLANDS`, `VIRGIN ISLANDS`},
			currency:  `US Dollar`,
			code:      `USD`,
			number:    `840`,
		},
	}, `GUATEMALA`: {
		{
			countries: Countries{`GUATEMALA`},
			currency:  `Quetzal`,
			code:      `GTQ`,
			number:    `320`,
		},
	}, `GUERNSEY`: {
		{
			countries: Countries{`GUERNSEY`, `ISLE OF MAN`, `JERSEY`, `UNITED KINGDOM OF GREAT BRITAIN AND NORTHERN IRELAND (THE)`, `UNITED KINGDOM OF GREAT BRITAIN AND NORTHERN IRELAND`},
			currency:  `Pound Sterling`,
			code:      `GBP`,
			number:    `826`,
		},
	}, `GUINEA`: {
		{
			countries: Countries{`GUINEA`},
			currency:  `Guinean Franc`,
			code:      `GNF`,
			number:    `324`,
		},
	}, `GUINEA-BISSAU`: {
		{
			countries: Countries{`BENIN`, `BURKINA FASO`, `CÔTE D'IVOIRE`, `GUINEA-BISSAU`, `MALI`, `NIGER (THE)`, `SENEGAL`, `TOGO`, `NIGER`},
			currency:  `CFA Franc BCEAO`,
			code:      `XOF`,
			number:    `952`,
		},
	}, `GUYANA`: {
		{
			countries: Countries{`GUYANA`},
			currency:  `Guyana Dollar`,
			code:      `GYD`,
			number:    `328`,
		},
	}, `HAITI`: {
		{
			countries: Countries{`HAITI`},
			currency:  `Gourde`,
			code:      `HTG`,
			number:    `332`,
		}, {
			countries: Countries{`AMERICAN SAMOA`, `BONAIRE, SINT EUSTATIUS AND SABA`, `BRITISH INDIAN OCEAN TERRITORY (THE)`, `ECUADOR`, `EL SALVADOR`, `GUAM`, `HAITI`, `MARSHALL ISLANDS (THE)`, `MICRONESIA (FEDERATED STATES OF)`, `NORTHERN MARIANA ISLANDS (THE)`, `PALAU`, `PANAMA`, `PUERTO RICO`, `TIMOR-LESTE`, `TURKS AND CAICOS ISLANDS (THE)`, `UNITED STATES MINOR OUTLYING ISLANDS (THE)`, `UNITED STATES OF AMERICA (THE)`, `VIRGIN ISLANDS (BRITISH)`, `VIRGIN ISLANDS (U.S.)`, `BRITISH INDIAN OCEAN TERRITORY`, `MARSHALL ISLANDS`, `MICRONESIA`, `NORTHERN MARIANA ISLANDS`, `TURKS AND CAICOS ISLANDS`, `UNITED STATES MINOR OUTLYING ISLANDS`, `UNITED STATES OF AMERICA`, `VIRGIN ISLANDS`, `VIRGIN ISLANDS`},
			currency:  `US Dollar`,
			code:      `USD`,
			number:    `840`,
		},
	}, `HEARD ISLAND AND McDONALD ISLANDS`: {
		{
			countries: Countries{`AUSTRALIA`, `CHRISTMAS ISLAND`, `COCOS (KEELING) ISLANDS (THE)`, `HEARD ISLAND AND McDONALD ISLANDS`, `KIRIBATI`, `NAURU`, `NORFOLK ISLAND`, `TUVALU`, `COCOS ISLANDS`},
			currency:  `Australian Dollar`,
			code:      `AUD`,
			number:    `036`,
		},
	}, `HOLY SEE (THE)`: {
		{
			countries: Countries{`ÅLAND ISLANDS`, `ANDORRA`, `AUSTRIA`, `BELGIUM`, `CYPRUS`, `ESTONIA`, `EUROPEAN UNION`, `FINLAND`, `FRANCE`, `FRENCH GUIANA`, `FRENCH SOUTHERN TERRITORIES (THE)`, `GERMANY`, `GREECE`, `GUADELOUPE`, `HOLY SEE (THE)`, `IRELAND`, `ITALY`, `LATVIA`, `LITHUANIA`, `LUXEMBOURG`, `MALTA`, `MARTINIQUE`, `MAYOTTE`, `MONACO`, `MONTENEGRO`, `NETHERLANDS (THE)`, `PORTUGAL`, `RÉUNION`, `SAINT BARTHÉLEMY`, `SAINT MARTIN (FRENCH PART)`, `SAINT PIERRE AND MIQUELON`, `SAN MARINO`, `SLOVAKIA`, `SLOVENIA`, `SPAIN`, `FRENCH SOUTHERN TERRITORIES`, `HOLY SEE`, `NETHERLANDS`, `SAINT MARTIN`},
			currency:  `Euro`,
			code:      `EUR`,
			number:    `978`,
		},
	}, `HONDURAS`: {
		{
			countries: Countries{`HONDURAS`},
			currency:  `Lempira`,
			code:      `HNL`,
			number:    `340`,
		},
	}, `HONG KONG`: {
		{
			countries: Countries{`HONG KONG`},
			currency:  `Hong Kong Dollar`,
			code:      `HKD`,
			number:    `344`,
		},
	}, `HUNGARY`: {
		{
			countries: Countries{`HUNGARY`},
			currency:  `Forint`,
			code:      `HUF`,
			number:    `348`,
		},
	}, `ICELAND`: {
		{
			countries: Countries{`ICELAND`},
			currency:  `Iceland Krona`,
			code:      `ISK`,
			number:    `352`,
		},
	}, `INDIA`: {
		{
			countries: Countries{`BHUTAN`, `INDIA`},
			currency:  `Indian Rupee`,
			code:      `INR`,
			number:    `356`,
		},
	}, `INDONESIA`: {
		{
			countries: Countries{`INDONESIA`},
			currency:  `Rupiah`,
			code:      `IDR`,
			number:    `360`,
		},
	}, `INTERNATIONAL MONETARY FUND (IMF) `: {
		{
			countries: Countries{`INTERNATIONAL MONETARY FUND (IMF) `, `INTERNATIONAL MONETARY FUND `},
			currency:  `SDR (Special Drawing Right)`,
			code:      `XDR`,
			number:    `960`,
		},
	}, `IRAN (ISLAMIC REPUBLIC OF)`: {
		{
			countries: Countries{`IRAN (ISLAMIC REPUBLIC OF)`, `IRAN`},
			currency:  `Iranian Rial`,
			code:      `IRR`,
			number:    `364`,
		},
	}, `IRAQ`: {
		{
			countries: Countries{`IRAQ`},
			currency:  `Iraqi Dinar`,
			code:      `IQD`,
			number:    `368`,
		},
	}, `IRELAND`: {
		{
			countries: Countries{`ÅLAND ISLANDS`, `ANDORRA`, `AUSTRIA`, `BELGIUM`, `CYPRUS`, `ESTONIA`, `EUROPEAN UNION`, `FINLAND`, `FRANCE`, `FRENCH GUIANA`, `FRENCH SOUTHERN TERRITORIES (THE)`, `GERMANY`, `GREECE`, `GUADELOUPE`, `HOLY SEE (THE)`, `IRELAND`, `ITALY`, `LATVIA`, `LITHUANIA`, `LUXEMBOURG`, `MALTA`, `MARTINIQUE`, `MAYOTTE`, `MONACO`, `MONTENEGRO`, `NETHERLANDS (THE)`, `PORTUGAL`, `RÉUNION`, `SAINT BARTHÉLEMY`, `SAINT MARTIN (FRENCH PART)`, `SAINT PIERRE AND MIQUELON`, `SAN MARINO`, `SLOVAKIA`, `SLOVENIA`, `SPAIN`, `FRENCH SOUTHERN TERRITORIES`, `HOLY SEE`, `NETHERLANDS`, `SAINT MARTIN`},
			currency:  `Euro`,
			code:      `EUR`,
			number:    `978`,
		},
	}, `ISLE OF MAN`: {
		{
			countries: Countries{`GUERNSEY`, `ISLE OF MAN`, `JERSEY`, `UNITED KINGDOM OF GREAT BRITAIN AND NORTHERN IRELAND (THE)`, `UNITED KINGDOM OF GREAT BRITAIN AND NORTHERN IRELAND`},
			currency:  `Pound Sterling`,
			code:      `GBP`,
			number:    `826`,
		},
	}, `ISRAEL`: {
		{
			countries: Countries{`ISRAEL`},
			currency:  `New Israeli Sheqel`,
			code:      `ILS`,
			number:    `376`,
		},
	}, `ITALY`: {
		{
			countries: Countries{`ÅLAND ISLANDS`, `ANDORRA`, `AUSTRIA`, `BELGIUM`, `CYPRUS`, `ESTONIA`, `EUROPEAN UNION`, `FINLAND`, `FRANCE`, `FRENCH GUIANA`, `FRENCH SOUTHERN TERRITORIES (THE)`, `GERMANY`, `GREECE`, `GUADELOUPE`, `HOLY SEE (THE)`, `IRELAND`, `ITALY`, `LATVIA`, `LITHUANIA`, `LUXEMBOURG`, `MALTA`, `MARTINIQUE`, `MAYOTTE`, `MONACO`, `MONTENEGRO`, `NETHERLANDS (THE)`, `PORTUGAL`, `RÉUNION`, `SAINT BARTHÉLEMY`, `SAINT MARTIN (FRENCH PART)`, `SAINT PIERRE AND MIQUELON`, `SAN MARINO`, `SLOVAKIA`, `SLOVENIA`, `SPAIN`, `FRENCH SOUTHERN TERRITORIES`, `HOLY SEE`, `NETHERLANDS`, `SAINT MARTIN`},
			currency:  `Euro`,
			code:      `EUR`,
			number:    `978`,
		},
	}, `JAMAICA`: {
		{
			countries: Countries{`JAMAICA`},
			currency:  `Jamaican Dollar`,
			code:      `JMD`,
			number:    `388`,
		},
	}, `JAPAN`: {
		{
			countries: Countries{`JAPAN`},
			currency:  `Yen`,
			code:      `JPY`,
			number:    `392`,
		},
	}, `JERSEY`: {
		{
			countries: Countries{`GUERNSEY`, `ISLE OF MAN`, `JERSEY`, `UNITED KINGDOM OF GREAT BRITAIN AND NORTHERN IRELAND (THE)`, `UNITED KINGDOM OF GREAT BRITAIN AND NORTHERN IRELAND`},
			currency:  `Pound Sterling`,
			code:      `GBP`,
			number:    `826`,
		},
	}, `JORDAN`: {
		{
			countries: Countries{`JORDAN`},
			currency:  `Jordanian Dinar`,
			code:      `JOD`,
			number:    `400`,
		},
	}, `KAZAKHSTAN`: {
		{
			countries: Countries{`KAZAKHSTAN`},
			currency:  `Tenge`,
			code:      `KZT`,
			number:    `398`,
		},
	}, `KENYA`: {
		{
			countries: Countries{`KENYA`},
			currency:  `Kenyan Shilling`,
			code:      `KES`,
			number:    `404`,
		},
	}, `KIRIBATI`: {
		{
			countries: Countries{`AUSTRALIA`, `CHRISTMAS ISLAND`, `COCOS (KEELING) ISLANDS (THE)`, `HEARD ISLAND AND McDONALD ISLANDS`, `KIRIBATI`, `NAURU`, `NORFOLK ISLAND`, `TUVALU`, `COCOS ISLANDS`},
			currency:  `Australian Dollar`,
			code:      `AUD`,
			number:    `036`,
		},
	}, `KOREA (THE DEMOCRATIC PEOPLE’S REPUBLIC OF)`: {
		{
			countries: Countries{`KOREA (THE DEMOCRATIC PEOPLE’S REPUBLIC OF)`, `KOREA`},
			currency:  `North Korean Won`,
			code:      `KPW`,
			number:    `408`,
		},
	}, `KOREA (THE REPUBLIC OF)`: {
		{
			countries: Countries{`KOREA (THE REPUBLIC OF)`, `KOREA`},
			currency:  `Won`,
			code:      `KRW`,
			number:    `410`,
		},
	}, `KUWAIT`: {
		{
			countries: Countries{`KUWAIT`},
			currency:  `Kuwaiti Dinar`,
			code:      `KWD`,
			number:    `414`,
		},
	}, `KYRGYZSTAN`: {
		{
			countries: Countries{`KYRGYZSTAN`},
			currency:  `Som`,
			code:      `KGS`,
			number:    `417`,
		},
	}, `LAO PEOPLE’S DEMOCRATIC REPUBLIC (THE)`: {
		{
			countries: Countries{`LAO PEOPLE’S DEMOCRATIC REPUBLIC (THE)`, `LAO PEOPLE’S DEMOCRATIC REPUBLIC`},
			currency:  `Lao Kip`,
			code:      `LAK`,
			number:    `418`,
		},
	}, `LATVIA`: {
		{
			countries: Countries{`ÅLAND ISLANDS`, `ANDORRA`, `AUSTRIA`, `BELGIUM`, `CYPRUS`, `ESTONIA`, `EUROPEAN UNION`, `FINLAND`, `FRANCE`, `FRENCH GUIANA`, `FRENCH SOUTHERN TERRITORIES (THE)`, `GERMANY`, `GREECE`, `GUADELOUPE`, `HOLY SEE (THE)`, `IRELAND`, `ITALY`, `LATVIA`, `LITHUANIA`, `LUXEMBOURG`, `MALTA`, `MARTINIQUE`, `MAYOTTE`, `MONACO`, `MONTENEGRO`, `NETHERLANDS (THE)`, `PORTUGAL`, `RÉUNION`, `SAINT BARTHÉLEMY`, `SAINT MARTIN (FRENCH PART)`, `SAINT PIERRE AND MIQUELON`, `SAN MARINO`, `SLOVAKIA`, `SLOVENIA`, `SPAIN`, `FRENCH SOUTHERN TERRITORIES`, `HOLY SEE`, `NETHERLANDS`, `SAINT MARTIN`},
			currency:  `Euro`,
			code:      `EUR`,
			number:    `978`,
		},
	}, `LEBANON`: {
		{
			countries: Countries{`LEBANON`},
			currency:  `Lebanese Pound`,
			code:      `LBP`,
			number:    `422`,
		},
	}, `LESOTHO`: {
		{
			countries: Countries{`LESOTHO`},
			currency:  `Loti`,
			code:      `LSL`,
			number:    `426`,
		}, {
			countries: Countries{`LESOTHO`, `NAMIBIA`, `SOUTH AFRICA`},
			currency:  `Rand`,
			code:      `ZAR`,
			number:    `710`,
		},
	}, `LIBERIA`: {
		{
			countries: Countries{`LIBERIA`},
			currency:  `Liberian Dollar`,
			code:      `LRD`,
			number:    `430`,
		},
	}, `LIBYA`: {
		{
			countries: Countries{`LIBYA`},
			currency:  `Libyan Dinar`,
			code:      `LYD`,
			number:    `434`,
		},
	}, `LIECHTENSTEIN`: {
		{
			countries: Countries{`LIECHTENSTEIN`, `SWITZERLAND`},
			currency:  `Swiss Franc`,
			code:      `CHF`,
			number:    `756`,
		},
	}, `LITHUANIA`: {
		{
			countries: Countries{`ÅLAND ISLANDS`, `ANDORRA`, `AUSTRIA`, `BELGIUM`, `CYPRUS`, `ESTONIA`, `EUROPEAN UNION`, `FINLAND`, `FRANCE`, `FRENCH GUIANA`, `FRENCH SOUTHERN TERRITORIES (THE)`, `GERMANY`, `GREECE`, `GUADELOUPE`, `HOLY SEE (THE)`, `IRELAND`, `ITALY`, `LATVIA`, `LITHUANIA`, `LUXEMBOURG`, `MALTA`, `MARTINIQUE`, `MAYOTTE`, `MONACO`, `MONTENEGRO`, `NETHERLANDS (THE)`, `PORTUGAL`, `RÉUNION`, `SAINT BARTHÉLEMY`, `SAINT MARTIN (FRENCH PART)`, `SAINT PIERRE AND MIQUELON`, `SAN MARINO`, `SLOVAKIA`, `SLOVENIA`, `SPAIN`, `FRENCH SOUTHERN TERRITORIES`, `HOLY SEE`, `NETHERLANDS`, `SAINT MARTIN`},
			currency:  `Euro`,
			code:      `EUR`,
			number:    `978`,
		},
	}, `LUXEMBOURG`: {
		{
			countries: Countries{`ÅLAND ISLANDS`, `ANDORRA`, `AUSTRIA`, `BELGIUM`, `CYPRUS`, `ESTONIA`, `EUROPEAN UNION`, `FINLAND`, `FRANCE`, `FRENCH GUIANA`, `FRENCH SOUTHERN TERRITORIES (THE)`, `GERMANY`, `GREECE`, `GUADELOUPE`, `HOLY SEE (THE)`, `IRELAND`, `ITALY`, `LATVIA`, `LITHUANIA`, `LUXEMBOURG`, `MALTA`, `MARTINIQUE`, `MAYOTTE`, `MONACO`, `MONTENEGRO`, `NETHERLANDS (THE)`, `PORTUGAL`, `RÉUNION`, `SAINT BARTHÉLEMY`, `SAINT MARTIN (FRENCH PART)`, `SAINT PIERRE AND MIQUELON`, `SAN MARINO`, `SLOVAKIA`, `SLOVENIA`, `SPAIN`, `FRENCH SOUTHERN TERRITORIES`, `HOLY SEE`, `NETHERLANDS`, `SAINT MARTIN`},
			currency:  `Euro`,
			code:      `EUR`,
			number:    `978`,
		},
	}, `MACAO`: {
		{
			countries: Countries{`MACAO`},
			currency:  `Pataca`,
			code:      `MOP`,
			number:    `446`,
		},
	}, `NORTH MACEDONIA`: {
		{
			countries: Countries{`NORTH MACEDONIA`},
			currency:  `Denar`,
			code:      `MKD`,
			number:    `807`,
		},
	}, `MADAGASCAR`: {
		{
			countries: Countries{`MADAGASCAR`},
			currency:  `Malagasy Ariary`,
			code:      `MGA`,
			number:    `969`,
		},
	}, `MALAWI`: {
		{
			countries: Countries{`MALAWI`},
			currency:  `Malawi Kwacha`,
			code:      `MWK`,
			number:    `454`,
		},
	}, `MALAYSIA`: {
		{
			countries: Countries{`MALAYSIA`},
			currency:  `Malaysian Ringgit`,
			code:      `MYR`,
			number:    `458`,
		},
	}, `MALDIVES`: {
		{
			countries: Countries{`MALDIVES`},
			currency:  `Rufiyaa`,
			code:      `MVR`,
			number:    `462`,
		},
	}, `MALI`: {
		{
			countries: Countries{`BENIN`, `BURKINA FASO`, `CÔTE D'IVOIRE`, `GUINEA-BISSAU`, `MALI`, `NIGER (THE)`, `SENEGAL`, `TOGO`, `NIGER`},
			currency:  `CFA Franc BCEAO`,
			code:      `XOF`,
			number:    `952`,
		},
	}, `MALTA`: {
		{
			countries: Countries{`ÅLAND ISLANDS`, `ANDORRA`, `AUSTRIA`, `BELGIUM`, `CYPRUS`, `ESTONIA`, `EUROPEAN UNION`, `FINLAND`, `FRANCE`, `FRENCH GUIANA`, `FRENCH SOUTHERN TERRITORIES (THE)`, `GERMANY`, `GREECE`, `GUADELOUPE`, `HOLY SEE (THE)`, `IRELAND`, `ITALY`, `LATVIA`, `LITHUANIA`, `LUXEMBOURG`, `MALTA`, `MARTINIQUE`, `MAYOTTE`, `MONACO`, `MONTENEGRO`, `NETHERLANDS (THE)`, `PORTUGAL`, `RÉUNION`, `SAINT BARTHÉLEMY`, `SAINT MARTIN (FRENCH PART)`, `SAINT PIERRE AND MIQUELON`, `SAN MARINO`, `SLOVAKIA`, `SLOVENIA`, `SPAIN`, `FRENCH SOUTHERN TERRITORIES`, `HOLY SEE`, `NETHERLANDS`, `SAINT MARTIN`},
			currency:  `Euro`,
			code:      `EUR`,
			number:    `978`,
		},
	}, `MARSHALL ISLANDS (THE)`: {
		{
			countries: Countries{`AMERICAN SAMOA`, `BONAIRE, SINT EUSTATIUS AND SABA`, `BRITISH INDIAN OCEAN TERRITORY (THE)`, `ECUADOR`, `EL SALVADOR`, `GUAM`, `HAITI`, `MARSHALL ISLANDS (THE)`, `MICRONESIA (FEDERATED STATES OF)`, `NORTHERN MARIANA ISLANDS (THE)`, `PALAU`, `PANAMA`, `PUERTO RICO`, `TIMOR-LESTE`, `TURKS AND CAICOS ISLANDS (THE)`, `UNITED STATES MINOR OUTLYING ISLANDS (THE)`, `UNITED STATES OF AMERICA (THE)`, `VIRGIN ISLANDS (BRITISH)`, `VIRGIN ISLANDS (U.S.)`, `BRITISH INDIAN OCEAN TERRITORY`, `MARSHALL ISLANDS`, `MICRONESIA`, `NORTHERN MARIANA ISLANDS`, `TURKS AND CAICOS ISLANDS`, `UNITED STATES MINOR OUTLYING ISLANDS`, `UNITED STATES OF AMERICA`, `VIRGIN ISLANDS`, `VIRGIN ISLANDS`},
			currency:  `US Dollar`,
			code:      `USD`,
			number:    `840`,
		},
	}, `MARTINIQUE`: {
		{
			countries: Countries{`ÅLAND ISLANDS`, `ANDORRA`, `AUSTRIA`, `BELGIUM`, `CYPRUS`, `ESTONIA`, `EUROPEAN UNION`, `FINLAND`, `FRANCE`, `FRENCH GUIANA`, `FRENCH SOUTHERN TERRITORIES (THE)`, `GERMANY`, `GREECE`, `GUADELOUPE`, `HOLY SEE (THE)`, `IRELAND`, `ITALY`, `LATVIA`, `LITHUANIA`, `LUXEMBOURG`, `MALTA`, `MARTINIQUE`, `MAYOTTE`, `MONACO`, `MONTENEGRO`, `NETHERLANDS (THE)`, `PORTUGAL`, `RÉUNION`, `SAINT BARTHÉLEMY`, `SAINT MARTIN (FRENCH PART)`, `SAINT PIERRE AND MIQUELON`, `SAN MARINO`, `SLOVAKIA`, `SLOVENIA`, `SPAIN`, `FRENCH SOUTHERN TERRITORIES`, `HOLY SEE`, `NETHERLANDS`, `SAINT MARTIN`},
			currency:  `Euro`,
			code:      `EUR`,
			number:    `978`,
		},
	}, `MAURITANIA`: {
		{
			countries: Countries{`MAURITANIA`},
			currency:  `Ouguiya`,
			code:      `MRU`,
			number:    `929`,
		},
	}, `MAURITIUS`: {
		{
			countries: Countries{`MAURITIUS`},
			currency:  `Mauritius Rupee`,
			code:      `MUR`,
			number:    `480`,
		},
	}, `MAYOTTE`: {
		{
			countries: Countries{`ÅLAND ISLANDS`, `ANDORRA`, `AUSTRIA`, `BELGIUM`, `CYPRUS`, `ESTONIA`, `EUROPEAN UNION`, `FINLAND`, `FRANCE`, `FRENCH GUIANA`, `FRENCH SOUTHERN TERRITORIES (THE)`, `GERMANY`, `GREECE`, `GUADELOUPE`, `HOLY SEE (THE)`, `IRELAND`, `ITALY`, `LATVIA`, `LITHUANIA`, `LUXEMBOURG`, `MALTA`, `MARTINIQUE`, `MAYOTTE`, `MONACO`, `MONTENEGRO`, `NETHERLANDS (THE)`, `PORTUGAL`, `RÉUNION`, `SAINT BARTHÉLEMY`, `SAINT MARTIN (FRENCH PART)`, `SAINT PIERRE AND MIQUELON`, `SAN MARINO`, `SLOVAKIA`, `SLOVENIA`, `SPAIN`, `FRENCH SOUTHERN TERRITORIES`, `HOLY SEE`, `NETHERLANDS`, `SAINT MARTIN`},
			currency:  `Euro`,
			code:      `EUR`,
			number:    `978`,
		},
	}, `MEMBER COUNTRIES OF THE AFRICAN DEVELOPMENT BANK GROUP`: {
		{
			countries: Countries{`MEMBER COUNTRIES OF THE AFRICAN DEVELOPMENT BANK GROUP`},
			currency:  `ADB Unit of Account`,
			code:      `XUA`,
			number:    `965`,
		},
	}, `MEXICO`: {
		{
			countries: Countries{`MEXICO`},
			currency:  `Mexican Peso`,
			code:      `MXN`,
			number:    `484`,
		}, {
			countries: Countries{`MEXICO`},
			currency:  `Mexican Unidad de Inversion (UDI)`,
			code:      `MXV`,
			number:    `979`,
		},
	}, `MICRONESIA (FEDERATED STATES OF)`: {
		{
			countries: Countries{`AMERICAN SAMOA`, `BONAIRE, SINT EUSTATIUS AND SABA`, `BRITISH INDIAN OCEAN TERRITORY (THE)`, `ECUADOR`, `EL SALVADOR`, `GUAM`, `HAITI`, `MARSHALL ISLANDS (THE)`, `MICRONESIA (FEDERATED STATES OF)`, `NORTHERN MARIANA ISLANDS (THE)`, `PALAU`, `PANAMA`, `PUERTO RICO`, `TIMOR-LESTE`, `TURKS AND CAICOS ISLANDS (THE)`, `UNITED STATES MINOR OUTLYING ISLANDS (THE)`, `UNITED STATES OF AMERICA (THE)`, `VIRGIN ISLANDS (BRITISH)`, `VIRGIN ISLANDS (U.S.)`, `BRITISH INDIAN OCEAN TERRITORY`, `MARSHALL ISLANDS`, `MICRONESIA`, `NORTHERN MARIANA ISLANDS`, `TURKS AND CAICOS ISLANDS`, `UNITED STATES MINOR OUTLYING ISLANDS`, `UNITED STATES OF AMERICA`, `VIRGIN ISLANDS`, `VIRGIN ISLANDS`},
			currency:  `US Dollar`,
			code:      `USD`,
			number:    `840`,
		},
	}, `MOLDOVA (THE REPUBLIC OF)`: {
		{
			countries: Countries{`MOLDOVA (THE REPUBLIC OF)`, `MOLDOVA`},
			currency:  `Moldovan Leu`,
			code:      `MDL`,
			number:    `498`,
		},
	}, `MONACO`: {
		{
			countries: Countries{`ÅLAND ISLANDS`, `ANDORRA`, `AUSTRIA`, `BELGIUM`, `CYPRUS`, `ESTONIA`, `EUROPEAN UNION`, `FINLAND`, `FRANCE`, `FRENCH GUIANA`, `FRENCH SOUTHERN TERRITORIES (THE)`, `GERMANY`, `GREECE`, `GUADELOUPE`, `HOLY SEE (THE)`, `IRELAND`, `ITALY`, `LATVIA`, `LITHUANIA`, `LUXEMBOURG`, `MALTA`, `MARTINIQUE`, `MAYOTTE`, `MONACO`, `MONTENEGRO`, `NETHERLANDS (THE)`, `PORTUGAL`, `RÉUNION`, `SAINT BARTHÉLEMY`, `SAINT MARTIN (FRENCH PART)`, `SAINT PIERRE AND MIQUELON`, `SAN MARINO`, `SLOVAKIA`, `SLOVENIA`, `SPAIN`, `FRENCH SOUTHERN TERRITORIES`, `HOLY SEE`, `NETHERLANDS`, `SAINT MARTIN`},
			currency:  `Euro`,
			code:      `EUR`,
			number:    `978`,
		},
	}, `MONGOLIA`: {
		{
			countries: Countries{`MONGOLIA`},
			currency:  `Tugrik`,
			code:      `MNT`,
			number:    `496`,
		},
	}, `MONTENEGRO`: {
		{
			countries: Countries{`ÅLAND ISLANDS`, `ANDORRA`, `AUSTRIA`, `BELGIUM`, `CYPRUS`, `ESTONIA`, `EUROPEAN UNION`, `FINLAND`, `FRANCE`, `FRENCH GUIANA`, `FRENCH SOUTHERN TERRITORIES (THE)`, `GERMANY`, `GREECE`, `GUADELOUPE`, `HOLY SEE (THE)`, `IRELAND`, `ITALY`, `LATVIA`, `LITHUANIA`, `LUXEMBOURG`, `MALTA`, `MARTINIQUE`, `MAYOTTE`, `MONACO`, `MONTENEGRO`, `NETHERLANDS (THE)`, `PORTUGAL`, `RÉUNION`, `SAINT BARTHÉLEMY`, `SAINT MARTIN (FRENCH PART)`, `SAINT PIERRE AND MIQUELON`, `SAN MARINO`, `SLOVAKIA`, `SLOVENIA`, `SPAIN`, `FRENCH SOUTHERN TERRITORIES`, `HOLY SEE`, `NETHERLANDS`, `SAINT MARTIN`},
			currency:  `Euro`,
			code:      `EUR`,
			number:    `978`,
		},
	}, `MONTSERRAT`: {
		{
			countries: Countries{`ANGUILLA`, `ANTIGUA AND BARBUDA`, `DOMINICA`, `GRENADA`, `MONTSERRAT`, `SAINT KITTS AND NEVIS`, `SAINT LUCIA`, `SAINT VINCENT AND THE GRENADINES`},
			currency:  `East Caribbean Dollar`,
			code:      `XCD`,
			number:    `951`,
		},
	}, `MOROCCO`: {
		{
			countries: Countries{`MOROCCO`, `WESTERN SAHARA`},
			currency:  `Moroccan Dirham`,
			code:      `MAD`,
			number:    `504`,
		},
	}, `MOZAMBIQUE`: {
		{
			countries: Countries{`MOZAMBIQUE`},
			currency:  `Mozambique Metical`,
			code:      `MZN`,
			number:    `943`,
		},
	}, `MYANMAR`: {
		{
			countries: Countries{`MYANMAR`},
			currency:  `Kyat`,
			code:      `MMK`,
			number:    `104`,
		},
	}, `NAMIBIA`: {
		{
			countries: Countries{`NAMIBIA`},
			currency:  `Namibia Dollar`,
			code:      `NAD`,
			number:    `516`,
		}, {
			countries: Countries{`LESOTHO`, `NAMIBIA`, `SOUTH AFRICA`},
			currency:  `Rand`,
			code:      `ZAR`,
			number:    `710`,
		},
	}, `NAURU`: {
		{
			countries: Countries{`AUSTRALIA`, `CHRISTMAS ISLAND`, `COCOS (KEELING) ISLANDS (THE)`, `HEARD ISLAND AND McDONALD ISLANDS`, `KIRIBATI`, `NAURU`, `NORFOLK ISLAND`, `TUVALU`, `COCOS ISLANDS`},
			currency:  `Australian Dollar`,
			code:      `AUD`,
			number:    `036`,
		},
	}, `NEPAL`: {
		{
			countries: Countries{`NEPAL`},
			currency:  `Nepalese Rupee`,
			code:      `NPR`,
			number:    `524`,
		},
	}, `NETHERLANDS (THE)`: {
		{
			countries: Countries{`ÅLAND ISLANDS`, `ANDORRA`, `AUSTRIA`, `BELGIUM`, `CYPRUS`, `ESTONIA`, `EUROPEAN UNION`, `FINLAND`, `FRANCE`, `FRENCH GUIANA`, `FRENCH SOUTHERN TERRITORIES (THE)`, `GERMANY`, `GREECE`, `GUADELOUPE`, `HOLY SEE (THE)`, `IRELAND`, `ITALY`, `LATVIA`, `LITHUANIA`, `LUXEMBOURG`, `MALTA`, `MARTINIQUE`, `MAYOTTE`, `MONACO`, `MONTENEGRO`, `NETHERLANDS (THE)`, `PORTUGAL`, `RÉUNION`, `SAINT BARTHÉLEMY`, `SAINT MARTIN (FRENCH PART)`, `SAINT PIERRE AND MIQUELON`, `SAN MARINO`, `SLOVAKIA`, `SLOVENIA`, `SPAIN`, `FRENCH SOUTHERN TERRITORIES`, `HOLY SEE`, `NETHERLANDS`, `SAINT MARTIN`},
			currency:  `Euro`,
			code:      `EUR`,
			number:    `978`,
		},
	}, `NEW CALEDONIA`: {
		{
			countries: Countries{`FRENCH POLYNESIA`, `NEW CALEDONIA`, `WALLIS AND FUTUNA`},
			currency:  `CFP Franc`,
			code:      `XPF`,
			number:    `953`,
		},
	}, `NEW ZEALAND`: {
		{
			countries: Countries{`COOK ISLANDS (THE)`, `NEW ZEALAND`, `NIUE`, `PITCAIRN`, `TOKELAU`, `COOK ISLANDS`},
			currency:  `New Zealand Dollar`,
			code:      `NZD`,
			number:    `554`,
		},
	}, `NICARAGUA`: {
		{
			countries: Countries{`NICARAGUA`},
			currency:  `Cordoba Oro`,
			code:      `NIO`,
			number:    `558`,
		},
	}, `NIGER (THE)`: {
		{
			countries: Countries{`BENIN`, `BURKINA FASO`, `CÔTE D'IVOIRE`, `GUINEA-BISSAU`, `MALI`, `NIGER (THE)`, `SENEGAL`, `TOGO`, `NIGER`},
			currency:  `CFA Franc BCEAO`,
			code:      `XOF`,
			number:    `952`,
		},
	}, `NIGERIA`: {
		{
			countries: Countries{`NIGERIA`},
			currency:  `Naira`,
			code:      `NGN`,
			number:    `566`,
		},
	}, `NIUE`: {
		{
			countries: Countries{`COOK ISLANDS (THE)`, `NEW ZEALAND`, `NIUE`, `PITCAIRN`, `TOKELAU`, `COOK ISLANDS`},
			currency:  `New Zealand Dollar`,
			code:      `NZD`,
			number:    `554`,
		},
	}, `NORFOLK ISLAND`: {
		{
			countries: Countries{`AUSTRALIA`, `CHRISTMAS ISLAND`, `COCOS (KEELING) ISLANDS (THE)`, `HEARD ISLAND AND McDONALD ISLANDS`, `KIRIBATI`, `NAURU`, `NORFOLK ISLAND`, `TUVALU`, `COCOS ISLANDS`},
			currency:  `Australian Dollar`,
			code:      `AUD`,
			number:    `036`,
		},
	}, `NORTHERN MARIANA ISLANDS (THE)`: {
		{
			countries: Countries{`AMERICAN SAMOA`, `BONAIRE, SINT EUSTATIUS AND SABA`, `BRITISH INDIAN OCEAN TERRITORY (THE)`, `ECUADOR`, `EL SALVADOR`, `GUAM`, `HAITI`, `MARSHALL ISLANDS (THE)`, `MICRONESIA (FEDERATED STATES OF)`, `NORTHERN MARIANA ISLANDS (THE)`, `PALAU`, `PANAMA`, `PUERTO RICO`, `TIMOR-LESTE`, `TURKS AND CAICOS ISLANDS (THE)`, `UNITED STATES MINOR OUTLYING ISLANDS (THE)`, `UNITED STATES OF AMERICA (THE)`, `VIRGIN ISLANDS (BRITISH)`, `VIRGIN ISLANDS (U.S.)`, `BRITISH INDIAN OCEAN TERRITORY`, `MARSHALL ISLANDS`, `MICRONESIA`, `NORTHERN MARIANA ISLANDS`, `TURKS AND CAICOS ISLANDS`, `UNITED STATES MINOR OUTLYING ISLANDS`, `UNITED STATES OF AMERICA`, `VIRGIN ISLANDS`, `VIRGIN ISLANDS`},
			currency:  `US Dollar`,
			code:      `USD`,
			number:    `840`,
		},
	}, `NORWAY`: {
		{
			countries: Countries{`BOUVET ISLAND`, `NORWAY`, `SVALBARD AND JAN MAYEN`},
			currency:  `Norwegian Krone`,
			code:      `NOK`,
			number:    `578`,
		},
	}, `OMAN`: {
		{
			countries: Countries{`OMAN`},
			currency:  `Rial Omani`,
			code:      `OMR`,
			number:    `512`,
		},
	}, `PAKISTAN`: {
		{
			countries: Countries{`PAKISTAN`},
			currency:  `Pakistan Rupee`,
			code:      `PKR`,
			number:    `586`,
		},
	}, `PALAU`: {
		{
			countries: Countries{`AMERICAN SAMOA`, `BONAIRE, SINT EUSTATIUS AND SABA`, `BRITISH INDIAN OCEAN TERRITORY (THE)`, `ECUADOR`, `EL SALVADOR`, `GUAM`, `HAITI`, `MARSHALL ISLANDS (THE)`, `MICRONESIA (FEDERATED STATES OF)`, `NORTHERN MARIANA ISLANDS (THE)`, `PALAU`, `PANAMA`, `PUERTO RICO`, `TIMOR-LESTE`, `TURKS AND CAICOS ISLANDS (THE)`, `UNITED STATES MINOR OUTLYING ISLANDS (THE)`, `UNITED STATES OF AMERICA (THE)`, `VIRGIN ISLANDS (BRITISH)`, `VIRGIN ISLANDS (U.S.)`, `BRITISH INDIAN OCEAN TERRITORY`, `MARSHALL ISLANDS`, `MICRONESIA`, `NORTHERN MARIANA ISLANDS`, `TURKS AND CAICOS ISLANDS`, `UNITED STATES MINOR OUTLYING ISLANDS`, `UNITED STATES OF AMERICA`, `VIRGIN ISLANDS`, `VIRGIN ISLANDS`},
			currency:  `US Dollar`,
			code:      `USD`,
			number:    `840`,
		},
	}, `PANAMA`: {
		{
			countries: Countries{`PANAMA`},
			currency:  `Balboa`,
			code:      `PAB`,
			number:    `590`,
		}, {
			countries: Countries{`AMERICAN SAMOA`, `BONAIRE, SINT EUSTATIUS AND SABA`, `BRITISH INDIAN OCEAN TERRITORY (THE)`, `ECUADOR`, `EL SALVADOR`, `GUAM`, `HAITI`, `MARSHALL ISLANDS (THE)`, `MICRONESIA (FEDERATED STATES OF)`, `NORTHERN MARIANA ISLANDS (THE)`, `PALAU`, `PANAMA`, `PUERTO RICO`, `TIMOR-LESTE`, `TURKS AND CAICOS ISLANDS (THE)`, `UNITED STATES MINOR OUTLYING ISLANDS (THE)`, `UNITED STATES OF AMERICA (THE)`, `VIRGIN ISLANDS (BRITISH)`, `VIRGIN ISLANDS (U.S.)`, `BRITISH INDIAN OCEAN TERRITORY`, `MARSHALL ISLANDS`, `MICRONESIA`, `NORTHERN MARIANA ISLANDS`, `TURKS AND CAICOS ISLANDS`, `UNITED STATES MINOR OUTLYING ISLANDS`, `UNITED STATES OF AMERICA`, `VIRGIN ISLANDS`, `VIRGIN ISLANDS`},
			currency:  `US Dollar`,
			code:      `USD`,
			number:    `840`,
		},
	}, `PAPUA NEW GUINEA`: {
		{
			countries: Countries{`PAPUA NEW GUINEA`},
			currency:  `Kina`,
			code:      `PGK`,
			number:    `598`,
		},
	}, `PARAGUAY`: {
		{
			countries: Countries{`PARAGUAY`},
			currency:  `Guarani`,
			code:      `PYG`,
			number:    `600`,
		},
	}, `PERU`: {
		{
			countries: Countries{`PERU`},
			currency:  `Sol`,
			code:      `PEN`,
			number:    `604`,
		},
	}, `PHILIPPINES (THE)`: {
		{
			countries: Countries{`PHILIPPINES (THE)`, `PHILIPPINES`},
			currency:  `Philippine Peso`,
			code:      `PHP`,
			number:    `608`,
		},
	}, `PITCAIRN`: {
		{
			countries: Countries{`COOK ISLANDS (THE)`, `NEW ZEALAND`, `NIUE`, `PITCAIRN`, `TOKELAU`, `COOK ISLANDS`},
			currency:  `New Zealand Dollar`,
			code:      `NZD`,
			number:    `554`,
		},
	}, `POLAND`: {
		{
			countries: Countries{`POLAND`},
			currency:  `Zloty`,
			code:      `PLN`,
			number:    `985`,
		},
	}, `PORTUGAL`: {
		{
			countries: Countries{`ÅLAND ISLANDS`, `ANDORRA`, `AUSTRIA`, `BELGIUM`, `CYPRUS`, `ESTONIA`, `EUROPEAN UNION`, `FINLAND`, `FRANCE`, `FRENCH GUIANA`, `FRENCH SOUTHERN TERRITORIES (THE)`, `GERMANY`, `GREECE`, `GUADELOUPE`, `HOLY SEE (THE)`, `IRELAND`, `ITALY`, `LATVIA`, `LITHUANIA`, `LUXEMBOURG`, `MALTA`, `MARTINIQUE`, `MAYOTTE`, `MONACO`, `MONTENEGRO`, `NETHERLANDS (THE)`, `PORTUGAL`, `RÉUNION`, `SAINT BARTHÉLEMY`, `SAINT MARTIN (FRENCH PART)`, `SAINT PIERRE AND MIQUELON`, `SAN MARINO`, `SLOVAKIA`, `SLOVENIA`, `SPAIN`, `FRENCH SOUTHERN TERRITORIES`, `HOLY SEE`, `NETHERLANDS`, `SAINT MARTIN`},
			currency:  `Euro`,
			code:      `EUR`,
			number:    `978`,
		},
	}, `PUERTO RICO`: {
		{
			countries: Countries{`AMERICAN SAMOA`, `BONAIRE, SINT EUSTATIUS AND SABA`, `BRITISH INDIAN OCEAN TERRITORY (THE)`, `ECUADOR`, `EL SALVADOR`, `GUAM`, `HAITI`, `MARSHALL ISLANDS (THE)`, `MICRONESIA (FEDERATED STATES OF)`, `NORTHERN MARIANA ISLANDS (THE)`, `PALAU`, `PANAMA`, `PUERTO RICO`, `TIMOR-LESTE`, `TURKS AND CAICOS ISLANDS (THE)`, `UNITED STATES MINOR OUTLYING ISLANDS (THE)`, `UNITED STATES OF AMERICA (THE)`, `VIRGIN ISLANDS (BRITISH)`, `VIRGIN ISLANDS (U.S.)`, `BRITISH INDIAN OCEAN TERRITORY`, `MARSHALL ISLANDS`, `MICRONESIA`, `NORTHERN MARIANA ISLANDS`, `TURKS AND CAICOS ISLANDS`, `UNITED STATES MINOR OUTLYING ISLANDS`, `UNITED STATES OF AMERICA`, `VIRGIN ISLANDS`, `VIRGIN ISLANDS`},
			currency:  `US Dollar`,
			code:      `USD`,
			number:    `840`,
		},
	}, `QATAR`: {
		{
			countries: Countries{`QATAR`},
			currency:  `Qatari Rial`,
			code:      `QAR`,
			number:    `634`,
		},
	}, `RÉUNION`: {
		{
			countries: Countries{`ÅLAND ISLANDS`, `ANDORRA`, `AUSTRIA`, `BELGIUM`, `CYPRUS`, `ESTONIA`, `EUROPEAN UNION`, `FINLAND`, `FRANCE`, `FRENCH GUIANA`, `FRENCH SOUTHERN TERRITORIES (THE)`, `GERMANY`, `GREECE`, `GUADELOUPE`, `HOLY SEE (THE)`, `IRELAND`, `ITALY`, `LATVIA`, `LITHUANIA`, `LUXEMBOURG`, `MALTA`, `MARTINIQUE`, `MAYOTTE`, `MONACO`, `MONTENEGRO`, `NETHERLANDS (THE)`, `PORTUGAL`, `RÉUNION`, `SAINT BARTHÉLEMY`, `SAINT MARTIN (FRENCH PART)`, `SAINT PIERRE AND MIQUELON`, `SAN MARINO`, `SLOVAKIA`, `SLOVENIA`, `SPAIN`, `FRENCH SOUTHERN TERRITORIES`, `HOLY SEE`, `NETHERLANDS`, `SAINT MARTIN`},
			currency:  `Euro`,
			code:      `EUR`,
			number:    `978`,
		},
	}, `ROMANIA`: {
		{
			countries: Countries{`ROMANIA`},
			currency:  `Romanian Leu`,
			code:      `RON`,
			number:    `946`,
		},
	}, `RUSSIAN FEDERATION (THE)`: {
		{
			countries: Countries{`RUSSIAN FEDERATION (THE)`, `RUSSIAN FEDERATION`},
			currency:  `Russian Ruble`,
			code:      `RUB`,
			number:    `643`,
		},
	}, `RWANDA`: {
		{
			countries: Countries{`RWANDA`},
			currency:  `Rwanda Franc`,
			code:      `RWF`,
			number:    `646`,
		},
	}, `SAINT BARTHÉLEMY`: {
		{
			countries: Countries{`ÅLAND ISLANDS`, `ANDORRA`, `AUSTRIA`, `BELGIUM`, `CYPRUS`, `ESTONIA`, `EUROPEAN UNION`, `FINLAND`, `FRANCE`, `FRENCH GUIANA`, `FRENCH SOUTHERN TERRITORIES (THE)`, `GERMANY`, `GREECE`, `GUADELOUPE`, `HOLY SEE (THE)`, `IRELAND`, `ITALY`, `LATVIA`, `LITHUANIA`, `LUXEMBOURG`, `MALTA`, `MARTINIQUE`, `MAYOTTE`, `MONACO`, `MONTENEGRO`, `NETHERLANDS (THE)`, `PORTUGAL`, `RÉUNION`, `SAINT BARTHÉLEMY`, `SAINT MARTIN (FRENCH PART)`, `SAINT PIERRE AND MIQUELON`, `SAN MARINO`, `SLOVAKIA`, `SLOVENIA`, `SPAIN`, `FRENCH SOUTHERN TERRITORIES`, `HOLY SEE`, `NETHERLANDS`, `SAINT MARTIN`},
			currency:  `Euro`,
			code:      `EUR`,
			number:    `978`,
		},
	}, `SAINT HELENA, ASCENSION AND TRISTAN DA CUNHA`: {
		{
			countries: Countries{`SAINT HELENA, ASCENSION AND TRISTAN DA CUNHA`},
			currency:  `Saint Helena Pound`,
			code:      `SHP`,
			number:    `654`,
		},
	}, `SAINT KITTS AND NEVIS`: {
		{
			countries: Countries{`ANGUILLA`, `ANTIGUA AND BARBUDA`, `DOMINICA`, `GRENADA`, `MONTSERRAT`, `SAINT KITTS AND NEVIS`, `SAINT LUCIA`, `SAINT VINCENT AND THE GRENADINES`},
			currency:  `East Caribbean Dollar`,
			code:      `XCD`,
			number:    `951`,
		},
	}, `SAINT LUCIA`: {
		{
			countries: Countries{`ANGUILLA`, `ANTIGUA AND BARBUDA`, `DOMINICA`, `GRENADA`, `MONTSERRAT`, `SAINT KITTS AND NEVIS`, `SAINT LUCIA`, `SAINT VINCENT AND THE GRENADINES`},
			currency:  `East Caribbean Dollar`,
			code:      `XCD`,
			number:    `951`,
		},
	}, `SAINT MARTIN (FRENCH PART)`: {
		{
			countries: Countries{`ÅLAND ISLANDS`, `ANDORRA`, `AUSTRIA`, `BELGIUM`, `CYPRUS`, `ESTONIA`, `EUROPEAN UNION`, `FINLAND`, `FRANCE`, `FRENCH GUIANA`, `FRENCH SOUTHERN TERRITORIES (THE)`, `GERMANY`, `GREECE`, `GUADELOUPE`, `HOLY SEE (THE)`, `IRELAND`, `ITALY`, `LATVIA`, `LITHUANIA`, `LUXEMBOURG`, `MALTA`, `MARTINIQUE`, `MAYOTTE`, `MONACO`, `MONTENEGRO`, `NETHERLANDS (THE)`, `PORTUGAL`, `RÉUNION`, `SAINT BARTHÉLEMY`, `SAINT MARTIN (FRENCH PART)`, `SAINT PIERRE AND MIQUELON`, `SAN MARINO`, `SLOVAKIA`, `SLOVENIA`, `SPAIN`, `FRENCH SOUTHERN TERRITORIES`, `HOLY SEE`, `NETHERLANDS`, `SAINT MARTIN`},
			currency:  `Euro`,
			code:      `EUR`,
			number:    `978`,
		},
	}, `SAINT PIERRE AND MIQUELON`: {
		{
			countries: Countries{`ÅLAND ISLANDS`, `ANDORRA`, `AUSTRIA`, `BELGIUM`, `CYPRUS`, `ESTONIA`, `EUROPEAN UNION`, `FINLAND`, `FRANCE`, `FRENCH GUIANA`, `FRENCH SOUTHERN TERRITORIES (THE)`, `GERMANY`, `GREECE`, `GUADELOUPE`, `HOLY SEE (THE)`, `IRELAND`, `ITALY`, `LATVIA`, `LITHUANIA`, `LUXEMBOURG`, `MALTA`, `MARTINIQUE`, `MAYOTTE`, `MONACO`, `MONTENEGRO`, `NETHERLANDS (THE)`, `PORTUGAL`, `RÉUNION`, `SAINT BARTHÉLEMY`, `SAINT MARTIN (FRENCH PART)`, `SAINT PIERRE AND MIQUELON`, `SAN MARINO`, `SLOVAKIA`, `SLOVENIA`, `SPAIN`, `FRENCH SOUTHERN TERRITORIES`, `HOLY SEE`, `NETHERLANDS`, `SAINT MARTIN`},
			currency:  `Euro`,
			code:      `EUR`,
			number:    `978`,
		},
	}, `SAINT VINCENT AND THE GRENADINES`: {
		{
			countries: Countries{`ANGUILLA`, `ANTIGUA AND BARBUDA`, `DOMINICA`, `GRENADA`, `MONTSERRAT`, `SAINT KITTS AND NEVIS`, `SAINT LUCIA`, `SAINT VINCENT AND THE GRENADINES`},
			currency:  `East Caribbean Dollar`,
			code:      `XCD`,
			number:    `951`,
		},
	}, `SAMOA`: {
		{
			countries: Countries{`SAMOA`},
			currency:  `Tala`,
			code:      `WST`,
			number:    `882`,
		},
	}, `SAN MARINO`: {
		{
			countries: Countries{`ÅLAND ISLANDS`, `ANDORRA`, `AUSTRIA`, `BELGIUM`, `CYPRUS`, `ESTONIA`, `EUROPEAN UNION`, `FINLAND`, `FRANCE`, `FRENCH GUIANA`, `FRENCH SOUTHERN TERRITORIES (THE)`, `GERMANY`, `GREECE`, `GUADELOUPE`, `HOLY SEE (THE)`, `IRELAND`, `ITALY`, `LATVIA`, `LITHUANIA`, `LUXEMBOURG`, `MALTA`, `MARTINIQUE`, `MAYOTTE`, `MONACO`, `MONTENEGRO`, `NETHERLANDS (THE)`, `PORTUGAL`, `RÉUNION`, `SAINT BARTHÉLEMY`, `SAINT MARTIN (FRENCH PART)`, `SAINT PIERRE AND MIQUELON`, `SAN MARINO`, `SLOVAKIA`, `SLOVENIA`, `SPAIN`, `FRENCH SOUTHERN TERRITORIES`, `HOLY SEE`, `NETHERLANDS`, `SAINT MARTIN`},
			currency:  `Euro`,
			code:      `EUR`,
			number:    `978`,
		},
	}, `SAO TOME AND PRINCIPE`: {
		{
			countries: Countries{`SAO TOME AND PRINCIPE`},
			currency:  `Dobra`,
			code:      `STN`,
			number:    `930`,
		},
	}, `SAUDI ARABIA`: {
		{
			countries: Countries{`SAUDI ARABIA`},
			currency:  `Saudi Riyal`,
			code:      `SAR`,
			number:    `682`,
		},
	}, `SENEGAL`: {
		{
			countries: Countries{`BENIN`, `BURKINA FASO`, `CÔTE D'IVOIRE`, `GUINEA-BISSAU`, `MALI`, `NIGER (THE)`, `SENEGAL`, `TOGO`, `NIGER`},
			currency:  `CFA Franc BCEAO`,
			code:      `XOF`,
			number:    `952`,
		},
	}, `SERBIA`: {
		{
			countries: Countries{`SERBIA`},
			currency:  `Serbian Dinar`,
			code:      `RSD`,
			number:    `941`,
		},
	}, `SEYCHELLES`: {
		{
			countries: Countries{`SEYCHELLES`},
			currency:  `Seychelles Rupee`,
			code:      `SCR`,
			number:    `690`,
		},
	}, `SIERRA LEONE`: {
		{
			countries: Countries{`SIERRA LEONE`},
			currency:  `Leone`,
			code:      `SLL`,
			number:    `694`,
		},
	}, `SINGAPORE`: {
		{
			countries: Countries{`SINGAPORE`},
			currency:  `Singapore Dollar`,
			code:      `SGD`,
			number:    `702`,
		},
	}, `SINT MAARTEN (DUTCH PART)`: {
		{
			countries: Countries{`CURAÇAO`, `SINT MAARTEN (DUTCH PART)`, `SINT MAARTEN`},
			currency:  `Netherlands Antillean Guilder`,
			code:      `ANG`,
			number:    `532`,
		},
	}, `SISTEMA UNITARIO DE COMPENSACION REGIONAL DE PAGOS "SUCRE"`: {
		{
			countries: Countries{`SISTEMA UNITARIO DE COMPENSACION REGIONAL DE PAGOS "SUCRE"`},
			currency:  `Sucre`,
			code:      `XSU`,
			number:    `994`,
		},
	}, `SLOVAKIA`: {
		{
			countries: Countries{`ÅLAND ISLANDS`, `ANDORRA`, `AUSTRIA`, `BELGIUM`, `CYPRUS`, `ESTONIA`, `EUROPEAN UNION`, `FINLAND`, `FRANCE`, `FRENCH GUIANA`, `FRENCH SOUTHERN TERRITORIES (THE)`, `GERMANY`, `GREECE`, `GUADELOUPE`, `HOLY SEE (THE)`, `IRELAND`, `ITALY`, `LATVIA`, `LITHUANIA`, `LUXEMBOURG`, `MALTA`, `MARTINIQUE`, `MAYOTTE`, `MONACO`, `MONTENEGRO`, `NETHERLANDS (THE)`, `PORTUGAL`, `RÉUNION`, `SAINT BARTHÉLEMY`, `SAINT MARTIN (FRENCH PART)`, `SAINT PIERRE AND MIQUELON`, `SAN MARINO`, `SLOVAKIA`, `SLOVENIA`, `SPAIN`, `FRENCH SOUTHERN TERRITORIES`, `HOLY SEE`, `NETHERLANDS`, `SAINT MARTIN`},
			currency:  `Euro`,
			code:      `EUR`,
			number:    `978`,
		},
	}, `SLOVENIA`: {
		{
			countries: Countries{`ÅLAND ISLANDS`, `ANDORRA`, `AUSTRIA`, `BELGIUM`, `CYPRUS`, `ESTONIA`, `EUROPEAN UNION`, `FINLAND`, `FRANCE`, `FRENCH GUIANA`, `FRENCH SOUTHERN TERRITORIES (THE)`, `GERMANY`, `GREECE`, `GUADELOUPE`, `HOLY SEE (THE)`, `IRELAND`, `ITALY`, `LATVIA`, `LITHUANIA`, `LUXEMBOURG`, `MALTA`, `MARTINIQUE`, `MAYOTTE`, `MONACO`, `MONTENEGRO`, `NETHERLANDS (THE)`, `PORTUGAL`, `RÉUNION`, `SAINT BARTHÉLEMY`, `SAINT MARTIN (FRENCH PART)`, `SAINT PIERRE AND MIQUELON`, `SAN MARINO`, `SLOVAKIA`, `SLOVENIA`, `SPAIN`, `FRENCH SOUTHERN TERRITORIES`, `HOLY SEE`, `NETHERLANDS`, `SAINT MARTIN`},
			currency:  `Euro`,
			code:      `EUR`,
			number:    `978`,
		},
	}, `SOLOMON ISLANDS`: {
		{
			countries: Countries{`SOLOMON ISLANDS`},
			currency:  `Solomon Islands Dollar`,
			code:      `SBD`,
			number:    `090`,
		},
	}, `SOMALIA`: {
		{
			countries: Countries{`SOMALIA`},
			currency:  `Somali Shilling`,
			code:      `SOS`,
			number:    `706`,
		},
	}, `SOUTH AFRICA`: {
		{
			countries: Countries{`LESOTHO`, `NAMIBIA`, `SOUTH AFRICA`},
			currency:  `Rand`,
			code:      `ZAR`,
			number:    `710`,
		},
	}, `SOUTH SUDAN`: {
		{
			countries: Countries{`SOUTH SUDAN`},
			currency:  `South Sudanese Pound`,
			code:      `SSP`,
			number:    `728`,
		},
	}, `SPAIN`: {
		{
			countries: Countries{`ÅLAND ISLANDS`, `ANDORRA`, `AUSTRIA`, `BELGIUM`, `CYPRUS`, `ESTONIA`, `EUROPEAN UNION`, `FINLAND`, `FRANCE`, `FRENCH GUIANA`, `FRENCH SOUTHERN TERRITORIES (THE)`, `GERMANY`, `GREECE`, `GUADELOUPE`, `HOLY SEE (THE)`, `IRELAND`, `ITALY`, `LATVIA`, `LITHUANIA`, `LUXEMBOURG`, `MALTA`, `MARTINIQUE`, `MAYOTTE`, `MONACO`, `MONTENEGRO`, `NETHERLANDS (THE)`, `PORTUGAL`, `RÉUNION`, `SAINT BARTHÉLEMY`, `SAINT MARTIN (FRENCH PART)`, `SAINT PIERRE AND MIQUELON`, `SAN MARINO`, `SLOVAKIA`, `SLOVENIA`, `SPAIN`, `FRENCH SOUTHERN TERRITORIES`, `HOLY SEE`, `NETHERLANDS`, `SAINT MARTIN`},
			currency:  `Euro`,
			code:      `EUR`,
			number:    `978`,
		},
	}, `SRI LANKA`: {
		{
			countries: Countries{`SRI LANKA`},
			currency:  `Sri Lanka Rupee`,
			code:      `LKR`,
			number:    `144`,
		},
	}, `SUDAN (THE)`: {
		{
			countries: Countries{`SUDAN (THE)`, `SUDAN`},
			currency:  `Sudanese Pound`,
			code:      `SDG`,
			number:    `938`,
		},
	}, `SURINAME`: {
		{
			countries: Countries{`SURINAME`},
			currency:  `Surinam Dollar`,
			code:      `SRD`,
			number:    `968`,
		},
	}, `SVALBARD AND JAN MAYEN`: {
		{
			countries: Countries{`BOUVET ISLAND`, `NORWAY`, `SVALBARD AND JAN MAYEN`},
			currency:  `Norwegian Krone`,
			code:      `NOK`,
			number:    `578`,
		},
	}, `SWEDEN`: {
		{
			countries: Countries{`SWEDEN`},
			currency:  `Swedish Krona`,
			code:      `SEK`,
			number:    `752`,
		},
	}, `SWITZERLAND`: {
		{
			countries: Countries{`LIECHTENSTEIN`, `SWITZERLAND`},
			currency:  `Swiss Franc`,
			code:      `CHF`,
			number:    `756`,
		}, {
			countries: Countries{`SWITZERLAND`},
			currency:  `WIR Euro`,
			code:      `CHE`,
			number:    `947`,
		}, {
			countries: Countries{`SWITZERLAND`},
			currency:  `WIR Franc`,
			code:      `CHW`,
			number:    `948`,
		},
	}, `SYRIAN ARAB REPUBLIC`: {
		{
			countries: Countries{`SYRIAN ARAB REPUBLIC`},
			currency:  `Syrian Pound`,
			code:      `SYP`,
			number:    `760`,
		},
	}, `TAIWAN (PROVINCE OF CHINA)`: {
		{
			countries: Countries{`TAIWAN (PROVINCE OF CHINA)`, `TAIWAN`},
			currency:  `New Taiwan Dollar`,
			code:      `TWD`,
			number:    `901`,
		},
	}, `TAJIKISTAN`: {
		{
			countries: Countries{`TAJIKISTAN`},
			currency:  `Somoni`,
			code:      `TJS`,
			number:    `972`,
		},
	}, `TANZANIA, UNITED REPUBLIC OF`: {
		{
			countries: Countries{`TANZANIA, UNITED REPUBLIC OF`},
			currency:  `Tanzanian Shilling`,
			code:      `TZS`,
			number:    `834`,
		},
	}, `THAILAND`: {
		{
			countries: Countries{`THAILAND`},
			currency:  `Baht`,
			code:      `THB`,
			number:    `764`,
		},
	}, `TIMOR-LESTE`: {
		{
			countries: Countries{`AMERICAN SAMOA`, `BONAIRE, SINT EUSTATIUS AND SABA`, `BRITISH INDIAN OCEAN TERRITORY (THE)`, `ECUADOR`, `EL SALVADOR`, `GUAM`, `HAITI`, `MARSHALL ISLANDS (THE)`, `MICRONESIA (FEDERATED STATES OF)`, `NORTHERN MARIANA ISLANDS (THE)`, `PALAU`, `PANAMA`, `PUERTO RICO`, `TIMOR-LESTE`, `TURKS AND CAICOS ISLANDS (THE)`, `UNITED STATES MINOR OUTLYING ISLANDS (THE)`, `UNITED STATES OF AMERICA (THE)`, `VIRGIN ISLANDS (BRITISH)`, `VIRGIN ISLANDS (U.S.)`, `BRITISH INDIAN OCEAN TERRITORY`, `MARSHALL ISLANDS`, `MICRONESIA`, `NORTHERN MARIANA ISLANDS`, `TURKS AND CAICOS ISLANDS`, `UNITED STATES MINOR OUTLYING ISLANDS`, `UNITED STATES OF AMERICA`, `VIRGIN ISLANDS`, `VIRGIN ISLANDS`},
			currency:  `US Dollar`,
			code:      `USD`,
			number:    `840`,
		},
	}, `TOGO`: {
		{
			countries: Countries{`BENIN`, `BURKINA FASO`, `CÔTE D'IVOIRE`, `GUINEA-BISSAU`, `MALI`, `NIGER (THE)`, `SENEGAL`, `TOGO`, `NIGER`},
			currency:  `CFA Franc BCEAO`,
			code:      `XOF`,
			number:    `952`,
		},
	}, `TOKELAU`: {
		{
			countries: Countries{`COOK ISLANDS (THE)`, `NEW ZEALAND`, `NIUE`, `PITCAIRN`, `TOKELAU`, `COOK ISLANDS`},
			currency:  `New Zealand Dollar`,
			code:      `NZD`,
			number:    `554`,
		},
	}, `TONGA`: {
		{
			countries: Countries{`TONGA`},
			currency:  `Pa’anga`,
			code:      `TOP`,
			number:    `776`,
		},
	}, `TRINIDAD AND TOBAGO`: {
		{
			countries: Countries{`TRINIDAD AND TOBAGO`},
			currency:  `Trinidad and Tobago Dollar`,
			code:      `TTD`,
			number:    `780`,
		},
	}, `TUNISIA`: {
		{
			countries: Countries{`TUNISIA`},
			currency:  `Tunisian Dinar`,
			code:      `TND`,
			number:    `788`,
		},
	}, `TURKEY`: {
		{
			countries: Countries{`TURKEY`},
			currency:  `Turkish Lira`,
			code:      `TRY`,
			number:    `949`,
		},
	}, `TURKMENISTAN`: {
		{
			countries: Countries{`TURKMENISTAN`},
			currency:  `Turkmenistan New Manat`,
			code:      `TMT`,
			number:    `934`,
		},
	}, `TURKS AND CAICOS ISLANDS (THE)`: {
		{
			countries: Countries{`AMERICAN SAMOA`, `BONAIRE, SINT EUSTATIUS AND SABA`, `BRITISH INDIAN OCEAN TERRITORY (THE)`, `ECUADOR`, `EL SALVADOR`, `GUAM`, `HAITI`, `MARSHALL ISLANDS (THE)`, `MICRONESIA (FEDERATED STATES OF)`, `NORTHERN MARIANA ISLANDS (THE)`, `PALAU`, `PANAMA`, `PUERTO RICO`, `TIMOR-LESTE`, `TURKS AND CAICOS ISLANDS (THE)`, `UNITED STATES MINOR OUTLYING ISLANDS (THE)`, `UNITED STATES OF AMERICA (THE)`, `VIRGIN ISLANDS (BRITISH)`, `VIRGIN ISLANDS (U.S.)`, `BRITISH INDIAN OCEAN TERRITORY`, `MARSHALL ISLANDS`, `MICRONESIA`, `NORTHERN MARIANA ISLANDS`, `TURKS AND CAICOS ISLANDS`, `UNITED STATES MINOR OUTLYING ISLANDS`, `UNITED STATES OF AMERICA`, `VIRGIN ISLANDS`, `VIRGIN ISLANDS`},
			currency:  `US Dollar`,
			code:      `USD`,
			number:    `840`,
		},
	}, `TUVALU`: {
		{
			countries: Countries{`AUSTRALIA`, `CHRISTMAS ISLAND`, `COCOS (KEELING) ISLANDS (THE)`, `HEARD ISLAND AND McDONALD ISLANDS`, `KIRIBATI`, `NAURU`, `NORFOLK ISLAND`, `TUVALU`, `COCOS ISLANDS`},
			currency:  `Australian Dollar`,
			code:      `AUD`,
			number:    `036`,
		},
	}, `UGANDA`: {
		{
			countries: Countries{`UGANDA`},
			currency:  `Uganda Shilling`,
			code:      `UGX`,
			number:    `800`,
		},
	}, `UKRAINE`: {
		{
			countries: Countries{`UKRAINE`},
			currency:  `Hryvnia`,
			code:      `UAH`,
			number:    `980`,
		},
	}, `UNITED ARAB EMIRATES (THE)`: {
		{
			countries: Countries{`UNITED ARAB EMIRATES (THE)`, `UNITED ARAB EMIRATES`},
			currency:  `UAE Dirham`,
			code:      `AED`,
			number:    `784`,
		},
	}, `UNITED KINGDOM OF GREAT BRITAIN AND NORTHERN IRELAND (THE)`: {
		{
			countries: Countries{`GUERNSEY`, `ISLE OF MAN`, `JERSEY`, `UNITED KINGDOM OF GREAT BRITAIN AND NORTHERN IRELAND (THE)`, `UNITED KINGDOM OF GREAT BRITAIN AND NORTHERN IRELAND`},
			currency:  `Pound Sterling`,
			code:      `GBP`,
			number:    `826`,
		},
	}, `UNITED STATES MINOR OUTLYING ISLANDS (THE)`: {
		{
			countries: Countries{`AMERICAN SAMOA`, `BONAIRE, SINT EUSTATIUS AND SABA`, `BRITISH INDIAN OCEAN TERRITORY (THE)`, `ECUADOR`, `EL SALVADOR`, `GUAM`, `HAITI`, `MARSHALL ISLANDS (THE)`, `MICRONESIA (FEDERATED STATES OF)`, `NORTHERN MARIANA ISLANDS (THE)`, `PALAU`, `PANAMA`, `PUERTO RICO`, `TIMOR-LESTE`, `TURKS AND CAICOS ISLANDS (THE)`, `UNITED STATES MINOR OUTLYING ISLANDS (THE)`, `UNITED STATES OF AMERICA (THE)`, `VIRGIN ISLANDS (BRITISH)`, `VIRGIN ISLANDS (U.S.)`, `BRITISH INDIAN OCEAN TERRITORY`, `MARSHALL ISLANDS`, `MICRONESIA`, `NORTHERN MARIANA ISLANDS`, `TURKS AND CAICOS ISLANDS`, `UNITED STATES MINOR OUTLYING ISLANDS`, `UNITED STATES OF AMERICA`, `VIRGIN ISLANDS`, `VIRGIN ISLANDS`},
			currency:  `US Dollar`,
			code:      `USD`,
			number:    `840`,
		},
	}, `UNITED STATES OF AMERICA (THE)`: {
		{
			countries: Countries{`AMERICAN SAMOA`, `BONAIRE, SINT EUSTATIUS AND SABA`, `BRITISH INDIAN OCEAN TERRITORY (THE)`, `ECUADOR`, `EL SALVADOR`, `GUAM`, `HAITI`, `MARSHALL ISLANDS (THE)`, `MICRONESIA (FEDERATED STATES OF)`, `NORTHERN MARIANA ISLANDS (THE)`, `PALAU`, `PANAMA`, `PUERTO RICO`, `TIMOR-LESTE`, `TURKS AND CAICOS ISLANDS (THE)`, `UNITED STATES MINOR OUTLYING ISLANDS (THE)`, `UNITED STATES OF AMERICA (THE)`, `VIRGIN ISLANDS (BRITISH)`, `VIRGIN ISLANDS (U.S.)`, `BRITISH INDIAN OCEAN TERRITORY`, `MARSHALL ISLANDS`, `MICRONESIA`, `NORTHERN MARIANA ISLANDS`, `TURKS AND CAICOS ISLANDS`, `UNITED STATES MINOR OUTLYING ISLANDS`, `UNITED STATES OF AMERICA`, `VIRGIN ISLANDS`, `VIRGIN ISLANDS`},
			currency:  `US Dollar`,
			code:      `USD`,
			number:    `840`,
		}, {
			countries: Countries{`UNITED STATES OF AMERICA (THE)`, `UNITED STATES OF AMERICA`},
			currency:  `US Dollar (Next day)`,
			code:      `USN`,
			number:    `997`,
		},
	}, `URUGUAY`: {
		{
			countries: Countries{`URUGUAY`},
			currency:  `Peso Uruguayo`,
			code:      `UYU`,
			number:    `858`,
		}, {
			countries: Countries{`URUGUAY`},
			currency:  `Uruguay Peso en Unidades Indexadas (UI)`,
			code:      `UYI`,
			number:    `940`,
		}, {
			countries: Countries{`URUGUAY`},
			currency:  `Unidad Previsional`,
			code:      `UYW`,
			number:    `927`,
		},
	}, `UZBEKISTAN`: {
		{
			countries: Countries{`UZBEKISTAN`},
			currency:  `Uzbekistan Sum`,
			code:      `UZS`,
			number:    `860`,
		},
	}, `VANUATU`: {
		{
			countries: Countries{`VANUATU`},
			currency:  `Vatu`,
			code:      `VUV`,
			number:    `548`,
		},
	}, `VENEZUELA (BOLIVARIAN REPUBLIC OF)`: {
		{
			countries: Countries{`VENEZUELA (BOLIVARIAN REPUBLIC OF)`, `VENEZUELA`},
			currency:  `Bolívar Soberano`,
			code:      `VES`,
			number:    `928`,
		},
	}, `VIET NAM`: {
		{
			countries: Countries{`VIET NAM`},
			currency:  `Dong`,
			code:      `VND`,
			number:    `704`,
		},
	}, `VIRGIN ISLANDS (BRITISH)`: {
		{
			countries: Countries{`AMERICAN SAMOA`, `BONAIRE, SINT EUSTATIUS AND SABA`, `BRITISH INDIAN OCEAN TERRITORY (THE)`, `ECUADOR`, `EL SALVADOR`, `GUAM`, `HAITI`, `MARSHALL ISLANDS (THE)`, `MICRONESIA (FEDERATED STATES OF)`, `NORTHERN MARIANA ISLANDS (THE)`, `PALAU`, `PANAMA`, `PUERTO RICO`, `TIMOR-LESTE`, `TURKS AND CAICOS ISLANDS (THE)`, `UNITED STATES MINOR OUTLYING ISLANDS (THE)`, `UNITED STATES OF AMERICA (THE)`, `VIRGIN ISLANDS (BRITISH)`, `VIRGIN ISLANDS (U.S.)`, `BRITISH INDIAN OCEAN TERRITORY`, `MARSHALL ISLANDS`, `MICRONESIA`, `NORTHERN MARIANA ISLANDS`, `TURKS AND CAICOS ISLANDS`, `UNITED STATES MINOR OUTLYING ISLANDS`, `UNITED STATES OF AMERICA`, `VIRGIN ISLANDS`, `VIRGIN ISLANDS`},
			currency:  `US Dollar`,
			code:      `USD`,
			number:    `840`,
		},
	}, `VIRGIN ISLANDS (U.S.)`: {
		{
			countries: Countries{`AMERICAN SAMOA`, `BONAIRE, SINT EUSTATIUS AND SABA`, `BRITISH INDIAN OCEAN TERRITORY (THE)`, `ECUADOR`, `EL SALVADOR`, `GUAM`, `HAITI`, `MARSHALL ISLANDS (THE)`, `MICRONESIA (FEDERATED STATES OF)`, `NORTHERN MARIANA ISLANDS (THE)`, `PALAU`, `PANAMA`, `PUERTO RICO`, `TIMOR-LESTE`, `TURKS AND CAICOS ISLANDS (THE)`, `UNITED STATES MINOR OUTLYING ISLANDS (THE)`, `UNITED STATES OF AMERICA (THE)`, `VIRGIN ISLANDS (BRITISH)`, `VIRGIN ISLANDS (U.S.)`, `BRITISH INDIAN OCEAN TERRITORY`, `MARSHALL ISLANDS`, `MICRONESIA`, `NORTHERN MARIANA ISLANDS`, `TURKS AND CAICOS ISLANDS`, `UNITED STATES MINOR OUTLYING ISLANDS`, `UNITED STATES OF AMERICA`, `VIRGIN ISLANDS`, `VIRGIN ISLANDS`},
			currency:  `US Dollar`,
			code:      `USD`,
			number:    `840`,
		},
	}, `WALLIS AND FUTUNA`: {
		{
			countries: Countries{`FRENCH POLYNESIA`, `NEW CALEDONIA`, `WALLIS AND FUTUNA`},
			currency:  `CFP Franc`,
			code:      `XPF`,
			number:    `953`,
		},
	}, `WESTERN SAHARA`: {
		{
			countries: Countries{`MOROCCO`, `WESTERN SAHARA`},
			currency:  `Moroccan Dirham`,
			code:      `MAD`,
			number:    `504`,
		},
	}, `YEMEN`: {
		{
			countries: Countries{`YEMEN`},
			currency:  `Yemeni Rial`,
			code:      `YER`,
			number:    `886`,
		},
	}, `ZAMBIA`: {
		{
			countries: Countries{`ZAMBIA`},
			currency:  `Zambian Kwacha`,
			code:      `ZMW`,
			number:    `967`,
		},
	}, `ZIMBABWE`: {
		{
			countries: Countries{`ZIMBABWE`},
			currency:  `Zimbabwe Dollar`,
			code:      `ZWL`,
			number:    `932`,
		},
	}, `BAHAMAS`: {
		{
			countries: Countries{`BAHAMAS (THE)`, `BAHAMAS`},
			currency:  `Bahamian Dollar`,
			code:      `BSD`,
			number:    `044`,
		},
	}, `BOLIVIA`: {
		{
			countries: Countries{`BOLIVIA (PLURINATIONAL STATE OF)`, `BOLIVIA`},
			currency:  `Boliviano`,
			code:      `BOB`,
			number:    `068`,
		}, {
			countries: Countries{`BOLIVIA (PLURINATIONAL STATE OF)`, `BOLIVIA`},
			currency:  `Mvdol`,
			code:      `BOV`,
			number:    `984`,
		},
	}, `BRITISH INDIAN OCEAN TERRITORY`: {
		{
			countries: Countries{`AMERICAN SAMOA`, `BONAIRE, SINT EUSTATIUS AND SABA`, `BRITISH INDIAN OCEAN TERRITORY (THE)`, `ECUADOR`, `EL SALVADOR`, `GUAM`, `HAITI`, `MARSHALL ISLANDS (THE)`, `MICRONESIA (FEDERATED STATES OF)`, `NORTHERN MARIANA ISLANDS (THE)`, `PALAU`, `PANAMA`, `PUERTO RICO`, `TIMOR-LESTE`, `TURKS AND CAICOS ISLANDS (THE)`, `UNITED STATES MINOR OUTLYING ISLANDS (THE)`, `UNITED STATES OF AMERICA (THE)`, `VIRGIN ISLANDS (BRITISH)`, `VIRGIN ISLANDS (U.S.)`, `BRITISH INDIAN OCEAN TERRITORY`, `MARSHALL ISLANDS`, `MICRONESIA`, `NORTHERN MARIANA ISLANDS`, `TURKS AND CAICOS ISLANDS`, `UNITED STATES MINOR OUTLYING ISLANDS`, `UNITED STATES OF AMERICA`, `VIRGIN ISLANDS`, `VIRGIN ISLANDS`},
			currency:  `US Dollar`,
			code:      `USD`,
			number:    `840`,
		},
	}, `CAYMAN ISLANDS`: {
		{
			countries: Countries{`CAYMAN ISLANDS (THE)`, `CAYMAN ISLANDS`},
			currency:  `Cayman Islands Dollar`,
			code:      `KYD`,
			number:    `136`,
		},
	}, `CENTRAL AFRICAN REPUBLIC`: {
		{
			countries: Countries{`CAMEROON`, `CENTRAL AFRICAN REPUBLIC (THE)`, `CHAD`, `CONGO (THE)`, `EQUATORIAL GUINEA`, `GABON`, `CENTRAL AFRICAN REPUBLIC`, `CONGO`},
			currency:  `CFA Franc BEAC`,
			code:      `XAF`,
			number:    `950`,
		},
	}, `COCOS ISLANDS`: {
		{
			countries: Countries{`AUSTRALIA`, `CHRISTMAS ISLAND`, `COCOS (KEELING) ISLANDS (THE)`, `HEARD ISLAND AND McDONALD ISLANDS`, `KIRIBATI`, `NAURU`, `NORFOLK ISLAND`, `TUVALU`, `COCOS ISLANDS`},
			currency:  `Australian Dollar`,
			code:      `AUD`,
			number:    `036`,
		},
	}, `COMOROS`: {
		{
			countries: Countries{`COMOROS (THE)`, `COMOROS`},
			currency:  `Comorian Franc `,
			code:      `KMF`,
			number:    `174`,
		},
	}, `CONGO`: {
		{
			countries: Countries{`CONGO (THE DEMOCRATIC REPUBLIC OF THE)`, `CONGO`},
			currency:  `Congolese Franc`,
			code:      `CDF`,
			number:    `976`,
		}, {
			countries: Countries{`CAMEROON`, `CENTRAL AFRICAN REPUBLIC (THE)`, `CHAD`, `CONGO (THE)`, `EQUATORIAL GUINEA`, `GABON`, `CENTRAL AFRICAN REPUBLIC`, `CONGO`},
			currency:  `CFA Franc BEAC`,
			code:      `XAF`,
			number:    `950`,
		},
	}, `COOK ISLANDS`: {
		{
			countries: Countries{`COOK ISLANDS (THE)`, `NEW ZEALAND`, `NIUE`, `PITCAIRN`, `TOKELAU`, `COOK ISLANDS`},
			currency:  `New Zealand Dollar`,
			code:      `NZD`,
			number:    `554`,
		},
	}, `DOMINICAN REPUBLIC`: {
		{
			countries: Countries{`DOMINICAN REPUBLIC (THE)`, `DOMINICAN REPUBLIC`},
			currency:  `Dominican Peso`,
			code:      `DOP`,
			number:    `214`,
		},
	}, `FALKLAND ISLANDS [MALVINAS]`: {
		{
			countries: Countries{`FALKLAND ISLANDS (THE) [MALVINAS]`, `FALKLAND ISLANDS [MALVINAS]`},
			currency:  `Falkland Islands Pound`,
			code:      `FKP`,
			number:    `238`,
		},
	}, `FAROE ISLANDS`: {
		{
			countries: Countries{`DENMARK`, `FAROE ISLANDS (THE)`, `GREENLAND`, `FAROE ISLANDS`},
			currency:  `Danish Krone`,
			code:      `DKK`,
			number:    `208`,
		},
	}, `FRENCH SOUTHERN TERRITORIES`: {
		{
			countries: Countries{`ÅLAND ISLANDS`, `ANDORRA`, `AUSTRIA`, `BELGIUM`, `CYPRUS`, `ESTONIA`, `EUROPEAN UNION`, `FINLAND`, `FRANCE`, `FRENCH GUIANA`, `FRENCH SOUTHERN TERRITORIES (THE)`, `GERMANY`, `GREECE`, `GUADELOUPE`, `HOLY SEE (THE)`, `IRELAND`, `ITALY`, `LATVIA`, `LITHUANIA`, `LUXEMBOURG`, `MALTA`, `MARTINIQUE`, `MAYOTTE`, `MONACO`, `MONTENEGRO`, `NETHERLANDS (THE)`, `PORTUGAL`, `RÉUNION`, `SAINT BARTHÉLEMY`, `SAINT MARTIN (FRENCH PART)`, `SAINT PIERRE AND MIQUELON`, `SAN MARINO`, `SLOVAKIA`, `SLOVENIA`, `SPAIN`, `FRENCH SOUTHERN TERRITORIES`, `HOLY SEE`, `NETHERLANDS`, `SAINT MARTIN`},
			currency:  `Euro`,
			code:      `EUR`,
			number:    `978`,
		},
	}, `GAMBIA`: {
		{
			countries: Countries{`GAMBIA (THE)`, `GAMBIA`},
			currency:  `Dalasi`,
			code:      `GMD`,
			number:    `270`,
		},
	}, `HOLY SEE`: {
		{
			countries: Countries{`ÅLAND ISLANDS`, `ANDORRA`, `AUSTRIA`, `BELGIUM`, `CYPRUS`, `ESTONIA`, `EUROPEAN UNION`, `FINLAND`, `FRANCE`, `FRENCH GUIANA`, `FRENCH SOUTHERN TERRITORIES (THE)`, `GERMANY`, `GREECE`, `GUADELOUPE`, `HOLY SEE (THE)`, `IRELAND`, `ITALY`, `LATVIA`, `LITHUANIA`, `LUXEMBOURG`, `MALTA`, `MARTINIQUE`, `MAYOTTE`, `MONACO`, `MONTENEGRO`, `NETHERLANDS (THE)`, `PORTUGAL`, `RÉUNION`, `SAINT BARTHÉLEMY`, `SAINT MARTIN (FRENCH PART)`, `SAINT PIERRE AND MIQUELON`, `SAN MARINO`, `SLOVAKIA`, `SLOVENIA`, `SPAIN`, `FRENCH SOUTHERN TERRITORIES`, `HOLY SEE`, `NETHERLANDS`, `SAINT MARTIN`},
			currency:  `Euro`,
			code:      `EUR`,
			number:    `978`,
		},
	}, `INTERNATIONAL MONETARY FUND `: {
		{
			countries: Countries{`INTERNATIONAL MONETARY FUND (IMF) `, `INTERNATIONAL MONETARY FUND `},
			currency:  `SDR (Special Drawing Right)`,
			code:      `XDR`,
			number:    `960`,
		},
	}, `IRAN`: {
		{
			countries: Countries{`IRAN (ISLAMIC REPUBLIC OF)`, `IRAN`},
			currency:  `Iranian Rial`,
			code:      `IRR`,
			number:    `364`,
		},
	}, `KOREA`: {
		{
			countries: Countries{`KOREA (THE DEMOCRATIC PEOPLE’S REPUBLIC OF)`, `KOREA`},
			currency:  `North Korean Won`,
			code:      `KPW`,
			number:    `408`,
		}, {
			countries: Countries{`KOREA (THE REPUBLIC OF)`, `KOREA`},
			currency:  `Won`,
			code:      `KRW`,
			number:    `410`,
		},
	}, `LAO PEOPLE’S DEMOCRATIC REPUBLIC`: {
		{
			countries: Countries{`LAO PEOPLE’S DEMOCRATIC REPUBLIC (THE)`, `LAO PEOPLE’S DEMOCRATIC REPUBLIC`},
			currency:  `Lao Kip`,
			code:      `LAK`,
			number:    `418`,
		},
	}, `MARSHALL ISLANDS`: {
		{
			countries: Countries{`AMERICAN SAMOA`, `BONAIRE, SINT EUSTATIUS AND SABA`, `BRITISH INDIAN OCEAN TERRITORY (THE)`, `ECUADOR`, `EL SALVADOR`, `GUAM`, `HAITI`, `MARSHALL ISLANDS (THE)`, `MICRONESIA (FEDERATED STATES OF)`, `NORTHERN MARIANA ISLANDS (THE)`, `PALAU`, `PANAMA`, `PUERTO RICO`, `TIMOR-LESTE`, `TURKS AND CAICOS ISLANDS (THE)`, `UNITED STATES MINOR OUTLYING ISLANDS (THE)`, `UNITED STATES OF AMERICA (THE)`, `VIRGIN ISLANDS (BRITISH)`, `VIRGIN ISLANDS (U.S.)`, `BRITISH INDIAN OCEAN TERRITORY`, `MARSHALL ISLANDS`, `MICRONESIA`, `NORTHERN MARIANA ISLANDS`, `TURKS AND CAICOS ISLANDS`, `UNITED STATES MINOR OUTLYING ISLANDS`, `UNITED STATES OF AMERICA`, `VIRGIN ISLANDS`, `VIRGIN ISLANDS`},
			currency:  `US Dollar`,
			code:      `USD`,
			number:    `840`,
		},
	}, `MICRONESIA`: {
		{
			countries: Countries{`AMERICAN SAMOA`, `BONAIRE, SINT EUSTATIUS AND SABA`, `BRITISH INDIAN OCEAN TERRITORY (THE)`, `ECUADOR`, `EL SALVADOR`, `GUAM`, `HAITI`, `MARSHALL ISLANDS (THE)`, `MICRONESIA (FEDERATED STATES OF)`, `NORTHERN MARIANA ISLANDS (THE)`, `PALAU`, `PANAMA`, `PUERTO RICO`, `TIMOR-LESTE`, `TURKS AND CAICOS ISLANDS (THE)`, `UNITED STATES MINOR OUTLYING ISLANDS (THE)`, `UNITED STATES OF AMERICA (THE)`, `VIRGIN ISLANDS (BRITISH)`, `VIRGIN ISLANDS (U.S.)`, `BRITISH INDIAN OCEAN TERRITORY`, `MARSHALL ISLANDS`, `MICRONESIA`, `NORTHERN MARIANA ISLANDS`, `TURKS AND CAICOS ISLANDS`, `UNITED STATES MINOR OUTLYING ISLANDS`, `UNITED STATES OF AMERICA`, `VIRGIN ISLANDS`, `VIRGIN ISLANDS`},
			currency:  `US Dollar`,
			code:      `USD`,
			number:    `840`,
		},
	}, `MOLDOVA`: {
		{
			countries: Countries{`MOLDOVA (THE REPUBLIC OF)`, `MOLDOVA`},
			currency:  `Moldovan Leu`,
			code:      `MDL`,
			number:    `498`,
		},
	}, `NETHERLANDS`: {
		{
			countries: Countries{`ÅLAND ISLANDS`, `ANDORRA`, `AUSTRIA`, `BELGIUM`, `CYPRUS`, `ESTONIA`, `EUROPEAN UNION`, `FINLAND`, `FRANCE`, `FRENCH GUIANA`, `FRENCH SOUTHERN TERRITORIES (THE)`, `GERMANY`, `GREECE`, `GUADELOUPE`, `HOLY SEE (THE)`, `IRELAND`, `ITALY`, `LATVIA`, `LITHUANIA`, `LUXEMBOURG`, `MALTA`, `MARTINIQUE`, `MAYOTTE`, `MONACO`, `MONTENEGRO`, `NETHERLANDS (THE)`, `PORTUGAL`, `RÉUNION`, `SAINT BARTHÉLEMY`, `SAINT MARTIN (FRENCH PART)`, `SAINT PIERRE AND MIQUELON`, `SAN MARINO`, `SLOVAKIA`, `SLOVENIA`, `SPAIN`, `FRENCH SOUTHERN TERRITORIES`, `HOLY SEE`, `NETHERLANDS`, `SAINT MARTIN`},
			currency:  `Euro`,
			code:      `EUR`,
			number:    `978`,
		},
	}, `NIGER`: {
		{
			countries: Countries{`BENIN`, `BURKINA FASO`, `CÔTE D'IVOIRE`, `GUINEA-BISSAU`, `MALI`, `NIGER (THE)`, `SENEGAL`, `TOGO`, `NIGER`},
			currency:  `CFA Franc BCEAO`,
			code:      `XOF`,
			number:    `952`,
		},
	}, `NORTHERN MARIANA ISLANDS`: {
		{
			countries: Countries{`AMERICAN SAMOA`, `BONAIRE, SINT EUSTATIUS AND SABA`, `BRITISH INDIAN OCEAN TERRITORY (THE)`, `ECUADOR`, `EL SALVADOR`, `GUAM`, `HAITI`, `MARSHALL ISLANDS (THE)`, `MICRONESIA (FEDERATED STATES OF)`, `NORTHERN MARIANA ISLANDS (THE)`, `PALAU`, `PANAMA`, `PUERTO RICO`, `TIMOR-LESTE`, `TURKS AND CAICOS ISLANDS (THE)`, `UNITED STATES MINOR OUTLYING ISLANDS (THE)`, `UNITED STATES OF AMERICA (THE)`, `VIRGIN ISLANDS (BRITISH)`, `VIRGIN ISLANDS (U.S.)`, `BRITISH INDIAN OCEAN TERRITORY`, `MARSHALL ISLANDS`, `MICRONESIA`, `NORTHERN MARIANA ISLANDS`, `TURKS AND CAICOS ISLANDS`, `UNITED STATES MINOR OUTLYING ISLANDS`, `UNITED STATES OF AMERICA`, `VIRGIN ISLANDS`, `VIRGIN ISLANDS`},
			currency:  `US Dollar`,
			code:      `USD`,
			number:    `840`,
		},
	}, `PHILIPPINES`: {
		{
			countries: Countries{`PHILIPPINES (THE)`, `PHILIPPINES`},
			currency:  `Philippine Peso`,
			code:      `PHP`,
			number:    `608`,
		},
	}, `RUSSIAN FEDERATION`: {
		{
			countries: Countries{`RUSSIAN FEDERATION (THE)`, `RUSSIAN FEDERATION`},
			currency:  `Russian Ruble`,
			code:      `RUB`,
			number:    `643`,
		},
	}, `SAINT MARTIN`: {
		{
			countries: Countries{`ÅLAND ISLANDS`, `ANDORRA`, `AUSTRIA`, `BELGIUM`, `CYPRUS`, `ESTONIA`, `EUROPEAN UNION`, `FINLAND`, `FRANCE`, `FRENCH GUIANA`, `FRENCH SOUTHERN TERRITORIES (THE)`, `GERMANY`, `GREECE`, `GUADELOUPE`, `HOLY SEE (THE)`, `IRELAND`, `ITALY`, `LATVIA`, `LITHUANIA`, `LUXEMBOURG`, `MALTA`, `MARTINIQUE`, `MAYOTTE`, `MONACO`, `MONTENEGRO`, `NETHERLANDS (THE)`, `PORTUGAL`, `RÉUNION`, `SAINT BARTHÉLEMY`, `SAINT MARTIN (FRENCH PART)`, `SAINT PIERRE AND MIQUELON`, `SAN MARINO`, `SLOVAKIA`, `SLOVENIA`, `SPAIN`, `FRENCH SOUTHERN TERRITORIES`, `HOLY SEE`, `NETHERLANDS`, `SAINT MARTIN`},
			currency:  `Euro`,
			code:      `EUR`,
			number:    `978`,
		},
	}, `SINT MAARTEN`: {
		{
			countries: Countries{`CURAÇAO`, `SINT MAARTEN (DUTCH PART)`, `SINT MAARTEN`},
			currency:  `Netherlands Antillean Guilder`,
			code:      `ANG`,
			number:    `532`,
		},
	}, `SUDAN`: {
		{
			countries: Countries{`SUDAN (THE)`, `SUDAN`},
			currency:  `Sudanese Pound`,
			code:      `SDG`,
			number:    `938`,
		},
	}, `TAIWAN`: {
		{
			countries: Countries{`TAIWAN (PROVINCE OF CHINA)`, `TAIWAN`},
			currency:  `New Taiwan Dollar`,
			code:      `TWD`,
			number:    `901`,
		},
	}, `TURKS AND CAICOS ISLANDS`: {
		{
			countries: Countries{`AMERICAN SAMOA`, `BONAIRE, SINT EUSTATIUS AND SABA`, `BRITISH INDIAN OCEAN TERRITORY (THE)`, `ECUADOR`, `EL SALVADOR`, `GUAM`, `HAITI`, `MARSHALL ISLANDS (THE)`, `MICRONESIA (FEDERATED STATES OF)`, `NORTHERN MARIANA ISLANDS (THE)`, `PALAU`, `PANAMA`, `PUERTO RICO`, `TIMOR-LESTE`, `TURKS AND CAICOS ISLANDS (THE)`, `UNITED STATES MINOR OUTLYING ISLANDS (THE)`, `UNITED STATES OF AMERICA (THE)`, `VIRGIN ISLANDS (BRITISH)`, `VIRGIN ISLANDS (U.S.)`, `BRITISH INDIAN OCEAN TERRITORY`, `MARSHALL ISLANDS`, `MICRONESIA`, `NORTHERN MARIANA ISLANDS`, `TURKS AND CAICOS ISLANDS`, `UNITED STATES MINOR OUTLYING ISLANDS`, `UNITED STATES OF AMERICA`, `VIRGIN ISLANDS`, `VIRGIN ISLANDS`},
			currency:  `US Dollar`,
			code:      `USD`,
			number:    `840`,
		},
	}, `UNITED ARAB EMIRATES`: {
		{
			countries: Countries{`UNITED ARAB EMIRATES (THE)`, `UNITED ARAB EMIRATES`},
			currency:  `UAE Dirham`,
			code:      `AED`,
			number:    `784`,
		},
	}, `UNITED KINGDOM OF GREAT BRITAIN AND NORTHERN IRELAND`: {
		{
			countries: Countries{`GUERNSEY`, `ISLE OF MAN`, `JERSEY`, `UNITED KINGDOM OF GREAT BRITAIN AND NORTHERN IRELAND (THE)`, `UNITED KINGDOM OF GREAT BRITAIN AND NORTHERN IRELAND`},
			currency:  `Pound Sterling`,
			code:      `GBP`,
			number:    `826`,
		},
	}, `UNITED STATES MINOR OUTLYING ISLANDS`: {
		{
			countries: Countries{`AMERICAN SAMOA`, `BONAIRE, SINT EUSTATIUS AND SABA`, `BRITISH INDIAN OCEAN TERRITORY (THE)`, `ECUADOR`, `EL SALVADOR`, `GUAM`, `HAITI`, `MARSHALL ISLANDS (THE)`, `MICRONESIA (FEDERATED STATES OF)`, `NORTHERN MARIANA ISLANDS (THE)`, `PALAU`, `PANAMA`, `PUERTO RICO`, `TIMOR-LESTE`, `TURKS AND CAICOS ISLANDS (THE)`, `UNITED STATES MINOR OUTLYING ISLANDS (THE)`, `UNITED STATES OF AMERICA (THE)`, `VIRGIN ISLANDS (BRITISH)`, `VIRGIN ISLANDS (U.S.)`, `BRITISH INDIAN OCEAN TERRITORY`, `MARSHALL ISLANDS`, `MICRONESIA`, `NORTHERN MARIANA ISLANDS`, `TURKS AND CAICOS ISLANDS`, `UNITED STATES MINOR OUTLYING ISLANDS`, `UNITED STATES OF AMERICA`, `VIRGIN ISLANDS`, `VIRGIN ISLANDS`},
			currency:  `US Dollar`,
			code:      `USD`,
			number:    `840`,
		},
	}, `UNITED STATES OF AMERICA`: {
		{
			countries: Countries{`AMERICAN SAMOA`, `BONAIRE, SINT EUSTATIUS AND SABA`, `BRITISH INDIAN OCEAN TERRITORY (THE)`, `ECUADOR`, `EL SALVADOR`, `GUAM`, `HAITI`, `MARSHALL ISLANDS (THE)`, `MICRONESIA (FEDERATED STATES OF)`, `NORTHERN MARIANA ISLANDS (THE)`, `PALAU`, `PANAMA`, `PUERTO RICO`, `TIMOR-LESTE`, `TURKS AND CAICOS ISLANDS (THE)`, `UNITED STATES MINOR OUTLYING ISLANDS (THE)`, `UNITED STATES OF AMERICA (THE)`, `VIRGIN ISLANDS (BRITISH)`, `VIRGIN ISLANDS (U.S.)`, `BRITISH INDIAN OCEAN TERRITORY`, `MARSHALL ISLANDS`, `MICRONESIA`, `NORTHERN MARIANA ISLANDS`, `TURKS AND CAICOS ISLANDS`, `UNITED STATES MINOR OUTLYING ISLANDS`, `UNITED STATES OF AMERICA`, `VIRGIN ISLANDS`, `VIRGIN ISLANDS`},
			currency:  `US Dollar`,
			code:      `USD`,
			number:    `840`,
		}, {
			countries: Countries{`UNITED STATES OF AMERICA (THE)`, `UNITED STATES OF AMERICA`},
			currency:  `US Dollar (Next day)`,
			code:      `USN`,
			number:    `997`,
		},
	}, `VENEZUELA`: {
		{
			countries: Countries{`VENEZUELA (BOLIVARIAN REPUBLIC OF)`, `VENEZUELA`},
			currency:  `Bolívar Soberano`,
			code:      `VES`,
			number:    `928`,
		},
	}, `VIRGIN ISLANDS`: {
		{
			countries: Countries{`AMERICAN SAMOA`, `BONAIRE, SINT EUSTATIUS AND SABA`, `BRITISH INDIAN OCEAN TERRITORY (THE)`, `ECUADOR`, `EL SALVADOR`, `GUAM`, `HAITI`, `MARSHALL ISLANDS (THE)`, `MICRONESIA (FEDERATED STATES OF)`, `NORTHERN MARIANA ISLANDS (THE)`, `PALAU`, `PANAMA`, `PUERTO RICO`, `TIMOR-LESTE`, `TURKS AND CAICOS ISLANDS (THE)`, `UNITED STATES MINOR OUTLYING ISLANDS (THE)`, `UNITED STATES OF AMERICA (THE)`, `VIRGIN ISLANDS (BRITISH)`, `VIRGIN ISLANDS (U.S.)`, `BRITISH INDIAN OCEAN TERRITORY`, `MARSHALL ISLANDS`, `MICRONESIA`, `NORTHERN MARIANA ISLANDS`, `TURKS AND CAICOS ISLANDS`, `UNITED STATES MINOR OUTLYING ISLANDS`, `UNITED STATES OF AMERICA`, `VIRGIN ISLANDS`, `VIRGIN ISLANDS`},
			currency:  `US Dollar`,
			code:      `USD`,
			number:    `840`,
		}, {
			countries: Countries{`AMERICAN SAMOA`, `BONAIRE, SINT EUSTATIUS AND SABA`, `BRITISH INDIAN OCEAN TERRITORY (THE)`, `ECUADOR`, `EL SALVADOR`, `GUAM`, `HAITI`, `MARSHALL ISLANDS (THE)`, `MICRONESIA (FEDERATED STATES OF)`, `NORTHERN MARIANA ISLANDS (THE)`, `PALAU`, `PANAMA`, `PUERTO RICO`, `TIMOR-LESTE`, `TURKS AND CAICOS ISLANDS (THE)`, `UNITED STATES MINOR OUTLYING ISLANDS (THE)`, `UNITED STATES OF AMERICA (THE)`, `VIRGIN ISLANDS (BRITISH)`, `VIRGIN ISLANDS (U.S.)`, `BRITISH INDIAN OCEAN TERRITORY`, `MARSHALL ISLANDS`, `MICRONESIA`, `NORTHERN MARIANA ISLANDS`, `TURKS AND CAICOS ISLANDS`, `UNITED STATES MINOR OUTLYING ISLANDS`, `UNITED STATES OF AMERICA`, `VIRGIN ISLANDS`, `VIRGIN ISLANDS`},
			currency:  `US Dollar`,
			code:      `USD`,
			number:    `840`,
		},
	},
}

var currenciesByCurrency = map[string]currency{
	`Afghani`: {
		countries: Countries{`AFGHANISTAN`},
		currency:  `Afghani`,
		code:      `AFN`,
		number:    `971`,
	},
	`Euro`: {
		countries: Countries{`ÅLAND ISLANDS`, `ANDORRA`, `AUSTRIA`, `BELGIUM`, `CYPRUS`, `ESTONIA`, `EUROPEAN UNION`, `FINLAND`, `FRANCE`, `FRENCH GUIANA`, `FRENCH SOUTHERN TERRITORIES (THE)`, `GERMANY`, `GREECE`, `GUADELOUPE`, `HOLY SEE (THE)`, `IRELAND`, `ITALY`, `LATVIA`, `LITHUANIA`, `LUXEMBOURG`, `MALTA`, `MARTINIQUE`, `MAYOTTE`, `MONACO`, `MONTENEGRO`, `NETHERLANDS (THE)`, `PORTUGAL`, `RÉUNION`, `SAINT BARTHÉLEMY`, `SAINT MARTIN (FRENCH PART)`, `SAINT PIERRE AND MIQUELON`, `SAN MARINO`, `SLOVAKIA`, `SLOVENIA`, `SPAIN`, `FRENCH SOUTHERN TERRITORIES`, `HOLY SEE`, `NETHERLANDS`, `SAINT MARTIN`},
		currency:  `Euro`,
		code:      `EUR`,
		number:    `978`,
	},
	`Lek`: {
		countries: Countries{`ALBANIA`},
		currency:  `Lek`,
		code:      `ALL`,
		number:    `008`,
	},
	`Algerian Dinar`: {
		countries: Countries{`ALGERIA`},
		currency:  `Algerian Dinar`,
		code:      `DZD`,
		number:    `012`,
	},
	`US Dollar`: {
		countries: Countries{`AMERICAN SAMOA`, `BONAIRE, SINT EUSTATIUS AND SABA`, `BRITISH INDIAN OCEAN TERRITORY (THE)`, `ECUADOR`, `EL SALVADOR`, `GUAM`, `HAITI`, `MARSHALL ISLANDS (THE)`, `MICRONESIA (FEDERATED STATES OF)`, `NORTHERN MARIANA ISLANDS (THE)`, `PALAU`, `PANAMA`, `PUERTO RICO`, `TIMOR-LESTE`, `TURKS AND CAICOS ISLANDS (THE)`, `UNITED STATES MINOR OUTLYING ISLANDS (THE)`, `UNITED STATES OF AMERICA (THE)`, `VIRGIN ISLANDS (BRITISH)`, `VIRGIN ISLANDS (U.S.)`, `BRITISH INDIAN OCEAN TERRITORY`, `MARSHALL ISLANDS`, `MICRONESIA`, `NORTHERN MARIANA ISLANDS`, `TURKS AND CAICOS ISLANDS`, `UNITED STATES MINOR OUTLYING ISLANDS`, `UNITED STATES OF AMERICA`, `VIRGIN ISLANDS`, `VIRGIN ISLANDS`},
		currency:  `US Dollar`,
		code:      `USD`,
		number:    `840`,
	},
	`Kwanza`: {
		countries: Countries{`ANGOLA`},
		currency:  `Kwanza`,
		code:      `AOA`,
		number:    `973`,
	},
	`East Caribbean Dollar`: {
		countries: Countries{`ANGUILLA`, `ANTIGUA AND BARBUDA`, `DOMINICA`, `GRENADA`, `MONTSERRAT`, `SAINT KITTS AND NEVIS`, `SAINT LUCIA`, `SAINT VINCENT AND THE GRENADINES`},
		currency:  `East Caribbean Dollar`,
		code:      `XCD`,
		number:    `951`,
	},
	`Argentine Peso`: {
		countries: Countries{`ARGENTINA`},
		currency:  `Argentine Peso`,
		code:      `ARS`,
		number:    `032`,
	},
	`Armenian Dram`: {
		countries: Countries{`ARMENIA`},
		currency:  `Armenian Dram`,
		code:      `AMD`,
		number:    `051`,
	},
	`Aruban Florin`: {
		countries: Countries{`ARUBA`},
		currency:  `Aruban Florin`,
		code:      `AWG`,
		number:    `533`,
	},
	`Australian Dollar`: {
		countries: Countries{`AUSTRALIA`, `CHRISTMAS ISLAND`, `COCOS (KEELING) ISLANDS (THE)`, `HEARD ISLAND AND McDONALD ISLANDS`, `KIRIBATI`, `NAURU`, `NORFOLK ISLAND`, `TUVALU`, `COCOS ISLANDS`},
		currency:  `Australian Dollar`,
		code:      `AUD`,
		number:    `036`,
	},
	`Azerbaijan Manat`: {
		countries: Countries{`AZERBAIJAN`},
		currency:  `Azerbaijan Manat`,
		code:      `AZN`,
		number:    `944`,
	},
	`Bahamian Dollar`: {
		countries: Countries{`BAHAMAS (THE)`, `BAHAMAS`},
		currency:  `Bahamian Dollar`,
		code:      `BSD`,
		number:    `044`,
	},
	`Bahraini Dinar`: {
		countries: Countries{`BAHRAIN`},
		currency:  `Bahraini Dinar`,
		code:      `BHD`,
		number:    `048`,
	},
	`Taka`: {
		countries: Countries{`BANGLADESH`},
		currency:  `Taka`,
		code:      `BDT`,
		number:    `050`,
	},
	`Barbados Dollar`: {
		countries: Countries{`BARBADOS`},
		currency:  `Barbados Dollar`,
		code:      `BBD`,
		number:    `052`,
	},
	`Belarusian Ruble`: {
		countries: Countries{`BELARUS`},
		currency:  `Belarusian Ruble`,
		code:      `BYN`,
		number:    `933`,
	},
	`Belize Dollar`: {
		countries: Countries{`BELIZE`},
		currency:  `Belize Dollar`,
		code:      `BZD`,
		number:    `084`,
	},
	`CFA Franc BCEAO`: {
		countries: Countries{`BENIN`, `BURKINA FASO`, `CÔTE D'IVOIRE`, `GUINEA-BISSAU`, `MALI`, `NIGER (THE)`, `SENEGAL`, `TOGO`, `NIGER`},
		currency:  `CFA Franc BCEAO`,
		code:      `XOF`,
		number:    `952`,
	},
	`Bermudian Dollar`: {
		countries: Countries{`BERMUDA`},
		currency:  `Bermudian Dollar`,
		code:      `BMD`,
		number:    `060`,
	},
	`Indian Rupee`: {
		countries: Countries{`BHUTAN`, `INDIA`},
		currency:  `Indian Rupee`,
		code:      `INR`,
		number:    `356`,
	},
	`Ngultrum`: {
		countries: Countries{`BHUTAN`},
		currency:  `Ngultrum`,
		code:      `BTN`,
		number:    `064`,
	},
	`Boliviano`: {
		countries: Countries{`BOLIVIA (PLURINATIONAL STATE OF)`, `BOLIVIA`},
		currency:  `Boliviano`,
		code:      `BOB`,
		number:    `068`,
	},
	`Mvdol`: {
		countries: Countries{`BOLIVIA (PLURINATIONAL STATE OF)`, `BOLIVIA`},
		currency:  `Mvdol`,
		code:      `BOV`,
		number:    `984`,
	},
	`Convertible Mark`: {
		countries: Countries{`BOSNIA AND HERZEGOVINA`},
		currency:  `Convertible Mark`,
		code:      `BAM`,
		number:    `977`,
	},
	`Pula`: {
		countries: Countries{`BOTSWANA`},
		currency:  `Pula`,
		code:      `BWP`,
		number:    `072`,
	},
	`Norwegian Krone`: {
		countries: Countries{`BOUVET ISLAND`, `NORWAY`, `SVALBARD AND JAN MAYEN`},
		currency:  `Norwegian Krone`,
		code:      `NOK`,
		number:    `578`,
	},
	`Brazilian Real`: {
		countries: Countries{`BRAZIL`},
		currency:  `Brazilian Real`,
		code:      `BRL`,
		number:    `986`,
	},
	`Brunei Dollar`: {
		countries: Countries{`BRUNEI DARUSSALAM`},
		currency:  `Brunei Dollar`,
		code:      `BND`,
		number:    `096`,
	},
	`Bulgarian Lev`: {
		countries: Countries{`BULGARIA`},
		currency:  `Bulgarian Lev`,
		code:      `BGN`,
		number:    `975`,
	},
	`Burundi Franc`: {
		countries: Countries{`BURUNDI`},
		currency:  `Burundi Franc`,
		code:      `BIF`,
		number:    `108`,
	},
	`Cabo Verde Escudo`: {
		countries: Countries{`CABO VERDE`},
		currency:  `Cabo Verde Escudo`,
		code:      `CVE`,
		number:    `132`,
	},
	`Riel`: {
		countries: Countries{`CAMBODIA`},
		currency:  `Riel`,
		code:      `KHR`,
		number:    `116`,
	},
	`CFA Franc BEAC`: {
		countries: Countries{`CAMEROON`, `CENTRAL AFRICAN REPUBLIC (THE)`, `CHAD`, `CONGO (THE)`, `EQUATORIAL GUINEA`, `GABON`, `CENTRAL AFRICAN REPUBLIC`, `CONGO`},
		currency:  `CFA Franc BEAC`,
		code:      `XAF`,
		number:    `950`,
	},
	`Canadian Dollar`: {
		countries: Countries{`CANADA`},
		currency:  `Canadian Dollar`,
		code:      `CAD`,
		number:    `124`,
	},
	`Cayman Islands Dollar`: {
		countries: Countries{`CAYMAN ISLANDS (THE)`, `CAYMAN ISLANDS`},
		currency:  `Cayman Islands Dollar`,
		code:      `KYD`,
		number:    `136`,
	},
	`Chilean Peso`: {
		countries: Countries{`CHILE`},
		currency:  `Chilean Peso`,
		code:      `CLP`,
		number:    `152`,
	},
	`Unidad de Fomento`: {
		countries: Countries{`CHILE`},
		currency:  `Unidad de Fomento`,
		code:      `CLF`,
		number:    `990`,
	},
	`Yuan Renminbi`: {
		countries: Countries{`CHINA`},
		currency:  `Yuan Renminbi`,
		code:      `CNY`,
		number:    `156`,
	},
	`Colombian Peso`: {
		countries: Countries{`COLOMBIA`},
		currency:  `Colombian Peso`,
		code:      `COP`,
		number:    `170`,
	},
	`Unidad de Valor Real`: {
		countries: Countries{`COLOMBIA`},
		currency:  `Unidad de Valor Real`,
		code:      `COU`,
		number:    `970`,
	},
	`Comorian Franc `: {
		countries: Countries{`COMOROS (THE)`, `COMOROS`},
		currency:  `Comorian Franc `,
		code:      `KMF`,
		number:    `174`,
	},
	`Congolese Franc`: {
		countries: Countries{`CONGO (THE DEMOCRATIC REPUBLIC OF THE)`, `CONGO`},
		currency:  `Congolese Franc`,
		code:      `CDF`,
		number:    `976`,
	},
	`New Zealand Dollar`: {
		countries: Countries{`COOK ISLANDS (THE)`, `NEW ZEALAND`, `NIUE`, `PITCAIRN`, `TOKELAU`, `COOK ISLANDS`},
		currency:  `New Zealand Dollar`,
		code:      `NZD`,
		number:    `554`,
	},
	`Costa Rican Colon`: {
		countries: Countries{`COSTA RICA`},
		currency:  `Costa Rican Colon`,
		code:      `CRC`,
		number:    `188`,
	},
	`Kuna`: {
		countries: Countries{`CROATIA`},
		currency:  `Kuna`,
		code:      `HRK`,
		number:    `191`,
	},
	`Cuban Peso`: {
		countries: Countries{`CUBA`},
		currency:  `Cuban Peso`,
		code:      `CUP`,
		number:    `192`,
	},
	`Peso Convertible`: {
		countries: Countries{`CUBA`},
		currency:  `Peso Convertible`,
		code:      `CUC`,
		number:    `931`,
	},
	`Netherlands Antillean Guilder`: {
		countries: Countries{`CURAÇAO`, `SINT MAARTEN (DUTCH PART)`, `SINT MAARTEN`},
		currency:  `Netherlands Antillean Guilder`,
		code:      `ANG`,
		number:    `532`,
	},
	`Czech Koruna`: {
		countries: Countries{`CZECHIA`},
		currency:  `Czech Koruna`,
		code:      `CZK`,
		number:    `203`,
	},
	`Danish Krone`: {
		countries: Countries{`DENMARK`, `FAROE ISLANDS (THE)`, `GREENLAND`, `FAROE ISLANDS`},
		currency:  `Danish Krone`,
		code:      `DKK`,
		number:    `208`,
	},
	`Djibouti Franc`: {
		countries: Countries{`DJIBOUTI`},
		currency:  `Djibouti Franc`,
		code:      `DJF`,
		number:    `262`,
	},
	`Dominican Peso`: {
		countries: Countries{`DOMINICAN REPUBLIC (THE)`, `DOMINICAN REPUBLIC`},
		currency:  `Dominican Peso`,
		code:      `DOP`,
		number:    `214`,
	},
	`Egyptian Pound`: {
		countries: Countries{`EGYPT`},
		currency:  `Egyptian Pound`,
		code:      `EGP`,
		number:    `818`,
	},
	`El Salvador Colon`: {
		countries: Countries{`EL SALVADOR`},
		currency:  `El Salvador Colon`,
		code:      `SVC`,
		number:    `222`,
	},
	`Nakfa`: {
		countries: Countries{`ERITREA`},
		currency:  `Nakfa`,
		code:      `ERN`,
		number:    `232`,
	},
	`Lilangeni`: {
		countries: Countries{`ESWATINI`},
		currency:  `Lilangeni`,
		code:      `SZL`,
		number:    `748`,
	},
	`Ethiopian Birr`: {
		countries: Countries{`ETHIOPIA`},
		currency:  `Ethiopian Birr`,
		code:      `ETB`,
		number:    `230`,
	},
	`Falkland Islands Pound`: {
		countries: Countries{`FALKLAND ISLANDS (THE) [MALVINAS]`, `FALKLAND ISLANDS [MALVINAS]`},
		currency:  `Falkland Islands Pound`,
		code:      `FKP`,
		number:    `238`,
	},
	`Fiji Dollar`: {
		countries: Countries{`FIJI`},
		currency:  `Fiji Dollar`,
		code:      `FJD`,
		number:    `242`,
	},
	`CFP Franc`: {
		countries: Countries{`FRENCH POLYNESIA`, `NEW CALEDONIA`, `WALLIS AND FUTUNA`},
		currency:  `CFP Franc`,
		code:      `XPF`,
		number:    `953`,
	},
	`Dalasi`: {
		countries: Countries{`GAMBIA (THE)`, `GAMBIA`},
		currency:  `Dalasi`,
		code:      `GMD`,
		number:    `270`,
	},
	`Lari`: {
		countries: Countries{`GEORGIA`},
		currency:  `Lari`,
		code:      `GEL`,
		number:    `981`,
	},
	`Ghana Cedi`: {
		countries: Countries{`GHANA`},
		currency:  `Ghana Cedi`,
		code:      `GHS`,
		number:    `936`,
	},
	`Gibraltar Pound`: {
		countries: Countries{`GIBRALTAR`},
		currency:  `Gibraltar Pound`,
		code:      `GIP`,
		number:    `292`,
	},
	`Quetzal`: {
		countries: Countries{`GUATEMALA`},
		currency:  `Quetzal`,
		code:      `GTQ`,
		number:    `320`,
	},
	`Pound Sterling`: {
		countries: Countries{`GUERNSEY`, `ISLE OF MAN`, `JERSEY`, `UNITED KINGDOM OF GREAT BRITAIN AND NORTHERN IRELAND (THE)`, `UNITED KINGDOM OF GREAT BRITAIN AND NORTHERN IRELAND`},
		currency:  `Pound Sterling`,
		code:      `GBP`,
		number:    `826`,
	},
	`Guinean Franc`: {
		countries: Countries{`GUINEA`},
		currency:  `Guinean Franc`,
		code:      `GNF`,
		number:    `324`,
	},
	`Guyana Dollar`: {
		countries: Countries{`GUYANA`},
		currency:  `Guyana Dollar`,
		code:      `GYD`,
		number:    `328`,
	},
	`Gourde`: {
		countries: Countries{`HAITI`},
		currency:  `Gourde`,
		code:      `HTG`,
		number:    `332`,
	},
	`Lempira`: {
		countries: Countries{`HONDURAS`},
		currency:  `Lempira`,
		code:      `HNL`,
		number:    `340`,
	},
	`Hong Kong Dollar`: {
		countries: Countries{`HONG KONG`},
		currency:  `Hong Kong Dollar`,
		code:      `HKD`,
		number:    `344`,
	},
	`Forint`: {
		countries: Countries{`HUNGARY`},
		currency:  `Forint`,
		code:      `HUF`,
		number:    `348`,
	},
	`Iceland Krona`: {
		countries: Countries{`ICELAND`},
		currency:  `Iceland Krona`,
		code:      `ISK`,
		number:    `352`,
	},
	`Rupiah`: {
		countries: Countries{`INDONESIA`},
		currency:  `Rupiah`,
		code:      `IDR`,
		number:    `360`,
	},
	`SDR (Special Drawing Right)`: {
		countries: Countries{`INTERNATIONAL MONETARY FUND (IMF) `, `INTERNATIONAL MONETARY FUND `},
		currency:  `SDR (Special Drawing Right)`,
		code:      `XDR`,
		number:    `960`,
	},
	`Iranian Rial`: {
		countries: Countries{`IRAN (ISLAMIC REPUBLIC OF)`, `IRAN`},
		currency:  `Iranian Rial`,
		code:      `IRR`,
		number:    `364`,
	},
	`Iraqi Dinar`: {
		countries: Countries{`IRAQ`},
		currency:  `Iraqi Dinar`,
		code:      `IQD`,
		number:    `368`,
	},
	`New Israeli Sheqel`: {
		countries: Countries{`ISRAEL`},
		currency:  `New Israeli Sheqel`,
		code:      `ILS`,
		number:    `376`,
	},
	`Jamaican Dollar`: {
		countries: Countries{`JAMAICA`},
		currency:  `Jamaican Dollar`,
		code:      `JMD`,
		number:    `388`,
	},
	`Yen`: {
		countries: Countries{`JAPAN`},
		currency:  `Yen`,
		code:      `JPY`,
		number:    `392`,
	},
	`Jordanian Dinar`: {
		countries: Countries{`JORDAN`},
		currency:  `Jordanian Dinar`,
		code:      `JOD`,
		number:    `400`,
	},
	`Tenge`: {
		countries: Countries{`KAZAKHSTAN`},
		currency:  `Tenge`,
		code:      `KZT`,
		number:    `398`,
	},
	`Kenyan Shilling`: {
		countries: Countries{`KENYA`},
		currency:  `Kenyan Shilling`,
		code:      `KES`,
		number:    `404`,
	},
	`North Korean Won`: {
		countries: Countries{`KOREA (THE DEMOCRATIC PEOPLE’S REPUBLIC OF)`, `KOREA`},
		currency:  `North Korean Won`,
		code:      `KPW`,
		number:    `408`,
	},
	`Won`: {
		countries: Countries{`KOREA (THE REPUBLIC OF)`, `KOREA`},
		currency:  `Won`,
		code:      `KRW`,
		number:    `410`,
	},
	`Kuwaiti Dinar`: {
		countries: Countries{`KUWAIT`},
		currency:  `Kuwaiti Dinar`,
		code:      `KWD`,
		number:    `414`,
	},
	`Som`: {
		countries: Countries{`KYRGYZSTAN`},
		currency:  `Som`,
		code:      `KGS`,
		number:    `417`,
	},
	`Lao Kip`: {
		countries: Countries{`LAO PEOPLE’S DEMOCRATIC REPUBLIC (THE)`, `LAO PEOPLE’S DEMOCRATIC REPUBLIC`},
		currency:  `Lao Kip`,
		code:      `LAK`,
		number:    `418`,
	},
	`Lebanese Pound`: {
		countries: Countries{`LEBANON`},
		currency:  `Lebanese Pound`,
		code:      `LBP`,
		number:    `422`,
	},
	`Loti`: {
		countries: Countries{`LESOTHO`},
		currency:  `Loti`,
		code:      `LSL`,
		number:    `426`,
	},
	`Rand`: {
		countries: Countries{`LESOTHO`, `NAMIBIA`, `SOUTH AFRICA`},
		currency:  `Rand`,
		code:      `ZAR`,
		number:    `710`,
	},
	`Liberian Dollar`: {
		countries: Countries{`LIBERIA`},
		currency:  `Liberian Dollar`,
		code:      `LRD`,
		number:    `430`,
	},
	`Libyan Dinar`: {
		countries: Countries{`LIBYA`},
		currency:  `Libyan Dinar`,
		code:      `LYD`,
		number:    `434`,
	},
	`Swiss Franc`: {
		countries: Countries{`LIECHTENSTEIN`, `SWITZERLAND`},
		currency:  `Swiss Franc`,
		code:      `CHF`,
		number:    `756`,
	},
	`Pataca`: {
		countries: Countries{`MACAO`},
		currency:  `Pataca`,
		code:      `MOP`,
		number:    `446`,
	},
	`Denar`: {
		countries: Countries{`NORTH MACEDONIA`},
		currency:  `Denar`,
		code:      `MKD`,
		number:    `807`,
	},
	`Malagasy Ariary`: {
		countries: Countries{`MADAGASCAR`},
		currency:  `Malagasy Ariary`,
		code:      `MGA`,
		number:    `969`,
	},
	`Malawi Kwacha`: {
		countries: Countries{`MALAWI`},
		currency:  `Malawi Kwacha`,
		code:      `MWK`,
		number:    `454`,
	},
	`Malaysian Ringgit`: {
		countries: Countries{`MALAYSIA`},
		currency:  `Malaysian Ringgit`,
		code:      `MYR`,
		number:    `458`,
	},
	`Rufiyaa`: {
		countries: Countries{`MALDIVES`},
		currency:  `Rufiyaa`,
		code:      `MVR`,
		number:    `462`,
	},
	`Ouguiya`: {
		countries: Countries{`MAURITANIA`},
		currency:  `Ouguiya`,
		code:      `MRU`,
		number:    `929`,
	},
	`Mauritius Rupee`: {
		countries: Countries{`MAURITIUS`},
		currency:  `Mauritius Rupee`,
		code:      `MUR`,
		number:    `480`,
	},
	`ADB Unit of Account`: {
		countries: Countries{`MEMBER COUNTRIES OF THE AFRICAN DEVELOPMENT BANK GROUP`},
		currency:  `ADB Unit of Account`,
		code:      `XUA`,
		number:    `965`,
	},
	`Mexican Peso`: {
		countries: Countries{`MEXICO`},
		currency:  `Mexican Peso`,
		code:      `MXN`,
		number:    `484`,
	},
	`Mexican Unidad de Inversion (UDI)`: {
		countries: Countries{`MEXICO`},
		currency:  `Mexican Unidad de Inversion (UDI)`,
		code:      `MXV`,
		number:    `979`,
	},
	`Moldovan Leu`: {
		countries: Countries{`MOLDOVA (THE REPUBLIC OF)`, `MOLDOVA`},
		currency:  `Moldovan Leu`,
		code:      `MDL`,
		number:    `498`,
	},
	`Tugrik`: {
		countries: Countries{`MONGOLIA`},
		currency:  `Tugrik`,
		code:      `MNT`,
		number:    `496`,
	},
	`Moroccan Dirham`: {
		countries: Countries{`MOROCCO`, `WESTERN SAHARA`},
		currency:  `Moroccan Dirham`,
		code:      `MAD`,
		number:    `504`,
	},
	`Mozambique Metical`: {
		countries: Countries{`MOZAMBIQUE`},
		currency:  `Mozambique Metical`,
		code:      `MZN`,
		number:    `943`,
	},
	`Kyat`: {
		countries: Countries{`MYANMAR`},
		currency:  `Kyat`,
		code:      `MMK`,
		number:    `104`,
	},
	`Namibia Dollar`: {
		countries: Countries{`NAMIBIA`},
		currency:  `Namibia Dollar`,
		code:      `NAD`,
		number:    `516`,
	},
	`Nepalese Rupee`: {
		countries: Countries{`NEPAL`},
		currency:  `Nepalese Rupee`,
		code:      `NPR`,
		number:    `524`,
	},
	`Cordoba Oro`: {
		countries: Countries{`NICARAGUA`},
		currency:  `Cordoba Oro`,
		code:      `NIO`,
		number:    `558`,
	},
	`Naira`: {
		countries: Countries{`NIGERIA`},
		currency:  `Naira`,
		code:      `NGN`,
		number:    `566`,
	},
	`Rial Omani`: {
		countries: Countries{`OMAN`},
		currency:  `Rial Omani`,
		code:      `OMR`,
		number:    `512`,
	},
	`Pakistan Rupee`: {
		countries: Countries{`PAKISTAN`},
		currency:  `Pakistan Rupee`,
		code:      `PKR`,
		number:    `586`,
	},
	`Balboa`: {
		countries: Countries{`PANAMA`},
		currency:  `Balboa`,
		code:      `PAB`,
		number:    `590`,
	},
	`Kina`: {
		countries: Countries{`PAPUA NEW GUINEA`},
		currency:  `Kina`,
		code:      `PGK`,
		number:    `598`,
	},
	`Guarani`: {
		countries: Countries{`PARAGUAY`},
		currency:  `Guarani`,
		code:      `PYG`,
		number:    `600`,
	},
	`Sol`: {
		countries: Countries{`PERU`},
		currency:  `Sol`,
		code:      `PEN`,
		number:    `604`,
	},
	`Philippine Peso`: {
		countries: Countries{`PHILIPPINES (THE)`, `PHILIPPINES`},
		currency:  `Philippine Peso`,
		code:      `PHP`,
		number:    `608`,
	},
	`Zloty`: {
		countries: Countries{`POLAND`},
		currency:  `Zloty`,
		code:      `PLN`,
		number:    `985`,
	},
	`Qatari Rial`: {
		countries: Countries{`QATAR`},
		currency:  `Qatari Rial`,
		code:      `QAR`,
		number:    `634`,
	},
	`Romanian Leu`: {
		countries: Countries{`ROMANIA`},
		currency:  `Romanian Leu`,
		code:      `RON`,
		number:    `946`,
	},
	`Russian Ruble`: {
		countries: Countries{`RUSSIAN FEDERATION (THE)`, `RUSSIAN FEDERATION`},
		currency:  `Russian Ruble`,
		code:      `RUB`,
		number:    `643`,
	},
	`Rwanda Franc`: {
		countries: Countries{`RWANDA`},
		currency:  `Rwanda Franc`,
		code:      `RWF`,
		number:    `646`,
	},
	`Saint Helena Pound`: {
		countries: Countries{`SAINT HELENA, ASCENSION AND TRISTAN DA CUNHA`},
		currency:  `Saint Helena Pound`,
		code:      `SHP`,
		number:    `654`,
	},
	`Tala`: {
		countries: Countries{`SAMOA`},
		currency:  `Tala`,
		code:      `WST`,
		number:    `882`,
	},
	`Dobra`: {
		countries: Countries{`SAO TOME AND PRINCIPE`},
		currency:  `Dobra`,
		code:      `STN`,
		number:    `930`,
	},
	`Saudi Riyal`: {
		countries: Countries{`SAUDI ARABIA`},
		currency:  `Saudi Riyal`,
		code:      `SAR`,
		number:    `682`,
	},
	`Serbian Dinar`: {
		countries: Countries{`SERBIA`},
		currency:  `Serbian Dinar`,
		code:      `RSD`,
		number:    `941`,
	},
	`Seychelles Rupee`: {
		countries: Countries{`SEYCHELLES`},
		currency:  `Seychelles Rupee`,
		code:      `SCR`,
		number:    `690`,
	},
	`Leone`: {
		countries: Countries{`SIERRA LEONE`},
		currency:  `Leone`,
		code:      `SLL`,
		number:    `694`,
	},
	`Singapore Dollar`: {
		countries: Countries{`SINGAPORE`},
		currency:  `Singapore Dollar`,
		code:      `SGD`,
		number:    `702`,
	},
	`Sucre`: {
		countries: Countries{`SISTEMA UNITARIO DE COMPENSACION REGIONAL DE PAGOS "SUCRE"`},
		currency:  `Sucre`,
		code:      `XSU`,
		number:    `994`,
	},
	`Solomon Islands Dollar`: {
		countries: Countries{`SOLOMON ISLANDS`},
		currency:  `Solomon Islands Dollar`,
		code:      `SBD`,
		number:    `090`,
	},
	`Somali Shilling`: {
		countries: Countries{`SOMALIA`},
		currency:  `Somali Shilling`,
		code:      `SOS`,
		number:    `706`,
	},
	`South Sudanese Pound`: {
		countries: Countries{`SOUTH SUDAN`},
		currency:  `South Sudanese Pound`,
		code:      `SSP`,
		number:    `728`,
	},
	`Sri Lanka Rupee`: {
		countries: Countries{`SRI LANKA`},
		currency:  `Sri Lanka Rupee`,
		code:      `LKR`,
		number:    `144`,
	},
	`Sudanese Pound`: {
		countries: Countries{`SUDAN (THE)`, `SUDAN`},
		currency:  `Sudanese Pound`,
		code:      `SDG`,
		number:    `938`,
	},
	`Surinam Dollar`: {
		countries: Countries{`SURINAME`},
		currency:  `Surinam Dollar`,
		code:      `SRD`,
		number:    `968`,
	},
	`Swedish Krona`: {
		countries: Countries{`SWEDEN`},
		currency:  `Swedish Krona`,
		code:      `SEK`,
		number:    `752`,
	},
	`WIR Euro`: {
		countries: Countries{`SWITZERLAND`},
		currency:  `WIR Euro`,
		code:      `CHE`,
		number:    `947`,
	},
	`WIR Franc`: {
		countries: Countries{`SWITZERLAND`},
		currency:  `WIR Franc`,
		code:      `CHW`,
		number:    `948`,
	},
	`Syrian Pound`: {
		countries: Countries{`SYRIAN ARAB REPUBLIC`},
		currency:  `Syrian Pound`,
		code:      `SYP`,
		number:    `760`,
	},
	`New Taiwan Dollar`: {
		countries: Countries{`TAIWAN (PROVINCE OF CHINA)`, `TAIWAN`},
		currency:  `New Taiwan Dollar`,
		code:      `TWD`,
		number:    `901`,
	},
	`Somoni`: {
		countries: Countries{`TAJIKISTAN`},
		currency:  `Somoni`,
		code:      `TJS`,
		number:    `972`,
	},
	`Tanzanian Shilling`: {
		countries: Countries{`TANZANIA, UNITED REPUBLIC OF`},
		currency:  `Tanzanian Shilling`,
		code:      `TZS`,
		number:    `834`,
	},
	`Baht`: {
		countries: Countries{`THAILAND`},
		currency:  `Baht`,
		code:      `THB`,
		number:    `764`,
	},
	`Pa’anga`: {
		countries: Countries{`TONGA`},
		currency:  `Pa’anga`,
		code:      `TOP`,
		number:    `776`,
	},
	`Trinidad and Tobago Dollar`: {
		countries: Countries{`TRINIDAD AND TOBAGO`},
		currency:  `Trinidad and Tobago Dollar`,
		code:      `TTD`,
		number:    `780`,
	},
	`Tunisian Dinar`: {
		countries: Countries{`TUNISIA`},
		currency:  `Tunisian Dinar`,
		code:      `TND`,
		number:    `788`,
	},
	`Turkish Lira`: {
		countries: Countries{`TURKEY`},
		currency:  `Turkish Lira`,
		code:      `TRY`,
		number:    `949`,
	},
	`Turkmenistan New Manat`: {
		countries: Countries{`TURKMENISTAN`},
		currency:  `Turkmenistan New Manat`,
		code:      `TMT`,
		number:    `934`,
	},
	`Uganda Shilling`: {
		countries: Countries{`UGANDA`},
		currency:  `Uganda Shilling`,
		code:      `UGX`,
		number:    `800`,
	},
	`Hryvnia`: {
		countries: Countries{`UKRAINE`},
		currency:  `Hryvnia`,
		code:      `UAH`,
		number:    `980`,
	},
	`UAE Dirham`: {
		countries: Countries{`UNITED ARAB EMIRATES (THE)`, `UNITED ARAB EMIRATES`},
		currency:  `UAE Dirham`,
		code:      `AED`,
		number:    `784`,
	},
	`US Dollar (Next day)`: {
		countries: Countries{`UNITED STATES OF AMERICA (THE)`, `UNITED STATES OF AMERICA`},
		currency:  `US Dollar (Next day)`,
		code:      `USN`,
		number:    `997`,
	},
	`Peso Uruguayo`: {
		countries: Countries{`URUGUAY`},
		currency:  `Peso Uruguayo`,
		code:      `UYU`,
		number:    `858`,
	},
	`Uruguay Peso en Unidades Indexadas (UI)`: {
		countries: Countries{`URUGUAY`},
		currency:  `Uruguay Peso en Unidades Indexadas (UI)`,
		code:      `UYI`,
		number:    `940`,
	},
	`Unidad Previsional`: {
		countries: Countries{`URUGUAY`},
		currency:  `Unidad Previsional`,
		code:      `UYW`,
		number:    `927`,
	},
	`Uzbekistan Sum`: {
		countries: Countries{`UZBEKISTAN`},
		currency:  `Uzbekistan Sum`,
		code:      `UZS`,
		number:    `860`,
	},
	`Vatu`: {
		countries: Countries{`VANUATU`},
		currency:  `Vatu`,
		code:      `VUV`,
		number:    `548`,
	},
	`Bolívar Soberano`: {
		countries: Countries{`VENEZUELA (BOLIVARIAN REPUBLIC OF)`, `VENEZUELA`},
		currency:  `Bolívar Soberano`,
		code:      `VES`,
		number:    `928`,
	},
	`Dong`: {
		countries: Countries{`VIET NAM`},
		currency:  `Dong`,
		code:      `VND`,
		number:    `704`,
	},
	`Yemeni Rial`: {
		countries: Countries{`YEMEN`},
		currency:  `Yemeni Rial`,
		code:      `YER`,
		number:    `886`,
	},
	`Zambian Kwacha`: {
		countries: Countries{`ZAMBIA`},
		currency:  `Zambian Kwacha`,
		code:      `ZMW`,
		number:    `967`,
	},
	`Zimbabwe Dollar`: {
		countries: Countries{`ZIMBABWE`},
		currency:  `Zimbabwe Dollar`,
		code:      `ZWL`,
		number:    `932`,
	},
}
