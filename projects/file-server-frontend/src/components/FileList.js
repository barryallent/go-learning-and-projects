import React from 'react';
import { fileService } from '../services/fileService';

const FileList = ({ files, loading, error, onDeleteFile, onRefresh }) => {
  const handleDownload = (file) => {
    const url = fileService.getFileUrl(file.id, file.filename);
    window.open(url, '_blank');
  };

  const handleDelete = (file) => {
    if (window.confirm(`Are you sure you want to delete "${file.filename}"?`)) {
      onDeleteFile(file);
    }
  };

  const formatFileSize = (bytes) => {
    if (bytes === 0) return '0 Bytes';
    const k = 1024;
    const sizes = ['Bytes', 'KB', 'MB', 'GB'];
    const i = Math.floor(Math.log(bytes) / Math.log(k));
    return parseFloat((bytes / Math.pow(k, i)).toFixed(2)) + ' ' + sizes[i];
  };

  return (
    <div className="file-list">
      <div className="file-list-header">
        <h2>Files on Server</h2>
        <button onClick={onRefresh} className="refresh-btn" disabled={loading}>
          {loading ? 'Loading...' : 'Refresh'}
        </button>
      </div>

      {error && (
        <div className="error-message">
          {error}
          <button onClick={onRefresh} className="retry-btn">
            Retry
          </button>
        </div>
      )}

      {loading ? (
        <div className="loading">Loading files...</div>
      ) : files.length === 0 ? (
        <p className="no-files">No files found on server.</p>
      ) : (
        <div className="files-grid">
          {files.map((file, index) => (
            <div key={`${file.id}-${file.filename}-${index}`} className="file-item">
              <div className="file-info">
                <h3>{file.filename}</h3>
                <p>ID: {file.id}</p>
                <p>Size: {formatFileSize(file.size)}</p>
                <p>Path: {file.path}</p>
              </div>
              <div className="file-actions">
                <button 
                  onClick={() => handleDownload(file)}
                  className="download-btn"
                >
                  Download
                </button>
                <button 
                  onClick={() => handleDelete(file)}
                  className="delete-btn"
                >
                  Delete
                </button>
              </div>
            </div>
          ))}
        </div>
      )}
    </div>
  );
};

export default FileList; 