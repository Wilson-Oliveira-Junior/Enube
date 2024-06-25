import React, { useState, useEffect } from 'react';
import axios from 'axios';

const EditData = () => {
  const [data, setData] = useState([]);
  const [selected, setSelected] = useState(null);
  const [name, setName] = useState('');

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

  const handleEdit = async (e) => {
    e.preventDefault();
    try {
      await axios.put(`/api/data/${selected.id}`, { name }, {
        headers: { Authorization: `Bearer ${localStorage.getItem('token')}` }
      });
      setName('');
      setSelected(null);
    } catch (error) {
      console.error('There was an error editing the data!', error);
    }
  };

  return (
    <div>
      <h1>Edit Data</h1>
      <ul>
        {data.map(item => (
          <li key={item.id} onClick={() => setSelected(item)}>
            {item.name}
          </li>
        ))}
      </ul>
      {selected && (
        <form onSubmit={handleEdit}>
          <input
            type="text"
            placeholder="Name"
            value={name}
            onChange={(e) => setName(e.target.value)}
            required
          />
          <button type="submit">Edit</button>
        </form>
      )}
    </div>
  );
};

export default EditData;
