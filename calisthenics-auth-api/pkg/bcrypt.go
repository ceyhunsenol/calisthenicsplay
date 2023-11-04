package pkg

import (
	"golang.org/x/crypto/bcrypt"
)

// HashPassword şifreyi bcrypt ile hashler
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	return string(bytes), err
}

// CheckPasswordHash verilen düz metin şifre ile hashlenmiş şifreyi karşılaştırır
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
