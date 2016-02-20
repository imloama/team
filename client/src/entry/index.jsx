import '../common/lib';
// import App from '../component/App';
import {render} from 'react-dom';
import React from 'react';
import Routers from './router.jsx';
import Router from 'react-router';
// ReactDOM.render(<App />, document.getElementById('react-content'));

render(Routers, document.getElementById('app'));
