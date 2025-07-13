import React, { useState } from 'react';
import { fileService } from '../services/fileService';

const FileUpload = ({ onUploadSuccess }) => {
  const [selectedFile, setSelectedFile] = useState(null);
  const [fileId, setFileId] = useState('1');
  const [isUploading, setIsUploading] = useState(false);
  const [error, setError] = useState('');

  const handleFileSelect = (event) => {
    const file = event.target.files[0];
    setSelectedFile(file);
    setError('');
  };

  const handleUpload = async (event) => {
    event.preventDefault();
    
    if (!selectedFile) {
      setError('Please select a file');
      return;
    }

    if (!fileId.trim()) {
      setError('Please enter a file ID');
      return;
    }

    setIsUploading(true);
    setError('');

    try {
      await fileService.uploadFile(fileId, selectedFile.name, selectedFile);
      
      // Reset form
      setSelectedFile(null);
      setFileId('1');
      event.target.reset();
      
      // Notify parent component
      if (onUploadSuccess) {
        onUploadSuccess({
          id: fileId,
          filename: selectedFile.name,
          size: selectedFile.size,
        });
      }
      
    } catch (error) {
      setError(error.message || 'Upload failed');
    } finally {
      setIsUploading(false);
    }
  };

  return (
    <div className="file-upload">
      <h2>Upload File</h2>
      <form onSubmit={handleUpload}>
        <div className="form-group">
          <label htmlFor="fileId">File ID:</label>
          <input
            type="text"
            id="fileId"
            value={fileId}
            onChange={(e) => setFileId(e.target.value)}
            placeholder="Enter file ID (e.g., 1)"
            disabled={isUploading}
          />
        </div>
        
        <div className="form-group">
          <label htmlFor="file">Select File:</label>
          <input
            type="file"
            id="file"
            onChange={handleFileSelect}
            disabled={isUploading}
          />
        </div>
        
        {selectedFile && (
          <div className="file-info">
            <p>Selected: {selectedFile.name} ({Math.round(selectedFile.size / 1024)} KB)</p>
          </div>
        )}
        
        {error && <div className="error">{error}</div>}
        
        <button type="submit" disabled={isUploading}>
          {isUploading ? 'Uploading...' : 'Upload File'}
        </button>
      </form>
    </div>
  );
};

export default FileUpload; 