.PHONY: all api web clean

# 백엔드와 프론트엔드를 동시에 빌드 및 실행
all:
	@echo "서적 관리 통합 빌드 시작..."
	@make -j2 api web

# 백엔드 서버 실행 (backGo/cmd/api)
api:
	@echo "백엔드 서버 실행 중..."
	@cd backGo && go run cmd/api/main.go

# 프론트엔드 개발 서버 실행 (web/)
web:
	@echo "프론트엔드 개발 서버 실행 중..."
	@cd web && npm run dev

# 빌드 산출물 정리
clean:
	@echo "정리 작업 중..."
	@rm -rf web/dist api server
