const fs = require("fs");
const transliterate = require("transliteration");
const yaml = require("yaml");

//For testing purpose we can use raw/countries
let codes = JSON.parse(fs.readFileSync(0, "utf8"));

codes = codes.map((code) => {
    let result = {};
    result.a2 = {
        key: `Alpha2${code["Alpha-2 code"]}`,
        value: code["Alpha-2 code"],
    };

    result.a3 = {
        key: `Alpha3${code["Alpha-3 code"]}`,
        value: code["Alpha-3 code"],
    };

    result.name = {value: code["English short name"]};

    let enumCountryKey = code["English short name"];

    if (enumCountryKey === "Korea (the Republic of)") {
        enumCountryKey = "South Korea";
    }

    switch (enumCountryKey) {
        case "Korea (the Democratic People's Republic of)": {
            enumCountryKey = "North Korea";
            break;
        }
        case "Korea (the Republic of)": {
            enumCountryKey = "South Korea";
            break;
        }
        case "Virgin Islands (British)": {
            enumCountryKey = "British Virgin Islands";
            break;
        }
        case "Virgin Islands (U.S.)": {
            enumCountryKey = "US Virgin Islands";
            break;
        }
        case "Palestine, State of": {
            enumCountryKey = "Palestine";
            break;
        }
        case "Congo (the Democratic Republic of the)": {
            enumCountryKey = "Democratic Republic of the Congo";
            break;
        }
    }

    enumCountryKey = transliterate.transliterate(enumCountryKey);
    enumCountryKey = enumCountryKey.replace(",", "");
    enumCountryKey = enumCountryKey.replace(".", "");
    enumCountryKey = enumCountryKey.replace("'", "");
    enumCountryKey = enumCountryKey.replace("-", "");
    enumCountryKey = enumCountryKey.replace("*", "");
    enumCountryKey = enumCountryKey.replace(/( ?)\(.*\)( ?)/, "");
    enumCountryKey = enumCountryKey.replace(/( ?)\[.*\]( ?)/, "");
    enumCountryKey = enumCountryKey
        .split(" ")
        .map((str) => `${str[0].toUpperCase()}${str.slice(1)}`)
        .join("");

    result.name.key = `Name${enumCountryKey}`;
    result.key = enumCountryKey;

    return result;
});

const countriesTemplate = `package name

import "github.com/mikekonan/go-types/v2/country"

const (
${codes.map(
        (code) => `\t// ${code.key} represents '${code.name.value}' country name
    ${code.key} = country.Name("${code.name.value}")`
    ).join("\n")}
) 
`;

const alpha2Template = `package alpha2

import "github.com/mikekonan/go-types/v2/country"

const (
${codes.map(
        (code) => `\t// ${code.a2.value} represents '${code.a2.value}' country alpha-2 code
    ${code.a2.value} = country.Alpha2Code("${code.a2.value}")`
    ).join("\n")}
) 
`;

const alpha3Template = `package alpha3

import "github.com/mikekonan/go-types/v2/country"

const (
${codes.map(
        (code) => `\t// ${code.a3.value} represents '${code.a3.value}' country alpha-3 code
    ${code.a3.value} = country.Alpha3Code("${code.a3.value}")`
    ).join("\n")}
) 
`;

const entitiesTemplate = `package country

var (
${codes.map(
    (code) => `\t// ${code.key} represents '${code.name.value}' country
    ${code.key} = Country{
		name:    \"${code.name.value}\",
		alpha2:  \"${code.a2.value}\",
		alpha3:  \"${code.a3.value}\",
	}`
)
    .join("\n")}
)
`;

const countryByCountryTemplate = `package country

var CountryByName = map[string]Country{
${codes.map((code) => `\t\"${code.name.value}\" : ${code.key}`).join(",\n")},
}

var CountryByAlpha2 = map[string]Country{
${codes.map((code) => `\t\"${code.a2.value}\" : ${code.key}`).join(",\n")},
}

var CountryByAlpha3 = map[string]Country{
${codes.map((code) => `\t\"${code.a3.value}\" : ${code.key}`).join(",\n")},
}
`;

const spec = {
    openapi: "3.0.0",
    components: {
        schemas: {
            CountryName: {
                example: "Austria",
                type: "string",
                format: "iso3166-country",
                enum: [...new Set(codes.map((code) => code.name.value))],
                "x-go-type": "github.com/mikekonan/go-types/v2/country.Name",
            },
            CountryAlpha2: {
                example: "AS",
                type: "string",
                format: "iso3166-alpha-2",
                minLength: 2,
                maxLength: 2,
                enum: [...new Set(codes.map((code) => code.a2.value))],
                "x-go-type": "github.com/mikekonan/go-types/v2/country.Alpha2Code",
            },
            CountryAlpha3: {
                example: "USA",
                type: "string",
                format: "iso3166-alpha-3",
                minLength: 3,
                maxLength: 3,
                enum: [...new Set(codes.map((code) => code.a3.value))],
                "x-go-type": "github.com/mikekonan/go-types/v2/country.Alpha3Code",
            }
        },
    },
};

let writeFiles = [
    fs.writeFileSync("../name/name_gen.go", countriesTemplate, {
        encoding: "utf8",
        flag: "w",
    }),
    fs.writeFileSync("../alpha2/alpha_2_gen.go", alpha2Template, {
        encoding: "utf8",
        flag: "w",
    }),
    fs.writeFileSync("../alpha3/alpha_3_gen.go", alpha3Template, {
        encoding: "utf8",
        flag: "w",
    }),
    fs.writeFileSync("../entities_gen.go", entitiesTemplate, {
        encoding: "utf8",
        flag: "w",
    }),
    fs.writeFileSync("../country_mapping_gen.go", countryByCountryTemplate, {
        encoding: "utf8",
        flag: "w",
    }),
    fs.writeFileSync("../swagger_gen.yaml", yaml.stringify(spec), {
        encoding: "utf8",
        flag: "w",
    })
];

Promise.all(writeFiles).then(() => console.log("Done"));
