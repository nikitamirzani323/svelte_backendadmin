package controller

import (
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/nikitamirzani323/svelte_backendadmin/config"
)

type response_dashboard struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Record  interface{} `json:"record"`
}

func Dashboardwinlose(c *fiber.Ctx) error {
	type payload_dashboardwinlose struct {
		Year string `json:"year"`
	}
	render_page := time.Now()
	bearToken := c.Get("Authorization")
	token := strings.Split(bearToken, " ")
	client := new(payload_dashboardwinlose)
	if err := c.BodyParser(client); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}
	axios := resty.New()
	resp, err := axios.R().
		SetAuthToken(token[1]).
		SetResult(response_dashboard{}).
		SetError(response_dashboard{}).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"year": client.Year,
		}).
		Post(config.Path_url() + "api/dashboardwinlose")
	if err != nil {
		log.Println(err.Error())
	}

	result := resp.Result().(*response_dashboard)
	if result.Status == 200 {
		c.Status(fiber.StatusOK)
		return c.JSON(fiber.Map{
			"status":  http.StatusOK,
			"message": result.Message,
			"record":  result.Record,
			"time":    time.Since(render_page).String(),
		})
	} else {
		result_error := resp.Error().(*response_dashboard)
		c.Status(result_error.Status)
		return c.JSON(fiber.Map{
			"status":  result_error.Status,
			"message": result_error.Message,
			"record":  nil,
			"time":    time.Since(render_page).String(),
		})
	}
}
func Dashboardwinlosepasaran(c *fiber.Ctx) error {
	type payload_dashboardwinlosepasaran struct {
		Year string `json:"year"`
	}
	render_page := time.Now()
	bearToken := c.Get("Authorization")
	token := strings.Split(bearToken, " ")
	client := new(payload_dashboardwinlosepasaran)
	if err := c.BodyParser(client); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}
	axios := resty.New()
	resp, err := axios.R().
		SetAuthToken(token[1]).
		SetResult(response_dashboard{}).
		SetError(response_dashboard{}).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"year": client.Year,
		}).
		Post(config.Path_url() + "api/dashboardpasaranwinlose")
	if err != nil {
		log.Println(err.Error())
	}

	result := resp.Result().(*response_dashboard)
	if result.Status == 200 {
		c.Status(fiber.StatusOK)
		return c.JSON(fiber.Map{
			"status":  http.StatusOK,
			"message": result.Message,
			"record":  result.Record,
			"time":    time.Since(render_page).String(),
		})
	} else {
		result_error := resp.Error().(*response_dashboard)
		c.Status(result_error.Status)
		return c.JSON(fiber.Map{
			"status":  result_error.Status,
			"message": result_error.Message,
			"record":  nil,
			"time":    time.Since(render_page).String(),
		})
	}
}
