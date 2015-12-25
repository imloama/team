//后台入口类，支持express,可以打包
var express = require('express'),
    bodyParser = require('body-parser'),
    methodOverride = require('method-override'),
    cookieParser = require('cookie-parser'),
    session = require('express-session'),
    path = require('path'),
    app = express();

var userRouter = require('./routers/user');


app.use(bodyParser.json());
app.use(bodyParser.urlencoded({extended: false}));
app.use(methodOverride());
app.use(cookieParser());
app.use(session({
    //store: new RedisStore(conf.redis),
    secret: 'team',
    resave: true,
    saveUninitialized: false
}));


app.get('/user',userRouter);


module.exports = app;
