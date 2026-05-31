import React from 'react';
import { Link, useLocation } from 'react-router-dom';

interface TopAppBarProps {
  isAdmin?: boolean;
}

// DESIGN.md 기반 디자인 토큰 적용
export const TopAppBar: React.FC<TopAppBarProps> = ({ isAdmin = false }) => {
  const location = useLocation();

  const navItems = isAdmin 
    ? [
        { label: 'Dashboard', path: '/admin' },
        { label: 'Inventory', path: '/admin/inventory' },
      ]
    : [
        { label: 'Dashboard', path: '/' },
        { label: 'Catalog', path: '/catalog' },
        { label: 'Settings', path: '/settings' },
      ];

  return (
    <header className="bg-surface w-full h-16 flex items-center sticky top-0 z-40 justify-between px-md md:px-lg border-b border-outline-variant">
      <div className="flex items-center gap-4">
        <h1 className="font-headline-md text-primary">KH Library</h1>
      </div>
      
      <div className="hidden md:flex items-center gap-lg h-full">
        <nav className="flex gap-md h-full">
          {navItems.map((item) => (
            <Link 
              key={item.label}
              to={item.path}
              className={`h-full flex items-center px-sm transition-colors duration-150 font-label-md ${
                location.pathname === item.path 
                  ? 'text-primary border-b-2 border-primary' 
                  : 'text-on-surface-variant hover:bg-surface-container-low'
              }`}
            >
              {item.label}
            </Link>
          ))}
        </nav>
        
        <div className="flex items-center gap-sm ml-lg pl-lg border-l border-outline-variant h-8">
          <span className="material-symbols-outlined text-on-surface-variant cursor-pointer">qr_code_scanner</span>
          <span className="material-symbols-outlined text-on-surface-variant cursor-pointer">account_circle</span>
        </div>
      </div>
    </header>
  );
};
