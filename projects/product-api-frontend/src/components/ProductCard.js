import React from 'react';

const ProductCard = ({ product, onEdit }) => {
  return (
    <div className="product-card">
      <div className="product-header">
        <h4>{product.name}</h4>
        <span className="product-sku">{product.sku}</span>
      </div>
      {product.description && (
        <p className="product-description">{product.description}</p>
      )}
      <div className="product-footer">
        <span className="product-price">${product.price.toFixed(2)}</span>
        <button 
          onClick={() => onEdit(product)} 
          className="btn btn-small btn-secondary"
        >
          Edit
        </button>
      </div>
    </div>
  );
};

export default ProductCard; 