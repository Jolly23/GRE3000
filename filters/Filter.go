package filters

import (
	"GRE3000/database"
	"GRE3000/types"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"strings"
	"time"
)

func IsAuthenticated(ctx *fiber.Ctx) bool {
	token, ok := getSecureCookie(ctx, "token")
	return ok && token != ""
}

func LoadUser(ctx *fiber.Ctx) *types.User {
	token, ok := getSecureCookie(ctx, "token")
	if !ok {
		return nil
	}

	db := ctx.Locals("db").(*database.Database)
	return db.FindUserByToken(token)
}

func getSecureCookie(ctx *fiber.Ctx, key string) (string, bool) {
	val := ctx.Cookies(key)
	if val == "" {
		return "", false
	}

	parts := strings.SplitN(val, "|", 3)

	if len(parts) != 3 {
		return "", false
	}

	vs := parts[0]
	timestamp := parts[1]
	sig := parts[2]

	h := hmac.New(sha256.New, []byte("3e7wQQ4BCc4C7MJq3ycn3kjYF3fNGXQT6TBWgKnTCKfinA8HVNyNjXGa4fJLJCsj"))
	fmt.Fprintf(h, "%s%s", vs, timestamp)

	if fmt.Sprintf("%02x", h.Sum(nil)) != sig {
		return "", false
	}
	res, _ := base64.URLEncoding.DecodeString(vs)
	return string(res), true
}

func SetSecureCookie(ctx *fiber.Ctx, name, value string) {
	vs := base64.URLEncoding.EncodeToString([]byte(value))
	timestamp := strconv.FormatInt(time.Now().UnixNano(), 10)

	h := hmac.New(sha256.New, []byte("3e7wQQ4BCc4C7MJq3ycn3kjYF3fNGXQT6TBWgKnTCKfinA8HVNyNjXGa4fJLJCsj"))
	fmt.Fprintf(h, "%s%s", vs, timestamp)
	sig := fmt.Sprintf("%02x", h.Sum(nil))
	cookie := strings.Join([]string{vs, timestamp, sig}, "|")

	ctx.Cookie(&fiber.Cookie{
		Name:     name,
		Value:    cookie,
		Expires:  time.Now().Add(365 * 24 * time.Hour),
		Secure:   true,
		HTTPOnly: true,
	})
}
