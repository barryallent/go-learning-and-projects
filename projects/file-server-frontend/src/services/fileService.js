const API_BASE_URL = 'http://localhost:9095';

// File service for handling API calls
export const fileService = {
  // Upload a file
  uploadFile: async (id, filename, file) => {
    try {
      const response = await fetch(`${API_BASE_URL}/images/${id}/${filename}`, {
        method: 'POST',
        body: file,
        headers: {
          'Content-Type': 'application/octet-stream',
        },
      });
      
      if (!response.ok) {
        throw new Error(`Upload failed: ${response.statusText}`);
      }
      
      return { success: true, message: 'File uploaded successfully' };
    } catch (error) {
      console.error('Upload error:', error);
      throw error;
    }
  },

  // Get all files from the backend
  getAllFiles: async () => {
    try {
      const response = await fetch(`${API_BASE_URL}/files`);
      if (!response.ok) {
        throw new Error(`Failed to fetch files: ${response.statusText}`);
      }
      
      const files = await response.json();
      return files || [];
    } catch (error) {
      console.error('Get files error:', error);
      throw error;
    }
  },

  // Delete a file
  deleteFile: async (id, filename) => {
    try {
      const response = await fetch(`${API_BASE_URL}/images/${id}/${filename}`, {
        method: 'DELETE',
      });
      
      if (!response.ok) {
        throw new Error(`Delete failed: ${response.statusText}`);
      }
      
      return { success: true, message: 'File deleted successfully' };
    } catch (error) {
      console.error('Delete error:', error);
      throw error;
    }
  },

  // Get file URL for download
  getFileUrl: (id, filename) => {
    return `${API_BASE_URL}/images/${id}/${filename}`;
  },

  // Health check
  healthCheck: async () => {
    try {
      const response = await fetch(`${API_BASE_URL}/health`);
      if (response.ok) {
        return await response.json();
      }
      throw new Error('Health check failed');
    } catch (error) {
      console.error('Health check error:', error);
      throw error;
    }
  },
}; 