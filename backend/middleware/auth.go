package middleware

import (
	"fmt"
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
	fmt.Println("Unanuthorized")
	return ctx.Status(401).SendString("Not authorized")
}
