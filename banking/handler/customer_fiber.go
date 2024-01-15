package handler

import (
	"banking/services"

	"github.com/gofiber/fiber/v2"
)

type customerHandlerFiber struct {
	custSrv services.CustomerService
}

func NewCustomerHandlerFiber(cusSrv services.CustomerService) customerHandlerFiber {
	return customerHandlerFiber{custSrv: cusSrv}
}

func (h customerHandlerFiber) GetCustomers(c *fiber.Ctx) error {
	customers, err := h.custSrv.GetCustomers()
	if err != nil {
		return err
	}
	return c.Status(fiber.StatusOK).JSON(customers)
}

func (h customerHandlerFiber) GetCustomer(c *fiber.Ctx) error {
	id, err := c.ParamsInt("customer_id")
	if err != nil {
		return err
	}

	customer, err := h.custSrv.GetCustomer(id)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(customer)
}
