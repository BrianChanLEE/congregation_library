import { Outlet } from 'react-router-dom';
import { TopAppBar } from '../common/TopAppBar';
import { SideNavBar } from '../common/SideNavBar';
import { MobileNavBar } from '../common/MobileNavBar';

export const DashboardLayout = () => {
  return (
    <div className="flex h-screen overflow-hidden bg-background">
      {/* 데스크탑 사이드바 - Level 1 */}
      <div className="hidden md:flex w-64 flex-col border-r border-outline-variant bg-surface-container-lowest">
        <SideNavBar />
      </div>

      <div className="flex flex-1 flex-col overflow-hidden">
        <TopAppBar />
        <main className="flex-1 overflow-y-auto p-md md:p-lg bg-surface">
          <div className="max-w-7xl mx-auto">
            <Outlet />
          </div>
        </main>
        
        {/* 모바일 하단바 */}
        <div className="md:hidden border-t border-outline-variant bg-surface-container-lowest">
          <MobileNavBar />
        </div>
      </div>
    </div>
  );
};
