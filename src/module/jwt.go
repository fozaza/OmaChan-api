package module

import (
	"os"
	//"os/user"
	//"os/user"
	"time"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	//"github.com/golang-jwt/jwt"
	"github.com/golang-jwt/jwt/v5"
)

// jwt use for auth
// Create_jwt
func Cr_jwt(email string, level int) (string, error) {
	// set claims
	claims := jwt.MapClaims{
		"email": email,
		"level": level,
		"exp":   time.Now().Add(time.Hour * 3).Unix(),
	}

	// create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	println(claims["level"].(int))

	// Check token
	t, err := token.SignedString([]byte(os.Getenv("key")))
	if err != nil {
		return "", err
	}
	return t, nil
}

// userContextKey is the key used to store user data in the fiber context
// UserReturn for jwt
type UserReturn struct {
	Email string
	Level int
}

const userContextKey = "user"

func ExtractUserFromJWT(app *fiber.App) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var user UserReturn
		// Extract the token from the Fiber context (inserted by the JWT middleware)

		token := c.Locals("user").(*jwt.Token)
		claims := token.Claims.(jwt.MapClaims)

		user.Email = claims["email"].(string)
		//user.Level = claims["level_user"].(float64)

		//Store the user data in the Fiber context
		//c.Locals("exUser", &user)
		return c.Next()
	}
}

// Read token
func Get_token(c *fiber.Ctx) (string, int) {
	user := c.Locals(userContextKey).(*UserReturn)
	return user.Email, user.Level
}

func Req_level(levelreq int) fiber.Handler {
	return func(c *fiber.Ctx) error {
		user, ok := c.Locals("user").(*jwt.Token)

		if !ok || user == nil {
			return c.Status(fiber.StatusUnauthorized).
				JSON(fiber.Map{"error": "OmaChan >>> you dont have jwt"})
		}

		claims := user.Claims.(jwt.MapClaims)
		_, emailOk := claims["email"].(string)
		level, levelOk := claims["level"].(float64)
		if !levelOk || !emailOk {
			return c.Status(fiber.StatusInternalServerError).
				JSON(fiber.Map{"Error": "OmaChan >>> Invalid server"})
		}

		if int(level) < levelreq {
			return c.Status(fiber.StatusForbidden).
				JSON(fiber.Map{"error": "OmaChan >>> you cannot see here"})
		}
		return c.Next()
	}
}

func Con_jwt(app *fiber.App) any {
	return app.Use(jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(os.Getenv("key"))},
		// est context key ex.. c.locals("jwt")
		//ContextKey:   "jwt",
		ErrorHandler: error_jwt,
		// SuccessHandler: func(c *fiber.Ctx) error {
		// 	return c.Next()
		// },
	}))
}

// error hander
func error_jwt(c *fiber.Ctx, err error) error {
	if err.Error() == "Missing or malformed JWT" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "OmaChan >>> Missing or malformed JWT",
			"data":    nil,
		})
	}

	// For other errors, such as invalid or expired tokens
	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"status":  "error",
		"message": "OmaChan >>> Invalid or expired JWT",
		"data":    nil,
	})
}
