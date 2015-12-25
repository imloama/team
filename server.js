var path = require('path');
var express = require('express');
var webpack = require('webpack');
var config = require('./webpack.config');
var mongodb = require('./mongodb');
var app = require('./src/server/index');
var compiler = webpack(config);

app.use('/bower_components', express.static("bower_components"));

app.use(require('webpack-dev-middleware')(compiler, {
  noInfo: true,
  publicPath: config.output.publicPath
}));

app.use(require('webpack-hot-middleware')(compiler));

app.get('/', function(req, res) {
  res.sendFile(path.join(__dirname, 'index.html'));
});

app.listen(3000, 'localhost', function(err) {
  if (err) {
    console.log(err);
    return;
  }

  console.log('Listening at http://localhost:3000');
});
