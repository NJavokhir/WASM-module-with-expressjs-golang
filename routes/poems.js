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
//   const sum = addTwoNumbers(a, b);
  poems = getPoems();
  console.log(poems)
  res.render('index', { title: b });
  // res.send(`Sum: ${sum}`);
});

module.exports = router;