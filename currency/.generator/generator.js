const request = require("sync-request");
const xml2js = require("xml2js");
const yaml = require("yaml");
const fs = require("fs");

let resp = request(
    "GET",
    "https://www.six-group.com/dam/download/financial-information/data-center/iso-currrency/lists/list-one.xml"
);

if (resp.statusCode !== 200) {
    console.log(`unexpected status code - '${resp.statusCode}'`);
    os.exit(1);
}

let currencyCodes = (currencies) => {
    let temp = {}
    currencies.forEach((c) => {
        if (!!!temp[c.code]) {
            temp[c.code] = {
                code: c.code,
            };
        }
    })

    return temp
}

let currenciesByCodeMap = (currencies) => {
    let temp = {}
    currencies.forEach((c) => {
        if (!!!temp[c.code]) {
            temp[c.code] = {
                countries: [c.country],
                currency: c.currency,
                code: c.code,
                number: c.number,
                decimalPlaces: c.decimalPlaces,
            };

            return;
        }

        temp[c.code].countries.push(c.country);
    });

    return temp
}

let currenciesByNumberMap = (currencies) => {
    let temp = {}
    currencies.forEach((c) => {
        if (!!!temp[c.number]) {
            temp[c.number] = {
                countries: [c.country],
                currency: c.currency,
                code: c.code,
                number: c.number,
                decimalPlaces: c.decimalPlaces,
            };

            return;
        }

        temp[c.number].countries.push(c.country);
    });

    return temp
}

let currenciesByCountryMap = (currencies) => {
    let temp = {}
    currencies.forEach((c) => {
        if (!!temp[c.country]) {
            temp[c.country].currencies.push(c);
            return;
        }

        temp[c.country] = {
            currencies: [c],
        };

        return;
    });

    return temp
}

let currenciesByCurrencyMap = (currencies) => {
    let temp = {}
    currencies.forEach((c) => {
        if (!!!temp[c.currency]) {
            temp[c.currency] = {
                countries: [c.country],
                currency: c.currency,
                code: c.code,
                number: c.number,
                decimalPlaces: c.decimalPlaces,
            };

            return;
        }

        temp[c.currency].countries.push(c.country);
    });

    return temp
}

let renderByCountry = (currenciesMap, currenciesByCodeMap) => {
    return Object.keys(currenciesMap).map(
        (key) => `\`${key}\`: {
          ${currenciesMap[key].currencies.map((c) => {
        return `{
              countries:   Countries{${currenciesByCodeMap[c.code].countries
              .map((country) => `\`${country}\``).join(", ")}},
              currency:    \`${c.currency}\`,
              code:        \`${c.code}\`,
              number:      \`${c.number}\`,
              decimalPlaces: ${c.decimalPlaces},
          }`;
        })
        .join(`,`)},
    }`
        )
        .join(`,`);
};

let renderCodes = (currenciesMap) => {
    return Object.keys(currenciesMap).map(
        (key) => `\t// ${key} represents '${key}' currency code
    ${key} = currency.Code("${key}")`
    ).join("\n");
};

let render = (currenciesMap) => {
    return Object.keys(currenciesMap).map(
        (key) => `\`${key}\`: {
      countries:  Countries{${currenciesMap[key].countries
            .map((country) => `\`${country}\``)
            .join(", ")}},
      currency:     \`${currenciesMap[key].currency}\`,
      code:         \`${currenciesMap[key].code}\`,
      number:       \`${currenciesMap[key].number}\`,
      decimalPlaces: ${currenciesMap[key].decimalPlaces},
    }`
    ).join(`,
    `);
};

let renderCurrencyCodesTemplate = (currencies) => {
    return `package code

import "github.com/mikekonan/go-types/v2/currency"

const (
${renderCodes(currencyCodes(currencies))}
)`
};

let renderCurrencyMappingTemplate = (currencies) => {
    return `package currency

var currenciesByCode = map[string]currency{
  ${render(currenciesByCodeMap(currencies))},
}

var currenciesByNumber = map[string]currency{
  ${render(currenciesByNumberMap(currencies))},
}

var currenciesByCountry = map[string]currencies{
  ${renderByCountry(currenciesByCountryMap(currencies), currenciesByCodeMap(currencies))},
}

var currenciesByCurrency = map[string]currency{
  ${render(currenciesByCurrencyMap(currencies))},
}`;
};

