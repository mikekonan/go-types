const request = require("sync-request");
const xml2js = require("xml2js");
const yaml = require("yaml");
const fs = require("fs");

const currencyOverrides = {
    XTS: {
      country: '[TEST] XTS',
      decimalPlaces: 1,
    }
}

let actualCurrencies = request(
    "GET",
    "https://www.six-group.com/dam/download/financial-information/data-center/iso-currrency/lists/list-one.xml"
);

let historicalCurrencies = request(
    "GET",
    "https://www.six-group.com/dam/download/financial-information/data-center/iso-currrency/lists/list-three.xml"
);

if (actualCurrencies.statusCode !== 200 || historicalCurrencies.statusCode !== 200) {
    console.log(`unexpected status code - actual currencies: '${actualCurrencies.statusCode}', historical: ${historicalCurrencies.statusCode}`);
    os.exit(1);
}

const buildCurrency = (c) => (
    {
        countries: [c.country],
        currency: c.currency,
        code: c.code,
        number: c.number,
        decimalPlaces: c.decimalPlaces,
        isHistorical: c.isHistorical
    }
)
const currenciesMapWithCountries = (currencies, keyField, multi = false) => 
    currencies.reduce((acc, c) => {
        if (multi) {
            acc[c[keyField]] = [...(acc[c[keyField]] || []), buildCurrency(c)]
        } else {
            acc[c[keyField]] && acc[c[keyField]].countries.push(c.country) || (acc[c[keyField]] = buildCurrency(c))
        }
        return acc
    }, {})
const currenciesByCodeMap = (currencies) => currenciesMapWithCountries(currencies, 'code')
const currenciesByNumberMap = (currencies) => currenciesMapWithCountries(currencies, 'number')
const currenciesByCurrencyMap = (currencies) => currenciesMapWithCountries(currencies, 'currency')
const currenciesByCountryMap = (currencies) => currenciesMapWithCountries(currencies, 'country', true)

const currencyCodes = (currencies) => Object.keys(currenciesByCodeMap(currencies))
const renderCodes = (currenciesCodes) => 
    currenciesCodes.map(
        (code) => `\t// ${code} represents '${code}' currency code
    ${code} = currency.Code("${code}")`
    ).join("\n")

