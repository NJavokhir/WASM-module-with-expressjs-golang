var express = require('express');
var router = express.Router();

const fs = require('fs');
const crypto = require("crypto").webcrypto;
globalThis.crypto = crypto;
require('../golang/wasm/wasm_exec.js');

function loadWebAssembly() {
  const wasmModule = fs.readFileSync('./golang/wasm/main.wasm');
    const go = new Go();
    const importObject = go.importObject;
    WebAssembly.instantiate(wasmModule, importObject).then((results) => {
      const instance = results.instance
      go.run(instance);
    });
}

loadWebAssembly();

/* GET home page. */
router.get('/', function(req, res, next) {
  a = 3;
  b = 3;
  const sum = addTwoNumbers(a, b);
  const poems = getPoems();
  console.log("POEMS", poems)
  console.log("SUM", sum)
  res.render('index', { title: sum });
  // res.send(`Sum: ${sum}`);
});

function getData() {
  return getPoems()
}

module.exports = router;