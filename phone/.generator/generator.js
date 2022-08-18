const https = require("https");
const fs = require("fs").promises;
const yaml = require("yaml");

const requestCountryDialKey = "Dial";
const requestCountryKey = "ISO3166-1-Alpha-2";

const innerCountryDialKey = "dial";
const innerCountryKey = "countryAlpha2";

const options = {
    hostname: "datahub.io",
    port: 443,
    path: "/core/country-codes/r/country-codes.json",
    method: "GET",
    headers: {
        "Content-Type": "application/json"
    },
};
const maxDownloadSize = 1000000 // 1Mb

getPhoneCountryCodes(options, maxDownloadSize)
    .then((codes) => {
        let countryCodes = preparePhoneCountryCodes(codes);
        // console.log(concatenateCountriesByDialCode(countryCodes))
        Promise.all([
            fs.writeFile("../dial_code_gen.go", renderCountryDialCodesTemplate(countryCodes)),
            fs.writeFile("../codes_mapping_gen.go", renderDialCodesMappingTemplate(countryCodes)),
            fs.writeFile("../swagger_gen.yaml", renderSpecification(countryCodes), {
                encoding: "utf8",
                flag: "w",
            })
        ]).then(() => console.log("Done"))
            .catch(err => console.log("Not Done :`(\nerr: ", err))
    })
    .catch((err) => console.error("woooooooops, seems like an error occurred...\n", err))

function getPhoneCountryCodes(options, maxDownloadSize) {
    return new Promise((resolve, reject) => {
        const req = https.request(options, res => {
            let body = "", alreadyDownloaded = 0;

            res.on("data", chunk => {
                alreadyDownloaded += chunk.length;

                if (alreadyDownloaded > maxDownloadSize) {
                    reject(`max download size of ${maxDownloadSize} bytes reached. it's better to search for less overhead data source`);
                    return
                }

                body += chunk;
            });

            res.on("end", () => {
                if (res.statusCode !== 200) {
                    reject(`got ${res.statusCode} response status code`);
                    return
                }

                if (alreadyDownloaded > maxDownloadSize) {
                    return
                }

                const parsedBody = JSON.parse(body);

                resolve(parsedBody);
            });

            req.on("error", error => {
                reject(error);
            });
        });

        req.end();
    })
}

function preparePhoneCountryCodes(rawCodes) {
    return rawCodes.reduce((result, codes) => {
        // check if there is null/undefined in request result dial country code value
        if (
            !(!codes[requestCountryDialKey] || !codes[requestCountryKey])
        ) {
            // check if there is countries with few dial codes (f.e. Dominican Republic)
            let codesByCountry = codes[requestCountryDialKey].split(",")
            if (codesByCountry.length > 1) {
                codesByCountry.forEach(code => {
                    result.push({ [innerCountryDialKey]: code, [innerCountryKey]: codes[requestCountryKey] })
                })
            } else {
                result[result.length] = { [innerCountryDialKey]: codes[requestCountryDialKey], [innerCountryKey]: codes[requestCountryKey] }
            }
        }

        return result
    }, []);
}

function renderCountryDialCodesTemplate(codes) {
return `package phone

const (
${renderByAlpha2Codes(trimRedundantDominicana(codes))}
)`
}

function trimRedundantDominicana(codes) {
    let redundant = ["1-809", "1-849"]
    return codes.filter(({ [innerCountryDialKey]: dialCode }) => !redundant.some(el => dialCode === el))
}

function renderByAlpha2Codes(dialCodes) {
    return dialCodes.map((codes) => `\t// ${codes[innerCountryDialKey]} is '${codes[innerCountryKey]}' country dial code
    ${codes[innerCountryKey]} = DialCode("${codes[innerCountryDialKey]}")`).join("\n");
}

function renderDialCodesMappingTemplate(codes) {
return `package phone

import "github.com/mikekonan/go-types/v2/country"

var dialCodeByCountryAlpha2Str = map[string]DialCode{
${renderByCountryMapping(trimRedundantDominicana(codes))}
}
    
var countryCodeByDialString = map[string][]country.Alpha2Code{
${renderByDialMapping(concatenateCountriesByDialCode(codes))}
}`
}

function renderByCountryMapping(rawCodes) {
    return rawCodes.map((codes) => `\t\"${codes[innerCountryKey]}\": ${codes[innerCountryKey]},`).join(`\n`);
}

function renderByDialMapping(rawCodes) {
    return Object.keys(rawCodes).map(dialCode => `\t\"${dialCode}\": {\n\t${rawCodes[dialCode].map(countryCode => `\tcountry.Alpha2Code("${countryCode}"),\n\t`).join(``)}},`).join(`\n`);
}

// there are countries with identical dial codes, so, concat them by dial code
function concatenateCountriesByDialCode(raw) {
    return raw.reduce((result, codes) => {
        if (result.hasOwnProperty(codes[innerCountryDialKey])) {
            result[codes[innerCountryDialKey]].push(codes[innerCountryKey])
        } else {
            result[codes[innerCountryDialKey]] = [codes[innerCountryKey]]
        }

        return result
    }, {})
}

function renderSpecification(codes) {
    const spec = {
        openapi: "3.0.0",
        components: {
            schemas: {
                DialCountryCode: {
                    example: "1-829",
                    type: "string",
                    format: "E.164",
                    enum: [...Object.keys(concatenateCountriesByDialCode(codes))],
                    "x-go-type": "github.com/mikekonan/go-types/v2/phone.DialCode",
                }
            },
        },
    };

    return yaml.stringify(spec)
}
