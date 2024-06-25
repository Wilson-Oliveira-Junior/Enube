import React, { useEffect, useState } from 'react';
import axios from 'axios';

const Home = () => {
  const [data, setData] = useState([]);

  useEffect(() => {
    axios.get('/api/data', {
      headers: { Authorization: `Bearer ${localStorage.getItem('token')}` }
    })
    .then(response => {
      setData(response.data);
    })
    .catch(error => {
      console.error('There was an error fetching the data!', error);
    });
  }, []);

  return (
    <div>
      <h1>Home</h1>
      <ul>
        {data.map(item => (
          <li key={item.id}>{item.name}</li>
        ))}
      </ul>
    </div>
  );
};

export default Home;
