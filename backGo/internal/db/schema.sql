-- 1. 사용자 (Users) - 승인/권한 관리 중심
CREATE TABLE users (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY COMMENT '사용자 고유 식별자',
    name VARCHAR(50) NOT NULL COMMENT '사용자 이름',
    role ENUM('ADMIN', 'USER') DEFAULT 'USER' COMMENT '사용자 역할',
    status ENUM('PENDING', 'APPROVED', 'REJECTED') DEFAULT 'PENDING' COMMENT '가입 승인 상태',
    jwhub_email VARCHAR(100) UNIQUE NOT NULL COMMENT 'JW Hub 이메일',
    password_hash VARCHAR(255) NOT NULL COMMENT '비밀번호 해시',
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '가입 신청 일시',
    deleted_at DATETIME NULL COMMENT '삭제 일시'
);

-- 2. 서적/품목 (Items)
CREATE TABLE items (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY COMMENT '품목 고유 식별자',
    code VARCHAR(50) UNIQUE NOT NULL COMMENT '품목 코드',
    name VARCHAR(100) NOT NULL COMMENT '품목 이름',
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '생성 일시',
    deleted_at DATETIME NULL COMMENT '삭제 일시'
);

-- 3. 활동 로그 (Activity Logs) - 모니터링 중심
CREATE TABLE activity_logs (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY COMMENT '로그 고유 식별자',
    user_id BIGINT UNSIGNED COMMENT '작업자 식별자',
    item_id BIGINT UNSIGNED NOT NULL COMMENT '품목 식별자',
    quantity INT NOT NULL COMMENT '수량 (+/-)',
    type ENUM('IN', 'OUT', 'CANCEL') NOT NULL COMMENT '입고/출고/취소 구분',
    method ENUM('WEB', 'QR') NOT NULL COMMENT '작업 방식',
    memo TEXT COMMENT '메모',
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '기록 일시',
    FOREIGN KEY (user_id) REFERENCES users (id),
    FOREIGN KEY (item_id) REFERENCES items (id)
);

-- 4. 공지사항 (Announcements)
CREATE TABLE announcements (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY COMMENT '공지 고유 식별자',
    title VARCHAR(255) NOT NULL COMMENT '제목',
    content TEXT NOT NULL COMMENT '내용',
    author_id BIGINT UNSIGNED NOT NULL COMMENT '작성자',
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '작성 일시',
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '수정 일시',
    deleted_at DATETIME NULL COMMENT '삭제 일시',
    FOREIGN KEY (author_id) REFERENCES users (id)
);
