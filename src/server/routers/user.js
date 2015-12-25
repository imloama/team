
var mongoose = require('mongoose');
var User = require('../models/user')


module.exports = function(req,res){
    console.log(" on /user");
    User.create({
        email : 'nowind_lee@qq.com',
        name : 'Freewind'
    }).then(function (user) {
      console.log(user);
          // var u = User.findById(user._id)
            //.populate('creater')
            //.populate('updater')
            // .exec();
            res.send(user);
    });
}
