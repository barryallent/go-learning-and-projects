import React, { useState, useEffect } from 'react';
import FileUpload from './FileUpload';
import FileList from './FileList';
import { fileService } from '../services/fileService';

const FileManager = () => {
  const [files, setFiles] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState('');

  // Load files from backend on component mount
  useEffect(() => {
    loadFiles();
  }, []);

  const loadFiles = async () => {
    try {
      setLoading(true);
      setError('');
      const filesFromBackend = await fileService.getAllFiles();
      setFiles(filesFromBackend);
    } catch (error) {
      console.error('Failed to load files:', error);
      setError('Failed to load files from server');
    } finally {
      setLoading(false);
    }
  };

  const handleUploadSuccess = async (fileInfo) => {
    // Refresh the file list from backend to get updated data
    await loadFiles();
  };

  const handleDeleteFile = async (file) => {
    try {
      await fileService.deleteFile(file.id, file.filename);
      // Refresh the file list after successful deletion
      await loadFiles();
    } catch (error) {
      console.error('Failed to delete file:', error);
      setError('Failed to delete file');
    }
  };

  const handleRefresh = () => {
    loadFiles();
  };

  return (
    <div className="container">
      <FileUpload onUploadSuccess={handleUploadSuccess} />
      <FileList 
        files={files} 
        loading={loading}
        error={error}
        onDeleteFile={handleDeleteFile}
        onRefresh={handleRefresh}
      />
    </div>
  );
};

export default FileManager; 