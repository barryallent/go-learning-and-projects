export const validateProduct = (formData) => {
  const errors = {};
  
  // Validate name
  if (!formData.name.trim()) {
    errors.name = 'Product name is required';
  }
  
  // Validate price
  if (!formData.price || formData.price <= 0) {
    errors.price = 'Price must be greater than 0';
  }
  
  // Validate SKU
  if (!formData.sku.trim()) {
    errors.sku = 'SKU is required';
  } else if (!/^SKU-[0-9]+$/.test(formData.sku)) {
    errors.sku = 'SKU must be in format SKU-123 (e.g., SKU-001)';
  }
  
  return errors;
}; 