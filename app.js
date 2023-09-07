var createError = require('http-errors');
var express = require('express');
var path = require('path');
var cookieParser = require('cookie-parser');
var logger = require('morgan');

var indexRouter = require('./routes/index');
var usersRouter = require('./routes/users');

var app = express();

// view engine setup
app.set('views', path.join(__dirname, 'views'));
app.set('view engine', 'jade');

app.use(logger('dev'));
app.use(express.json());
app.use(express.urlencoded({ extended: false }));
app.use(cookieParser());
app.use(express.static(path.join(__dirname, 'public')));

const fs = require('fs');
const crypto = require("crypto").webcrypto;
globalThis.crypto = crypto;
require('./public/wasm/wasm_exec.js');

function loadWebAssembly() {
  const wasmModule = fs.readFileSync('./public/wasm/main.wasm');
    const go = new Go();
    const importObject = go.importObject;
    WebAssembly.instantiate(wasmModule, importObject).then((results) => {
      const instance = results.instance
      go.run(instance);
    });
}

loadWebAssembly();

app.use(express.static('public'));
app.get('/', (req, res) => {
  const sum = addTwoNumbers();
  res.send(`Sum: ${sum}`);
});

app.use('/users', usersRouter);

// catch 404 and forward to error handler
app.use(function(req, res, next) {
  next(createError(404));
});

// error handler
app.use(function(err, req, res, next) {
  // set locals, only providing error in development
  res.locals.message = err.message;
  res.locals.error = req.app.get('env') === 'development' ? err : {};

  // render the error page
  res.status(err.status || 500);
  res.render('error');
});

module.exports = app;
