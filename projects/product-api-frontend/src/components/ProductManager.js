import React, { useState, useEffect } from 'react';
import ProductForm from './ProductForm';
import ProductList from './ProductList';
import ErrorBanner from './ErrorBanner';
import LoadingSpinner from './LoadingSpinner';
import { productService } from '../services/productService';

const ProductManager = () => {
  const [products, setProducts] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);
  const [editingProduct, setEditingProduct] = useState(null);

  // Fetch products from API
  const fetchProducts = async () => {
    try {
      setLoading(true);
      const data = await productService.getProducts();
      setProducts(data || []);
      setError(null);
    } catch (err) {
      setError(`Failed to fetch products: ${err.message}`);
      console.error('Error fetching products:', err);
    } finally {
      setLoading(false);
    }
  };

  // Add new product
  const handleAddProduct = async (productData) => {
    try {
      const newProduct = await productService.createProduct(productData);
      setProducts(prev => [...prev, newProduct]);
      setError(null);
    } catch (err) {
      setError(`Failed to add product: ${err.message}`);
      throw err;
    }
  };

  // Update existing product
  const handleUpdateProduct = async (productData) => {
    try {
      const updatedProduct = await productService.updateProduct(productData);
      setProducts(prev => prev.map(p => p.id === updatedProduct.id ? updatedProduct : p));
      setEditingProduct(null);
      setError(null);
    } catch (err) {
      setError(`Failed to update product: ${err.message}`);
      throw err;
    }
  };

  // Handle edit action
  const handleEdit = (product) => {
    setEditingProduct(product);
  };

  // Handle cancel edit
  const handleCancelEdit = () => {
    setEditingProduct(null);
  };

  // Load products on component mount
  useEffect(() => {
    fetchProducts();
  }, []);

  if (loading && products.length === 0) {
    return (
      <div className="container">
        <h1>Product Management</h1>
        <LoadingSpinner />
      </div>
    );
  }

  return (
    <div className="container">
      <header className="app-header">
        <h1>🛍️ Product Management</h1>
        <p>Manage your products with ease</p>
      </header>

      {error && (
        <ErrorBanner 
          message={error} 
          onRetry={fetchProducts}
        />
      )}

      <div className="app-content">
        <div className="form-section">
          <ProductForm
            product={editingProduct}
            onSave={editingProduct ? handleUpdateProduct : handleAddProduct}
            onCancel={handleCancelEdit}
            isEditing={!!editingProduct}
          />
        </div>

        <div className="products-section">
          <ProductList
            products={products}
            onEdit={handleEdit}
            onRefresh={fetchProducts}
          />
        </div>
      </div>
    </div>
  );
};

export default ProductManager; 