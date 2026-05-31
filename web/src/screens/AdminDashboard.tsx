import React, { useEffect, useState } from 'react';
import { apiClient } from '../lib/axios';

// Note: 백엔드에서 가져올 통계 데이터 타입 정의
interface Stats {
  total_items: number;
  recent_activity_count: number;
  pending_user_count: number;
}

// Note: 백엔드에서 가져올 활동 로그 데이터 타입 정의
interface ActivityLog {
  id: number;
  created_at: string;
  user_name: string;
  item_name: string;
  quantity: number;
  type: string; // 'IN' or 'OUT'
  method: string;
  memo: string;
}

export const AdminDashboard: React.FC = () => {
  const [stats, setStats] = useState<Stats | null>(null);
  const [logs, setLogs] = useState<ActivityLog[]>([]);

  useEffect(() => {
    // Note: 통계 및 활동 로그 데이터를 병렬로 가져와 로딩 워터폴(Waterfalls)을 방지합니다.
    Promise.all([
      apiClient.get('/admin/stats'),
      apiClient.get('/history')
    ])
      .then(([statsRes, logsRes]) => {
        setStats(statsRes.data);
        setLogs(logsRes.data || []);
      })
      .catch(err => console.error('대시보드 데이터 로딩 실패', err));
  }, []);

  return (
    <div className="flex-1 p-md lg:p-lg max-w-container-max mx-auto w-full">
      <div className="mb-lg">
        <h1 className="font-headline-sm text-primary">관리자 대시보드</h1>
        <p className="font-body-md text-on-surface-variant">시스템 운영 상태 및 활동 현황을 모니터링합니다.</p>
      </div>

      {/* Stats Cards */}
      <div className="grid grid-cols-1 md:grid-cols-3 gap-md mb-lg">
        <div className="bg-surface-container-lowest border border-outline-variant p-md rounded-default">
            <p className="font-label-md text-on-surface-variant uppercase tracking-wider">전체 품목 수</p>
            <p className="font-headline-sm text-on-surface">{stats?.total_items || 0}개</p>
        </div>
        <div className="bg-surface-container-lowest border border-outline-variant p-md rounded-default">
            <p className="font-label-md text-on-surface-variant uppercase tracking-wider">최근 7일 활동</p>
            <p className="font-headline-sm text-on-surface">{stats?.recent_activity_count || 0}건</p>
        </div>
        <div className="bg-surface-container-lowest border border-outline-variant p-md rounded-default">
            <p className="font-label-md text-on-surface-variant uppercase tracking-wider">가입 대기 사용자</p>
            <p className="font-headline-sm text-primary font-bold">{stats?.pending_user_count || 0}명</p>
        </div>
      </div>

      {/* Activity Timeline */}
      <div className="bg-surface-container-lowest border border-outline-variant rounded-default overflow-hidden">
        <div className="p-md border-b border-outline-variant bg-primary/5">
            <h3 className="font-headline-sm text-primary">최근 활동 로그</h3>
        </div>
        <div className="divide-y divide-outline-variant">
          {logs.map((log) => (
            <div key={log.id} className="p-md hover:bg-secondary-container/5 transition-colors">
              <div className="flex justify-between items-start mb-sm">
                <span className={`px-2 py-0.5 rounded text-[10px] font-bold uppercase ${
                    log.type === 'IN' ? 'bg-tertiary-container text-on-tertiary-container' : 'bg-secondary-container text-on-secondary-container'
                }`}>
                    {log.type === 'IN' ? '입고' : '출고'}
                </span>
                <span className="font-label-sm text-on-surface-variant">{new Date(log.created_at).toLocaleString()}</span>
              </div>
              <p className="font-body-md text-on-surface font-semibold">{log.item_name} <span className="font-normal text-on-surface-variant">({log.quantity}개)</span></p>
              <p className="font-body-sm text-on-surface-variant">사용자: {log.user_name} | 방식: {log.method}</p>
              {log.memo && <p className="font-body-sm mt-sm italic text-on-surface-variant">메모: {log.memo}</p>}
            </div>
          ))}
          {logs.length === 0 && <div className="p-md text-center text-on-surface-variant">최근 활동이 없습니다.</div>}
        </div>
      </div>
    </div>
  );
};
