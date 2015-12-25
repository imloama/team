import Button from 'react-bootstrap/lib/Button';
import React, {Component} from 'react';

export default class App extends Component {
    displayName = 'App component'
    render() {
        const {history} = this.props;
        return (
            <div>
              <Button bsStyle="success">success</Button>
            </div>
        );
    }
}
