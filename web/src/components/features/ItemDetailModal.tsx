import React, { useState } from 'react';
import type { Item } from '../../types/models';

interface ItemDetailModalProps {
  isOpen: boolean;
  onClose: () => void;
  item: Item;
  onConfirm: (quantity: number) => void;
}

export const ItemDetailModal: React.FC<ItemDetailModalProps> = ({ isOpen, onClose, item, onConfirm }) => {
  const [quantity, setQuantity] = useState(1);

  if (!isOpen) return null;

  return (
    <div className="fixed inset-0 z-50 flex items-center justify-center p-md bg-black/40 backdrop-blur-sm" id="modal-container">
      <div className="bg-surface-container-lowest w-full max-w-[500px] rounded-xl shadow-xl overflow-hidden flex flex-col max-h-[90vh]">

        {/* Header Image Section */}
        <div className="relative h-64 w-full bg-surface-container">
          {/* 이미지 URL이 존재할 때만 렌더링하여 빈 src 속성으로 인한 오류 방지 */}
        {item.imageUrl && (
          <img src={item.imageUrl} alt={item.name} className="w-full h-full object-cover" />
        )}
          <button className="absolute top-md right-md bg-surface/80 hover:bg-surface text-on-surface p-base rounded-full transition-colors flex items-center justify-center" onClick={onClose}>
            <span className="material-symbols-outlined text-[24px]">close</span>
          </button>
          <div className="absolute bottom-md left-md">
            <span className="bg-primary text-white font-label-md px-sm py-[4px] rounded-full uppercase tracking-wider">{item.category}</span>
          </div>
        </div>

        {/* Content Body */}
        <div className="flex-1 overflow-y-auto p-lg">
          <div className="flex justify-between items-start mb-sm">
            <h1 className="font-headline-md text-primary">{item.name}</h1>
            <div className="flex flex-col items-end">
              <span className="font-label-md text-on-surface-variant">현재 재고</span>
              <span className="font-headline-sm text-secondary">{item.stock}</span>
            </div>
          </div>

          <div className="bg-surface-container-low p-md rounded-lg border border-outline-variant mb-lg">
            <div className="flex items-center gap-sm mb-md">
              <span className="material-symbols-outlined text-primary">info</span>
              <span className="font-label-md text-on-surface-variant">가져갈 수량 선택</span>
            </div>

            <div className="flex items-center justify-between gap-md">
              <button 
                className="w-14 h-14 flex items-center justify-center bg-surface-container-highest hover:bg-surface-variant text-on-surface rounded-xl transition-all active:scale-90"
                onClick={() => setQuantity(Math.max(1, quantity - 1))}
              >
                <span className="material-symbols-outlined text-[28px]">remove</span>
              </button>
              <div className="flex-1 text-center">
                <input 
                  className="w-full text-center font-headline-md bg-transparent border-none focus:ring-0 text-primary"
                  type="number" 
                  value={quantity}
                  readOnly
                />
                <div className="h-[2px] bg-primary mx-auto w-16 mt-xs"></div>
              </div>
              <button 
                className="w-14 h-14 flex items-center justify-center bg-surface-container-highest hover:bg-surface-variant text-on-surface rounded-xl transition-all active:scale-90"
                onClick={() => setQuantity(Math.min(item.stock, quantity + 1))}
              >
                <span className="material-symbols-outlined text-[28px]">add</span>
              </button>
            </div>
          </div>
        </div>

        {/* Action Footer */}
        <div className="p-lg pt-0 flex gap-md">
          <button 
            className="flex-1 py-md rounded-xl font-label-md bg-surface-variant text-on-surface-variant hover:bg-outline-variant transition-colors"
            onClick={onClose}
          >
            취소
          </button>
          <button 
            className="flex-[2] py-md rounded-xl font-headline-sm bg-primary-container text-white hover:opacity-90 transition-all active:scale-95 flex items-center justify-center gap-sm"
            onClick={() => onConfirm(quantity)}
          >
            <span className="material-symbols-outlined">inventory_2</span>
            가져가기
          </button>
        </div>
      </div>
    </div>
  );
};
