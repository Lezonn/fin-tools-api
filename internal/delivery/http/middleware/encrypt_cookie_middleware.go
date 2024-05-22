package middleware

import (
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/encryptcookie"
)

var encryptCookieConfig = encryptcookie.Config{
	Next:      nil,
	Except:    []string{},
	Key:       encryptcookie.GenerateKey(),
	Encryptor: encryptcookie.EncryptCookie,
	Decryptor: encryptcookie.DecryptCookie,
}

func NewEncryptCookie() fiber.Handler {
	return encryptcookie.New(encryptCookieConfig)
}
