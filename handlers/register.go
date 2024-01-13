package handlers

import (
	"GRE3000/database"
	"GRE3000/filters"
	"github.com/gofiber/fiber/v2"
)

func RegisterPage(ctx *fiber.Ctx) error {
	if filters.IsAuthenticated(ctx) {
		_ = ctx.Redirect("/", 302)
	}

	res := fiber.Map{
		"PageTitle": "Register",
	}

	return ctx.Render("register", res)
}

func RegisterHandler(ctx *fiber.Ctx) error {
	username := ctx.FormValue("username")
	password := ctx.FormValue("password")

	if len(username) == 0 || len(password) == 0 {
		return ctx.Render("login", fiber.Map{
			"PageTitle":   "Register",
			"flash_error": "Empty username or password",
		})
	}

	db := ctx.Locals("db").(*database.Database)

	if userID, token, err := db.SaveUsername(username, password); err == nil {
		filters.SetSecureCookie(ctx, "token", token)
		filters.SetSecureCookie(ctx, "name", username)
		db.GenerateUserWord(userID)

		return ctx.Redirect("/", 302)
	}
	return ctx.Render("login", fiber.Map{
		"PageTitle":   "Register",
		"flash_error": "Username already exists",
	})
}
