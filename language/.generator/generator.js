const https = require("https");
const fs = require("fs").promises;
const transliterate = require("transliteration");
const yaml = require("yaml");

const Alpha2KeyName = "alpha2"
const Alpha3KeyName = "alpha3-b"
const LanguageNameKey = "English"
const maxDownloadSize = 1000000 // 1Mb

getData("https://datahub.io/core/language-codes/r/language-codes-3b2.json", maxDownloadSize)
    .then((codes) => {
        let languages = prepareLanguageCodes(codes);
        return Promise.all(prepareFilesContent(languages))
            .then(() => console.log("Done"))
            .catch(() => console.log("Not Done :`("))
    })
    .catch((err) => console.error("woooooooops, seems like an error occurred...\n", err))

function getLanguages(url, maxSize, resolve, reject) {
    https.get(url, (res) => {
        if(res.statusCode === 301 || res.statusCode === 302) {
            return getLanguages(res.headers.location, maxSize, resolve, reject)
        }

        if (res.statusCode !== 200) {
            reject(`got ${res.statusCode} response status code`);
            return
        }

        let body = [], alreadyDownloaded = 0;

        res.on("data", (chunk) => {
            alreadyDownloaded += chunk.length;

            if (alreadyDownloaded > maxDownloadSize) {
                reject(`max download size of ${maxDownloadSize} bytes reached. it's better to search for less overhead data source`);
                return
            }

            body.push(chunk);
        });

        res.on("end", () => {
            try {
                // remove JSON.parse(...) for plain data
                resolve(JSON.parse(Buffer.concat(body).toString()));
            } catch (err) {
                reject(err);
            }
        });
    });
}

async function getData(url, maxSize) {
    return new Promise((resolve, reject) => getLanguages(url, maxSize, resolve, reject))
}

function prepareLanguageCodes(codes) {
    return codes.map((code) => {
        let result = {};
        result.a2 = {
            key: `Alpha2${code[Alpha2KeyName]}`,
            value: code[Alpha2KeyName].toLowerCase(),
        };

        result.a3 = {
            key: `Alpha3${code[Alpha3KeyName]}`,
            value: code[Alpha3KeyName].toLowerCase(),
        };

        let languageName = code[LanguageNameKey];
        let semicolonSplitted = languageName.split(";");
        if (semicolonSplitted.length > 1) {
            languageName = semicolonSplitted[0];
        }

        let parenthesisSplitted = languageName.split("(");
        if (parenthesisSplitted.length > 1) {
            languageName = parenthesisSplitted[0].trim(" ")
        }

        let commaSplitted = languageName.split(",");
        if (commaSplitted.length > 1) {
            languageName = commaSplitted.join("");
        }

        result.name = {value: languageName};

        let enumLanguageKey = languageName;

        enumLanguageKey = transliterate.transliterate(enumLanguageKey);
        enumLanguageKey = enumLanguageKey.replace(",", "");
        enumLanguageKey = enumLanguageKey.replace(".", "");
        enumLanguageKey = enumLanguageKey.replace("'", "");
        enumLanguageKey = enumLanguageKey.replace("-", "");
        enumLanguageKey = enumLanguageKey.replace("*", "");
        enumLanguageKey = enumLanguageKey.replace(/( ?)\(.*\)( ?)/, "");
        enumLanguageKey = enumLanguageKey.replace(/( ?)\[.*\]( ?)/, "");
        enumLanguageKey = enumLanguageKey
            .split(" ")
            .map((str) => `${str[0].toUpperCase()}${str.slice(1)}`)
            .join("");

        result.name.key = `Name${enumLanguageKey}`;
        result.key = enumLanguageKey;

        return result;
    });
}

const prepareFilesContent = (codes) => {
const nameTemplate = `package name

import "github.com/mikekonan/go-types/v2/language"

const (
${codes.map(
        (code) => `\t// Represents '${code.name.value}' country name
    ${code.key} = language.Name("${code.name.value}")`
    ).join("\n")}
)
`;

const alpha2Template = `package alpha2

import "github.com/mikekonan/go-types/v2/language"

const (
${codes.map(
        (code) => `\t// ${code.a2.value} represents '${code.a2.value}' country alpha-2 code
    ${code.a2.value} = language.Alpha2Code("${code.a2.value}")`
    ).join("\n")}
)
`;

const alpha3Template = `package alpha3

import "github.com/mikekonan/go-types/v2/language"

const (
${codes.map(
        (code) => `\t// ${code.a3.value} represents '${code.a3.value}' country alpha-3 code
    ${code.a3.value} = language.Alpha3Code("${code.a3.value}")`
    ).join("\n")}
)
`;

const entitiesTemplate = `package language

var (
${codes.map(
        (code) => `\t// ${code.key} represents '${code.name.value}' language
    ${code.key} = Language{
        name:    \"${code.name.value}\",
        alpha2:  \"${code.a2.value}\",
        alpha3:  \"${code.a3.value}\",
    }`
    )
        .join("\n")}
)
`;

const spec = {
    openapi: "3.0.0",
    components: {
        schemas: {
            LanguageName: {
                example: "Norwegian",
                type: "string",
                format: "language-name",
                enum: [...new Set(codes.map((code) => code.name.value))],
                "x-go-type": "github.com/mikekonan/go-types/v2/language.Name",
            },
            CountryAlpha2: {
                example: "no",
                type: "string",
                format: "iso639-1",
                minLength: 2,
                maxLength: 2,
                enum: [...new Set(codes.map((code) => code.a2.value))],
                "x-go-type": "github.com/mikekonan/go-types/v2/language.Alpha2Code",
            },
            CountryAlpha3: {
                example: "nor",
                type: "string",
                format: "iso639-2",
                minLength: 3,
                maxLength: 3,
                enum: [...new Set(codes.map((code) => code.a3.value))],
                "x-go-type": "github.com/mikekonan/go-types/v2/language.Alpha3Code",
            }
        },
    },
};

const languageByPropertyTemplate = `package language

var LanguageByName = map[string]Language{
${codes.map((code) => `\t\"${code.name.value}\" : ${code.key}`).join(",\n")},
}

var LanguageByAlpha2 = map[string]Language{
${codes.map((code) => `\t\"${code.a2.value}\" : ${code.key}`).join(",\n")},
}

var LanguageByAlpha3 = map[string]Language{
${codes.map((code) => `\t\"${code.a3.value}\" : ${code.key}`).join(",\n")},
}
`;

    return [
        fs.writeFile("../name/name_gen.go", nameTemplate, {
            encoding: "utf8",
            flag: "w",
        }),
        fs.writeFile("../alpha2/alpha_2_gen.go", alpha2Template, {
            encoding: "utf8",
            flag: "w",
        }),
        fs.writeFile("../alpha3/alpha_3_gen.go", alpha3Template, {
            encoding: "utf8",
            flag: "w",
        }),
        fs.writeFile("../entities_gen.go", entitiesTemplate, {
            encoding: "utf8",
            flag: "w",
        }),
        fs.writeFile("../language_mapping_gen.go", languageByPropertyTemplate, {
            encoding: "utf8",
            flag: "w",
        }),
        fs.writeFile("../swagger_gen.yaml", yaml.stringify(spec), {
            encoding: "utf8",
            flag: "w",
        })
    ];
}
