import { BrowserRouter, Routes, Route, Navigate } from 'react-router-dom';
import { Suspense } from 'react';
import { InstallPrompt } from './components/common/InstallPrompt';
import { ProtectedRoute } from './components/common/ProtectedRoute';
import { DashboardLayout } from './components/layouts/DashboardLayout';

// Available screens
import { LoginScreen } from './screens/LoginScreen';
import { PublicLibraryScreen } from './screens/PublicLibraryScreen';
import { MyLibraryScreen } from './screens/MyLibraryScreen';
import { MovementHistoryScreen } from './screens/MovementHistoryScreen';
import { AdminDashboard } from './screens/AdminDashboard';
import { QrScanScreen } from './screens/QrScanScreen';

function App() {
  return (
    <>
      <BrowserRouter basename={import.meta.env.BASE_URL}>
        <Suspense fallback={<div className="flex h-screen items-center justify-center">로딩 중...</div>}>
          <Routes>
            <Route path="/login" element={<LoginScreen />} />
            
            {/* 보호된 라우트 */}
            <Route element={<ProtectedRoute />}>
              <Route element={<DashboardLayout />}>
                <Route path="/" element={<PublicLibraryScreen />} />
                <Route path="/catalog" element={<PublicLibraryScreen />} />
                <Route path="/my-library" element={<MyLibraryScreen />} />
                <Route path="/history" element={<MovementHistoryScreen />} />
                <Route path="/scan" element={<QrScanScreen />} />
                
                {/* 관리자 전용 라우트 */}
                <Route element={<ProtectedRoute requiredRole="admin" />}>
                  <Route path="/admin" element={<AdminDashboard />} />
                </Route>
              </Route>
            </Route>

            <Route path="*" element={<Navigate to="/" replace />} />
          </Routes>
        </Suspense>
      </BrowserRouter>
      <InstallPrompt />
    </>
  );
}

export default App;