const renderCurrency = (key, currency, countries) =>
`${key ? `\`${key}\`: ` : ''}{
  countries:     Countries{${[...new Set(countries)].map((country) => `\`${country}\``).join(", ")}},
  currency:      \`${currency.currency}\`,
  code:          \`${currency.code}\`,
  number:        \`${currency.number}\`,
  decimalPlaces: ${currency.decimalPlaces},
  isHistorical:  ${currency.isHistorical ? 'true' : 'false'},
}`
const renderCurrenciesByKey = (currenciesMap) => 
    Object.entries(currenciesMap).map(([key, c]) => renderCurrency(key, c, c.countries))
    .join(`,\n`)
const renderCurrenciesByCountry = (currenciesByCountryMap, currenciesByCodeMap) =>
    Object.entries(currenciesByCountryMap).map(([countryCode, countries]) =>
`\`${countryCode}\`: {
    ${countries.map((c) => renderCurrency(null, c, currenciesByCodeMap[c.code].countries)).join(`, `)},
}`)
    .join(`, `)

const renderCurrencyCodesTemplate = (currencies) => {
    return `package code

import "github.com/mikekonan/go-types/v2/currency"

const (
${renderCodes(currencyCodes(currencies))}
)
`
};

const renderCurrencyMappingTemplate = (currencies) => {
    return `package currency

var currenciesByCode = map[string]currency{
  ${renderCurrenciesByKey(currenciesByCodeMap(currencies))},
}

var currenciesByNumber = map[string]currency{
  ${renderCurrenciesByKey(currenciesByNumberMap(currencies))},
}

var currenciesByCountry = map[string]currencies{
  ${renderCurrenciesByCountry(currenciesByCountryMap(currencies), currenciesByCodeMap(currencies))},
}

var currenciesByCurrency = map[string]currency{
  ${renderCurrenciesByKey(currenciesByCurrencyMap(currencies))},
}
`;
};

const normalizeCurrency = (currency) => {
    const normalizeCountry = (country) => {
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

        return { country: country, bracketsFound: bracketsFound };
    };

    const normalizedResult = normalizeCountry(currency.country);
    if (normalizedResult.bracketsFound) {
        return {
            country: normalizedResult.country,
            currency: currency.currency,
            code: currency.code,
            number: currency.number,
            decimalPlaces: currency.decimalPlaces,
            isHistorical: currency.isHistorical,
        };
    }
}

const parseCurrenciesFromXML = (rows, useOverride = true, isHistorical = false) => {
    return rows
        .map((row) => {
            if (!row["Ccy"]) {
                return
            }

            const currencyCode = row["Ccy"][0]
            const override = useOverride && !isHistorical && currencyOverrides[currencyCode]

            if (row["CcyMnrUnts"] && row["CcyMnrUnts"][0] === "N.A." && (!override || !override.decimalPlaces)) {
                return
            }

            if (row["CcyNm"][0] === "No universal currency" && (!override || !override.country)) {
                return
            }

            const parsedCurrency = {
                country: row["CtryNm"][0],
                currency:
                    typeof row["CcyNm"][0] !== "object"
                        ? row["CcyNm"][0]
                        : row["CcyNm"][0]["_"],
                code: row["Ccy"][0],
                number: row["CcyNbr"] && row["CcyNbr"][0] || 0,
                decimalPlaces: row["CcyMnrUnts"] && row["CcyMnrUnts"][0] || 0,
                isHistorical: isHistorical
            }

            if (override && typeof override === "object") {
                Object.entries(override).forEach(([key, value]) => {
                    parsedCurrency[key] = value;
                })
            }

            return parsedCurrency
        })
        .filter((r) => !!r)
        .filter((r) => !r.country.startsWith("ZZ"))
}

const parseXMLResponse = (response, isHistorical = false) =>
    xml2js
    .parseStringPromise(response, { mergeAttrs: true })
    .then(result => {
        const firstKey = isHistorical ? "HstrcCcyTbl" : "CcyTbl"
        const secondKey = isHistorical ? "HstrcCcyNtry" : "CcyNtry"
        return Promise.resolve(result["ISO_4217"][firstKey][0][secondKey])
    })
    .catch(err => {
        console.log(err)
        os.exit(1);
    })

let goCodePromise = 
    parseXMLResponse(actualCurrencies.body.toString())
    .then(actual => {
        return parseXMLResponse(historicalCurrencies.body.toString(), true)
            .then(historical => parseCurrenciesFromXML(historical, false, true))
            .then(historical => parseCurrenciesFromXML(actual).concat(historical))
    })
    .then((parsedCurrencies) => {
        let normalizedCountries = []

        parsedCurrencies.forEach((currency) => {
            let normalized = normalizeCurrency(currency)
            if (!!normalized) {
                normalizedCountries.push(normalized)
            }
        });

        return parsedCurrencies.concat(normalizedCountries)
    })
    .then((currencies) => {
        fs.writeFileSync("../code/code_gen.go", renderCurrencyCodesTemplate(currencies))
        fs.writeFileSync("../currencies_mapping_gen.go", renderCurrencyMappingTemplate(currencies))
    });

let oas3Promise = xml2js
    .parseStringPromise(actualCurrencies.body.toString(), {mergeAttrs: true})
    .then((result) => {
        return Promise.resolve(result["ISO_4217"]["CcyTbl"][0]["CcyNtry"]);
    })
    .catch((err) => {
        console.log(err);
        os.exit(1);
    })
    .then((currencies) => {
        let parsedCurrencies = parseCurrenciesFromXML(currencies, useOverride = false)

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
                        enum: [...new Set(parsedCurrencies.map((currency) => currency.code))],
                        "x-go-type": "github.com/mikekonan/go-types/v2/currency.Code",
                    },
                    CurrencyName: {
                        example: "Euro",
                        type: "string",
                        enum: [...new Set(parsedCurrencies.map((currency) => currency.currency))],
                        "x-go-type": "github.com/mikekonan/go-types/v2/currency.Currency",
                    },
                    CurrencyCountry: {
                        example: "PUERTO RICO",
                        type: "string",
                        enum: [...new Set(parsedCurrencies.map((currency) => currency.country))],
                        "x-go-type": "github.com/mikekonan/go-types/v2/currency.Country",
                    },
                    CurrencyNumber: {
                        example: "840",
                        type: "string",
                        minLength: 3,
                        maxLength: 3,
                        enum: [...new Set(parsedCurrencies.map((currency) => currency.number))],
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
