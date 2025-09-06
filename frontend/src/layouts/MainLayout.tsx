import React from 'react';
import Header from '../components/Header';
import Footer from '../components/Footer';


interface MainLayoutProps {
  children: React.ReactNode;
  cartLength: number;
  onShowCart: () => void;
  showCart: boolean;
  setShowCart: (show: boolean) => void;
}

const MainLayout: React.FC<MainLayoutProps> = ({ children, cartLength, onShowCart, showCart, setShowCart }) => {
  return (
    <div className="min-h-screen bg-gradient-to-br from-orange-50 to-red-50">
      <Header cartLength={cartLength} onShowCart={onShowCart} showCart={showCart} setShowCart={setShowCart} />
      <main>{children}</main>
      <Footer />
    </div>
  );
};

export default MainLayout;
