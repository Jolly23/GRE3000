package handlers

import (
	"GRE3000/database"
	"GRE3000/filters"
	"github.com/gofiber/fiber/v2"
)

func LoginPage(ctx *fiber.Ctx) error {
	if filters.IsAuthenticated(ctx) {
		_ = ctx.Redirect("/", 302)
	}

	res := fiber.Map{
		"PageTitle": "登录",
	}

	return ctx.Render("login", res)
}

func LoginHandler(ctx *fiber.Ctx) error {
	username := ctx.FormValue("username")
	password := ctx.FormValue("password")

	if len(username) == 0 || len(password) == 0 {
		return ctx.Render("login", fiber.Map{
			"PageTitle":   "Login",
			"flash_error": "Empty username or password",
		})
	}

	db := ctx.Locals("db").(*database.Database)

	if token, ok := db.AuthUser(username, password); ok {
		filters.SetSecureCookie(ctx, "token", token)
		filters.SetSecureCookie(ctx, "name", username)
		return ctx.Redirect("/", 302)
	}

	return ctx.Render("login", fiber.Map{
		"PageTitle":   "Login",
		"flash_error": "Wrong username or password",
	})
}
