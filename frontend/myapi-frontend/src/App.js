import React, { useState } from 'react';
import { BrowserRouter as Router, Route, Switch, Redirect } from 'react-router-dom';
import Login from './components/Login';
import Dashboard from './components/Dashboard';
import './App.css';


const App = () => {
  const [token, setToken] = useState(localStorage.getItem('token'));

  return (
    <Router>
      <Switch>
        <Route path="/login">
          <Login setToken={setToken} />
        </Route>
        <Route path="/dashboard">
          {token ? <Dashboard /> : <Redirect to="/login" />}
        </Route>
        <Route path="/">
          <Redirect to="/login" />
        </Route>
      </Switch>
    </Router>
  );
};

export default App;
