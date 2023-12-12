package handlers

import (
	"GRE3000/database"
	"GRE3000/filters"
	"github.com/gofiber/fiber/v2"
)

func VocabularyPage(ctx *fiber.Ctx) error {
	user := filters.LoadUser(ctx)
	db := ctx.Locals("db").(*database.Database)
	isRandomSort := ctx.QueryBool("RandomSort")

	res := fiber.Map{
		"PageTitle": "GRE3000",
		"UserInfo":  user,
	}
	if user != nil {
		userWordsList := db.LoadUserWords(user.ID, isRandomSort)
		res["UserWords"] = &userWordsList
		res["PageTitle"] = user.Username + "的单词表"
	} else {
		rawWordsList := db.LoadRawWords(isRandomSort)
		res["RawWords"] = &rawWordsList
		res["PageTitle"] = "GRE单词表"
	}

	res["IsWordsPage"] = true

	res["RandomSort"] = isRandomSort

	return ctx.Render("words/vocabulary", res)
}
