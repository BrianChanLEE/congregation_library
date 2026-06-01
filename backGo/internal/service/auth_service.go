package service

import (
	"boock/backGo/internal/models"
	"boock/backGo/internal/repository"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"os"
	"strconv"
	"strings"
	"time"
)

// AuthServiceInterface 는 인증 관련 비즈니스 로직을 정의하는 인터페이스입니다.
type AuthServiceInterface interface {
	Register(name, password string) error
	Login(congCode, email, password string) (string, error)
}

// AuthService 는 AuthServiceInterface 를 구현하는 실제 서비스 구조체입니다.
type AuthService struct {
	userRepo repository.UserRepositoryInterface
}

// NewAuthService 는 새로운 AuthService 인스턴스를 생성합니다.
func NewAuthService(repo repository.UserRepositoryInterface) *AuthService {
	return &AuthService{userRepo: repo}
}

// Register 는 새로운 사용자를 등록합니다. 비밀번호는 bcrypt로 암호화되어 저장됩니다.
func (s *AuthService) Register(name, password string) error {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user := &models.User{
		Name:         name,
		PasswordHash: string(hashed),
	}
	return s.userRepo.Create(user)
}

// Login 은 사용자 정보를 확인하고 JWT 토큰을 발급합니다.
func (s *AuthService) Login(congCode, email, password string) (string, error) {
	// 1. 회중 코드 변환
	congID, err := strconv.ParseInt(congCode, 10, 64)
	if err != nil {
		return "", errors.New("잘못된 회중 코드")
	}

	// 2. 이메일과 회중 ID로 사용자 조회
	user, err := s.userRepo.GetByJwhubEmailAndCongID(email, congID)
	if err != nil {
		return "", errors.New("인증 실패: 사용자를 찾을 수 없습니다.")
	}

	// 3. 비밀번호 검증
	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
		return "", errors.New("인증 실패: 비밀번호가 일치하지 않습니다.")
	}

	// 4. JWT 클레임 설정 및 토큰 생성
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  user.JWhubEmail,
		"userId": user.ID,
		"role":   strings.ToLower(user.Role),
		"exp":    time.Now().Add(time.Hour * 24).Unix(), // 24시간 후 만료
	})

	// 5. 비밀키로 서명하여 토큰 반환
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		return "", errors.New("JWT_SECRET 환경 변수가 설정되지 않았습니다.")
	}
	return token.SignedString([]byte(secret))
}
