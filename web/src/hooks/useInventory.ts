import { useQuery } from '@tanstack/react-query';
import { apiClient } from '../lib/axios';
import type { Inventory } from '../types/models';

// Note: TanStack Query를 사용한 재고 정보 조회 훅입니다.
// 데이터 캐싱과 자동 갱신을 지원합니다.
export const useInventory = (congId: number, itemId: number) => {
  return useQuery<Inventory, Error>({
    queryKey: ['inventory', congId, itemId],
    queryFn: async () => {
      const { data } = await apiClient.get(`/inventory?cong_id=${congId}&item_id=${itemId}`);
      return data;
    },
  });
};
