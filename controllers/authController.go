package controllers

import (
	"strconv"
	"time"

	"github.com/RyanCheungJF/CVWO-Backend/database"
	"github.com/RyanCheungJF/CVWO-Backend/models"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

// Secret Key used for hashing in JWTs
const SECRET_KEY = "secret"

func Register(c *fiber.Ctx) error {
	var data map[string]string

	err := c.BodyParser(&data)
	if err != nil {
		return err
	}

	// Hash password
	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)

	// Defining user fields
	user := models.User{
		Email:    data["email"],
		Password: password,
	}

	//Insertion of user into database
	database.DB.Create(&user)

	return c.JSON(user)
}

func Login(c *fiber.Ctx) error {
	var data map[string]string

	err := c.BodyParser(&data)
	if err != nil {
		return err
	}

	var user models.User

	// Get user
	database.DB.Where("email = ?", data["email"]).First(&user)

	// Check if user exists
	if user.ID == 0 {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"meesage": "User not found",
		})
	}

	// Check if password matches hash
	err = bcrypt.CompareHashAndPassword(user.Password, []byte(data["password"]))
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "Incorrect Password",
		})
	}

	// Issue JSON Web Token
	// claims are info to be exchanged
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, &jwt.RegisteredClaims{
		Issuer:    strconv.Itoa(int(user.ID)),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
	})

	token, err := claims.SignedString([]byte(SECRET_KEY))
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "Failed to login",
		})
	}

	// Create Cookie (Only stored in frontend without access)
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Minute * 30),
		HTTPOnly: true,
	}
	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "Success",
		"userid": user.ID,
	})
}

func User(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")

	token, err := jwt.ParseWithClaims(cookie, &jwt.RegisteredClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(SECRET_KEY), nil
		},
	)

	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	claims := token.Claims.(*jwt.RegisteredClaims)

	var user models.User

	// SELECT * FROM users WHERE id = claims.Issuer
	database.DB.Where("id = ?", claims.Issuer).First(&user)

	return c.JSON(user)
}

func Logout(c *fiber.Ctx) error {
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}
	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "Success",
	})
}

func Status(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "Server up and running!",
	})
}