let normalizeCurrency = (currency) => {
    let normalizeCountry;

    normalizeCountry = (country) => {
        let bracketsFound = false;

        if (country.includes(" (") && country.includes(")")) {
            bracketsFound = true;
            let lastIndexOfOpeningBrackets = country.lastIndexOf(" (");
            let lastIndexOfClosingBrackets = country.lastIndexOf(")");
            let bracked = country.substring(
                lastIndexOfOpeningBrackets,
                lastIndexOfClosingBrackets + 1
            );
            country = country.replace(bracked, "");

            for (let nested = true; nested;) {
                let result = normalizeCountry(country);
                if (result.bracketsFound) {
                    country = result.country;
                }

                nested = result.bracketsFound;
            }
        }

        return {country: country, bracketsFound: bracketsFound};
    };

    let normalizedResult = normalizeCountry(currency.country);
    if (normalizedResult.bracketsFound) {
        return {
            country: normalizedResult.country,
            currency: currency.currency,
            code: currency.code,
            number: currency.number,
            decimalPlaces: currency.decimalPlaces,
        };
    }

    return null;
};

let goCodePromise = xml2js
    .parseStringPromise(resp.body.toString(), {mergeAttrs: true})
    .then((result) => {
        return Promise.resolve(result["ISO_4217"]["CcyTbl"][0]["CcyNtry"]);
    })
    .catch((err) => {
        console.log(err);
        os.exit(1);
    })
    .then((currencies) => {
        let result = currencies.map((row) => {
            if (!!!row["Ccy"]) {
                return;
            }

            return {
                country: row["CtryNm"][0],
                currency:
                    typeof row["CcyNm"][0] !== "object"
                        ? row["CcyNm"][0]
                        : row["CcyNm"][0]["_"],
                code: row["Ccy"][0],
                number: row["CcyNbr"][0],
                decimalPlaces: isNaN(row["CcyMnrUnts"]) ? 0 : parseInt(row["CcyMnrUnts"]),
            };
        });

        result = result
            .filter((r) => !!r)
            .filter((r) => !r.country.startsWith("ZZ"));

        let normalizedCountries = [];

        result.forEach((currency) => {
            let normalized = normalizeCurrency(currency);
            if (!!normalized) {
                normalizedCountries.push(normalized);
            }
        });

        return result.concat(normalizedCountries);
    })
    .then((currencies) => {
            fs.writeFileSync("../code/code_gen.go", renderCurrencyCodesTemplate(currencies))
            fs.writeFileSync("../currencies_mapping_gen.go", renderCurrencyMappingTemplate(currencies))
    });

let oas3Promise = xml2js
    .parseStringPromise(resp.body.toString(), {mergeAttrs: true})
    .then((result) => {
        return Promise.resolve(result["ISO_4217"]["CcyTbl"][0]["CcyNtry"]);
    })
    .catch((err) => {
        console.log(err);
        os.exit(1);
    })
    .then((currencies) => {
        let result = currencies.map((row) => {
            if (!!!row["Ccy"]) {
                return;
            }

            return {
                country: row["CtryNm"][0],
                currency:
                    typeof row["CcyNm"][0] !== "object"
                        ? row["CcyNm"][0]
                        : row["CcyNm"][0]["_"],
                code: row["Ccy"][0],
                number: row["CcyNbr"][0],
            };
        });
        result = result
            .filter((r) => !!r)
            .filter((r) => !r.country.startsWith("ZZ"));

        const spec = {
            openapi: "3.0.0",
            components: {
                schemas: {
                    CurrencyCode: {
                        example: "EUR",
                        type: "string",
                        minLength: 3,
                        maxLength: 3,
                        format: "iso4217-currency-code",
                        enum: [...new Set(result.map((currency) => currency.code))],
                        "x-go-type": "github.com/mikekonan/go-types/v2/currency.Code",
                    },
                    CurrencyName: {
                        example: "Euro",
                        type: "string",
                        enum: [...new Set(result.map((currency) => currency.currency))],
                        "x-go-type": "github.com/mikekonan/go-types/v2/currency.Currency",
                    },
                    CurrencyCountry: {
                        example: "PUERTO RICO",
                        type: "string",
                        enum: [...new Set(result.map((currency) => currency.country))],
                        "x-go-type": "github.com/mikekonan/go-types/v2/currency.Country",
                    },
                    CurrencyNumber: {
                        example: "840",
                        type: "string",
                        minLength: 3,
                        maxLength: 3,
                        enum: [...new Set(result.map((currency) => currency.number))],
                    },
                },
            },
        };

        return spec;
    })
    .then((spec) =>
        fs.writeFileSync("../swagger_gen.yaml", yaml.stringify(spec))
    );

Promise.all([goCodePromise, oas3Promise]).then(() => console.log("Done"));
