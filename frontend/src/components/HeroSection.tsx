import React from 'react';

const HeroSection: React.FC = () => {
  return (
    <section className="relative py-20 bg-gradient-to-r from-red-900 to-orange-800 text-white">
      <div className="container mx-auto px-4 text-center">
        <div className="max-w-4xl mx-auto">
          <h2 className="text-5xl font-bold mb-6">60 Years of Tea Excellence</h2>
          <p className="text-xl mb-8 text-orange-100">
            From our family to yours, experience the finest traditional teas and signature blends 
            crafted with love in the heart of Jaipur.
          </p>
          <div className="grid md:grid-cols-3 gap-8 mt-12">
            <div className="text-center">
              <div className="text-6xl mb-4">🍃</div>
              <h3 className="text-xl font-semibold mb-2">Premium Quality</h3>
              <p className="text-orange-200">Hand-selected tea leaves from the finest gardens</p>
            </div>
            <div className="text-center">
              <div className="text-6xl mb-4">👑</div>
              <h3 className="text-xl font-semibold mb-2">Royal Heritage</h3>
              <p className="text-orange-200">Traditional recipes passed down through generations</p>
            </div>
            <div className="text-center">
              <div className="text-6xl mb-4">🚚</div>
              <h3 className="text-xl font-semibold mb-2">Fresh Delivery</h3>
              <p className="text-orange-200">Direct from our store to your doorstep</p>
            </div>
          </div>
        </div>
      </div>
    </section>
  );
};

export default HeroSection;
