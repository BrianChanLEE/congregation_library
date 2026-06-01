import React from 'react';
import { useMovementHistory } from '../hooks/useMovementHistory';

export const MovementHistoryScreen: React.FC = () => {
  const { data: transactions = [], isError } = useMovementHistory();

  return (
    <div className="flex-1 p-md lg:p-lg max-w-container-max mx-auto w-full">
      {/* Header */}
      <div className="mb-lg flex flex-col md:flex-row md:items-end justify-between gap-md">
        <div>
          <h2 className="font-headline-lg text-headline-lg text-primary">물품 이동 기록</h2>
          <p className="font-body-md text-body-md text-on-surface-variant">비치 물품 및 서적 이동 내역을 확인하세요.</p>
        </div>
      </div>

      {/* Table */}
      <div className="bg-surface border border-outline-variant rounded-xl overflow-hidden shadow-sm">
        <table className="w-full border-collapse text-left">
          <thead className="bg-surface-container-low border-b border-outline-variant">
            <tr>
              <th className="p-md font-label-md text-on-surface-variant">시간</th>
              <th className="p-md font-label-md text-on-surface-variant">품목</th>
              <th className="p-md font-label-md text-on-surface-variant text-right">수량</th>
              <th className="p-md font-label-md text-on-surface-variant">처리자</th>
              <th className="p-md font-label-md text-on-surface-variant">메모</th>
              <th className="p-md font-label-md text-on-surface-variant">구분</th>
            </tr>
          </thead>
          <tbody className="divide-y divide-outline-variant">
            {isError && (
              <tr>
                <td className="p-md font-body-sm text-error" colSpan={6}>이동 기록을 불러오지 못했습니다.</td>
              </tr>
            )}
            {transactions.map((tx) => (
              <tr key={tx.id} className="hover:bg-surface-container-low transition-colors">
                <td className="p-md font-body-sm text-on-surface-variant">{tx.time ?? tx.created_at}</td>
                <td className="p-md font-body-md font-semibold text-primary">{tx.item ?? tx.item_name}</td>
                <td className="p-md font-body-md text-right">{tx.qty ?? tx.quantity}</td>
                <td className="p-md font-body-md">{tx.user ?? tx.user_name}</td>
                <td className="p-md font-body-sm text-on-surface-variant">{tx.memo}</td>
                <td className="p-md">
                    <span className={`px-2 py-0.5 rounded text-[10px] font-bold uppercase ${
                        tx.type === 'IN' ? 'bg-tertiary-container text-on-tertiary-container' : 'bg-secondary-container text-on-secondary-container'
                    }`}>
                        {tx.type === 'IN' ? '입고' : '출고'}
                    </span>
                </td>
              </tr>
            ))}
          </tbody>
        </table>
      </div>
    </div>
  );
};
