package handlers

import (
	"GRE3000/types"
	"github.com/gofiber/fiber/v2"
)

func AboutPage(ctx *fiber.Ctx) error {
	res := fiber.Map{
		"PageTitle": "About",
	}

	res["IsLogin"] = true
	res["UserInfo"] = types.User{
		Username:  "",
		Token:     "",
		Avatar:    "",
		Email:     "",
		Signature: "",
	}

	return ctx.Render("about", res)
}
