package core

import (

)

/*
type LangsShare struct {
    Used
    Total
}
*/

type CountryShare struct {
    Used int
    Total int
}

type Country struct {
    Code string
    Name string
    Population int
}

type Lang struct {
    Name string
    CountryNames []string
}


type Calculations struct {
    //LangsShare LangsShare
    CountryShare CountryShare
    Country Country
}


func (app *App) GetCalculations() (*Calculations, error) {
    var calc = Calculations{Country: ALL_COUNTRIES[0]}
    return &calc, nil
}

var ALL_LANGS = []Lang{
   Lang{
        Name: "Mandarin Chinese",
        CountryNames: []string{"China", "Singapore"},
   },
   Lang{
        Name: "Hindi",
        CountryNames: []string{"India", "Fiji"},
   },
   Lang{
        Name: "Spanish",
        CountryNames: []string{"Argentina", "Bolivia", "Chile", "Colombia", "Costa Rica", "Cuba", "Dominican Republic", "Ecuador", "El Salvador", "Equatorial Guinea", "Guatemala", "Honduras", "Mexico", "Nicaragua", "Panama", "Paraguay", "Peru", "Spain", "United States", "New Mexico", "Puerto Rico", "Uruguay", "Venezuela"},
   },
   Lang{
        Name: "English",
        CountryNames: []string{"Antigua and Barbuda", "Australia", "The Bahamas", "Bangladesh", "Barbados", "Belize", "Botswana", "Brunei", "Cameroon", "Canada", "Dominica", "Ethiopia", "Eritrea", "Fiji", "The Gambia", "Ghana", "Grenada", "Guyana", "Hong Kong", "India", "Ireland", "Jamaica", "Kenya", "Kiribati", "Lesotho", "Liberia", "Malawi", "Maldives", "Malta", "Marshall Islands", "Maritius", "Micronesia", "Namibia", "Nauru", "New Zealand", "Nigeria", "Pakistan", "Palau", "Papua New Guinea", "Philippines", "Rwanda", "Saint Kitts and Nevs", "Saint Lucia", "Saint Vincent and the Grenadines", "Samoa", "Seychelles", "Sierra Leone", "Singapore", "Solomon Islands", "Somolia", "South Africa", "Sri Lanka", "Swaziland", "Tanzania", "Tonga", "Trinidad and Tobago", "Tuvalu", "Uganda", "United Kingdom", "United States", "Vanuatu", "Zambia", "Zimbabwe"},
   },
   Lang{
        Name: "Arabic",
        CountryNames: []string{"Algeria", "Bahrain", "Chad", "Comoros", "Djibouti", "Egypt", "Eritrea", "Iraq", "Israel", "Jordan", "Kuwait", "Lebanon", "Libya", "Morocco", "Niger", "Oman", "Palestinian Territories", "Quatar", "Saudi Arabia", "Somalia", "Sudan", "Syria", "Tunisia", "United Arab Emirates", "Western Sahara", "Yemen", "Mauritania", "Senegal"},
   },
   Lang{
        Name: "Portuguese",
        CountryNames: []string{"Angola", "Brazil", "Cape Verde", "East Timor", "Guinea-Bissau", "Macau", "Mozambique", "Portugal"},
   },
   Lang{
        Name: "Russian",
        CountryNames: []string{"Belarus", "Kazakhstan", "Kyyrgyzstan", "Russia"},
   },
   Lang{
        Name: "Japanese",
        CountryNames: []string{"Japan", "Palau"},
   },
   Lang{
        Name: "German",
        CountryNames: []string{"Austria", "Belgium", "Germany", "Italy", "Liechtenstein", "Luxembourg", "Poland", "Siwtzerland"},
   },
   Lang{
        Name: "Korean",
        CountryNames: []string{"North Korea", "South Korea"},
   },
   Lang{
        Name: "Vietnamese",
        CountryNames: []string{"Vietnam"},
   },
   Lang{
        Name: "French",
        CountryNames: []string{"Belgium", "Benin", "Burkina Faso", "Burundi", "Cameroon", "Canada", "Central African Republic", "Chad", "Comoros", "Congo-Brazzaville", "Congo-Kinshasa", "Côte d'Ivoire", "Djibouti", "Equatorial Guinea", "France", "French Polynesia", "Gabon", "Guernsey", "Guinea", "Haiti", "Italy", "Jersey", "Lebanon", "Luxembourg", "Madagascar", "Mali", "Martinique", "Mauritius", "Mayotte", "Monaco", "New Caledonia", "Niger", "Rwanda", "Senegal", "Seychelles", "Switzerland", "Togo", "Vanuatu"},
   },
   Lang{
        Name: "Italian",
        CountryNames: []string{"Croatia", "Italy", "San Marino", "Slovenia", "Switzerland"},
   },
   Lang{
        Name: "Turkish",
        CountryNames: []string{"Bulgaria", "Cyprus", "Turkey"},
   },
   Lang{
        Name: "Polish",
        CountryNames: []string{"Poland"},
   },
   Lang{
        Name: "Thai",
        CountryNames: []string{"Thailand"},
   },
}


