package middleware

import (
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/csrf"
	"github.com/gofiber/fiber/v3/middleware/session"
	"github.com/gofiber/utils/v2"
)

var csrfConfig = csrf.Config{
	KeyLookup:         "header:" + csrf.HeaderName,
	CookieName:        "__Host-csrf_",
	CookieSameSite:    "Strict",
	CookieSecure:      true,
	CookieSessionOnly: true,
	CookieHTTPOnly:    true,
	Expiration:        1 * time.Hour,
	KeyGenerator:      utils.UUIDv4,
	ErrorHandler:      csrf.ConfigDefault.ErrorHandler,
	Extractor:         csrf.FromHeader(csrf.HeaderName),
	Session:           session.New(),
	SessionKey:        "csrfToken",
}

func NewCsrf() fiber.Handler {
	return csrf.New(csrfConfig)
}
