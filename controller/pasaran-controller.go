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

type pasarandetail struct {
	Idpasaran int `json:"idpasaran"`
}
type pasaransave struct {
	Idpasaran         int    `json:"idpasaran"`
	Page              string `json:"page"`
	Pasaran_diundi    string `json:"pasaran_diundi"`
	Pasaran_url       string `json:"pasaran_url"`
	Pasaran_jamtutup  string `json:"pasaran_jamtutup"`
	Pasaran_jamjadwal string `json:"pasaran_jamjadwal"`
	Pasaran_jamopen   string `json:"pasaran_jamopen"`
	Pasaran_display   int    `json:"pasaran_display"`
	Pasaran_status    string `json:"pasaran_status"`
}
type pasaransaveonline struct {
	Idpasaran   int    `json:"idpasaran" validate:"required"`
	Page        string `json:"page"`
	Haripasaran string `json:"pasaran_hariraya" validate:"required,alpha"`
}
type pasarandeleteonline struct {
	Idpasaran      int    `json:"idpasaran" validate:"required"`
	Idpasaraonline int    `json:"idpasaraonline" validate:"required,numeric"`
	Page           string `json:"page"`
}
type pasaransavelimit struct {
	Idpasaran            int    `json:"idpasaran"`
	Page                 string `json:"page"`
	Pasaran_limitline4d  int    `json:"pasaran_limitline4d"`
	Pasaran_limitline3d  int    `json:"pasaran_limitline3d"`
	Pasaran_limitline3dd int    `json:"pasaran_limitline3dd"`
	Pasaran_limitline2d  int    `json:"pasaran_limitline2d"`
	Pasaran_limitline2dd int    `json:"pasaran_limitline2dd"`
	Pasaran_limitline2dt int    `json:"pasaran_limitline2dt"`
}

