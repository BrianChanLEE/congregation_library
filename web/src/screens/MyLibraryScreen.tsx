import React, { useState, useEffect } from 'react';
import { apiClient } from '../lib/axios';
import { useAuthStore } from '../store/authStore';

interface InventoryItem {
  id: number;
  name: string;
  category: string;
  stock: number;
  image_url: string;
}

export const MyLibraryScreen: React.FC = () => {
  const [items, setItems] = useState<InventoryItem[]>([]);
  const userId = useAuthStore(state => state.userId);

  useEffect(() => {
    if (userId) {
        // userId 기반으로 회중 정보 조회
        apiClient.get(`/user/profile?id=${userId}`)
            .then(res => {
                const congId = res.data.congregation_id;
                return apiClient.get(`/inventory?cong_id=${congId}`);
            })
            .then(res => setItems(res.data))
            .catch(err => console.error('내 서재 조회 실패', err));
    }
  }, [userId]);

  return (
    <div className="flex-1 p-md lg:p-lg max-w-container-max mx-auto w-full">
      {/* ... 기존 UI ... */}
      {/* Grid Layout for Items */}
      <div className="grid grid-cols-1 md:grid-cols-2 xl:grid-cols-3 gap-md">
        {items.map((item) => (
          <div key={item.id} className="bg-surface-container-lowest border border-outline-variant rounded-xl overflow-hidden hover:shadow-xl transition-all group">
            <div className="flex">
                <div className="w-1/3 aspect-[3/4] relative overflow-hidden bg-surface-container">
                <img alt={item.name} className="w-full h-full object-cover group-hover:scale-105 transition-transform duration-500" src={item.image_url} />
                </div>
                <div className="p-md flex-1">
                <span className="font-label-md text-label-md text-secondary uppercase">{item.category}</span>
                <h3 className="font-headline-sm text-headline-sm text-primary mt-1 truncate">{item.name}</h3>
                <div className="mt-lg flex items-center justify-between">
                    <div className="flex flex-col">
                        <span className="font-label-md text-label-md text-on-surface-variant">보유 수량</span>
                        <span className={`font-body-md text-body-md font-semibold text-primary`}>
                            {item.stock}권
                        </span>
                    </div>
                </div>
                </div>
            </div>
          </div>
        ))}
      </div>
    </div>
  );
};
