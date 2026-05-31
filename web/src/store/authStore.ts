import { create } from 'zustand';
import { persist } from 'zustand/middleware';

export type UserRole = 'user' | 'admin';

interface AuthState {
  token: string | null;
  role: UserRole | null;
  userId: number | null;
  setAuth: (token: string, role: UserRole, userId: number) => void;
  logout: () => void;
}

export const useAuthStore = create<AuthState>()(
  persist(
    (set) => ({
      token: null,
      role: null,
      userId: null,
      setAuth: (token, role, userId) => set({ token, role, userId }),
      logout: () => set({ token: null, role: null, userId: null }),
    }),
    {
      name: 'auth-storage',
    }
  )
);
