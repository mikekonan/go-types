const fs = require("fs");
const moment = require("moment-timezone");
const names = moment.tz.names()
const yaml = require("yaml");

let render = (names) => {

    let mapKV = names.map(key => `"${key}": Timezone("${key}"),`).join(`
    `);

    let template =
        `package timezone
        
        var timezonesByName = map[string]Timezone{
            ${mapKV}
        }`;

    return template
}

fs.writeFileSync("../timezone_gen.go", render(names))

const spec = {
    openapi: "3.0.0",
    components: {
        schemas: {
            Timezone: {
                example: "Europe/Minsk",
                type: "string",
                format: "rfc6557-time-zone",
                enum: [...new Set(names)],
                "x-go-type": "github.com/mikekonan/go-types/v2/timezone.Timezone",
            },
        },
    },
};

fs.writeFileSync("../swagger_gen.yaml", yaml.stringify(spec));
console.log("Done");