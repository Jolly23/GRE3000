package handlers

import (
	"GRE3000/filters"
	"github.com/gofiber/fiber/v2"
)

func LogoutHandler(ctx *fiber.Ctx) error {
	filters.SetSecureCookie(ctx, "token", "")
	filters.SetSecureCookie(ctx, "name", "")

	return ctx.Redirect("/", 302)
}
