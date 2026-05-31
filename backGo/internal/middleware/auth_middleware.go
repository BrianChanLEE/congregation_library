package middleware

import (
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// AuthMiddleware 는 JWT 토큰을 검증하고 사용자 정보를 컨텍스트에 저장하는 미들웨어입니다.
// 이 미들웨어는 모든 보호된 경로에 대해 인증 여부를 확인합니다.
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 1. Authorization 헤더 추출
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization 헤더가 필요합니다."})
			return
		}

		// 2. Bearer 토큰 형식 확인
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Bearer {token} 형식이 아닙니다."})
			return
		}

		// 3. 토큰 파싱 및 검증
		tokenString := parts[1]
		jwtSecret := os.Getenv("JWT_SECRET")
		if jwtSecret == "" {
			jwtSecret = "default_secret" // 보안을 위해 .env 설정을 권장
		}

		claims := jwt.MapClaims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtSecret), nil
		})

		// 4. 유효하지 않은 토큰 처리
		if err != nil || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "유효하지 않은 토큰입니다."})
			return
		}

		// 5. 클레임에서 사용자 정보 추출 및 컨텍스트 저장
		// float64로 저장된 숫자를 int64로 변환할 때 타입 단언(Type Assertion) 주의
		userIdFloat, ok := claims["userId"].(float64)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "토큰 내 사용자 정보가 올바르지 않습니다."})
			return
		}

		userRole, ok := claims["role"].(string)
		if !ok {
			userRole = "user" // 기본 역할 설정
		}

		c.Set("userId", int64(userIdFloat))
		c.Set("userRole", userRole)
		c.Next()
	}
}

// AdminMiddleware 는 관리자 권한(admin)이 있는지 확인하는 미들웨어입니다.
// AuthMiddleware 다음에 실행되어야 하며, 컨텍스트의 userRole을 검사합니다.
func AdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		role, exists := c.Get("userRole")
		if !exists || role != "admin" {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "관리자 권한이 필요합니다."})
			return
		}
		c.Next()
	}
}
