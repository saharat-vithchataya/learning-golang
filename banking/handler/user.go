package handler

import (
	"banking/services"

	"github.com/gofiber/fiber/v2"
)

type userHandlerFiber struct {
	userSrv services.UserService
}

func NewUserHandlerFiber(userSrv services.UserService) userHandlerFiber {
	return userHandlerFiber{userSrv: userSrv}
}

func (h userHandlerFiber) SignUp(c *fiber.Ctx) error {
	request := services.SignUpRequest{}
	err := c.BodyParser(&request)
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).SendString(err.Error())
	}
	response, err := h.userSrv.SignUp(request)
	if err != nil {
		return err
	}
	return c.Status(fiber.StatusCreated).JSON(response)
}

func (h userHandlerFiber) Login(c *fiber.Ctx) error {
	request := services.LoginRequest{}
	err := c.BodyParser(&request)
	if err != nil {
		return err
	}

	response, err := h.userSrv.Login(request)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusCreated).JSON(response)
}

func (h userHandlerFiber) WhoAmI(c *fiber.Ctx) error {

	// user, ok := c.Locals("user").(*jwtware.)
	// jwtware.KeyRefreshSuccessHandler()
	// if !ok {
	// 	return fiber.ErrUnauthorized
	// }

	// // Now you can access user information from the claims
	// userId := user.Claims["issuer"].(string)
	// // Add your logic to fetch user details from the database using the userId

	// // Example response
	// return c.JSON(fiber.Map{
	// 	"userId": userId,
	// })
	return nil
}
