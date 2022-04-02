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

type periodedetail struct {
	Idtrxkeluaran int `json:"idinvoice"`
}
type periodedetailbettable struct {
	Idtrxkeluaran int    `json:"idinvoice"`
	Permainan     string `json:"permainan"`
}
type periodedetailmembernomor struct {
	Idtrxkeluaran int    `json:"idinvoice" validate:"required"`
	Permainan     string `json:"permainan" validate:"required"`
	Nomor         string `json:"nomor" validate:"required"`
}
type periodelistbet struct {
	Idtrxkeluaran int    `json:"idinvoice"`
	Permainan     string `json:"permainan"`
}
type periodelistbetstatus struct {
	Idtrxkeluaran int    `json:"idinvoice"`
	Status        string `json:"status"`
}
type periodelistbetusername struct {
	Idtrxkeluaran int    `json:"idinvoice"`
	Username      string `json:"username"`
}
type periodeSave struct {
	Sdata          string `json:"sData" validate:"required"`
	Page           string `json:"page"`
	Idtrxkeluaran  int    `json:"idinvoice" validate:"required"`
	Nomorkeluaran  string `json:"nomorkeluaran" validate:"required,min=4,max=4"`
	Idpasarantogel string `json:"idpasarancode" validate:"required"`
}
type periodesavenew struct {
	Sdata         string `json:"sData" validate:"required"`
	Page          string `json:"page"`
	Idcomppasaran int    `json:"pasaran_code" validate:"required"`
}
type periodesaverevisi struct {
	Sdata         string `json:"sData" validate:"required"`
	Page          string `json:"page"`
	Idtrxkeluaran int    `json:"idinvoice" validate:"required"`
	Msgrevisi     string `json:"msgrevisi" validate:"required"`
}

type periodeprediksi struct {
	Nomorkeluaran string `json:"nomorkeluaran" validate:"required,min=4,max=4"`
	Idcomppasaran int    `json:"pasaran_code" validate:"required"`
}
type response_periode struct {
	Status        int         `json:"status"`
	Message       string      `json:"message"`
	Record        interface{} `json:"record"`
	Pasaranonline interface{} `json:"pasaranonline"`
}
type response_periodedetail struct {
	Status      int         `json:"status"`
	Message     string      `json:"message"`
	Record      interface{} `json:"record"`
	Totalbet    int         `json:"totalbet"`
	Subtotal    int         `json:"subtotal"`
	Subtotalwin int         `json:"subtotalwin"`
}
type response_periodelistbet struct {
	Status      int         `json:"status"`
	Message     string      `json:"message"`
	Record      interface{} `json:"record"`
	Totalbet    int         `json:"totalbet"`
	Subtotal    int         `json:"subtotal"`
	Subtotalwin int         `json:"subtotalwin"`
}

