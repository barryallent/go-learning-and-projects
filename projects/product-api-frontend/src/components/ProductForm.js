import React, { useState, useEffect } from 'react';
import { validateProduct } from '../utils/validation';

const ProductForm = ({ product, onSave, onCancel, isEditing = false }) => {
  const [formData, setFormData] = useState({
    name: '',
    description: '',
    price: '',
    sku: ''
  });
  const [errors, setErrors] = useState({});
  const [loading, setLoading] = useState(false);

  useEffect(() => {
    if (product) {
      setFormData({
        name: product.name || '',
        description: product.description || '',
        price: product.price || '',
        sku: product.sku || ''
      });
    } else {
      setFormData({
        name: '',
        description: '',
        price: '',
        sku: ''
      });
    }
  }, [product]);

  const handleSubmit = async (e) => {
    e.preventDefault();
    
    const validationErrors = validateProduct(formData);
    if (Object.keys(validationErrors).length > 0) {
      setErrors(validationErrors);
      return;
    }
    
    setLoading(true);
    
    try {
      const productData = {
        ...formData,
        price: parseFloat(formData.price)
      };
      
      if (isEditing) {
        productData.id = product.id;
      }
      
      await onSave(productData);
      
      // Reset form if not editing
      if (!isEditing) {
        setFormData({ name: '', description: '', price: '', sku: '' });
      }
      setErrors({});
    } catch (error) {
      console.error('Error saving product:', error);
    } finally {
      setLoading(false);
    }
  };

  const handleChange = (e) => {
    const { name, value } = e.target;
    setFormData(prev => ({ ...prev, [name]: value }));
    
    // Clear error when user starts typing
    if (errors[name]) {
      setErrors(prev => ({ ...prev, [name]: '' }));
    }
  };

  return (
    <div className="product-form">
      <h3>{isEditing ? 'Edit Product' : 'Add New Product'}</h3>
      <form onSubmit={handleSubmit}>
        <div className="form-group">
          <label htmlFor="name">Product Name *</label>
          <input
            type="text"
            id="name"
            name="name"
            value={formData.name}
            onChange={handleChange}
            className={errors.name ? 'error' : ''}
            placeholder="Enter product name"
          />
          {errors.name && <span className="error-message">{errors.name}</span>}
        </div>

        <div className="form-group">
          <label htmlFor="description">Description</label>
          <textarea
            id="description"
            name="description"
            value={formData.description}
            onChange={handleChange}
            placeholder="Enter product description"
            rows="3"
          />
        </div>

        <div className="form-group">
          <label htmlFor="price">Price *</label>
          <input
            type="number"
            id="price"
            name="price"
            value={formData.price}
            onChange={handleChange}
            step="0.01"
            min="0"
            className={errors.price ? 'error' : ''}
            placeholder="0.00"
          />
          {errors.price && <span className="error-message">{errors.price}</span>}
        </div>

        <div className="form-group">
          <label htmlFor="sku">SKU *</label>
          <input
            type="text"
            id="sku"
            name="sku"
            value={formData.sku}
            onChange={handleChange}
            className={errors.sku ? 'error' : ''}
            placeholder="SKU-001"
          />
          {errors.sku && <span className="error-message">{errors.sku}</span>}
        </div>

        <div className="form-actions">
          <button type="submit" disabled={loading} className="btn btn-primary">
            {loading ? 'Saving...' : (isEditing ? 'Update Product' : 'Add Product')}
          </button>
          {isEditing && (
            <button type="button" onClick={onCancel} className="btn btn-secondary">
              Cancel
            </button>
          )}
        </div>
      </form>
    </div>
  );
};

export default ProductForm; 