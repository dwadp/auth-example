package bcrypt

import (
	"github.com/dwadp/auth-example/pkg/hash"
	"golang.org/x/crypto/bcrypt"
)

type Bcrypt struct{}

func NewBcrypt() hash.Hash {
	return &Bcrypt{}
}

func (b *Bcrypt) Make(plainText string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(plainText), bcrypt.DefaultCost)

	return string(hashed), err
}

func (b *Bcrypt) Check(plainText, hashedText string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedText), []byte(plainText))

	return err == nil
}
