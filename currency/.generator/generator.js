const request = require("sync-request");
const xml2js = require("xml2js");
const yaml = require("yaml");
const fs = require("fs");

let resp = request(
    "GET",
    "https://www.currency-iso.org/dam/downloads/lists/list_one.xml"
);

if (resp.statusCode !== 200) {
    console.log(`unexpected status code - '${resp.statusCode}'`);
    os.exit(1);
}

let render = (currencies) => {
    let currenciesByCodeMap = {};
    let currenciesByNumberMap = {};
    let currenciesByCountryMap = {};
    let currenciesByCurrencyMap = {};

    currencies.forEach((c) => {
        if (!!!currenciesByCodeMap[c.code]) {
            currenciesByCodeMap[c.code] = {
                countries: [c.country],
                currency: c.currency,
                code: c.code,
                number: c.number,
            };

            return;
        }

        currenciesByCodeMap[c.code].countries.push(c.country);
    });

    currencies.forEach((c) => {
        if (!!!currenciesByNumberMap[c.number]) {
            currenciesByNumberMap[c.number] = {
                countries: [c.country],
                currency: c.currency,
                code: c.code,
                number: c.number,
            };

            return;
        }

        currenciesByNumberMap[c.number].countries.push(c.country);
    });

    currencies.forEach((c) => {
        if (!!currenciesByCountryMap[c.country]) {
            currenciesByCountryMap[c.country].currencies.push(c);
            return;
        }

        currenciesByCountryMap[c.country] = {
            currencies: [c],
        };

        return;
    });

    currencies.forEach((c) => {
        if (!!!currenciesByCurrencyMap[c.currency]) {
            currenciesByCurrencyMap[c.currency] = {
                countries: [c.country],
                currency: c.currency,
                code: c.code,
                number: c.number,
            };

            return;
        }

        currenciesByCurrencyMap[c.currency].countries.push(c.country);
    });

    let renderByCountry = (currenciesMap) => {
        return Object.keys(currenciesMap)
            .map(
                (key) => `\`${key}\`: {
                      ${currenciesMap[key].currencies
                    .map((c) => {
                        return `{
                        countries:  Countries{${currenciesByCodeMap[
                            c.code
                            ].countries
                            .map((country) => `\`${country}\``)
                            .join(", ")}},
                        currency:   \`${c.currency}\`,
                        code:       \`${c.code}\`,
                        number:     \`${c.number}\`,
                        }`;
                    })
                    .join(`,`)},
                    }`
            )
            .join(`,`);
    };

    let render = (currenciesMap) => {
        return Object.keys(currenciesMap).map(
            (key) => `\`${key}\`: {
      countries:  Countries{${currenciesMap[key].countries
                .map((country) => `\`${country}\``)
                .join(", ")}},
      currency:   \`${currenciesMap[key].currency}\`,
      code:       \`${currenciesMap[key].code}\`,
      number:     \`${currenciesMap[key].number}\`,
    }`
        ).join(`,
    `);
    };

    let template = `package currency

var currenciesByCode = map[string]currency{
    ${render(currenciesByCodeMap)},
}

var currenciesByNumber = map[string]currency{
  ${render(currenciesByNumberMap)},
}

var currenciesByCountry = map[string]currencies{
  ${renderByCountry(currenciesByCountryMap)},
}

var currenciesByCurrency = map[string]currency{
  ${render(currenciesByCurrencyMap)},
}
`;

    return template;
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
    .then((currencies) =>
        fs.writeFileSync("../currency_gen.go", render(currencies))
    );

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
                        format: "iso4217-currency-code",
                        enum: [...new Set(result.map((currency) => currency.code))],
                        "x-go-type": "github.com/mikekonan/go-types/currency.Code",
                    },
                    CurrencyName: {
                        example: "Euro",
                        type: "string",
                        enum: [...new Set(result.map((currency) => currency.currency))],
                        "x-go-type": "github.com/mikekonan/go-types/currency.Currency",
                    },
                    CurrencyCountry: {
                        example: "PUERTO RICO",
                        type: "string",
                        enum: [...new Set(result.map((currency) => currency.country))],
                        "x-go-type": "github.com/mikekonan/go-types/currency.Country",
                    },
                    CurrencyNumber: {
                        example: "840",
                        type: "string",
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
