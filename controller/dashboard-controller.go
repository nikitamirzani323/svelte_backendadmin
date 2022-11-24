package controller

import (
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/go-resty/resty/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/nikitamirzani323/svelte_backendadmin/config"
	"github.com/nikitamirzani323/svelte_backendadmin/helpers"
)

type response_dashboard struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Record  interface{} `json:"record"`
}

func Dashboardlistpasaran(c *fiber.Ctx) error {
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
		Post(config.Path_url() + "api/allpasarandashboard")
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
func GeneratorSave(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(pasaransaveonline)
	validate := validator.New()
	if err := c.BodyParser(client); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}
	err := validate.Struct(client)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element helpers.ErrorResponse
			element.Field = err.StructField()
			element.Tag = err.Tag()
			errors = append(errors, &element)
		}
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": "validation",
			"record":  errors,
		})
	}
	render_page := time.Now()
	bearToken := c.Get("Authorization")
	token := strings.Split(bearToken, " ")
	axios := resty.New()
	resp, err := axios.R().
		SetAuthToken(token[1]).
		SetResult(response_pasaransave{}).
		SetError(response_pasaransave{}).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"idpasaran":        client.Idpasaran,
			"page":             client.Page,
			"pasaran_hariraya": client.Haripasaran,
		}).
		Post(config.Path_url() + "api/savepasaranonline")
	if err != nil {
		log.Println(err.Error())
	}
	result := resp.Result().(*response_pasaransave)
	if result.Status == 200 {
		c.Status(fiber.StatusOK)
		return c.JSON(fiber.Map{
			"status":  http.StatusOK,
			"message": result.Message,
			"record":  result.Record,
			"time":    time.Since(render_page).String(),
		})
	} else {
		result_error := resp.Error().(*response_pasaransave)
		c.Status(result_error.Status)
		return c.JSON(fiber.Map{
			"status":  result_error.Status,
			"message": result_error.Message,
			"record":  nil,
			"time":    time.Since(render_page).String(),
		})
	}
}
