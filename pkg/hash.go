package pkg

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type (
	IHash interface {
		Bcrypt(data []byte) (string, error)
		CompareBcrypt(paintext []byte, hash []byte) bool
	}

	h struct {
	}
)

// bcrypt implements IHash.
func (*h) Bcrypt(data []byte) (string, error) {
	if hash, err := bcrypt.GenerateFromPassword(data, bcrypt.DefaultCost); err != nil {
		return "", errors.New("error: failed to hash password")
	} else {
		return string(hash), nil
	}
}

// compareBcrypt implements IHash.
func (*h) CompareBcrypt(paintext []byte, hash []byte) bool {
	if err := bcrypt.CompareHashAndPassword(hash, paintext); err != nil {
		return false
	}
	return true
}

func NewHash() IHash {
	return &h{}
}
