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

type home struct {
	Page string `json:"page" validate:"required"`
}
type response_login struct {
	Token string `json:"token"`
	Key   string `json:"key"`
}
type response_loginhome struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func Login(c *fiber.Ctx) error {
	type payload_login struct {
		Username  string `json:"username" validate:"required"`
		Password  string `json:"password" validate:"required"`
		Ipaddress string `json:"ipaddress" validate:"required"`
		Timezone  string `json:"timezone" validate:"required"`
	}
	var errors []*helpers.ErrorResponse
	client := new(payload_login)
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
	axios := resty.New()
	resp, err := axios.R().
		SetResult(response_login{}).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"username":  client.Username,
			"password":  client.Password,
			"ipaddress": client.Ipaddress,
			"timezone":  client.Timezone,
		}).
		Post(config.Path_url() + "api/login")
	if err != nil {
		log.Println(err.Error())
	}
	result := resp.Result().(*response_login)
	c.Status(fiber.StatusOK)
	return c.JSON(fiber.Map{
		"status": http.StatusOK,
		"token":  result.Token,
		"key":    result.Key,
		"time":   time.Since(render_page).String(),
	})
}
func Home(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(home)
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
	bearToken := c.Get("Authorization")
	token := strings.Split(bearToken, " ")
	render_page := time.Now()
	axios := resty.New()
	resp, err := axios.R().
		SetAuthToken(token[1]).
		SetResult(response_loginhome{}).
		SetError(response_loginhome{}).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"page": client.Page,
		}).
		Post(config.Path_url() + "api/home")
	if err != nil {
		log.Println(err.Error())
	}

	result := resp.Result().(*response_loginhome)
	if result.Status == 200 {
		c.Status(fiber.StatusOK)
		return c.JSON(fiber.Map{
			"status":  http.StatusOK,
			"message": result.Message,
			"time":    time.Since(render_page).String(),
		})
	} else {
		result_error := resp.Error().(*response_loginhome)
		c.Status(result_error.Status)
		return c.JSON(fiber.Map{
			"status":  result_error.Status,
			"message": result_error.Message,
			"record":  nil,
			"time":    time.Since(render_page).String(),
		})
	}
}
