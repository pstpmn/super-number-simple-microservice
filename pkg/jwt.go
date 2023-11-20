package pkg

import (
	"math"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
)

type (
	IJwt interface {
		SignToken(claims any) string
	}

	Claims struct {
		data any
		jwt.RegisteredClaims
	}

	j struct {
		Secret   []byte
		Duration int64
	}
)

// SignToken implements IJwt.
func (s *j) SignToken(data any) string {
	var mapClaims = &Claims{
		data: data,
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   "access-token",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(s.Duration * int64(math.Pow10(9))))),
			NotBefore: jwt.NewNumericDate(time.Now()),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, mapClaims)
	ss, _ := token.SignedString(s.Secret)
	return ss
}

func NewJwt(secret []byte, duration int64) IJwt {
	return &j{
		Secret:   secret,
		Duration: duration,
	}
}
