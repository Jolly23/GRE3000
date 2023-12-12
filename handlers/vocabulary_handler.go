package handlers

import (
	"GRE3000/database"
	"GRE3000/filters"
	"errors"
	"github.com/gofiber/fiber/v2"
)

func UserVocabularyRebuild(ctx *fiber.Ctx) error {
	user := filters.LoadUser(ctx)
	if user == nil {
		return errors.New("user not found")
	}

	db := ctx.Locals("db").(*database.Database)
	db.GenerateUserWord(user.ID)

	return ctx.Redirect("/", 302)
}

func UserWordMarkIncHandler(ctx *fiber.Ctx) error {
	userID, wordID, err := getUserIDAndWordID(ctx)
	if err != nil {
		return err
	}

	db := ctx.Locals("db").(*database.Database)
	return db.MarkUserWord(userID, wordID)
}

func UserWordDelete(ctx *fiber.Ctx) error {
	userID, wordID, err := getUserIDAndWordID(ctx)
	if err != nil {
		return err
	}

	db := ctx.Locals("db").(*database.Database)
	return db.DeleteUserWord(userID, wordID)
}

func getUserIDAndWordID(ctx *fiber.Ctx) (int, int, error) {
	wordID, err := ctx.ParamsInt("id")
	if err != nil {
		return 0, 0, err
	}

	user := filters.LoadUser(ctx)
	if user != nil {
		return user.ID, wordID, nil
	}
	return 0, 0, errors.New("user not found")
}
