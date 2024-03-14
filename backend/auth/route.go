package auth

import (
	"github.com/gofiber/fiber/v2"
)

func InitRoutes(app *fiber.App) {
	app.Post("auth/login", func(c *fiber.Ctx) error {
		c.Accepts(fiber.MIMEApplicationJSON)

		var request LoginRequest
		err := c.BodyParser(&request)
		if err != nil {
			return c.Status(409).SendString("Cannot parse body")
		}
		success, user := IsAuthAndGetUser(request.UserName, request.Password)
		if !success {
			return NotAuthorizedResponse(c)
		}
		session, err := CreateSession(user)
		if err != nil {
			return NotAuthorizedResponse(c)
		}
		return c.JSON(LoginResponse{
			User:         user,
			SessionToken: session.Id,
		})
	})
	app.Get("auth/logout", func(c *fiber.Ctx) error {
		sessionToken := c.Get("Authorization")
		session, err := GetSession(sessionToken)
		if err != nil {
			return NotAuthorizedResponse(c)
		}
		SessionRepository.Delete(session)
		return c.Status(200).JSON(session)
	})

	app.Get("auth/session", func(c *fiber.Ctx) error {
		sessionToken := c.Get("Authorization")
		session, err := GetSession(sessionToken)
		if err != nil {
			return NotAuthorizedResponse(c)
		}
		user, err := Repository.FindByIdOrErr(session.UserId)
		if err != nil {
			return NotAuthorizedResponse(c)
		}
		return c.Status(200).JSON(user)
	})
}

func NotAuthorizedResponse(ctx *fiber.Ctx) error {
	return ctx.Status(401).SendString("Not authorized")
}

type LoginResponse struct {
	User         *User  `json:"user"`
	SessionToken string `json:"sessionToken"`
}
type LoginRequest struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}
