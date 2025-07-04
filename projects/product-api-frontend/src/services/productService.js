// API base URL - adjust if your backend runs on a different port
const API_BASE_URL = 'http://localhost:9080';

export const productService = {
  // Get all products
  async getProducts() {
    const response = await fetch(`${API_BASE_URL}/`);
    
    if (!response.ok) {
      throw new Error(`HTTP error! status: ${response.status}`);
    }
    
    return await response.json();
  },

  // Create a new product
  async createProduct(productData) {
    const response = await fetch(`${API_BASE_URL}/product`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(productData),
    });

    if (!response.ok) {
      const errorText = await response.text();
      throw new Error(`HTTP error! status: ${response.status}, message: ${errorText}`);
    }

    return await response.json();
  },

  // Update an existing product
  async updateProduct(productData) {
    const response = await fetch(`${API_BASE_URL}/product/${productData.id}`, {
      method: 'PUT',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(productData),
    });

    if (!response.ok) {
      const errorText = await response.text();
      throw new Error(`HTTP error! status: ${response.status}, message: ${errorText}`);
    }

    return await response.json();
  },
}; 