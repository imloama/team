import React from 'react'
import { render } from 'react-dom'
//import { Provider } from 'react-redux'
import App from './views/App'


// function renderDevTools(store) {
//     if (__DEBUG__) {
//         let {DevTools, DebugPanel, LogMonitor} = require('redux-devtools/lib/react')
//
//         return (
//             <DebugPanel top right bottom>
//             <DevTools store={store} monitor={LogMonitor} />
//             </DebugPanel>
//         )
//     }
//
//     return null
// }

render(
        <App />,
    document.getElementById('app'))
