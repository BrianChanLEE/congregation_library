# PRD: 서적 관리 사용자 및 어드민 권한 기반 라우팅

## Problem Statement
현재 프로젝트는 사용자(회중)와 어드민(LCC)의 화면이 분리되어 있지 않아, 접근 제어 없이 모든 화면에 접근 가능한 상태입니다. 이는 보안상 취약하며, UX 측면에서도 역할에 맞는 최적화된 경험을 제공하기 어렵습니다.

## Solution
사용자와 어드민 역할을 구분하는 권한 기반 라우팅 시스템을 도입합니다. JWT의 `role` 정보를 활용하여 인증된 사용자만 적절한 레이아웃에 접근하도록 제어합니다.

## User Stories
1. As a User, I want to access my dashboard, so that I can view my personal information and tasks.
2. As an Admin, I want to access the LCC Dashboard, so that I can manage library inventory and literature records.
3. As an Unauthenticated User, I want to be redirected to the login page, so that I can securely authenticate.
4. As a User, I want to be denied access to Admin pages, so that security is maintained.
5. As an Admin, I want my dashboard layout to be distinct, so that I can focus on management tasks.

## Implementation Decisions
- **Router**: `BrowserRouter` (React Router) 사용.
- **State Management**: `Zustand`를 사용하여 JWT 상태 및 사용자 role 관리.
- **Access Control**: `ProtectedRoute` 컴포넌트로 역할 기반 필터링 구현.
- **Layouts**: `UserLayout`과 `AdminLayout`으로 시각적/구조적 분리.
- **Auth**: JWT 토큰을 통한 인증 및 role 추출.

## Testing Decisions
- 인증/인가 로직(라우팅 가드)에 대한 단위 테스트.
- 각 역할별 접근 권한 테스트 (권한이 없는 경로 접근 시 리다이렉트).
- Zustand 스토어 상태 변경 테스트.

## Out of Scope
- OAuth/SSO 통합 (기본 JWT 인증 중심).
- 백엔드 관리자용 추가 API 개발 (권한 미들웨어는 별도).

## Further Notes
- `LccDashboard`, `InventoryManage`, `EditLiterature`는 관리자 전용으로 즉시 이동.
