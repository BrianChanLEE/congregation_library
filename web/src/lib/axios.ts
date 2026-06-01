import axios from 'axios';
import { useAuthStore } from '../store/authStore';

// Note 1: API 호출을 위한 공통 Axios 인스턴스입니다.
export const apiClient = axios.create({
  baseURL: '/api', // Vite 프록시 설정 활용
});

// Note 2: 요청 인터셉터를 사용하여 모든 요청에 JWT 토큰을 자동으로 주입합니다.
apiClient.interceptors.request.use((config) => {
  const { token } = useAuthStore.getState();
  if (token) {
    config.headers.Authorization = `Bearer ${token}`;
  }
  return config;
});

// Note 3: 인증 만료/무효(401)일 때는 토큰을 제거하고 로그인 화면으로 이동합니다.
apiClient.interceptors.response.use(
  (res) => res,
  (err) => {
    if (err?.response?.status === 401) {
      useAuthStore.getState().logout();
      if (typeof window !== 'undefined') {
        window.location.hash = '/login';
      }
    }
    return Promise.reject(err);
  }
);
