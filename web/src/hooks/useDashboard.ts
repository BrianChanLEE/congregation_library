import { useQuery } from '@tanstack/react-query';
import { apiClient } from '../lib/axios';
import type { Transaction, AdminStats, Alert } from '../types/models';

export const useTransactions = (congId?: number) => {
  return useQuery<Transaction[], Error>({
    queryKey: ['transactions', congId],
    queryFn: async () => {
      const url = congId ? `/transactions?cong_id=${congId}` : '/transactions';
      const { data } = await apiClient.get(url);
      return data;
    },
  });
};

export const useAdminStats = () => {
  return useQuery<AdminStats, Error>({
    queryKey: ['admin-stats'],
    queryFn: async () => {
      const { data } = await apiClient.get('/admin/stats');
      return data;
    },
  });
};

export const useAlerts = () => {
  return useQuery<Alert[], Error>({
    queryKey: ['alerts'],
    queryFn: async () => {
      const { data } = await apiClient.get('/admin/alerts');
      return data;
    },
  });
};
