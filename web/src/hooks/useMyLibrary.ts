import { useQuery } from '@tanstack/react-query';
import { apiClient } from '../lib/axios';

// Note: 사용자의 서재 자료 조회 훅입니다.
export const useMyLibraryItems = () => {
  return useQuery({
    queryKey: ['my-library-items'],
    queryFn: async () => {
      const { data } = await apiClient.get('/inventory');
      return data;
    },
  });
};
