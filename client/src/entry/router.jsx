//路由
import React from 'react';
//import { render } from 'react-dom'
import { Router, Route, Link } from 'react-router';
import app from '../view/app.jsx';
import index from '../view/index.jsx';
import login from '../view/login.jsx';



const routers = (
    <Router>
        <Route path="/" component={app}>
            <Route path="login" component={login} />
            <Route path="index" component={index} />
        </Route>
    </Router>
)

export default routers