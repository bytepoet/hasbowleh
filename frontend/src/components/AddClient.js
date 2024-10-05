import React, { useState } from 'react';
import { addClient } from '../services/api';

function AddClient() {
  const [username, setUsername] = useState('');

  const handleSubmit = async (e) => {
    e.preventDefault();
    await addClient(username);
    setUsername('');
    // Redirect to client list or show success message
  };

  return (
    <form onSubmit={handleSubmit}>
      <input
        type="text"
        value={username}
        onChange={(e) => setUsername(e.target.value)}
        placeholder="Enter username"
        required
      />
      <button type="submit">Add Client</button>
    </form>
  );
}

export default AddClient;