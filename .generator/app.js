const yaml = require("yaml");
const glob = require('glob-promise');
const fs = require('fs');

glob.promise(__dirname + '/../*/*.yaml')
    .then(files => {
        console.log(files);
        return files.map(f => fs.readFileSync(f, "utf8"))
    })
    .then(files => files.map(f => yaml.parse(f)))
    .then(files => {
        let schemasArr = files.map(f => Object.keys(f.components.schemas)
            .map(k => Object.fromEntries([[k, f.components.schemas[k]]]))).flat(1);
        let schemas = Object.fromEntries(schemasArr.map(schema => [Object.keys(schema)[0], schema[Object.keys(schema)[0]]]));

        return {openapi: "3.0.0", components: {schemas: schemas}};
    })
    .then(result => yaml.stringify(result))
    .then(result => fs.writeFileSync("../swagger.yaml", result, {encoding: "utf8", flag: "w"}))
