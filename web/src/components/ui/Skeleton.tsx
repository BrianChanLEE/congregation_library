import React from 'react';

interface SkeletonProps {
  className?: string;
}

// Note: DESIGN.md의 디자인 토큰을 활용한 재사용 가능한 Skeleton 컴포넌트
export const Skeleton: React.FC<SkeletonProps> = ({ className }) => {
  return (
    <div 
      className={`animate-pulse bg-surface-container-high rounded-md ${className}`} 
    />
  );
};