var ALL_COUNTRIES = []Country{
    Country{
        Code: "AF",
        Name: "Afghanistan",
        Population: 1856781,
    },
    Country{
        Code: "AL",
        Name: "Albania",
        Population: 1798686,
    },
    Country{
        Code: "DZ",
        Name: "Algeria",
        Population: 6669927,
    },
    Country{
        Code: "AD",
        Name: "Andorra",
        Population: 71575,
    },
    Country{
        Code: "AO",
        Name: "Angola",
        Population: 4286821,
    },
    Country{
        Code: "AG",
        Name: "Antigua and Barbuda",
        Population: 81545,
    },
    Country{
        Code: "AR",
        Name: "Argentina",
        Population: 24973660,
    },
    Country{
        Code: "AM",
        Name: "Armenia",
        Population: 1300013,
    },
    Country{
        Code: "AW",
        Name: "Aruba",
        Population: 81945,
    },
    Country{
        Code: "AU",
        Name: "Australia",
        Population: 21176595,
    },
    Country{
        Code: "AT",
        Name: "Austria",
        Population: 7135168,
    },
    Country{
        Code: "AZ",
        Name: "Azerbaijan",
        Population: 5737223,
    },
    Country{
        Code: "BS",
        Name: "Bahamas",
        Population: 293875,
    },
    Country{
        Code: "BH",
        Name: "Bahrain",
        Population: 1297500,
    },
    Country{
        Code: "BD",
        Name: "Bangladesh",
        Population: 10867567,
    },
    Country{
        Code: "BB",
        Name: "Barbados",
        Population: 224588,
    },
    Country{
        Code: "BY",
        Name: "Belarus",
        Population: 4856969,
    },
    Country{
        Code: "BE",
        Name: "Belgium",
        Population: 9441116,
    },
    Country{
        Code: "BZ",
        Name: "Belize",
        Population: 90939,
    },
    Country{
        Code: "BJ",
        Name: "Benin",
        Population: 460232,
    },
    Country{
        Code: "BM",
        Name: "Bermuda",
        Population: 63987,
    },
    Country{
        Code: "BT",
        Name: "Bhutan",
        Population: 211896,
    },
    Country{
        Code: "BO",
        Name: "Bolivia",
        Population: 3970587,
    },
    Country{
        Code: "BA",
        Name: "Bosnia Herzegovina",
        Population: 2582502,
    },
    Country{
        Code: "BW",
        Name: "Botswana",
        Population: 268038,
    },
    Country{
        Code: "BR",
        Name: "Brazil",
        Population: 107822831,
    },
    Country{
        Code: "BN",
        Name: "Brunei",
        Population: 277589,
    },
    Country{
        Code: "BG",
        Name: "Bulgaria",
        Population: 4083950,
    },
    Country{
        Code: "BF",
        Name: "Burkina Faso",
        Population: 741888,
    },
    Country{
        Code: "BI",
        Name: "Burundi",
        Population: 146219,
    },
    Country{
        Code: "KH",
        Name: "Cambodia",
        Population: 828317,
    },
    Country{
        Code: "CM",
        Name: "Cameroon",
        Population: 1486815,
    },
    Country{
        Code: "CA",
        Name: "Canada",
        Population: 33000381,
    },
    Country{
        Code: "KY",
        Name: "Cayman Islands",
        Population: 47003,
    },
    Country{
        Code: "CF",
        Name: "Central African Republic",
        Population: 161524,
    },
    Country{
        Code: "TD",
        Name: "Chad",
        Population: 317197,
    },
    Country{
        Code: "CL",
        Name: "Chile",
        Population: 11686746,
    },
    Country{
        Code: "CN",
        Name: "China",
        Population: 641601070,
    },
    Country{
        Code: "CO",
        Name: "Colombia",
        Population: 25660725,
    },
    Country{
        Code: "KM",
        Name: "Comoros",
        Population: 49320,
    },
    Country{
        Code: "CG",
        Name: "Congo",
        Population: 87559,
    },
    Country{
        Code: "CK",
        Name: "Cook Islands",
        Population: 1378,
    },
    Country{
        Code: "CR",
        Name: "Costa Rica",
        Population: 2511139,
    },
    Country{
        Code: "HR",
        Name: "Croatia",
        Population: 2780534,
    },
    Country{
        Code: "CU",
        Name: "Cuba",
        Population: 3090796,
    },
    Country{
        Code: "CY",
        Name: "Cyprus",
        Population: 726663,
    },
    Country{
        Code: "CZ",
        Name: "Czech Republic",
        Population: 8322168,
    },
    Country{
        Code: "CI",
        Name: "Côte d'Ivoire",
        Population: 565874,
    },
    Country{
        Code: "DK",
        Name: "Denmark",
        Population: 5419113,
    },
    Country{
        Code: "DJ",
        Name: "Djibouti",
        Population: 80378,
    },
    Country{
        Code: "DM",
        Name: "Dominica",
        Population: 42735,
    },
    Country{
        Code: "DO",
        Name: "Dominican Republic",
        Population: 5072674,
    },
    Country{
        Code: "EC",
        Name: "Ecuador",
        Population: 6012003,
    },
    Country{
        Code: "EG",
        Name: "Egypt",
        Population: 40311562,
    },
    Country{
        Code: "SV",
        Name: "El Salvador",
        Population: 1742832,
    },
    Country{
        Code: "GQ",
        Name: "Equatorial Guinea",
        Population: 124035,
    },
    Country{
        Code: "ER",
        Name: "Eritrea",
        Population: 59784,
    },
    Country{
        Code: "EE",
        Name: "Estonia",
        Population: 1047772,
    },
    Country{
        Code: "ET",
        Name: "Ethiopia",
        Population: 1636099,
    },
    Country{
        Code: "FO",
        Name: "Faeroe Islands",
        Population: 43605,
    },
    Country{
        Code: "FJ",
        Name: "Fiji",
        Population: 325717,
    },
    Country{
        Code: "FI",
        Name: "Finland",
        Population: 5117660,
    },
    Country{
        Code: "FR",
        Name: "France",
        Population: 55429382,
    },
    Country{
        Code: "PF",
        Name: "French Polynesia",
        Population: 161025,
    },
    Country{
        Code: "GA",
        Name: "Gabon",
        Population: 168592,
    },
    Country{
        Code: "GM",
        Name: "Gambia",
        Population: 271711,
    },
    Country{
        Code: "GE",
        Name: "Georgia",
        Population: 2188311,
    },
    Country{
        Code: "DE",
        Name: "Germany",
        Population: 71727551,
    },
    Country{
        Code: "GH",
        Name: "Ghana",
        Population: 5171993,
    },
    Country{
        Code: "GI",
        Name: "Greece",
        Population: 6438325,
    },
    Country{
        Code: "GR",
        Name: "Greenland",
        Population: 39717,
    },
    Country{
        Code: "GD",
        Name: "Grenada",
        Population: 47903,
    },
    Country{
        Code: "GU",
        Name: "Guam",
        Population: 112196,
    },
    Country{
        Code: "GT",
        Name: "Guatemala",
        Population: 2716781,
    },
    Country{
        Code: "GN",
        Name: "Guinea",
        Population: 205194,
    },
    Country{
        Code: "GW",
        Name: "Guinea-Bissau",
        Population: 57764,
    },
    Country{
        Code: "GY",
        Name: "Guyana",
        Population: 295200,
    },
    Country{
        Code: "HT",
        Name: "Haiti",
        Population: 1217505,
    },
    Country{
        Code: "HN",
        Name: "Honduras",
        Population: 1602558,
    },
    Country{
        Code: "HK",
        Name: "Hong Kong SAR",
        Population: 5751357,
    },
    Country{
        Code: "HU",
        Name: "Hungary",
        Population: 7388776,
    },
    Country{
        Code: "IS",
        Name: "Iceland",
        Population: 321475,
    },
    Country{
        Code: "IN",
        Name: "India",
        Population: 243198922,
    },
    Country{
        Code: "ID",
        Name: "Indonesia",
        Population: 42258824,
    },
    Country{
        Code: "IR",
        Name: "Iran",
        Population: 22200708,
    },
    Country{
        Code: "IQ",
        Name: "Iraq",
        Population: 2707928,
    },
    Country{
        Code: "IE",
        Name: "Ireland",
        Population: 3817491,
    },
    Country{
        Code: "IL",
        Name: "Israel",
        Population: 5928772,
    },
    Country{
        Code: "IT",
        Name: "Italy",
        Population: 36593969,
    },
    Country{
        Code: "JM",
        Name: "Jamaica",
        Population: 1393381,
    },
    Country{
        Code: "JP",
        Name: "Japan",
        Population: 109252912,
    },
    Country{
        Code: "JO",
        Name: "Jordan",
        Population: 3375307,
    },
    Country{
        Code: "KZ",
        Name: "Kazakhstan",
        Population: 9850123,
    },
    Country{
        Code: "KE",
        Name: "Kenya",
        Population: 16713319,
    },
    Country{
        Code: "KI",
        Name: "Kiribati",
        Population: 12156,
    },
    Country{
        Code: "KR",
        Name: "South Korea",
        Population: 45314248,
    },
    Country{
        Code: "KW",
        Name: "Kuwait",
        Population: 3022010,
    },
    Country{
        Code: "KG",
        Name: "Kyrgyzstan",
        Population: 1359416,
    },
    Country{
        Code: "LV",
        Name: "Latvia",
        Population: 1560452,
    },
    Country{
        Code: "LB",
        Name: "Lebanon",
        Population: 3336517,
    },
    Country{
        Code: "LS",
        Name: "Lesotho",
        Population: 110065,
    },
    Country{
        Code: "LR",
        Name: "Liberia",
        Population: 190731,
    },
    Country{
        Code: "LY",
        Name: "Libya",
        Population: 1362604,
    },
    Country{
        Code: "LI",
        Name: "Liechtenstein",
        Population: 34356,
    },
    Country{
        Code: "LT",
        Name: "Lithuania",
        Population: 2113393,
    },
    Country{
        Code: "LU",
        Name: "Luxembourg",
        Population: 510177,
    },
    Country{
        Code: "MG",
        Name: "Madagascar",
        Population: 17321756,
    },
    Country{
        Code: "MW",
        Name: "Malawi",
        Population: 12150362,
    },
    Country{
        Code: "MY",
        Name: "Malaysia",
        Population: 675074,
    },
    Country{
        Code: "MV",
        Name: "Maldives",
        Population: 16645,
    },
    Country{
        Code: "ML",
        Name: "Mali",
        Population: 11862559,
    },
    Country{
        Code: "MT",
        Name: "Malta",
        Population: 173003,
    },
    Country{
        Code: "MH",
        Name: "Marshall Islands",
        Population: 1246,
    },
    Country{
        Code: "MQ",
        Name: "Martinique",
        Population: 303302,
    },
    Country{
        Code: "MR",
        Name: "Mauritania",
        Population: 455553,
    },
    Country{
        Code: "MU",
        Name: "Mauritius",
        Population: 76681,
    },
    Country{
        Code: "YT",
        Name: "Mayotte",
        Population: 107940,
    },
    Country{
        Code: "MX",
        Name: "Mexico",
        Population: 50923060,
    },
    Country{
        Code: "FM",
        Name: "Micronesia",
        Population: 29370,
    },
    Country{
        Code: "MD",
        Name: "Moldova",
        Population: 1550925,
    },
    Country{
        Code: "MC",
        Name: "Monaco",
        Population: 34214,
    },
    Country{
        Code: "MN",
        Name: "Mongolia",
        Population: 514254,
    },
    Country{
        Code: "ME",
        Name: "Montenegro",
        Population: 364978,
    },
    Country{
        Code: "MA",
        Name: "Morocco",
        Population: 20207154,
    },
    Country{
        Code: "MZ",
        Name: "Mozambique",
        Population: 1467687,
    },
    Country{
        Code: "MM",
        Name: "Myanmar",
        Population: 624991,
    },
    Country{
        Code: "NA",
        Name: "Namibia",
        Population: 347414,
    },
    Country{
        Code: "NP",
        Name: "Nepal",
        Population: 3411948,
    },
    Country{
        Code: "NL",
        Name: "Netherlands",
        Population: 16143879,
    },
    Country{
        Code: "NC",
        Name: "New Caledonia",
        Population: 163997,
    },
    Country{
        Code: "NZ",
        Name: "New Zealand",
        Population: 4162209,
    },
    Country{
        Code: "NI",
        Name: "Nicaragua",
        Population: 891675,
    },
    Country{
        Code: "NE",
        Name: "Niger",
        Population: 298310,
    },
    Country{
        Code: "NG",
        Name: "Nigeria",
        Population: 67101452,
    },
    Country{
        Code: "NU",
        Name: "Niue",
        Population: 617,
    },
    Country{
        Code: "NO",
        Name: "Norway",
        Population: 4895885,
    },
    Country{
        Code: "OM",
        Name: "Oman",
        Population: 2584316,
    },
    Country{
        Code: "PK",
        Name: "Pakistan",
        Population: 20073929,
    },
    Country{
        Code: "PA",
        Name: "Panama",
        Population: 1899892,
    },
    Country{
        Code: "PG",
        Name: "Papua New Guinea",
        Population: 187284,
    },
    Country{
        Code: "PY",
        Name: "Paraguay",
        Population: 2005278,
    },
    Country{
        Code: "PE",
        Name: "Peru",
        Population: 12583953,
    },
    Country{
        Code: "PH",
        Name: "Philippines",
        Population: 39470845,
    },
    Country{
        Code: "PL",
        Name: "Poland",
        Population: 25666238,
    },
    Country{
        Code: "PT",
        Name: "Portugal",
        Population: 7015519,
    },
    Country{
        Code: "PR",
        Name: "Puerto Rico",
        Population: 2027549,
    },
    Country{
        Code: "QA",
        Name: "Qatar",
        Population: 2191866,
    },
    Country{
        Code: "RO",
        Name: "Romania",
        Population: 11178477,
    },
    Country{
        Code: "RU",
        Name: "Russia",
        Population: 84437793,
    },
    Country{
        Code: "RW",
        Name: "Rwanda",
        Population: 1110043,
    },
    Country{
        Code: "WS",
        Name: "Samoa",
        Population: 26977,
    },
    Country{
        Code: "SM",
        Name: "San Marino",
        Population: 16631,
    },
    Country{
        Code: "ST",
        Name: "Sao Tome and Principe",
        Population: 48806,
    },
    Country{
        Code: "SA",
        Name: "Saudi Arabia",
        Population: 17397179,
    },
    Country{
        Code: "SN",
        Name: "Senegal",
        Population: 3194190,
    },
    Country{
        Code: "RS",
        Name: "Serbia",
        Population: 4705141,
    },
    Country{
        Code: "SC",
        Name: "Seychelles",
        Population: 50220,
    },
    Country{
        Code: "SL",
        Name: "Sierra Leone",
        Population: 92232,
    },
    Country{
        Code: "SG",
        Name: "Singapore",
        Population: 4453859,
    },
    Country{
        Code: "SK",
        Name: "Slovakia",
        Population: 4507849,
    },
    Country{
        Code: "SI",
        Name: "Slovenia",
        Population: 1501039,
    },
    Country{
        Code: "SB",
        Name: "Solomon Islands",
        Population: 43623,
    },
    Country{
        Code: "SO",
        Name: "Somalia",
        Population: 163185,
    },
    Country{
        Code: "ZA",
        Name: "South Africa",
        Population: 24909854,
    },
    Country{
        Code: "ES",
        Name: "Spain",
        Population: 35010273,
    },
    Country{
        Code: "LK",
        Name: "Sri Lanka",
        Population: 4267507,
    },
    Country{
        Code: "SD",
        Name: "Sudan",
        Population: 9307189,
    },
    Country{
        Code: "SR",
        Name: "Suriname",
        Population: 201963,
    },
    Country{
        Code: "SZ",
        Name: "Swaziland",
        Population: 301211,
    },
    Country{
        Code: "SE",
        Name: "Sweden",
        Population: 8581261,
    },
    Country{
        Code: "CH",
        Name: "Switzerland",
        Population: 7180749,
    },
    Country{
        Code: "SY",
        Name: "Syria",
        Population: 5860788,
    },
    Country{
        Code: "TJ",
        Name: "Tajikistan",
        Population: 1357400,
    },
    Country{
        Code: "TZ",
        Name: "Tanzania",
        Population: 7590794,
    },
    Country{
        Code: "TH",
        Name: "Thailand",
        Population: 19386154,
    },
    Country{
        Code: "TL",
        Name: "Timor-Leste",
        Population: 11472,
    },
    Country{
        Code: "TG",
        Name: "Togo",
        Population: 319822,
    },
    Country{
        Code: "TO",
        Name: "Tonga",
        Population: 40131,
    },
    Country{
        Code: "TT",
        Name: "Trinidad and Tobago",
        Population: 856544,
    },
    Country{
        Code: "TN",
        Name: "Tunisia",
        Population: 5053704,
    },
    Country{
        Code: "TR",
        Name: "Turkey",
        Population: 35358888,
    },
    Country{
        Code: "TM",
        Name: "Turkmenistan",
        Population: 424855,
    },
    Country{
        Code: "TV",
        Name: "Tuvalu",
        Population: 3768,
    },
    Country{
        Code: "UG",
        Name: "Uganda",
        Population: 6523949,
    },
    Country{
        Code: "UA",
        Name: "Ukraine",
        Population: 16849008,
    },
    Country{
        Code: "AE",
        Name: "United Arab Emirates",
        Population: 8807226,
    },
    Country{
        Code: "GB",
        Name: "United Kingdom",
        Population: 57075826,
    },
    Country{
        Code: "US",
        Name: "United States",
        Population: 279834232,
    },
    Country{
        Code: "UY",
        Name: "Uruguay",
        Population: 2017280,
    },
    Country{
        Code: "UZ",
        Name: "Uzbekistan",
        Population: 11914665,
    },
    Country{
        Code: "VU",
        Name: "Vanuatu",
        Population: 29791,
    },
    Country{
        Code: "VE",
        Name: "Venezuela",
        Population: 14548421,
    },
    Country{
        Code: "VN",
        Name: "Viet Nam",
        Population: 39772424,
    },
    Country{
        Code: "YE",
        Name: "Yemen",
        Population: 4778488,
    },
    Country{
        Code: "ZM",
        Name: "Zambia",
        Population: 2313013,
    },
    Country{
        Code: "ZW",
        Name: "Zimbabwe",
        Population: 2852757,
    },
}
