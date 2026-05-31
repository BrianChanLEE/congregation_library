import { useQuery } from '@tanstack/react-query';
import { apiClient } from '../lib/axios';
import type { UserProfile } from '../types/models';

// Note: 사용자 프로필 조회 훅 (ID 기반)
export const useProfile = (userId: number) => {
  return useQuery<UserProfile, Error>({
    queryKey: ['profile', userId],
    queryFn: async () => {
      if (!userId || userId <= 0) throw new Error('Invalid user ID');
      const { data } = await apiClient.get(`/user/profile?id=${userId}`);
      return data;
    },
    enabled: !!userId && userId > 0,
  });
};
