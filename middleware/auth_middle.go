package middleware

import (
	"github.com/gofiber/fiber/v2"
	"virtual-windows-manager/auth"
)

func AuthMiddleware(ctx *fiber.Ctx) error {
	sessionToken := ctx.Cookies("session_token")
	if ctx.OriginalURL() == "/login" {
		if auth.IsAuth(sessionToken) {
			return ctx.Redirect("/")
		}
		return ctx.Next()
	}
	if auth.IsAuth(sessionToken) {
		return ctx.Next()
	}
	return ctx.Redirect("/login")
}
