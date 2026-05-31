import React, { useState } from 'react';
import { BrowserMultiFormatReader } from '@zxing/library';
import { ItemDetailModal } from '../components/features/ItemDetailModal';
import { apiClient } from '../lib/axios';
import type { Item } from '../types/models';

export const QrScanScreen: React.FC = () => {
  const [selectedItem, setSelectedItem] = useState<Item | null>(null);
  const [error, setError] = useState<string | null>(null);

  React.useEffect(() => {
    const reader = new BrowserMultiFormatReader();
    reader.decodeFromConstraints({ video: { facingMode: 'environment' } }, 'video', (result) => {
      if (result) {
        const code = result.getText();
        // 품목 코드로 조회
        apiClient.get(`/catalog?code=${code}`)
            .then(res => setSelectedItem(res.data))
            .catch(() => setError('품목을 찾을 수 없습니다.'));
      }
    });
    return () => reader.reset();
  }, []);

  const handleConfirm = (quantity: number) => {
    apiClient.post('/transactions', {
        to_cong_id: 1,
        item_id: selectedItem?.id,
        quantity: quantity,
        type: 'OUT'
    }).then(() => {
        alert('성공적으로 차감되었습니다.');
        setSelectedItem(null);
    });
  };

  return (
    <div className="flex-1 flex flex-col items-center justify-center p-md">
      <video id="video" className="w-full max-w-sm aspect-square bg-black rounded-xl" />
      <p className="mt-md text-on-surface-variant">QR 코드를 스캔하세요.</p>
      {error && <p className="text-error mt-sm">{error}</p>}
      
      {selectedItem && (
        <ItemDetailModal 
          isOpen={!!selectedItem}
          onClose={() => setSelectedItem(null)}
          item={selectedItem}
          onConfirm={handleConfirm}
        />
      )}
    </div>
  );
};
