package handlers

import (
	"GRE3000/database"
	"GRE3000/filters"
	"github.com/gofiber/fiber/v2"
)

func VocabularyPage(ctx *fiber.Ctx) error {
	user := filters.LoadUser(ctx)
	isRandomSort := ctx.QueryBool("rs")

	res := fiber.Map{
		"PageTitle": "GRE3000",
		"UserInfo":  user,
	}
	if user != nil {
		res["PageTitle"] = user.Username + "的单词表"
	} else {
		res["PageTitle"] = "GRE单词表"
	}

	res["IsWordsPage"] = true

	res["RandomSort"] = isRandomSort

	return ctx.Render("words/vocabulary", res)
}

func LoadWordsHandler(ctx *fiber.Ctx) error {
	user := filters.LoadUser(ctx)
	db := ctx.Locals("db").(*database.Database)
	isRandomSort := ctx.QueryBool("rs")

	if user != nil {
		userWordsList := db.LoadUserWords(user.ID, isRandomSort)
		return ctx.JSON(userWordsList)
	}

	rawWordsList := db.LoadRawWords(isRandomSort)
	return ctx.JSON(rawWordsList)
}
