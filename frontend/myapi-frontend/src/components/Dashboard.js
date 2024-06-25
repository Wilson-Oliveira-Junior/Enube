import React, { useEffect, useState } from 'react';
import axios from 'axios';

const Dashboard = ({ token }) => {
  const [data, setData] = useState([]);
  const [error, setError] = useState(null);

  useEffect(() => {
    const fetchData = async () => {
      try {
        const response = await axios.get('/api/data', {
          headers: { Authorization: `Bearer ${token}` }
        });
        setData(response.data);
      } catch (error) {
        setError('Error fetching data');
        console.error('Error fetching data:', error);
      }
    };
  
    fetchData();
  }, [token]);
  

  if (error) {
    return <div>Error: {error}</div>;
  }

  return (
    <div>
      <h1>Dashboard</h1>
      <ul>
        {data.map(item => (
          <li key={item.PartnerId}>
            PartnerId: {item.PartnerId}, PartnerName: {item.PartnerName}, CustomerId: {item.CustomerId}, CustomerName: {item.CustomerName}
          </li>
        ))}
      </ul>
    </div>
  );
};

export default Dashboard;
