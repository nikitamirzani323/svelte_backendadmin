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

type admindetail struct {
	Username string `json:"username" validate:"required"`
}

type adminsaveiplist struct {
	Sdata     string `json:"sdata" validate:"required"`
	Page      string `json:"page"`
	Username  string `json:"username" validate:"required,alphanum,max=20"`
	Ipaddress string `json:"ipaddress" validate:"required,max=20"`
}
type response_admin struct {
	Status        int         `json:"status"`
	Message       string      `json:"message"`
	Record        interface{} `json:"record"`
	Listruleadmin interface{} `json:"listruleadmin"`
	Listip        interface{} `json:"listip"`
}
type response_adminsave struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Record  interface{} `json:"record"`
}

func Admin(c *fiber.Ctx) error {
	render_page := time.Now()
	bearToken := c.Get("Authorization")
	token := strings.Split(bearToken, " ")
	axios := resty.New()
	resp, err := axios.R().
		SetAuthToken(token[1]).
		SetResult(response_admin{}).
		SetError(response_admin{}).
		SetHeader("Content-Type", "application/json").
		Post(config.Path_url() + "api/alladmin")
	if err != nil {
		log.Println(err.Error())
	}
	result := resp.Result().(*response_admin)
	if result.Status == 200 {
		c.Status(fiber.StatusOK)
		return c.JSON(fiber.Map{
			"status":        http.StatusOK,
			"message":       result.Message,
			"record":        result.Record,
			"listruleadmin": result.Listruleadmin,
			"time":          time.Since(render_page).String(),
		})
	} else {
		result_error := resp.Error().(*response_admin)
		c.Status(result_error.Status)
		return c.JSON(fiber.Map{
			"status":  result_error.Status,
			"message": result_error.Message,
			"record":  nil,
			"time":    time.Since(render_page).String(),
		})
	}
}
func Admindetail(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(admindetail)
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
		SetResult(response_admin{}).
		SetError(response_admin{}).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"username": client.Username,
		}).
		Post(config.Path_url() + "api/editadmin")
	if err != nil {
		log.Println(err.Error())
	}
	result := resp.Result().(*response_admin)
	if result.Status == 200 {
		c.Status(fiber.StatusOK)
		return c.JSON(fiber.Map{
			"status":        http.StatusOK,
			"message":       result.Message,
			"record":        result.Record,
			"listip":        result.Listip,
			"listruleadmin": result.Listruleadmin,
			"time":          time.Since(render_page).String(),
		})
	} else {
		result_error := resp.Error().(*response_admin)
		c.Status(result_error.Status)
		return c.JSON(fiber.Map{
			"status":  result_error.Status,
			"message": result_error.Message,
			"record":  nil,
			"time":    time.Since(render_page).String(),
		})
	}
}
func Adminsave(c *fiber.Ctx) error {
	type payload_adminsave struct {
		Sdata       string `json:"sdata" validate:"required"`
		Page        string `json:"page"`
		Idruleadmin int    `json:"idruleadmin" `
		Username    string `json:"username" validate:"required,alphanum,max=20"`
		Password    string `json:"password" `
		Name        string `json:"nama" validate:"required,alphanum,max=70"`
		Status      string `json:"status" validate:"required,alpha"`
	}
	var errors []*helpers.ErrorResponse
	client := new(payload_adminsave)
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
		SetResult(response_adminsave{}).
		SetError(response_adminsave{}).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"sdata":       client.Sdata,
			"page":        client.Page,
			"idruleadmin": client.Idruleadmin,
			"username":    client.Username,
			"password":    client.Password,
			"nama":        client.Name,
			"status":      client.Status,
		}).
		Post(config.Path_url() + "api/saveadmin")
	if err != nil {
		log.Println(err.Error())
	}
	log.Println("Response Info:")
	log.Println("  Error      :", err)
	log.Println("  Status Code:", resp.StatusCode())
	log.Println("  Status     :", resp.Status())
	log.Println("  Proto      :", resp.Proto())
	log.Println("  Time       :", resp.Time())
	log.Println("  Received At:", resp.ReceivedAt())
	log.Println("  Body       :\n", resp)
	log.Println()
	result := resp.Result().(*response_adminsave)
	if result.Status == 200 {
		c.Status(fiber.StatusOK)
		return c.JSON(fiber.Map{
			"status":  http.StatusOK,
			"message": result.Message,
			"record":  result.Record,
			"time":    time.Since(render_page).String(),
		})
	} else {
		result_error := resp.Error().(*response_adminsave)
		c.Status(result_error.Status)
		return c.JSON(fiber.Map{
			"status":  result_error.Status,
			"message": result_error.Message,
			"record":  nil,
			"time":    time.Since(render_page).String(),
		})
	}
}
func Adminsaveiplist(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(adminsaveiplist)
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
		SetResult(response_adminsave{}).
		SetError(response_adminsave{}).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"sdata":     client.Sdata,
			"page":      client.Page,
			"username":  client.Username,
			"ipaddress": client.Ipaddress,
		}).
		Post(config.Path_url() + "api/saveadminiplist")
	if err != nil {
		log.Println(err.Error())
	}
	result := resp.Result().(*response_adminsave)
	if result.Status == 200 {
		c.Status(fiber.StatusOK)
		return c.JSON(fiber.Map{
			"status":  http.StatusOK,
			"message": result.Message,
			"record":  result.Record,
			"time":    time.Since(render_page).String(),
		})
	} else {
		result_error := resp.Error().(*response_adminsave)
		c.Status(result_error.Status)
		return c.JSON(fiber.Map{
			"status":  result_error.Status,
			"message": result_error.Message,
			"record":  nil,
			"time":    time.Since(render_page).String(),
		})
	}
}
func Deleteiplist(c *fiber.Ctx) error {
	type payload_deleteiplist struct {
		Page         string `json:"page"`
		Username     string `json:"username" validate:"required,alphanum,max=20"`
		Idcompiplist int    `json:"idcompiplist" validate:"required"`
	}

	var errors []*helpers.ErrorResponse
	client := new(payload_deleteiplist)
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
		SetResult(response_adminsave{}).
		SetError(response_adminsave{}).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"page":         client.Page,
			"username":     client.Username,
			"idcompiplist": client.Idcompiplist,
		}).
		Post(config.Path_url() + "api/deleteadminiplist")
	if err != nil {
		log.Println(err.Error())
	}
	result := resp.Result().(*response_adminsave)
	if result.Status == 200 {
		c.Status(fiber.StatusOK)
		return c.JSON(fiber.Map{
			"status":  http.StatusOK,
			"message": result.Message,
			"record":  result.Record,
			"time":    time.Since(render_page).String(),
		})
	} else {
		result_error := resp.Error().(*response_adminsave)
		c.Status(result_error.Status)
		return c.JSON(fiber.Map{
			"status":  result_error.Status,
			"message": result_error.Message,
			"record":  nil,
			"time":    time.Since(render_page).String(),
		})
	}
}
