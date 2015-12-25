//项目实际运行时将执行该界面
var path = require('path');
var express = require('express');
var app = express();
var conf = require("./conf");//配置条件
var mongodb = require('./mongodb');

app.use('/bower_components', express.static("bower_components"));
app.use('/static', express.static("dist"));

app.get('*', function(req, res) {
  res.sendFile(path.join(__dirname, 'index.html'));
});

app.listen(conf.port, conf.host, function(err) {
  if (err) {
    console.log(err);
    return;
  }

  console.log('Listening at http://'+conf.host+':'+conf.port);
});
