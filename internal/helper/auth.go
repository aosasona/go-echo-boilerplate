package helper

import (
	"crypto/rand"
	"errors"
	"io"
	"time"

	"github.com/golang-jwt/jwt"
	"gopi/internal/config"
	apperrors "gopi/pkg/app_errors"

	"go.uber.org/zap"
)

type TokenType int

const (
	ACCESS_TOKEN TokenType = iota
	REFRESH_TOKEN
)

const (
	ACCESS_TOKEN_DURATION  = 5 * time.Minute
	REFRESH_TOKEN_DURATION = 30 * 24 * time.Hour
)

var ALLOWED_OTP_CONTENT = [...]byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}

func (h *helper) GenerateOTP(length int) string {
	otp := make([]byte, length)

	n, err := io.ReadAtLeast(rand.Reader, otp, length)
	if n != length || err != nil {
		h.logger.Error("failed to generate OTP", zap.Error(err))
		return ""
	}

	for i := 0; i < length; i++ {
		otp[i] = ALLOWED_OTP_CONTENT[otp[i]%byte(len(ALLOWED_OTP_CONTENT))]
	}

	return string(otp)
}

func (h *helper) GenerateJWT(userId uint64, tokenType TokenType) (string, error) {
	var (
		token  string
		secret string
		err    error
		ttl    time.Duration
	)

	switch tokenType {
	case ACCESS_TOKEN:
		ttl = ACCESS_TOKEN_DURATION
		if h.config.AppEnv == config.DEVELOPMENT {
			ttl = 21 * 24 * 60 * time.Minute
		}
		secret = h.config.AccessTokenSecret
	case REFRESH_TOKEN:
		ttl = REFRESH_TOKEN_DURATION
		secret = h.config.RefreshTokenSecret
	default:
		return token, errors.New(apperrors.ErrInvalidTokenType)
	}

	claims := jwt.MapClaims{
		"id":  userId,
		"exp": time.Now().Add(ttl).Unix(),
	}

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err = t.SignedString([]byte(secret))
	return token, err
}
