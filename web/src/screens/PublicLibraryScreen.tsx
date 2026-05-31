import React, { useState } from 'react';
import { ItemDetailModal } from '../components/features/ItemDetailModal';
import { apiClient } from '../lib/axios';
import type { Item } from '../types/models';
import { useCatalog } from '../hooks/useCatalog';

export const PublicLibraryScreen: React.FC = () => {
  const [activeCategory, setActiveCategory] = useState('전체');
  const { data: items = [], isLoading, error } = useCatalog();
  const [selectedItem, setSelectedItem] = useState<Item | null>(null);

  const filteredItems = items.filter(item => 
    activeCategory === '전체' || item.category === activeCategory
  );

  const handleConfirm = (quantity: number) => {
    apiClient.post('/transactions', {
        to_cong_id: 1,
        item_id: selectedItem?.id,
        quantity: quantity,
        type: 'OUT'
    })
    .then(() => {
        alert('성공적으로 가져갔습니다.');
        setSelectedItem(null);
    })
    .catch(() => alert('처리 실패'));
  };

  if (isLoading) return <div>로딩 중...</div>;
  if (error) return <div>에러 발생</div>;

  return (
    <div className="flex-1">
      {/* ... 기존 JSX ... */}

      <div className="sticky top-0 z-40 bg-background/95 backdrop-blur-md pb-md">
        <div className="flex flex-col gap-sm">
          <div className="relative">
            <div className="absolute inset-y-0 left-0 pl-sm flex items-center pointer-events-none text-on-surface-variant">
              <span className="material-symbols-outlined">search</span>
            </div>
            <input 
              className="w-full pl-lg pr-md py-md bg-surface-container-lowest border border-outline-variant rounded-default outline-none font-body-md text-on-surface focus:border-primary" 
              placeholder="품목 코드 또는 이름 검색..." 
              type="text"
            />
          </div>
          <div className="flex overflow-x-auto gap-xs">
            {['전체', '성서', '봉사 도구함', '공개 증거', '서책', '팜플렛', '양식 및 공급품'].map((cat) => (
              <button 
                key={cat}
                onClick={() => setActiveCategory(cat)}
                className={`px-md py-sm whitespace-nowrap rounded-full font-label-md transition-colors ${
                  activeCategory === cat 
                    ? 'bg-primary text-on-primary' 
                    : 'bg-surface-container-lowest text-on-surface-variant border border-outline-variant'
                }`}
              >
                {cat}
              </button>
            ))}
          </div>
        </div>
      </div>

      {/* Item Grid */}
      <div className="grid grid-cols-2 md:grid-cols-3 xl:grid-cols-4 gap-md mt-md">
        {filteredItems.map((item) => (
          <div key={item.id} onClick={() => setSelectedItem(item)} className="bg-surface-container-lowest border border-outline-variant rounded-default overflow-hidden cursor-pointer">
            <div className="aspect-[3/4] relative overflow-hidden bg-surface-container">
              {item.imageUrl ? (
                <img alt={item.name} className="w-full h-full object-cover" src={item.imageUrl} />
              ) : (
                <div className="w-full h-full flex items-center justify-center text-on-surface-variant">
                  <span className="material-symbols-outlined text-4xl">image</span>
                </div>
              )}
            </div>
            <div className="p-md">
              <h3 className="font-headline-sm text-primary mb-xs">{item.name}</h3>
              <p className="font-body-sm text-on-surface-variant mb-md">{item.code}</p>
              <span className={`px-sm py-xs rounded-full font-label-md ${
                item.stock > 10 ? 'bg-tertiary-container text-on-tertiary-container' :
                item.stock > 0 ? 'bg-error-container text-on-error-container' :
                'bg-surface-container text-on-surface-variant'
              }`}>
                {item.stock > 0 ? '재고 ' + item.stock : '품절'}
              </span>
            </div>
          </div>
        ))}
      </div>

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
