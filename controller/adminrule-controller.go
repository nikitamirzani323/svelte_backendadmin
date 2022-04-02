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

type adminruledetail struct {
	Idrule int `json:"idrule" validate:"required"`
}
type adminrulesavedetail struct {
	Idrule int    `json:"idrule" validate:"required"`
	Page   string `json:"page"`
	Sdata  string `json:"sData" validate:"required"`
	Nama   string `json:"nama" validate:"required"`
}
type adminrulesaveconf struct {
	Idrule int    `json:"idrule" validate:"required"`
	Page   string `json:"page"`
	Sdata  string `json:"sData" validate:"required"`
	Rule   string `json:"rule" validate:"required"`
}
type response_adminrule struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Record  interface{} `json:"record"`
}

func Adminrule(c *fiber.Ctx) error {
	render_page := time.Now()
	bearToken := c.Get("Authorization")
	token := strings.Split(bearToken, " ")
	axios := resty.New()
	resp, err := axios.R().
		SetAuthToken(token[1]).
		SetResult(response_adminrule{}).
		SetError(response_adminrule{}).
		SetHeader("Content-Type", "application/json").
		Post(config.Path_url() + "api/alladminrule")
	if err != nil {
		log.Println(err.Error())
	}
	result := resp.Result().(*response_adminrule)
	if result.Status == 200 {
		c.Status(fiber.StatusOK)
		return c.JSON(fiber.Map{
			"status":  http.StatusOK,
			"message": result.Message,
			"record":  result.Record,
			"time":    time.Since(render_page).String(),
		})
	} else {
		result_error := resp.Error().(*response_adminrule)
		c.Status(result_error.Status)
		return c.JSON(fiber.Map{
			"status":  result_error.Status,
			"message": result_error.Message,
			"record":  nil,
			"time":    time.Since(render_page).String(),
		})
	}
}
func Adminruledetail(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(adminruledetail)
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
		SetResult(response_adminrule{}).
		SetError(response_adminrule{}).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"idrule": client.Idrule,
		}).
		Post(config.Path_url() + "api/editadminrule")
	if err != nil {
		log.Println(err.Error())
	}
	result := resp.Result().(*response_adminrule)
	if result.Status == 200 {
		c.Status(fiber.StatusOK)
		return c.JSON(fiber.Map{
			"status":  http.StatusOK,
			"message": result.Message,
			"record":  result.Record,
			"time":    time.Since(render_page).String(),
		})
	} else {
		result_error := resp.Error().(*response_adminrule)
		c.Status(result_error.Status)
		return c.JSON(fiber.Map{
			"status":  result_error.Status,
			"message": result_error.Message,
			"record":  nil,
			"time":    time.Since(render_page).String(),
		})
	}
}
func Adminrulesave(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(adminrulesavedetail)
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
		SetResult(response_adminrule{}).
		SetError(response_adminrule{}).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"idrule": client.Idrule,
			"page":   client.Page,
			"sData":  client.Sdata,
			"nama":   client.Nama,
		}).
		Post(config.Path_url() + "api/saveadminrule")
	if err != nil {
		log.Println(err.Error())
	}
	result := resp.Result().(*response_adminrule)
	if result.Status == 200 {
		c.Status(fiber.StatusOK)
		return c.JSON(fiber.Map{
			"status":  http.StatusOK,
			"message": result.Message,
			"record":  result.Record,
			"time":    time.Since(render_page).String(),
		})
	} else {
		result_error := resp.Error().(*response_adminrule)
		c.Status(result_error.Status)
		return c.JSON(fiber.Map{
			"status":  result_error.Status,
			"message": result_error.Message,
			"record":  nil,
			"time":    time.Since(render_page).String(),
		})
	}
}
func Adminruleconfsave(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(adminrulesaveconf)
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
		SetResult(response_adminrule{}).
		SetError(response_adminrule{}).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"idrule": client.Idrule,
			"page":   client.Page,
			"sData":  client.Sdata,
			"rule":   client.Rule,
		}).
		Post(config.Path_url() + "api/saveadminruleconf")
	if err != nil {
		log.Println(err.Error())
	}
	result := resp.Result().(*response_adminrule)
	if result.Status == 200 {
		c.Status(fiber.StatusOK)
		return c.JSON(fiber.Map{
			"status":  http.StatusOK,
			"message": result.Message,
			"record":  result.Record,
			"time":    time.Since(render_page).String(),
		})
	} else {
		result_error := resp.Error().(*response_adminrule)
		c.Status(result_error.Status)
		return c.JSON(fiber.Map{
			"status":  result_error.Status,
			"message": result_error.Message,
			"record":  nil,
			"time":    time.Since(render_page).String(),
		})
	}
}
