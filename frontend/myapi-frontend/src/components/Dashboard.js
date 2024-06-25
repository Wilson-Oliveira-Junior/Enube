import React from 'react';
import { BrowserRouter as Router, Route, Switch, Link, useRouteMatch } from 'react-router-dom';
import Home from './Home';
import AddData from './AddData';
import EditData from './EditData';
import Sidebar from './Sidebar';

const Dashboard = () => {
  let { path, url } = useRouteMatch();

  return (
    <div className="dashboard">
      <Router>
        <Sidebar />
        <div className="content">
          <Switch>
            <Route exact path={`${path}/`} component={Home} />
            <Route path={`${path}/add`} component={AddData} />
            <Route path={`${path}/edit`} component={EditData} />
          </Switch>
        </div>
      </Router>
    </div>
  );
};

export default Dashboard;
