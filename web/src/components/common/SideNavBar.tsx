import React, { useState } from 'react';
import { NavLink } from 'react-router-dom';
import logo from '/logo.png';
import { useAuthStore } from '../../store/authStore';

export const SideNavBar: React.FC = () => {
  const role = useAuthStore(state => state.role);
  const [isAdminOpen, setIsAdminOpen] = useState(false);
  
  const navItems = role === 'admin' 
    ? [
        { label: 'Dashboard', icon: 'dashboard', path: '/admin' },
        { 
          label: 'Admin', 
          icon: 'admin_panel_settings', 
          isSubmenu: true,
          subItems: [
            { label: 'Users', path: '/admin/users' },
            { label: 'Announcements', path: '/admin/announcements' },
            { label: 'Inventory', path: '/admin/inventory' },
            { label: 'System', path: '/admin/system' },
          ]
        },
        { label: 'History', icon: 'history', path: '/history' },
      ]
    : [
        { label: 'Public', icon: 'home', path: '/catalog' },
        { label: 'My Library', icon: 'inventory_2', path: '/my-library' },
        { label: 'History', icon: 'history', path: '/history' },
        { label: 'Profile', icon: 'person', path: '/profile' },
      ];

  return (
    <aside className="h-full border-r border-outline-variant bg-surface-container-low w-64 flex flex-col">
      <div className="p-lg flex items-center gap-sm">
        <img src={logo} alt="Logo" className="w-10 h-10 rounded-default object-cover" />
        <div>
          <h1 className="font-headline-sm text-primary">KH Library</h1>
          <p className="font-label-md text-on-surface-variant uppercase tracking-wider">Inventory System</p>
        </div>
      </div>
      <nav className="flex-1 px-sm mt-md space-y-xs">
        {navItems.map(item => (
          <div key={item.label}>
            {item.isSubmenu ? (
              <>
                <button
                  onClick={() => setIsAdminOpen(!isAdminOpen)}
                  className="w-full rounded-full p-md flex items-center gap-md cursor-pointer transition-all text-on-surface-variant hover:bg-surface-variant"
                >
                  <span className="material-symbols-outlined">{item.icon}</span>
                  <span className="font-label-md flex-1 text-left">{item.label}</span>
                  <span className="material-symbols-outlined">
                    {isAdminOpen ? 'expand_less' : 'expand_more'}
                  </span>
                </button>
                {isAdminOpen && (
                  <div className="ml-xl mt-xs space-y-xs">
                    {item.subItems?.map(sub => (
                      <NavLink
                        key={sub.label}
                        to={sub.path as string}
                        className={({ isActive }) =>
                          `rounded-full p-sm pl-lg flex items-center gap-md cursor-pointer transition-all text-sm ${
                            isActive 
                              ? 'text-primary font-bold' 
                              : 'text-on-surface-variant hover:text-primary'
                          }`
                        }
                      >
                        {sub.label}
                      </NavLink>
                    ))}
                  </div>
                )}
              </>
            ) : (
              <NavLink
                to={item.path as string}
                className={({ isActive }) =>
                  `rounded-full p-md flex items-center gap-md cursor-pointer transition-all ${
                    isActive 
                      ? 'bg-primary-container text-on-primary-container' 
                      : 'text-on-surface-variant hover:bg-surface-variant'
                  }`
                }
              >
                <span className="material-symbols-outlined">{item.icon}</span>
                <span className="font-label-md">{item.label}</span>
              </NavLink>
            )}
          </div>
        ))}
      </nav>
    </aside>
  );
};
