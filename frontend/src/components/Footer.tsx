import React from 'react';

const Footer: React.FC = () => {
  return (
    <footer className="bg-gray-800 text-white py-12">
      <div className="container mx-auto px-4 text-center">
        <h3 className="text-2xl font-bold mb-4">Mangal Chai</h3>
        <p className="text-gray-300 mb-4">
          Experience the finest teas from our 60-year-old family business in Jaipur, India
        </p>
        <p className="text-gray-400 text-sm">
          © 2025 Mangal Chai. All rights reserved. | Serving quality since 1965
        </p>
      </div>
    </footer>
  );
};

export default Footer;
