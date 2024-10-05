import React from 'react';
import { BrowserRouter as Router, Route, Switch } from 'react-router-dom';
import ClientList from './components/ClientList';
import AddClient from './components/AddClient';
import Navbar from './components/Navbar';

function App() {
  return (
    <Router>
      <div className="App">
        <Navbar />
        <Switch>
          <Route exact path="/" component={ClientList} />
          <Route path="/add" component={AddClient} />
        </Switch>
      </div>
    </Router>
  );
}

export default App;