import React, { useState, useEffect } from 'react';
import { apiClient } from '../lib/axios';

// Note: 품목 정보를 관리하는 인터페이스 정의
interface Item {
  id: number;
  name: string;
  code: string;
}

export const InventoryInScreen: React.FC = () => {
  const [items, setItems] = useState<Item[]>([]);
  const [itemName, setItemName] = useState('');
  const [itemCode, setItemCode] = useState('');
  const [selectedItemId, setSelectedItemId] = useState<number | ''>('');
  const [quantity, setQuantity] = useState(0);
  const [memo, setMemo] = useState('');

  // Note: 품목 목록을 불러오는 함수
  const fetchItems = () => {
    apiClient.get('/admin/items')
      .then(res => setItems(res.data || []))
      .catch(err => console.error('품목 조회 실패', err));
  };

  useEffect(() => {
    fetchItems();
  }, []);

  // Note: 신규 품목 등록
  const handleAddItem = async (e: React.FormEvent) => {
    e.preventDefault();
    try {
      await apiClient.post('/admin/items', { name: itemName, code: itemCode });
      setItemName('');
      setItemCode('');
      fetchItems();
      alert('품목이 추가되었습니다.');
    } catch (err) {
      console.error('품목 추가 실패', err);
    }
  };

  // Note: 입고 거래 생성
  const handleInbound = async (e: React.FormEvent) => {
    e.preventDefault();
    if (!selectedItemId) return;
    try {
      await apiClient.post('/transactions', {
        user_id: 1, // 임시 사용자 ID (추후 인증 연동 필요)
        item_id: selectedItemId,
        quantity: quantity,
        type: 'IN',
        method: 'WEB',
        memo: memo
      });
      setQuantity(0);
      setMemo('');
      alert('입고 처리되었습니다.');
    } catch (err) {
      console.error('입고 처리 실패', err);
    }
  };

  return (
    <div className="flex-1 p-md lg:p-lg max-w-container-max mx-auto w-full">
      <div className="mb-lg">
        <h1 className="font-headline-sm text-primary">입고 관리</h1>
        <p className="font-body-md text-on-surface-variant">신규 품목 등록 및 물품 입고를 처리합니다.</p>
      </div>

      <div className="grid grid-cols-1 md:grid-cols-2 gap-lg">
        {/* 신규 품목 등록 */}
        <div className="bg-surface-container-lowest border border-outline-variant p-md rounded-default">
          <h3 className="font-headline-sm text-primary mb-md">신규 품목 등록</h3>
          <form onSubmit={handleAddItem} className="space-y-md">
            <input className="w-full p-sm border rounded" placeholder="품목 이름" value={itemName} onChange={e => setItemName(e.target.value)} required />
            <input className="w-full p-sm border rounded" placeholder="품목 코드" value={itemCode} onChange={e => setItemCode(e.target.value)} required />
            <button type="submit" className="bg-primary text-on-primary px-md py-sm rounded">품목 추가</button>
          </form>
        </div>

        {/* 입고 처리 */}
        <div className="bg-surface-container-lowest border border-outline-variant p-md rounded-default">
          <h3 className="font-headline-sm text-primary mb-md">물품 입고</h3>
          <form onSubmit={handleInbound} className="space-y-md">
            <select className="w-full p-sm border rounded" value={selectedItemId} onChange={e => setSelectedItemId(Number(e.target.value))} required>
              <option value="">품목 선택</option>
              {items.map(item => <option key={item.id} value={item.id}>{item.name} ({item.code})</option>)}
            </select>
            <input type="number" className="w-full p-sm border rounded" placeholder="수량" value={quantity} onChange={e => setQuantity(Number(e.target.value))} required />
            <input className="w-full p-sm border rounded" placeholder="메모" value={memo} onChange={e => setMemo(e.target.value)} />
            <button type="submit" className="bg-tertiary text-on-tertiary px-md py-sm rounded">입고 처리</button>
          </form>
        </div>
      </div>
    </div>
  );
};