type pasaranconfcbebas struct {
	Idpasaran                  int     `json:"idpasaran"`
	Idpasarantogel             string  `json:"idpasarantogel"`
	Page                       string  `json:"page"`
	Pasaran_minbet_cbebas      int     `json:"pasaran_minbet_cbebas"`
	Pasaran_maxbet_cbebas      int     `json:"pasaran_maxbet_cbebas"`
	Pasaran_limitotal_cbebas   int     `json:"pasaran_limitotal_cbebas"`
	Pasaran_limitglobal_cbebas int     `json:"pasaran_limitglobal_cbebas"`
	Pasaran_win_cbebas         float32 `json:"pasaran_win_cbebas"`
	Pasaran_disc_cbebas        float32 `json:"pasaran_disc_cbebas"`
}
type pasaranconfcmacau struct {
	Idpasaran                  int     `json:"idpasaran"`
	Idpasarantogel             string  `json:"idpasarantogel"`
	Page                       string  `json:"page"`
	Pasaran_minbet_cmacau      int     `json:"pasaran_minbet_cmacau"`
	Pasaran_maxbet_cmacau      int     `json:"pasaran_maxbet_cmacau"`
	Pasaran_limitotal_cmacau   int     `json:"pasaran_limitotal_cmacau"`
	Pasaran_limitglobal_cmacau int     `json:"pasaran_limitglobal_cmacau"`
	Pasaran_win2_cmacau        float32 `json:"pasaran_win2_cmacau"`
	Pasaran_win3_cmacau        float32 `json:"pasaran_win3_cmacau"`
	Pasaran_win4_cmacau        float32 `json:"pasaran_win4_cmacau"`
	Pasaran_disc_cmacau        float32 `json:"pasaran_disc_cmacau"`
}
type pasaranconfcnaga struct {
	Idpasaran                 int     `json:"idpasaran"`
	Idpasarantogel            string  `json:"idpasarantogel"`
	Page                      string  `json:"page"`
	Pasaran_minbet_cnaga      int     `json:"pasaran_minbet_cnaga"`
	Pasaran_maxbet_cnaga      int     `json:"pasaran_maxbet_cnaga"`
	Pasaran_limittotal_cnaga  int     `json:"pasaran_limittotal_cnaga"`
	Pasaran_limitglobal_cnaga int     `json:"pasaran_limitglobal_cnaga"`
	Pasaran_win3_cnaga        float32 `json:"pasaran_win3_cnaga"`
	Pasaran_win4_cnaga        float32 `json:"pasaran_win4_cnaga"`
	Pasaran_disc_cnaga        float32 `json:"pasaran_disc_cnaga"`
}
type pasaranconfcjitu struct {
	Idpasaran                 int     `json:"idpasaran"`
	Idpasarantogel            string  `json:"idpasarantogel"`
	Page                      string  `json:"page"`
	Pasaran_minbet_cjitu      int     `json:"pasaran_minbet_cjitu"`
	Pasaran_maxbet_cjitu      int     `json:"pasaran_maxbet_cjitu"`
	Pasaran_limittotal_cjitu  int     `json:"pasaran_limittotal_cjitu"`
	Pasaran_limitglobal_cjitu int     `json:"pasaran_limitglobal_cjitu"`
	Pasaran_winas_cjitu       float32 `json:"pasaran_winas_cjitu"`
	Pasaran_winkop_cjitu      float32 `json:"pasaran_winkop_cjitu"`
	Pasaran_winkepala_cjitu   float32 `json:"pasaran_winkepala_cjitu"`
	Pasaran_winekor_cjitu     float32 `json:"pasaran_winekor_cjitu"`
	Pasaran_desc_cjitu        float32 `json:"pasaran_desc_cjitu"`
}
type pasaranconfc5050 struct {
	Idpasaran                    int     `json:"idpasaran"`
	Idpasarantogel               string  `json:"idpasarantogel"`
	Page                         string  `json:"page"`
	Pasaran_minbet_5050umum      int     `json:"pasaran_minbet_5050umum"`
	Pasaran_maxbet_5050umum      int     `json:"pasaran_maxbet_5050umum"`
	Pasaran_limittotal_5050umum  int     `json:"pasaran_limittotal_5050umum"`
	Pasaran_limitglobal_5050umum int     `json:"pasaran_limitglobal_5050umum"`
	Pasaran_keibesar_5050umum    float32 `json:"pasaran_keibesar_5050umum"`
	Pasaran_keikecil_5050umum    float32 `json:"pasaran_keikecil_5050umum"`
	Pasaran_keigenap_5050umum    float32 `json:"pasaran_keigenap_5050umum"`
	Pasaran_keiganjil_5050umum   float32 `json:"pasaran_keiganjil_5050umum"`
	Pasaran_keitengah_5050umum   float32 `json:"pasaran_keitengah_5050umum"`
	Pasaran_keitepi_5050umum     float32 `json:"pasaran_keitepi_5050umum"`
	Pasaran_discbesar_5050umum   float32 `json:"pasaran_discbesar_5050umum"`
	Pasaran_disckecil_5050umum   float32 `json:"pasaran_disckecil_5050umum"`
	Pasaran_discgenap_5050umum   float32 `json:"pasaran_discgenap_5050umum"`
	Pasaran_discganjil_5050umum  float32 `json:"pasaran_discganjil_5050umum"`
	Pasaran_disctengah_5050umum  float32 `json:"pasaran_disctengah_5050umum"`
	Pasaran_disctepi_5050umum    float32 `json:"pasaran_disctepi_5050umum"`
}
type pasaranconfc5050special struct {
	Idpasaran                            int     `json:"idpasaran"`
	Idpasarantogel                       string  `json:"idpasarantogel"`
	Page                                 string  `json:"page"`
	Pasaran_minbet_5050special           int     `json:"pasaran_minbet_5050special"`
	Pasaran_maxbet_5050special           int     `json:"pasaran_maxbet_5050special"`
	Pasaran_limittotal_5050special       int     `json:"pasaran_limittotal_5050special"`
	Pasaran_limitglobal_5050special      int     `json:"pasaran_limitglobal_5050special"`
	Pasaran_keiasganjil_5050special      float32 `json:"pasaran_keiasganjil_5050special"`
	Pasaran_keiasgenap_5050special       float32 `json:"pasaran_keiasgenap_5050special"`
	Pasaran_keiasbesar_5050special       float32 `json:"pasaran_keiasbesar_5050special"`
	Pasaran_keiaskecil_5050special       float32 `json:"pasaran_keiaskecil_5050special"`
	Pasaran_keikopganjil_5050special     float32 `json:"pasaran_keikopganjil_5050special"`
	Pasaran_keikopgenap_5050special      float32 `json:"pasaran_keikopgenap_5050special"`
	Pasaran_keikopbesar_5050special      float32 `json:"pasaran_keikopbesar_5050special"`
	Pasaran_keikopkecil_5050special      float32 `json:"pasaran_keikopkecil_5050special"`
	Pasaran_keikepalaganjil_5050special  float32 `json:"pasaran_keikepalaganjil_5050special"`
	Pasaran_keikepalagenap_5050special   float32 `json:"pasaran_keikepalagenap_5050special"`
	Pasaran_keikepalabesar_5050special   float32 `json:"pasaran_keikepalabesar_5050special"`
	Pasaran_keikepalakecil_5050special   float32 `json:"pasaran_keikepalakecil_5050special"`
	Pasaran_keiekorganjil_5050special    float32 `json:"pasaran_keiekorganjil_5050special"`
	Pasaran_keiekorgenap_5050special     float32 `json:"pasaran_keiekorgenap_5050special"`
	Pasaran_keiekorbesar_5050special     float32 `json:"pasaran_keiekorbesar_5050special"`
	Pasaran_keiekorkecil_5050special     float32 `json:"pasaran_keiekorkecil_5050special"`
	Pasaran_discasganjil_5050special     float32 `json:"pasaran_discasganjil_5050special"`
	Pasaran_discasgenap_5050special      float32 `json:"pasaran_discasgenap_5050special"`
	Pasaran_discasbesar_5050special      float32 `json:"pasaran_discasbesar_5050special"`
	Pasaran_discaskecil_5050special      float32 `json:"pasaran_discaskecil_5050special"`
	Pasaran_disckopganjil_5050special    float32 `json:"pasaran_disckopganjil_5050special"`
	Pasaran_disckopgenap_5050special     float32 `json:"pasaran_disckopgenap_5050special"`
	Pasaran_disckopbesar_5050special     float32 `json:"pasaran_disckopbesar_5050special"`
	Pasaran_disckopkecil_5050special     float32 `json:"pasaran_disckopkecil_5050special"`
	Pasaran_disckepalaganjil_5050special float32 `json:"pasaran_disckepalaganjil_5050special"`
	Pasaran_disckepalagenap_5050special  float32 `json:"pasaran_disckepalagenap_5050special"`
	Pasaran_disckepalabesar_5050special  float32 `json:"pasaran_disckepalabesar_5050special"`
	Pasaran_disckepalakecil_5050special  float32 `json:"pasaran_disckepalakecil_5050special"`
	Pasaran_discekorganjil_5050special   float32 `json:"pasaran_discekorganjil_5050special"`
	Pasaran_discekorgenap_5050special    float32 `json:"pasaran_discekorgenap_5050special"`
	Pasaran_discekorbesar_5050special    float32 `json:"pasaran_discekorbesar_5050special"`
	Pasaran_discekorkecil_5050special    float32 `json:"pasaran_discekorkecil_5050special"`
}
type pasaranconfc5050kombinasi struct {
	Idpasaran                                 int     `json:"idpasaran"`
	Idpasarantogel                            string  `json:"idpasarantogel"`
	Page                                      string  `json:"page"`
	Pasaran_minbet_5050kombinasi              int     `json:"pasaran_minbet_5050kombinasi"`
	Pasaran_maxbet_5050kombinasi              int     `json:"pasaran_maxbet_5050kombinasi"`
	Pasaran_limittotal_5050kombinasi          int     `json:"pasaran_limittotal_5050kombinasi"`
	Pasaran_limitglobal_5050kombinasi         int     `json:"pasaran_limitglobal_5050kombinasi"`
	Pasaran_belakangkeimono_5050kombinasi     float32 `json:"pasaran_belakangkeimono_5050kombinasi"`
	Pasaran_belakangkeistereo_5050kombinasi   float32 `json:"pasaran_belakangkeistereo_5050kombinasi"`
	Pasaran_belakangkeikembang_5050kombinasi  float32 `json:"pasaran_belakangkeikembang_5050kombinasi"`
	Pasaran_belakangkeikempis_5050kombinasi   float32 `json:"pasaran_belakangkeikempis_5050kombinasi"`
	Pasaran_belakangkeikembar_5050kombinasi   float32 `json:"pasaran_belakangkeikembar_5050kombinasi"`
	Pasaran_tengahkeimono_5050kombinasi       float32 `json:"pasaran_tengahkeimono_5050kombinasi"`
	Pasaran_tengahkeistereo_5050kombinasi     float32 `json:"pasaran_tengahkeistereo_5050kombinasi"`
	Pasaran_tengahkeikembang_5050kombinasi    float32 `json:"pasaran_tengahkeikembang_5050kombinasi"`
	Pasaran_tengahkeikempis_5050kombinasi     float32 `json:"pasaran_tengahkeikempis_5050kombinasi"`
	Pasaran_tengahkeikembar_5050kombinasi     float32 `json:"pasaran_tengahkeikembar_5050kombinasi"`
	Pasaran_depankeimono_5050kombinasi        float32 `json:"pasaran_depankeimono_5050kombinasi"`
	Pasaran_depankeistereo_5050kombinasi      float32 `json:"pasaran_depankeistereo_5050kombinasi"`
	Pasaran_depankeikembang_5050kombinasi     float32 `json:"pasaran_depankeikembang_5050kombinasi"`
	Pasaran_depankeikempis_5050kombinasi      float32 `json:"pasaran_depankeikempis_5050kombinasi"`
	Pasaran_depankeikembar_5050kombinasi      float32 `json:"pasaran_depankeikembar_5050kombinasi"`
	Pasaran_belakangdiscmono_5050kombinasi    float32 `json:"pasaran_belakangdiscmono_5050kombinasi"`
	Pasaran_belakangdiscstereo_5050kombinasi  float32 `json:"pasaran_belakangdiscstereo_5050kombinasi"`
	Pasaran_belakangdisckembang_5050kombinasi float32 `json:"pasaran_belakangdisckembang_5050kombinasi"`
	Pasaran_belakangdisckempis_5050kombinasi  float32 `json:"pasaran_belakangdisckempis_5050kombinasi"`
	Pasaran_belakangdisckembar_5050kombinasi  float32 `json:"pasaran_belakangdisckembar_5050kombinasi"`
	Pasaran_tengahdiscmono_5050kombinasi      float32 `json:"pasaran_tengahdiscmono_5050kombinasi"`
	Pasaran_tengahdiscstereo_5050kombinasi    float32 `json:"pasaran_tengahdiscstereo_5050kombinasi"`
	Pasaran_tengahdisckembang_5050kombinasi   float32 `json:"pasaran_tengahdisckembang_5050kombinasi"`
	Pasaran_tengahdisckempis_5050kombinasi    float32 `json:"pasaran_tengahdisckempis_5050kombinasi"`
	Pasaran_tengahdisckembar_5050kombinasi    float32 `json:"pasaran_tengahdisckembar_5050kombinasi"`
	Pasaran_depandiscmono_5050kombinasi       float32 `json:"pasaran_depandiscmono_5050kombinasi"`
	Pasaran_depandiscstereo_5050kombinasi     float32 `json:"pasaran_depandiscstereo_5050kombinasi"`
	Pasaran_depandisckembang_5050kombinasi    float32 `json:"pasaran_depandisckembang_5050kombinasi"`
	Pasaran_depandisckempis_5050kombinasi     float32 `json:"pasaran_depandisckempis_5050kombinasi"`
	Pasaran_depandisckembar_5050kombinasi     float32 `json:"pasaran_depandisckembar_5050kombinasi"`
}
type pasaranconfmakaukombinasi struct {
	Idpasaran                     int     `json:"idpasaran"`
	Idpasarantogel                string  `json:"idpasarantogel"`
	Page                          string  `json:"page"`
	Pasaran_minbet_kombinasi      int     `json:"pasaran_minbet_kombinasi"`
	Pasaran_maxbet_kombinasi      int     `json:"pasaran_maxbet_kombinasi"`
	Pasaran_limittotal_kombinasi  int     `json:"pasaran_limittotal_kombinasi"`
	Pasaran_limitglobal_kombinasi int     `json:"pasaran_limitglobal_kombinasi"`
	Pasaran_win_kombinasi         float32 `json:"pasaran_win_kombinasi"`
	Pasaran_disc_kombinasi        float32 `json:"pasaran_disc_kombinasi"`
}
type pasaranconfdasar struct {
	Idpasaran                 int     `json:"idpasaran"`
	Idpasarantogel            string  `json:"idpasarantogel"`
	Page                      string  `json:"page"`
	Pasaran_minbet_dasar      int     `json:"pasaran_minbet_dasar"`
	Pasaran_maxbet_dasar      int     `json:"pasaran_maxbet_dasar"`
	Pasaran_limittotal_dasar  int     `json:"pasaran_limittotal_dasar"`
	Pasaran_limitglobal_dasar int     `json:"pasaran_limitglobal_dasar"`
	Pasaran_keibesar_dasar    float32 `json:"pasaran_keibesar_dasar"`
	Pasaran_keikecil_dasar    float32 `json:"pasaran_keikecil_dasar"`
	Pasaran_keigenap_dasar    float32 `json:"pasaran_keigenap_dasar"`
	Pasaran_keiganjil_dasar   float32 `json:"pasaran_keiganjil_dasar"`
	Pasaran_discbesar_dasar   float32 `json:"pasaran_discbesar_dasar"`
	Pasaran_disckecil_dasar   float32 `json:"pasaran_disckecil_dasar"`
	Pasaran_discgenap_dasar   float32 `json:"pasaran_discgenap_dasar"`
	Pasaran_discganjil_dasar  float32 `json:"pasaran_discganjil_dasar"`
}
type pasaranconfshio struct {
	Idpasaran                int     `json:"idpasaran"`
	Idpasarantogel           string  `json:"idpasarantogel"`
	Page                     string  `json:"page"`
	Pasaran_minbet_shio      int     `json:"pasaran_minbet_shio"`
	Pasaran_maxbet_shio      int     `json:"pasaran_maxbet_shio"`
	Pasaran_limittotal_shio  int     `json:"pasaran_limittotal_shio"`
	Pasaran_limitglobal_shio int     `json:"pasaran_limitglobal_shio"`
	Pasaran_shioyear_shio    string  `json:"pasaran_shioyear_shio"`
	Pasaran_disc_shio        float32 `json:"pasaran_disc_shio"`
	Pasaran_win_shio         float32 `json:"pasaran_win_shio"`
}
type response_pasaran struct {
	Status        int         `json:"status"`
	Message       string      `json:"message"`
	Record        interface{} `json:"record"`
	Pasaranonline interface{} `json:"pasaranonline"`
}
type response_pasarandefault struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Record  interface{} `json:"record"`
}
type response_pasaransave struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Record  interface{} `json:"record"`
}
type responseerror struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func Pasaran(c *fiber.Ctx) error {
	render_page := time.Now()
	bearToken := c.Get("Authorization")
	token := strings.Split(bearToken, " ")
	axios := resty.New()
	resp, err := axios.R().
		SetAuthToken(token[1]).
		SetResult(response_pasarandefault{}).
		SetError(response_pasarandefault{}).
		SetHeader("Content-Type", "application/json").
		Post(config.Path_url() + "api/allpasaran")
	if err != nil {
		log.Println(err.Error())
	}
	result := resp.Result().(*response_pasarandefault)
	if result.Status == 200 {
		c.Status(fiber.StatusOK)
		return c.JSON(fiber.Map{
			"status":  http.StatusOK,
			"message": result.Message,
			"record":  result.Record,
			"time":    time.Since(render_page).String(),
		})
	} else {
		result_error := resp.Error().(*response_pasarandefault)
		c.Status(result_error.Status)
		return c.JSON(fiber.Map{
			"status":  result_error.Status,
			"message": result_error.Message,
			"record":  nil,
			"time":    time.Since(render_page).String(),
		})
	}
}
func Pasarandetail(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(pasarandetail)
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
		SetResult(response_pasaran{}).
		SetError(response_pasaran{}).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"idpasaran": client.Idpasaran,
		}).
		Post(config.Path_url() + "api/editpasaran")
	if err != nil {
		log.Println(err.Error())
	}
	result := resp.Result().(*response_pasaran)
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
		result_error := resp.Error().(*response_pasaran)
		c.Status(result_error.Status)
		return c.JSON(fiber.Map{
			"status":  result_error.Status,
			"message": result_error.Message,
			"record":  nil,
			"time":    time.Since(render_page).String(),
		})
	}
}
func Pasaransave(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(pasaransave)
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
			"idpasaran":         client.Idpasaran,
			"page":              client.Page,
			"pasaran_diundi":    client.Pasaran_diundi,
			"pasaran_url":       client.Pasaran_url,
			"pasaran_jamtutup":  client.Pasaran_jamtutup,
			"pasaran_jamjadwal": client.Pasaran_jamjadwal,
			"pasaran_jamopen":   client.Pasaran_jamopen,
			"pasaran_display":   client.Pasaran_display,
			"pasaran_status":    client.Pasaran_status,
		}).
		Post(config.Path_url() + "api/savepasaran")
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
	result := resp.Result().(*response_pasaransave)
	if result.Status == 200 {
		return c.JSON(fiber.Map{
			"status":  result.Status,
			"message": result.Message,
			"record":  result.Record,
			"time":    time.Since(render_page).String(),
		})
	} else {
		result_error := resp.Error().(*response_pasaransave)
		return c.JSON(fiber.Map{
			"status":  result_error.Status,
			"message": result_error.Message,
			"record":  result.Record,
			"time":    time.Since(render_page).String(),
		})
	}
}
func Pasaransaveonline(c *fiber.Ctx) error {
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
func Pasarandeleteonline(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(pasarandeleteonline)
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
			"idpasaran":      client.Idpasaran,
			"idpasaraonline": client.Idpasaraonline,
			"page":           client.Page,
		}).
		Post(config.Path_url() + "api/deletepasaranonline")
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
func Pasaransavelimit(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(pasaransavelimit)
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
			"idpasaran":            client.Idpasaran,
			"page":                 client.Page,
			"pasaran_limitline4d":  client.Pasaran_limitline4d,
			"pasaran_limitline3d":  client.Pasaran_limitline3d,
			"pasaran_limitline3dd": client.Pasaran_limitline3dd,
			"pasaran_limitline2d":  client.Pasaran_limitline2d,
			"pasaran_limitline2dd": client.Pasaran_limitline2dd,
			"pasaran_limitline2dt": client.Pasaran_limitline2dt,
		}).
		Post(config.Path_url() + "api/savepasaranlimitline")
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
func Pasaransaveconf432d(c *fiber.Ctx) error {
	type payload_pasaranconf432 struct {
		Idpasaran                   int     `json:"idpasaran"`
		Idpasarantogel              string  `json:"idpasarantogel"`
		Page                        string  `json:"page"`
		Pasaran_minbet_432d         int     `json:"pasaran_minbet_432d"`
		Pasaran_maxbet4d_432d       int     `json:"pasaran_maxbet4d_432d"`
		Pasaran_maxbet3d_432d       int     `json:"pasaran_maxbet3d_432d"`
		Pasaran_maxbet3dd_432d      int     `json:"pasaran_maxbet3dd_432d"`
		Pasaran_maxbet2d_432d       int     `json:"pasaran_maxbet2d_432d"`
		Pasaran_maxbet2dd_432d      int     `json:"pasaran_maxbet2dd_432d"`
		Pasaran_maxbet2dt_432d      int     `json:"pasaran_maxbet2dt_432d"`
		Pasaran_limitotal4d_432d    int     `json:"pasaran_limitotal4d_432d"`
		Pasaran_limitotal3d_432d    int     `json:"pasaran_limitotal3d_432d"`
		Pasaran_limitotal3dd_432d   int     `json:"pasaran_limitotal3dd_432d"`
		Pasaran_limitotal2d_432d    int     `json:"pasaran_limitotal2d_432d"`
		Pasaran_limitotal2dd_432d   int     `json:"pasaran_limitotal2dd_432d"`
		Pasaran_limitotal2dt_432d   int     `json:"pasaran_limitotal2dt_432d"`
		Pasaran_limitglobal4d_432d  int     `json:"pasaran_limitglobal4d_432d"`
		Pasaran_limitglobal3d_432d  int     `json:"pasaran_limitglobal3d_432d"`
		Pasaran_limitglobal3dd_432d int     `json:"pasaran_limitglobal3dd_432d"`
		Pasaran_limitglobal2d_432d  int     `json:"pasaran_limitglobal2d_432d"`
		Pasaran_limitglobal2dd_432d int     `json:"pasaran_limitglobal2dd_432d"`
		Pasaran_limitglobal2dt_432d int     `json:"pasaran_limitglobal2dt_432d"`
		Pasaran_win4d_432d          int     `json:"pasaran_win4d_432d"`
		Pasaran_win3d_432d          int     `json:"pasaran_win3d_432d"`
		Pasaran_win3dd_432d         int     `json:"pasaran_win3dd_432d"`
		Pasaran_win2d_432d          int     `json:"pasaran_win2d_432d"`
		Pasaran_win2dd_432d         int     `json:"pasaran_win2dd_432d"`
		Pasaran_win2dt_432d         int     `json:"pasaran_win2dt_432d"`
		Pasaran_win4dnodisc_432d    int     `json:"pasaran_win4dnodisc_432d"`
		Pasaran_win3dnodisc_432d    int     `json:"pasaran_win3dnodisc_432d"`
		Pasaran_win3ddnodisc_432d   int     `json:"pasaran_win3ddnodisc_432d"`
		Pasaran_win2dnodisc_432d    int     `json:"pasaran_win2dnodisc_432d"`
		Pasaran_win2ddnodisc_432d   int     `json:"pasaran_win2ddnodisc_432d"`
		Pasaran_win2dtnodisc_432d   int     `json:"pasaran_win2dtnodisc_432d"`
		Pasaran_win4d_bb_432d       int     `json:"pasaran_win4d_bb_432d" `
		Pasaran_win3d_bb_432d       int     `json:"pasaran_win3d_bb_432d" `
		Pasaran_win3dd_bb_432d      int     `json:"pasaran_win3dd_bb_432d" `
		Pasaran_win2d_bb_432d       int     `json:"pasaran_win2d_bb_432d" `
		Pasaran_win2dd_bb_432d      int     `json:"pasaran_win2dd_bb_432d" `
		Pasaran_win2dt_bb_432d      int     `json:"pasaran_win2dt_bb_432d" `
		Pasaran_win4d_bb_kena_432d  int     `json:"pasaran_win4d_bb_kena_432d" `
		Pasaran_win3d_bb_kena_432d  int     `json:"pasaran_win3d_bb_kena_432d" `
		Pasaran_win3dd_bb_kena_432d int     `json:"pasaran_win3dd_bb_kena_432d" `
		Pasaran_win2d_bb_kena_432d  int     `json:"pasaran_win2d_bb_kena_432d" `
		Pasaran_win2dd_bb_kena_432d int     `json:"pasaran_win2dd_bb_kena_432d" `
		Pasaran_win2dt_bb_kena_432d int     `json:"pasaran_win2dt_bb_kena_432d" `
		Pasaran_disc4d_432d         float32 `json:"pasaran_disc4d_432d"`
		Pasaran_disc3d_432d         float32 `json:"pasaran_disc3d_432d"`
		Pasaran_disc3dd_432d        float32 `json:"pasaran_disc3dd_432d"`
		Pasaran_disc2d_432d         float32 `json:"pasaran_disc2d_432d"`
		Pasaran_disc2dd_432d        float32 `json:"pasaran_disc2dd_432d"`
		Pasaran_disc2dt_432d        float32 `json:"pasaran_disc2dt_432d"`
	}
	var errors []*helpers.ErrorResponse
	client := new(payload_pasaranconf432)
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
			"idpasaran":                   client.Idpasaran,
			"idpasarantogel":              client.Idpasarantogel,
			"page":                        client.Page,
			"pasaran_minbet_432d":         client.Pasaran_minbet_432d,
			"pasaran_maxbet4d_432d":       client.Pasaran_maxbet4d_432d,
			"pasaran_maxbet3d_432d":       client.Pasaran_maxbet3d_432d,
			"pasaran_maxbet3dd_432d":      client.Pasaran_maxbet3dd_432d,
			"pasaran_maxbet2d_432d":       client.Pasaran_maxbet2d_432d,
			"pasaran_maxbet2dd_432d":      client.Pasaran_maxbet2dd_432d,
			"pasaran_maxbet2dt_432d":      client.Pasaran_maxbet2dt_432d,
			"pasaran_limitotal4d_432d":    client.Pasaran_limitotal4d_432d,
			"pasaran_limitotal3d_432d":    client.Pasaran_limitotal3d_432d,
			"pasaran_limitotal3dd_432d":   client.Pasaran_limitotal3dd_432d,
			"pasaran_limitotal2d_432d":    client.Pasaran_limitotal2d_432d,
			"pasaran_limitotal2dd_432d":   client.Pasaran_limitotal2dd_432d,
			"pasaran_limitotal2dt_432d":   client.Pasaran_limitotal2dt_432d,
			"pasaran_limitglobal4d_432d":  client.Pasaran_limitglobal4d_432d,
			"pasaran_limitglobal3d_432d":  client.Pasaran_limitglobal3d_432d,
			"pasaran_limitglobal3dd_432d": client.Pasaran_limitglobal3dd_432d,
			"pasaran_limitglobal2d_432d":  client.Pasaran_limitglobal2d_432d,
			"pasaran_limitglobal2dd_432d": client.Pasaran_limitglobal2dt_432d,
			"pasaran_limitglobal2dt_432d": client.Pasaran_limitglobal2dt_432d,
			"pasaran_win4d_432d":          client.Pasaran_win4d_432d,
			"pasaran_win3d_432d":          client.Pasaran_win3d_432d,
			"pasaran_win3dd_432d":         client.Pasaran_win3dd_432d,
			"pasaran_win2d_432d":          client.Pasaran_win2d_432d,
			"pasaran_win2dd_432d":         client.Pasaran_win2dd_432d,
			"pasaran_win2dt_432d":         client.Pasaran_win2dt_432d,
			"pasaran_win4dnodisc_432d":    client.Pasaran_win4dnodisc_432d,
			"pasaran_win3dnodisc_432d":    client.Pasaran_win3dnodisc_432d,
			"pasaran_win3ddnodisc_432d":   client.Pasaran_win3ddnodisc_432d,
			"pasaran_win2dnodisc_432d":    client.Pasaran_win2dnodisc_432d,
			"pasaran_win2ddnodisc_432d":   client.Pasaran_win2ddnodisc_432d,
			"pasaran_win2dtnodisc_432d":   client.Pasaran_win2dtnodisc_432d,
			"pasaran_win4dbb_kena_432d":   client.Pasaran_win4d_bb_kena_432d,
			"pasaran_win3dbb_kena_432d":   client.Pasaran_win3d_bb_kena_432d,
			"pasaran_win3ddbb_kena_432d":  client.Pasaran_win3dd_bb_kena_432d,
			"pasaran_win2dbb_kena_432d":   client.Pasaran_win2d_bb_kena_432d,
			"pasaran_win2ddbb_kena_432d":  client.Pasaran_win2dd_bb_kena_432d,
			"pasaran_win2dtbb_kena_432d":  client.Pasaran_win2dt_bb_kena_432d,
			"pasaran_win4dbb_432d":        client.Pasaran_win4d_bb_432d,
			"pasaran_win3dbb_432d":        client.Pasaran_win3d_bb_432d,
			"pasaran_win3ddbb_432d":       client.Pasaran_win3dd_bb_432d,
			"pasaran_win2dbb_432d":        client.Pasaran_win2d_bb_432d,
			"pasaran_win2ddbb_432d":       client.Pasaran_win2dd_bb_432d,
			"pasaran_win2dtbb_432d":       client.Pasaran_win2dt_bb_432d,
			"pasaran_disc4d_432d":         client.Pasaran_disc4d_432d,
			"pasaran_disc3d_432d":         client.Pasaran_disc3d_432d,
			"pasaran_disc3dd_432d":        client.Pasaran_disc3dd_432d,
			"pasaran_disc2d_432d":         client.Pasaran_disc2d_432d,
			"pasaran_disc2dd_432d":        client.Pasaran_disc2dd_432d,
			"pasaran_disc2dt_432d":        client.Pasaran_disc2dt_432d,
		}).
		Post(config.Path_url() + "api/savepasaranconf432")
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
func Pasaransaveconfcbebas(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(pasaranconfcbebas)
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
			"idpasaran":                  client.Idpasaran,
			"idpasarantogel":             client.Idpasarantogel,
			"page":                       client.Page,
			"pasaran_minbet_cbebas":      client.Pasaran_minbet_cbebas,
			"pasaran_maxbet_cbebas":      client.Pasaran_maxbet_cbebas,
			"pasaran_limitotal_cbebas":   client.Pasaran_limitotal_cbebas,
			"pasaran_limitglobal_cbebas": client.Pasaran_limitglobal_cbebas,
			"pasaran_win_cbebas":         client.Pasaran_win_cbebas,
			"pasaran_disc_cbebas":        client.Pasaran_disc_cbebas,
		}).
		Post(config.Path_url() + "api/savepasaranconfcolokbebas")
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
func Pasaransaveconfcmacau(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(pasaranconfcmacau)
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
			"idpasaran":                  client.Idpasaran,
			"idpasarantogel":             client.Idpasarantogel,
			"page":                       client.Page,
			"pasaran_minbet_cmacau":      client.Pasaran_minbet_cmacau,
			"pasaran_maxbet_cmacau":      client.Pasaran_maxbet_cmacau,
			"pasaran_limitotal_cmacau":   client.Pasaran_limitotal_cmacau,
			"pasaran_limitglobal_cmacau": client.Pasaran_limitglobal_cmacau,
			"pasaran_win2_cmacau":        client.Pasaran_win2_cmacau,
			"pasaran_win3_cmacau":        client.Pasaran_win3_cmacau,
			"pasaran_win4_cmacau":        client.Pasaran_win4_cmacau,
			"pasaran_disc_cmacau":        client.Pasaran_disc_cmacau,
		}).
		Post(config.Path_url() + "api/savepasaranconfcolokmacau")
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
func Pasaransaveconfcnaga(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(pasaranconfcnaga)
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
			"idpasaran":                 client.Idpasaran,
			"idpasarantogel":            client.Idpasarantogel,
			"page":                      client.Page,
			"pasaran_minbet_cnaga":      client.Pasaran_minbet_cnaga,
			"pasaran_maxbet_cnaga":      client.Pasaran_maxbet_cnaga,
			"pasaran_limittotal_cnaga":  client.Pasaran_limittotal_cnaga,
			"pasaran_limitglobal_cnaga": client.Pasaran_limitglobal_cnaga,
			"pasaran_win3_cnaga":        client.Pasaran_win3_cnaga,
			"pasaran_win4_cnaga":        client.Pasaran_win4_cnaga,
			"pasaran_disc_cnaga":        client.Pasaran_disc_cnaga,
		}).
		Post(config.Path_url() + "api/savepasaranconfcoloknaga")
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
func Pasaransaveconfcjitu(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(pasaranconfcjitu)
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
			"idpasaran":                 client.Idpasaran,
			"idpasarantogel":            client.Idpasarantogel,
			"page":                      client.Page,
			"pasaran_minbet_cjitu":      client.Pasaran_minbet_cjitu,
			"pasaran_maxbet_cjitu":      client.Pasaran_maxbet_cjitu,
			"pasaran_limittotal_cjitu":  client.Pasaran_limittotal_cjitu,
			"pasaran_limitglobal_cjitu": client.Pasaran_limitglobal_cjitu,
			"pasaran_winas_cjitu":       client.Pasaran_winas_cjitu,
			"pasaran_winkop_cjitu":      client.Pasaran_winkop_cjitu,
			"pasaran_winkepala_cjitu":   client.Pasaran_winkepala_cjitu,
			"pasaran_winekor_cjitu":     client.Pasaran_winekor_cjitu,
			"pasaran_desc_cjitu":        client.Pasaran_desc_cjitu,
		}).
		Post(config.Path_url() + "api/savepasaranconfcolokjitu")
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
func Pasaransaveconf5050umum(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(pasaranconfc5050)
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
			"idpasaran":                    client.Idpasaran,
			"idpasarantogel":               client.Idpasarantogel,
			"page":                         client.Page,
			"pasaran_minbet_5050umum":      client.Pasaran_minbet_5050umum,
			"pasaran_maxbet_5050umum":      client.Pasaran_maxbet_5050umum,
			"pasaran_limittotal_5050umum":  client.Pasaran_limittotal_5050umum,
			"pasaran_limitglobal_5050umum": client.Pasaran_limitglobal_5050umum,
			"pasaran_keibesar_5050umum":    client.Pasaran_keibesar_5050umum,
			"pasaran_keikecil_5050umum":    client.Pasaran_keikecil_5050umum,
			"pasaran_keigenap_5050umum":    client.Pasaran_keigenap_5050umum,
			"pasaran_keiganjil_5050umum":   client.Pasaran_keiganjil_5050umum,
			"pasaran_keitengah_5050umum":   client.Pasaran_keitengah_5050umum,
			"pasaran_keitepi_5050umum":     client.Pasaran_keitepi_5050umum,
			"pasaran_discbesar_5050umum":   client.Pasaran_discbesar_5050umum,
			"pasaran_disckecil_5050umum":   client.Pasaran_disckecil_5050umum,
			"pasaran_discgenap_5050umum":   client.Pasaran_discgenap_5050umum,
			"pasaran_discganjil_5050umum":  client.Pasaran_discganjil_5050umum,
			"pasaran_disctengah_5050umum":  client.Pasaran_disctengah_5050umum,
			"pasaran_disctepi_5050umum":    client.Pasaran_disctepi_5050umum,
		}).
		Post(config.Path_url() + "api/savepasaranconf5050umum")
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
func Pasaransaveconf5050special(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(pasaranconfc5050special)
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
			"idpasaran":                            client.Idpasaran,
			"idpasarantogel":                       client.Idpasarantogel,
			"page":                                 client.Page,
			"pasaran_minbet_5050special":           client.Pasaran_minbet_5050special,
			"pasaran_maxbet_5050special":           client.Pasaran_maxbet_5050special,
			"pasaran_limittotal_5050special":       client.Pasaran_limittotal_5050special,
			"pasaran_limitglobal_5050special":      client.Pasaran_limitglobal_5050special,
			"pasaran_keiasganjil_5050special":      client.Pasaran_keiasganjil_5050special,
			"pasaran_keiasgenap_5050special":       client.Pasaran_keiasgenap_5050special,
			"pasaran_keiasbesar_5050special":       client.Pasaran_keiasbesar_5050special,
			"pasaran_keiaskecil_5050special":       client.Pasaran_keiaskecil_5050special,
			"pasaran_keikopganjil_5050special":     client.Pasaran_keikopganjil_5050special,
			"pasaran_keikopgenap_5050special":      client.Pasaran_keikopgenap_5050special,
			"pasaran_keikopbesar_5050special":      client.Pasaran_keikopbesar_5050special,
			"pasaran_keikopkecil_5050special":      client.Pasaran_keikopkecil_5050special,
			"pasaran_keikepalaganjil_5050special":  client.Pasaran_keikepalaganjil_5050special,
			"pasaran_keikepalagenap_5050special":   client.Pasaran_keikepalagenap_5050special,
			"pasaran_keikepalabesar_5050special":   client.Pasaran_keikepalabesar_5050special,
			"pasaran_keikepalakecil_5050special":   client.Pasaran_keikepalakecil_5050special,
			"pasaran_keiekorganjil_5050special":    client.Pasaran_keiekorganjil_5050special,
			"pasaran_keiekorgenap_5050special":     client.Pasaran_keiekorgenap_5050special,
			"pasaran_keiekorbesar_5050special":     client.Pasaran_keiekorbesar_5050special,
			"pasaran_keiekorkecil_5050special":     client.Pasaran_keiekorkecil_5050special,
			"pasaran_discasganjil_5050special":     client.Pasaran_discasganjil_5050special,
			"pasaran_discasgenap_5050special":      client.Pasaran_discasgenap_5050special,
			"pasaran_discasbesar_5050special":      client.Pasaran_discasbesar_5050special,
			"pasaran_discaskecil_5050special":      client.Pasaran_discaskecil_5050special,
			"pasaran_disckopganjil_5050special":    client.Pasaran_disckopganjil_5050special,
			"pasaran_disckopgenap_5050special":     client.Pasaran_disckopgenap_5050special,
			"pasaran_disckopbesar_5050special":     client.Pasaran_disckopbesar_5050special,
			"pasaran_disckopkecil_5050special":     client.Pasaran_disckopkecil_5050special,
			"pasaran_disckepalaganjil_5050special": client.Pasaran_disckepalaganjil_5050special,
			"pasaran_disckepalagenap_5050special":  client.Pasaran_disckepalagenap_5050special,
			"pasaran_disckepalabesar_5050special":  client.Pasaran_disckepalabesar_5050special,
			"pasaran_disckepalakecil_5050special":  client.Pasaran_disckepalakecil_5050special,
			"pasaran_discekorganjil_5050special":   client.Pasaran_discekorganjil_5050special,
			"pasaran_discekorgenap_5050special":    client.Pasaran_discekorgenap_5050special,
			"pasaran_discekorbesar_5050special":    client.Pasaran_discekorbesar_5050special,
			"pasaran_discekorkecil_5050special":    client.Pasaran_discekorkecil_5050special,
		}).
		Post(config.Path_url() + "api/savepasaranconf5050special")
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
func Pasaransaveconf5050kombinasi(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(pasaranconfc5050kombinasi)
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
			"idpasaran":                                 client.Idpasaran,
			"idpasarantogel":                            client.Idpasarantogel,
			"page":                                      client.Page,
			"pasaran_minbet_5050kombinasi":              client.Pasaran_minbet_5050kombinasi,
			"pasaran_maxbet_5050kombinasi":              client.Pasaran_maxbet_5050kombinasi,
			"pasaran_limittotal_5050kombinasi":          client.Pasaran_limittotal_5050kombinasi,
			"pasaran_limitglobal_5050kombinasi":         client.Pasaran_limitglobal_5050kombinasi,
			"pasaran_belakangkeimono_5050kombinasi":     client.Pasaran_belakangkeimono_5050kombinasi,
			"pasaran_belakangkeistereo_5050kombinasi":   client.Pasaran_belakangkeistereo_5050kombinasi,
			"pasaran_belakangkeikembang_5050kombinasi":  client.Pasaran_belakangkeikembang_5050kombinasi,
			"pasaran_belakangkeikempis_5050kombinasi":   client.Pasaran_belakangkeikempis_5050kombinasi,
			"pasaran_belakangkeikembar_5050kombinasi":   client.Pasaran_belakangkeikembar_5050kombinasi,
			"pasaran_tengahkeimono_5050kombinasi":       client.Pasaran_tengahkeimono_5050kombinasi,
			"pasaran_tengahkeistereo_5050kombinasi":     client.Pasaran_tengahkeistereo_5050kombinasi,
			"pasaran_tengahkeikembang_5050kombinasi":    client.Pasaran_tengahkeikembang_5050kombinasi,
			"pasaran_tengahkeikempis_5050kombinasi":     client.Pasaran_tengahkeikempis_5050kombinasi,
			"pasaran_tengahkeikembar_5050kombinasi":     client.Pasaran_tengahkeikembar_5050kombinasi,
			"pasaran_depankeimono_5050kombinasi":        client.Pasaran_depankeimono_5050kombinasi,
			"pasaran_depankeistereo_5050kombinasi":      client.Pasaran_depankeistereo_5050kombinasi,
			"pasaran_depankeikembang_5050kombinasi":     client.Pasaran_depankeikembang_5050kombinasi,
			"pasaran_depankeikempis_5050kombinasi":      client.Pasaran_depankeikempis_5050kombinasi,
			"pasaran_depankeikembar_5050kombinasi":      client.Pasaran_depankeikembar_5050kombinasi,
			"pasaran_belakangdiscmono_5050kombinasi":    client.Pasaran_belakangdiscmono_5050kombinasi,
			"pasaran_belakangdiscstereo_5050kombinasi":  client.Pasaran_belakangdiscstereo_5050kombinasi,
			"pasaran_belakangdisckembang_5050kombinasi": client.Pasaran_belakangdisckembang_5050kombinasi,
			"pasaran_belakangdisckempis_5050kombinasi":  client.Pasaran_belakangdisckempis_5050kombinasi,
			"pasaran_belakangdisckembar_5050kombinasi":  client.Pasaran_belakangdisckembar_5050kombinasi,
			"pasaran_tengahdiscmono_5050kombinasi":      client.Pasaran_tengahdiscmono_5050kombinasi,
			"pasaran_tengahdiscstereo_5050kombinasi":    client.Pasaran_tengahdiscstereo_5050kombinasi,
			"pasaran_tengahdisckembang_5050kombinasi":   client.Pasaran_tengahdisckembang_5050kombinasi,
			"pasaran_tengahdisckempis_5050kombinasi":    client.Pasaran_tengahdisckempis_5050kombinasi,
			"pasaran_tengahdisckembar_5050kombinasi":    client.Pasaran_tengahdisckembar_5050kombinasi,
			"pasaran_depandiscmono_5050kombinasi":       client.Pasaran_depandiscmono_5050kombinasi,
			"pasaran_depandiscstereo_5050kombinasi":     client.Pasaran_depandiscstereo_5050kombinasi,
			"pasaran_depandisckembang_5050kombinasi":    client.Pasaran_depandisckembang_5050kombinasi,
			"pasaran_depandisckempis_5050kombinasi":     client.Pasaran_depandisckempis_5050kombinasi,
			"pasaran_depandisckembar_5050kombinasi":     client.Pasaran_depandisckembar_5050kombinasi,
		}).
		Post(config.Path_url() + "api/savepasaranconf5050kombinasi")
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
func Pasaransaveconfmacaukombinasi(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(pasaranconfmakaukombinasi)
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
			"idpasaran":                     client.Idpasaran,
			"idpasarantogel":                client.Idpasarantogel,
			"page":                          client.Page,
			"pasaran_minbet_kombinasi":      client.Pasaran_minbet_kombinasi,
			"pasaran_maxbet_kombinasi":      client.Pasaran_maxbet_kombinasi,
			"pasaran_limittotal_kombinasi":  client.Pasaran_limittotal_kombinasi,
			"pasaran_limitglobal_kombinasi": client.Pasaran_limitglobal_kombinasi,
			"pasaran_win_kombinasi":         client.Pasaran_win_kombinasi,
			"pasaran_disc_kombinasi":        client.Pasaran_disc_kombinasi,
		}).
		Post(config.Path_url() + "api/savepasaranconfmacaukombinasi")
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
func Pasaransaveconfdasar(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(pasaranconfdasar)
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
			"idpasaran":                 client.Idpasaran,
			"idpasarantogel":            client.Idpasarantogel,
			"page":                      client.Page,
			"pasaran_minbet_dasar":      client.Pasaran_minbet_dasar,
			"pasaran_maxbet_dasar":      client.Pasaran_maxbet_dasar,
			"pasaran_limittotal_dasar":  client.Pasaran_limittotal_dasar,
			"pasaran_limitglobal_dasar": client.Pasaran_limitglobal_dasar,
			"pasaran_keibesar_dasar":    client.Pasaran_keibesar_dasar,
			"pasaran_keikecil_dasar":    client.Pasaran_keikecil_dasar,
			"pasaran_keigenap_dasar":    client.Pasaran_keigenap_dasar,
			"pasaran_keiganjil_dasar":   client.Pasaran_keiganjil_dasar,
			"pasaran_discbesar_dasar":   client.Pasaran_discbesar_dasar,
			"pasaran_disckecil_dasar":   client.Pasaran_disckecil_dasar,
			"pasaran_discgenap_dasar":   client.Pasaran_discgenap_dasar,
			"pasaran_discganjil_dasar":  client.Pasaran_discganjil_dasar,
		}).
		Post(config.Path_url() + "api/savepasaranconfdasar")
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
func Pasaransaveconfshio(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(pasaranconfshio)
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
			"idpasaran":                client.Idpasaran,
			"idpasarantogel":           client.Idpasarantogel,
			"page":                     client.Page,
			"pasaran_minbet_shio":      client.Pasaran_minbet_shio,
			"pasaran_maxbet_shio":      client.Pasaran_maxbet_shio,
			"pasaran_limittotal_shio":  client.Pasaran_limittotal_shio,
			"pasaran_limitglobal_shio": client.Pasaran_limitglobal_shio,
			"pasaran_shioyear_shio":    client.Pasaran_shioyear_shio,
			"pasaran_disc_shio":        client.Pasaran_disc_shio,
			"pasaran_win_shio":         client.Pasaran_win_shio,
		}).
		Post(config.Path_url() + "api/savepasaranconfshio")
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
