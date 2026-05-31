import React, { useState, useEffect } from 'react';
import { apiClient } from '../lib/axios';

// Note: 시스템 에러 로그 정보를 관리하는 인터페이스 정의
interface ErrorLog {
  id: number;
  error: string;
  time: string;
}

// Note: 시스템 상태 정보를 관리하는 인터페이스 정의
interface SystemStatus {
  status: string;
  error_logs: ErrorLog[];
  system_time: string;
}

export const SystemStatusScreen: React.FC = () => {
  const [status, setStatus] = useState<SystemStatus | null>(null);

  // Note: 시스템 상태 및 로그 조회
  const fetchStatus = () => {
    apiClient.get('/admin/system-status')
      .then(res => setStatus(res.data))
      .catch(err => console.error('시스템 상태 조회 실패', err));
  };

  useEffect(() => {
    fetchStatus();
  }, []);

  return (
    <div className="flex-1 p-md lg:p-lg max-w-container-max mx-auto w-full">
      <div className="mb-lg">
        <h1 className="font-headline-sm text-primary">시스템 상태 모니터링</h1>
        <p className="font-body-md text-on-surface-variant">서버의 실시간 상태 및 최근 에러 로그를 확인합니다.</p>
      </div>

      {/* 상태 요약 */}
      <div className="bg-surface-container-lowest border border-outline-variant p-md rounded-default mb-lg flex items-center justify-between">
        <span className="font-headline-sm text-on-surface">현재 상태:</span>
        <span className={`px-3 py-1 rounded-full font-bold text-sm ${status?.status === 'ok' ? 'bg-tertiary-container text-on-tertiary-container' : 'bg-error-container text-on-error-container'}`}>
            {status?.status?.toUpperCase() || 'UNKNOWN'}
        </span>
        <span className="font-body-sm text-on-surface-variant">기준 시각: {status?.system_time}</span>
      </div>

      {/* 에러 로그 */}
      <div className="bg-surface-container-lowest border border-outline-variant rounded-default overflow-hidden">
        <div className="p-md border-b border-outline-variant bg-primary/5">
            <h3 className="font-headline-sm text-error">최근 에러 로그 (상위 10건)</h3>
        </div>
        <div className="divide-y divide-outline-variant">
          {status?.error_logs.map((log) => (
            <div key={log.id} className="p-md hover:bg-secondary-container/5 transition-colors">
              <p className="font-body-md text-error font-semibold">{log.error}</p>
              <p className="font-label-sm text-on-surface-variant mt-xs">{new Date(log.time).toLocaleString()}</p>
            </div>
          ))}
          {(!status || status.error_logs.length === 0) && <div className="p-md text-center text-on-surface-variant">로그된 에러가 없습니다.</div>}
        </div>
      </div>
    </div>
  );
};
