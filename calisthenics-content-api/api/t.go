package api

import (
	"fmt"
	"github.com/o1egl/paseto"
	"time"
)

func main() {
	// PASETO V2 anahtarı oluştur
	v2 := paseto.NewV2()

	// Symmetric anahtar (32 byte)
	key := []byte("supersecretkeyyoushouldnotcommit") // Uzunluğu tam olarak 32 byte olmalı

	// Token için payload oluştur
	payload := paseto.JSONToken{
		Audience:   "example",
		Issuer:     "example.com",
		Jti:        "token-id",
		Subject:    "subject",
		IssuedAt:   time.Now(),
		Expiration: time.Now().Add(24 * time.Hour),
		NotBefore:  time.Now(),
	}

	// Ek veriler (isteğe bağlı)
	footer := A2{
		Name:       "Khan",
		FamilyName: "Calisthenics",
	}

	// Token'ı oluştur
	token, err := v2.Encrypt(key, payload, footer)
	if err != nil {
		fmt.Println("Token oluşturulurken hata:", err)
		return
	}

	fmt.Println("Oluşturulan Token:", token)

	// Token'ı doğrula
	var newPayload paseto.JSONToken
	var newFooter A2
	err = v2.Decrypt(token, key, &newPayload, &newFooter)
	if err != nil {
		fmt.Println("Token doğrulanırken hata:", err)
		return
	}

	// Payload ve footer'ı yazdır
	fmt.Printf("Payload: %+v\n", newPayload)
	fmt.Println("Footer:", newFooter)
}

type A2 struct {
	Name       string
	FamilyName string
}
