import React, { useState, useEffect } from 'react';
import { fileService } from '../services/fileService';

const ServerStatus = () => {
  const [serverStatus, setServerStatus] = useState('checking');

  useEffect(() => {
    checkServerHealth();
  }, []);

  const checkServerHealth = async () => {
    try {
      await fileService.healthCheck();
      setServerStatus('connected');
    } catch (error) {
      setServerStatus('disconnected');
    }
  };

  const getStatusColor = () => {
    switch (serverStatus) {
      case 'connected': return '#4caf50';
      case 'disconnected': return '#f44336';
      default: return '#ff9800';
    }
  };

  return (
    <div className="server-status">
      <span 
        className="status-indicator"
        style={{ backgroundColor: getStatusColor() }}
      ></span>
      <span>Server: {serverStatus}</span>
      <button onClick={checkServerHealth} className="refresh-btn">
        Refresh
      </button>
    </div>
  );
};

export default ServerStatus; 