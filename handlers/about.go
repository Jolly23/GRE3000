package handlers

import (
	"GRE3000/filters"
	"github.com/gofiber/fiber/v2"
)

func AboutPage(ctx *fiber.Ctx) error {
	user := filters.LoadUser(ctx)

	res := fiber.Map{
		"PageTitle": "About",
		"UserInfo":  user,
	}

	return ctx.Render("about", res)
}
