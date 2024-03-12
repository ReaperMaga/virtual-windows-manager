package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/valyala/bytebufferpool"
	"strings"
	"virtual-windows-manager/auth"
)

func AuthMiddleware(ctx *fiber.Ctx) error {
	sessionToken := ctx.Cookies("session_token")
	if strings.Contains(ctx.OriginalURL(), "/login") {
		if auth.IsAuth(sessionToken) {
			return ctx.Redirect("/")
		}
		return ctx.Next()
	}
	if auth.IsAuth(sessionToken) {
		return ctx.Next()
	}
	bb := bytebufferpool.Get()
	defer bytebufferpool.Put(bb)
	bb.WriteString("/login?")
	bb.Write(ctx.Request().URI().QueryString())
	return ctx.Redirect(bb.String())
}