func Periode(c *fiber.Ctx) error {
	render_page := time.Now()
	bearToken := c.Get("Authorization")
	token := strings.Split(bearToken, " ")
	axios := resty.New()
	resp, err := axios.R().
		SetAuthToken(token[1]).
		SetResult(response_periode{}).
		SetError(response_periode{}).
		SetHeader("Content-Type", "application/json").
		Post(config.Path_url() + "api/allperiode")
	if err != nil {
		log.Println(err.Error())
	}
	result := resp.Result().(*response_periode)
	if result.Status == 200 {
		c.Status(fiber.StatusOK)
		return c.JSON(fiber.Map{
			"status":        http.StatusOK,
			"message":       result.Message,
			"record":        result.Record,
			"pasaranonline": result.Pasaranonline,
			"time":          time.Since(render_page).String(),
		})
	} else {
		result_error := resp.Error().(*response_periode)
		c.Status(result_error.Status)
		return c.JSON(fiber.Map{
			"status":  result_error.Status,
			"message": result_error.Message,
			"record":  nil,
			"time":    time.Since(render_page).String(),
		})
	}
}
func Periodedetail(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(periodedetail)
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
		SetResult(response_periodedetail{}).
		SetError(response_periodedetail{}).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"idinvoice": client.Idtrxkeluaran,
		}).
		Post(config.Path_url() + "api/editperiode")
	if err != nil {
		log.Println(err.Error())
	}
	result := resp.Result().(*response_periodedetail)
	if result.Status == 200 {
		c.Status(fiber.StatusOK)
		return c.JSON(fiber.Map{
			"status":  http.StatusOK,
			"message": result.Message,
			"record":  result.Record,
			"time":    time.Since(render_page).String(),
		})
	} else {
		result_error := resp.Error().(*response_periodedetail)
		c.Status(result_error.Status)
		return c.JSON(fiber.Map{
			"status":  result_error.Status,
			"message": result_error.Message,
			"record":  nil,
			"time":    time.Since(render_page).String(),
		})
	}
}
func Periodesave(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(periodeSave)
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
		SetResult(response_periodedetail{}).
		SetError(response_periodedetail{}).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"sData":         client.Sdata,
			"idinvoice":     client.Idtrxkeluaran,
			"nomorkeluaran": client.Nomorkeluaran,
			"idpasarancode": client.Idpasarantogel,
		}).
		Post(config.Path_url() + "api/saveperiode")
	if err != nil {
		log.Println(err.Error())
	}

	result := resp.Result().(*response_periodedetail)
	if result.Status == 200 {
		c.Status(fiber.StatusOK)
		return c.JSON(fiber.Map{
			"status":  http.StatusOK,
			"message": result.Message,
			"record":  result.Record,
			"time":    time.Since(render_page).String(),
		})
	} else {
		result_error := resp.Error().(*response_periodedetail)
		c.Status(result_error.Status)
		return c.JSON(fiber.Map{
			"status":  result_error.Status,
			"message": result_error.Message,
			"record":  nil,
			"time":    time.Since(render_page).String(),
		})
	}
}
func Periodesavenew(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(periodesavenew)
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
		SetResult(response_periodedetail{}).
		SetError(response_periodedetail{}).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"sData":        client.Sdata,
			"page":         client.Page,
			"pasaran_code": client.Idcomppasaran,
		}).
		Post(config.Path_url() + "api/savepasarannew")
	if err != nil {
		log.Println(err.Error())
	}

	result := resp.Result().(*response_periodedetail)
	if result.Status == 200 {
		c.Status(fiber.StatusOK)
		return c.JSON(fiber.Map{
			"status":  http.StatusOK,
			"message": result.Message,
			"record":  result.Record,
			"time":    time.Since(render_page).String(),
		})
	} else {
		result_error := resp.Error().(*response_periodedetail)
		c.Status(result_error.Status)
		return c.JSON(fiber.Map{
			"status":  result_error.Status,
			"message": result_error.Message,
			"record":  nil,
			"time":    time.Since(render_page).String(),
		})
	}
}
func Periodesaverevisi(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(periodesaverevisi)
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
		SetResult(response_periodedetail{}).
		SetError(response_periodedetail{}).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"sData":     client.Sdata,
			"idinvoice": client.Idtrxkeluaran,
			"page":      client.Page,
			"msgrevisi": client.Msgrevisi,
		}).
		Post(config.Path_url() + "api/saveperioderevisi")
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
	result := resp.Result().(*response_periodedetail)
	if result.Status == 200 {
		c.Status(fiber.StatusOK)
		return c.JSON(fiber.Map{
			"status":  http.StatusOK,
			"message": result.Message,
			"record":  result.Record,
			"time":    time.Since(render_page).String(),
		})
	} else {
		result_error := resp.Error().(*response_periodedetail)
		c.Status(result_error.Status)
		return c.JSON(fiber.Map{
			"status":  result_error.Status,
			"message": result_error.Message,
			"time":    time.Since(render_page).String(),
		})
	}
}
func Periodecancelbet(c *fiber.Ctx) error {
	type payload_periodecancelbet struct {
		Sdata               string `json:"sData" validate:"required"`
		Page                string `json:"page"`
		Permainan           string `json:"permainan" validate:"required"`
		Idtrxkeluaran       int    `json:"idinvoice" validate:"required"`
		Idtrxkeluarandetail int    `json:"idinvoicedetail" validate:"required"`
	}
	var errors []*helpers.ErrorResponse
	client := new(payload_periodecancelbet)
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
		SetResult(response_periodedetail{}).
		SetError(response_periodedetail{}).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"sData":           client.Sdata,
			"page":            client.Page,
			"permainan":       client.Permainan,
			"idinvoice":       client.Idtrxkeluaran,
			"idinvoicedetail": client.Idtrxkeluarandetail,
		}).
		Post(config.Path_url() + "api/cancelbet")
	if err != nil {
		log.Println(err.Error())
	}

	result := resp.Result().(*response_periodedetail)
	if result.Status == 200 {
		c.Status(fiber.StatusOK)
		return c.JSON(fiber.Map{
			"status":  http.StatusOK,
			"message": result.Message,
			"record":  result.Record,
			"time":    time.Since(render_page).String(),
		})
	} else {
		result_error := resp.Error().(*response_periodedetail)
		c.Status(result_error.Status)
		return c.JSON(fiber.Map{
			"status":  result_error.Status,
			"message": result_error.Message,
			"record":  nil,
			"time":    time.Since(render_page).String(),
		})
	}
}
func Periodelistmember(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(periodedetail)
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
		SetResult(response_periodedetail{}).
		SetError(response_periodedetail{}).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"idinvoice": client.Idtrxkeluaran,
		}).
		Post(config.Path_url() + "api/periodelistmember")
	if err != nil {
		log.Println(err.Error())
	}
	result := resp.Result().(*response_periodedetail)
	if result.Status == 200 {
		c.Status(fiber.StatusOK)
		return c.JSON(fiber.Map{
			"status":  http.StatusOK,
			"message": result.Message,
			"record":  result.Record,
			"time":    time.Since(render_page).String(),
		})
	} else {
		result_error := resp.Error().(*response_periodedetail)
		c.Status(result_error.Status)
		return c.JSON(fiber.Map{
			"status":  result_error.Status,
			"message": result_error.Message,
			"record":  nil,
			"time":    time.Since(render_page).String(),
		})
	}
}
func Periodelistmemberbynomor(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(periodedetailmembernomor)
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
		SetResult(response_periodedetail{}).
		SetError(response_periodedetail{}).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"idinvoice": client.Idtrxkeluaran,
			"permainan": client.Permainan,
			"nomor":     client.Nomor,
		}).
		Post(config.Path_url() + "api/periodelistmemberbynomor")
	if err != nil {
		log.Println(err.Error())
	}
	result := resp.Result().(*response_periodedetail)
	if result.Status == 200 {
		c.Status(fiber.StatusOK)
		return c.JSON(fiber.Map{
			"status":  http.StatusOK,
			"message": result.Message,
			"record":  result.Record,
			"time":    time.Since(render_page).String(),
		})
	} else {
		result_error := resp.Error().(*response_periodedetail)
		c.Status(result_error.Status)
		return c.JSON(fiber.Map{
			"status":  result_error.Status,
			"message": result_error.Message,
			"record":  nil,
			"time":    time.Since(render_page).String(),
		})
	}
}
func Periodelistbet(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(periodelistbet)
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
		SetResult(response_periodelistbet{}).
		SetError(response_periodelistbet{}).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"idinvoice": client.Idtrxkeluaran,
			"permainan": client.Permainan,
		}).
		Post(config.Path_url() + "api/periodelistbet")
	if err != nil {
		log.Println(err.Error())
	}
	result := resp.Result().(*response_periodelistbet)
	if result.Status == 200 {
		c.Status(fiber.StatusOK)
		return c.JSON(fiber.Map{
			"status":      http.StatusOK,
			"message":     result.Message,
			"record":      result.Record,
			"totalbet":    result.Totalbet,
			"subtotal":    result.Subtotal,
			"subtotalwin": result.Subtotalwin,
			"time":        time.Since(render_page).String(),
		})
	} else {
		result_error := resp.Error().(*response_periodelistbet)
		c.Status(result_error.Status)
		return c.JSON(fiber.Map{
			"status":  result_error.Status,
			"message": result_error.Message,
			"record":  nil,
			"time":    time.Since(render_page).String(),
		})
	}
}
func Periodelistbetstatus(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(periodelistbetstatus)
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
		SetResult(response_periodelistbet{}).
		SetError(response_periodelistbet{}).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"idinvoice": client.Idtrxkeluaran,
			"status":    client.Status,
		}).
		Post(config.Path_url() + "api/periodelistbetstatus")
	if err != nil {
		log.Println(err.Error())
	}
	result := resp.Result().(*response_periodelistbet)
	if result.Status == 200 {
		c.Status(fiber.StatusOK)
		return c.JSON(fiber.Map{
			"status":      http.StatusOK,
			"message":     result.Message,
			"record":      result.Record,
			"totalbet":    result.Totalbet,
			"subtotal":    result.Subtotal,
			"subtotalwin": result.Subtotalwin,
			"time":        time.Since(render_page).String(),
		})
	} else {
		result_error := resp.Error().(*response_periodelistbet)
		c.Status(result_error.Status)
		return c.JSON(fiber.Map{
			"status":  result_error.Status,
			"message": result_error.Message,
			"record":  nil,
			"time":    time.Since(render_page).String(),
		})
	}
}
func Periodelistbetbyusername(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(periodelistbetusername)
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
		SetResult(response_periodelistbet{}).
		SetError(response_periodelistbet{}).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"idinvoice": client.Idtrxkeluaran,
			"username":  client.Username,
		}).
		Post(config.Path_url() + "api/periodelistbetusername")
	if err != nil {
		log.Println(err.Error())
	}
	result := resp.Result().(*response_periodelistbet)
	if result.Status == 200 {
		c.Status(fiber.StatusOK)
		return c.JSON(fiber.Map{
			"status":      http.StatusOK,
			"message":     result.Message,
			"record":      result.Record,
			"totalbet":    result.Totalbet,
			"subtotal":    result.Subtotal,
			"subtotalwin": result.Subtotalwin,
			"time":        time.Since(render_page).String(),
		})
	} else {
		result_error := resp.Error().(*response_periodelistbet)
		c.Status(result_error.Status)
		return c.JSON(fiber.Map{
			"status":  result_error.Status,
			"message": result_error.Message,
			"record":  nil,
			"time":    time.Since(render_page).String(),
		})
	}
}
func Periodelistbettable(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(periodedetail)
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
		SetResult(response_periodedetail{}).
		SetError(response_periodedetail{}).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"idinvoice": client.Idtrxkeluaran,
		}).
		Post(config.Path_url() + "api/periodelistbettable")
	if err != nil {
		log.Println(err.Error())
	}
	result := resp.Result().(*response_periodedetail)
	if result.Status == 200 {
		c.Status(fiber.StatusOK)
		return c.JSON(fiber.Map{
			"status":  http.StatusOK,
			"message": result.Message,
			"record":  result.Record,
			"time":    time.Since(render_page).String(),
		})
	} else {
		result_error := resp.Error().(*response_periodedetail)
		c.Status(result_error.Status)
		return c.JSON(fiber.Map{
			"status":  result_error.Status,
			"message": result_error.Message,
			"record":  nil,
			"time":    time.Since(render_page).String(),
		})
	}
}
func Periodebettable(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(periodedetailbettable)
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
		SetResult(response_periodedetail{}).
		SetError(response_periodedetail{}).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"idinvoice": client.Idtrxkeluaran,
			"permainan": client.Permainan,
		}).
		Post(config.Path_url() + "api/periodebettable")
	if err != nil {
		log.Println(err.Error())
	}
	result := resp.Result().(*response_periodedetail)
	if result.Status == 200 {
		c.Status(fiber.StatusOK)
		return c.JSON(fiber.Map{
			"status":  http.StatusOK,
			"message": result.Message,
			"record":  result.Record,
			"time":    time.Since(render_page).String(),
		})
	} else {
		result_error := resp.Error().(*response_periodedetail)
		c.Status(result_error.Status)
		return c.JSON(fiber.Map{
			"status":  result_error.Status,
			"message": result_error.Message,
			"record":  nil,
			"time":    time.Since(render_page).String(),
		})
	}
}
func Periodelistpasaran(c *fiber.Ctx) error {
	render_page := time.Now()
	bearToken := c.Get("Authorization")
	token := strings.Split(bearToken, " ")
	axios := resty.New()
	resp, err := axios.R().
		SetAuthToken(token[1]).
		SetResult(response_periode{}).
		SetError(response_periode{}).
		SetHeader("Content-Type", "application/json").
		Post(config.Path_url() + "api/listpasaran")
	if err != nil {
		log.Println(err.Error())
	}
	result := resp.Result().(*response_periode)
	if result.Status == 200 {
		c.Status(fiber.StatusOK)
		return c.JSON(fiber.Map{
			"status":  http.StatusOK,
			"message": result.Message,
			"record":  result.Record,
			"time":    time.Since(render_page).String(),
		})
	} else {
		result_error := resp.Error().(*response_periode)
		c.Status(result_error.Status)
		return c.JSON(fiber.Map{
			"status":  result_error.Status,
			"message": result_error.Message,
			"record":  nil,
			"time":    time.Since(render_page).String(),
		})
	}
}
func Periodeprediksi(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(periodeprediksi)
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
		SetResult(response_periodedetail{}).
		SetError(response_periodedetail{}).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"nomorkeluaran": client.Nomorkeluaran,
			"Idcomppasaran": client.Idcomppasaran,
		}).
		Post(config.Path_url() + "api/listprediksi")
	if err != nil {
		log.Println(err.Error())
	}
	result := resp.Result().(*response_periodedetail)
	if result.Status == 200 {
		c.Status(fiber.StatusOK)
		return c.JSON(fiber.Map{
			"status":      http.StatusOK,
			"message":     result.Message,
			"record":      result.Record,
			"totalbet":    result.Totalbet,
			"subtotal":    result.Subtotal,
			"subtotalwin": result.Subtotalwin,
			"time":        time.Since(render_page).String(),
		})
	} else {
		result_error := resp.Error().(*response_periodedetail)
		c.Status(result_error.Status)
		return c.JSON(fiber.Map{
			"status":      result_error.Status,
			"message":     result_error.Message,
			"record":      nil,
			"totalbet":    0,
			"subtotal":    0,
			"subtotalwin": 0,
			"time":        time.Since(render_page).String(),
		})
	}
}
