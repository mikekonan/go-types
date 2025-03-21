package timezone

var timezonesByName = map[string]Timezone{
	"Africa/Abidjan":                   Timezone("Africa/Abidjan"),
	"Africa/Accra":                     Timezone("Africa/Accra"),
	"Africa/Addis_Ababa":               Timezone("Africa/Addis_Ababa"),
	"Africa/Algiers":                   Timezone("Africa/Algiers"),
	"Africa/Asmara":                    Timezone("Africa/Asmara"),
	"Africa/Asmera":                    Timezone("Africa/Asmera"),
	"Africa/Bamako":                    Timezone("Africa/Bamako"),
	"Africa/Bangui":                    Timezone("Africa/Bangui"),
	"Africa/Banjul":                    Timezone("Africa/Banjul"),
	"Africa/Bissau":                    Timezone("Africa/Bissau"),
	"Africa/Blantyre":                  Timezone("Africa/Blantyre"),
	"Africa/Brazzaville":               Timezone("Africa/Brazzaville"),
	"Africa/Bujumbura":                 Timezone("Africa/Bujumbura"),
	"Africa/Cairo":                     Timezone("Africa/Cairo"),
	"Africa/Casablanca":                Timezone("Africa/Casablanca"),
	"Africa/Ceuta":                     Timezone("Africa/Ceuta"),
	"Africa/Conakry":                   Timezone("Africa/Conakry"),
	"Africa/Dakar":                     Timezone("Africa/Dakar"),
	"Africa/Dar_es_Salaam":             Timezone("Africa/Dar_es_Salaam"),
	"Africa/Djibouti":                  Timezone("Africa/Djibouti"),
	"Africa/Douala":                    Timezone("Africa/Douala"),
	"Africa/El_Aaiun":                  Timezone("Africa/El_Aaiun"),
	"Africa/Freetown":                  Timezone("Africa/Freetown"),
	"Africa/Gaborone":                  Timezone("Africa/Gaborone"),
	"Africa/Harare":                    Timezone("Africa/Harare"),
	"Africa/Johannesburg":              Timezone("Africa/Johannesburg"),
	"Africa/Juba":                      Timezone("Africa/Juba"),
	"Africa/Kampala":                   Timezone("Africa/Kampala"),
	"Africa/Khartoum":                  Timezone("Africa/Khartoum"),
	"Africa/Kigali":                    Timezone("Africa/Kigali"),
	"Africa/Kinshasa":                  Timezone("Africa/Kinshasa"),
	"Africa/Lagos":                     Timezone("Africa/Lagos"),
	"Africa/Libreville":                Timezone("Africa/Libreville"),
	"Africa/Lome":                      Timezone("Africa/Lome"),
	"Africa/Luanda":                    Timezone("Africa/Luanda"),
	"Africa/Lubumbashi":                Timezone("Africa/Lubumbashi"),
	"Africa/Lusaka":                    Timezone("Africa/Lusaka"),
	"Africa/Malabo":                    Timezone("Africa/Malabo"),
	"Africa/Maputo":                    Timezone("Africa/Maputo"),
	"Africa/Maseru":                    Timezone("Africa/Maseru"),
	"Africa/Mbabane":                   Timezone("Africa/Mbabane"),
	"Africa/Mogadishu":                 Timezone("Africa/Mogadishu"),
	"Africa/Monrovia":                  Timezone("Africa/Monrovia"),
	"Africa/Nairobi":                   Timezone("Africa/Nairobi"),
	"Africa/Ndjamena":                  Timezone("Africa/Ndjamena"),
	"Africa/Niamey":                    Timezone("Africa/Niamey"),
	"Africa/Nouakchott":                Timezone("Africa/Nouakchott"),
	"Africa/Ouagadougou":               Timezone("Africa/Ouagadougou"),
	"Africa/Porto-Novo":                Timezone("Africa/Porto-Novo"),
	"Africa/Sao_Tome":                  Timezone("Africa/Sao_Tome"),
	"Africa/Timbuktu":                  Timezone("Africa/Timbuktu"),
	"Africa/Tripoli":                   Timezone("Africa/Tripoli"),
	"Africa/Tunis":                     Timezone("Africa/Tunis"),
	"Africa/Windhoek":                  Timezone("Africa/Windhoek"),
	"America/Adak":                     Timezone("America/Adak"),
	"America/Anchorage":                Timezone("America/Anchorage"),
	"America/Anguilla":                 Timezone("America/Anguilla"),
	"America/Antigua":                  Timezone("America/Antigua"),
	"America/Araguaina":                Timezone("America/Araguaina"),
	"America/Argentina/Buenos_Aires":   Timezone("America/Argentina/Buenos_Aires"),
	"America/Argentina/Catamarca":      Timezone("America/Argentina/Catamarca"),
	"America/Argentina/ComodRivadavia": Timezone("America/Argentina/ComodRivadavia"),
	"America/Argentina/Cordoba":        Timezone("America/Argentina/Cordoba"),
	"America/Argentina/Jujuy":          Timezone("America/Argentina/Jujuy"),
	"America/Argentina/La_Rioja":       Timezone("America/Argentina/La_Rioja"),
	"America/Argentina/Mendoza":        Timezone("America/Argentina/Mendoza"),
	"America/Argentina/Rio_Gallegos":   Timezone("America/Argentina/Rio_Gallegos"),
	"America/Argentina/Salta":          Timezone("America/Argentina/Salta"),
	"America/Argentina/San_Juan":       Timezone("America/Argentina/San_Juan"),
	"America/Argentina/San_Luis":       Timezone("America/Argentina/San_Luis"),
	"America/Argentina/Tucuman":        Timezone("America/Argentina/Tucuman"),
	"America/Argentina/Ushuaia":        Timezone("America/Argentina/Ushuaia"),
	"America/Aruba":                    Timezone("America/Aruba"),
	"America/Asuncion":                 Timezone("America/Asuncion"),
	"America/Atikokan":                 Timezone("America/Atikokan"),
	"America/Atka":                     Timezone("America/Atka"),
	"America/Bahia":                    Timezone("America/Bahia"),
	"America/Bahia_Banderas":           Timezone("America/Bahia_Banderas"),
	"America/Barbados":                 Timezone("America/Barbados"),
	"America/Belem":                    Timezone("America/Belem"),
	"America/Belize":                   Timezone("America/Belize"),
	"America/Blanc-Sablon":             Timezone("America/Blanc-Sablon"),
	"America/Boa_Vista":                Timezone("America/Boa_Vista"),
	"America/Bogota":                   Timezone("America/Bogota"),
	"America/Boise":                    Timezone("America/Boise"),
	"America/Buenos_Aires":             Timezone("America/Buenos_Aires"),
	"America/Cambridge_Bay":            Timezone("America/Cambridge_Bay"),
	"America/Campo_Grande":             Timezone("America/Campo_Grande"),
	"America/Cancun":                   Timezone("America/Cancun"),
	"America/Caracas":                  Timezone("America/Caracas"),
	"America/Catamarca":                Timezone("America/Catamarca"),
	"America/Cayenne":                  Timezone("America/Cayenne"),
	"America/Cayman":                   Timezone("America/Cayman"),
	"America/Chicago":                  Timezone("America/Chicago"),
	"America/Chihuahua":                Timezone("America/Chihuahua"),
	"America/Ciudad_Juarez":            Timezone("America/Ciudad_Juarez"),
	"America/Coral_Harbour":            Timezone("America/Coral_Harbour"),
	"America/Cordoba":                  Timezone("America/Cordoba"),
	"America/Costa_Rica":               Timezone("America/Costa_Rica"),
	"America/Creston":                  Timezone("America/Creston"),
	"America/Cuiaba":                   Timezone("America/Cuiaba"),
	"America/Curacao":                  Timezone("America/Curacao"),
	"America/Danmarkshavn":             Timezone("America/Danmarkshavn"),
	"America/Dawson":                   Timezone("America/Dawson"),
	"America/Dawson_Creek":             Timezone("America/Dawson_Creek"),
	"America/Denver":                   Timezone("America/Denver"),
	"America/Detroit":                  Timezone("America/Detroit"),
	"America/Dominica":                 Timezone("America/Dominica"),
	"America/Edmonton":                 Timezone("America/Edmonton"),
	"America/Eirunepe":                 Timezone("America/Eirunepe"),
	"America/El_Salvador":              Timezone("America/El_Salvador"),
	"America/Ensenada":                 Timezone("America/Ensenada"),
	"America/Fort_Nelson":              Timezone("America/Fort_Nelson"),
	"America/Fort_Wayne":               Timezone("America/Fort_Wayne"),
	"America/Fortaleza":                Timezone("America/Fortaleza"),
	"America/Glace_Bay":                Timezone("America/Glace_Bay"),
	"America/Godthab":                  Timezone("America/Godthab"),
	"America/Goose_Bay":                Timezone("America/Goose_Bay"),
	"America/Grand_Turk":               Timezone("America/Grand_Turk"),
	"America/Grenada":                  Timezone("America/Grenada"),
	"America/Guadeloupe":               Timezone("America/Guadeloupe"),
	"America/Guatemala":                Timezone("America/Guatemala"),
	"America/Guayaquil":                Timezone("America/Guayaquil"),
	"America/Guyana":                   Timezone("America/Guyana"),
	"America/Halifax":                  Timezone("America/Halifax"),
	"America/Havana":                   Timezone("America/Havana"),
	"America/Hermosillo":               Timezone("America/Hermosillo"),
	"America/Indiana/Indianapolis":     Timezone("America/Indiana/Indianapolis"),
	"America/Indiana/Knox":             Timezone("America/Indiana/Knox"),
	"America/Indiana/Marengo":          Timezone("America/Indiana/Marengo"),
	"America/Indiana/Petersburg":       Timezone("America/Indiana/Petersburg"),
	"America/Indiana/Tell_City":        Timezone("America/Indiana/Tell_City"),
	"America/Indiana/Vevay":            Timezone("America/Indiana/Vevay"),
	"America/Indiana/Vincennes":        Timezone("America/Indiana/Vincennes"),
	"America/Indiana/Winamac":          Timezone("America/Indiana/Winamac"),
	"America/Indianapolis":             Timezone("America/Indianapolis"),
	"America/Inuvik":                   Timezone("America/Inuvik"),
	"America/Iqaluit":                  Timezone("America/Iqaluit"),
	"America/Jamaica":                  Timezone("America/Jamaica"),
	"America/Jujuy":                    Timezone("America/Jujuy"),
	"America/Juneau":                   Timezone("America/Juneau"),
	"America/Kentucky/Louisville":      Timezone("America/Kentucky/Louisville"),
	"America/Kentucky/Monticello":      Timezone("America/Kentucky/Monticello"),
	"America/Knox_IN":                  Timezone("America/Knox_IN"),
	"America/Kralendijk":               Timezone("America/Kralendijk"),
	"America/La_Paz":                   Timezone("America/La_Paz"),
	"America/Lima":                     Timezone("America/Lima"),
	"America/Los_Angeles":              Timezone("America/Los_Angeles"),
	"America/Louisville":               Timezone("America/Louisville"),
	"America/Lower_Princes":            Timezone("America/Lower_Princes"),
	"America/Maceio":                   Timezone("America/Maceio"),
	"America/Managua":                  Timezone("America/Managua"),
	"America/Manaus":                   Timezone("America/Manaus"),
	"America/Marigot":                  Timezone("America/Marigot"),
	"America/Martinique":               Timezone("America/Martinique"),
	"America/Matamoros":                Timezone("America/Matamoros"),
	"America/Mazatlan":                 Timezone("America/Mazatlan"),
	"America/Mendoza":                  Timezone("America/Mendoza"),
	"America/Menominee":                Timezone("America/Menominee"),
	"America/Merida":                   Timezone("America/Merida"),
	"America/Metlakatla":               Timezone("America/Metlakatla"),
	"America/Mexico_City":              Timezone("America/Mexico_City"),
	"America/Miquelon":                 Timezone("America/Miquelon"),
	"America/Moncton":                  Timezone("America/Moncton"),
	"America/Monterrey":                Timezone("America/Monterrey"),
	"America/Montevideo":               Timezone("America/Montevideo"),
	"America/Montreal":                 Timezone("America/Montreal"),
	"America/Montserrat":               Timezone("America/Montserrat"),
	"America/Nassau":                   Timezone("America/Nassau"),
	"America/New_York":                 Timezone("America/New_York"),
	"America/Nipigon":                  Timezone("America/Nipigon"),
	"America/Nome":                     Timezone("America/Nome"),
	"America/Noronha":                  Timezone("America/Noronha"),
	"America/North_Dakota/Beulah":      Timezone("America/North_Dakota/Beulah"),
	"America/North_Dakota/Center":      Timezone("America/North_Dakota/Center"),
	"America/North_Dakota/New_Salem":   Timezone("America/North_Dakota/New_Salem"),
	"America/Nuuk":                     Timezone("America/Nuuk"),
	"America/Ojinaga":                  Timezone("America/Ojinaga"),
	"America/Panama":                   Timezone("America/Panama"),
	"America/Pangnirtung":              Timezone("America/Pangnirtung"),
	"America/Paramaribo":               Timezone("America/Paramaribo"),
	"America/Phoenix":                  Timezone("America/Phoenix"),
	"America/Port-au-Prince":           Timezone("America/Port-au-Prince"),
	"America/Port_of_Spain":            Timezone("America/Port_of_Spain"),
	"America/Porto_Acre":               Timezone("America/Porto_Acre"),
	"America/Porto_Velho":              Timezone("America/Porto_Velho"),
	"America/Puerto_Rico":              Timezone("America/Puerto_Rico"),
	"America/Punta_Arenas":             Timezone("America/Punta_Arenas"),
	"America/Rainy_River":              Timezone("America/Rainy_River"),
	"America/Rankin_Inlet":             Timezone("America/Rankin_Inlet"),
	"America/Recife":                   Timezone("America/Recife"),
	"America/Regina":                   Timezone("America/Regina"),
	"America/Resolute":                 Timezone("America/Resolute"),
	"America/Rio_Branco":               Timezone("America/Rio_Branco"),
	"America/Rosario":                  Timezone("America/Rosario"),
	"America/Santa_Isabel":             Timezone("America/Santa_Isabel"),
	"America/Santarem":                 Timezone("America/Santarem"),
	"America/Santiago":                 Timezone("America/Santiago"),
	"America/Santo_Domingo":            Timezone("America/Santo_Domingo"),
	"America/Sao_Paulo":                Timezone("America/Sao_Paulo"),
	"America/Scoresbysund":             Timezone("America/Scoresbysund"),
	"America/Shiprock":                 Timezone("America/Shiprock"),
	"America/Sitka":                    Timezone("America/Sitka"),
	"America/St_Barthelemy":            Timezone("America/St_Barthelemy"),
	"America/St_Johns":                 Timezone("America/St_Johns"),
	"America/St_Kitts":                 Timezone("America/St_Kitts"),
	"America/St_Lucia":                 Timezone("America/St_Lucia"),
	"America/St_Thomas":                Timezone("America/St_Thomas"),
	"America/St_Vincent":               Timezone("America/St_Vincent"),
	"America/Swift_Current":            Timezone("America/Swift_Current"),
	"America/Tegucigalpa":              Timezone("America/Tegucigalpa"),
	"America/Thule":                    Timezone("America/Thule"),
	"America/Thunder_Bay":              Timezone("America/Thunder_Bay"),
	"America/Tijuana":                  Timezone("America/Tijuana"),
	"America/Toronto":                  Timezone("America/Toronto"),
	"America/Tortola":                  Timezone("America/Tortola"),
	"America/Vancouver":                Timezone("America/Vancouver"),
	"America/Virgin":                   Timezone("America/Virgin"),
	"America/Whitehorse":               Timezone("America/Whitehorse"),
	"America/Winnipeg":                 Timezone("America/Winnipeg"),
	"America/Yakutat":                  Timezone("America/Yakutat"),
	"America/Yellowknife":              Timezone("America/Yellowknife"),
	"Antarctica/Casey":                 Timezone("Antarctica/Casey"),
	"Antarctica/Davis":                 Timezone("Antarctica/Davis"),
	"Antarctica/DumontDUrville":        Timezone("Antarctica/DumontDUrville"),
	"Antarctica/Macquarie":             Timezone("Antarctica/Macquarie"),
	"Antarctica/Mawson":                Timezone("Antarctica/Mawson"),
	"Antarctica/McMurdo":               Timezone("Antarctica/McMurdo"),
	"Antarctica/Palmer":                Timezone("Antarctica/Palmer"),
	"Antarctica/Rothera":               Timezone("Antarctica/Rothera"),
	"Antarctica/South_Pole":            Timezone("Antarctica/South_Pole"),
	"Antarctica/Syowa":                 Timezone("Antarctica/Syowa"),
	"Antarctica/Troll":                 Timezone("Antarctica/Troll"),
	"Antarctica/Vostok":                Timezone("Antarctica/Vostok"),
	"Arctic/Longyearbyen":              Timezone("Arctic/Longyearbyen"),
	"Asia/Aden":                        Timezone("Asia/Aden"),
	"Asia/Almaty":                      Timezone("Asia/Almaty"),
	"Asia/Amman":                       Timezone("Asia/Amman"),
	"Asia/Anadyr":                      Timezone("Asia/Anadyr"),
	"Asia/Aqtau":                       Timezone("Asia/Aqtau"),
	"Asia/Aqtobe":                      Timezone("Asia/Aqtobe"),
	"Asia/Ashgabat":                    Timezone("Asia/Ashgabat"),
	"Asia/Ashkhabad":                   Timezone("Asia/Ashkhabad"),
	"Asia/Atyrau":                      Timezone("Asia/Atyrau"),
	"Asia/Baghdad":                     Timezone("Asia/Baghdad"),
	"Asia/Bahrain":                     Timezone("Asia/Bahrain"),
	"Asia/Baku":                        Timezone("Asia/Baku"),
	"Asia/Bangkok":                     Timezone("Asia/Bangkok"),
	"Asia/Barnaul":                     Timezone("Asia/Barnaul"),
	"Asia/Beirut":                      Timezone("Asia/Beirut"),
	"Asia/Bishkek":                     Timezone("Asia/Bishkek"),
	"Asia/Brunei":                      Timezone("Asia/Brunei"),
	"Asia/Calcutta":                    Timezone("Asia/Calcutta"),
	"Asia/Chita":                       Timezone("Asia/Chita"),
	"Asia/Choibalsan":                  Timezone("Asia/Choibalsan"),
	"Asia/Chongqing":                   Timezone("Asia/Chongqing"),
	"Asia/Chungking":                   Timezone("Asia/Chungking"),
	"Asia/Colombo":                     Timezone("Asia/Colombo"),
	"Asia/Dacca":                       Timezone("Asia/Dacca"),
	"Asia/Damascus":                    Timezone("Asia/Damascus"),
	"Asia/Dhaka":                       Timezone("Asia/Dhaka"),
	"Asia/Dili":                        Timezone("Asia/Dili"),
	"Asia/Dubai":                       Timezone("Asia/Dubai"),
	"Asia/Dushanbe":                    Timezone("Asia/Dushanbe"),
	"Asia/Famagusta":                   Timezone("Asia/Famagusta"),
	"Asia/Gaza":                        Timezone("Asia/Gaza"),
	"Asia/Harbin":                      Timezone("Asia/Harbin"),
	"Asia/Hebron":                      Timezone("Asia/Hebron"),
	"Asia/Ho_Chi_Minh":                 Timezone("Asia/Ho_Chi_Minh"),
	"Asia/Hong_Kong":                   Timezone("Asia/Hong_Kong"),
	"Asia/Hovd":                        Timezone("Asia/Hovd"),
	"Asia/Irkutsk":                     Timezone("Asia/Irkutsk"),
	"Asia/Istanbul":                    Timezone("Asia/Istanbul"),
	"Asia/Jakarta":                     Timezone("Asia/Jakarta"),
	"Asia/Jayapura":                    Timezone("Asia/Jayapura"),
	"Asia/Jerusalem":                   Timezone("Asia/Jerusalem"),
	"Asia/Kabul":                       Timezone("Asia/Kabul"),
	"Asia/Kamchatka":                   Timezone("Asia/Kamchatka"),
	"Asia/Karachi":                     Timezone("Asia/Karachi"),
	"Asia/Kashgar":                     Timezone("Asia/Kashgar"),
	"Asia/Kathmandu":                   Timezone("Asia/Kathmandu"),
	"Asia/Katmandu":                    Timezone("Asia/Katmandu"),
	"Asia/Khandyga":                    Timezone("Asia/Khandyga"),
	"Asia/Kolkata":                     Timezone("Asia/Kolkata"),
	"Asia/Krasnoyarsk":                 Timezone("Asia/Krasnoyarsk"),
	"Asia/Kuala_Lumpur":                Timezone("Asia/Kuala_Lumpur"),
	"Asia/Kuching":                     Timezone("Asia/Kuching"),
	"Asia/Kuwait":                      Timezone("Asia/Kuwait"),
	"Asia/Macao":                       Timezone("Asia/Macao"),
	"Asia/Macau":                       Timezone("Asia/Macau"),
	"Asia/Magadan":                     Timezone("Asia/Magadan"),
	"Asia/Makassar":                    Timezone("Asia/Makassar"),
	"Asia/Manila":                      Timezone("Asia/Manila"),
	"Asia/Muscat":                      Timezone("Asia/Muscat"),
	"Asia/Nicosia":                     Timezone("Asia/Nicosia"),
	"Asia/Novokuznetsk":                Timezone("Asia/Novokuznetsk"),
	"Asia/Novosibirsk":                 Timezone("Asia/Novosibirsk"),
	"Asia/Omsk":                        Timezone("Asia/Omsk"),
	"Asia/Oral":                        Timezone("Asia/Oral"),
	"Asia/Phnom_Penh":                  Timezone("Asia/Phnom_Penh"),
	"Asia/Pontianak":                   Timezone("Asia/Pontianak"),
	"Asia/Pyongyang":                   Timezone("Asia/Pyongyang"),
	"Asia/Qatar":                       Timezone("Asia/Qatar"),
	"Asia/Qostanay":                    Timezone("Asia/Qostanay"),
	"Asia/Qyzylorda":                   Timezone("Asia/Qyzylorda"),
	"Asia/Rangoon":                     Timezone("Asia/Rangoon"),
	"Asia/Riyadh":                      Timezone("Asia/Riyadh"),
	"Asia/Saigon":                      Timezone("Asia/Saigon"),
	"Asia/Sakhalin":                    Timezone("Asia/Sakhalin"),
	"Asia/Samarkand":                   Timezone("Asia/Samarkand"),
	"Asia/Seoul":                       Timezone("Asia/Seoul"),
	"Asia/Shanghai":                    Timezone("Asia/Shanghai"),
	"Asia/Singapore":                   Timezone("Asia/Singapore"),
	"Asia/Srednekolymsk":               Timezone("Asia/Srednekolymsk"),
	"Asia/Taipei":                      Timezone("Asia/Taipei"),
	"Asia/Tashkent":                    Timezone("Asia/Tashkent"),
	"Asia/Tbilisi":                     Timezone("Asia/Tbilisi"),
	"Asia/Tehran":                      Timezone("Asia/Tehran"),
	"Asia/Tel_Aviv":                    Timezone("Asia/Tel_Aviv"),
	"Asia/Thimbu":                      Timezone("Asia/Thimbu"),
	"Asia/Thimphu":                     Timezone("Asia/Thimphu"),
	"Asia/Tokyo":                       Timezone("Asia/Tokyo"),
	"Asia/Tomsk":                       Timezone("Asia/Tomsk"),
	"Asia/Ujung_Pandang":               Timezone("Asia/Ujung_Pandang"),
	"Asia/Ulaanbaatar":                 Timezone("Asia/Ulaanbaatar"),
	"Asia/Ulan_Bator":                  Timezone("Asia/Ulan_Bator"),
	"Asia/Urumqi":                      Timezone("Asia/Urumqi"),
	"Asia/Ust-Nera":                    Timezone("Asia/Ust-Nera"),
	"Asia/Vientiane":                   Timezone("Asia/Vientiane"),
	"Asia/Vladivostok":                 Timezone("Asia/Vladivostok"),
	"Asia/Yakutsk":                     Timezone("Asia/Yakutsk"),
	"Asia/Yangon":                      Timezone("Asia/Yangon"),
	"Asia/Yekaterinburg":               Timezone("Asia/Yekaterinburg"),
	"Asia/Yerevan":                     Timezone("Asia/Yerevan"),
	"Atlantic/Azores":                  Timezone("Atlantic/Azores"),
	"Atlantic/Bermuda":                 Timezone("Atlantic/Bermuda"),
	"Atlantic/Canary":                  Timezone("Atlantic/Canary"),
	"Atlantic/Cape_Verde":              Timezone("Atlantic/Cape_Verde"),
	"Atlantic/Faeroe":                  Timezone("Atlantic/Faeroe"),
	"Atlantic/Faroe":                   Timezone("Atlantic/Faroe"),
	"Atlantic/Jan_Mayen":               Timezone("Atlantic/Jan_Mayen"),
	"Atlantic/Madeira":                 Timezone("Atlantic/Madeira"),
	"Atlantic/Reykjavik":               Timezone("Atlantic/Reykjavik"),
	"Atlantic/South_Georgia":           Timezone("Atlantic/South_Georgia"),
	"Atlantic/St_Helena":               Timezone("Atlantic/St_Helena"),
	"Atlantic/Stanley":                 Timezone("Atlantic/Stanley"),
	"Australia/ACT":                    Timezone("Australia/ACT"),
	"Australia/Adelaide":               Timezone("Australia/Adelaide"),
	"Australia/Brisbane":               Timezone("Australia/Brisbane"),
	"Australia/Broken_Hill":            Timezone("Australia/Broken_Hill"),
	"Australia/Canberra":               Timezone("Australia/Canberra"),
	"Australia/Currie":                 Timezone("Australia/Currie"),
	"Australia/Darwin":                 Timezone("Australia/Darwin"),
	"Australia/Eucla":                  Timezone("Australia/Eucla"),
	"Australia/Hobart":                 Timezone("Australia/Hobart"),
	"Australia/LHI":                    Timezone("Australia/LHI"),
	"Australia/Lindeman":               Timezone("Australia/Lindeman"),
	"Australia/Lord_Howe":              Timezone("Australia/Lord_Howe"),
	"Australia/Melbourne":              Timezone("Australia/Melbourne"),
	"Australia/NSW":                    Timezone("Australia/NSW"),
	"Australia/North":                  Timezone("Australia/North"),
	"Australia/Perth":                  Timezone("Australia/Perth"),
	"Australia/Queensland":             Timezone("Australia/Queensland"),
	"Australia/South":                  Timezone("Australia/South"),
	"Australia/Sydney":                 Timezone("Australia/Sydney"),
	"Australia/Tasmania":               Timezone("Australia/Tasmania"),
	"Australia/Victoria":               Timezone("Australia/Victoria"),
	"Australia/West":                   Timezone("Australia/West"),
	"Australia/Yancowinna":             Timezone("Australia/Yancowinna"),
	"Brazil/Acre":                      Timezone("Brazil/Acre"),
	"Brazil/DeNoronha":                 Timezone("Brazil/DeNoronha"),
	"Brazil/East":                      Timezone("Brazil/East"),
	"Brazil/West":                      Timezone("Brazil/West"),
	"CET":                              Timezone("CET"),
	"CST6CDT":                          Timezone("CST6CDT"),
	"Canada/Atlantic":                  Timezone("Canada/Atlantic"),
	"Canada/Central":                   Timezone("Canada/Central"),
	"Canada/Eastern":                   Timezone("Canada/Eastern"),
	"Canada/Mountain":                  Timezone("Canada/Mountain"),
	"Canada/Newfoundland":              Timezone("Canada/Newfoundland"),
	"Canada/Pacific":                   Timezone("Canada/Pacific"),
	"Canada/Saskatchewan":              Timezone("Canada/Saskatchewan"),
	"Canada/Yukon":                     Timezone("Canada/Yukon"),
	"Chile/Continental":                Timezone("Chile/Continental"),
	"Chile/EasterIsland":               Timezone("Chile/EasterIsland"),
	"Cuba":                             Timezone("Cuba"),
	"EET":                              Timezone("EET"),
	"EST":                              Timezone("EST"),
	"EST5EDT":                          Timezone("EST5EDT"),
	"Egypt":                            Timezone("Egypt"),
	"Eire":                             Timezone("Eire"),
	"Etc/GMT":                          Timezone("Etc/GMT"),
	"Etc/GMT+0":                        Timezone("Etc/GMT+0"),
	"Etc/GMT+1":                        Timezone("Etc/GMT+1"),
	"Etc/GMT+10":                       Timezone("Etc/GMT+10"),
	"Etc/GMT+11":                       Timezone("Etc/GMT+11"),
	"Etc/GMT+12":                       Timezone("Etc/GMT+12"),
	"Etc/GMT+2":                        Timezone("Etc/GMT+2"),
	"Etc/GMT+3":                        Timezone("Etc/GMT+3"),
	"Etc/GMT+4":                        Timezone("Etc/GMT+4"),
	"Etc/GMT+5":                        Timezone("Etc/GMT+5"),
	"Etc/GMT+6":                        Timezone("Etc/GMT+6"),
	"Etc/GMT+7":                        Timezone("Etc/GMT+7"),
	"Etc/GMT+8":                        Timezone("Etc/GMT+8"),
	"Etc/GMT+9":                        Timezone("Etc/GMT+9"),
	"Etc/GMT-0":                        Timezone("Etc/GMT-0"),
	"Etc/GMT-1":                        Timezone("Etc/GMT-1"),
	"Etc/GMT-10":                       Timezone("Etc/GMT-10"),
	"Etc/GMT-11":                       Timezone("Etc/GMT-11"),
	"Etc/GMT-12":                       Timezone("Etc/GMT-12"),
	"Etc/GMT-13":                       Timezone("Etc/GMT-13"),
	"Etc/GMT-14":                       Timezone("Etc/GMT-14"),
	"Etc/GMT-2":                        Timezone("Etc/GMT-2"),
	"Etc/GMT-3":                        Timezone("Etc/GMT-3"),
	"Etc/GMT-4":                        Timezone("Etc/GMT-4"),
	"Etc/GMT-5":                        Timezone("Etc/GMT-5"),
	"Etc/GMT-6":                        Timezone("Etc/GMT-6"),
	"Etc/GMT-7":                        Timezone("Etc/GMT-7"),
	"Etc/GMT-8":                        Timezone("Etc/GMT-8"),
	"Etc/GMT-9":                        Timezone("Etc/GMT-9"),
	"Etc/GMT0":                         Timezone("Etc/GMT0"),
	"Etc/Greenwich":                    Timezone("Etc/Greenwich"),
	"Etc/UCT":                          Timezone("Etc/UCT"),
	"Etc/UTC":                          Timezone("Etc/UTC"),
	"Etc/Universal":                    Timezone("Etc/Universal"),
	"Etc/Zulu":                         Timezone("Etc/Zulu"),
	"Europe/Amsterdam":                 Timezone("Europe/Amsterdam"),
	"Europe/Andorra":                   Timezone("Europe/Andorra"),
	"Europe/Astrakhan":                 Timezone("Europe/Astrakhan"),
	"Europe/Athens":                    Timezone("Europe/Athens"),
	"Europe/Belfast":                   Timezone("Europe/Belfast"),
	"Europe/Belgrade":                  Timezone("Europe/Belgrade"),
	"Europe/Berlin":                    Timezone("Europe/Berlin"),
	"Europe/Bratislava":                Timezone("Europe/Bratislava"),
	"Europe/Brussels":                  Timezone("Europe/Brussels"),
	"Europe/Bucharest":                 Timezone("Europe/Bucharest"),
	"Europe/Budapest":                  Timezone("Europe/Budapest"),
	"Europe/Busingen":                  Timezone("Europe/Busingen"),
	"Europe/Chisinau":                  Timezone("Europe/Chisinau"),
	"Europe/Copenhagen":                Timezone("Europe/Copenhagen"),
	"Europe/Dublin":                    Timezone("Europe/Dublin"),
	"Europe/Gibraltar":                 Timezone("Europe/Gibraltar"),
	"Europe/Guernsey":                  Timezone("Europe/Guernsey"),
	"Europe/Helsinki":                  Timezone("Europe/Helsinki"),
	"Europe/Isle_of_Man":               Timezone("Europe/Isle_of_Man"),
	"Europe/Istanbul":                  Timezone("Europe/Istanbul"),
	"Europe/Jersey":                    Timezone("Europe/Jersey"),
	"Europe/Kaliningrad":               Timezone("Europe/Kaliningrad"),
	"Europe/Kiev":                      Timezone("Europe/Kiev"),
	"Europe/Kirov":                     Timezone("Europe/Kirov"),
	"Europe/Kyiv":                      Timezone("Europe/Kyiv"),
	"Europe/Lisbon":                    Timezone("Europe/Lisbon"),
	"Europe/Ljubljana":                 Timezone("Europe/Ljubljana"),
	"Europe/London":                    Timezone("Europe/London"),
	"Europe/Luxembourg":                Timezone("Europe/Luxembourg"),
	"Europe/Madrid":                    Timezone("Europe/Madrid"),
	"Europe/Malta":                     Timezone("Europe/Malta"),
	"Europe/Mariehamn":                 Timezone("Europe/Mariehamn"),
	"Europe/Minsk":                     Timezone("Europe/Minsk"),
	"Europe/Monaco":                    Timezone("Europe/Monaco"),
	"Europe/Moscow":                    Timezone("Europe/Moscow"),
	"Europe/Nicosia":                   Timezone("Europe/Nicosia"),
	"Europe/Oslo":                      Timezone("Europe/Oslo"),
	"Europe/Paris":                     Timezone("Europe/Paris"),
	"Europe/Podgorica":                 Timezone("Europe/Podgorica"),
	"Europe/Prague":                    Timezone("Europe/Prague"),
	"Europe/Riga":                      Timezone("Europe/Riga"),
	"Europe/Rome":                      Timezone("Europe/Rome"),
	"Europe/Samara":                    Timezone("Europe/Samara"),
	"Europe/San_Marino":                Timezone("Europe/San_Marino"),
	"Europe/Sarajevo":                  Timezone("Europe/Sarajevo"),
	"Europe/Saratov":                   Timezone("Europe/Saratov"),
	"Europe/Simferopol":                Timezone("Europe/Simferopol"),
	"Europe/Skopje":                    Timezone("Europe/Skopje"),
	"Europe/Sofia":                     Timezone("Europe/Sofia"),
	"Europe/Stockholm":                 Timezone("Europe/Stockholm"),
	"Europe/Tallinn":                   Timezone("Europe/Tallinn"),
	"Europe/Tirane":                    Timezone("Europe/Tirane"),
	"Europe/Tiraspol":                  Timezone("Europe/Tiraspol"),
	"Europe/Ulyanovsk":                 Timezone("Europe/Ulyanovsk"),
	"Europe/Uzhgorod":                  Timezone("Europe/Uzhgorod"),
	"Europe/Vaduz":                     Timezone("Europe/Vaduz"),
	"Europe/Vatican":                   Timezone("Europe/Vatican"),
	"Europe/Vienna":                    Timezone("Europe/Vienna"),
	"Europe/Vilnius":                   Timezone("Europe/Vilnius"),
	"Europe/Volgograd":                 Timezone("Europe/Volgograd"),
	"Europe/Warsaw":                    Timezone("Europe/Warsaw"),
	"Europe/Zagreb":                    Timezone("Europe/Zagreb"),
	"Europe/Zaporozhye":                Timezone("Europe/Zaporozhye"),
	"Europe/Zurich":                    Timezone("Europe/Zurich"),
	"GB":                               Timezone("GB"),
	"GB-Eire":                          Timezone("GB-Eire"),
	"GMT":                              Timezone("GMT"),
	"GMT+0":                            Timezone("GMT+0"),
	"GMT-0":                            Timezone("GMT-0"),
	"GMT0":                             Timezone("GMT0"),
	"Greenwich":                        Timezone("Greenwich"),
	"HST":                              Timezone("HST"),
	"Hongkong":                         Timezone("Hongkong"),
	"Iceland":                          Timezone("Iceland"),
	"Indian/Antananarivo":              Timezone("Indian/Antananarivo"),
	"Indian/Chagos":                    Timezone("Indian/Chagos"),
	"Indian/Christmas":                 Timezone("Indian/Christmas"),
	"Indian/Cocos":                     Timezone("Indian/Cocos"),
	"Indian/Comoro":                    Timezone("Indian/Comoro"),
	"Indian/Kerguelen":                 Timezone("Indian/Kerguelen"),
	"Indian/Mahe":                      Timezone("Indian/Mahe"),
	"Indian/Maldives":                  Timezone("Indian/Maldives"),
	"Indian/Mauritius":                 Timezone("Indian/Mauritius"),
	"Indian/Mayotte":                   Timezone("Indian/Mayotte"),
	"Indian/Reunion":                   Timezone("Indian/Reunion"),
	"Iran":                             Timezone("Iran"),
	"Israel":                           Timezone("Israel"),
	"Jamaica":                          Timezone("Jamaica"),
	"Japan":                            Timezone("Japan"),
	"Kwajalein":                        Timezone("Kwajalein"),
	"Libya":                            Timezone("Libya"),
	"MET":                              Timezone("MET"),
	"MST":                              Timezone("MST"),
	"MST7MDT":                          Timezone("MST7MDT"),
	"Mexico/BajaNorte":                 Timezone("Mexico/BajaNorte"),
	"Mexico/BajaSur":                   Timezone("Mexico/BajaSur"),
	"Mexico/General":                   Timezone("Mexico/General"),
	"NZ":                               Timezone("NZ"),
	"NZ-CHAT":                          Timezone("NZ-CHAT"),
	"Navajo":                           Timezone("Navajo"),
	"PRC":                              Timezone("PRC"),
	"PST8PDT":                          Timezone("PST8PDT"),
	"Pacific/Apia":                     Timezone("Pacific/Apia"),
	"Pacific/Auckland":                 Timezone("Pacific/Auckland"),
	"Pacific/Bougainville":             Timezone("Pacific/Bougainville"),
	"Pacific/Chatham":                  Timezone("Pacific/Chatham"),
	"Pacific/Chuuk":                    Timezone("Pacific/Chuuk"),
	"Pacific/Easter":                   Timezone("Pacific/Easter"),
	"Pacific/Efate":                    Timezone("Pacific/Efate"),
	"Pacific/Enderbury":                Timezone("Pacific/Enderbury"),
	"Pacific/Fakaofo":                  Timezone("Pacific/Fakaofo"),
	"Pacific/Fiji":                     Timezone("Pacific/Fiji"),
	"Pacific/Funafuti":                 Timezone("Pacific/Funafuti"),
	"Pacific/Galapagos":                Timezone("Pacific/Galapagos"),
	"Pacific/Gambier":                  Timezone("Pacific/Gambier"),
	"Pacific/Guadalcanal":              Timezone("Pacific/Guadalcanal"),
	"Pacific/Guam":                     Timezone("Pacific/Guam"),
	"Pacific/Honolulu":                 Timezone("Pacific/Honolulu"),
	"Pacific/Johnston":                 Timezone("Pacific/Johnston"),
	"Pacific/Kanton":                   Timezone("Pacific/Kanton"),
	"Pacific/Kiritimati":               Timezone("Pacific/Kiritimati"),
	"Pacific/Kosrae":                   Timezone("Pacific/Kosrae"),
	"Pacific/Kwajalein":                Timezone("Pacific/Kwajalein"),
	"Pacific/Majuro":                   Timezone("Pacific/Majuro"),
	"Pacific/Marquesas":                Timezone("Pacific/Marquesas"),
	"Pacific/Midway":                   Timezone("Pacific/Midway"),
	"Pacific/Nauru":                    Timezone("Pacific/Nauru"),
	"Pacific/Niue":                     Timezone("Pacific/Niue"),
	"Pacific/Norfolk":                  Timezone("Pacific/Norfolk"),
	"Pacific/Noumea":                   Timezone("Pacific/Noumea"),
	"Pacific/Pago_Pago":                Timezone("Pacific/Pago_Pago"),
	"Pacific/Palau":                    Timezone("Pacific/Palau"),
	"Pacific/Pitcairn":                 Timezone("Pacific/Pitcairn"),
	"Pacific/Pohnpei":                  Timezone("Pacific/Pohnpei"),
	"Pacific/Ponape":                   Timezone("Pacific/Ponape"),
	"Pacific/Port_Moresby":             Timezone("Pacific/Port_Moresby"),
	"Pacific/Rarotonga":                Timezone("Pacific/Rarotonga"),
	"Pacific/Saipan":                   Timezone("Pacific/Saipan"),
	"Pacific/Samoa":                    Timezone("Pacific/Samoa"),
	"Pacific/Tahiti":                   Timezone("Pacific/Tahiti"),
	"Pacific/Tarawa":                   Timezone("Pacific/Tarawa"),
	"Pacific/Tongatapu":                Timezone("Pacific/Tongatapu"),
	"Pacific/Truk":                     Timezone("Pacific/Truk"),
	"Pacific/Wake":                     Timezone("Pacific/Wake"),
	"Pacific/Wallis":                   Timezone("Pacific/Wallis"),
	"Pacific/Yap":                      Timezone("Pacific/Yap"),
	"Poland":                           Timezone("Poland"),
	"Portugal":                         Timezone("Portugal"),
	"ROC":                              Timezone("ROC"),
	"ROK":                              Timezone("ROK"),
	"Singapore":                        Timezone("Singapore"),
	"Turkey":                           Timezone("Turkey"),
	"UCT":                              Timezone("UCT"),
	"US/Alaska":                        Timezone("US/Alaska"),
	"US/Aleutian":                      Timezone("US/Aleutian"),
	"US/Arizona":                       Timezone("US/Arizona"),
	"US/Central":                       Timezone("US/Central"),
	"US/East-Indiana":                  Timezone("US/East-Indiana"),
	"US/Eastern":                       Timezone("US/Eastern"),
	"US/Hawaii":                        Timezone("US/Hawaii"),
	"US/Indiana-Starke":                Timezone("US/Indiana-Starke"),
	"US/Michigan":                      Timezone("US/Michigan"),
	"US/Mountain":                      Timezone("US/Mountain"),
	"US/Pacific":                       Timezone("US/Pacific"),
	"US/Samoa":                         Timezone("US/Samoa"),
	"UTC":                              Timezone("UTC"),
	"Universal":                        Timezone("Universal"),
	"W-SU":                             Timezone("W-SU"),
	"WET":                              Timezone("WET"),
	"Zulu":                             Timezone("Zulu"),
}
