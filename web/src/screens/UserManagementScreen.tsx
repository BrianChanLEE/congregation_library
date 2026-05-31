import React, { useState, useEffect } from 'react';
import { apiClient } from '../lib/axios';

// Note: 사용자 정보를 관리하는 인터페이스 정의
interface User {
  id: number;
  name: string;
  jwhub_email: string;
  status: 'PENDING' | 'APPROVED' | 'REJECTED';
}

export const UserManagementScreen: React.FC = () => {
  const [users, setUsers] = useState<User[]>([]);

  // Note: 가입 대기 사용자 목록을 불러오는 함수
  const fetchPendingUsers = () => {
    apiClient.get('/admin/users/pending')
      .then(res => setUsers(res.data || []))
      .catch(err => console.error('대기 사용자 조회 실패', err));
  };

  useEffect(() => {
    fetchPendingUsers();
  }, []);

  // Note: 사용자 승인 상태 업데이트
  const handleUpdateStatus = async (id: number, status: 'APPROVED' | 'REJECTED') => {
    try {
      await apiClient.put(`/admin/users/${id}/status`, { status });
      fetchPendingUsers(); // 목록 새로고침
      alert(`사용자가 ${status === 'APPROVED' ? '승인' : '거절'}되었습니다.`);
    } catch (err) {
      console.error('상태 변경 실패', err);
    }
  };

  return (
    <div className="flex-1 p-md lg:p-lg max-w-container-max mx-auto w-full">
      <div className="mb-lg">
        <h1 className="font-headline-sm text-primary">사용자 관리</h1>
        <p className="font-body-md text-on-surface-variant">신규 가입 사용자를 승인하거나 거절합니다.</p>
      </div>

      {/* 가입 대기 목록 */}
      <div className="bg-surface-container-lowest border border-outline-variant rounded-default overflow-hidden">
        <div className="p-md border-b border-outline-variant bg-primary/5">
            <h3 className="font-headline-sm text-primary">가입 대기 목록</h3>
        </div>
        <div className="divide-y divide-outline-variant">
          {users.map((user) => (
            <div key={user.id} className="p-md flex justify-between items-center hover:bg-secondary-container/5 transition-colors">
              <div>
                <p className="font-body-md text-on-surface font-semibold">{user.name}</p>
                <p className="font-body-sm text-on-surface-variant">{user.jwhub_email}</p>
              </div>
              <div className="space-x-sm">
                <button 
                    onClick={() => handleUpdateStatus(user.id, 'APPROVED')} 
                    className="bg-primary text-on-primary px-sm py-xs rounded text-label-md"
                >
                    승인
                </button>
                <button 
                    onClick={() => handleUpdateStatus(user.id, 'REJECTED')} 
                    className="bg-error text-on-error px-sm py-xs rounded text-label-md"
                >
                    거절
                </button>
              </div>
            </div>
          ))}
          {users.length === 0 && <div className="p-md text-center text-on-surface-variant">대기 중인 사용자가 없습니다.</div>}
        </div>
      </div>
    </div>
  );
};
