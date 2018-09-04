var fibos = require('fibos');
var fs = require("fs");
var config = {
    "producer-name": "{{#BP_ACCOUNT_NAME#}}",
    "public-key": "{{#BP_PUBLIC_KEY#}}",
    "private-key": "{{#BP_PRIVATE_KEY#}}"
};

console.notice("Starting up FIBOS name:", config["producer-name"]);
fibos.config_dir = config["producer-name"] + "_Dir";
fibos.data_dir = config["producer-name"] + "_Dir";

console.notice("config_dir:", fibos.config_dir);
console.notice("data_dir:", fibos.data_dir);

if (fs.exists(fibos.data_dir) || fs.exists(fibos.config_dir)) {

    console.notice("data_dir or config_dir is exists");
    fibos.load("chain");

} else {

    console.notice("I will run as a blank node");

    fibos.load("chain", {
        "genesis-json": "genesis.json",
        "delete-all-blocks": true,
    });

}
fibos.load("http", {
    "http-server-address": "127.0.0.1:65342",
    "access-control-allow-origin": "*",

    "http-validate-host": false
});

fibos.load("net", {
    "p2p-listen-endpoint": "0.0.0.0:9870",
    "p2p-peer-address": [
        "p2p-mainnet.fibos123.com:9977",
        "p2p.foshenzhenbp.com:9877",
        "fibos.qubitfund.com:9870",
        "api.fibosgenesis.com:9870",
        "se-p2p.fibos.io:9870",
        "seed.fibos.rocks:10100",
        "fibos-node.slowmist.io:9870",
        "p2p.mainnet.fibos.me:80",
        "fibos.qubitfund.com:9870",
        "p2p.eoschina.me:10300",
        "sl-p2p.fibos.io:9870",
        "to-p2p.fibos.io:9870",
        "ca-p2p.fibos.io:9870",
        "ln-p2p.fibos.io:9870",
        "va-p2p.fibos.io:9870"
    ]
});

fibos.load("producer", {
    'producer-name': config["producer-name"],
    'enable-stale-production': true,
    'private-key': JSON.stringify([config["public-key"], config["private-key"]])
});

fibos.load("chain");
fibos.load("chain_api");
fibos.load("wallet");
fibos.load("wallet_api");

fibos.start();