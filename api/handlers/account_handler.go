package handlers

import (
	"role-based/config/env"
	"role-based/models"
	"role-based/services"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

// Services Injection
type AccountHandlers struct {
	service services.AccountServices
}

// To Initialize this Handlers
func AccountHandlersInit(service services.AccountServices) *AccountHandlers {
	return &AccountHandlers{service}
}

func (hh *AccountHandlers) CreateAccount(h *fiber.Ctx) error {
	var userAccount models.Account

	if err := h.BodyParser(&userAccount); err != nil {
		return h.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}
	err := hh.service.CreateAccountService(userAccount)
	if err != nil {
		return err
	}
	return h.SendStatus(fiber.StatusAccepted)
}

func (hh *AccountHandlers) AccountLogin(h *fiber.Ctx) error {

	var userLogin models.LoginCred

	err := h.BodyParser(&userLogin)
	if err != nil {
		return h.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	userAccount, mess := hh.service.AccountLoginService(userLogin)

	switch mess {
	case "User Does not exist exist":
		return h.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "User not exist",
		})

	case "Error in Database":
		return h.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error in Database",
		})

	case "Password does not match":
		return h.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Password does not match",
		})

	case "Account match":
	}
	// Create the Claims
	claims := jwt.MapClaims{
		"id":    userAccount.ID,
		"name":  userAccount.Username,
		"role": userAccount.Role,
		"exp":   time.Now().Add(time.Minute * 60).Unix(),
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(env.Config("COOKIES_SECRET_KEY")))
	if err != nil {
		return h.SendStatus(fiber.StatusInternalServerError)
	}

	//Creating Cookies struct may other way setcookie
	h.Cookie(&fiber.Cookie{
		Name:     "token",
		Value:    t,
		Expires:  time.Now().Add(60 * time.Minute),
		HTTPOnly: true,
		Secure:   false,
		SameSite: "Lax",
	})

	return h.Status(fiber.StatusOK).JSON(fiber.Map{
		"alert": "succesfull login",
	})
}

//Get Profile
func (hh *AccountHandlers) GetProfile (h *fiber.Ctx) error {

	//Get the id from claims that store in locals
	id := h.Locals("id")
	if id == nil {
		return h.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}
	//convert to float and to int before passing
	idFloat, ok := id.(float64)
	if !ok {
		return h.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Invalid user ID type"})
	}

	profile,err := hh.service.GetProfileService(int(idFloat))

	if err != nil {
		return h.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return h.Status(fiber.StatusOK).JSON(profile)
}
