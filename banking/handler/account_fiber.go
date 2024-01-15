package handler

import (
	"banking/services"

	"github.com/gofiber/fiber/v2"
)

type accountHandlerFiber struct {
	accSrv services.AccountService
}

func NewAccountHandlerFiber(accSrv services.AccountService) accountHandlerFiber {
	return accountHandlerFiber{accSrv: accSrv}
}

func (h accountHandlerFiber) NewAccount(c *fiber.Ctx) error {
	customerID, err := c.ParamsInt("customer_id")
	if err != nil {
		return err
	}
	request := services.NewAccountRequest{}
	err = c.BodyParser(&request)
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"message": err,
		})
	}
	response, err := h.accSrv.NewAccount(customerID, request)
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"message": err,
		})
	}
	return c.Status(fiber.StatusCreated).JSON(response)
}

func (h accountHandlerFiber) GetAccounts(c *fiber.Ctx) error {
	customerID, err := c.ParamsInt("customer_id")
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"message": err,
		})
	}
	responses, err := h.accSrv.GetAccounts(customerID)
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"message": err,
		})
	}
	return c.Status(fiber.StatusOK).JSON(responses)

}
