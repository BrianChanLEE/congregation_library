# Backend Test Coverage & Refactoring Issues

1. **[AFK] 리팩토링: `AnnouncementService` 및 `AnnouncementRepository` 의존성 주입**
   - **무엇을**: `AnnouncementService`를 인터페이스 기반으로 DI 수정.
   - **완료 조건**: `AnnouncementService`가 `NewAnnouncementService(repo AnnouncementRepoInterface)` 구조를 가짐.

2. **[AFK] 단위 테스트: `AnnouncementService` (커버리지 80% 달성)**
   - **무엇을**: `Table-driven` 테스트를 사용하여 Happy/Error 케이스 작성.
   - **완료 조건**: `announcement_service_test.go` 커버리지 80% 이상.

3. **[AFK] 리팩토링: `AuthService` 및 `UserRepository` 의존성 주입**
   - **무엇을**: `AuthService`를 인터페이스 기반으로 DI 수정.
   - **완료 조건**: `AuthService`가 DI 구조를 가짐.

4. **[AFK] 단위 테스트: `AuthService` (커버리지 80% 달성)**
   - **무엇을**: `Table-driven` 테스트 작성. Happy/Error(로그인 실패, 유저 없음 등) 필수 포함.
   - **완료 조건**: `auth_service_test.go` 커버리지 80% 이상.

5. **[AFK] 리팩토링: `CatalogService` 및 `ItemRepository` 의존성 주입**
   - **무엇을**: `CatalogService`를 인터페이스 기반으로 DI 수정.
   - **완료 조건**: `CatalogService`가 DI 구조를 가짐.

6. **[AFK] 단위 테스트: `CatalogService` (커버리지 80% 달성)**
   - **무엇을**: `Table-driven` 테스트 작성.
   - **완료 조건**: `catalog_service_test.go` 커버리지 80% 이상.
