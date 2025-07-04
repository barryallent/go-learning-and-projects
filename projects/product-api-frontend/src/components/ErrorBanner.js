import React from 'react';

const ErrorBanner = ({ message, onRetry }) => {
  return (
    <div className="error-banner">
      <p>⚠️ {message}</p>
      <button onClick={onRetry} className="btn btn-small">
        Retry
      </button>
    </div>
  );
};

export default ErrorBanner; 