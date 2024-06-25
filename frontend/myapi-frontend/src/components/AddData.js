import React, { useState } from 'react';
import axios from 'axios';

const AddData = () => {
  const [name, setName] = useState('');

  const handleSubmit = async (e) => {
    e.preventDefault();
    try {
      await axios.post('/api/data', { name }, {
        headers: { Authorization: `Bearer ${localStorage.getItem('token')}` }
      });
      setName('');
    } catch (error) {
      console.error('There was an error adding the data!', error);
    }
  };

  return (
    <div>
      <h1>Add Data</h1>
      <form onSubmit={handleSubmit}>
        <input
          type="text"
          placeholder="Name"
          value={name}
          onChange={(e) => setName(e.target.value)}
          required
        />
        <button type="submit">Add</button>
      </form>
    </div>
  );
};

export default AddData;
