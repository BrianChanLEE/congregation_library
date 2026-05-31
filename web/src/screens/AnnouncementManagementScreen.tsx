import React, { useState, useEffect } from 'react';
import { apiClient } from '../lib/axios';

// Note: 공지사항 데이터를 관리하는 인터페이스 정의
interface Announcement {
  id: number;
  title: string;
  content: string;
  created_at: string;
}

export const AnnouncementManagementScreen: React.FC = () => {
  const [announcements, setAnnouncements] = useState<Announcement[]>([]);
  const [title, setTitle] = useState('');
  const [content, setContent] = useState('');

  // Note: 공지사항 목록 조회
  const fetchAnnouncements = () => {
    apiClient.get('/admin/announcements')
      .then(res => setAnnouncements(res.data || []))
      .catch(err => console.error('공지사항 조회 실패', err));
  };

  useEffect(() => {
    fetchAnnouncements();
  }, []);

  // Note: 공지사항 등록
  const handleAddAnnouncement = async (e: React.FormEvent) => {
    e.preventDefault();
    try {
      await apiClient.post('/admin/announcements', { title, content, author_id: 1 }); // 임시 author_id
      setTitle('');
      setContent('');
      fetchAnnouncements();
      alert('공지사항이 등록되었습니다.');
    } catch (err) {
      console.error('공지사항 등록 실패', err);
    }
  };

  // Note: 공지사항 삭제
  const handleDelete = async (id: number) => {
    if (!confirm('정말 삭제하시겠습니까?')) return;
    try {
      await apiClient.delete(`/admin/announcements/${id}`);
      fetchAnnouncements();
    } catch (err) {
      console.error('공지사항 삭제 실패', err);
    }
  };

  return (
    <div className="flex-1 p-md lg:p-lg max-w-container-max mx-auto w-full">
      <div className="mb-lg">
        <h1 className="font-headline-sm text-primary">공지사항 관리</h1>
        <p className="font-body-md text-on-surface-variant">시스템 공지사항을 등록하고 관리합니다.</p>
      </div>

      {/* 등록 폼 */}
      <div className="bg-surface-container-lowest border border-outline-variant p-md rounded-default mb-lg">
        <h3 className="font-headline-sm text-primary mb-md">신규 공지 등록</h3>
        <form onSubmit={handleAddAnnouncement} className="space-y-md">
          <input className="w-full p-sm border rounded" placeholder="제목" value={title} onChange={e => setTitle(e.target.value)} required />
          <textarea className="w-full p-sm border rounded h-32" placeholder="내용" value={content} onChange={e => setContent(e.target.value)} required />
          <button type="submit" className="bg-primary text-on-primary px-md py-sm rounded">공지 등록</button>
        </form>
      </div>

      {/* 목록 */}
      <div className="bg-surface-container-lowest border border-outline-variant rounded-default overflow-hidden">
        <div className="p-md border-b border-outline-variant bg-primary/5">
            <h3 className="font-headline-sm text-primary">등록된 공지사항</h3>
        </div>
        <div className="divide-y divide-outline-variant">
          {announcements.map((ann) => (
            <div key={ann.id} className="p-md flex justify-between items-start hover:bg-secondary-container/5 transition-colors">
              <div>
                <p className="font-body-md text-on-surface font-semibold">{ann.title}</p>
                <p className="font-body-sm text-on-surface-variant whitespace-pre-line">{ann.content}</p>
                <p className="font-label-sm text-on-surface-variant mt-sm">{new Date(ann.created_at).toLocaleString()}</p>
              </div>
              <button 
                onClick={() => handleDelete(ann.id)} 
                className="bg-error text-on-error px-sm py-xs rounded text-label-md ml-md"
              >
                삭제
              </button>
            </div>
          ))}
          {announcements.length === 0 && <div className="p-md text-center text-on-surface-variant">등록된 공지사항이 없습니다.</div>}
        </div>
      </div>
    </div>
  );
};
