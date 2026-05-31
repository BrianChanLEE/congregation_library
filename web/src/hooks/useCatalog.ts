import { useQuery } from '@tanstack/react-query';
import { apiClient } from '../lib/axios';
import type { Item } from '../types/models';

// Note: API 호출 및 캐싱 훅
export const useCatalog = () => {
  return useQuery<Item[], Error>({
    queryKey: ['catalog'],
    queryFn: async () => {
      const { data } = await apiClient.get('/catalog');
      return data;
    },
  });
};

export const useCategories = () => {
  return useQuery<string[], Error>({
    queryKey: ['categories'],
    queryFn: async () => {
      const { data } = await apiClient.get('/catalog/categories');
      return data;
    },
  });
};
