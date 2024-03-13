package middleware

import (
	"github.com/gofiber/fiber/v2"
	"strings"
	"virtual-windows-manager/auth"
)

func AuthMiddleware(ctx *fiber.Ctx) error {
	sessionToken := ctx.Get("Authorization")
	if auth.IsAuth(sessionToken) {
		return ctx.Next()
	}
	if strings.HasPrefix(ctx.OriginalURL(), "/auth/login") {
		return ctx.Next()
	}
	return ctx.Status(401).SendString("Not authorized")
}
