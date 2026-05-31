// Note 1: PWA 설치 안내 스낵바 컴포넌트.
// Note 2: 브라우저의 beforeinstallprompt 이벤트를 감지하여 설치 유도.
import React, { useState, useEffect } from 'react';

// Note 3: 브라우저의 BeforeInstallPromptEvent 타입 정의.
interface BeforeInstallPromptEvent extends Event {
  prompt: () => void;
  userChoice: Promise<{ outcome: 'accepted' | 'dismissed' }>;
}

export const InstallPrompt: React.FC = () => {
  const [deferredPrompt, setDeferredPrompt] = useState<BeforeInstallPromptEvent | null>(null);
  const [show, setShow] = useState<boolean>(false);

  useEffect(() => {
    // Note 4: 이벤트 리스너 등록.
    const handler = (e: Event) => {
      e.preventDefault();
      setDeferredPrompt(e as BeforeInstallPromptEvent);
      setShow(true);
    };

    window.addEventListener('beforeinstallprompt', handler as EventListener);
    return () => window.removeEventListener('beforeinstallprompt', handler as EventListener);
  }, []);

  // Note 5: 설치 유도 처리.
  const handleInstall = () => {
    if (deferredPrompt) {
      deferredPrompt.prompt();
      setShow(false);
    }
  };

  if (!show) return null;

  return (
    <div className="fixed bottom-md left-md right-md md:left-auto md:right-md md:w-80 bg-surface-container-high border border-outline-variant p-md rounded-default shadow-lg flex items-center justify-between gap-md z-[60]">
      <p className="font-body-md text-on-surface">앱을 설치하여 더 빠르게 사용하세요!</p>
      <div className="flex gap-sm">
        <button className="font-label-md text-on-surface-variant px-sm" onClick={() => setShow(false)}>닫기</button>
        <button className="font-label-md text-secondary px-sm" onClick={handleInstall}>설치</button>
      </div>
    </div>
  );
};
