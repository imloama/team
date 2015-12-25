//mongodb配置
var mongoose = require('mongoose');
var url = 'mongodb://10.10.4.73:27017/team';
mongoose.connect(url);


function installDB(){
  //TODO 数据库初始数据
}


module.exports.mongoose = mongoose;
module.exports.install = installDB;
