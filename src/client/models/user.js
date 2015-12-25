import mongoose from 'mongoose'
let Schema = mongoose.Schema;

// Define User schema
let _User = new Schema({
    email : String,
    name : String,
    salt : String,
    password : String
});
// export them
exports.User = mongoose.model('User', _User);
