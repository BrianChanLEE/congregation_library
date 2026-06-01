import { useQuery } from '@tanstack/react-query';
import { apiClient } from '../lib/axios';

export interface MovementHistoryItem {
  id: number;
  time?: string;
  created_at?: string;
  item?: string;
  item_name?: string;
  qty?: number;
  quantity?: number;
  user?: string;
  user_name?: string;
  memo?: string;
  type: string;
}

export const useMovementHistory = () => {
  return useQuery<MovementHistoryItem[], Error>({
    queryKey: ['movement-history'],
    queryFn: async () => {
      const { data } = await apiClient.get('/history');
      return data || [];
    },
  });
};
