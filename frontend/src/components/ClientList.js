import React, { useState, useEffect } from 'react';
import { getClients, deleteClient, downloadConfig } from '../services/api';

function ClientList() {
  const [clients, setClients] = useState([]);
  const [searchTerm, setSearchTerm] = useState('');

  useEffect(() => {
    fetchClients();
  }, []);

  const fetchClients = async () => {
    const fetchedClients = await getClients();
    setClients(fetchedClients);
  };

  const handleDelete = async (clientId) => {
    if (window.confirm('Are you sure you want to delete this client?')) {
      await deleteClient(clientId);
      fetchClients();
    }
  };

  const handleDownload = async (clientId) => {
    const config = await downloadConfig(clientId);
    // Logic to download the config file
  };

  const filteredClients = clients.filter(client =>
    client.username.toLowerCase().includes(searchTerm.toLowerCase()) ||
    client.ip.includes(searchTerm)
  );

  return (
    <div>
      <input
        type="text"
        placeholder="Search clients..."
        value={searchTerm}
        onChange={(e) => setSearchTerm(e.target.value)}
      />
      <table>
        <thead>
          <tr>
            <th>Username</th>
            <th>IP</th>
            <th>Actions</th>
          </tr>
        </thead>
        <tbody>
          {filteredClients.map(client => (
            <tr key={client.id}>
              <td>{client.username}</td>
              <td>{client.ip}</td>
              <td>
                <button onClick={() => handleDownload(client.id)}>Download Config</button>
                <button onClick={() => handleDelete(client.id)}>Delete</button>
              </td>
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  );
}

export default ClientList;