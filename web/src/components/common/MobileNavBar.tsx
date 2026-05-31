import React from 'react';
import { NavLink } from 'react-router-dom';
import { useAuthStore } from '../../store/authStore';

export const MobileNavBar: React.FC = () => {
  const role = useAuthStore(state => state.role);

  const navItems = role === 'admin'
    ? [
        { label: 'Admin', icon: 'dashboard', path: '/admin' },
        { label: 'Catalog', icon: 'inventory_2', path: '/catalog' },
        { label: 'History', icon: 'history', path: '/history' },
      ]
    : [
        { label: 'Public', icon: 'home', path: '/catalog' },
        { label: 'My Library', icon: 'inventory_2', path: '/my-library' },
        { label: 'History', icon: 'history', path: '/history' },
        { label: 'Profile', icon: 'person', path: '/profile' },
      ];

  return (
    <nav className="flex justify-around items-center px-sm py-xs h-16 bg-surface-container-lowest">
      {navItems.map(item => (
        <NavLink
          key={item.label}
          to={item.path}
          className={({ isActive }) =>
            `flex flex-col items-center justify-center py-xs px-md rounded-default transition-all duration-150 ${
              isActive 
                ? 'bg-secondary-container text-on-secondary-container' 
                : 'text-on-surface-variant'
            }`
          }
        >
          <span className="material-symbols-outlined" style={{ fontVariationSettings: "'FILL' 1" }}>{item.icon}</span>
          <span className="font-label-md">{item.label}</span>
        </NavLink>
      ))}
    </nav>
  );
};
