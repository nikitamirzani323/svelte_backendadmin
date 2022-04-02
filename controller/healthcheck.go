package controller

import (
	"github.com/gofiber/fiber/v2"
)

func HealthCheck(c *fiber.Ctx) error {

	realip := "0.0.0.0"

	if c.IPs() != nil {
		realip = c.IPs()[0]
	}

	res := map[string]interface{}{
		"status":       fiber.StatusOK,
		"data":         "Server is up and running",
		"container_ip": c.IP(),
		"real_ip":      realip,
	}

	if err := c.JSON(res); err != nil {
		return err
	}

	return nil
}
