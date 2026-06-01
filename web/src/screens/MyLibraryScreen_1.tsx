import React, { useState } from 'react';
import { useMyLibraryItems } from '../hooks/useMyLibrary';

/**
 * Note 1
 * MyLibraryScreen_1 컴포넌트: React의 함수형 컴포넌트입니다.
 * props나 상태(state)를 이용해 UI를 렌더링합니다.
 */

/**
 * Note 2
 * 함수형 컴포넌트는 UI를 반환하는 자바스크립트 함수입니다.
 * React의 Hook(useState 등)을 사용하여 상태를 관리할 수 있습니다.
 */
export const MyLibraryScreen_1: React.FC = () => {
  /**
   * Note 3
   * useState Hook: 컴포넌트 내부에서 상태를 관리합니다.
   * rotation은 현재 캐러셀의 회전 각도 상태를 저장합니다.
   */
  const [rotation, setRotation] = useState(0);
  const { data: carouselItems, isLoading, error } = useMyLibraryItems();

  /**
   * Note 5
   * stepAngle: 캐러셀의 각 항목이 차지하는 각도를 계산합니다.
   * 계산된 값을 이용해 3D 회전을 구현합니다.
   */
  const stepAngle = carouselItems ? 360 / carouselItems.length : 0;

  /**
   * Note 6
   * rotateBy 함수: 사용자가 버튼을 클릭하면 회전 각도를 변경합니다.
   * React의 상태 업데이트 함수 setRotation을 호출하여 재렌더링을 유도합니다.
   */
  const rotateBy = (direction: number) => {
    setRotation(prev => prev + direction * stepAngle);
  };

  /**
   * Note 7
   * isActive 함수: 특정 인덱스의 항목이 현재 캐러셀의 정면(활성 상태)에 있는지 확인합니다.
   */
  const isActive = (index: number) => {
    if (!carouselItems) return false;
    const angle = stepAngle * index;
    const currentRingRot = (rotation % 360 + 360) % 360;
    const targetRingRot = (360 - angle) % 360;
    const diff = Math.abs(currentRingRot - targetRingRot);
    return diff < 10 || diff > 350;
  };

  if (isLoading) return <div className="p-lg">Loading...</div>;
  if (error) return <div className="p-lg">Error: {error.message}</div>;
  if (!carouselItems) return <div className="p-lg">No items.</div>;

  return (
    <div className="bg-background text-on-surface min-h-screen flex flex-col font-body-md">
      {/* TopAppBar - 상단 앱 바 */}
      <header className="sticky top-0 z-50 flex justify-between items-center w-full px-md h-16 max-w-7xl mx-auto bg-surface border-b border-outline-variant glass-header">
        <div className="flex items-center gap-sm">
          <span className="font-headline-md text-headline-md font-bold text-primary">KH Library</span>
        </div>
        <nav className="hidden md:flex items-center gap-lg">
          <a className="text-primary font-bold border-b-2 border-primary pb-1 font-label-md text-label-md" href="#">Public</a>
          <a className="text-on-surface-variant hover:bg-surface-container-low transition-colors font-label-md text-label-md" href="#">My Library</a>
          <a className="text-on-surface-variant hover:bg-surface-container-low transition-colors font-label-md text-label-md" href="#">History</a>
        </nav>
      </header>

      <main className="flex-1 w-full max-w-7xl mx-auto px-md py-lg overflow-x-hidden">
        {/* Carousel Section */}
        <section className="mb-xl relative rounded-xl py-xl px-md overflow-hidden bg-gradient-to-b from-surface to-surface-container">
          <h2 className="font-headline-lg text-headline-lg mb-xl relative z-10">주요 자료</h2>
          
          <div className="carousel-scene flex items-center justify-center h-[440px]">
            <div className="camera-tilt" style={{ transform: 'rotateX(12deg) translateY(20px)' }}>
              <div style={{ transform: `rotateY(${rotation}deg)` }} className="carousel-ring transition-transform duration-600 ease-[cubic-bezier(0.2,1,0.3,1)]">
                {carouselItems.map((item: any, index: number) => (
                  <div key={index} className="carousel-card flex flex-col bg-white rounded-xl overflow-hidden cursor-pointer shadow-lg" 
                       style={{ 
                         transform: `rotateY(${stepAngle * index}deg) translateZ(420px)`,
                         opacity: isActive(index) ? 1 : 0.5 
                       }}>
                    <div className="flex-1 relative bg-surface-container-low overflow-hidden">
                      <img alt={item.name} src={item.image_url} className="w-full h-full object-cover"/>
                      <div className={`bg-tertiary-fixed text-on-tertiary-fixed absolute top-4 right-4 px-3 py-1 rounded-full text-label-md flex items-center gap-1 shadow-sm`}>
                        <span className={`bg-tertiary-fixed-dim w-2 h-2 rounded-full`}></span> {item.stock}권
                      </div>
                    </div>
                    <div className="p-lg bg-white border-t border-outline-variant/20">
                      <span className="text-on-surface-variant font-label-md text-label-md">{item.category}</span>
                      <h3 className="font-headline-sm text-headline-sm text-primary mt-1">{item.name}</h3>
                    </div>
                  </div>
                ))}
              </div>
            </div>
            {/* Navigation */}
            <div className="absolute top-1/2 -translate-y-1/2 w-full flex justify-between px-4 z-40">
              <button onClick={() => rotateBy(1)} className="w-12 h-12 rounded-full bg-white/90 shadow-lg flex items-center justify-center border border-outline-variant hover:bg-surface-container text-primary">
                <span className="material-symbols-outlined">chevron_left</span>
              </button>
              <button onClick={() => rotateBy(-1)} className="w-12 h-12 rounded-full bg-white/90 shadow-lg flex items-center justify-center border border-outline-variant hover:bg-surface-container text-primary">
                <span className="material-symbols-outlined">chevron_right</span>
              </button>
            </div>
          </div>
        </section>
      </main>

      {/* BottomNavBar - 하단 네비게이션 바 */}
      <nav className="lg:hidden fixed bottom-0 left-0 w-full z-50 flex justify-around items-center px-sm py-xs bg-surface border-t border-outline-variant shadow-sm">
        {/* 네비게이션 아이템 */}
      </nav>
    </div>
  );
};
