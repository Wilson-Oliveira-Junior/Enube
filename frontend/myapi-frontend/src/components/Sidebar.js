import React from 'react';
import { Link, useRouteMatch } from 'react-router-dom';

const Sidebar = () => {
  let { url } = useRouteMatch();

  return (
    <div className="sidebar">
      <ul>
        <li>
          <Link to={`${url}/`}>Home</Link>
        </li>
        <li>
          <Link to={`${url}/add`}>Add Data</Link>
        </li>
        <li>
          <Link to={`${url}/edit`}>Edit Data</Link>
        </li>
      </ul>
    </div>
  );
};

export default Sidebar;
