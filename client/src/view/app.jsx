import React from 'react';
// import Router from 'react-router'

// let RouteHandler  = Router.RouteHandler;

const App = React.createClass({
  render() {
    return (
      <div>
        <h1>app</h1>
        {this.props.children}
      </div>
    );
  },
});

export default App;