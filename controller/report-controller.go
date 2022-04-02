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

type winlose struct {
	Client_start string `json:"client_start" validate:"required"`
	Client_end   string `json:"client_end" validate:"required"`
}
type response_winlose struct {
	Status                 int         `json:"status"`
	Message                string      `json:"message"`
	Record                 interface{} `json:"record"`
	Subtotalturnover       int         `json:"subtotalturnover"`
	Subtotalwinlose        int         `json:"subtotalwinlose"`
	Subtotalwinlosecompany int         `json:"subtotalwinlosecompany"`
}

func Reportwinlose(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(winlose)
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
		SetResult(response_winlose{}).
		SetError(response_winlose{}).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"client_start": client.Client_start,
			"client_end":   client.Client_end,
		}).
		Post(config.Path_url() + "api/reportwinlose")
	if err != nil {
		log.Println(err.Error())
	}
	result := resp.Result().(*response_winlose)
	if result.Status == 200 {
		c.Status(fiber.StatusOK)
		return c.JSON(fiber.Map{
			"status":                 http.StatusOK,
			"message":                result.Message,
			"record":                 result.Record,
			"subtotalturnover":       result.Subtotalturnover,
			"subtotalwinlose":        result.Subtotalwinlose,
			"subtotalwinlosecompany": result.Subtotalwinlosecompany,
			"time":                   time.Since(render_page).String(),
		})
	} else {
		result_error := resp.Error().(*response_winlose)
		c.Status(result_error.Status)
		return c.JSON(fiber.Map{
			"status":  result_error.Status,
			"message": result_error.Message,
			"record":  nil,
			"time":    time.Since(render_page).String(),
		})
	}
}
