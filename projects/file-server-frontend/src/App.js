import React from 'react';
import FileManager from './components/FileManager';
import ServerStatus from './components/ServerStatus';
import './App.css';

function App() {
  return (
    <div className="App">
      <header className="App-header">
        <h1>File Server</h1>
        <ServerStatus />
      </header>

      <main className="App-main">
        <FileManager />
      </main>

      <footer className="App-footer">
        <p>Simple file upload and download service</p>
      </footer>
    </div>
  );
}

export default App;
