var express = require('express');
var router = express.Router();
const axios = require('axios');

// const fs = require('fs');
// const crypto = require("crypto").webcrypto;
// globalThis.crypto = crypto;
// require('../golang/wasm/wasm_exec.js');

// function loadWebAssembly() {
//   const wasmModule = fs.readFileSync('./golang/wasm/main.wasm');
//     const go = new Go();
//     const importObject = go.importObject;
//     WebAssembly.instantiate(wasmModule, importObject).then((results) => {
//       const instance = results.instance
//       go.run(instance);
//     });
// }

// loadWebAssembly();

// Function to fetch data from localhost:9000/poems
async function getPoems() {
  try {
    const response = await axios.get('http://localhost:9000/poems');
    return response.data;
  } catch (error) {
    console.error(error);
    throw error;
  }
}

/* GET home page. */
router.get('/', async function(req, res, next) {
  // Start the server
  // a = 3;
  // b = 3;
  // const sum = addTwoNumbers(a, b);
  console.log("POEMS", await getPoems());
  // console.log("SUM", sum)
  res.render('index', { poems:  await getPoems()});
});

module.exports = router;