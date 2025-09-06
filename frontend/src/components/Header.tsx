import React from 'react';

interface HeaderProps {
  cartLength: number;
  onShowCart: () => void;
  showCart: boolean;
  setShowCart: (show: boolean) => void;
}

const Header: React.FC<HeaderProps> = ({ cartLength, onShowCart }) => {
  return (
    <header className="bg-gradient-to-r from-red-800 to-orange-700 text-white shadow-xl">
      <div className="container mx-auto px-4 py-6">
        <div className="flex justify-between items-center">
          <div>
            <h1 className="text-4xl font-bold mb-2">Mangal Chai</h1>
            <p className="text-orange-100 text-lg">Serving Premium Teas Since 1965 • Jaipur, India</p>
          </div>
          <button
            onClick={onShowCart}
            className="bg-orange-600 hover:bg-orange-500 px-6 py-3 rounded-full font-semibold transition-all duration-300 flex items-center gap-2 shadow-lg"
          >
            <span>🛒</span>
            Cart ({cartLength})
          </button>
        </div>
      </div>
    </header>
  );
};

export default Header;