import React, { useState, useEffect } from 'react';
import { apiClient } from '../lib/axios';

// Note 1: 타입 정의(Interface)는 코드의 가독성과 타입 안정성을 높여줍니다.
// TypeScript를 사용할 때 데이터를 구조화하여 관리하는 핵심 요소입니다.
interface Transaction {
  id: number;
  time: string;
  item: string;
  qty: number;
  user: string;
  memo: string;
  type: string;
}

// Note 2: React 함수형 컴포넌트는 UI를 함수로 표현합니다.
// React.FC를 사용하여 타입 정의가 명확한 컴포넌트를 만듭니다.
export const MovementHistoryScreen: React.FC = () => {
  // Note 3: useState 훅은 컴포넌트의 상태를 관리합니다.
  // 여기서는 거래 목록(transactions)을 배열 상태로 관리합니다.
  // 초기값으로 빈 배열([])을 사용하여 렌더링 시점에 오류를 방지합니다.
  const [transactions, setTransactions] = useState<Transaction[]>([]);

  // Note 4: useEffect 훅은 컴포넌트 생명주기에 따라 부수 효과(Side Effect)를 수행합니다.
  // 컴포넌트가 처음 마운트(화면에 나타남)될 때 API를 호출하여 데이터를 가져옵니다.
  // 빈 배열([])을 의존성 배열로 전달하면 마운트 시에만 실행됩니다.
  useEffect(() => {
    // 이동 내역 조회 API 호출
    apiClient.get('/history')
        // Note 5: API 응답이 성공하면 setTransactions를 사용하여 상태를 업데이트합니다.
        // 데이터가 없거나 null일 경우를 대비해 OR 연산자(||)로 빈 배열을 기본값으로 사용합니다.
        // 이는 'null'로 인한 런타임 오류(TypeError: map of null)를 방지하는 방어적 프로그래밍 기법입니다.
        .then(res => setTransactions(res.data || []))
        // Note 6: catch 블록은 에러가 발생했을 때 상황을 처리합니다.
        // 콘솔에 로그를 남겨 디버깅을 돕습니다.
        .catch(err => console.error('이동 기록 조회 실패', err));
  }, []);

  // Note 7: JSX는 React에서 UI를 작성하는 자바스크립트의 확장 문법입니다.
  // 가독성이 좋고 컴포넌트 구조를 한눈에 파악할 수 있게 합니다.
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
            {/* Note 8: map 함수를 사용하여 배열 데이터를 리스트 형태의 UI로 변환(렌더링)합니다.
                리스트 렌더링 시에는 각 항목에 유일한 'key' 값을 주어야 React가 성능을 최적화할 수 있습니다. */}
            {transactions.map((tx) => (
              <tr key={tx.id} className="hover:bg-surface-container-low transition-colors">
                <td className="p-md font-body-sm text-on-surface-variant">{tx.time}</td>
                <td className="p-md font-body-md font-semibold text-primary">{tx.item}</td>
                <td className="p-md font-body-md text-right">{tx.qty}</td>
                <td className="p-md font-body-md">{tx.user}</td>
                <td className="p-md font-body-sm text-on-surface-variant">{tx.memo}</td>
                <td className="p-md">
                    {/* Note 9: 조건부 렌더링을 사용하여 데이터의 'type' 값에 따라 다른 스타일을 적용합니다.
                        CSS 클래스를 동적으로 조합하여 UI에 피드백을 제공합니다. */}
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
