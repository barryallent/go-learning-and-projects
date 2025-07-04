import React from 'react';
import ProductCard from './ProductCard';

const ProductList = ({ products, onEdit, onRefresh }) => {
  return (
    <div className="products-section">
      <div className="products-header">
        <h2>Products ({products.length})</h2>
        <button onClick={onRefresh} className="btn btn-secondary btn-small">
          Refresh
        </button>
      </div>

      {products.length === 0 ? (
        <div className="empty-state">
          <p>No products found. Add your first product above!</p>
        </div>
      ) : (
        <div className="products-grid">
          {products.map(product => (
            <ProductCard
              key={product.id}
              product={product}
              onEdit={onEdit}
            />
          ))}
        </div>
      )}
    </div>
  );
};

export default ProductList; 