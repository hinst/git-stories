import React, { ChangeEvent, Component } from 'react';
import './w3.css';
import './git-stories.css'
import { RepoHistoryViewer } from './RepoHistoryViewer';

class App extends Component {
    render() {
        document.title = 'Git Stories';
        return <RepoHistoryViewer/>;
    }
}

export default App;
